package docker_helper

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os/exec"
	"strings"
)

var IMAGE = "alpine"

var _ = Describe("Docker Helper", func() {
	Context("docker execution", Label("docker"), func() {
		It("Executes a simple program", func() {
			err, stdout, stderr := Run("echo 'Hello World!'", IMAGE, map[string]interface{}{})
			Expect(stderr).To(Equal(""))
			Expect(err).To(BeNil())
			Expect(stdout).To(MatchRegexp("Hello World!"))
			Expect(stderr).To(Equal(""))
		})

		It("reports unsuccessful execution correctly", func() {
			err, stdout, stderr := Run("cat not_there", IMAGE, map[string]interface{}{DockerOptionRetries: 0})
			Expect(err).NotTo(BeNil())
			Expect(stdout).To(Equal(""))
			Expect(stderr).To(MatchRegexp("can't open"))
		})
	})

	Context("indirect docker execution - for when you need more detailed control of the exec.Command", Label("docker"), func() {
		It("Executes a simple program - gathers combined output", func() {
			var output string
			err := RunIndirect("echo 'Hello World!", IMAGE, map[string]interface{}{}, func(cmd *exec.Cmd) error {
				std, e := cmd.CombinedOutput()
				output = string(std)
				return e
			})
			Expect(err).To(BeNil())
			Expect(output).To(MatchRegexp("Hello World!"))
		})

		It("Executes a simple program - separate stderr and stdout", func() {
			var stdout string
			var stderr string
			err := RunIndirect("echo 'Hello World!'", IMAGE, map[string]interface{}{}, func(cmd *exec.Cmd) error {
				std, e := cmd.Output()
				stdout = string(std)
				if e != nil {
					stderr = e.Error()
				}
				return e
			})
			Expect(err).To(BeNil())
			Expect(stdout).To(MatchRegexp("Hello World!"))
			Expect(stderr).To(Equal(""))
		})

		It("reports unsuccessful execution correctly", func() {
			var stdout string
			var stderr string
			err := RunIndirect("cat not_there", IMAGE, map[string]interface{}{}, func(cmd *exec.Cmd) error {
				std, e := cmd.Output()
				stdout = string(std)
				fmt.Println(stdout)
				if e != nil {
					if err, ok := e.(*exec.ExitError); ok {
						stderr = string(err.Stderr)
					} else {
						stderr = e.Error()
					}
				}
				return e
			})
			Expect(err).NotTo(BeNil())
			Expect(stdout).To(Equal(""))
			Expect(stderr).To(MatchRegexp("can't open"))
		})

		It("retries the expected number of times", func() {
			var executionCount int
			RunIndirect("cat not_there", IMAGE, map[string]interface{}{DockerOptionRetries: 4}, func(cmd *exec.Cmd) error {
				_, e := cmd.Output()
				executionCount += 1
				return e
			})
			Expect(executionCount).To(Equal(5))
		})
	})

	Context("formatOptions", func() {
		It("adds the docker socket when requested", func() {
			var args []string
			RunIndirect("ls", IMAGE, map[string]interface{}{DockerOptionDockerSocket: true}, func(cmd *exec.Cmd) error {
				// Don't execute, only capturing the args
				args = cmd.Args
				return nil
			})
			formattedOptions := strings.Join(args, " ")
			Expect(formattedOptions).To(ContainSubstring("-v /var/run/docker.sock:/var/run/docker.sock"))
		})

		It("doesn't add the docker socket when not requested", func() {
			var args []string
			RunIndirect("ls", IMAGE, map[string]interface{}{}, func(cmd *exec.Cmd) error {
				// Don't execute, only capturing the args
				args = cmd.Args
				return nil
			})
			formattedOptions := strings.Join(args, " ")
			Expect(formattedOptions).NotTo(ContainSubstring("-v /var/run/docker.sock:/var/run/docker.sock"))
		})
		It("allows docker sock to be overridden -v", func() {
			options := map[string]interface{}{
				"v": map[string]string{
					"$(which docker)":      "/bin/docker",
					"/var/run/docker.sock": "moofy",
				},
			}

			var args []string
			RunIndirect("ls", IMAGE, options, func(cmd *exec.Cmd) error {
				// Don't execute, only capturing the args
				args = cmd.Args
				return nil
			})
			formattedOptions := strings.Join(args, " ")
			Expect(formattedOptions).To(ContainSubstring("-v $(which docker):/bin/docker"))
			Expect(formattedOptions).To(ContainSubstring("-v /var/run/docker.sock:moofy"))
			Expect(formattedOptions).NotTo(ContainSubstring("-v /var/run/docker.sock:/var/run/docker.sock"))
		})
		It("allows docker sock to be overridden --volume", func() {
			options := map[string]interface{}{
				"volume": map[string]string{
					"$(which docker)":      "/bin/docker",
					"/var/run/docker.sock": "moofy",
				},
			}

			var args []string
			RunIndirect("ls", IMAGE, options, func(cmd *exec.Cmd) error {
				// Don't execute, only capturing the args
				args = cmd.Args
				return nil
			})
			formattedOptions := strings.Join(args, " ")
			Expect(formattedOptions).To(ContainSubstring("-volume $(which docker):/bin/docker"))
			Expect(formattedOptions).To(ContainSubstring("-volume /var/run/docker.sock:moofy"))
			Expect(formattedOptions).NotTo(ContainSubstring("-volume /var/run/docker.sock:/var/run/docker.sock"))
		})

		It("automatically adds --rm", func() {
			options := map[string]interface{}{
				"test": "value",
			}
			formattedOptions := strings.Join(formatOptions(options), " ")
			Expect(formattedOptions).To(ContainSubstring("--rm"))
		})

		It("doesn't add rm if no_remove specified adds --rm", func() {
			options := map[string]interface{}{
				"test":               "value",
				DockerOptionNoRemove: true,
			}
			formattedOptions := strings.Join(formatOptions(options), " ")
			Expect(formattedOptions).NotTo(ContainSubstring("--rm"))
		})

		It("retries are not included in command line options", func() {
			options := map[string]interface{}{
				"test":              "value",
				DockerOptionRetries: 50,
			}
			formattedOptions := strings.Join(formatOptions(options), " ")
			Expect(formattedOptions).NotTo(ContainSubstring(DockerOptionRetries))
		})

		It("correctly formats string options", func() {
			options := map[string]interface{}{
				"test": "value",
			}
			formattedOptions := strings.Join(formatOptions(options), " ")
			Expect(formattedOptions).To(ContainSubstring("--test=value"))
		})

		It("correctly formats string slice options", func() {
			options := map[string]interface{}{
				"test": []string{"value1", "value2"},
			}
			formattedOptions := strings.Join(formatOptions(options), " ")
			Expect(formattedOptions).To(ContainSubstring("--test value1=value1"))
			Expect(formattedOptions).To(ContainSubstring("--test value2=value2"))
		})

		It("correctly formats hash options", func() {
			options := map[string]interface{}{
				"e": map[string]string{
					"POSTGRES_USER":     "common_ingress",
					"POSTGRES_PASSWORD": "simple_pw_for_local_database",
				},
			}
			formattedOptions := strings.Join(formatOptions(options), " ")
			Expect(formattedOptions).To(ContainSubstring("-e POSTGRES_USER=common_ingress"))
			Expect(formattedOptions).To(ContainSubstring("-e POSTGRES_PASSWORD=simple_pw_for_local_database"))
		})

		It("correctly formats string slice link options", func() {
			options := map[string]interface{}{
				"link": []string{"pg", "redis"},
			}
			formattedOptions := strings.Join(formatOptions(options), " ")
			Expect(formattedOptions).To(ContainSubstring("--link pg:pg"))
			Expect(formattedOptions).To(ContainSubstring("--link redis:redis"))
		})

		It("correctly formats hash link options", func() {
			options := map[string]interface{}{
				"link": map[string]string{"pg": "my_pg"},
			}
			formattedOptions := strings.Join(formatOptions(options), " ")
			Expect(formattedOptions).To(ContainSubstring("--link pg:my_pg"))
		})

		It("correctly formats hash volume options", func() {
			options := map[string]interface{}{
				"v": map[string]string{"$(which docker)": "/bin/docker"},
			}
			formattedOptions := strings.Join(formatOptions(options), " ")
			Expect(formattedOptions).To(ContainSubstring("-v $(which docker):/bin/docker"))
		})
	})
})

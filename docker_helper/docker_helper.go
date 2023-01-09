package docker_helper

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/koanhealth/gotools/convert"
	"github.com/koanhealth/gotools/fixgo"
	"github.com/koanhealth/gotools/slices"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	DockerOptionNoRemove = "no_remove"
	DockerOptionRetries  = "retries"
)

var (
	specialCases = slices.StringSlice{DockerOptionNoRemove}
)

type TempDirCallback func(path string)
type DockerCallback func(cmd *exec.Cmd) error

func UsingTempDir(block TempDirCallback) {
	path := filepath.Join(dockerTempDir(), uuid.New().String())
	os.MkdirAll(path, os.ModePerm)
	block(path)
	os.RemoveAll(path)
}

func EnsureFileAccessible(path string) (string, error) {
	if strings.Index(path, dockerExchange()) == 0 {
		return path, nil
	}
	sharePath := filepath.Join(shareRoot(), path)
	shareDir := filepath.Dir(sharePath)

	os.MkdirAll(shareDir, os.ModePerm)
	source, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer source.Close()

	share, err := os.Create(sharePath)
	if err != nil {
		return "", err
	}
	defer share.Close()

	share.ReadFrom(source)
	return sharePath, nil
}

func EnsurePathAccessible(path string) string {
	if strings.Index(path, dockerExchange()) == 0 {
		return path
	}
	return filepath.Join(shareRoot(), path)
}

// Run a command in a docker container created for the purpose. Pass DockerOptionNoRemove: true in the options to preserve the
// container after the command completes.
func Run(cmd, image string, options DockerOptions) (err error, stdout, stderr string) {
	err = RunIndirect(cmd, image, options, func(cmd *exec.Cmd) error {
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
	return
}

// RunIndirect a command in a docker container created for the purpose. Pass DockerOptionNoRemove: true in the options to preserve the
// container after the command completes.
// The resulting Command is yielded to the caller in the callback function. The called is responsible for actually
// executing the command, generally with cmd.CombinedOutput() which invokes the process, waits for the return and
// provides stderr & stdout in one byte stream. See os/exec Command for more options.
func RunIndirect(cmd, image string, options DockerOptions, block DockerCallback) error {

	var retries int64
	if r, found := options[DockerOptionRetries]; found {
		retries = convert.ToInt(r)
		delete(options, DockerOptionRetries)
	} else {
		retries = 2
	}

	mergeDefaultVolumes(options)
	runOptions := formatDockerRunOptions(cmd, image, options)
	log.Printf("Docker RUN: %s", strings.Join(runOptions, " "))
	return invokeCmd(runOptions, retries, block)
}

// always pass the docker socket to the invoked container
// the container needs to install the docker cli to use it (and can just ignore this if it doesn't)
// Allows docker containers to invoke docker containers (it's turtles all the way down)
// For now hard coding to the known path of the host (where the daemon runs and to whom all paths must make sense)
func defaultVolumeMappings() map[string]string {
	return map[string]string{"/var/run/docker.sock": "/var/run/docker.sock"}
}

func mergeDefaultVolumes(options map[string]interface{}) {
	var volumeOptions map[string]string

	if v, found := options["v"]; found {
		switch volumes := v.(type) {
		case map[string]string:
			volumeOptions = volumes
		}
	}
	if volumeOptions == nil {
		volumeOptions = make(map[string]string)
		options["v"] = volumeOptions
	}
	for k, v := range defaultVolumeMappings() {
		if _, found := volumeOptions[k]; !found {
			volumeOptions[k] = v
		}
	}
}
func formatDockerRunOptions(cmd, imageName string, options DockerOptions) (runOptions []string) {
	runOptions = []string{
		"run",
		"-v", fmt.Sprintf("%s:%s", dockerExchange(), dockerExchange()),
		"-e", fmt.Sprintf("DOCKER_EXCHANGE=%s", dockerExchange()),
	}
	runOptions = append(runOptions, formatOptions(options)...)
	runOptions = append(runOptions, imageName)
	runOptions = append(runOptions, strings.Split(cmd, " ")...)

	return
}

func dockerExchange() string {
	if evar, found := os.LookupEnv("DOCKER_EXCHANGE"); found && len(evar) > 0 {
		return evar
	} else {
		return makeDockerExchange()
	}
}

func makeDockerExchange() string {
	root, _ := os.Getwd()
	path := filepath.Join(root, "tmp", "docker_exchange")
	os.MkdirAll(path, os.ModePerm)
	return path
}

func dockerTempDir() string {
	path := filepath.Join(dockerExchange(), "tmp")
	os.MkdirAll(path, os.ModePerm)
	return path
}

var _shareRoot string

func shareRoot() string {
	if _shareRoot == "" {
		_shareRoot = filepath.Join(dockerExchange(), fmt.Sprintf("share_%s", uuid.New().String()))
	}
	return _shareRoot
}

func invokeCmd(runOptions []string, retries int64, callback DockerCallback) (err error) {
	dockerPath, _ := exec.LookPath("docker")
	for attemptCount := int64(0); attemptCount <= retries; attemptCount++ {
		cmd := exec.Command(dockerPath, runOptions...)
		if err = callback(cmd); err == nil {
			return
		}
	}
	return
}

func formatOptions(options DockerOptions) (formattedOptions []string) {
	for name, values := range options {
		if specialCases.Contains(name) {
			continue
		}
		formattedOptions = append(formattedOptions, formatOption(name, values)...)
	}

	if noRemove, found := options[DockerOptionNoRemove]; !(found && convert.ToBool(noRemove)) {
		formattedOptions = append(formattedOptions, "--rm")
	}

	return
}

func formatOption(name string, values interface{}) (options []string) {
	if name == "link" || name == "v" || name == "volume" {
		return formatLinkOption(name, values)
	}

	dash := fixgo.Ternary(len(name) > 1, "--", "-")

	if values == nil {
		options = append(options, fmt.Sprintf("%s%s", dash, name))
	} else {
		switch v := values.(type) {
		case map[string]string:
			for ok, ov := range v {
				options = append(options,
					fmt.Sprintf("%s%s", dash, name),
					fmt.Sprintf("%s=%s", ok, ov),
				)
			}
		case string:
			if len(name) > 1 {
				options = append(options, fmt.Sprintf("%s%s=%s", dash, name, v))
			} else {
				options = append(options, fmt.Sprintf("%s%s", dash, name), v)
			}
		case []string:
			for _, ov := range v {
				options = append(options,
					fmt.Sprintf("%s%s", dash, name),
					fmt.Sprintf("%s=%s", ov, ov),
				)
			}
		}
	}
	return
}

func formatLinkOption(name string, values interface{}) (options []string) {
	dash := fixgo.Ternary(len(name) > 1, "--", "-")

	switch v := values.(type) {
	case map[string]string:
		for ok, ov := range v {
			options = append(options,
				fmt.Sprintf("%s%s", dash, name),
				fmt.Sprintf("%s:%s", ok, ov),
			)
		}
	case []string:
		for _, ov := range v {
			options = append(options,
				fmt.Sprintf("%s%s", dash, name),
				fmt.Sprintf("%s:%s", ov, ov),
			)
		}
	}

	return
}

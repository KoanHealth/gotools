FROM golang
COPY . /gotools
RUN go install github.com/onsi/ginkgo/ginkgo@latest
WORKDIR /gotools
COPY .github/compose/docker_entrypoint_test.sh ./
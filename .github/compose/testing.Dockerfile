FROM golang
COPY . /gotools
RUN go install github.com/onsi/ginkgo/v2/ginkgo@latest
WORKDIR /gotools
COPY .github/compose/docker_entrypoint_test.sh ./
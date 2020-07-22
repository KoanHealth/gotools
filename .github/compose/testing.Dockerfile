
FROM golang
COPY . /gotools
RUN go get -u github.com/onsi/ginkgo/ginkgo
WORKDIR /gotools
COPY .github/compose/docker_entrypoint_test.sh ./
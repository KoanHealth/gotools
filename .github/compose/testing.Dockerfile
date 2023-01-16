FROM golang
COPY . /gotools
WORKDIR /gotools
COPY .github/compose/docker_entrypoint_test.sh ./
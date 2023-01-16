#!/usr/bin/env bash
set -euo pipefail

make build && go test ./... -ginkgo.label-filter='!docker' -test.v
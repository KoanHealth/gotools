#!/usr/bin/env bash
set -euo pipefail

make build && ginkgo -r --label-filter='!docker'
#!/bin/bash
set -euo pipefail

go build .
./tritou-proxy

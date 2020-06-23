#!/bin/bash

set -e -o pipefail

if [ "$DISABLE_LINTER" == "true" ]
then
  exit 0
fi

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

if ! [ -x "$(command -v golangci-lint)" ]; then
	echo "Installing GolangCI-Lint"
	${DIR}/install_golint.sh -b $GOPATH/bin v1.27.0
fi

export GOGC=10 GO111MODULE=on
golangci-lint run \
	--no-config \
  -E gofmt \
  -E goimports \
  -E gosec \
  -E interfacer \
	-E misspell \
	-E unconvert \
  -E unparam \
  -E bodyclose \
  -E dupl \
  -E asciicheck \
  -E dogsled \
  -E goconst \
  --timeout 15m \
  --verbose \
  --build-tags build

# -E gocyclo \
# -E nestif \
# -E gocritic \
# -E golint \
# -E godox \
# -E funlen \
# -E goerr113 \

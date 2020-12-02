#!/bin/bash

set -e -o pipefail

if [ "$DISABLE_LINTER" == "true" ]
then
  exit 0
fi

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

if ! [ -x "$(command -v golangci-lint)" ]; then
	echo "Installing GolangCI-Lint"
	${DIR}/install_golint.sh -b $GOPATH/bin v1.31.0
fi

export GOGC=10 GO111MODULE=on
golangci-lint run \
  --timeout 15m \
  --verbose \
  --build-tags build \
  --skip-dirs pkg/client \
  --skip-files pkg/apis/core/v4beta1/zz_generated.deepcopy.go \
  --skip-files pkg/apis/jenkins.io/v1/zz_generated.deepcopy.go

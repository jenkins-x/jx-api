
#!/usr/bin/env bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

#CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ./cmd/code-generator)}
GENERATOR_VERSION=v0.20.2
(
  # To support running this script from anywhere, we have to first cd into this directory
  # so we can install the tools.
  cd "$(dirname "${0}")"
  go get k8s.io/code-generator/cmd/{defaulter-gen,client-gen,lister-gen,informer-gen,deepcopy-gen}@$GENERATOR_VERSION
)

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
rm -rf "${SCRIPT_ROOT}"/pkg/client
# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.
bash hack/generate-groups.sh all \
  github.com/jenkins-x/jx-api/v4/pkg/client github.com/jenkins-x/jx-api/v4/pkg/apis \
  jenkins.io:v1 \
  --output-base "$(dirname "${BASH_SOURCE[0]}")/../../../.." \
  --go-header-file "${SCRIPT_ROOT}"/hack/custom-boilerplate.go.txt

cp -R "${SCRIPT_ROOT}"/v4/pkg/client/ "${SCRIPT_ROOT}"/pkg/client
cp -R "${SCRIPT_ROOT}"/v4/pkg/apis/jenkins.io/v1/zz_generated.deepcopy.go "${SCRIPT_ROOT}"/pkg/apis/jenkins.io/v1/zz_generated.deepcopy.go

rm -rf "${SCRIPT_ROOT}"/v4

#bash hack/generate-groups.sh all \
#  github.com/jenkins-x/jx-api/v4/pkg/generated/core github.com/jenkins-x/jx-api/v4/pkg/apis \
#  core:v4beta1 \
#  --output-base "$(dirname "${BASH_SOURCE[0]}")/../../../.." \
#  --go-header-file "${SCRIPT_ROOT}"/hack/custom-boilerplate.go.txt
#
#
#cp -R "${SCRIPT_ROOT}"/v4/pkg/generated/ "${SCRIPT_ROOT}"/pkg/generated
#cp -R "${SCRIPT_ROOT}"/v4/pkg/apis/core/v4beta1/zz_generated.deepcopy.go "${SCRIPT_ROOT}"/pkg/apis/core/v4beta1/zz_generated.deepcopy.go
#
#rm -rf "${SCRIPT_ROOT}"/v4

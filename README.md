# jx-api

[![Documentation](https://godoc.org/github.com/jenkins-x/jx-api?status.svg)](https://pkg.go.dev/mod/github.com/jenkins-x/jx-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/jenkins-x/jx-api)](https://goreportcard.com/report/github.com/jenkins-x/jx-api)
[![Releases](https://img.shields.io/github/release-pre/jenkins-x/jx-api.svg)](https://github.com/jenkins-x/jx-api/releases)
[![LICENSE](https://img.shields.io/github/license/jenkins-x/jx-api.svg)](https://github.com/jenkins-x/jx-api/blob/master/LICENSE)
[![Slack Status](https://img.shields.io/badge/slack-join_chat-white.svg?logo=slack&style=social)](https://slack.k8s.io/)

Provides an API for JX, can be used with [Jenkins X Kube Client](https://github.com/jenkins-x/jx-kube-client) to create
a programatic interface

Here's an example which also uses [Jenkins X logging](https://github.com/jenkins-x/jx-logging)

```go
import (
    "github.com/jenkins-x/jx-kube-client/v3/pkg/kubeclient"
    "github.com/jenkins-x/jx-logging/v3/pkg/log"
    "github.com/jenkins-x/jx-api/v3/pkg/client/clientset/versioned"
)

func main() {
    f := kubeclient.NewFactory()
    cfg, err := f.CreateKubeConfig()
    if err != nil {
        log.Logger().Fatalf("failed to get kubernetes config: %v", err)
    }


    jxClient, err := versioned.NewForConfig(cfg)
    if err != nil {
        log.Logger().Fatalf("error building jx client: %v", err)
    }
}
```


See the [other modules available](https://github.com/jenkins-x/jx-cli#plugins)

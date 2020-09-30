# jx-api

Provides an API for JX, can be used with [Jenkins X Kube Client](https://github.com/jenkins-x/jx-kube-client) to create
a programatic interface

Here's an example which also uses [Jenkins X logging](https://github.com/jenkins-x/jx-logging)

```go
import (
    "github.com/jenkins-x/jx-kube-client/v3/pkg/kubeclient"
    "github.com/jenkins-x/jx-logging/v3/pkg/log"
    "github.com/jenkins-x/jx-api/pkg/client/clientset/versioned"
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

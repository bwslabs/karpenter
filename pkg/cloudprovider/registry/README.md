# Cloud Provider Registry
This package enables cloud providers to embed themselves into the Karpenter binary without bundling all cloud providers simultaneously. We use mutually exclusive go build tags to register the cloud provider into the import tree. The default implementation is a neutral "mock" cloud provider that implements no-op behavior.

## Add your cloud provider in this directory:
```
// +build <YOUR_PROVIDER_NAME>
import (
	"github.com/awslabs/karpenter/pkg/cloudprovider/<YOUR_PROVIDER_NAME>"
)

func NewFactory() cloudprovider.Factory {
	return <YOUR_PROVIDER_NAME>.NewFactory()
}
```

## Build your customized binary
```
CLOUD_PROVIDER=<YOUR_PROVIDER_NAME> make apply
```

## Add a negative flag to mock.go
```
// +build !<YOUR_PROVIDER_NAME>
```

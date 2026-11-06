# PEIRAVELA Open Go SDK

The Apache-2.0 Go SDK exposes the public SUT adapter SPI so an external
repository can implement and register a System-Under-Test adapter without
depending on the AGPL-3.0 core.

```go
import (
	"context"

	peirasut "github.com/axisvela/peiravela-open/sdk/go/v3/sut"
)

type adapter struct{}

func (adapter) ID() string { return "example" }

func (adapter) Capabilities() peirasut.Capabilities {
	return peirasut.Capabilities{Mode: peirasut.Experiment}
}

func (adapter) Configure(context.Context, peirasut.Config) (peirasut.Session, error) {
	return nil, nil
}

func register() error {
	return peirasut.RegisterAdapter(peirasut.NewRegistry(), adapter{})
}
```

The module is `github.com/axisvela/peiravela-open/sdk/go/v3`. It has no
third-party dependencies. The control plane still owns execution and evidence;
adapters only expose identity, capabilities, configuration, run, observe, and
close behavior.

## Validation

```sh
cd sdk/go
go vet ./...
go test ./...
```

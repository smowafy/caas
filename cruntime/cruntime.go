package cruntime

import(
  "context"
  "github.com/containerd/containerd"
  "github.com/containerd/containerd/namespaces"
)

const ContainerdSocket string = "/run/containerd/containerd.sock"
const NS string = "NS0"

type CRuntime struct {
  ctx context.Context
  client *containerd.Client
}

func SetupContainerd() (runtime CRuntime, err error) {
  client, err := containerd.New(ContainerdSocket)

  if err != nil {
    return CRuntime{}, err
  }

  ctx := namespaces.WithNamespace(context.Background(), NS)

  return CRuntime { ctx: ctx, client: client }, nil
}

func(cruntime *CRuntime) ListContainers() ([]containerd.Container, error) {
  containers, err := cruntime.client.Containers(cruntime.ctx)

  return containers, err
}

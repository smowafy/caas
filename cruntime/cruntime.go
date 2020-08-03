package cruntime

import(
  "log"
  "context"
  "github.com/containerd/containerd"
  "github.com/containerd/containerd/namespaces"
  "github.com/containerd/containerd/oci"
)

const ContainerdSocket string = "/run/containerd/containerd.sock"
const NS string = "NS0"
const SnapshotId string = "alpine-snapshot"

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

func(cruntime *CRuntime) CreateContainer(id string) (containerd.Container, error) {
  var container containerd.Container
  // TODO: see if we need find or create kinda thing. In this case it can be
  // from client.ImageService()
  image, err := cruntime.client.Pull(
    cruntime.ctx,
    "docker.io/library/alpine:latest",
    containerd.WithPullUnpack,
  )

  container, err = cruntime.client.NewContainer(
    cruntime.ctx,
    id,
    cruntime.findOrCreateSnapshot(SnapshotId, image),
    containerd.WithNewSpec(oci.WithImageConfig(image)),
  )

  return container, err
}

func(cruntime *CRuntime) DeleteContainer(id string) error {
  possibleContainer, err := cruntime.findContainerById(id)

  if err != nil {
    return err
  }

  if possibleContainer == nil {
    return nil
  }

  log.Printf("[DeleteContainer] containers found with ID:\n%v\n", possibleContainer)

  return possibleContainer.Delete(cruntime.ctx, containerd.WithSnapshotCleanup)
}

func (cruntime *CRuntime) findOrCreateSnapshot(id string, image containerd.Image) containerd.NewContainerOpts {
  // returns the default (for Linux it's `overlayfs`)
  snapshotter := cruntime.client.SnapshotService("")

  info, err := snapshotter.Stat(cruntime.ctx, id)

  if err != nil || info.Name != id {
    return containerd.WithNewSnapshot(id, image)
  }

  return containerd.WithSnapshot(id)
}

func (cruntime *CRuntime) findContainerById(id string) (containerd.Container, error) {
  var resContainer containerd.Container

  containers, err := cruntime.ListContainers()

  if err != nil {
    return resContainer, err
  }

  var lastError error

  for _, container := range containers {
    info, err := container.Info(cruntime.ctx)

    if err != nil {
      lastError = err
      continue
    }

    log.Printf("current container ID: %s\n", info.ID)

    if info.ID == id {
      return container, nil
    }
  }

  return resContainer, lastError
}

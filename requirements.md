# Blueprint

## What is this project about?

A way to convert a machine (or eventually a cluster of machines) into host(s)
for containers that can be operated by different entities, and provides the
interface for managing these containers.

Think cloud providers but containers instead of VMs.

## Very high level components

#### Daemon
The main component, runs on the machine(s).

Contains all the heavy lifting of container and cluster management.

#### Clients
Interfaces with a "cluster" with the daemon running on the machines.

Can be a CLI, can be a web app, can be both.


### Daemon
The daemon plays these roles:

- Manages the containers on the machine
  - Container operations and management
  - Resource accounting and usage calculation of the containers
  - Coordination with other machines in the cluster

##### Main dependencies
- [containerd](https://containerd.io) for container management.
- (Not yet in place but a potential dependency) [etcd](https://etcd.io) for when
  the cluster part kicks in.
- (Not yet in place) [gRPC](https://grpc.io) for interfacing with clients.



## Very high level milestones

#### Daemon

1. Build CRUD operations for a single type of containers (currently Alpine Linux containers)

2. Build the capability to get shell access to a container

3. Add more options for the container type (preset images for example)

4. Access management (accounts, authorization, etc.)

5. Build resource accounting and enforcement

6. Build container lifecycle, failure recovery, backup etc.

7. Build cluster support

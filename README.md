# The krl
krl : kubernetes resource list.

this tool will list all resource from kubernetes in the table format.

Now release: v0.0.1
## Usage:

if not given target ns, it will show target resource from all namespace.

kubectl rl rs 
```
+===================================+=============+===========+============+
|name                               |  namespace  | api-group |  api-kind  |
+===================================+=============+===========+============+
|drone-runner-8698b9d977            |   drones    |   apps    | ReplicaSet |
+-----------------------------------+-------------+-----------+------------+
|drone-server-96bffc8ff             |   drones    |   apps    | ReplicaSet |
+-----------------------------------+-------------+-----------+------------+
|calico-kube-controllers-bcc6f659f  | kube-system |   apps    | ReplicaSet |
+-----------------------------------+-------------+-----------+------------+
|coredns-7f89b7bc75                 | kube-system |   apps    | ReplicaSet |
+-----------------------------------+-------------+-----------+------------+
|grafana-f8cd57fcf                  | monitoring  |   apps    | ReplicaSet |
+-----------------------------------+-------------+-----------+------------+
|kube-state-metrics-587bfd4f97      | monitoring  |   apps    | ReplicaSet |
+-----------------------------------+-------------+-----------+------------+
|prometheus-adapter-69b8496df6      | monitoring  |   apps    | ReplicaSet |
+-----------------------------------+-------------+-----------+------------+
|prometheus-operator-7649c7454f     | monitoring  |   apps    | ReplicaSet |
+-----------------------------------+-------------+-----------+------------+
```

kubectl rl all -n kube-system
```
+=========================================+=============+===========+============+
|name                                     |  namespace  | api-group |  api-kind  |
+=========================================+=============+===========+============+
|calico-kube-controllers-bcc6f659f-6mvjr  | kube-system |           |    Pod     |
+-----------------------------------------+-------------+-----------+------------+
|calico-node-5vfjc                        | kube-system |           |    Pod     |
+-----------------------------------------+-------------+-----------+------------+
|calico-node-6shfz                        | kube-system |           |    Pod     |
+-----------------------------------------+-------------+-----------+------------+
|calico-node-tkfn8                        | kube-system |           |    Pod     |
+-----------------------------------------+-------------+-----------+------------+
|coredns-7f89b7bc75-5vgs4                 | kube-system |           |    Pod     |
+-----------------------------------------+-------------+-----------+------------+
|coredns-7f89b7bc75-6kbq2                 | kube-system |           |    Pod     |
+-----------------------------------------+-------------+-----------+------------+
|etcd-master                              | kube-system |           |    Pod     |
+-----------------------------------------+-------------+-----------+------------+
|kube-apiserver-master                    | kube-system |           |    Pod     |
+-----------------------------------------+-------------+-----------+------------+
|kube-controller-manager-master           | kube-system |           |    Pod     |
+-----------------------------------------+-------------+-----------+------------+
|kube-proxy-62nfn                         | kube-system |           |    Pod     |
+-----------------------------------------+-------------+-----------+------------+
|kube-proxy-bp86v                         | kube-system |           |    Pod     |
+-----------------------------------------+-------------+-----------+------------+
|kube-proxy-gntzv                         | kube-system |           |    Pod     |
+-----------------------------------------+-------------+-----------+------------+
|kube-scheduler-master                    | kube-system |           |    Pod     |
+-----------------------------------------+-------------+-----------+------------+
|kube-dns                                 | kube-system |           |  Service   |
+-----------------------------------------+-------------+-----------+------------+
|kubelet                                  | kube-system |           |  Service   |
+-----------------------------------------+-------------+-----------+------------+
|calico-node                              | kube-system |   apps    | DaemonSet  |
+-----------------------------------------+-------------+-----------+------------+
|kube-proxy                               | kube-system |   apps    | DaemonSet  |
+-----------------------------------------+-------------+-----------+------------+
|calico-kube-controllers                  | kube-system |   apps    | Deployment |
+-----------------------------------------+-------------+-----------+------------+
|coredns                                  | kube-system |   apps    | Deployment |
+-----------------------------------------+-------------+-----------+------------+
|calico-kube-controllers-bcc6f659f        | kube-system |   apps    | ReplicaSet |
+-----------------------------------------+-------------+-----------+------------+
|coredns-7f89b7bc75                       | kube-system |   apps    | ReplicaSet |
+-----------------------------------------+-------------+-----------+------------+
```

kubectl rl deployment -n kube-system
```
+=========================+=============+===========+============+
|name                     |  namespace  | api-group |  api-kind  |
+=========================+=============+===========+============+
|calico-kube-controllers  | kube-system |   apps    | Deployment |
+-------------------------+-------------+-----------+------------+
|coredns                  | kube-system |   apps    | Deployment |
+-------------------------+-------------+-----------+------------+
```

## For build
### drone 
[.drone.yml](.drone.yml)

The krl use drone to test and build.
High availability, Disaster Recovery, Scalabilty.

- Ingress handles every incoming request.
- Ingress is also replicated.
- Ingress will forward request to a Service.
- Service is a load balancer.
- If Node dies, Controller Manager schedule new replicas.
- Etcd tracks cluster state and we take snapshots of it in remote storage.
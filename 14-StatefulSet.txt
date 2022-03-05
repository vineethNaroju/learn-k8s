*** About Me ***

- Database is a stateful set.

Nodejs(stateless) - passthrough for data query / update.
MongoDB(stateful) - update data based on previous state, query data.

Deployment => abstraction of pods for stateless applications.
StatefulSet => deploy stateful apps and replicate pods.

- Both manage pods and replicate pods.

*** Deployment ***

- Identical and interchangable
- Created in random order with random hashes.
- One Service that load balances to any pod.

*** StatefulSet ***

- Can't be randomly deleted / addressed => Pod Identity.
- Sticky identity for each pod, persistent across any re-scheduling.
- Need synchronization - between Leader and Follower pods.
- New Follower clones from previous pod in sequence and will listen from Leader.

- All data will be lost when all pods die !!!!

- Each pod state and data is stored in Persistent Volume.
- Re-scheduled pod is re-attached to a specific Persistent Volume.

- StatefulSet pods are name $(statefulset name)-$(ordinal)
mysql-0, mysql-1, mysql-2
- Leader, Follower, Follower

- Next pod is created only if previous pod is up and running.
- Deletion happens in reverse order starting from last one.

- Each pod has it's DNS endpoint from service.
mysql-0.svc2, mysql-1.svc2

1. Predictable pod name.
2. fixed individual DNS name.

=> IP address changes but name and endpoint stays same.

- Need to configure cloning and data synchronization.
- Make remote storage available.
- Managing and backup.
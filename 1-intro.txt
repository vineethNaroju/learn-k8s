*** About me ***

To manage containerized servers in physical / cloud / hybrid

*** Concepts ***

A Pod can have many containers. Generally a pod has just one container.
A Pod can die and has its own IP address.
That's why we have Service and we give config about pods in Deployment file.
A Service has permanent IP and services stays infront of Pod.
A Virtual Network is maintained across all Worker Nodes and Master Nodes.

*** Architecure ***

Worker Node 
It's a machine. Each Worker Node has 3 process:
1.Container Runtime process.
2.Kubelet process
Creates Nodes, manages CPU, RAM and storage.
3. Kube-proxy
Forwards requests to pods.

Master Node - has 4 components
1. API server process (UI / kubectl, cluster gateway, authentication).
2. Scheduler process (schedules pods on nodes).
3. Controller Manager (Detects cluster state changes and asks Scheduler to do make changes).
4. Etcd (Stores cluster state as key-value store).

Scheduler process asks kubelet of a Node to create a pod on itself.
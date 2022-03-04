*** About me ***

> kubectl cluster-info
> kubectl get namespaces

NAME              STATUS   AGE
default           Active   80m
kube-node-lease   Active   80m
kube-public       Active   80m
kube-system       Active   80m


1. kube-system
Dont modify or create.
System processes.
Master and kubectl processes.

2. kube-public
Publicily accessible data.
ConfigMap containing cluster information.

3. kube-node-lease
Hearbeats of nodes.
Each node has a lease object here.
Determines availability of a node.

4. Default
Obvious.

*** Create Namespace ***

> kubectl create namespace my-namespace 

Create namespace with a config file as,
apiVersion: v1
kind: ConfigMap
metadata:
    name: mongo-configmap
    namespace: my-namespace
data:
    database_url: mongo-service


*** why namespaces ? ***

1. Group resources (database, elastic stack, monitoring, nginx-ingress).

2. To avoid conflicts in deployments of different teams in same cluster.

3. Staging, Development -> resuse components in both environment and share 
Nginx-Controller, Elastic Stack cluster among Blue/Green - different versions.

4. Different teams can access only their namespace and limit resources. To 
limit CPU, RAM, Storage per Namespace.

*** Properties of Namespace ***

1. You can't access most resources from another Namespace.

2. Each namespace must have its own ConfigMap and Secret.

3. Service can be shared across multiple in Namespaces.
ex database_url: SERVICE-NAME.NAMESPACE

4. Some things live globally and you can't isolate them like Volume, node

> kubectl api-resources --namespaced=false
> kubectl get configmaps -n=my-namespace
> kubectl get all -n my-namespace
> kubectl apply apply -f mongo-config.yaml --namespace=my-namespace
> metadata.namespace : my-namespace

"ConfigMap namespace is recommended over kubectl cmd"

*** Active Namespace ***
Change the active namespace with kubens
No default tool available for k8s

> brew install kubectx
> kubens
default ---------> Highlighted
kube-node-lease
kube-public
kube-system
my-namespace

To change active namespace
> kubens my-namespace
*** About me ***

Instead of exposing a service as external and invoking it using ip:port.
I can use ingress and link a domain to above service(as internal).
Ip and port of internal service is not opened.

No nodePort and ClusterIP (default type) in internal service.

*** Example ***

apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
    name: myapp-ingress
spec:
    rules:
    - host: myapp.com -------> Domain address (forward req to internal service)
     http: -----> Protocol of incoming request that gets forwarded to internal service.
        paths: ------> Everything after domain (http://myapp.com/...THIS...)
        - backend:
            serviceName: myapp-internal-service
            servicePort: 8080 ----> Internal service port


Host:
- valid domain address
- map domain name to Node's IP address, the entrypoint.

*** Configure Ingress in Cluster ***

- Need implementation for ingress -> Ingress Controller.
- Evaluates and processes Ingress rules.
- Entrypoint to cluster based on domain and sub-domain.
- Many 3rd-party ones.
- K8s Nginx Ingress Controller.

AWS, GCP, Linode =>
Ext Req -> Cloud LB (provided by Cloud Service) -> Ingress Controller Pod

Bare Metal:
we have to configure some kind of entrypoint, either inside of cluster or outside
as separate server.
https://kubernetes.github.io/ingress-nginx/deploy/baremetal/

External Proxy Server

> minikube addons enable ingress

Enable kubernetes dashboard
> minikube dashboard --url
> kubens kubernetes-dashboard
> kubectl get all

> kubectl apply -f dashboard-ingress.yaml
> kubectl describe ingress dashboard-ingress
> kubectl get ingress

To get external ip address
> minikube ip
192.168.49.2

> sudo vim /etc/hosts
192.168.49.2 dashboard.com

----------------------------------------
wtf is "nginx.ingress.kubernetes.io/rewrite-target: /" ?


host: myapp.com

path: /analytics
backend:

path: /shopping
backend:

or 

host: analytics.myapp.com
http:
    paths:
        backend:
host: shopping.myapp.com
http:
    paths:
        backend:

-------------------------------------------

*** Configure for TLS ***

spec:
    tls:
    - hosts:
        - myapp.com
        secretName: myapp-secret-tls


apiVersion: v1
kind: Secret
metadata:
    name: myapp-secret-tls
    namespace: default -----> same namespace as ingress 
data:
    tls.crt: base64(cert) --------> Mandatory file contents not paths
    tls.key: base64(key) --------> Mandatory file contents not paths
type: kubernetes.io/tls
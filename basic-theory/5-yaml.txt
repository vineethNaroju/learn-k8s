
*** Deployment ***

apiVersion: apps/v1
kind: Deployment
metadata:
    name: nginx-deployment
    labels:
        app: nginx
spec:
    replicas: 2
    selector:
        matchLabels:
            app: nginx
    template:
        metadata:
            labels:
                app: nginx
        spec: ------------> POD BLUEPRINT
            containers:
            - name: nginx
              image: nginx:1.16
              ports:
              - containerPort: 8080


*** Service ***

apiVersion: v1
kind: Service
metadata:
    name: nginx-service
spec:
    selector:
        app: nginx
    ports:
    - protocol: TCP
      port: 80 ------------------> Service Port
      targetPort: 8080  ----------------> container port of pod

1. Metadata (name, labels)
2. Spec (replicas, selector, port) - Attributes specific to the "kind"
3. Status is maintained in etcd.

We can store these files in code.

Deployment -> Replicaset -> Pod -> Container

*** Labels & Selectors ***

Template contains "labels" and Spec contains "selectors"

Service spec selector connects with Deployment metadata-labels and with
spec-template-metadata-labels based on spec-selector-matchLabels

-------------------------------------------

kubectl get services
kubectl describe service nginx-service 

Now check Endpoint in it

To find ip address of pods
kubectl get pod -o wide

To get status of deployment
kubectl get deployment nginx-deployment -o yaml > nginx-deployment-result.yaml
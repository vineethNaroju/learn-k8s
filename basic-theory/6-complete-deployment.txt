Mongo-Express(Web App) -> MongoDB

2 Deployment / Pod
2 Service
1 ConfigMap
1 Secret

1. MongoDB pod -> internal service (no external requests)
2. Mongo Express -> external service
3. ConfigMap[DB Url]
4. Secret[DB User, DB Pwd]
5. Mongo Express can be called using http://NODE-IP-ADDRESS:PORT

----------------------------------------------

kubectl apply -f mongodb-deployment.yaml

mongodb listens on 27017

Deploy configs, secrets and then deployment, services.

To make external service
- type -> LoadBalancer
- nodePort for external IP addresss 30000 - 32767

---------------------------------------------------------

In minikube we don't get external ip address

minikube service mongo-express-service

minikube service list
minikube service --url example-service
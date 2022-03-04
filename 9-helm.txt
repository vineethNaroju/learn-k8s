*** About Me ***
- Package manager for k8s
- To package YAML and distribute them in public and private repos.

? Helm , Helm Charts, Tiller ?

*** Example *** 
Say we want Elastic stack for Logging,
we need StatefulSet, ConfigMap, Secret, K8s User with permissions,
Services.

Bundle of YAML files => Helm Chart and push to Helm repo.
Download and use existing ones.

We can reuse same configuration.

We have public and private repos.

*** Templating Engine *** 

1. Define common blueprint.

2. Dynamic values replaced by placeholders.
name: {{ .Values.name }}
values.yaml
name: my-app
container:
    name: myapp-container
    image: myapp-image
    port: 9001

3. Package them as my own helm chart to deploy in different environment
clusters.

*** Directory Structure ***
mychar/
    chart.yaml
    values.yaml
    charts/ ----> chart dependencies
    templates/ -------> template files

helm install <chartname>
helm install --values=my-values.yaml <chartname>
helm install --set version=2.0.0

*** Release Management ***
v2 - Helm CLI (Client) and Tiller (Server)
helm install <chartname> cli -> Tiller (executes in k8s )

Tiller keep tracks of all chart executions.

Tiller has too much power - create , update and delete stuff.
- Security Issue

v3 - Hard to use but no more Tiller :)
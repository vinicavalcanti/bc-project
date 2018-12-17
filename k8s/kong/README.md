# Kong Ingress

## Deploy Kong as ingress controller on Kubernetes

##### Before to deploy Kong Ingress on Kubernetes do you need to provide the PostgreSQL with Amazon RDS. The manifest files of Terraform are located on folder `rds-database`

### Install

```
kubectl create -f deploy.yaml

```

The command above will create this resources on your Kubernetes:

```
namespace/kong created
customresourcedefinition.apiextensions.k8s.io/kongplugins.configuration.konghq.com created
customresourcedefinition.apiextensions.k8s.io/kongconsumers.configuration.konghq.com created
customresourcedefinition.apiextensions.k8s.io/kongcredentials.configuration.konghq.com created
customresourcedefinition.apiextensions.k8s.io/kongingresses.configuration.konghq.com created
serviceaccount/kong-serviceaccount created
clusterrole.rbac.authorization.k8s.io/kong-ingress-clusterrole created
role.rbac.authorization.k8s.io/kong-ingress-role created
rolebinding.rbac.authorization.k8s.io/kong-ingress-role-nisa-binding created
clusterrolebinding.rbac.authorization.k8s.io/kong-ingress-clusterrole-nisa-binding created
service/kong-ingress-controller created
deployment.extensions/kong-ingress-controller created
service/kong-proxy created
deployment.extensions/kong created
job.batch/kong-migrations created
```

Now, you can check if your deployment was successfull:

```
$ kubectl get pods -n kong

NAME                                      READY     STATUS      RESTARTS   AGE
kong-bd6f8f695-2vfv5                      1/1       Running     1          10s
kong-ingress-controller-5f79fffdd-dstx8   2/2       Running     2          10s
kong-migrations-tv5s4                     0/1       Completed   0          10s
```


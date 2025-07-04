# openshift-gin-api

```
oc expose deployment gin-api --port=8080 --name=gin-api

oc expose svc/gin-api

oc get route gin-api

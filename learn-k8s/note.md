### nslookup
```bash
nslookup google-service.default.svc.cluster.local
```

- Set up local ingress on minikube: https://kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/

- Install metric-server
```bash
minikube addons enable metrics-server
```

### RBAC
- Export serviceaccount into Pod 
```bash
export SERVICEACCOUNT=/var/run/secrets/kubernetes.io/serviceaccount/
export TOKEN=$(cat ${SERVICEACCOUNT}/token)
export CACERT=${SERVICEACCOUNT}/ca.crt

curl --cacert ${CACERT} --header "Authorization: Bearer ${TOKEN}" -X GET https://kubernetes.default.svc/api/v1/namespaces/default/pods
```

- Gen key: 
```bash
openssl genrsa -out developer-be.key 2048
openssl req -new -key developer-be.key -out developer-be.csr -subj "/CN=developer-be/O=staging"
openssl x509 -req -in developer-be.csr -CA /etc/kubernetes/pki/ca.crt -CAkey /etc/kubernetes/pki/ca.key -CAcreateserial -out developer-be.crt -days 365

openssl x509 -req -in developer-be.csr -CA $HOME/.minikube/ca.crt -CAkey $HOME/.minikube/ca.key -CAcreateserial -out developer-be.crt -days 365
```

- Set new user
```bash
kubectl config set-credentials developer-be --client-certificate=developer-be.crt --client-key=developer-be.key
```

- Set context
```bash
kubectl config set-context developer-be-context --cluster=minikube --user=developer-be
```

- Với statefulset thì thường triển khai theo headless service

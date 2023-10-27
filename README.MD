### we have to ensure that the images we will build, will be inside the minikube registery not on our local machine registery, so minikube has its own running docker ..
```cmd
➜ docker-with-golang git:(working_with_k8s_minikube) minikube docker-env
W1027 16:19:18.777975   11884 main.go:291] Unable to resolve the current Docker CLI context "default": context "default": context not found: open C:\Users\gamil\.docker\contexts\meta\37a8eec1ce19687d132fe29051dca629d164e2c4958ba141d5f4133a33f0688f\meta.json: The system cannot find the path specified.
$Env:DOCKER_TLS_VERIFY = "1"
$Env:DOCKER_HOST = "tcp://127.0.0.1:57372"
$Env:DOCKER_CERT_PATH = "C:\Users\gamil\.minikube\certs"
$Env:MINIKUBE_ACTIVE_DOCKERD = "minikube"
# To point your shell to minikube's docker-daemon, run:
# & minikube -p minikube docker-env --shell powershell | Invoke-Expression
➜ docker-with-golang git:(working_with_k8s_minikube)
```

### now we need to open the minikube inside our terminal to execute docker commands agains it 
```cmd
# ==============================> before minikube
➜ docker-with-golang git:(working_with_k8s_minikube) docker images
REPOSITORY                                                       TAG                                                                          IMAGE ID       CREATED         SIZE
golang-api-image                                                 2.0                                                                          50e0a12feabb   4 hours ago     714MB
golang-api-image                                                 1.0                                                                          08accb44db2d   21 hours ago    16.1MB
<none>                                                           <none>                                                                       db785e32b8e3   6 days ago      195MB
<none>                                                           <none>                                                                       1b27590f5b2d   6 days ago      193MB
service-amd64                                                    1.0                                                                          8d29568822d5   8 days ago      9.29MB
fadygamil/auth                                                   ebb20cdbf148b8b8025a12a233220fd9ea15c5eb5e330badde38dfd7b6610640             ebb20cdbf148   8 days ago      259MB
fadygamil/auth                                                   latest                                                                       ebb20cdbf148   8 days ago      259MB
fadygamil/auth                                                   2956b2550e3aee71ff9d53ede396e6ec8b6c352ef5097938a5b2676d65b7571e             2956b2550e3a   8 days ago      259MB
postgres                                                         14                                                                           251b1e989f6e   2 months ago    408MB
gcr.io/k8s-minikube/kicbase                                      v0.0.40                                                                      c6cc01e60919   3 months ago    1.19GB
hubproxy.docker.internal:5555/docker/desktop-kubernetes          kubernetes-v1.27.2-cni-v1.2.0-critools-v1.27.0-cri-dockerd-v0.3.2-1-debian   c763812a4530   4 months ago    418MB
registry.k8s.io/kube-apiserver                                   v1.27.2                                                                      c5b13e4f7806   5 months ago    121MB
registry.k8s.io/kube-controller-manager                          v1.27.2                                                                      ac2b7465ebba   5 months ago    112MB
registry.k8s.io/kube-scheduler                                   v1.27.2                                                                      89e70da428d2   5 months ago    58.4MB
registry.k8s.io/kube-proxy                                       v1.27.2                                                                      b8aa50768fd6   5 months ago    71.1MB
docker/desktop-vpnkit-controller                                 dc331cb22850be0cdd97c84a9cfecaf44a1afb6e                                     556098075b3d   5 months ago    36.2MB
registry.k8s.io/coredns/coredns                                  v1.10.1                                                                      ead0a4a53df8   8 months ago    53.6MB
registry.k8s.io/etcd                                             3.5.7-0                                                                      86b6af7dd652   9 months ago    296MB
registry.k8s.io/pause                                            3.9                                                                          e6f181688397   12 months ago   744kB
k8s.gcr.io/ingress-nginx/controller                              <none>                                                                       ae1a7201ec95   23 months ago   285MB
k8s.gcr.io/ingress-nginx/controller                              <none>                                                                       ac0e4fe3e6b0   2 years ago     279MB
kindest/node                                                     v1.21.1                                                                      65d38077cb24   2 years ago     931MB
docker/desktop-storage-provisioner                               v2.0                                                                         99f89471f470   2 years ago     41.9MB
jettech/kube-webhook-certgen                                     v1.5.1                                                                       a013daf8730d   2 years ago     44.7MB
quay.io/kubernetes-ingress-controller/nginx-ingress-controller   0.30.0                                                                       89ccad40ce8e   3 years ago     323MB

# ==============================> after minikube
➜ docker-with-golang git:(working_with_k8s_minikube) minikube -p minikube docker-env --shell powershell | Invoke-Expression
W1027 16:21:18.316739   12724 main.go:291] Unable to resolve the current Docker CLI context "default": context "default": context not found: open C:\Users\gamil\.docker\contexts\meta\37a8eec1ce19687d132fe29051dca629d164e2c4958ba141d5f4133a33f0688f\meta.json: The system cannot find the path specified.
➜ docker-with-golang git:(working_with_k8s_minikube) docker images
REPOSITORY                                TAG       IMAGE ID       CREATED         SIZE
registry.k8s.io/kube-apiserver            v1.27.4   e7972205b661   3 months ago    121MB
registry.k8s.io/kube-proxy                v1.27.4   6848d7eda034   3 months ago    71.1MB
registry.k8s.io/kube-scheduler            v1.27.4   98ef2570f3cd   3 months ago    58.4MB
registry.k8s.io/kube-controller-manager   v1.27.4   f466468864b7   3 months ago    113MB
registry.k8s.io/coredns/coredns           v1.10.1   ead0a4a53df8   8 months ago    53.6MB
registry.k8s.io/etcd                      3.5.7-0   86b6af7dd652   9 months ago    296MB
registry.k8s.io/pause                     3.9       e6f181688397   12 months ago   744kB
gcr.io/k8s-minikube/storage-provisioner   v5        6e38f40d628d   2 years ago     31.5MB
➜ docker-with-golang git:(working_with_k8s_minikube)
```

### now lets first build the image 
```cmd
docker build -t golang-api:1.0 .
```

### next lets create the deployment to manage our pod
```cmd
➜ docker-with-golang git:(working_with_k8s_minikube) kubectl create -f .\infra\api-depl.yaml
deployment.apps/api-depl created
```

### now lets explore all the created resources within the current namespace 
```cmd
➜ docker-with-golang git:(working_with_k8s_minikube) kubectl get all
NAME                            READY   STATUS         RESTARTS   AGE
pod/api-depl-75466df88b-rmklc   0/1     ErrImagePull   0          7s

NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   134m

NAME                       READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/api-depl   0/1     1            0           7s

NAME                                  DESIRED   CURRENT   READY   AGE
replicaset.apps/api-depl-75466df88b   1         1         0       7s
➜ docker-with-golang git:(working_with_k8s_minikube)
```

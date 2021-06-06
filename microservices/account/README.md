### 总体规划
1. 编写go程序
2. 编译运行go程序
3. 安装docker
4. docker register
5. docker运行go程序
6. 搭建kubernetes集群
7. 搭建ingress
8. golang部署到kubernetes（deployment,service,ingress）

### 编译步骤
1. go build
   ```shell
    GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o accountservice-linux-amd64
   ```
2. create docker images from Dockerfile
   ```Dockerfile
   FROM alpine
   ADD accountservice-linux-amd64 /
   ENTRYPOINT ["/accountservice-linux-amd64"]
   ```
   ```shell
    # 先删除旧镜像
    docker images |grep account | awk '{print $3}' |xargs docker rmi -f
    # 执行编译镜像
    docker build . -t b.io:5000/account
    # 推送镜像
    docker push b.io:5000/account
    # 测试是否可用
    docker run -it --rm b.io:5000/account  /bin/sh
    # 启动 docker 容器
    docker run -d b.io:5000/account
   ```
2. deploy by docker-compose
   ```yaml
   version: '3.3'
   services:
     account:
       restart: always
       image: b.io:5000/account
       ports:
         - 6767:6767
   ```
4. if not 3,deploy by kubernetes
   ```yaml
   # kubernetes.yml
   # deployment
   apiVersion: apps/v1
   kind: Deployment
   metadata:
     name: account-deployment
     namespace: microservices
   spec:
     selector:
       matchLabels:
         app: account
     replicas: 2 # tells deployment to run 2 pods matching the template
     template:
       metadata:
         labels:
           app: account
       spec:
         containers:
           - name: account
             image: b.io:5000/account
   ---
   #service
   apiVersion: v1
   kind: Service
   metadata:
     name: account-service
     namespace: microservices
   spec:
     selector:
       app: account
     ports:
       - name: http
         protocol: TCP
         port: 6767
         targetPort: 6767


   ---
   # ingress
   apiVersion: traefik.containo.us/v1alpha1
   kind: IngressRoute
   metadata:
     name: account-api
     namespace: microservices
   spec:
     entryPoints:
       - web
     routes:
       - match: Host(`account.b.io`)
         kind: Rule
         services:
           - name: account-service
             port: 6767
   ```
   ```
   kubectl apply -f kubernetes.yml
   ```

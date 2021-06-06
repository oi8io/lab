all:kubernetes
build:
	GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o accountservice-linux-amd64
docker-delete-image:
	docker images |grep account | awk '{print $$3}' |xargs -r docker rmi -f
docker-build:build|docker-delete-image
	docker build . -t b.io:5000/account
docker-push:docker-build
	docker push b.io:5000/account

docker:docker-push

#todo don't execute with not exist
kubernetes-delete:
	kubectl delete -f kubernetes.yml
kubernetes-apply:docker|kubernetes-delete
	kubectl apply -f kubernetes.yml
kubernetes:kubernetes-apply
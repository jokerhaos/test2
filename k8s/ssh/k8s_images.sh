images=(
	kube-apiserver:v1.27.1
	kube-controller-manager:v1.27.1
	kube-scheduler:v1.27.1
	kube-proxy:v1.27.1
	pause:3.9
	etcd:3.5.7-0
	coredns:v1.10.1
)

for imageName in ${images[@]}; do
	docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/$imageName
	docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/$imageName k8s.gcr.io/$imageName
	docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/$imageName
done
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kindest/node:v1.13.4
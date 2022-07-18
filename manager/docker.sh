image=scyna/manager:$1

docker build -t ${image} -f Dockerfile .
docker push ${image}
image=scyna/engine:$1

docker build -t ${image} -f Dockerfile .
docker push ${image}
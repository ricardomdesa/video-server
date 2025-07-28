# how to run minikube

minikube start

## how to deploy

Rodar o comando para criar a imagem na pasta raiz:
docker build -t go-api-minikube:v1 .

kubectl apply -f deployment.yaml -n video-str
kubectl apply -f api-key-secret.yaml -n video-str
kubectl apply -f service.yaml -n video-str

### check

kubectl get svc -n video-str
kubectl get pods -n video-str

### port forward
kubectl port-forward service/go-api 8010:80 -n video-str 

curl localhost:8010/ping --header 'X-api-key: <apikey>'


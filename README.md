# microservice-architecture-03
Prometheus. Grafana

Deployment

Create namespace:

    kubectl create namespace monitoring
    kubectl config set-context --current --namespace=monitoring

Install and start helm services:

    helm repo add prometheus-community
    helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
    helm repo add prometheus-community https://prometheus-community.github.io/helm-charts

    helm repo update

    helm install prom prometheus-community/kube-prometheus-stack -f prometheus.yaml --atomic
    helm install nginx ingress-nginx/ingress-nginx -f nginx-ingress.yaml --atomic
    helm install prom-postgres prometheus-community/prometheus-postgres-exporter -f prometheus-postgres.yaml

Port forwarding: 
    
    Grafana (http://localhost:9000):
    
        kubectl port-forward service/prom-grafana 9000:80
    
    Prometheus (http://localhost:9090):

        kubectl port-forward service/prom-kube-prometheus-stack-prometheus 9090

Add crudapp grafana dashboard:

    kubectl apply -f grafana.yaml

Run crudapp:

    helm install crudapp ./crud-chart

Test:

    newman run ./crud-service.postman_collection.json

Delete: 

    helm uninstall crudapp
    helm uninstall prom
    helm uninstall prom-postgres
    helm uninstall nginx

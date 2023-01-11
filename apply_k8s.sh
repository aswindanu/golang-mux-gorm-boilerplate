kubectl apply -f . --validate=false
kubectl port-forward svc/backend 3000:3000 # development
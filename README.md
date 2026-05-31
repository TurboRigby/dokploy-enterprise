# Free Dokploy Enterprise License

A fully featured enterprise license server for Dokploy.

## Installation
go run main.go

## Usage
Point LICENSE_KEY_URL to this server.
You're now enterprise. Yeah. That's all.

## Contributing
PRs welcome I guess, but what are you will change in it? how it returns json? lol

## License
MIT, cuz why not, let's copyright a fucking crack that's 28 lines of code

## Deployment

### Local (recommended, it's 28 lines bro)
```bash
go run main.go
```

### Docker (if you're fancy)
```bash
docker-compose up
```

### Kubernetes (if you're insane)
```bash
kubectl apply -f k8s/deployment.yaml
```

and yes, the k8s manifest has:
- **3 replicas** because returning `{"valid": true}` is serious business
- **HorizontalPodAutoscaler** scaling up to 10 pods in case the load of returning true gets too heavy
- **liveness and readiness probes** to make sure the JSON is still alive and well
- **512Mi memory limit** which is coincidentally exactly what Dokploy uses at idle, for an app that uses 8MB
- a dedicated **namespace** because this crack deserves isolation from lesser workloads

an AI agent generated this manifest with the same energy as deploying a banking system. we kept it as-is for comedic and production purposes.

### Dokploy (poetic justice)
idk figure it out, their UI eats 500mb of ram anyway
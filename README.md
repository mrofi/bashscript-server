# bashscript-server

A lightweight Go server to serve and run Bash scripts via HTTP.  
Supports piping scripts directly with `curl`/`wget` and passing arguments.  
Configurable script directory and port via environment variables. Ideal as a base Docker image or Go module for script distribution.

---

## Features

- Serve Bash scripts over HTTP.
- Pass arguments to scripts: `curl | bash -s -- arg1 arg2`.
- Configurable `SCRIPT_DIR` and `PORT`.
- Can be used as a base Docker image or Go module.
- Optional Kubernetes deployment.

---

## 🔹 Installation

### Using Go module
```bash
go install github.com/mrofi/bashscript-server@latest
````

### Using Docker

#### Build with bundled scripts

```bash
docker build -t bashscript-server .
docker run -p 8080:8080 bashscript-server
```

#### Use as base image

```dockerfile
FROM ghcr.io/mrofi/bashscript-server:latest
COPY scripts /app/scripts
```

#### Mount scripts dynamically

```bash
docker run -p 8080:8080 \
  -e SCRIPT_DIR=/scripts \
  -v $(pwd)/scripts:/scripts \
  ghcr.io/mrofi/bashscript-server
```

---

## 🔹 Usage

Place your `.sh` scripts in the `scripts` directory.
Example structure:

```
scripts/
├── foo.sh
└── bar.sh
```

Start the server (default port 8080):

```bash
docker run -p 8080:8080 -v $(pwd)/scripts:/app/scripts bashscript-server
```

Run a script with arguments:

```bash
curl -sL http://localhost:8080/foo.sh | bash -s -- arg1 arg2
```

Inside `foo.sh`:

```bash
#!/bin/bash
echo "Foo script running with args: $@"
```

Output:

```
Foo script running with args: arg1 arg2
```

---

## 🔹 Environment Variables

| Variable   | Default        | Description                         |
| ---------- | -------------- | ----------------------------------- |
| SCRIPT_DIR | `/app/scripts` | Directory where scripts are stored. |
| PORT       | `8080`         | Port to serve HTTP requests.        |

---

## 🔹 Kubernetes Deployment Example

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: bashscript-server
spec:
  replicas: 1
  selector:
    matchLabels:
      app: bashscript-server
  template:
    metadata:
      labels:
        app: bashscript-server
    spec:
      containers:
        - name: server
          image: your-dockerhub/bashscript-server:latest
          ports:
            - containerPort: 8080
---
apiVersion: v1
kind: Service
metadata:
  name: bashscript-server
spec:
  selector:
    app: bashscript-server
  ports:
    - port: 80
      targetPort: 8080
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: bashscript-ingress
  annotations:
    kubernetes.io/ingress.class: "nginx"
spec:
  rules:
    - host: myscripts.io
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: bashscript-server
                port:
                  number: 80
```

---

## 🔹 License

MIT License. See [LICENSE](LICENSE) for details.


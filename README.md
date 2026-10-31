# nginx-go-postgres-cicd

Multi-service application (Nginx frontend, Go backend, PostgreSQL database)
built and deployed with Docker Compose and a GitLab CI/CD pipeline.


## Architecture

Browser → Nginx (port 80) → Go backend (port 8080) → PostgreSQL

- Nginx serves the static page and proxies `/api/` to the backend.
- The backend exposes `/health` and checks database availability.
- PostgreSQL stores data in a named volume.


## Project structure

```
frontend/   index.html, nginx.conf, Dockerfile
backend/    main.go, go.mod, Dockerfile
docker-compose.yml
.gitlab-ci.yml
```

## CI/CD pipeline

| Stage | What it does |
|---|---|
| test | `go vet` and `go test` on the backend |
| build | Builds frontend and backend images, pushes them to the GitLab Container Registry tagged with the commit SHA |
| deploy | Connects to the server over SSH, uploads the compose file, generates `.env` from CI variables, pulls images and runs `docker compose up -d` |

> The GitLab Container Registry can be replaced with another registry such as Nexus. 

## CI/CD variables

`SERVER_IP`, `SERVER_USER`, `SSH_PRIVATE_KEY_B64`,
`POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB`

> Values are stored in GitLab CI/CD settings and never committed.

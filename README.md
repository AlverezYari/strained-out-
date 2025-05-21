# Imperial Assault Companion Prototype

This repository contains a minimal scaffolding for a lightweight Imperial Assault campaign assistant.
The goal is to keep the stack simple and maintainable while allowing future expansion.

## Structure

```
backend/   Go server exposing REST endpoints
  cmd/server    entrypoint for running the API server
  internal/game game logic and data models
frontend/  Vue 3 application (TypeScript + Pinia + Vue Router)
  src/          application source
  public/       static assets / index.html
```

## Development

This project uses [Task](https://taskfile.dev) to simplify common workflows. After installing Task, you can bootstrap both the backend and frontend in one command:

```bash
task dev
```

The `dev` task installs frontend dependencies if needed, starts the Go API server on `:8080`, and launches the Vite dev server with hot reload.

Other helpful tasks include:

```bash
task build   # build Go binary and frontend assets
task lint    # run Go vet and frontend lint
task test    # run backend and frontend tests
task docker  # build Docker image containing the full app
```

Refer to `Taskfile.yml` for the exact commands.

### Docker Images

Three Dockerfiles are provided:

1. `Dockerfile.backend` – builds and runs the Go server alone.
2. `Dockerfile.frontend` – builds and serves the compiled Vue frontend with nginx.
3. `Dockerfile` – multi-stage build running both pieces together.

Run `task docker` to build the combined image tagged `imperial-app`.

## Game Models

The core data models live in [`backend/internal/game/models.go`](backend/internal/game/models.go):
- `Unit` – player or enemy with position, health, actions, and conditions
- `Tile` – single map tile with an image and grid position
- `Trigger` – basic mission scripting hook (round/location/condition)
- `Mission` – collection of tiles, units, and triggers representing a scenario
- `State` – top level object served to the frontend

## Frontend Components

- **pages/** – route level components (e.g. `HomePage.vue`)
- **components/** – smaller reusable pieces (map renderer, unit info, etc.)
- **store/** – Pinia stores for game state
- **router/** – Vue Router setup
- **types/** – shared TypeScript interfaces matching backend models

Keep components focused and as stateless as possible. Use composables or store actions for logic.

## Next Steps

- Flesh out map rendering from `mission.map` data
- Add interactions for units and triggers
- Implement save/load of `State` to disk on the server

This scaffolding should provide a clean starting point for further development.

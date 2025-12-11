# Thunder Sporting Event Platform

Thunder is a dual-stack platform that mirrors the aesthetic and experience of [suffix.events](https://suffix.events) while adding a production-ready backend for managing premium sporting experiences.

## Architecture

- **Backend (`/backend`)** — Golang + Echo API that surfaces events and captures leads. It authenticates admin-only routes with Keycloak-issued JWTs and persists content through PocketBase.
- **Frontend (`/frontend`)** — Next.js 14 + Tailwind single-page experience that borrows the dark, neon aesthetic from suffix.events and consumes the backend API.
- **PocketBase** — Acts as the headless database for `events` and `leads` collections. The backend uses the REST API through Resty.
- **Keycloak** — Provides OAuth2/OpenID Connect flows and token introspection for the `/api/events` write endpoints.

```
frontend (Next.js)  <--->  backend (Echo)  <--->  PocketBase
        |                                  \
        └-------------- Keycloak ------------\
```

## Requirements

- Go 1.21+
- Node.js 20+ and npm
- PocketBase instance with `events` and `leads` collections
- Keycloak realm + confidential client for the API

## Backend Setup

```bash
cp backend/.env.example backend/.env
# edit backend/.env with your values
cd backend
go run ./cmd/server
```

Environment variables:

| Variable | Purpose |
| --- | --- |
| `SERVER_PORT` | Port Echo listens on (default `8080`) |
| `CORS_ALLOWED_ORIGINS` | Comma-separated list for the frontend origin |
| `KEYCLOAK_*` | Base URL, realm, client credentials used for token introspection |
| `POCKETBASE_*` | PocketBase REST endpoint, admin API token, and public files URL |

### PocketBase collections

Create collections named `events` and `leads` with fields that match the payloads in `internal/domain`. Files (cover + gallery) should use PocketBase file fields so the API can construct public URLs.

### Keycloak

1. Create a confidential client (e.g., `thunder-api`) in your realm.
2. Enable service accounts and client credentials.
3. Copy the client ID and secret into `backend/.env`.
4. Protect the POST `/api/events` route by sending `Authorization: Bearer <token>` from a Keycloak client.

## Frontend Setup

```bash
cp frontend/.env.example frontend/.env
# update NEXT_PUBLIC_API_BASE_URL if backend runs elsewhere
cd frontend
npm install
npm run dev
```

The homepage (`/`) is a server component that fetches live events. If the API is unavailable it falls back to seeded concept events to preserve the marketing experience.

## One-command stack (Docker Compose)

Use the provided Dockerfiles + compose file when you want a reproducible stack that runs PocketBase, the Echo API, and the Next.js frontend together.

```bash
cp backend/.env.example backend/.env
cp frontend/.env.example frontend/.env
# set KEYCLOAK_* secrets and point POCKETBASE_BASE_URL to http://pocketbase:8090 for container-to-container comms
# keep POCKETBASE_FILES_BASE_URL pointing to http://localhost:8090 so browsers can load media
cd /workspace
docker compose up --build
```

What you get:

- `http://localhost:8090` → PocketBase admin UI/storage (data persisted in the `pocketbase_data` volume)
- `http://localhost:8080` → Echo API (still protected by Keycloak for POST `/api/events`)
- `http://localhost:3000` → Thunder marketing experience

Stop the stack with `docker compose down` (add `-v` if you also want to wipe the PocketBase volume).

## Development Tips

- Run `go test ./...` inside `backend` to make sure the API compiles.
- Run `npm run lint` inside `frontend` to keep the UI clean.
- Update both `.env` files before deploying; neither is committed to git.

## Directory Layout

```
backend/
  cmd/server/main.go         # Echo bootstrap
  internal/
    config/                  # env parsing
    domain/                  # core models
    http/handler             # Echo handlers
    http/middleware          # Keycloak JWT validation
    http/router              # route + middleware wiring
    repository/pocketbase    # PocketBase REST integration
    service/                 # business logic
frontend/
  app/                       # Next.js App Router pages
  components/                # Client-side widgets
  public/                    # (optional) static assets
```

## Next Steps

- Script PocketBase schema migrations for automated provisioning.
- Add integration tests that stub PocketBase and Keycloak.
- Expand the frontend into additional routes (case studies, insights) and connect to real event detail pages.

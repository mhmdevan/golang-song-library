
# Song Library API — Go + Gin + PostgreSQL + Swagger

A **production-style REST API** for managing a song catalog, built with Go, Gin, PostgreSQL and Swagger.

> **Resume angle:** This project is intentionally structured like a real-world backend service: layered architecture, clean separation of concerns, environment-driven config, tests, Docker support, and API documentation. It is designed so I can walk an interviewer through my engineering decisions, trade-offs, and how I validate and operate a service in production-like conditions.

---

## 1. Problem & Context

### 1.1 What problem does this API solve?

A lot of tutorials show “CRUD over a database” with minimal structure. They are fine for learning syntax but completely useless as portfolio pieces:

- No clear **domain model**.
- No separation between **handler**, **business logic**, and **persistence**.
- No **testing strategy**, no **logging**, no readiness for **Docker** or configuration via environment variables.

The idea behind **Song Library API**:

- Take a very simple domain (songs) and implement it in a way that looks and behaves like a small **production service**.
- Show how I structure Go projects, how I handle **HTTP**, **validation**, **DB access**, **logging**, **configuration**, and **documentation**.
- Create a base that can be realistically extended with **authentication**, **rate limiting**, **search**, or **analytics**.

**Domain:**

- A **Song** has fields such as `id`, `title`, `artist`, `album`, `duration`, `release_year`, etc.
- The API allows clients to **create**, **read**, **update**, and **delete** songs via RESTful endpoints.

### 1.2 Why is this on my resume?

This project demonstrates that I can:

- Design and implement a **layered Go service** using:
  - Gin for HTTP routing.
  - PostgreSQL as the data store.
  - A repository/service pattern to keep logic testable.
- Document APIs using **Swagger/OpenAPI** and keep docs in sync with implementation.
- Build and run the service using **Docker**, with environment-based configuration.
- Write and run **unit and integration tests** with `go test`.
- Think about **metrics**, **logging**, and **operational concerns**, not just “make the code compile”.

In an interview, I can use this project to talk through:

- How I structure new Go services from scratch.
- How I validate input and handle errors consistently.
- How I evolve an API while keeping backward compatibility in mind.
- How I would scale from this simple API to a larger system (more services, auth, caching, etc.).

---

## 2. Tech Stack & Key Decisions

### 2.1 Core technologies

- **Language:** Go (Golang) **1.22+**
- **HTTP framework:** [Gin](https://github.com/gin-gonic/gin)
- **Database:** PostgreSQL
- **Documentation:** Swagger (OpenAPI) via `swag`
- **Containerization:** Docker

### 2.2 Why Gin?

I evaluated a few options:

- **net/http** (standard library)
  - ✅ Zero dependencies, very flexible.
  - ❌ Requires more boilerplate for routing, middleware, error handling.
- **Echo**
  - ✅ Good performance, similar to Gin.
  - ❌ Less familiar ecosystem for some teams.
- **Gin**
  - ✅ Widely used in Go community, great ecosystem.
  - ✅ Simple middleware model, easy JSON handling.
  - ✅ Good developer experience.

> **Decision:** Use **Gin** for a balance between performance, simplicity, and ecosystem support.

### 2.3 Why PostgreSQL?

Alternatives considered:

- **SQLite**
  - ✅ Extremely easy for local development.
  - ❌ Less representative of typical production setups.
- **MySQL**
  - ✅ Common and battle-tested.
  - ❌ Postgres has a richer feature set and is often the default choice in modern backend stacks.
- **PostgreSQL**
  - ✅ Strong data types, good indexing, transactional features.
  - ✅ Works well with Docker and cloud providers.
  - ✅ Familiar choice for production-ready Go services.

> **Decision:** Use **PostgreSQL** to reflect a realistic backend stack and allow scaling beyond a toy project.

### 2.4 Architecture style

The project follows a **layered architecture**:

- **Handler layer (`internal/handler`)**
  - Responsible for HTTP concerns: parsing requests, binding/validating input, returning JSON responses.
- **Router layer (`internal/router`)**
  - Sets up all routes, attaches middleware, groups endpoints by resource.
- **Service layer (`internal/service`)**
  - Contains business logic: validation rules, orchestration of repository calls, domain-level decisions.
- **Repository layer (`internal/repository`)**
  - Handles all DB interactions, hides SQL/queries from upper layers.
- **DB layer (`internal/db`)**
  - Manages database connection, pooling, and migration hooks.
- **Model layer (`internal/model`)**
  - Declares domain structs that represent Songs and related types.
- **Logger (`pkg/logger`)**
  - Centralized logging utilities used across layers.

This structure keeps the codebase **testable**, **extensible**, and **readable**, especially for teams.

---

## 3. Metrics, Quality & Operational Focus

Even though this is a demo project, it is designed to expose the same levers you care about in production:

- **Performance metrics**
  - Average and p95 latency per endpoint (e.g. `GET /api/v1/songs`).
  - Request throughput (requests per second).
- **Reliability metrics**
  - Error rate (4xx/5xx ratio).
  - DB connection errors/timeouts.
- **Data quality metrics**
  - Percentage of invalid requests rejected by validation.
  - Consistency between API responses and DB state in tests.

In a real deployment, you would wire these into Prometheus/Grafana via:

- HTTP middleware to measure request latency and response codes.
- DB wrapper to count connection errors and query durations.

> In interviews, I use this project to talk about **how I would add metrics** even if the sample repo doesn’t include a full monitoring stack.

---

## 4. Project Structure

```plaintext
.
├── cmd/
│   └── server/
│       └── main.go        # Application entry point (bootstrap, DI, server start)
├── internal/
│   ├── handler/           # HTTP handlers (controllers)
│   ├── router/            # Route registration and middleware
│   ├── service/           # Business logic / use cases
│   ├── repository/        # Database access layer
│   ├── db/                # DB connection, migrations
│   ├── model/             # Domain models (Song, etc.)
├── pkg/
│   └── logger/            # Logging utilities
├── docs/                  # Swagger / OpenAPI documentation
├── go.mod                 # Module definition
├── go.sum                 # Dependency checksums
├── .env                   # Environment variables (not committed)
└── README.md              # Project documentation (this file)
```

**Key ideas:**

- `cmd/server/main.go` is intentionally thin: it wires dependencies and starts the HTTP server.
- `internal` is used so no external packages accidentally depend on these modules.
- The separation into handler/service/repository makes it easy to mock out dependencies in tests.

---

## 5. Setup & Installation

### 5.1 Prerequisites

- Go **1.22+**
- PostgreSQL
- Docker (optional, but recommended for local DB)

### 5.2 Environment configuration

Create a `.env` file in the root directory:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=song_library
SERVER_PORT=8080
```

These variables control how the app connects to the database and which port it listens on.

### 5.3 Install dependencies

```bash
go mod tidy
```

### 5.4 Database setup

Create the database (if not already created):

```bash
createdb song_library
```

Run migrations (depending on how migrations are implemented, e.g. using a tool like `golang-migrate`, `goose`, or custom code in `internal/db`). For example:

```bash
# Example pattern (pseudo-command):
go run ./cmd/server/main.go migrate
```

If you do not have a migration tool wired yet, you can manually create the `songs` table using a simple SQL script such as:

```sql
CREATE TABLE IF NOT EXISTS songs (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    artist TEXT NOT NULL,
    album TEXT,
    duration_seconds INT,
    release_year INT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
```

### 5.5 Running tests

Run unit and integration tests:

```bash
go test ./...
```

In an interview, I can discuss:

- What is covered by tests (service and repository layers).
- How I would structure integration tests using a test database or Docker.

### 5.6 Generate Swagger documentation

Swagger docs are generated from annotations in the Go code using `swag`:

```bash
swag init --output ./docs --generalInfo ./cmd/server/main.go
```

This generates the OpenAPI specification and static assets in `./docs`, which are then served by the application.

### 5.7 Run the application (locally)

Start the server:

```bash
go run ./cmd/server/main.go
```

The API is now available at:

```text
http://localhost:8080/api/v1/songs
```

Swagger UI is available at:

```text
http://localhost:8080/swagger/index.html
```

---

## 6. Running with Docker

Containerization support makes the app easier to run in a clean environment and closer to how it might be deployed in production.

### 6.1 Build the Docker image

```bash
docker build -t song-library .
```

### 6.2 Run the container

```bash
docker run --name song-library \
  -p 8080:8080 \
  --env-file .env \
  song-library
```

You can also create a `docker-compose.yml` to run the API and PostgreSQL together. For example:

```yaml
version: "3.8"
services:
  db:
    image: postgres:16
    environment:
      POSTGRES_USER: your_user
      POSTGRES_PASSWORD: your_password
      POSTGRES_DB: song_library
    ports:
      - "5432:5432"
    volumes:
      - dbdata:/var/lib/postgresql/data

  api:
    build: .
    depends_on:
      - db
    environment:
      DB_HOST: db
      DB_PORT: 5432
      DB_USER: your_user
      DB_PASSWORD: your_password
      DB_NAME: song_library
      SERVER_PORT: 8080
    ports:
      - "8080:8080"

volumes:
  dbdata:
```

Run everything with:

```bash
docker compose up --build
```

---

## 7. API Design

### 7.1 Base URL

```text
/api/v1
```

### 7.2 Endpoints

| Method | Endpoint            | Description               |
|--------|---------------------|---------------------------|
| GET    | /songs              | Retrieve all songs        |
| GET    | /songs/:id          | Retrieve a song by ID     |
| POST   | /songs              | Add a new song            |
| PUT    | /songs/:id          | Update an existing song   |
| DELETE | /songs/:id          | Delete a song by ID       |

### 7.3 Example: Create a song

**Request**

```http
POST /api/v1/songs
Content-Type: application/json

{
  "title": "Everlong",
  "artist": "Foo Fighters",
  "album": "The Colour and the Shape",
  "duration_seconds": 250,
  "release_year": 1997
}
```

**Possible validation rules:**

- `title`: required, non-empty, reasonable max length (e.g. 255).
- `artist`: required, non-empty.
- `duration_seconds`: optional but if present, must be positive.
- `release_year`: optional but if present, must be within a sane range (e.g. 1900–current_year).

**Response**

```http
201 Created
Content-Type: application/json

{
  "id": 1,
  "title": "Everlong",
  "artist": "Foo Fighters",
  "album": "The Colour and the Shape",
  "duration_seconds": 250,
  "release_year": 1997,
  "created_at": "2025-01-01T12:00:00Z",
  "updated_at": "2025-01-01T12:00:00Z"
}
```

### 7.4 Error handling

Errors are returned in a consistent JSON format, for example:

```json
{
  "error": "validation_failed",
  "message": "title is required"
}
```

In an interview, I can explain:

- How I centralize error handling in middleware.
- How I differentiate between client errors (4xx) and server errors (5xx).
- How I log errors with context (request ID, path, etc.).

---

## 8. Swagger / OpenAPI

The project uses **Swagger** for:

- Documenting each endpoint: parameters, responses, error codes.
- Providing a **live playground** at `/swagger/index.html` where developers can test the API.
- Keeping API docs close to the code via annotations.

This demonstrates that I can:

- Maintain up-to-date API documentation.
- Make services easier to integrate with other teams/clients.

Swagger UI is accessible at:

```text
/swagger/index.html
```

---

## 9. Logging & Observability

### 9.1 Logging

The `pkg/logger` package centralizes logging configuration so that:

- Handlers and services can log structured messages.
- Logs include useful context: request path, status code, errors.
- In production, logging can be switched to JSON for easier ingestion into tools like ELK or Loki.

Examples of events we log:

- Incoming requests (method, path).
- DB errors (query failed, connection issues).
- Business-level events (song created, updated, deleted).

### 9.2 Metrics (future enhancement)

The service is structured so that adding metrics is straightforward:

- Wrap Gin handlers with middleware to track:
  - Request duration.
  - Response status codes.
- Wrap DB operations to measure query times and errors.

This can be exported to **Prometheus** and visualized in **Grafana** as:

- Latency histograms per endpoint.
- Error rate dashboards.
- DB health and usage metrics.

---

## 10. Testing Strategy

Even though this is a relatively small service, the structure is **test-friendly**:

- **Unit tests** for `service` layer:
  - Validate business rules (e.g. cannot create song with empty title).
  - Ensure we handle repository errors correctly.
- **Unit/integration tests** for `repository` layer:
  - Use a test database (or Dockerized Postgres) to ensure queries behave as expected.
- **Handler tests (optional)**:
  - Use Gin’s testing helpers or `httptest` to test endpoint behavior end-to-end.

Command to run all tests:

```bash
go test ./...
```

In an interview, I talk about:

- How I would add **CI** to run `go test ./...` and `golangci-lint` on every push.
- How I would structure **integration tests** with Docker Compose.

---

## 11. Challenges & How They Were Solved

This project intentionally touches a few non-trivial areas that often show up in real services.

### 11.1 Consistent error handling in Gin

**Problem:** Without a clear pattern, handlers return different error formats and status codes, making the API hard to consume.

**Solution:**

- Introduce a standard error response structure.
- Use helper functions (e.g. `respondError`) that map internal errors to HTTP codes.
- Optionally use middleware for panic recovery and capturing unhandled errors.

### 11.2 Clean separation of layers

**Problem:** It’s easy to end up with SQL and business logic mixed directly into handlers.

**Solution:**

- Enforce a strict separation:
  - Handlers: HTTP, request/response.
  - Services: business rules, orchestration.
  - Repositories: DB details.
- This makes it much easier to test and evolve the code.

### 11.3 Local vs. containerized environment

**Problem:** Code that works on a local machine doesn’t always work inside a Docker container (e.g. hostnames, ports, env vars).

**Solution:**

- Use environment variables for all DB and server configuration.
- Provide a `docker-compose.yml` that sets matching variables for API and DB.
- In the README, document both `go run` and `docker compose` flows.

---

## 12. Possible Extensions

If I had more time or in a real production context, natural extensions would be:

- **Authentication & authorization**
  - JWT-based auth or API keys.
  - Role-based access control (admin vs. regular user).
- **Search & filtering**
  - Query parameters for filtering songs by artist, album, or year.
  - Pagination and sorting.
- **Caching**
  - In-memory cache or Redis for frequently accessed reads.
- **Rate limiting**
  - Protect the API from abuse using middleware (e.g., token bucket per IP/API key).
- **Async processing**
  - For heavy tasks (batch imports, analytics), use background workers or message queues.

These are exactly the kinds of “how would you evolve this” questions that often appear in interviews — and this API is structured to support those answers.

---

## 13. Interview Cheat Sheet

When talking about **Song Library API** in an interview, I focus on:

- **Problem:** Build a small, but realistic REST service that manages songs and is structured like a production backend.
- **Tech:** Go 1.22, Gin, PostgreSQL, Swagger, Docker.
- **Architecture:** Layered design (handler, router, service, repository, db, model, logger). Environment-driven config, swagger docs, Docker support.
- **Quality:** Tests via `go test ./...`, clean error handling, consistent JSON responses.
- **Metrics & ops:** Designed to expose latency/error metrics and integrate with Prometheus/Grafana. Structured logging with a centralized logger.
- **Challenges:** Consistent error handling, clear separation of concerns, making local + Docker workflows work smoothly.
- **Future work:** Auth, search/filtering, caching, rate limiting, multi-service architecture.

This way, the project is not “just CRUD”, but a concrete example of how I design and build backend services in Go.

---

## 14. License

```text
MIT License

Copyright (c) 2025 Mohammad Eslamnia
```

# Connexions Codegen Template

Generate a Go mock server from OpenAPI specs with [Connexions](https://github.com/mockzilla/connexions).

Each service is a Go package with generated handlers, embedded OpenAPI specs, and optional custom logic.
Includes an API Explorer UI at `/api-explorer`.

## Quick start

1. Click [**Use this template**](https://github.com/mockzilla/connexions-codegen-template/generate) to create your own repository
2. Add services (see below)
3. Regenerate and discover:
   ```bash
   make generate && make discover
   ```
4. Push to main - binaries for Linux, macOS, and Windows are built automatically and published to **Releases**

## Adding services

### From an OpenAPI spec

```bash
make service name=my-api
```

This creates `pkg/my_api/` with a scaffold. Replace `pkg/my_api/setup/openapi.yml` with your spec, then run:

```bash
make generate && make discover
```

### From static responses

```bash
make service-from-static name=my-api
```

Add response files under `pkg/my_api/setup/data/` organized by method and path:

```
pkg/my_api/setup/data/
  get/
    users/
      index.json            -> GET /users
    users/{id}/
      index.json            -> GET /users/{id}
  post/
    users/
      index.json            -> POST /users
```

Then regenerate:

```bash
make generate && make discover
```

## Service structure

Each service lives in `pkg/{service_name}/` with:

```
pkg/petstore/
  setup/
    openapi.yml             # OpenAPI specification
    config.yml              # Latency, errors, upstream, caching
    codegen.yml             # Code generation settings
    context.yml             # Custom values for mock data
  generate.go               # go:generate directive
  gen.go                    # Generated handler (do not edit)
  service.go                # Custom logic (optional overrides)
```

## API Explorer UI

The built-in UI is available at `/` (configurable in `resources/data/app.yml`).
It lets you browse services, view specs, test endpoints, and inspect request/response history.

## Local development

```bash
make build
.build/server/server
```

See the [Connexions docs](https://github.com/mockzilla/connexions) for more options.

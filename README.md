# Mockzilla Codegen Template

Generate a Go mock server from OpenAPI specs with [Mockzilla](https://github.com/mockzilla/mockzilla).

Each service is a Go package with generated handlers, embedded OpenAPI specs, and optional custom logic.
Includes an API Explorer UI at `/api-explorer`.

## Quick start

1. Click [**Use this template**](https://github.com/mockzilla/mockzilla-codegen-template/generate) to create your own repository
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

See the [Mockzilla docs](https://github.com/mockzilla/mockzilla) for more options.

## Release

Every push to main/master generates code, builds binaries for Linux, macOS, and Windows (amd64 and arm64), 
and publishes them to the **Releases** page. 
Download the binary for your platform from the `latest` release and run it locally:

```bash
./your-repo-name
```

## Mockzilla workflow

The included GitHub Actions workflow (`.github/workflows/mockzilla.yml`) publishes your server to [Mockzilla](https://mockzilla.org) automatically:

- **Push to main/master** — builds and publishes the server to your main simulation
- **Pull request with `Mockzilla` label** — deploys a preview simulation for the PR (torn down when the PR is closed)

The `Mockzilla` label is created automatically on first push via the setup workflow.

Your simulation will be available at:
- `https://api.mockzilla.org/gh/{org}/{repo}/` — main branch
- `https://api.mockzilla.org/gh/{org}/{repo}/pr-{n}/` — per pull request

### Action inputs

You can customize the action in `.github/workflows/mockzilla.yml`:

```yaml
- uses: mockzilla/actions/codegen@v1
  with:
    token: ${{ secrets.GITHUB_TOKEN }}
    region: us-east-1        # optional — preferred AWS region, used as a hint on first deploy only
    memory-size: 256         # optional — memory in MB (default: 128)
    timeout: 60              # optional — request timeout in seconds
    environment: '{"ENV":"production","DEBUG":"true"}'  # optional
    host: api.mockzilla.net  # optional — API host for the simulation URL
    timeout-minutes: 5       # optional — max minutes to wait for simulation to become active (default: 5)
    delete: false            # optional — remove this repository from Mockzilla (default: false)
```

| Input | Required | Description |
|---|---|---|
| `token` | yes | `GITHUB_TOKEN` — used to verify repo identity |
| `region` | no | Preferred AWS region (e.g. `us-east-1`, `ap-southeast-1`). Used as a hint on first deploy — if at capacity, the nearest available region is used. Has no effect after the simulation is deployed. |
| `memory-size` | no | Memory in megabytes (e.g. `128`, `256`, `512`). Defaults to `128`. |
| `timeout` | no | Request timeout for the simulation in seconds (e.g. `30`, `60`). |
| `environment` | no | JSON object of environment variables to set in the simulation (e.g. `'{"ENV":"production"}'`). |
| `host` | no | API host for the simulation URL (`api.mockzilla.org`, `api.mockzilla.de`, or `api.mockzilla.net`). Defaults to org setting or `api.mockzilla.org`. |
| `timeout-minutes` | no | Max minutes the action polls for the simulation to become active. Defaults to `5`. |
| `delete` | no | Remove this repository from Mockzilla. When set to `true`, the action skips publishing and deletes all mock APIs for this repo. Useful on the free plan to free up your slot before connecting a different repository. Defaults to `false`. |

### Removing this repository from Mockzilla

On the free plan you can only have one repository connected to Mockzilla at a time. To switch to a different repo, run the action with `delete: true` on the old one first:

```yaml
name: mockzilla-remove

on:
  workflow_dispatch:

jobs:
  remove:
    runs-on: ubuntu-latest
    steps:
      - uses: mockzilla/actions/codegen@v1
        with:
          token: ${{ secrets.GITHUB_TOKEN }}
          delete: true
```

Trigger it manually from the **Actions** tab when you're ready.

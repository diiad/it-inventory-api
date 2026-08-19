# IT Inventory API

REST API in Go for managing IT assets: machines, collaborators, sites, and workstation deployment events. Personal side project to go deeper into Go + PostgreSQL.

## Stack

- **Go** (`net/http`, native Go 1.22+ routing)
- **PostgreSQL** (`github.com/lib/pq`)

## Features

Full CRUD on:
- `machine` — device fleet, linked to a `model`
- `collabs` — collaborators
- `site` — physical sites
- `soclage_event` — workstation deployment history (date, IP, project, floor, desk...)

## Run it

1. Create the PostgreSQL database with `script postgreSQL.sql`
2. Adjust the connection settings in `db/db.go` (host, user, password, dbname)
3. Start the API:

```bash
go run .
```

The API starts on `http://localhost:8080`.

## Main routes

| Method | Route | Description |
|---|---|---|
| GET | `/` | Health check (DB connection test) |
| POST/GET/PATCH/DELETE | `/machine`, `/machine/{id}` | Machine CRUD |
| POST/GET/PATCH/DELETE | `/collabs`, `/collabs/{id}` | Collaborator CRUD |

## Structure

```
api.go            ← entry point, routes
app/              ← HTTP handlers (machine, collabs, site, soclage_event)
db/               ← PostgreSQL connection + queries
script postgreSQL.sql ← database schema
```

---

## Version française

API REST en Go pour la gestion d'un parc informatique : machines, collaborateurs, sites et événements de "soclage" (déploiement/imagerie de poste de travail). Side project personnel pour approfondir Go + PostgreSQL.

### Stack

- **Go** (`net/http`, routage natif Go 1.22+)
- **PostgreSQL** (`github.com/lib/pq`)

### Lancer le projet

1. Créer la base PostgreSQL avec le script `script postgreSQL.sql`
2. Adapter la configuration de connexion dans `db/db.go` (host, user, password, dbname)
3. Lancer l'API : `go run .` (démarre sur `http://localhost:8080`)

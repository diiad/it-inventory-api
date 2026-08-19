# IT Inventory API

API REST en Go pour la gestion d'un parc informatique : machines, collaborateurs, sites et événements de "soclage" (déploiement/imagerie de poste de travail). Side project personnel pour approfondir Go + PostgreSQL.

## Stack

- **Go** (`net/http`, routage natif Go 1.22+)
- **PostgreSQL** (`github.com/lib/pq`)

## Fonctionnalités

CRUD complet sur :
- `machine` — parc de machines, liées à un modèle (`model`)
- `collabs` — collaborateurs
- `site` — sites physiques
- `soclage_event` — historique des déploiements de poste (date, IP, projet, étage, bureau...)

## Lancer le projet

1. Créer la base PostgreSQL avec le script `script postgreSQL.sql`
2. Adapter la configuration de connexion dans `db/db.go` (host, user, password, dbname)
3. Lancer l'API :

```bash
go run .
```

L'API démarre sur `http://localhost:8080`.

## Routes principales

| Méthode | Route | Description |
|---|---|---|
| GET | `/` | Health check (test de connexion DB) |
| POST/GET/PATCH/DELETE | `/machine`, `/machine/{id}` | CRUD machines |
| POST/GET/PATCH/DELETE | `/collabs`, `/collabs/{id}` | CRUD collaborateurs |

## Structure

```
api.go            ← point d'entrée, routes
app/              ← handlers HTTP (machine, collabs, site, soclage_event)
db/               ← connexion PostgreSQL + requêtes SQL
script postgreSQL.sql ← schéma de la base
```

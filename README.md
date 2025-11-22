# Projects API

**Note:** This project is no longer running on production. I am now storing the data on the [front-end](https://github.com/andrewjamesmoore/andrew-projects).

---

A GraphQL API built with Go and MongoDB for portfolio projects and notes.

Access the production endpoint at: [https://api.andr3w.sh](https://api.andr3w.sh). 

View the [light-weight front-end](https://github.com/andrewjamesmoore/andrew-projects) built using Astro.

---

## Stack

- **Go** (with gqlgen) — GraphQL server
- **MongoDB** — Database
- **Docker Compose** — Container orchestration
- **Nginx** — Reverse proxy and SSL
- **Certbot** — HTTPS via Let's Encrypt
- **Terraform** — Infrastructure management on AWS EC2

## Deployment + Workflow

- Mutations are locked down (CORS + no public access).
- Projects are added via CLI or mutation playground.
- On mutation, the frontend fetches new data automatically via SSR.
- No rebuilds. No re-deploys. Data changes appear instantly.

# Projects API

A GraphQL API built with Go, MongoDB, and deployed using Docker on an EC2 instance.  
Access the production endpoint at: [https://api.andr3w.sh](https://api.andr3w.sh)

Is this overkill? Yes, but in the process of learning some GraphQL I set up an API for my portfolio projects using Go.

---

## Stack

- **Go** (with gqlgen) — GraphQL server
- **MongoDB** — Database
- **Docker Compose** — Container orchestration
- **Nginx** — Reverse proxy and SSL termination
- **Certbot** — HTTPS via Let's Encrypt
- **Terraform** — Infrastructure management on AWS EC2

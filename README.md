# Tezos Delegation Service

This service provides an API to fetch and manage Tezos delegation data.

## Features

- Fetch delegation data from the Tezos blockchain
- Store delegation data in a PostgreSQL database
- Provide API endpoints to retrieve delegation information

## Prerequisites

- Go 1.22 or higher
- Make
- Docker and Docker Compose (optional but recommended)

## Getting Started

Clone the repository:

```bash
git clone https://github.com/lat1992/tezos-delegation-service.git
cd tezos-delegation-service
```

## Build with docker (recommended)
Build and run the service using docker build and compose:

```bash
make install
```

This command will build the Docker image and start the service along with a PostgreSQL database.

The API will be available at `http://localhost:8080`

## Build without docker

1. Setup a PostgresSQL database and init with the ./scripts/postgres.sql file

2. Add env variable
```bash
export PORT=8080
export DATABASE=postgres_password
```

3. Build and start the service
```bash
make
./build/tezos-delegation-service
```

## API Endpoints

- `GET /`: Health check
- `GET /health`: Health check
- `POST /xtz/delegations/[:year]`: Fetch delegations with or without a specific year

## Configuration

The service can be configured using environment variables or a configuration file. The following variables are available:

- `PORT`: API server port (default: 8080)
- `DATABASE_HOST`: PostgreSQL host (default: localhost)
- `DATABASE_PORT`: PostgreSQL port (default: 5432)
- `DATABASE_DATABASE`: PostgreSQL database name (default: app)
- `DATABASE_USER`: PostgreSQL user (default: postgres)
- `DATABASE_PASSWORD`: PostgreSQL password
- `TEZOS_API`: Tezos API endpoint (default: https://api.tzkt.io/v1)

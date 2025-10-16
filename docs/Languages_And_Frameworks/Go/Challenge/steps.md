# Steps

1. Initialized a Go module

```shell
go mod init gochat
```

2. Installed required packages

```shell
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
go get github.com/gin-gonic/gin
go get github.com/jmoiron/sqlx
go get github.com/lib/pq
```

3. Install podman on Debian

```shell
sudo apt-get -y install podman
```

4. Installed official Postgres image

```shell
podman pull docker.io/postgres:latest
```

5. Started the container with

```shell
podman run -d --name postgres-db -e POSTGRES_DB=gochat -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=secret -p 5432:5432 docker.io/postgres:latest
```

6. Execed into the container with

```shell
podman exec -it <container-id> /bin/bash
```

7. Logged into Postgres DB

```shell
psql -U admin -d gochat
```

3. Defined the User Model and DB Schema

- In Postgres, created the `users` table:

```shell
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);
```




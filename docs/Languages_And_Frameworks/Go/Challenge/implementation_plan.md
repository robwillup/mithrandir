# Implementation Plan

## 1. Project Layout & Structure

- Use Go modules for depencency management
- Organize code into packages
  - `api`: HTTP handler for login, registration, chat messaging, and commands.
  - `auth`: Local user registration, password hashing (bcrypt), JWT issuance and middleware.
  - `chat`: Chatroom state, message storage, retrieval, message ordering.
  - `bot`: Decoupled bot logic subscribing to command queue, querying stock API, posting stock quotes via RabbitMQ.
  - `models`: Shared data structures (User, Message, Token)
  - `db`: Database connection and queries (Postgres preferred)
  - `mq`: RabbitMQ client abstraction
  - `tests`: Unit tests for core logic.
- Use `.http` files for API testing (login, chat, commands)

## 2. User Authentication with Local DB

- Users stored securely in Postgress with bcrypt password hashes.
- API endpoints:
  - `POST /register` - register new username + password.
  - `POST /login` - authenticate and generate JWT token.
- JWT middleware to secure chat and command endpoints.

## 3. Chatroom and Messaging

- Single chatroom initially.
- Store messages in Postgres with timestamp and sender ID.
- Endpoint:
  - `POST /messages` - send message or command.
  - `GET /messages` - fetch the last 50 messages ordered by timestamp.
- Filter out `/stock=...` messages from storage; instead publish to RabbitMQ.

## 4. Bot Service with RabbitMQ

- Bot subscribes to RabbitMQ queue for `/stock=stock_code` commands.
- Fetch stock CSV from `https://stooq.com/q/l/?s=<stock_code>&f=sd2t2ohlcv&h&e=csv`.
- Parse CSV and post back a formatted quote message to RabbitMQ.
- Main backend subscribes to bot's messages and persists them in the chat as messages from the bot.

## 5. Containerization using Podman

- Backend container:
  - Build Go app static binary
  - Include config for DB and RabbitMQ connection via environment variables.
  - Expose service port (e.g., :8080)

- Postgress container:
  - Use official Postgres image with volume for persistent storage.
  - Set environment vars for user/password/db.

- RabbitMQ container:
  - Use official RabbitMQ image.
  - Set ports and credentials.

- Networking:
  - Podman network connecting containers by hostname.

- Orchestration:
  - Provide a `podman-compose.yml` or scripts to:
    - Build backend image.
    - Start containers with correct environment/volumes.
    - Tear down and cleanup.

## 6. Security and Best Practices

- Store secrets (DB password, JWT secret) safely using environment variables.
- Use bcrypt for password hashing, with secure cost factor.
- Secure RabbitMQ connection with authentication.
- JWT middleware to check token expiry and validity.
- Validate and sanitize user inputs.
- Limit chat messages fetched to the latest 50 to avoid excessive resource use.

## 7. Testing and Documentation

- Provide `.http` request files for:
  - Registering users.
  - Logging in users and extracting JWT.
  - Sending chat messages.
  - Posting stock commands and verifying bot responses.
  - Retrieving the last 50 messages.
- Document setup and startup instructions with Podman:
  - Building backend container.
  - Running all containers with `podman-compose.yml`.
  - How to execute `.http` tests with tools like VSCode REST Client.

## 8. Optional Future Enhancements (Bonus)

- Support multiple chatrooms with messages scoped per room.
- Bot handles unrecognized commands gracefully.
- Frontend with simple UI.
- More sophisticated stock API handling.

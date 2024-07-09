# Trello Clone

This project is a Trello clone written in Golang. The application replicates the core functionalities of Trello,
allowing users to create
workspace, users, tasks, subTasks and roles for every user in every workspace.

## Technologies used

- [Golang](https://golang.org/), Programming language.
- [Echo](https://echo.labstack.com/), HTTP web framework.
- [Gorm](https://gorm.io/), ORM library for Golang.
- [PostgreSQL](https://www.postgresql.org/), database management system.
- [Docker](https://www.docker.com/), Containerization platform.

## Features

- CRUD operations for users, workspaces, tasks, and subtasks
- User authentication using JWT tokens
- Authorization to ensure user-specific limitations and permissions

## Getting Started

### Prerequisites

- Golang
- Docker

### Setting up

Build and run the application using Docker Compose:

```shell
docker-compose -f deployment/docker-compose.yml up -d --build
```

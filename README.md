# Trello Clone

## Introduction

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

Structure Details
----------------------

### Workspace

- **Fields:**
    - `id`: Unique identifier
    - `name`: Name of the workspace (NOT NULL)
    - `description`: A brief description of the workspace
    - `created_at`: Date and time when the workspace was created
    - `updated_at`: Date and time when the workspace was last updated

### Task

- **Fields:**
    - `id`: Unique identifier
    - `title`: Title of the task (NOT NULL)
    - `description`: Detailed description of the task
    - `status`: Current status of the task (Planned, In Progress, Completed)
    - `estimated_time`: Estimated time required to complete the task
    - `actual_time`: Actual time spent on the task
    - `due_date`: Deadline for the task
    - `priority`: Priority level of the task
    - `workspace_id`: Identifier for the workspace to which the task belongs (NOT NULL, Foreign key)
    - `created_at`: Date and time when the task was created
    - `updated_at`: Date and time when the task was last updated
    - `image_url`: URL of the image associated with the task

### SubTask

- **Fields:**
    - `id`: Unique identifier
    - `task_id`: Identifier of the main task (NOT NULL, Foreign key)
    - `title`: Title of the subtask (NOT NULL)
    - `is_completed`: Completion status of the subtask (Yes/No)
    - `created_at`: Date and time when the subtask was created
    - `updated_at`: Date and time when the subtask was last updated

### User

- **Fields:**
    - `id`: Unique identifier
    - `username`: Username of the user (UNIQUE, NOT NULL)
    - `email`: Email address of the user (UNIQUE, NOT NULL)
    - `password_hash`: Hashed password for user authentication
    - `created_at`: Date and time when the user account was created
    - `updated_at`: Date and time when the user account was last updated

### UserWorkspaceRole (Pivot Table)

- **Fields:**
    - `id`: Unique identifier
    - `user_id`: Identifier of the user (NOT NULL, Foreign key)
    - `workspace_id`: Identifier of the workspace (NOT NULL, Foreign key)
    - `role`: Role of the user in the workspace (Admin, Standard User)
    - `created_at`: Date and time when the association was created
    - `updated_at`: Date and time when the association was last updated

<br>

### Authentication & Authorization

Each user can perform CRUD operations only on their respective workspaces after authentication.
Token-based(JWT) authentication is implemented to authenticate users.

Users receive a token upon successful login, which they must use in API requests to access their workspaces.
CRUD operations on baskets are restricted to the authenticated user.

## Getting Started

### Prerequisites

- Golang
- Docker

### Setting up

Build and run the application using Docker Compose:

```shell
docker-compose -f deployment/docker-compose.yml up -d --build
```

Once the application is up and running, you can access it at http://localhost:8080 .


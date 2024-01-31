# Gin Bootstrapper

Welcome to the Gin Bootstrapper – your starting point for building Go web applications using the Gin framework!

## Introduction

Have you ever found yourself searching for a Gin project boilerplate or bootstrapper that matches your development style? I did, and that's why I created the Gin Bootstrapper. This project aims to provide you with a foundation for your Go web applications using Gin, tailored to my preferences. While it's opinionated and follows my development practices, I believe it can serve as a great starting point for other Go enthusiasts as well.

Feel free to use, modify, and adapt this bootstrapper to your liking. Let's start building awesome Go web applications together!

## Table of Contents

- [Gin Bootstrapper](#gin-bootstrapper)
  - [Introduction](#introduction)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
    - [Usage](#usage)
    - [Contributing](#contributing)
    - [License](#license)

## Features

- Easy-to-use Gin framework setup.
- User deserialisation middleware included, authentication to come.
- Database connection and ORM (GORM) preconfigured.
- Structured project layout for maintainability.
- API route examples for quick reference.
- Error handling and validation best practices.
- Under active development and improvement, might see major changes as my development style changes/improves.

## Getting Started

Follow the steps below to get started with this project.

### Prerequisites

Before you begin, ensure you have met the following requirements:

- [Go](https://golang.org/) installed on your system.
- [Git](https://git-scm.com/) for version control.
- A database server (e.g., PostgreSQL) set up and configured.

### Installation

1. Clone the repository:

```bash
git clone https://github.com/carboncody/gin-bootrapper.git
```

2. Change to the project directory:

```bash
cd enterprise-task-management
```

3. Install dependencies:

```bash
go mod download
```

3. Configure your environment variables by creating a .env file in the project root. You can use the .env.example file as a reference.

4. Run the project:

```bash
go run main.go
```

5. Access the server at http://localhost:8080.

### Usage

```plaintext



├── controllers/
│ ├── user.controller.go # User controller - all your controllers go here
│ └── ...
├── initializers/
│ ├── connectDB.go # Create a connection pool to DB
│ ├── loadEnv.go # Load environment variables
│ └── ...
├── middleware/
│ ├── deserialize-user.go # Using bearer token currently in the repo, ? to Auth0 soon
│ └── ...
├── migrate/
│ ├── migrate.go # Create a db migration when finished defining data model
│ └── ...
├── models/
│ ├── user.go # User model, define other models here
│ └── ...
├── routes/
│ ├── user.routes.go # User db method routes
│ └── ...
├── utils/  # Utils is a darkpit so try sorting your utils based on models or business logic
│ ├── ...
└── main.go # Main application file
```

### Contributing

Would love to take take some PRs

### License

This project is licensed under the MIT License.

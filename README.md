# Gin-Restful-Boilerplate

## Description

This repository provides a boilerplate for building a RESTful API using Go, Gin, GORM and Postgres. This boilerplate is designed to help you kickstart your API development with a well-structured project setup and common functionalities.

## Features

- **Gin Framework**: Utilizes the fast and lightweight Gin web framework to handle HTTP routing, middleware, and request handling efficiently.

- **GORM ORM**: Integrates GORM for database interaction, making it easier to work with databases and manage models in a structured manner.

- **CRUD Operations**: Includes basic CRUD (Create, Read, Update, Delete) operations for a sample resource, demonstrating how to set up database models and perform API actions.

- **Live reload**: Configured for live reload using **air** cli command to refresh the server upon code changes during development.

- **Testing**: Provides a comprehensive testing setup using Go's built-in testing package and **testify** for testing API endpoints and components.

- **Configuration**: Implements a configuration management system, allowing separation of environment-specific settings from the codebase.

## API Endpoints

The following API endpoints are available in this boilerplate:

- `GET /api/resource`: Fetch a list of resources.
- `GET /api/resource/:id`: Fetch details of a specific resource.
- `POST /api/resource`: Create a new resource.
- `PUT /api/resource/:id`: Update details of a specific resource.
- `DELETE /api/resource/:id`: Delete a specific resource.

## Contributing

Contributions to this boilerplate are welcome! Feel free to submit pull requests or open issues for any improvements, bug fixes, or additional features you'd like to see.

## License

This project is licensed under the [MIT License](LICENSE). You are free to use, modify, and distribute the code in this repository as long as you include the original license file.

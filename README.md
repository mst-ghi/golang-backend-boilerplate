<!-- @format -->

# Golang backend boilerplate

Welcome to the Golang Backend Boilerplate project! This boilerplate provides a solid foundation for building backend applications in Go, equipped with essential features to kickstart your development process.

## Technical Features

-   **Gin Framework**: Utilizes Gin, a high-performance HTTP web framework for building web applications in Go, providing robust routing and middleware capabilities.
-   **Gorm ORM**: Integrates Gorm, an ORM (Object-Relational Mapping) library for Go, simplifying database interactions and providing convenient query building.
-   **Swagger API Documentation**: Includes Swagger for documenting REST APIs, ensuring clear and comprehensive documentation for developers.
-   **Custom JWT Authentication**: Implements custom JWT (JSON Web Token) authentication for securing API endpoints and managing user authentication.
-   **Custom Commands**: Provides custom commands for various tasks, streamlining development processes and enhancing developer productivity.
-   **Middlewares**: Supports middleware functionality for extending the functionality of HTTP handlers, allowing for features like logging, authentication, and error handling.
-   **Base Response**: Includes a base response structure for consistent and standardized API responses, enhancing the clarity and maintainability of the codebase.
-   **SocketIO Integration**: Integrates SocketIO for real-time communication within the application, enabling seamless bidirectional communication between clients and the server.
-   **Scheduler**: Incorporates a scheduler for executing periodic tasks or background jobs, automating repetitive tasks and enhancing application functionality.

With these features, the Golang Backend Boilerplate offers a robust and flexible foundation for building backend applications in Go. Whether you're starting a new project or looking to enhance an existing one, this boilerplate provides the tools and structure to accelerate your development process.

## Clone and Run

```bash
-> go install github.com/swaggo/swag/cmd/swag@latest
-> git clone git@github.com:mst-ghi/app.git
-> cd ./app
-> cp .env.example .env #Update Database configuration
-> go mod download
-> go run . db:migrate
-> go run . serve
```

## Commands

-   **serve**: Starts the application server, allowing it to handle incoming requests and serve content to clients.
-   **swag**: Initiates the generation process for Swagger documentation, facilitating comprehensive API documentation for better understanding and integration.
-   **db:migrate**: Executes database table migrations, ensuring that the database structure is up-to-date with the latest changes in the application.
-   **db:seeder**: Inserts seed data into the database, providing initial data for development, testing, or demonstration purposes. This helps populate the database with predefined data for a smoother setup and testing process.

## Gin Framework [gin-gonic](https://gin-gonic.com/)

Our application leverages the Gin framework, a high-performance HTTP web framework for building web applications in Go. Gin provides a robust set of features, including routing, middleware support, and rendering utilities, making it ideal for developing scalable and efficient web services.
</br>

With Gin, we can easily define routes, handle requests, and manage middleware to streamline our application logic. Its minimalist design and fast performance make it a popular choice for building APIs and web applications in Go. We've utilized Gin to create a reliable and flexible foundation for our project, ensuring smooth development and high performance.

## API Documentation with Swagger [gin-swagger](https://github.com/swaggo/gin-swagger)

We utilize [gin-swagger](https://github.com/swaggo/gin-swagger) to comprehensively document our REST APIs.
To access the Swagger documentation, navigate to: **/api/docs/index.html**

## SocketIO Integration [googollee/go-socket.io](github.com/googollee/go-socket.io)

We've incorporated SocketIO to enable real-time communication within our application. This feature facilitates bidirectional communication between clients and the server, allowing for instant updates and interactions. With SocketIO, we can build dynamic and responsive applications that meet the demands of modern web development.
</br>

SocketIO is accessible via the URL path **/socket.io/\***, providing seamless integration for real-time features in our application. One of the available events is **user:get**, which allows the backend to retrieve user information. Additionally, the **user:me** event on the client side can be used to fetch details about the currently logged-in user. This setup ensures efficient communication and data exchange between the server and clients.

## Source structure

1. **core**:

    - **bootstrap**: Initializes the application.
    - **cmd**: Houses command-line interface (CLI) commands for tasks like database migration and seeding.
    - **config**: Manages application configuration settings.
    - **engine**: Initializes and configures the application engine.
    - **swagger**: Handles Swagger documentation setup.

2. **database**:
    - **db_drivers**: Contains database driver implementations.
    - **db_scopes**: Defines database query scopes.
    - **migrations**: Manages database schema migrations.
    - **models**: Defines database models/entities.
    - **repositories**: Contains repository implementations for database operations.
    - **seeder**: Manages seed data for populating the database.
    - **database.go**: Initializes the database connection and ORM.
3. docs:
    - **docs.go**: Initializes documentation generation.
    - **swagger.json/yaml**: Generated Swagger documentation files.
4. **internal**:

    - **middlewares**: Includes middleware functions for handling cross-origin resource sharing (CORS) and JSON Web Token (JWT) authentication.
    - **modules**: Contains app-specific logic and business logic.
        - **auth**: Handles authentication-related logic, including controllers, services, and routes.
        - **gateway**: Implements gateway logic and handlers.
        - **users**: Manages user-related logic, including controllers, services, and routes.
        - **routing.go**: Initializes application routes and routing logic.
    - **scheduler**: Manages scheduled tasks or jobs.
    - **kernel.go**: Initializes the application kernel.
    - **provider.go**: Manages dependency injection and service providers.

5. **pkg**:
    - **handlers**: Provides error handling utilities.
    - **helpers**: Contains helper functions for common tasks like data conversion and cryptography.
    - **messages**: Defines message constants and utilities for message formatting.
    - **validation**: Manages validation logic and error messages.
6. **go.mod/go.sum**: Define the project's module and dependencies.
7. **main.go**: Main entry point for the application.
8. **README.md**: Documentation file providing essential information about the project.

## Author

**[Mostafa Gholami](https://mst-ghi.github.io/)**


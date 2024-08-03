# Simple CRUD in Go

This project is a simple CRUD (Create, Read, Update, Delete) developed in Go. It uses a series of modern technologies and practices to ensure a robust and high-quality application.

## Technologies Used

-   **Go**: Main language for application development.
-   **MongoDB**: NoSQL database for data storage.
-   **Docker**: Containerization of the application to facilitate development and deployment.
-   **JWT (JSON Web Tokens)**: Secure authentication and authorization.
-   **Environment Variables**: Flexible application configuration.
-   **Data Entry Validation**: Guarantee that the data received is valid and safe.
-   **SOLID Principles**: Application of design principles to ensure a clean and scalable code structure.
-   **Error and Log Handling**: Implementation of a robust strategy to deal with errors and record important information about the execution of the application.

## Project Structure

-   **`/cmd`**: Contains the application initialization code.
-   **`/internal/service`**: Contains service implementations, following SOLID principles.
-   **`/internal/model`**: Contains the data models and data structures.
-   **`/internal/handler`**: Contains the route handlers and API endpoint implementations.
-   **`/pkg`**: Contains reusable packages such as data validation and helper functions.
-   **`/docker`**: Files related to Docker configuration.
-   **`/config`**: Configuration files, such as environment variables.
-   **`/logs`**: Configuration and strategies for logging.

## Settings

1. **Clone the repository:**

    ```bash
     git clone <REPOSITORY_URL>
     cd <REPOSITORY_NAME>

    ```

2. **Configure the environment variables:**

    Create a .env file in the project root and add the following variables:

    ```env
        MONGO_URI=mongodb://localhost:27017
        JWT_SECRET=secret_key

        LOG_OUTPUT=
        LOG_LEVEL=

        MONGODB_URL=
        MONGODB_USER=
        MONGODB_USER_COLLECTION=

        JWT_SECRET=
    ```

3. **Build and run Docker:**

    To build the Docker image:

    ```bash
    docker build -t image-name .
    ```

4. **Run the project locally (without Docker):**

    Install dependencies:

    ```bash
    go mod download
    ```

    Run the application:

    ```bash
    go run main.goS
    ```

# Endpoints

1. POST /login: Logs a user in and returns a JWT.
2. GET /getUserById/
   : Retrieves information for a specific user by ID. Requires JWT authentication.
3. GET /getUserByEmail/
   : Retrieves information from a specific user via email. Requires JWT authentication.
4. POST /createUser: Creates a new user.
5. PUT /updateUser/
   : Updates information for a specific user by ID. Requires JWT authentication.
6. DELETE /deleteUser/
   : Deletes a specific user by ID. Requires JWT authentication.

7. \*GET /swagger/any: Accesses the Swagger API documentation.

# cs-exp-go-api

cs-exp-go-api is a Go-based REST API for managing users with JWT authentication. This project demonstrates how to build a secure API with user management functionalities, including registration, login, and CRUD operations with JWT authentication.

## Table of Contents
1. Prerequisites 
2. Setup and Installation 
3. Configuration 
4. Running the Project 
5. API Endpoints 
6. Testing with cURL 
7. Troubleshooting 
8. License
### Prerequisites
Before you begin, ensure you have the following installed:
- Docker 
- Docker Compose
- (for local development, if not using Docker)
### Setup and Installation
1. Clone the Repository:
        git clone https://github.com/yourusername/cs-exp-go-api.git
            cd cs-exp-go-api
2.  Build and Start the Containers:
        docker-compose up --build
    This will build and start the application and PostgreSQL database using Docker Compose.

3.	Database Initialization:
    Ensure that init.sql is properly set up in the docker-compose.yml file under the db service to initialize the database schema.

### Configuration

1.	Environment Variables:
        Create a .env file in the root directory with the following content:

			SER=postgres
           	PASSWORD=postgres
           	DB_NAME=mydb
           	DB_HOST=db
           	DB_PORT=5432
           	JWT_SECRET=your_jwt_secret
           
       Replace your_jwt_secret with a secure secret key for JWT authentication.

2.	Docker Configuration:
        Ensure the docker-compose.yml file is configured with correct database and application settings. It should look something like this:

### Running the Project

1.	Start the Application:
        Run the following command to start the application:
        
            docker-compose up

2.	Access the API:
    The API will be available at http://localhost:8989.

    API Endpoints
    
    Public Routes

	•	Register User
        http:
        
            POST /register

    Request Body:

        {
            "name": "John Doe",
            "email": "john@example.com",
            "username": "johndoe",
            "password": "password123"
        }
    
    Login User:

        POST /login

    Request Body:

        {
            "username": "johndoe",
            "password": "password123"
        }

    Response:

        {
            "user": {
                "id": 1,
                "name": "John Doe",
                "email": "john@example.com",
                "username": "johndoe"
            },
            "token": "your_jwt_token"
        }

    Protected Routes
    - Get All Users

            GET /users

        Headers:

            Authorization: Bearer <your_jwt_token>

    - Create User

            POST /users

         Request Body:

            {
                "name": "Jane Doe",
                "email": "jane@example.com",
                "username": "janedoe",
                "password": "password123"
            }

        Headers:

            Authorization: Bearer <your_jwt_token>

    - Update User

            PUT /users

        Request Body:

            {
                "id": 1,
                "name": "Jane Doe Updated",
                "email": "jane.updated@example.com",
                "username": "janedoe"
            }
        Headers

            Authorization: Bearer <your_jwt_token>

    - Delete User

            DELETE /users/{id}

        Headers:

            Authorization: Bearer <your_jwt_token>

### Testing with cURL

#### Register User
bash:

            curl --location 'http://localhost:8989/register' \
            --header 'Content-Type: application/json' \
            --data-raw '{
                "name": "John Doe",
                "email": "john@example.com",
                "username": "johndoe",
                "password": "password123"
            }'

#### Login User
    bash:
    curl --location 'http://localhost:8989/login' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "username": "johndoe",
        "password": "password123"
    }'
#### Get All Users
    bash:
    curl --location 'http://localhost:8989/users' \
    --header 'Authorization: Bearer <your_jwt_token>'
#### Create User
    bash:
    curl --location 'http://localhost:8989/users' \
    --header 'Authorization: Bearer <your_jwt_token>' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "name": "Jane Doe",
        "email": "jane@example.com",
        "username": "janedoe",
        "password": "password123"
    }'
#### Update User
    bash:
    curl --location 'http://localhost:8989/users' \
    --header 'Authorization: Bearer <your_jwt_token>' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "id": 1,
        "name": "Jane Doe Updated",
        "email": "jane.updated@example.com",
        "username": "janedoe"
    }'
#### Delete User
    bash:
            
        curl --location 'http://localhost:8989/users/{id}' \
        --header 'Authorization: Bearer <your_jwt_token>' \
        --request DELETE

### Troubleshooting
- Invalid Token Errors: Ensure you are using a valid JWT token generated during login. Verify that the token is correctly set in the Authorization header.

- Database Connection Issues: Check your database settings and ensure the database is running.

### License

This project is licensed under the MIT License. See the LICENSE file for details.

# API-BACKEND-CHALLENGE

This application allows managers and technicians to keep track of tasks.

## Requirements
*  Go 1.19 or higher. If you don't have Golang installed, you can download it from https://go.dev/doc/install)
*  Docker Compose (https://docs.docker.com/compose/install/)

## Setup

1. Install the project dependencies:

   ```
   make install
   ```
   
2. Start local environment:

   ```
   make local-up
   ```

3. Start the project:

   ```
   make start
   ```

   The project should now be running at http://localhost:8088.


4. Run unit tests:

   ```
   make test
   ```
   
5. Run test coverage:

   ```
   make cover
   ```

## Usage

### Endpoints

The API has the following endpoints:

------------------

#### User

Create a user.
Required fields:
* username
* password
* role (manager, technician)

##### Request

```
POST /v1/user
Content-Type: application/json

{
    "username": "john.doe",
    "password": "123456",
    "role": "manager"
 }
```

##### Response

```
HTTP/1.1 201 Created
Content-Type: application/json

{
    "meta": {
        "offset": 0,
        "limit": 0,
        "record_count": 1
    },
    "records": [
        {
            "id": "125edfd8-a7bd-4ddd-8429-278e935def82",
            "username": "john.doee",
            "role": "manager"
        }
    ]
}

```
------------------
#### Login

Login with a user.
Required fields:
* username
* password

##### Request

```
POST /v1/user/login
Content-Type: application/json

{
    "username": "john.doe",
    "password": "123456",
}
```

##### Response

```
HTTP/1.1 201 Created
Content-Type: application/json

{
    "meta": {
        "offset": 0,
        "limit": 0,
        "record_count": 1
    },
    "records": [
        {
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJ1c2VyIjoiam9obi5kb2UifQ.cqDH000_9wpwtp2pmrAgUmcPvlyNDxObz8ks6ohBiUU"
        }
    ]
}
```

------------------

#### Task

Create a task.
Required fields:
* summary
* performed_at

##### Request

```
POST /v1/task
Content-Type: application/json

{
    "summary":"12345",
    "performed_at": "2023-04-27T08:05:14Z"
}
```

##### Response

```
HTTP/1.1 201 Created
Content-Type: application/json
Authorization: Bearer <token>
{
    "meta": {
        "offset": 0,
        "limit": 0,
        "record_count": 1
    },
    "records": [
        {
            "id": "ae92d1d3-70f5-49c5-8528-5643d48f9738",
            "summary": "12345",
            "performed_at": "2023-04-27T08:05:14Z"
        }
    ]
}
```

------------------

#### TASKS

List all task for a given user. It will evaluate role as:
* Technicians will be able to see only tasks that they performed
* Manager will see all tasks

##### Request

```
GET /v1/task
```

##### Response

```
HTTP/1.1 201 Created
Content-Type: application/json
Authorization: Bearer <token>
{
    "meta": {
        "offset": 0,
        "limit": 0,
        "record_count": 1
    },
    "records": [
        {
            "id": "ae92d1d3-70f5-49c5-8528-5643d48f9738",
            "summary": "12345",
            "performed_at": "2023-04-27T08:05:14Z"
        }
    ]
}
```

### Error Handling

If an error occurs, the API will return a JSON object with an error message:

```
{
    "developer_message": string,
    "user_message": string,
    "status_code": int
}
```

Possible HTTP status codes for errors include:

- `400 Bad Request` for invalid request data
- `401 Unauthorized` missing a valid authorization token
- `409 Unauthorized` username already present in the database
- `500 Internal Server Error` for server-side errors

## Swagger API Documentation

This project uses Swagger for API documentation. Swagger provides a user-friendly interface for exploring and testing the API.

To access the Swagger page:

1. Start the application if it's not already running.
2. Open a web browser and navigate to `http://localhost:8088/v1/swagger/index.html#/`.
3. The Swagger page should load, displaying a list of available endpoints.

From here, you can explore the available endpoints, see what parameters they require, and test them out.

If you have any questions or issues with the Swagger page, please refer to the API documentation or contact the project maintainers.
# Todo REST API 

This is a Golang application with a postgresDb data store providing Restful routes for a TODO app supporting Tasks and TaskCategories. 

It can be built and deployed with simple Docker commands. Docker must be installed on your local machine for this to work. 

After building and running the containers with docker, your local docker environment will contain a `Tasks` and `TaskCategories` table in a postgres server exposed on port 5432 as well as a RestApi exposed on port 8080. Postgres credentials are stored in the .env file.

The database will have some initial seed data and so making a GET to  http://localhost:8080/tasks or http://localhost:8080/taskCategories will result in 200 responses with data returned.

## Build Docker Image
`docker compose build`

## Build and run docker containers
`docker compose up`

NOTE: localhost:8080 may not be exposed after the very first execution of `docker compose up`. If this is the case please simply kill the server and run `docker compose up` again.

## REST API
Provides CRUD actions on Tasks and TaskCategories.

## Task Categories
### Create a new TaskCategory
`POST http://localhost:8080/taskCategories`
#### Example Request Body:

```
{
    "task_category_name": "Groceries",
}
```

### Get list of all TaskCategories
`GET localhost:8080/taskCategories`

### Get TaskCategory
`GET http://localhost:8080/taskCategories/:id`

### Update TaskCategory
`PUT http://localhost:8080/taskCategories/:id`
#### Example Request Body:

```
{
    "task_category_name": "House Chores",
}
```

### Delete a Task
`DELETE /taskCategories/:id`

## Tasks
### Create a new Task
`POST http://localhost:8080/tasks`
#### Example Request Body:

```
{
    "task_name": "Tomatoes",
    "task_description": "Get cherry tomatoes",
    "task_category_id": 4,
    "task_is_complete": false
}
```

### Get list of all Tasks
`GET localhost:8080/tasks`

### Get task
`GET http://localhost:8080/tasks/:id`

### Update task
`PUT http://localhost:8080/tasks/:id`
#### Example Request Body:

```
{
    "task_name": "Tomatoes",
    "task_description": "Get cherry tomatoes",
    "task_is_complete": true,
    "task_category_id": 4
}
```

### Delete a Task
`DELETE /tasks/:id`

## Troubleshooting
1. The first attempt at running the application with `docker compose up` may result in only booting the database server but not the Rest Api. If this is the case please terminate the server and run `docker compose up` again. This should result in both the database and rest api server booting up successfully. 

2. Database tables not creating or seeding properly: run the following from root of project `docker-compose down --volumes`

## Future Improvements
- Add system/integration tests to validate controller behavior
- Add logging
- Implement HATEOS in response headers
- Add swagger or swagger equivalent
- Add support for multiple users
- Deploy to a free cloud service
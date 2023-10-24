# Todo REST API 

This is a Golang application with a postgresDb data store providing Restful routes for a TODO app supporting Tasks and TaskCategories. It can be built and deployed with simple Docker commands. Docker must be installed on your local machine for this to work.

## Build Docker Image
    docker compose build

## Build and run docker containers
    docker compose up

## REST API
Allows CRUD actions on Tasks and TaskCategories.

## Task Categories
### Create a new TaskCategory
`POST http://localhost:8080/taskCategories`

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
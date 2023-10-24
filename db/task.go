package db

import (
    "database/sql"
    "github.com/rwas2505/go-chi/models"
)

func (db Database) GetAllTasks() (*models.TaskList, error) {
    list := &models.TaskList{}

    rows, err := db.Conn.Query("SELECT * FROM tasks ORDER BY TaskId DESC")

    if err != nil {
        return list, err
    }

    for rows.Next() {
        var task models.Task

        err := rows.Scan(&task.TaskId, &task.TaskName, &task.TaskDescription, &task.TaskCreatedAt, &task.TaskIsComplete)

        if err != nil {
            return list, err
        }

        list.Tasks = append(list.Tasks, task)
    }

    return list, nil
}

func (db Database) AddTask(task *models.Task) error {
    var taskId int
    var taskCreatedAt string
    var taskIsComplete bool

    query := `INSERT INTO tasks (taskname, taskdescription) VALUES ($1, $2) RETURNING taskid, taskcreatedat, taskiscomplete`

    err := db.Conn.QueryRow(query, task.TaskName, task.TaskDescription).Scan(&taskId, &taskCreatedAt, &taskIsComplete)

    if err != nil {
        return err
    }

    task.TaskId = taskId
    task.TaskCreatedAt = taskCreatedAt
    task.TaskIsComplete= taskIsComplete

    return nil
}

func (db Database) GetTaskById(taskId int) (models.Task, error) {
    task := models.Task{}

    query := `SELECT * FROM tasks WHERE taskid = $1;`

    row := db.Conn.QueryRow(query, taskId)

    switch err := row.Scan(&task.TaskId, &task.TaskName, &task.TaskDescription, &task.TaskCreatedAt, &task.TaskIsComplete); err {
    case sql.ErrNoRows:
        return task, ErrNoMatch
    default:
        return task, err
    }
}

func (db Database) DeleteTask(taskId int) error {
    query := `DELETE FROM tasks WHERE taskId = $1;`

    _, err := db.Conn.Exec(query, taskId)

    switch err {
    case sql.ErrNoRows:
        return ErrNoMatch
    default:
        return err
    }
}

func (db Database) UpdateTask(taskId int, taskData models.Task) (models.Task, error) {
    task := models.Task{}

    query := `UPDATE tasks SET taskName=$1, taskDescription=$2, taskIsComplete=$3 WHERE taskId=$4 RETURNING taskId, taskName, taskDescription, taskIsComplete, taskCreatedAt;`

    err := db.Conn.QueryRow(query, taskData.TaskName, taskData.TaskDescription, taskData.TaskIsComplete, taskId).Scan(&task.TaskId, &task.TaskName, &task.TaskDescription, &task.TaskIsComplete, &task.TaskCreatedAt)

    if err != nil {
        if err == sql.ErrNoRows {
            return task, ErrNoMatch
        }
        return task, err
    }
	
    return task, nil
}
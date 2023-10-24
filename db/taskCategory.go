package db

import (
    "database/sql"
    "github.com/rwas2505/go-chi/models"
)

func (db Database) GetAllTaskCategories() (*models.TaskCategories, error) {
    list := &models.TaskCategories{}

    rows, err := db.Conn.Query("SELECT * FROM taskCategories ORDER BY TaskCategoryId DESC")

    if err != nil {
        return list, err
    }

    for rows.Next() {
        var taskCategory models.TaskCategory

        err := rows.Scan(&taskCategory.TaskCategoryId, &taskCategory.TaskCategoryName)

        if err != nil {
            return list , err
        }

        list.TaskCategories = append(list.TaskCategories, taskCategory)
    }

    return list, nil
}

func (db Database) AddTaskCategory(taskCategory *models.TaskCategory) error {
    var taskCategoryId int

    query := `INSERT INTO taskCategories (taskCategoryName) VALUES ($1) RETURNING taskCategoryId`

    err := db.Conn.QueryRow(query, taskCategory.TaskCategoryName).Scan(&taskCategoryId)

    if err != nil {
        return err
    }
 
    taskCategory.TaskCategoryId = taskCategoryId

    return nil
}

func (db Database) GetTaskCategoryById(taskCategoryId int) (models.TaskCategory, error) {
    taskCategory := models.TaskCategory{}

    query := `SELECT * FROM taskCategories WHERE taskCategoryId = $1;`

    row := db.Conn.QueryRow(query, taskCategoryId)

    switch err := row.Scan(&taskCategory.TaskCategoryId, &taskCategory.TaskCategoryName); err {
    case sql.ErrNoRows:
        return taskCategory, ErrNoMatch
    default:
        return taskCategory, err
    }
}

func (db Database) DeleteTaskCategory(taskCategoryId int) error {
    query := `DELETE FROM taskCategories WHERE taskCategoryId = $1;`

    _, err := db.Conn.Exec(query, taskCategoryId)

    switch err {
    case sql.ErrNoRows:
        return ErrNoMatch
    default:
        return err
    }
}

func (db Database) UpdateTaskCategory(taskCategoryId int, taskCategoryData models.TaskCategory) (models.TaskCategory, error) {
    taskCategory := models.TaskCategory{}

    query := `UPDATE taskCategories SET taskCategoryName=$1 WHERE taskCategoryId=$2 RETURNING taskCategoryId, taskCategoryName;`

    err := db.Conn.QueryRow(query, taskCategoryData.TaskCategoryName, taskCategoryData.TaskCategoryId).Scan(&taskCategory.TaskCategoryId, &taskCategory.TaskCategoryName)

    if err != nil {
        if err == sql.ErrNoRows {
            return taskCategory, ErrNoMatch
        }
        return taskCategory, err
    }
	
    return taskCategory, nil
}
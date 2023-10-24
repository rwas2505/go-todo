package models

import (
    "fmt"
    "net/http"
)

type TaskCategory struct {
    TaskCategoryId int `json:"task_category_id"`
    TaskCategoryName string `json:"task_category_name"`
}

type TaskCategories struct {
    TaskCategories []TaskCategory `json:"task_categories"`
}

func (i *TaskCategory) Bind(r *http.Request) error {
    if i.TaskCategoryName == "" {
        return fmt.Errorf("task_category_name is a required field")
    }
    return nil
}

func (*TaskCategories) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}

func (*TaskCategory) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}
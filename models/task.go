package models

import (
    "fmt"
    "net/http"
)

type Task struct {
    TaskId int `json:"task_id"`
    TaskName string `json:"task_name"`
    TaskDescription string `json:"task_description"`
    TaskCreatedAt string `json:"task_created_at"`
    TaskIsComplete bool `json:"task_is_complete"`
}

type TaskList struct {
    Tasks []Task `json:"tasks"`
}

func (i *Task) Bind(r *http.Request) error {
    if i.TaskName == "" {
        return fmt.Errorf("task_name is a required field")
    }
    return nil
}

func (*TaskList) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}

func (*Task) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}
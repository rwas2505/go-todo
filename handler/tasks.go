package handler

import (
    "context"
    "fmt"
    "net/http"
    "strconv"
    "github.com/go-chi/chi"
    "github.com/go-chi/render"
    "github.com/rwas2505/go-chi/db"
    "github.com/rwas2505/go-chi/models"
)

var taskIdKey = "taskId"

func tasks(router chi.Router) {
    router.Get("/", getAllTasks)
    router.Post("/", createTask)
    router.Route("/{taskId}", func(router chi.Router) {
        router.Use(TaskContext)
        router.Get("/", getTask)
        router.Put("/", updateTask)
        router.Delete("/", deleteTask)
    })
}

func TaskContext(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        taskId := chi.URLParam(r, "taskId")

        if taskId == "" {
            render.Render(w, r, ErrorRenderer(fmt.Errorf("task ID is required")))
            return
        }

        id, err := strconv.Atoi(taskId)

        if err != nil {
            render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid task ID")))
        }

        ctx := context.WithValue(r.Context(), taskIdKey, id)

        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func getAllTasks(w http.ResponseWriter, r *http.Request) {
    tasks, err := dbInstance.GetAllTasks()

    if err != nil {
        render.Render(w, r, ServerErrorRenderer(err))
        return
    }

    if err := render.Render(w, r, tasks); err != nil {
        render.Render(w, r, ErrorRenderer(err))
    }
}

func getTask(w http.ResponseWriter, r *http.Request) {
    taskId := r.Context().Value(taskIdKey).(int)

    task, err := dbInstance.GetTaskById(taskId)

    if err != nil {
        if err == db.ErrNoMatch {
            render.Render(w, r, ErrNotFound)
        } else {
            render.Render(w, r, ErrorRenderer(err))
        }
        return
    }

    if err := render.Render(w, r, &task); err != nil {
        render.Render(w, r, ServerErrorRenderer(err))
        return
    }
}

func createTask(w http.ResponseWriter, r *http.Request) {
    task := &models.Task{}

    if err := render.Bind(r, task); err != nil {
        render.Render(w, r, ErrorRenderer(err))
        return
    }

    if err := dbInstance.AddTask(task); err != nil {
        render.Render(w, r, ErrorRenderer(err))
        return
    }

    if err := render.Render(w, r, task); err != nil {
        render.Render(w, r, ServerErrorRenderer(err))
        return
    }
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	taskId := r.Context().Value(taskIdKey).(int)

	taskData := models.Task{}

    // If the request body cannot bind to the Task object, render an error
	if err := render.Bind(r, &taskData); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	task, err := dbInstance.UpdateTask(taskId, taskData)

	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrorNotFound(err, taskId))
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}

	if err := render.Render(w, r, &task); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
    taskId := r.Context().Value(taskIdKey).(int)
	
    err := dbInstance.DeleteTask(taskId)
	
    if err != nil {
        if err == db.ErrNoMatch {
            render.Render(w, r, ErrNotFound)
        } else {
            render.Render(w, r, ServerErrorRenderer(err))
        }
        return
    }
}
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

var taskCategoryIdKey = "taskCategoryId"

func taskCategories(router chi.Router) {
    router.Get("/", getAllTaskCategories)
    router.Post("/", createTaskCategory)
    router.Route("/{taskCategoryId}", func(router chi.Router) {
        router.Use(TaskCategoryContext)
        router.Get("/", getTaskCategory)
        router.Put("/", updateTaskCategory)
        router.Delete("/", deleteTaskCategory)
    })
}

func TaskCategoryContext(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        taskCategoryId := chi.URLParam(r, "taskCategoryId")

        if taskCategoryId == "" {
            render.Render(w, r, ErrorRenderer(fmt.Errorf("taskCategory ID is required")))
            return
        }

        id, err := strconv.Atoi(taskCategoryId)

        if err != nil {
            render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid taskCategory ID")))
        }

        ctx := context.WithValue(r.Context(), taskCategoryIdKey, id)

        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func getAllTaskCategories(w http.ResponseWriter, r *http.Request) {
    taskCategories, err := dbInstance.GetAllTaskCategories()

    if err != nil {
        render.Render(w, r, ServerErrorRenderer(err))
        return
    }

    if err := render.Render(w, r, taskCategories); err != nil {
        render.Render(w, r, ErrorRenderer(err))
    }
}

func getTaskCategory(w http.ResponseWriter, r *http.Request) {
    taskCategoryId := r.Context().Value(taskCategoryIdKey).(int)

    taskCategory, err := dbInstance.GetTaskCategoryById(taskCategoryId)

    if err != nil {
        if err == db.ErrNoMatch {
            render.Render(w, r, ErrNotFound)
        } else {
            render.Render(w, r, ErrorRenderer(err))
        }
        return
    }

    if err := render.Render(w, r, &taskCategory); err != nil {
        render.Render(w, r, ServerErrorRenderer(err))
        return
    }
}

func createTaskCategory(w http.ResponseWriter, r *http.Request) {
    taskCategory := &models.TaskCategory{}

    if err := render.Bind(r, taskCategory); err != nil {
        render.Render(w, r, ErrorRenderer(err))
        return
    }

    if err := dbInstance.AddTaskCategory(taskCategory); err != nil {
        render.Render(w, r, ErrorRenderer(err))
        return
    }

    if err := render.Render(w, r, taskCategory); err != nil {
        render.Render(w, r, ServerErrorRenderer(err))
        return
    }
}

func updateTaskCategory(w http.ResponseWriter, r *http.Request) {
	taskCategoryId := r.Context().Value(taskCategoryIdKey).(int)

	taskCategoryData := models.TaskCategory{}

    // If the request body cannot bind to the TaskCategory object, render an error
	if err := render.Bind(r, &taskCategoryData); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	taskCategory, err := dbInstance.UpdateTaskCategory(taskCategoryId, taskCategoryData)

	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrorNotFound(err))
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}

	if err := render.Render(w, r, &taskCategory); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deleteTaskCategory(w http.ResponseWriter, r *http.Request) {
    taskCategoryId := r.Context().Value(taskCategoryIdKey).(int)
	
    err := dbInstance.DeleteTaskCategory(taskCategoryId)
	
    if err != nil {
        if err == db.ErrNoMatch {
            render.Render(w, r, ErrNotFound)
        } else {
            render.Render(w, r, ServerErrorRenderer(err))
        }
        return
    }
}
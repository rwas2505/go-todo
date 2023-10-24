package handler

import (
    "net/http"
    "github.com/go-chi/chi"
    "github.com/go-chi/render"
    "github.com/rwas2505/go-chi/db"
)

var dbInstance db.Database

func NewHandler(db db.Database) http.Handler {
    router := chi.NewRouter()
    dbInstance = db
    router.MethodNotAllowed(methodNotAllowedHandler)
    router.NotFound(notFoundHandler)
    router.Route("/tasks", tasks) //root for tasks, children defined in handler.tasks.go
    return router
}
func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")
    w.WriteHeader(405)
    render.Render(w, r, ErrMethodNotAllowed)
}
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")
    w.WriteHeader(400)
    render.Render(w, r, ErrNotFound)
}
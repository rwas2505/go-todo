package main

import (
    "context"
    "fmt"
    "github.com/rwas2505/go-chi/db"
    "github.com/rwas2505/go-chi/handler"
    "log"
    "net"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    addr := ":8080"
    listener, err := net.Listen("tcp", addr)

    if err != nil {
        log.Fatalf("Error occurred: %s", err.Error())
    }

    dbUser, dbPassword, dbName :=
        os.Getenv("POSTGRES_USER"),
        os.Getenv("POSTGRES_PASSWORD"),
        os.Getenv("POSTGRES_DB")

    database, err := db.Initialize(dbUser, dbPassword, dbName)

    if err != nil {
        log.Fatalf("Could not set up database: %v", err)
    }

	// Ensure that the database connection is kept alive while the application is running
    defer database.Conn.Close()

    httpHandler := handler.NewHandler(database)

    server := &http.Server{
        Handler: httpHandler,
    }

	// The API server is started on a separate goroutine
    go func() {
        server.Serve(listener)
    }()

    defer Stop(server)

    log.Printf("Started server on %s", addr)

    ch := make(chan os.Signal, 1)

	// Server keeps running until it receives a SIGINT or SIGTERM signal 
	// After which it calls the Stop function to clean up and shut down the server.
    signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

    log.Println(fmt.Sprint(<-ch))

    log.Println("Stopping API server.")
}

func Stop(server *http.Server) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

    defer cancel()

    if err := server.Shutdown(ctx); err != nil {
        log.Printf("Could not shut down server correctly: %v\n", err)
        os.Exit(1)
    }
}
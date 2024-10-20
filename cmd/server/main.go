package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"log/slog"
	"net/http"
	jobcontroller "rollout/internal/controller/job"
	jobdomain "rollout/internal/domain/job"
	jobrepo "rollout/internal/repository/job"
	jobusecase "rollout/internal/usecase/job"
)

func main() {
	// Initialise the database connection
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		slog.Error("could not create database connection: %v", err)
	}

	// Migrate the database schema. TODO this is not production-ready.
	err = db.AutoMigrate(&jobdomain.Job{})
	if err != nil {
		slog.Error("could not migrate database: %v", err)
	}

	// Initialise repositories
	jobRepo := jobrepo.NewRepository(db)

	// Initialise use cases
	createJobUseCase := jobusecase.NewCreate(jobRepo)

	// Initialise controllers
	jobCtrl := jobcontroller.NewController(createJobUseCase)

	// Set up the router
	router := mux.NewRouter()

	// Register endpoints and hanlders
	router.HandleFunc("/jobs", jobCtrl.CreateJob).Methods("POST")

	// Start server
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

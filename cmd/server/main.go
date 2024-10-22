package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"log/slog"
	"net/http"
	foldercontroller "rollout/internal/controller/folder"
	jobcontroller "rollout/internal/controller/job"
	folderrepo "rollout/internal/repository/folder"
	jobrepo "rollout/internal/repository/job"
	folderusecase "rollout/internal/usecase/folder"
	jobusecase "rollout/internal/usecase/job"
)

func main() {
	// Initialise the database connection
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		slog.Error("could not create database connection: %v", err)
	}

	// Migrate the database schema. TODO this is not production-ready.
	err = db.AutoMigrate(&folderrepo.Folder{}, &jobrepo.Job{})
	if err != nil {
		slog.Error("could not migrate database: %v", err)
	}

	// Initialise repositories
	folderRepo := folderrepo.NewRepository(db)
	jobRepo := jobrepo.NewRepository(db)

	// Initialise use cases
	createFolderUseCase := folderusecase.NewCreate(folderRepo)
	createJobUseCase := jobusecase.NewCreate(jobRepo)

	// Initialise controllers
	// TODO there could be lots of use cases per controller so we possibly need a better pattern
	folderCtrl := foldercontroller.NewController(createFolderUseCase)
	jobCtrl := jobcontroller.NewController(createJobUseCase)

	// Set up the router
	router := mux.NewRouter()

	// Register endpoints and handlers
	// TODO REST naming conventions
	router.HandleFunc("/folders", folderCtrl.CreateFolder).Methods("POST")
	router.HandleFunc("/jobs", jobCtrl.CreateJob).Methods("POST")

	// Start server
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

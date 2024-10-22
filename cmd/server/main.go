package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"log/slog"
	"rollout/api/http"
	foldercontroller "rollout/internal/controller/folder"
	folderrepo "rollout/internal/repository/folder"
	jobrepo "rollout/internal/repository/job"
	folderusecase "rollout/internal/usecase/folder"
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
	//jobRepo := jobrepo.NewRepository(db)

	// Initialise use cases
	createFolderUseCase := folderusecase.NewCreate(folderRepo)
	//createJobUseCase := jobusecase.NewCreate(jobRepo)

	// Initialise controllers
	// TODO there could be lots of use cases per controller so we possibly need a better pattern
	folderCtrl := foldercontroller.NewController(createFolderUseCase)
	//jobCtrl := jobcontroller.NewController(createJobUseCase)

	// Initialise HTTP server
	server := http.NewServer(folderCtrl)

	// Start server
	if err := server.Start(8080); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

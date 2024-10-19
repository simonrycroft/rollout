package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log/slog"
	"rollout/internal/jobrepository"
	"rollout/internal/jobservice"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		slog.Error("could not create database connection: %v", err)
	}

	err = db.AutoMigrate(&jobrepository.Job{})
	if err != nil {
		slog.Error("could not migrate database: %v", err)
	}

	jobRepo := jobrepository.NewJobRepository(db)
	jobService := jobservice.NewJobService(jobRepo)

	job, err := jobService.CreateJob("My First Job")
	if err != nil {
		slog.Error("could not create job: %v", err)
	}

	fmt.Println("New Job Created")
	fmt.Printf("%#v\n", job)
}

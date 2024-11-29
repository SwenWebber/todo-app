package main

import (
	"log"
	"time"

	"github.com/swenwebber/todo-app/config"
	"github.com/swenwebber/todo-app/internal/handler"
	"github.com/swenwebber/todo-app/internal/model"
	"github.com/swenwebber/todo-app/internal/repository/db"
	"github.com/swenwebber/todo-app/internal/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/swenwebber/todo-app/internal/service"
)

func main() {
	//database setup
	dbConfig := config.NewDBConfig()
	dsn := dbConfig.GetDSN()
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection error")
	}
	database.AutoMigrate(&model.Task{})

	repoDB, err_ := db.NewDBTaskRepo(database)

	if err_ != nil {
		log.Printf(err_.Error())
	}

	//in memory storage with map[id]Tasks
	//repo := repository.NewMemoryTaskRepository()
	//service := service.NewTaskService(repo)
	service := service.NewTaskService(repoDB)
	api_handler := handler.NewTaskHandler(service)
	templateHandler := handler.NewTemplateHandler(service)

	config := &server.Config{
		Port:         "8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	srv := server.NewServer(api_handler, templateHandler, config)

	if err := srv.Run(); err != nil {
		log.Fatalf("Server error: %v", err)
	}

}

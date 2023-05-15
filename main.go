package main

import (
	"log"
	activityHandler "todolistapi/activity/handler/api"
	activityRepository "todolistapi/activity/repository"
	activityUsecase "todolistapi/activity/usecase"
	"todolistapi/database"
	"todolistapi/router"
	todoHandler "todolistapi/todo/handler/api"
	todoRepository "todolistapi/todo/repository"
	todoUsecase "todolistapi/todo/usecase"

	"github.com/go-gormigrate/gormigrate/v2"
)

func main() {
	db := database.InitDB()

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		&database.V2023051500001,
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	todoRepository := todoRepository.NewTodoRepository(db)
	activityRepository := activityRepository.NewActivityRepository(db)

	activityUsecase := activityUsecase.NewActivityUsecase(activityRepository)
	todoUsecase := todoUsecase.NewTodoUsecase(todoRepository)

	handler := router.BuildHandler(
		activityHandler.NewActivityHandler(activityUsecase),
		todoHandler.NewTodoHandler(todoUsecase),
	)
	router.RunServer(handler)
}

package main

import (
	activityHandler "todolistapi/activity/handler/api"
	activityRepository "todolistapi/activity/repository"
	activityUsecase "todolistapi/activity/usecase"
	"todolistapi/database"
	"todolistapi/router"
	todoHandler "todolistapi/todo/handler/api"
	todoRepository "todolistapi/todo/repository"
	todoUsecase "todolistapi/todo/usecase"
)

func main() {
	db := database.InitDB()

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

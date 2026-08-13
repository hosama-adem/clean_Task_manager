package main

import (
	"task_manager/Delivery/controllers"
	"task_manager/Delivery/routers"
	"task_manager/Repositories"
	"task_manager/Usecases"
)

func main() {
	taskrepository := Repositories.NewTaskStore()
	userrepository := Repositories.NewUserStore()

	taskusecase := Usecases.NewTaskUsecase(taskrepository)
	userusecase := Usecases.NewUserUsecase(userrepository)

	controller := controllers.NewTaskController(taskusecase, userusecase)

	r := routers.Router(controller)
	r.Run()

}

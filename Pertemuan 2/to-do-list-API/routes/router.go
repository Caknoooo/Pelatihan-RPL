package routes

import (
	"net/http"
	"github.com/Caknoooo/to-do-list-api/controller"
)

func Router() *http.ServeMux{
	router := http.NewServeMux()

	router.HandleFunc("/api/getData/", controller.Get)
	router.HandleFunc("/api/getData", controller.GetAll)
	router.HandleFunc("/api/createData", controller.Add)
	router.HandleFunc("/api/updateData/", controller.Update)
	router.HandleFunc("/api/deleteData/", controller.Delete)

	return router
}
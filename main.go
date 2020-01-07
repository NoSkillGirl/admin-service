package main

import (
	"net/http"

	"github.com/NoSkillGirl/admin-service/routers"
)

func main() {
	// AdminRoutes Initilization
	routers.AdminRoutes()
	http.ListenAndServe(":8084", nil)
}

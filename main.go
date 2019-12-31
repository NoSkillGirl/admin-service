package main

import (
	"net/http"

	"github.com/NoSkillGirl/admin-service/routers"
)

func main() {
	// AuthRoutes Initilization
	routers.AuthRoutes()
	http.ListenAndServe(":8084", nil)
}

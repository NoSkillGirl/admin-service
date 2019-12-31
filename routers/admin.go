package routers

import (
	"net/http"

	"github.com/NoSkillGirl/admin-service/controllers"
)

// AuthRoutes - All User related Routes
func AuthRoutes() {
	http.HandleFunc("/admin/company/new", controllers.NewCompany)
	http.HandleFunc("/admin/buses/new", controllers.NewBus)
}

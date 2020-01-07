package routers

import (
	"net/http"

	"github.com/NoSkillGirl/admin-service/controllers"
)

// AdminRoutes - All User related Routes
func AdminRoutes() {
	http.HandleFunc("/admin/company/new", controllers.NewCompany)
	http.HandleFunc("/admin/buses/new", controllers.NewBus)
}

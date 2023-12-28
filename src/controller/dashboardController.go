package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GoGym/src/domain/service"
	"github.com/GoGym/src/utils"
	lib "github.com/aag2807/triplex-to-go"
)

type DashboardController struct {
	Arguments    lib.Arguments
	State        lib.State
	Renderer     *utils.PartialRenderer
	userService  *service.LoginService
	pageRenderer *utils.TemplateRenderer
}

func NewDashboardController() *DashboardController {
	return &DashboardController{
		Arguments:    lib.Arguments{},
		State:        lib.State{},
		Renderer:     utils.NewPartialRenderer(),
		userService:  service.NewLoginService(),
		pageRenderer: utils.NewTemplateRenderer("src/templates"),
	}
}

func (dc *DashboardController) RenderInitialDashboardHome(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	log.Println(req.Cookies())
	userId := utils.GetUserSession(req).Values
	// user, _ := dc.userService.GetUserByID(userId.(int))
	log.Println(userId)
	err := dc.pageRenderer.LoadAndExecuteTemplateWithDashboardLayout(w, "pages/home.html", nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func (dc *DashboardController) RenderAdminProfile(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	err := dc.pageRenderer.LoadAndExecuteTemplateWithDashboardLayout(w, "pages/admin-profile.html", nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
func (dc *DashboardController) RenderCoachesPage(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	err := dc.pageRenderer.LoadAndExecuteTemplateWithDashboardLayout(w, "pages/coaches.html", nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func (dc *DashboardController) RenderInventoryPage(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	err := dc.pageRenderer.LoadAndExecuteTemplateWithDashboardLayout(w, "pages/inventory.html", nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func (dc *DashboardController) RenderMembersPage(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	err := dc.pageRenderer.LoadAndExecuteTemplateWithDashboardLayout(w, "pages/members.html", nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func (dc *DashboardController) RenderPaymentPage(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	err := dc.pageRenderer.LoadAndExecuteTemplateWithDashboardLayout(w, "pages/payment.html", nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func (dc *DashboardController) RenderPlansPage(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	err := dc.pageRenderer.LoadAndExecuteTemplateWithDashboardLayout(w, "pages/plans.html", nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func (dc *DashboardController) RenderRegistrationPage(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	err := dc.pageRenderer.LoadAndExecuteTemplateWithDashboardLayout(w, "pages/registration.html", nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func (dc *DashboardController) RenderReportsPage(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	err := dc.pageRenderer.LoadAndExecuteTemplateWithDashboardLayout(w, "pages/reports.html", nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

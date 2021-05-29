package routes

import (
	"healtEvaluationApp/cmd/webapp/response"
	healthevaluationapp "healtEvaluationApp/internal/healthEvaluationApp"
	"healtEvaluationApp/internal/healthEvaluationApp/common"
	"healtEvaluationApp/internal/pkg/service"
	"healtEvaluationApp/internal/utility"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/go-chi/cors"
)

func Set(httpService *service.Service) {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "DELETE", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-TOKEN"},
		ExposedHeaders:   []string{"Link", "Content-Disposition"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	routeHandler := newRoutes()
	httpService.Router.Use(cors.Handler)
	initiateRoute(httpService, routeHandler)
}

func initiateRoute(httpService *service.Service, routeHandler *routesHandler) {
	httpService.Router.Get(common.GetUserResource, routeHandler.GetUserLogin)
}

type routesHandler struct {
	HealthEvaluationApp healthevaluationapp.IMySql
}
type IUserRoutes interface {
	GetUserLogin(w http.ResponseWriter, r *http.Request)
}

func newRoutes() *routesHandler {
	return &routesHandler{
		HealthEvaluationApp: healthevaluationapp.NewMySqlRepo(),
	}
}

// GetUserLogin
func (svc *routesHandler) GetUserLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("GET USER Invoked")
	utility.Debug("GET", "INVOKED")
	response.Success(w, "hey")
}

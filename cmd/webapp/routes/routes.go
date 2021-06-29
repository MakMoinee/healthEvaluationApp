package routes

import (
	"encoding/json"
	"fmt"
	"healtEvaluationApp/cmd/webapp/response"
	healthevaluationapp "healtEvaluationApp/internal/healthEvaluationApp"
	"healtEvaluationApp/internal/healthEvaluationApp/common"
	"healtEvaluationApp/internal/healthEvaluationApp/models"
	"healtEvaluationApp/internal/pkg/service"
	"healtEvaluationApp/internal/utility"
	"io/ioutil"
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

// initiateRote initialize the routes available in the application
func initiateRoute(httpService *service.Service, routeHandler *routesHandler) {
	httpService.Router.Post(common.GetUserResource, routeHandler.GetUserLogin)
	httpService.Router.Post(common.SaveUserResource, routeHandler.SaveUser)
	httpService.Router.Put(common.UpdateUserResource, routeHandler.UpdateUser)
}

type routesHandler struct {
	HealthEvaluationApp healthevaluationapp.IMySql
}
type IUserRoutes interface {
	GetUserLogin(w http.ResponseWriter, r *http.Request)
	SaveUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
}

func newRoutes() *routesHandler {
	return &routesHandler{
		HealthEvaluationApp: healthevaluationapp.NewMySqlRepo(),
	}
}

// GetUserLogin route handler
func (svc *routesHandler) GetUserLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("GET USER Invoked")
	utility.Debug("GET", "INVOKED")
	users := models.Users{}
	resp := IResponse{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("Error reading body: %v", err)
		utility.LogPrometheusErr("Error reading body: %v", err.Error())
		errBody := response.BuildErrorMessage(http.StatusBadRequest, err.Error())
		response.Error(w, errBody, http.StatusBadRequest)
		return
	}
	fmt.Println(fmt.Sprintf("%v", string(body)))
	unmarshalErr := json.Unmarshal(body, &users)
	if unmarshalErr != nil {
		log.Errorf("Error unmarshalling body: %v", unmarshalErr)
		utility.LogPrometheusErr("Unmarshal error", unmarshalErr.Error())
		errBody := response.BuildErrorMessage(http.StatusBadRequest, err.Error())
		response.Error(w, errBody, http.StatusBadRequest)
		return
	}
	validationErr := utility.ValidateUser(users, 3)
	if validationErr != nil {
		errBody := response.BuildErrorMessage(http.StatusBadRequest, validationErr.Error())
		response.Error(w, errBody, http.StatusBadRequest)
		return
	}

	finalUser := svc.HealthEvaluationApp.GetUserLogin(users.Username, users.Password)
	if finalUser.UserID > 0 {

		response.Success(w, finalUser)
		return
	}
	resp.IsExist = false
	response.Success(w, resp)

}

// SaveUser route handler
func (svc *routesHandler) SaveUser(w http.ResponseWriter, r *http.Request) {
	log.Println("POST USER Invoked")
	utility.Debug("POST", "INVOKED")
	users := models.Users{}
	resp := IResponse{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		utility.LogPrometheusErr("Error reading body: %v", err.Error())
		errBody := response.BuildErrorMessage(http.StatusBadRequest, err.Error())
		response.Error(w, errBody, http.StatusBadRequest)
		return
	}

	unmarshalErr := json.Unmarshal(body, &users)
	if unmarshalErr != nil {
		log.Errorf("Error unmarshal body: %v", unmarshalErr)
		utility.LogPrometheusErr("Unmarshal error", unmarshalErr.Error())
		errBody := response.BuildErrorMessage(http.StatusBadRequest, unmarshalErr.Error())
		response.Error(w, errBody, http.StatusBadRequest)
		return
	}

	validationErr := utility.ValidateUser(users, 3)
	if validationErr != nil {
		errBody := response.BuildErrorMessage(http.StatusBadRequest, validationErr.Error())
		response.Error(w, errBody, http.StatusBadRequest)
		return
	}

	isSuccess, sErr := svc.HealthEvaluationApp.SaveUser(users)
	if sErr != nil {
		errBody := response.BuildErrorMessage(http.StatusBadRequest, sErr.Error())
		response.Error(w, errBody, getProperStatCodes(sErr.Error(), http.StatusBadRequest))
		return
	}
	if isSuccess {
		resp.InsertSuccessful = true
		response.Success(w, resp)
		return
	}

	resp.InsertSuccessful = false
	response.Success(w, resp)

}

func (svc *routesHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("PUT USER Invoked")
	utility.Debug("PUT", "INVOKED")
	users := models.Users{}
	resp := IResponse{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		utility.LogPrometheusErr("Error reading body: %v", err.Error())
		errBody := response.BuildErrorMessage(http.StatusBadRequest, err.Error())
		response.Error(w, errBody, http.StatusBadRequest)
		return
	}

	unmarshalErr := json.Unmarshal(body, &users)
	if unmarshalErr != nil {
		log.Errorf("Error unmarshal body: %v", unmarshalErr)
		utility.LogPrometheusErr("Unmarshal error", unmarshalErr.Error())
		errBody := response.BuildErrorMessage(http.StatusBadRequest, unmarshalErr.Error())
		response.Error(w, errBody, http.StatusBadRequest)
		return
	}
	validationErr := utility.ValidateUser(users, 1)
	if validationErr != nil {
		errBody := response.BuildErrorMessage(http.StatusBadRequest, validationErr.Error())
		response.Error(w, errBody, http.StatusBadRequest)
		return
	}
	if svc.HealthEvaluationApp.UpdateUser(users) {
		resp.UpdateSuccessful = true
		response.Success(w, resp)
		return
	}

	resp.UpdateSuccessful = false
	response.Success(w, resp)
}

type IResponse struct {
	UpdateSuccessful bool `json:"isUpdateSuccessful"`
	InsertSuccessful bool `json:"isInsertSuccessful"`
	IsExist          bool `json:"isExist"`
}

func getProperStatCodes(msg string, statCode int) int {
	if msg == common.TokenNotPresent {
		statCode = http.StatusForbidden
	}

	return statCode
}

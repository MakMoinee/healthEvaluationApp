package healthevaluationapp

import (
	"healtEvaluationApp/internal/healthEvaluationApp/models"
	"healtEvaluationApp/internal/healthEvaluationApp/repository/mysql"
)

type service struct {
	MySql mysql.IMySql
}

type IMySql interface {
	GetUserLogin(un string, pw string) models.Users
	SaveUser(user models.Users) bool
}

func NewMySqlRepo() IMySql {
	svc := service{}
	svc.MySql = mysql.NewService()
	return &svc
}

// GetUserLogin retrieves the login credentials
func (svc *service) GetUserLogin(un string, pw string) models.Users {
	return svc.MySql.GetUserLogin(un, pw)
}

func (svc *service) SaveUser(user models.Users) bool {
	return svc.MySql.SaveUser(user)
}

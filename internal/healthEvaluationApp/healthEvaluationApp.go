package healthevaluationapp

import (
	"healtEvaluationApp/internal/healthEvaluationApp/models"
	"healtEvaluationApp/internal/healthEvaluationApp/repository/mysql"

	"github.com/prometheus/common/log"
)

type service struct {
	MySql mysql.IMySql
}

type IMySql interface {
	GetUserLogin(un string, pw string) models.Users
	SaveUser(user models.Users) (bool, error)
	UpdateUser(user models.Users) bool
	GetAssessments(category int) []models.Assessment
	CreateAssessmentDetails(id int, userId int, ans string) bool
	GetHabits(id int) []models.Habit
	DeleteHabit(id int) bool
}

func NewMySqlRepo() IMySql {
	svc := service{}
	svc.MySql = mysql.NewService()
	return &svc
}

// GetUserLogin retrieves the login credentials
func (svc *service) GetUserLogin(un string, pw string) models.Users {
	mUsers := svc.MySql.GetUserLogin(un, pw)
	if mUsers.UserID > 0 {
		tokenStr, err := svc.MySql.GenerateNewToken(mUsers.UserID)
		if err != nil {
			log.Error("Error generating token")
			return models.Users{}
		}
		mUsers.UserToken = tokenStr.TokenString
	}

	return mUsers
}

func (svc *service) SaveUser(user models.Users) (bool, error) {
	isSaved, err := svc.MySql.SaveUser(user)
	if err != nil {
		return false, err
	}
	return isSaved, nil
}

func (svc *service) UpdateUser(user models.Users) bool {
	return svc.MySql.UpdateUser(user)
}

func (svc *service) GetAssessments(category int) []models.Assessment {
	return svc.MySql.GetAssessments(category)
}

func (svc *service) CreateAssessmentDetails(id int, userId int, ans string) bool {
	return svc.MySql.CreateAssessmentDetails(id, userId, ans)
}

func (svc *service) GetHabits(id int) []models.Habit {
	return svc.MySql.GetHabits(id)

}

func (svc *service) DeleteHabit(id int) bool {
	return svc.MySql.DeleteHabit(id)
}

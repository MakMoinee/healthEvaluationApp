package mysql

import (
	"database/sql"
	"fmt"
	"healtEvaluationApp/cmd/webapp/config"
	"healtEvaluationApp/internal/healthEvaluationApp/common"
	"healtEvaluationApp/internal/healthEvaluationApp/models"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type service struct {
	DatabaseName     string
	QueryString      string
	DbDriver         string
	ConnectionString string
	Db               *sql.DB
}

type IMySql interface {
	// GetUserLogin - Retrieves User Details by Username and Password.
	GetUserLogin(un string, pw string) models.Users
	// GetUserById - Retrievs User Details by the specified user Id
	GetUserById(userId int) models.Users
	// SaveUser - Inserts new user in db.
	SaveUser(user models.Users) bool
	// UpdateUser - Updates the user in db
	UpdateUser(user models.Users) bool
}

func NewService() IMySql {
	svc := service{}
	svc.Set()
	return &svc
}

// Set() - setting all the needed configurations in mysql_repository
func (svc *service) Set() {
	svc.DatabaseName = config.Registry.GetString("DB_NAME")
	svc.DbDriver = config.Registry.GetString("DB_DRIVER")
	svc.ConnectionString = config.Registry.GetString("MYSQL_USERNAME") + ":" + config.Registry.GetString("MYSQL_PASSWORD") + "@" + config.Registry.GetString("CONNECTION_STRING") + svc.DatabaseName
	svc.Db = svc.openDBConnection()
	defer svc.Db.Close()
	log.Info("DATABASE_CONNECTED")
}

func (svc *service) GetUserLogin(un string, pw string) models.Users {
	user := models.Users{}
	svc.Db = svc.openDBConnection()
	str := common.GetUserLoginQueryFirstClause + "`username`='" + un + "' and `password`='" + pw + "'" + common.GetUserLoginQueryEndClause
	// perform a db.Query insert
	result, err := svc.Db.Query(str)
	defer svc.Db.Close()
	// if there is an error inserting, handle it
	if err != nil {
		log.Error(err.Error())
	}
	// be careful deferring Queries if you are using transactions

	for result.Next() {
		err = result.Scan(&user.UserID, &user.Username, &user.Password, &user.UserType, &user.LastModifiedDate)
		if err != nil {
			log.Errorf(err.Error())
		}
		// and then print out the tag's Name attribute
		log.Printf("MYSQL Query was successful")
	}

	defer result.Close()
	return user
}

func (svc *service) GetUserById(userId int) models.Users {
	users := models.Users{}
	queryString := fmt.Sprintf(common.GetUserByIdQueryClause, strconv.Itoa(userId))

	svc.Db = svc.openDBConnection()

	result, err := svc.Db.Query(queryString)
	if err != nil {
		log.Error(err.Error())
		return users
	}

	defer svc.Db.Close()

	for result.Next() {
		err := result.Scan(&users.UserID, &users.Username, &users.Password, &users.UserType, &users.LastModifiedDate)
		if err != nil {
			log.Error(err.Error())
			return users
		}

	}
	defer result.Close()

	return users

}

func (svc *service) SaveUser(user models.Users) bool {
	isExist := true
	// checks the user if its already in db
	newUser := svc.GetUserLogin(user.Username, user.Password)
	if newUser.UserID > 0 {
		log.Warn("Username and Password are already in DB. Try again!")
		return false
	}
	svc.Db = svc.openDBConnection()
	queryString := common.InsertUserQueryFirstClause + "'" + user.Username + "','" + user.Password + "'," + strconv.Itoa(user.UserType) + ",now()" + common.InsertUserQueryEndClause

	insertResult, err := svc.Db.Query(queryString)
	defer svc.Db.Close()
	if err != nil {
		log.Errorf(err.Error())
		return false
	}

	defer insertResult.Close()
	log.Printf("Successfully inserted record!")
	return isExist
}

func (svc *service) UpdateUser(user models.Users) bool {
	isUpdate := false
	// checks the user if exist
	newUser := svc.GetUserById(user.UserID)
	if newUser.UserID > 0 {
		svc.Db = svc.openDBConnection()
		queryString := fmt.Sprintf(common.UpdateUserQueryClause, newUser.Username, newUser.Password, strconv.Itoa(newUser.UserID))
		updateResult, err := svc.Db.Query(queryString)
		defer svc.Db.Close()
		if err != nil {
			log.Errorf(err.Error())
			return false
		}

		defer updateResult.Close()
		log.Printf("Successfully updated record!")
		isUpdate = true
	} else {
		log.Error("Failed to update, User is not in the db. ")
	}
	return isUpdate
}

func (svc *service) openDBConnection() *sql.DB {
	db, err := sql.Open(svc.DbDriver, svc.ConnectionString)
	if err != nil {
		log.Error(err.Error())
	}
	return db
}

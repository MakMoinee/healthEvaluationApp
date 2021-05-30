package mysql

import (
	"database/sql"
	"fmt"
	"healtEvaluationApp/cmd/webapp/config"
	"healtEvaluationApp/internal/healthEvaluationApp/common"
	"healtEvaluationApp/internal/healthEvaluationApp/models"

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
	GetUserLogin(un string, pw string) models.Users
}

func NewService() IMySql {
	svc := service{}
	svc.Set()
	return &svc
}

func (svc *service) Set() {
	svc.DatabaseName = config.Registry.GetString("DB_NAME")
	svc.DbDriver = config.Registry.GetString("DB_DRIVER")
	svc.ConnectionString = config.Registry.GetString("MYSQL_USERNAME") + ":" + config.Registry.GetString("MYSQL_PASSWORD") + "@" + config.Registry.GetString("CONNECTION_STRING") + svc.DatabaseName
	svc.Db = svc.openDBConnection()
	defer svc.Db.Close()
	log.Info("DATABASE_CONNECTED")
}

// GetUserLogin Retrieves User Details by Username and Password
func (svc *service) GetUserLogin(un string, pw string) models.Users {
	user := models.Users{}
	svc.Db = svc.openDBConnection()
	str := common.GetUserLoginQueryFirstClause + "`username`='" + un + "' and `password`='" + pw + "'" + common.GetUserLoginQueryEndClause
	fmt.Println(str)
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
		log.Printf(fmt.Sprintf("%v", user.UserID))
		log.Printf(user.Username)
		log.Printf(user.Password)
		log.Printf(fmt.Sprintf("%v", user.UserType))
		log.Printf(user.LastModifiedDate)
	}

	defer result.Close()
	return user
}

func (svc *service) openDBConnection() *sql.DB {
	db, err := sql.Open(svc.DbDriver, svc.ConnectionString)
	if err != nil {
		log.Error(err.Error())
	}
	return db
}

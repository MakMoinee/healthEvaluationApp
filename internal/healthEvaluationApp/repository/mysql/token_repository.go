package mysql

import (
	"fmt"
	"healtEvaluationApp/internal/healthEvaluationApp/common"
	"healtEvaluationApp/internal/healthEvaluationApp/models"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func (svc *service) GetTokenById(userToken string) models.Tokens {
	tokens := models.Tokens{}

	svc.Db = svc.openDBConnection()
	queryString := fmt.Sprintf(common.GetTokenQueryClause, userToken)

	result, err := svc.Db.Query(queryString)
	if err != nil {
		log.Error(err.Error())
	}
	defer svc.Db.Close()
	for result.Next() {
		scanErr := result.Scan(&tokens.TokenID, &tokens.TokenString, &tokens.UserID, &tokens.DateCreated)
		if scanErr != nil {
			log.Error(scanErr.Error())
			return tokens
		}
	}
	defer result.Close()
	return tokens
}

func (svc *service) GenerateNewToken(userId int) (models.Tokens, error) {
	tokens := models.Tokens{}
	randomToken := generateRandomString(strconv.Itoa(userId))
	svc.Db = svc.openDBConnection()
	queryString := fmt.Sprintf(common.GenerateTokenQueryClause, randomToken, strconv.Itoa(userId))

	isInsert, err := svc.Db.Query(queryString)
	defer svc.Db.Close()
	if err != nil {
		log.Error(err.Error())
		return tokens, err
	}
	defer isInsert.Close()

	tokens = svc.GetTokenById(randomToken)
	return tokens, nil
}

func generateRandomString(str string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}

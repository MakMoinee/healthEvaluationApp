package utility

import (
	"fmt"
	"healtEvaluationApp/internal/healthEvaluationApp/models"
)

func ValidateUser(users models.Users, flag int) error {
	var err error
	switch flag {
	case 1:
		if users.UserID > 0 && len(users.Username) > 0 && len(users.Password) > 0 {
			err = nil
		} else {
			err = fmt.Errorf("required paramaters are not supplied")
		}
		break
	case 2:
		if users.UserID > 0 {
			err = nil
		} else {
			err = fmt.Errorf("required user id is not supplied")
		}
		break
	case 3:
		if len(users.Username) > 0 && len(users.Password) > 0 {
			err = nil
		} else {
			err = fmt.Errorf("required username and password are not supplied")
		}
		break
	}
	return err
}

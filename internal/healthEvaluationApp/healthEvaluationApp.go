package healthevaluationapp

type service struct {
	Username string
	Password string
	DbName   string
}

type IMySql interface {
	GetUserLogin(un string, pw string)
}

func NewMySqlRepo() IMySql {
	svc := service{}

	return &svc
}

func (svc *service) set() {
	svc.DbName = ""
}

// GetUserLogin retrieves the login credentials
func (svc *service) GetUserLogin(un string, pw string) {

}

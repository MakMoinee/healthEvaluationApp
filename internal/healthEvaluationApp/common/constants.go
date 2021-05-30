package common

const (
	GetUserResource              = "/get/user"
	SaveUserResource             = "/save/user"
	ContentTypeKey               = "Content-Type"
	ContentTypeValue             = "application/json; charset=UTF-8"
	GetUserLoginQueryFirstClause = "Select * from `users` where "
	GetUserLoginQueryEndClause   = " limit 1"

	InsertUserQueryFirstClause = "INSERT INTO `users` (Username,Password,UserType,LastModifiedDate) VALUES ("
	InsertUserQueryEndClause   = ")"
)

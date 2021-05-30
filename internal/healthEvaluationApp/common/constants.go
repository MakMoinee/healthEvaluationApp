package common

const (
	GetUserResource    = "/get/user"
	SaveUserResource   = "/save/user"
	UpdateUserResource = "/update/user"

	GetTokenResource      = "/token/user"
	GenerateTokenResource = "/token/generate"

	ContentTypeKey   = "Content-Type"
	ContentTypeValue = "application/json; charset=UTF-8"
	TimeFormat       = "2006-01-02 15:04:05"

	GetUserByIdQueryClause = "SELECT * from `users` where `UserID`=%v"

	GetUserLoginQueryFirstClause = "Select * from `users` where "
	GetUserLoginQueryEndClause   = " limit 1"

	InsertUserQueryFirstClause = "INSERT INTO `users` (Username,Password,UserType,LastModifiedDate) VALUES ("
	InsertUserQueryEndClause   = ")"

	UpdateUserQueryClause = "UPDATE `users` SET `Username`='%v',`Password`='%v',LastModifiedDate=NOW() WHERE UserID=%v"

	// tokens
	GetTokenQueryClause = "SELECT * FROM `tokens` where `TokenString`='%v' limit 1"
	TokenNotPresent     = "there is no valid token passed. login again"

	GenerateTokenQueryClause = "INSERT INTO `tokens` (`TokenString`,`UserID`,`DateCreated`) VALUES ('%v',%v,now())"
)

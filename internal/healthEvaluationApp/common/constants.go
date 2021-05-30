package common

const (
	GetUserResource    = "/get/user"
	SaveUserResource   = "/save/user"
	UpdateUserResource = "/update/user"
	ContentTypeKey     = "Content-Type"
	ContentTypeValue   = "application/json; charset=UTF-8"

	GetUserByIdQueryClause = "SELECT * from `users` where `UserID`=%v"

	GetUserLoginQueryFirstClause = "Select * from `users` where "
	GetUserLoginQueryEndClause   = " limit 1"

	InsertUserQueryFirstClause = "INSERT INTO `users` (Username,Password,UserType,LastModifiedDate) VALUES ("
	InsertUserQueryEndClause   = ")"

	UpdateUserQueryClause = "UPDATE `users` SET `Username`='%v',`Password`='%v',LastModifiedDate=NOW() WHERE UserID=%v"
)

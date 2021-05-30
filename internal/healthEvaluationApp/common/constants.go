package common

const (
	GetUserResource              = "/get/user"
	ContentTypeKey               = "Content-Type"
	ContentTypeValue             = "application/json; charset=UTF-8"
	GetUserLoginQueryFirstClause = "Select * from `users` where "
	GetUserLoginQueryEndClause   = " limit 1"
)

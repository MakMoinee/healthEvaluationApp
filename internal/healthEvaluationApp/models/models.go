package models

// Students data model
type Students struct {
	UserID           int    `json:"userId,omitempty"`
	StudentID        string `json:"studentId,omitempty"`
	FirstName        string `json:"firstName,omitempty"`
	MiddleName       string `json:"middleName,omitempty"`
	LastName         string `json:"lastName,omitempty"`
	Address          string `json:"address,omitempty"`
	BirthDate        string `json:"birthDate,omitempty"`
	LastModifiedDate string `json:"lastModifiedDate,omitempty"`
}

// Users data model
type Users struct {
	UserID           int    `json:"userId,omitempty"`
	Username         string `json:"username,omitempty"`
	Password         string `json:"password,omitempty"`
	UserType         int    `json:"userType,omitempty"`
	LastModifiedDate string `json:"lastModifiedDate,omitempty"`
	UserToken        string `json:"userToken,omitempty"`
}

// Teachers data model
type Teachers struct {
	UserID           int    `json:"userId,omitempty"`
	TeacherID        int    `json:"teacherId,omitempty"`
	ProfessionID     int    `json:"professionId,omitempty"`
	FirstName        string `json:"firstName,omitempty"`
	MiddleName       string `json:"middleName,omitempty"`
	LastName         string `json:"lastName,omitempty"`
	Address          string `json:"address,omitempty"`
	BirthDate        string `json:"birthDate,omitempty"`
	LastModifiedDate string `json:"lastModifiedDate,omitempty"`
}

// Profession data model
type Profession struct {
	ProfessionID     int    `json:"professionId,omitempty"`
	ProfessionName   string `json:"professionName,omitempty"`
	LastModifiedDate string `json:"lastModifiedDate,omitempty"`
}

// Token data model
type Tokens struct {
	TokenID     int    `json:"tokenId,omitempty"`
	TokenString string `json:"tokenString,omitempty"`
	UserID      int    `json:"userId,omitempty"`
	DateCreated string `json:"dateCreation,omitempty"`
}

type Habit struct {
	HabitID     int    `json:"HabitID"`
	UserID      int    `json:"UserID"`
	HabitName   string `json:"HabitName"`
	HabitTime   string `json:"HabitTime"`
	HabitDate   string `json:"HabitDate"`
	HabitStatus int    `json:"HabitStatus"`
}

type Todo struct {
	TodoID    int    `json:"TodoID"`
	Note      string `json:"Note"`
	StartDate string `json:"StartDate"`
	EndDate   string `json:"EndDate"`
}

type Assessment struct {
	AssessmentID  int    `json:"AssessmentID"`
	Sequence      int    `json:"Sequence"`
	Question      string `json:"Question"`
	AnswerA       string `json:"AnswerA"`
	AnswerB       string `json:"AnswerB"`
	AnswerC       string `json:"AnswerC`
	AnswerD       string `json:"AnswerD"`
	AnswerAPoints string `json:"AnswerAPoints"`
	AnswerBPoints string `json:"AnswerBPoints"`
	AnswerCPoints string `json:"AnswerCPoints`
	AnswerDPoints string `json:"AnswerDPoints"`
	Status        int    `json:"Status"`
	Category      int    `json:"Category"`
}

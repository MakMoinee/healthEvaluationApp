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

package Models

type Student struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	DOB       string `json:"dob"`
	Address   string `json:"address"`
}

func (b *Student) TableName() string {
	return "student"
}

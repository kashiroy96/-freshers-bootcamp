package Models

type StudentAndMarks struct {
	Id        uint    `json:"id"`
	FirstName string  `json:"first-name"`
	LastName  string  `json:"last-name"`
	DOB       string  `json:"dob"`
	Address   string  `json:"address"`
	Marks     []Marks `json:"marks"`
}

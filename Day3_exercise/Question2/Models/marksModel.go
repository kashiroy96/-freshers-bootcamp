package Models

type Marks struct {
	Id        uint   `json:"id"`
	StudentId uint   `json:"id" gorm:"foreignKey"`
	Subject   string `json:"subject"`
	Mark      uint   `json:"marks"`
}

func (b *Marks) TableName() string {
	return "marks"
}

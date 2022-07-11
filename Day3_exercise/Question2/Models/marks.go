package Models

import (
	"Question2/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// GetAllStudents Fetch all student data
func GetAllMarks(marks *[]Marks) (err error) {
	if err = Config.DB.Find(marks).Error; err != nil {
		return err
	}
	return nil
}

// CreateMarks ... Insert New data
func CreateMarks(marks *[]Marks) (err error) {
	if err = Config.DB.Create(marks).Error; err != nil {
		return err
	}
	return nil
}

// GetMarksByID ... Fetch only one marks by
func GetMarksByID(marks *[]Marks, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(marks).Error; err != nil {
		return err
	}
	return nil
}

// UpdateMarks ... Update user
func UpdateMarks(marks *[]Marks, id string) (err error) {
	fmt.Println(marks)
	Config.DB.Save(marks)
	return nil
}

// DeleteMarks ... Delete user
func DeleteMarks(marks *[]Marks, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(marks)
	return nil
}

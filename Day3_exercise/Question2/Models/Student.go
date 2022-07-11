package Models

import (
	"Question2/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// GetAllStudents Fetch all student data
func GetAllStudents(student *[]Student) (err error) {
	if err = Config.DB.Find(student).Error; err != nil {
		return err
	}
	return nil
}

// CreateStudent ... Insert New data
func CreateStudent(student *[]Student) (err error) {
	if err = Config.DB.Create(student).Error; err != nil {
		return err
	}
	return nil
}

// GetStudentByID ... Fetch only one user by Id
func GetStudentByID(student *[]Student, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(student).Error; err != nil {
		return err
	}
	return nil
}

// UpdateStudent ... Update user
func UpdateStudent(student *[]Student, id string) (err error) {
	fmt.Println(student)
	Config.DB.Save(student)
	return nil
}

// DeleteUser ... Delete user
func DeleteUser(student *[]Student, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(student)
	return nil
}

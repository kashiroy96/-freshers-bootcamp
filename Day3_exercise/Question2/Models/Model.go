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
func CreateStudent(student *Student) (err error) {
	if err = Config.DB.Create(student).Error; err != nil {
		return err
	}
	return nil
}

// GetStudentByID ... Fetch only one user by Id
func GetStudentByID(student *Student, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(student).Error; err != nil {
		return err
	}
	return nil
}

// UpdateStudent ... Update user
func UpdateStudent(student *Student, id string) (err error) {
	fmt.Println(student)
	Config.DB.Save(student)
	return nil
}

// DeleteStudent ... Delete student
func DeleteStudent(student *Student, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(student)
	return nil
}

// GetAllMarks Fetch all student data
func GetAllMarks(marks *[]Marks) (err error) {
	if err = Config.DB.Find(marks).Error; err != nil {
		return err
	}
	return nil
}

// CreateMarks ... Insert New data
func CreateMarks(marks *Marks) (err error) {
	if err = Config.DB.Create(marks).Error; err != nil {
		return err
	}
	return nil
}

// GetMarksByID ... Fetch only one user by Id
func GetMarksByID(marks *Marks, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(marks).Error; err != nil {
		return err
	}
	return nil
}

// GetMarksByStudentID to getting marks of particular student
func GetMarksByStudentID(marks *[]Marks, id string) (err error) {
	if err = Config.DB.Where("student_id = ?", id).Find(marks).Error; err != nil {
		return err
	}
	return nil
}

// UpdateMarks ... Update user
func UpdateMarks(marks *Marks, id string) (err error) {
	fmt.Println(marks)
	Config.DB.Save(marks)
	return nil
}

// DeleteMarks ... Delete user
func DeleteMarks(marks *Marks, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(marks)
	return nil
}

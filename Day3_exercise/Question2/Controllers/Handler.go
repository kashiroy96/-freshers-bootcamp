package Controllers

import (
	Models "Question2/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetStudent ... Get all Students
func GetStudent(c *gin.Context) {
	var student []Models.Student
	err := Models.GetAllStudents(&student)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

// CreateStudent ... Create Student
func CreateStudent(c *gin.Context) {
	var student Models.Student
	c.BindJSON(&student)
	err := Models.CreateStudent(&student)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

// GetStudentByID ... Get the Student by id
func GetStudentByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var student Models.Student
	err := Models.GetStudentByID(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

// UpdateStudent ... Update the Student information
func UpdateStudent(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")
	err := Models.GetStudentByID(&student, id)
	if err != nil {
		c.JSON(http.StatusNotFound, student)
	}
	c.BindJSON(&student)
	err = Models.UpdateStudent(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

// DeleteStudent ... Delete the Student
func DeleteStudent(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")
	err := Models.DeleteStudent(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

// GetMarks ... Get all Marks
func GetMarks(c *gin.Context) {
	var marks []Models.Marks
	err := Models.GetAllMarks(&marks)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, marks)
	}
}

// CreateMarks ... Create Marks
func CreateMarks(c *gin.Context) {
	var marks Models.Marks
	c.BindJSON(&marks)
	err := Models.CreateMarks(&marks)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, marks)
	}
}

// GetMarksByID ... Get the Marks by id
func GetMarksByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var marks Models.Marks
	err := Models.GetMarksByID(&marks, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, marks)
	}
}

// UpdateMarks ... Update the Marks information
func UpdateMarks(c *gin.Context) {
	var marks Models.Marks
	id := c.Params.ByName("id")
	err := Models.GetMarksByID(&marks, id)
	if err != nil {
		c.JSON(http.StatusNotFound, marks)
	}
	c.BindJSON(&marks)
	err = Models.UpdateMarks(&marks, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, marks)
	}
}

// DeleteMarks ... Delete the Marks
func DeleteMarks(c *gin.Context) {
	var marks Models.Marks
	id := c.Params.ByName("id")
	err := Models.DeleteMarks(&marks, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

func StudentsMarks(c *gin.Context) {
	studentId := c.Params.ByName("id")
	var student Models.Student
	err := Models.GetStudentByID(&student, studentId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var marks []Models.Marks
		err1 := Models.GetMarksByStudentID(&marks, studentId)
		if err1 != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {

			studentAndMarks := Models.StudentAndMarks{Id: student.Id, FirstName: student.FirstName, LastName: student.LastName, DOB: student.DOB, Address: student.Address, Marks: marks}
			c.JSON(http.StatusOK, studentAndMarks)
		}
	}

}

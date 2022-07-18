package Routes

import (
	"Question2/Controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/student-api")
	{
		grp1.GET("students", Controllers.GetStudent)
		grp1.POST("student", Controllers.CreateStudent)
		grp1.GET("student/:id", Controllers.GetStudentByID)
		grp1.PUT("student/:id", Controllers.UpdateStudent)
		grp1.DELETE("student/:id", Controllers.DeleteStudent)

		grp1.GET("student_marks/:id", Controllers.StudentsMarks)

		grp1.GET("marks", Controllers.GetMarks)
		grp1.POST("marks", Controllers.CreateMarks)
		grp1.GET("marks/:id", Controllers.GetMarksByID)
		grp1.PUT("marks/:id", Controllers.UpdateMarks)
		grp1.DELETE("marks/:id", Controllers.DeleteMarks)
	}
	return r
}

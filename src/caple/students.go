package caple

import (
	"fmt"

	"github.com/kataras/iris"
)

// Student represents a student
type Student struct {
	ID    int64
	Name  string
	Email string
}

func (s Student) String() string {
	return fmt.Sprintf("Student<%d %s %s>", s.ID, s.Name, s.Email)
}

func studentsHandler(c *iris.Context) {
	var students []Student
	err := db.Model(&students).Select()
	if err != nil {
		c.Log(err.Error())
		c.JSON(iris.StatusServiceUnavailable, iris.Map{
			"error": err.Error(),
		})
	} else {
		c.JSON(iris.StatusOK, students)
	}
}

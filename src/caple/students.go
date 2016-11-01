package caple

import (
	"fmt"

	pg "gopkg.in/pg.v5"

	"github.com/kataras/iris"
)

// Student represents a student
type Student struct {
	ID    uint64
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

func studentByIDHandler(c *iris.Context) {
	var student Student
	studentID, err := c.ParamInt("id")
	if err != nil {
		c.JSON(iris.StatusServiceUnavailable, iris.Map{
			"error": err.Error(),
		})
	}
	c.Log("Query Student ID: %d", studentID)
	err = db.Model(&student).Where("id = ?", studentID).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			c.JSON(iris.StatusNotFound, iris.Map{})
		} else {
			c.Log(err.Error())
			c.JSON(iris.StatusServiceUnavailable, iris.Map{
				"error": err.Error(),
			})
		}
	} else {
		c.JSON(iris.StatusOK, student)
	}
}

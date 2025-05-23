package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"

)

type Student struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Course string `json:"course"`
	Fees   string `json:"fees"`
}

var students = []Student{
	{ID: 1, Name: "John", Gender: "Male", Course: "Java", Fees: "$300.00"},
	{ID: 2, Name: "Dara", Gender: "Male", Course: "C++", Fees: "$300.00"},
	{ID: 3, Name: "Bora", Gender: "Male", Course: "C#", Fees: "$300.00"},
}

// Get all students
func getStudents(c *gin.Context) {
	c.JSON(http.StatusOK, students)
}

// Add a new student
func addStudent(c *gin.Context) {
	var newStudent Student
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newStudent.ID = len(students) + 1
	students = append(students, newStudent)
	c.JSON(http.StatusCreated, newStudent)
}

// Update a student
func updateStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedStudent Student
	if err := c.ShouldBindJSON(&updatedStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, student := range students {
		if student.ID == id {
			students[i] = updatedStudent
			students[i].ID = id // Ensure the ID remains unchanged
			c.JSON(http.StatusOK, updatedStudent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}

// Delete a student
func deleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, student := range students {
		if student.ID == id {
			students = append(students[:i], students[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}

func main() {
	r := gin.Default()

	// Serve static files
	r.Static("/static", "./bat")

	// API routes
	r.GET("/students", getStudents)
	r.POST("/students", addStudent)
	r.PUT("/students/:id", updateStudent)
	r.DELETE("/students/:id", deleteStudent)

	// Serve the frontend
	r.StaticFile("/", "./bat/form.html")

	r.Run(":8080") // Run on localhost:8080
}

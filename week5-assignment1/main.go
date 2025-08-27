package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// โครงสร้าง Todo
type Todo struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Details string `json:"details"`
	Done    bool   `json:"done"`
}

// เก็บข้อมูล todo ในหน่วยความจำ
var todos = []Todo{
	{ID: 1, Title: "ทำการบ้าน", Details: "การบ้านวิชา Database", Done: false},
	{ID: 2, Title: "อ่านหนังสือ", Details: "อ่านบทที่ 3 วิชา Network", Done: false},
}

// ฟังก์ชันแสดงรายการทั้งหมด
func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

// ฟังก์ชันเพิ่มรายการใหม่
func addTodo(c *gin.Context) {
	var newTodo Todo

	// ผูก JSON จาก request body เข้ากับ struct newTodo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// กำหนด ID ใหม่ (id ล่าสุด + 1)
	newTodo.ID = len(todos) + 1
	todos = append(todos, newTodo)

	c.IndentedJSON(http.StatusCreated, newTodo)
}

func main() {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.GET("/todos", getTodos) // แสดง todo ทั้งหมด
		api.POST("/todos", addTodo) // เพิ่ม todo ใหม่
	}

	r.Run(":8080") // รันที่พอร์ต 8080
}

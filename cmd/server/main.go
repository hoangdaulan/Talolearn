package main

import (
	db2 "Talolearn-Project/internal/db"
	"Talolearn-Project/internal/middleware"
	"Talolearn-Project/internal/module/category"
	"Talolearn-Project/internal/module/course"
	"Talolearn-Project/internal/module/teacher"

	"github.com/gin-gonic/gin"
)

func main() {
	dsn := "host=localhost user=postgres password=123 dbname=talolearn12 port=5432 sslmode=disable"

	db := db2.InitPostgresDB(dsn)

	err := db.AutoMigrate(&teacher.Teacher{})
	if err != nil {
		panic("failed to migrate teacher database")
	}
	err = db.AutoMigrate(&category.Category{})
	if err != nil {
		panic("failed to migrate category database")
	}
	err = db.AutoMigrate(&course.Course{})
	if err != nil {
		panic("failed to migrate course database")
	}

	r := gin.Default()
	r.Use(middleware.Logger())

	teacher.TeacherRouter(r, teacher.NewTeacherHandler(*teacher.NewGormRepository(db)))
	category.CategoryRouter(r, category.NewCategoryHandler(*category.NewGormRepository(db)))
	course.CourseRouter(r, course.NewCourseHandler(*course.NewGormRepository(db)))

	err = r.Run(":8080")
	if err != nil {
		return
	}
}

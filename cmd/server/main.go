package main

import (
	db2 "Talolearn-Project/internal/db"
	"Talolearn-Project/internal/middleware"
	"Talolearn-Project/internal/module/category"
	"Talolearn-Project/internal/module/course"
	"Talolearn-Project/internal/module/student"
	"Talolearn-Project/internal/module/teacher"
	"Talolearn-Project/internal/module/user"

	"github.com/gin-gonic/gin"
)

func main() {
	dsn := "host=localhost user=postgres password=123 dbname=talolearn12 port=5432 sslmode=disable"
	//dsn := fmt.Sprintf(
	//	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	//	os.Getenv("DB_HOST"),
	//	os.Getenv("DB_USER"),
	//	os.Getenv("DB_PASSWORD"),
	//	os.Getenv("DB_NAME"),
	//	os.Getenv("DB_PORT"),
	//)

	db := db2.InitPostgresDB(dsn)

	err := db.AutoMigrate(&user.User{})
	if err != nil {
		panic("failed to migrate user database")
	}
	err = db.AutoMigrate(&teacher.Teacher{})
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
	err = db.AutoMigrate(&course.CourseTeacher{})
	if err != nil {
		panic("failed to migrate course_teacher database")
	}
	err = db.AutoMigrate(&student.Student{})
	if err != nil {
		panic("failed to migrate student database")
	}

	r := gin.Default()
	r.Use(middleware.Logger())

	userRepo := user.NewGormRepository(db)
	categoryRepo := category.NewGormRepository(db)
	courseRepo := course.NewGormRepository(db)
	teacherRepo := teacher.NewGormRepository(db)
	studentRepo := student.NewGormRepository(db)

	userService := user.NewService(userRepo)
	categoryService := category.NewCategoryService(categoryRepo)
	courseService := course.NewCourseService(courseRepo, teacherRepo)
	teacherService := teacher.NewTeacherService(teacherRepo, userRepo)
	studentService := student.NewStudentService(studentRepo)

	userHandler := user.NewHandler(*userService)
	categoryHandler := category.NewCategoryHandler(*categoryService)
	courseHandler := course.NewCourseHandler(*courseService)
	teacherHandler := teacher.NewTeacherHandler(*teacherService)
	studentHandler := student.NewStudentHandler(*studentService)

	user.UserRouter(r, userHandler)
	course.CourseRouter(r, courseHandler)
	category.CategoryRouter(r, categoryHandler)
	teacher.TeacherRouter(r, teacherHandler)
	student.StudentRouter(r, studentHandler)

	err = r.Run(":8080")
	if err != nil {
		return
	}
}

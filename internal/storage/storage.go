package storage

import "github.com/Aniket03g/students-api/internal/types"

//using interface concept here

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
	UpdateStudent(id int64, name string, email string, age int) error // NEW
	DeleteStudent(id int64) error                                     // NEW METHOD
}

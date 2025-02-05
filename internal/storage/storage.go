package storage

//using interface concept here

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
}

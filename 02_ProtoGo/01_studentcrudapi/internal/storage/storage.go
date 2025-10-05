package storage


type Storage interface{
	CreateStudent(nam,email string,age int)(int64,error)
}
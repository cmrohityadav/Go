package sqlite

import (
	"database/sql"
	"fmt"
	"main/internal/config"
	"main/internal/types"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct{
	Db *sql.DB
}



func New(cfg *config.Config)(*Sqlite,error){
	db,err:=sql.Open("sqlite3",cfg.StoragePath);
	if err!=nil {
		return  nil,err
	}

	_,err=db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT,
		age INTEGER
	)`)

	if err!=nil {
		return nil,err
	}

	return &Sqlite{
		Db: db,
	}, nil
}


func (s *Sqlite) CreateStudent(name,email string,age int)(int64,error){

	stmt,err:=s.Db.Prepare("INSERT INTO students (name,email,age) VALUES (?,?,?)");
    if err!=nil {
		return 0,err;
	}

	defer stmt.Close();

	result,err:=stmt.Exec(name,email,age);
	if err!=nil {
		return 0,err;
	}


	id,err:=result.LastInsertId();
	if err!=nil {
		return 0,err
	}

	return id,nil;
}

func (s *Sqlite)GetStudentById(id int64)(types.Student,error){
	stmt,err:=s.Db.Prepare("SELECT * FROM students WHERE id=? LIMIT 1");
	if err!=nil{
		return types.Student{},err;
	}

	defer stmt.Close();

	var student types.Student;
	err=stmt.QueryRow(id).Scan(&student.Id,&student.Name,&student.Email,&student.Age);
	if err!=nil{
		if err==sql.ErrNoRows{
			return types.Student{},fmt.Errorf("no student found with id %d",id);
		}
		return types.Student{},fmt.Errorf("query error: %w",err);
	}

	return student,nil;

}

func (s *Sqlite) GetStudentList()([]types.Student,error){

	stmt,err:=s.Db.Prepare("SELECT id,name,email,age FROM students");

	if err!=nil {
		return []types.Student{},err;
	}

	defer stmt.Close();

	var studentList []types.Student;

	rows,err:=stmt.Query();

	if err!=nil{
		return []types.Student{},err;
	}
	defer rows.Close();

	for rows.Next(){
		var student types.Student;
		err=rows.Scan(&student.Id,&student.Name,&student.Email,&student.Age);
		if err!=nil {
			return []types.Student{},err
		}
		studentList=append(studentList, student);
	}

	return studentList,nil


}

func (s *Sqlite) DeleteStudentById(id int64) (bool,error){
	stmt,err:=s.Db.Prepare("DELETE FROM students WHERE id=?");

	if err!=nil{
		return false,err;
	}
	defer stmt.Close();

	result,err:=stmt.Exec(id);
	if err!= nil{
		return false,err;
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	
	if rowsAffected == 0 {
		return false, nil
	}

	return true,nil;

}
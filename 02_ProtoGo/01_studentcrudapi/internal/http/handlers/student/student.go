package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"main/internal/storage"
	"main/internal/types"
	"main/internal/utils/response"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)


func New(storage storage.Storage) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		slog.Info("create student");
		var student types.Student;

		err:=json.NewDecoder(r.Body).Decode(&student);
		if errors.Is(err,io.EOF){
			response.WriteJson(w,http.StatusBadRequest,response.GeneralError(err));
			return;
		}

		if err!=nil {
			response.WriteJson(w,http.StatusBadRequest,response.GeneralError(err));
			return;
		}


		// request validation
		if err:=validator.New().Struct(student); err!=nil{
			validateErrs:=err.(validator.ValidationErrors);
			response.WriteJson(w,http.StatusBadRequest,response.ValidateError(validateErrs));

			return;
		}

		lastId,err:=storage.CreateStudent(student.Name,student.Email,student.Age);

		if err!= nil {
			response.WriteJson(w,http.StatusInternalServerError,err);
			return;
		}

		slog.Info("create student succesfully",slog.String("userId",fmt.Sprint(lastId)));
		response.WriteJson(w,200,map[string]int64{"id":lastId});

		// w.Write([]byte("Welcome to student api handler"));

		// response.WriteJson(w,http.StatusCreated,map[string]string{"success":"OK"})
	}
}



func GetById(s storage.Storage) http.HandlerFunc{
	return func (w http.ResponseWriter,r *http.Request){
		id:=r.PathValue("id");
		slog.Info("Getting a student",slog.String("id",id));

		intId,err:=strconv.ParseInt(id,10,64);
		if err!=nil{
			response.WriteJson(w,400,response.GeneralError(err));
			return;
		}
		student,err:=s.GetStudentById(intId);

		if err!=nil{

			response.WriteJson(w,http.StatusInternalServerError,response.GeneralError(err));

			return;
		}

		response.WriteJson(w,200,student);

	}
}
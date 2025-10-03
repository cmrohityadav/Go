package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"main/internal/types"
	"main/internal/utils/response"
	"net/http"
)


func New() http.HandlerFunc{
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



		slog.Info("create student");
		// w.Write([]byte("Welcome to student api handler"));

		response.WriteJson(w,http.StatusCreated,map[string]string{"success":"OK"})
	}
}
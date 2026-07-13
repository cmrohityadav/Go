package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)
func uploadHandler(w http.ResponseWriter,r *http.Request){
	err:=r.ParseMultipartForm(5000000)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return;
	}


	file,header,err:=r.FormFile("userimg")
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	defer file.Close()
	
	fmt.Println("File Name: ",header.Filename)
	fmt.Println("Size of File: ",header.Size)
	fmt.Println("Header: ",header.Header)

	os.MkdirAll("uploads", 0755)
	dst, err :=os.Create("uploads/"+ header.Filename)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	io.Copy(dst,file)


}

func multipleUploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(20 << 20) // 20 MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	os.MkdirAll("uploads", 0755)

	// "userimgs" is the name attribute of the input
	files := r.MultipartForm.File["userimgs"]

	if len(files) == 0 {
		http.Error(w, "No files uploaded", http.StatusBadRequest)
		return
	}

	for _, header := range files {

		fmt.Println("File:", header.Filename)
		fmt.Println("Size:", header.Size)

		file, err := header.Open()
		if err != nil {
			fmt.Println(err)
			continue
		}

		dst, err := os.Create(filepath.Join("uploads", header.Filename))
		if err != nil {
			file.Close()
			fmt.Println(err)
			continue
		}

		_, err = io.Copy(dst, file)

		dst.Close()
		file.Close()

		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Fprintln(w, "All files uploaded successfully")
}
func main() {

	mux:=http.NewServeMux()

	mux.HandleFunc("/upload",uploadHandler)
	mux.HandleFunc("/uploads",multipleUploadHandler)

	http.ListenAndServe(":3000",mux)

}
package main

import (
	"bufio"
	"fmt"
	// "io"
	// "io/ioutil"
	"os"
)

func main() {

	f,err:=os.Open("example.txt");

	if err!=nil{
		fmt.Println(err);
	}
	defer f.Close();

	fileInfo,err:=f.Stat();

	if err != nil {
		fmt.Println(err);

	}

	fmt.Println("file name: ",fileInfo.Name());
	fmt.Println("is directory?: ",fileInfo.IsDir());
	fmt.Println("size in Bytes: ",fileInfo.Size());
	fmt.Println("file permission: ",fileInfo.Mode());
	fmt.Println("file modified at: ",fileInfo.ModTime());

	buf:=make([]byte,3);

	d,err:=f.Read(buf);

	if err!= nil {
		panic(err);
	}

	fmt.Println("buffer data: ",buf,"size: ",d);
	fmt.Println("buffer in string: ",string(buf));



	//-------------for small file reading--------
	data,err:=os.ReadFile("example.txt");
	if err!= nil{
		panic(err);
	}

	fmt.Println("Reading data using ReadFile(): ",string(data));
	


	//-------Read folder---------

	dir,err:=os.Open("../../");
	if err!=nil {
		panic(err);
	}
	defer dir.Close();


	dirInfo,err:=dir.ReadDir(-1);
	if err!= nil{
		panic(err);
	}
	
	for i:=0;i<len(dirInfo);i++ {
		fmt.Println("\n dir info at index ",i,"value at this ",dirInfo[i]);
		fmt.Println("is  folder: ",dirInfo[i].IsDir()," -> name of it: ",dirInfo[i].Name());
	}


	// -------Create File-----------

	pCreateFile,err:=os.Create("output.txt");
	if err!=nil {
		panic(err);
	}
	defer pCreateFile.Close();

	pCreateFile.WriteString("Hello writing string");
	pCreateFile.WriteString("Writing second time\n");

	pCreateFile.Write([]byte("Hello bro"))




	// -------Streaming---------

	pSourceFile,err:=os.Open("./output.txt");
	if err!= nil {
		panic(err);
	}
	defer pSourceFile.Close()


	pDestFile,err:=os.Create("destfile.txt")
	if err!= nil {
		panic(err);
	}
	defer pDestFile.Close();

	reader:=bufio.NewReader(pSourceFile);
	writer:=bufio.NewWriter(pDestFile);

	for {
		b,err:=reader.ReadByte();
		if err!=nil{
			if err.Error()!="EOF"{
				panic(err);
			}
			break;
		}

		er:=writer.WriteByte(b);
		if er!=nil {
			panic(er);
		}
	}

	writer.Flush();
	fmt.Println("Writen to destfile successfully")



	





	/*
		fmt.Println("Welcome to files in golang")

		content:="This needs to go in file - cmrohityadav.in"

		anyFile,error:=os.Create("./fileCreatedByGo.txt")

		if error!=nil{
			panic(error)
		}

		length,err:=io.WriteString(anyFile,content)
		
		if err !=nil{
			panic(err)
		}
		fmt.Println("lenght is : ",length)

		defer anyFile.Close()


		readFile("./fileCreatedByGo.txt")
	*/
}


/*
func readFile(filename string){
	itsDataInByte,err:=ioutil.ReadFile(filename)
	if err !=nil{
		panic(err)
	}

	fmt.Println("Text data inside the file (Raw) is \n",itsDataInByte)
	fmt.Println("Text data inside the file (string) is \n",string(itsDataInByte))
}
*/
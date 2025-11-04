package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ParseFileNameWithDate(url string,time *time.Time) string {
	dd:=fmt.Sprintf("%02d",time.Day());
	mm:=fmt.Sprintf("%02d",time.Month());
	yyyy:=fmt.Sprintf("%04d",time.Year());

	url=strings.ReplaceAll(url,"%dd",dd);
	url=strings.ReplaceAll(url,"%mm",mm);
	url=strings.ReplaceAll(url,"%yyyy",yyyy);

	return url;
}

func ExtractCSVFromZip(zipPath, destDir string) {
			pZipReaderCloser,err:=zip.OpenReader(zipPath);
			if err!=nil{
				log.Println("Error while opening zip file :",err);
				return;
			}
			
			defer pZipReaderCloser.Close();

			for _,fileInZip:=range pZipReaderCloser.File{
				if filepath.Ext(fileInZip.Name)!=".csv"{
					continue;
				}

				pfileContentStream,err:=fileInZip.Open();
				if err!=nil{
					log.Println("error while reading file of zip:",err);
				}


				outPutPath:=filepath.Join(destDir,fileInZip.Name);

				pOsFileOutPut,err:=os.Create(outPutPath);

				if err!=nil{
					log.Println("error while creating output file from zip:",err);
				};
				
				io.Copy(pOsFileOutPut,pfileContentStream);

				pOsFileOutPut.Close();
				pfileContentStream.Close();

				log.Printf("Extracted from zip %s to %s",zipPath,outPutPath);

			}
}



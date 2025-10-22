package utils

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func ParseUrlWithPresentDate(url string) string {
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



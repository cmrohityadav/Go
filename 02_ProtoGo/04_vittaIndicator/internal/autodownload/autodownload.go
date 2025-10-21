package autodownload

import (
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"path/filepath"
	"time"
	"vittaindicator/internal/config"
	"vittaindicator/internal/utils"
)

func CheckExists() {

}

func DownloadFile(url string) {
	jar,_:=cookiejar.New(nil);
	client := &http.Client{
		Timeout: time.Minute*2,
		Jar: jar,
	}

	pHttpRequestHomePage,err:=http.NewRequest("GET","https://www.nseindia.com",nil);
	if err!=nil{
		log.Println("Unable to Make Request NSE homepage:", err)
		return
	}
	pHttpRequestHomePage.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)");
	pHttpRequestHomePage.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

	homeResp, err := client.Do(pHttpRequestHomePage)
	if err != nil {
		log.Println("Unable to fetch NSE homepage:", err)
		return
	}
	homeResp.Body.Close();


	pHttpRequest, err := http.NewRequest("GET", url, nil);
	if err != nil {
		log.Println("Unable to create request:", err);
		return
	}

	pHttpRequest.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)");
	pHttpRequest.Header.Set("Accept", "*/*");
	pHttpRequest.Header.Set("Connection", "keep-alive");
	pHttpRequest.Header.Set("Referer", "https://www.nseindia.com/");

	pHttpResponse, err := client.Do(pHttpRequest);
	if err != nil {
		log.Println("Unable to download:", err);
		return
	}
	defer pHttpResponse.Body.Close()

	if pHttpResponse.StatusCode != http.StatusOK {
		log.Printf("Bad HTTP status: %s\n", pHttpResponse.Status)
		return
	}

	currentWorkingDir, err := os.Getwd()
	if err != nil {
		log.Println("Unable to get working directory: ", err)
		return
	}

	storagePath := filepath.Join(currentWorkingDir, "storage", "download")

	err = os.MkdirAll(storagePath, 0755)
	if err != nil {
		log.Println("Unable to create download file: ", err)
	}

	fileName := filepath.Base(url)
	fileFullPath := filepath.Join(storagePath, fileName)

	pOsFile, err := os.Create(fileFullPath)
	if err != nil {
		log.Println("Fail while creating file :", err)
	}
	defer pOsFile.Close()

	_, err = io.Copy(pOsFile, pHttpResponse.Body)
	if err != nil {
		log.Println("Error while copying : ", err)
	}

	log.Println("Successfully download: ",url);

	if filepath.Ext(fileFullPath) == ".zip" {
		utils.ExtractCSVFromZip(fileFullPath, storagePath)
	}
}

func AutoDownload(cfg config.Config) {
	go scheduleDownload(cfg.Bhavcopyurl.NSE.URL, (cfg.Bhavcopyurl.NSE.Time))
	go scheduleDownload(cfg.Priceband.NSE.URL, cfg.Priceband.NSE.Time)

}

func scheduleDownload(url, timeStr string) {
	for {
		now := time.Now()
		parsedTime, err := time.Parse("15:04", timeStr)
		if err != nil {
			log.Println("Invalid time format:", timeStr)
			return
		}

		next := time.Date(now.Year(), now.Month(), now.Day(),
			parsedTime.Hour(), parsedTime.Minute(), 0, 0, now.Location())
		if next.Before(now) {
			next = next.Add(24 * time.Hour)
		}

		waitDuration := next.Sub(now);
		log.Printf("Next download for %s in %v\n", url, waitDuration);
		time.Sleep(waitDuration);

		DownloadFile(url)
	}
}

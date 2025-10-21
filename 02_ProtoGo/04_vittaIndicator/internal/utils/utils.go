package utils

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func ParseUrlWithPresentDate(url string) string {
	return url
}

func ExtractCSVFromZip(zipPath, destDir string) {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		log.Println("Failed to open zip:", err)
		return
	}
	defer r.Close()

	for _, f := range r.File {
		if filepath.Ext(f.Name) != ".csv" {
			continue
		}

		rc, _ := f.Open()
		outPath := filepath.Join(destDir, f.Name)
		outFile, _ := os.Create(outPath)
		_, _ = io.Copy(outFile, rc)
		rc.Close()
		outFile.Close()
		log.Println("Extracted CSV:", outPath)
	}
}
package aula_moodle

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func (mc *MoodleClient) DownloadPdf(fileUrl, filename string) (string, error) {
	fullURL := fmt.Sprintf("%s&token=%s", fileUrl, mc.token)

	resp, err := http.Get(fullURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	downloadDir := "downloads"
	if err := os.MkdirAll(downloadDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("error creando carpeta de descargas: %v", err)
	}

	safeFilename := filepath.Base(filename)
	savePath := filepath.Join(downloadDir, safeFilename)

	file, err := os.Create(savePath)
	if err != nil {
		return "", fmt.Errorf("error creando el archivo local: %v", err)
	}
	defer file.Close()

	if _, err = io.Copy(file, resp.Body); err != nil {
		return "", fmt.Errorf("error guardando el PDF: %v", err)
	}

	absPath, err := filepath.Abs(savePath)
	if err != nil {
		return savePath, nil
	}
	return absPath, nil
}

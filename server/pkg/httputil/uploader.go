package httputil

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func UploadHandler(r *http.Request, formName string) (*string, error) {
	file, handler, err := r.FormFile(formName)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileName := generateFileName(handler.Filename)

	// Create a new file in the server's upload directory
	// /project/uploads
	f, err := os.OpenFile(filepath.Join("uploads", fileName), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		slog.Error("UploadHandler.OpenFile", slog.String("details", err.Error()))
		return nil, err
	}
	defer f.Close()

	// Copy the file to the destination path
	_, err = io.Copy(f, file)
	if err != nil {
		slog.Error("UploadHandler.Copy", slog.String("details:", err.Error()))
		return nil, err
	}

	return &fileName, nil
}
func DeleteUpload(fileName string) {
	_ = os.Remove(filepath.Join("uploads", fileName))
}

func generateFileName(originalFileName string) string {
	extension := filepath.Ext(originalFileName)
	newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), extension)
	return newFileName
}

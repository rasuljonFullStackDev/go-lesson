package files

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func UploadFile(file multipart.File, header *multipart.FileHeader, folder string, allowed []string) (string, error) {
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(header.Filename))
	ext = strings.TrimPrefix(ext, ".")

	// Agar ruxsat etilgan formatlar bo‘lsa, tekshiramiz
	if len(allowed) > 0 {
		valid := false
		for _, a := range allowed {
			if ext == strings.ToLower(a) {
				valid = true
				break
			}
		}
		if !valid {
			return "", fmt.Errorf("❌ format '%s' ruxsat etilmagan", ext)
		}
	}

	// Saqlash joyi
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), header.Filename)
	filePath := filepath.Join(folder, fileName)

	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	return filePath, nil
}

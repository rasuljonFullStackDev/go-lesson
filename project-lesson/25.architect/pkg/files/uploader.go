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

// UploadFile - universal fayl yuklash
func UploadFile(file multipart.File, header *multipart.FileHeader, folder string, allowed []string, maxSize int64) (string, error) {
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(header.Filename))
	ext = strings.TrimPrefix(ext, ".")

	if maxSize > 0 && header.Size > maxSize {
		return "", fmt.Errorf("❌ fayl hajmi %d MB dan oshmasin", maxSize/1024/1024)
	}

	// format check
	if len(allowed) > 0 {
		valid := false
		for _, a := range allowed {
			if ext == strings.ToLower(a) {
				valid = true
				break
			}
		}
		if !valid {
			return "", fmt.Errorf("❌ '%s' formati ruxsat etilmagan", ext)
		}
	}

	os.MkdirAll(folder, os.ModePerm)

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

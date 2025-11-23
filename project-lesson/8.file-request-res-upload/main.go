package main

import (
	// "encoding/json"
	// "net/http"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"bytes"
)


// | Amal                     | Funksiya                        | Tavsif                        |
// | ------------------------ | ------------------------------- | ----------------------------- |
// | Papka yaratish           | `os.Mkdir`, `os.MkdirAll`       | Fayl tizimida papka yaratadi  |
// | Papka o‘qish             | `os.ReadDir`                    | Papkadagi fayllarni qaytaradi |
// | Rekursiv o‘qish          | `filepath.WalkDir`              | Ichma-ich fayllarni chiqaradi |
// | Fayl o‘qish              | `os.ReadFile`                   | Faylni o‘qish                 |
// | Fayl yozish              | `os.Create`, `file.WriteString` | Yangi faylga yozish           |
// | Fayl o‘chirish           | `os.Remove`, `os.RemoveAll`     | Fayl yoki papkani o‘chirish   |
// | Request hajmini cheklash | `http.MaxBytesReader`           | Umumiy POST limit             |
// | RAM limit                | `r.ParseMultipartForm`          | RAMda saqlash hajmi           |
// | MIME aniqlash            | `http.DetectContentType`        | Fayl turini tekshirish        |
// | Temp fayl tozalash       | `r.MultipartForm.RemoveAll()`   | Yuklashdan keyin tozalaydi    |


func main() {
	// papaka yaratish
	// err := os.Mkdir("files", 0755)
	// if err != nil  && !os.IsExist(err) {
	// 	fmt.Println("files folder mavjud ")
	// 	return
	// 	// panic(err)
	// }
	// fmt.Println("files folder created ")

	// ichma ichma icha
	err := os.MkdirAll("files/a/b/c", 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Println("folders mavjud")
		return
	}
	fmt.Println("folders yaratildi")

	// falyllar joyichi oqish
	entries, err1 := os.ReadDir("files")
	if err1 != nil {
		panic(err1)
	}

	for _, e := range entries {
		if e.IsDir() {
			fmt.Println("[DIR] ", e.Name())
		} else {
			fmt.Println("      ", e.Name())
		}
	}
	// Papka mavjudligini tekshirish
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		fmt.Println("❌ Papka mavjud emas")
	} else {
		fmt.Println("✅ Papka mavjud")
	}

	os.Remove("uploads/temp")      // faqat bo‘sh papka
	os.RemoveAll("uploads/images") // ichidagilari bilan birga o‘chiradi

	// Fayl yaratish va yozish
	os.Mkdir("uploads", 0755)
	file, err := os.Create("uploads/hello.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("Salom, bu Go orqali yozilgan fayl!\n")
	fmt.Println("✅ Fayl yaratildi va yozildi.")

	// Faylni o‘qish
	data, err := os.ReadFile("uploads/hello.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	// Fayl haqida ma’lumot
	info, _ := os.Stat("uploads/hello.txt")
	fmt.Println("Hajm:", info.Size(), "bayt")
	fmt.Println("Oxirgi o‘zgarish:", info.ModTime())
	_ = os.MkdirAll("uploads", 0755)

	http.HandleFunc("/upload", uploadHandler)
	fmt.Println("Server: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	json.NewEncoder(w).Encode(map[string]string{"message": "Hello World"})
	// })
	// http.ListenAndServe(":8080", nil)

}


func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST", http.StatusMethodNotAllowed)
		return
	}

	// 1️⃣ Umumiy request hajmini cheklash (masalan, 10MB)
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)

	// 2️⃣ Multipart parsing (RAMda 8MB saqlansin, ortig‘i temp fayl)
	if err := r.ParseMultipartForm(8 << 20); err != nil {
		http.Error(w, "Too large or invalid multipart", http.StatusRequestEntityTooLarge)
		return
	}
	defer r.MultipartForm.RemoveAll() // temp fayllarni tozalash

	// 3️⃣ Faylni olish
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "file required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 4️⃣ Fayl nomini xavfsiz qilish
	filename := filepath.Base(header.Filename)

	// 5️⃣ MIME sniff (birinchi 512 bayt)
	var sniff bytes.Buffer
	tee := io.TeeReader(file, &sniff)
	head := make([]byte, 512)
	n, _ := io.ReadFull(tee, head)
	mime := http.DetectContentType(head[:n])

	allowed := map[string]bool{"image/jpeg": true, "image/png": true}
	if !allowed[mime] {
		http.Error(w, "Unsupported file type", http.StatusUnsupportedMediaType)
		return
	}

	// 6️⃣ Faylni saqlash
	dst, _ := os.Create(filepath.Join("uploads", filename))
	defer dst.Close()
	io.Copy(dst, &sniff)
	io.Copy(dst, file)

	// 7️⃣ JSON javob
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"message":   "uploaded",
		"filename":  filename,
		"mime_type": mime,
	})
}
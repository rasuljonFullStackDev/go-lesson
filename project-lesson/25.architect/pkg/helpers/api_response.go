package helpers

import (
	"encoding/json"
	"net/http"
)

// JsonResponse — yagona response tuzilmasi
type JsonResponse struct {
	Result interface{} `json:"result"`
	Errors interface{} `json:"errors"`
}

// Success — muvaffaqiyatli javob
func Success(w http.ResponseWriter, result interface{}, code int) {
	if code == 0 {
		code = http.StatusOK
	}
	response := JsonResponse{
		Result: result,
		Errors: nil,
	}
	writeJSON(w, response, code)
}

// Error — xato javob
func Error(w http.ResponseWriter, err interface{}, code int) {
	if code == 0 {
		code = http.StatusBadRequest
	}
	response := JsonResponse{
		Result: nil,
		Errors: err,
	}
	writeJSON(w, response, code)
}

// NotAccess — ruxsat yo‘q holat
func NotAccess(w http.ResponseWriter, message string) {
	Error(w, map[string]string{
		"message": message,
	}, http.StatusMethodNotAllowed)
}

// Yordamchi funksiya JSON yozish uchun
func writeJSON(w http.ResponseWriter, data JsonResponse, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

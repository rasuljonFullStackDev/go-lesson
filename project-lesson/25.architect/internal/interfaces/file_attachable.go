package interfaces

// Fayl maydonlari va formatlarni model darajasida aniqlash uchun interfeys
type FileAttachable interface {
	FileFields() map[string][]string  // ustun -> ruxsat formatlari
	MaxFileSize() int64               // bayt (masalan 5MB)
}

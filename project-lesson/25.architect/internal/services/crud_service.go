package services

import (
	"fmt"
	"net/http"
	"reflect"

	"25.architect/internal/config"
	"25.architect/internal/interfaces"
	"25.architect/pkg/files"
)

type CrudService[T any] struct{}

// CREATE
func (s *CrudService[T]) Create(model *T, r *http.Request) error {
	if err := s.handleFiles(model, r); err != nil {
		return err
	}
	return config.DB.Create(model).Error
}

// UPDATE
func (s *CrudService[T]) Update(id uint, model *T, r *http.Request) error {
	if err := s.handleFiles(model, r); err != nil {
		return err
	}
	return config.DB.Model(model).Where("id = ?", id).Updates(model).Error
}

// GET ALL
func (s *CrudService[T]) FindAll(list *[]T) error {
	return config.DB.Find(list).Error
}

// DELETE
func (s *CrudService[T]) Delete(id uint, model *T) error {
	return config.DB.Delete(model, id).Error
}

// FILE HANDLER
func (s *CrudService[T]) handleFiles(model *T, r *http.Request) error {
	m, ok := any(model).(interfaces.FileAttachable)
	if !ok {
		return nil
	}

	fileFields := m.FileFields()
	maxSize := m.MaxFileSize()
	if len(fileFields) == 0 {
		return nil
	}

	r.ParseMultipartForm(maxSize + (2 << 20)) // buffer

	for field, allowed := range fileFields {
		file, header, err := r.FormFile(field)
		if err == http.ErrMissingFile {
			continue
		}
		if err != nil {
			return fmt.Errorf("❌ fayl olishda xato: %v", err)
		}

		path, err := files.UploadFile(file, header, fmt.Sprintf("uploads/%s", field), allowed, maxSize)
		if err != nil {
			return err
		}

		v := reflect.ValueOf(model).Elem()
		f := v.FieldByName(field)
		if f.IsValid() && f.CanSet() {
			f.SetString(path)
		}
	}
	return nil
}

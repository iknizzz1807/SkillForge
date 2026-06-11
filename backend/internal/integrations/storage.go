package integrations

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Nhiệm vụ: Tích hợp lưu trữ file (local hoặc cloud)
// Liên kết:
// - Được gọi từ services/file_service.go để upload/download file
// Vai trò trong flow:
// - Xử lý lưu trữ file cho avatar, project documents, v.v.
//
// Current implementation: Local filesystem storage under ./storage/
// TODO: Add cloud storage backends (AWS S3, GCS, etc.)

// StorageBackend defines the interface for file storage operations
type StorageBackend interface {
	Upload(filename string, reader io.Reader) error
	Download(filename string) (io.ReadCloser, error)
	Delete(filename string) error
	Exists(filename string) bool
}

// LocalStorage stores files on the local filesystem
type LocalStorage struct {
	BasePath string
}

// NewLocalStorage creates a LocalStorage instance with the given base path.
// The base path defaults to "./storage" if empty.
func NewLocalStorage(basePath string) *LocalStorage {
	if basePath == "" {
		basePath = "./storage"
	}
	if err := os.MkdirAll(basePath, 0755); err != nil {
		log.Printf("[Storage] Warning: could not create base path %s: %v", basePath, err)
	}
	return &LocalStorage{BasePath: basePath}
}

// Upload writes a file to the local filesystem
func (s *LocalStorage) Upload(filename string, reader io.Reader) error {
	fullPath := filepath.Join(s.BasePath, filename)
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	dst, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, reader)
	return err
}

// Download reads a file from the local filesystem
func (s *LocalStorage) Download(filename string) (io.ReadCloser, error) {
	fullPath := filepath.Join(s.BasePath, filename)
	return os.Open(fullPath)
}

// Delete removes a file from the local filesystem
func (s *LocalStorage) Delete(filename string) error {
	return os.Remove(filepath.Join(s.BasePath, filename))
}

// Exists checks if a file exists on the local filesystem
func (s *LocalStorage) Exists(filename string) bool {
	_, err := os.Stat(filepath.Join(s.BasePath, filename))
	return err == nil
}

// NewStorageClient returns the default storage backend.
// Currently returns a LocalStorage instance. Swap this function
// to return an S3-compatible backend for production deployments.
func NewStorageClient(basePath string) StorageBackend {
	if basePath == "" {
		basePath = "./storage"
	}
	storage := NewLocalStorage(basePath)
	log.Printf("[Storage] Initialized local storage at %s", basePath)
	return storage
}

// ErrNotFound is returned when a file is not found
var ErrNotFound = errors.New("file not found")

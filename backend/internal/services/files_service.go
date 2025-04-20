// internal/services/file_service.go
// File này chịu trách nhiệm lưu file vào trong thư mục, không chịu trách nhiệm thao tác trên cơ sở dữ liệu
package services

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/iknizzz1807/SkillForge/internal/repositories"
)

// Định nghĩa đường dẫn lưu avatar ở đây hoặc trong config
// Thêm đường dẫn ví dụ dành cho file trong messange hoặc file đính kèm trong project describtion chẳng hạn
const AvatarStoragePath = "./storage/avatars"

type FileService struct {
	fileRepo *repositories.FileRepository
	// Có thể thêm các dependencies khác nếu cần (ví dụ: config)
}

// NewFileService khởi tạo FileService
func NewFileService() *FileService {
	return &FileService{}
}

// SaveAvatar lưu file avatar, xóa file cũ và trả về tên file mới
// Input: userID string, file multipart.File, header *multipart.FileHeader
// Return: string (tên file mới bao gồm extension), error
func (s *FileService) SaveAvatar(userID string, file multipart.File, header *multipart.FileHeader) (string, error) {
	if userID == "" {
		return "", errors.New("user ID is required")
	}
	if file == nil || header == nil {
		return "", errors.New("file data is missing")
	}

	// 1. Xác định phần mở rộng file
	ext := filepath.Ext(header.Filename)
	if ext == "" {
		return "", errors.New("could not determine file extension")
	}
	ext = strings.ToLower(ext)
	// Kiểm tra loại file đơn giản
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
		return "", errors.New("invalid file type. Only JPG, PNG, GIF allowed")
	}

	// 2. Tạo tên file mới
	newFilename := userID + ext

	// 3. Đảm bảo thư mục lưu trữ tồn tại
	err := os.MkdirAll(AvatarStoragePath, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating storage directory %s: %v\n", AvatarStoragePath, err)
		return "", errors.New("could not create storage directory")
	}

	// 4. Tạo đường dẫn file đầy đủ
	filePath := filepath.Join(AvatarStoragePath, newFilename)

	// 5. Xóa file avatar cũ (nếu có extension khác)
	existingFiles, _ := filepath.Glob(filepath.Join(AvatarStoragePath, userID+".*"))
	for _, f := range existingFiles {
		if f != filePath { // Chỉ xóa nếu tên file khác file mới
			fmt.Printf("Attempting to remove old avatar: %s\n", f)
			os.Remove(f) // Bỏ qua lỗi nếu xóa không được
		}
	}

	// 6. Tạo file đích
	dst, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", filePath, err)
		return "", errors.New("could not save file")
	}
	defer dst.Close() // Đảm bảo file được đóng

	// 7. Copy dữ liệu
	_, err = io.Copy(dst, file)
	if err != nil {
		fmt.Printf("Error copying file content to %s: %v\n", filePath, err)
		// Cố gắng xóa file vừa tạo nếu copy lỗi
		os.Remove(filePath)
		return "", errors.New("could not write file content")
	}

	// 8. Trả về tên file mới thành công
	return newFilename, nil
}

// GetAvatarFilePath (Optional) - Trả về đường dẫn đầy đủ tới file avatar
// Input: filename string (ví dụ: "userID.png")
// Return: string (đường dẫn tuyệt đối), error
func (s *FileService) GetAvatarFilePath(filename string) (string, error) {
	if filename == "" {
		return "", errors.New("filename is required")
	}
	// Security check đơn giản
	if strings.Contains(filename, "..") || strings.Contains(filename, "/") || strings.Contains(filename, "\\") {
		return "", errors.New("invalid filename")
	}

	filePath := filepath.Join(AvatarStoragePath, filename)

	// Kiểm tra file tồn tại
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", errors.New("avatar not found")
	} else if err != nil {
		fmt.Printf("Error accessing file %s: %v\n", filePath, err)
		return "", errors.New("could not access avatar file")
	}

	return filePath, nil
}

// func (f *FileService) UploadFile(ctx context.Context, FileName string, FolderName string, UploadedByID string) (*models.File, error) {
// 	if FileName == "" || FolderName == "" || UploadedByID == "" {
// 		return nil, errors.New("invalid input data")
// 	}

// 	file := &models.File{
// 		ID:           utils.GenerateUUID(),
// 		FileName:     FileName,
// 		Path:         FolderName + "/" + FileName,
// 		UploadedAt:   time.Now(),
// 		UploadedByID: UploadedByID,
// 	}

// 	fileRepo := repositories.NewFileRepository(f.fileCollection)
// 	err := fileRepo.InsertFile(ctx, file)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return file, nil
// }

// func (f *FileService) GetFileByID(fileID string) (*models.File, error) {
// 	fileRepo := repositories.NewFileRepository(f.fileCollection)
// 	return fileRepo.GetFileByID(fileID)
// }

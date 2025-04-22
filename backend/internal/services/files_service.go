package services

import (
	"context"
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
const AvatarStoragePath = "./storage/avatars"

type FileService struct {
	// fileRepo *repositories.FileRepository
	userRepo *repositories.UserRepository
	// Có thể thêm các dependencies khác nếu cần (ví dụ: config)
}

// NewFileService khởi tạo FileService với dependencies
func NewFileService(userRepo *repositories.UserRepository) *FileService {
	return &FileService{
		// fileRepo: fileRepo,
		userRepo: userRepo,
	}
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

// GetAvatarFilePath trả về đường dẫn đầy đủ tới file avatar
// Input: filename string (ví dụ: "userID.png")
// Return: string (đường dẫn tuyệt đối), error
func (s *FileService) GetAvatarFilePath(filename string) (string, error) {
	if filename == "" {
		return "", errors.New("filename is required")
	}
	// Security check đơn giản
	if strings.Contains(filename, "..") || strings.ContainsAny(filename, "/\\") {
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

// FindAvatarByUserID tìm file avatar dựa trên ID người dùng
// - Chuyển từ handler sang service để đảm bảo separation of concerns
func (s *FileService) FindAvatarByUserID(userID string) (string, error) {
	if userID == "" {
		return "", errors.New("user ID is required")
	}

	// Security check
	if strings.Contains(userID, "..") || strings.ContainsAny(userID, "/\\") {
		return "", errors.New("invalid user ID")
	}

	// 1. Tìm thông tin user từ database thông qua UserRepository
	user, err := s.userRepo.FindUserByID(context.Background(), userID)
	if err != nil {
		return "", fmt.Errorf("failed to find user: %w", err)
	}

	// 2. Kiểm tra xem user có avatar hay không
	if user.AvatarName == "" {
		// Nếu không có avatar name trong DB, trả về default avatar hoặc error
		return "", fmt.Errorf("user has no avatar")
	}

	// 3. Tạo đường dẫn đầy đủ từ tên file avatar đã lưu trong DB
	avatarPath := filepath.Join(AvatarStoragePath, user.AvatarName)

	// 4. Kiểm tra file có tồn tại không
	if _, err := os.Stat(avatarPath); os.IsNotExist(err) {
		return "", fmt.Errorf("avatar file does not exist: %s", user.AvatarName)
	} else if err != nil {
		return "", fmt.Errorf("error accessing avatar file: %w", err)
	}

	// 5. Trả về đường dẫn đầy đủ
	return avatarPath, nil
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

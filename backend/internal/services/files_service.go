package services

import (
	"context"
	"errors"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/utils"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type FileService struct {
	fileCollection *mongo.Collection
}

func NewFileService(fileCollection *mongo.Collection) *FileService {
	return &FileService{fileCollection: fileCollection}
}

func (f *FileService) UploadFile(ctx context.Context, FileName string, FolderName string, UploadedByID string) (*models.File, error) {
	if FileName == "" || FolderName == "" || UploadedByID == "" {
		return nil, errors.New("invalid input data")
	}

	file := &models.File{
		ID:         utils.GenerateUUID(),
		FileName:   FileName,
		Path:       FolderName + "/" + FileName,
		UploadedAt: time.Now(),
		UploadedByID:     UploadedByID,
	}

	fileRepo := repositories.NewFileRepository(f.fileCollection)
	err := fileRepo.InsertFile(ctx, file)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (f *FileService) GetFileByID(fileID string) (*models.File, error) {
	fileRepo := repositories.NewFileRepository(f.fileCollection)
	return fileRepo.GetFileByID(fileID)
}

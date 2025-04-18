package repositories

import (
	"context"
	"errors"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type FileRepository struct {
	fileCollection *mongo.Collection
}

func NewFileRepository(fileCollection *mongo.Collection) *FileRepository {
	return &FileRepository{fileCollection: fileCollection}
}

func (f *FileRepository) InsertFile(ctx context.Context, file *models.File) error {
	_, err := f.fileCollection.InsertOne(ctx, file)
	if err != nil {
		return err
	}
	return nil
}

func (f *FileRepository) GetFileByID(fileID string) (*models.File, error) {
	var file models.File
	err := f.fileCollection.FindOne(context.Background(), bson.M{"_id": fileID}).Decode(&file)
	if err != nil {
		return nil, errors.New("file not found")
	}
	return &file, nil
}
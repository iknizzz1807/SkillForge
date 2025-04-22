package repositories

import (
	"context"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BusinessInfoRepository struct {
	collection *mongo.Collection
}

// NewBusinessInfoRepository initializes a new BusinessInfoRepository
func NewBusinessInfoRepository(db *mongo.Database) *BusinessInfoRepository {
	return &BusinessInfoRepository{
		collection: db.Collection("business_info"),
	}
}

// FindByUserID retrieves business info by user ID
func (r *BusinessInfoRepository) FindByUserID(ctx context.Context, userID string) (*models.BusinessInfo, error) {
	filter := bson.M{"_id": userID}
	var businessInfo models.BusinessInfo

	err := r.collection.FindOne(ctx, filter).Decode(&businessInfo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No business info found
		}
		return nil, err
	}

	return &businessInfo, nil
}

// UpsertBusinessInfo creates or updates business info
func (r *BusinessInfoRepository) UpsertBusinessInfo(ctx context.Context, businessInfo *models.BusinessInfo) error {
	businessInfo.UpdatedAt = time.Now()

	// If this is a new record, set creation time
	if businessInfo.CreatedAt.IsZero() {
		businessInfo.CreatedAt = businessInfo.UpdatedAt
	}

	filter := bson.M{"_id": businessInfo.UserID}
	update := bson.M{"$set": businessInfo}
	opts := options.Update().SetUpsert(true)

	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	return err
}

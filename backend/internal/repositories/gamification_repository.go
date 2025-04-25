package repositories

import (
	"context"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GamificationRepository struct {
	levelCollection *mongo.Collection
	skillCollection *mongo.Collection
}

func NewGamificationRepository(db *mongo.Database) *GamificationRepository {
	return &GamificationRepository{
		levelCollection: db.Collection("user_levels"),
		skillCollection: db.Collection("user_skills"),
	}
}

func (r *GamificationRepository) GetUserLevel(ctx context.Context, userID string) (*models.UserLevel, error) {
	var level models.UserLevel
	err := r.levelCollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&level)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &level, err
}

func (r *GamificationRepository) UpdateUserLevel(ctx context.Context, level *models.UserLevel) error {
	level.UpdatedAt = time.Now()
	_, err := r.levelCollection.UpdateOne(
		ctx,
		bson.M{"user_id": level.UserID},
		bson.M{"$set": level},
	)
	return err
}

func (r *GamificationRepository) GetUserSkill(ctx context.Context, userID string, skillName string) (*models.UserSkill, error) {
	var skill models.UserSkill
	err := r.skillCollection.FindOne(ctx, bson.M{
		"user_id": userID,
		"skill":   skillName,
	}).Decode(&skill)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &skill, err
}

func (r *GamificationRepository) UpdateUserSkill(ctx context.Context, skill *models.UserSkill) error {
	skill.UpdatedAt = time.Now()
	_, err := r.skillCollection.UpdateOne(
		ctx,
		bson.M{
			"user_id": skill.UserID,
			"skill":   skill.Skill,
		},
		bson.M{"$set": skill},
	)
	return err
}

func (r *GamificationRepository) GetAllUserSkills(ctx context.Context, userID string) ([]models.UserSkill, error) {
	cursor, err := r.skillCollection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var skills []models.UserSkill
	if err := cursor.All(ctx, &skills); err != nil {
		return nil, err
	}
	return skills, nil
}

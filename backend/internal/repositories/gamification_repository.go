package repositories

import (
	"context"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GamificationRepository struct {
	levelCollection *mongo.Collection
	skillCollection *mongo.Collection
}

func NewGamificationRepository(db *mongo.Database) *GamificationRepository {
	levelCollection := db.Collection("user_levels")
	skillCollection := db.Collection("user_skills")

	_, _ = levelCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{{Key: "user_id", Value: 1}},
	})
	_, _ = skillCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{{Key: "user_id", Value: 1}, {Key: "skill", Value: 1}},
	})

	return &GamificationRepository{
		levelCollection: levelCollection,
		skillCollection: skillCollection,
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
		options.Update().SetUpsert(true),
	)
	return err
}

// AddXPAtomic atomically increments xp_current by the given amount using $inc.
// xpNeedOnInsert is the XP needed for the first level (caller provides from service config).
// Returns the updated document after increment.
func (r *GamificationRepository) AddXPAtomic(ctx context.Context, userID string, xp int, xpNeedOnInsert int) (*models.UserLevel, error) {
	after := options.After
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(after)
	var level models.UserLevel
	err := r.levelCollection.FindOneAndUpdate(
		ctx,
		bson.M{"user_id": userID},
		bson.M{
			"$inc": bson.M{"xp_current": xp},
			"$setOnInsert": bson.M{
				"id":         utils.GenerateUUID(),
				"user_id":    userID,
				"level":      1,
				"xp_need":    xpNeedOnInsert,
				"created_at": time.Now(),
			},
			"$set": bson.M{"updated_at": time.Now()},
		},
		opts,
	).Decode(&level)
	if err != nil {
		return nil, err
	}
	return &level, nil
}

// LevelUpAtomically atomically increases level and adjusts XP. This should be called
// after detecting a level-up condition from the post-increment document.
func (r *GamificationRepository) LevelUpAtomically(ctx context.Context, userID string, currentXP int, newLevel int, newXPNeed int) error {
	_, err := r.levelCollection.UpdateOne(
		ctx,
		bson.M{"user_id": userID},
		bson.M{
			"$set": bson.M{
				"level":      newLevel,
				"xp_current": currentXP,
				"xp_need":    newXPNeed,
				"updated_at": time.Now(),
			},
		},
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
		options.Update().SetUpsert(true),
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

// AddSkillPointAtomic atomically increments point_current for a user skill.
// pointNeededOnInsert is the initial point_needed value if the document is created.
func (r *GamificationRepository) AddSkillPointAtomic(ctx context.Context, userID string, skillName string, pointNeededOnInsert int) (*models.UserSkill, error) {
	after := options.After
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(after)
	var skill models.UserSkill
	err := r.skillCollection.FindOneAndUpdate(
		ctx,
		bson.M{"user_id": userID, "skill": skillName},
		bson.M{
			"$inc": bson.M{"point_current": 1},
			"$setOnInsert": bson.M{
				"id":           utils.GenerateUUID(),
				"user_id":      userID,
				"skill":        skillName,
				"level":        1,
				"point_needed": pointNeededOnInsert,
				"created_at":   time.Now(),
			},
			"$set": bson.M{"updated_at": time.Now()},
		},
		opts,
	).Decode(&skill)
	if err != nil {
		return nil, err
	}
	return &skill, nil
}

// SkillLevelUpAtomically atomically updates level and current points.
func (r *GamificationRepository) SkillLevelUpAtomically(ctx context.Context, userID string, skillName string, currentPoints int, newLevel int, newPointsNeeded int) error {
	_, err := r.skillCollection.UpdateOne(
		ctx,
		bson.M{"user_id": userID, "skill": skillName},
		bson.M{
			"$set": bson.M{
				"level":         newLevel,
				"point_current": currentPoints,
				"point_needed":  newPointsNeeded,
				"updated_at":    time.Now(),
			},
		},
	)
	return err
}

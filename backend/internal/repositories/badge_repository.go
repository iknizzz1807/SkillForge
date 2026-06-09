package repositories

import (
	"context"
	"errors"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// BadgeRepository xử lý tương tác với collection "badges" và "user_badges" trong MongoDB
type BadgeRepository struct {
	badgeCollection     *mongo.Collection
	userBadgeCollection *mongo.Collection
}

// NewBadgeRepository khởi tạo BadgeRepository với database
func NewBadgeRepository(db *mongo.Database) *BadgeRepository {
	return &BadgeRepository{
		badgeCollection:     db.Collection("badges"),
		userBadgeCollection: db.Collection("user_badges"),
	}
}

// InsertBadge thêm badge mới vào collection "badges"
func (r *BadgeRepository) InsertBadge(ctx context.Context, badge *models.Badge) error {
	_, err := r.badgeCollection.InsertOne(ctx, badge)
	return err
}

// FindBadgeByID tìm badge theo ID
func (r *BadgeRepository) FindBadgeByID(ctx context.Context, badgeID string) (*models.Badge, error) {
	var filter bson.M
	if isValidObjectID(badgeID) {
		objID, _ := primitive.ObjectIDFromHex(badgeID)
		filter = bson.M{"_id": objID}
	} else {
		filter = bson.M{"_id": badgeID}
	}

	var badge models.Badge
	err := r.badgeCollection.FindOne(ctx, filter).Decode(&badge)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("badge not found")
		}
		return nil, err
	}

	return &badge, nil
}

// FindAllBadges lấy tất cả badges
func (r *BadgeRepository) FindAllBadges(ctx context.Context) ([]*models.Badge, error) {
	cursor, err := r.badgeCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var badges []*models.Badge
	for cursor.Next(ctx) {
		var badge models.Badge
		if err := cursor.Decode(&badge); err != nil {
			return nil, err
		}
		badges = append(badges, &badge)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return badges, nil
}

// FindBadgesByType lấy badges theo type
func (r *BadgeRepository) FindBadgesByType(ctx context.Context, badgeType string) ([]*models.Badge, error) {
	cursor, err := r.badgeCollection.Find(ctx, bson.M{"type": badgeType})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var badges []*models.Badge
	for cursor.Next(ctx) {
		var badge models.Badge
		if err := cursor.Decode(&badge); err != nil {
			return nil, err
		}
		badges = append(badges, &badge)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return badges, nil
}

// AwardBadgeToUser gán badge cho user
func (r *BadgeRepository) AwardBadgeToUser(ctx context.Context, userBadge *models.UserBadge) error {
	_, err := r.userBadgeCollection.InsertOne(ctx, userBadge)
	return err
}

// FindBadgesByIDs lấy nhiều badges theo danh sách ID (batch query)
func (r *BadgeRepository) FindBadgesByIDs(ctx context.Context, badgeIDs []string) ([]*models.Badge, error) {
	if len(badgeIDs) == 0 {
		return nil, nil
	}

	var ids []interface{}
	for _, id := range badgeIDs {
		if isValidObjectID(id) {
			objID, _ := primitive.ObjectIDFromHex(id)
			ids = append(ids, objID)
		} else {
			ids = append(ids, id)
		}
	}

	cursor, err := r.badgeCollection.Find(ctx, bson.M{"_id": bson.M{"$in": ids}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var badges []*models.Badge
	for cursor.Next(ctx) {
		var badge models.Badge
		if err := cursor.Decode(&badge); err != nil {
			return nil, err
		}
		badges = append(badges, &badge)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return badges, nil
}

// FindUserBadges lấy tất cả badge của một user
func (r *BadgeRepository) FindUserBadges(ctx context.Context, userID string) ([]*models.UserBadge, error) {
	cursor, err := r.userBadgeCollection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var userBadges []*models.UserBadge
	var badgeIDs []string
	badgeIndexMap := make(map[string]int)

	for cursor.Next(ctx) {
		var userBadge models.UserBadge
		if err := cursor.Decode(&userBadge); err != nil {
			return nil, err
		}
		badgeIDs = append(badgeIDs, userBadge.BadgeID)
		badgeIndexMap[userBadge.BadgeID] = len(userBadges)
		userBadges = append(userBadges, &userBadge)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(badgeIDs) > 0 {
		badges, err := r.FindBadgesByIDs(ctx, badgeIDs)
		if err == nil {
			for _, badge := range badges {
				if idx, ok := badgeIndexMap[badge.ID]; ok {
					userBadges[idx].Badge = badge
				}
			}
		}
	}

	return userBadges, nil
}

// HasUserEarnedBadge kiểm tra xem user đã đạt được badge chưa
func (r *BadgeRepository) HasUserEarnedBadge(ctx context.Context, userID string, badgeID string) (bool, error) {
	count, err := r.userBadgeCollection.CountDocuments(ctx, bson.M{
		"user_id":  userID,
		"badge_id": badgeID,
	})

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

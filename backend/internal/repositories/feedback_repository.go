package repositories

import (
    "context"
    "time"

    "github.com/iknizzz1807/SkillForge/internal/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

type FeedbackRepository struct {
    collection *mongo.Collection
}

func NewFeedbackRepository(db *mongo.Database) *FeedbackRepository {
    return &FeedbackRepository{
        collection: db.Collection("feedbacks"),
    }
}

func (r *FeedbackRepository) CreateFeedback(ctx context.Context, feedback *models.Feedback) error {
    feedback.CreatedAt = time.Now()
    _, err := r.collection.InsertOne(ctx, feedback)
    return err
}

func (r *FeedbackRepository) GetFeedbacksForUser(ctx context.Context, userID string, feedbackType string) ([]*models.Feedback, error) {
    cursor, err := r.collection.Find(ctx, bson.M{
        "to": userID,
        "type": feedbackType,
    })
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var feedbacks []*models.Feedback
    if err = cursor.All(ctx, &feedbacks); err != nil {
        return nil, err
    }
    return feedbacks, nil
}

func (r *FeedbackRepository) GetAverageRating(ctx context.Context, userID string) (float64, error) {
    filter := bson.M{
		"to": userID,	
	}

	pipeline := []bson.M{
		{
			"$match": filter,
		},
		{
			"$group": bson.M{
				"_id": nil,
				"averageRating": bson.M{
					"$avg": "$rating",
				},			
			},
		},
	}

    cursor, err := r.collection.Aggregate(ctx, pipeline)
    if err != nil {
        return 0, err
    }
    defer cursor.Close(ctx)

    var result []struct {
        AverageRating float64 `bson:"averageRating"`
    }
    
    if err = cursor.All(ctx, &result); err != nil {
        return 0, err
    }

    if len(result) == 0 {
        return 0, nil
    }
    return result[0].AverageRating, nil
}

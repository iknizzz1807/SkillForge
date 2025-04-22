package services

import (
	"context"
	"errors"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
)

type BusinessInfoService struct {
	businessInfoRepo *repositories.BusinessInfoRepository
	userRepo         *repositories.UserRepository
}

// NewBusinessInfoService initializes a new BusinessInfoService
func NewBusinessInfoService(businessInfoRepo *repositories.BusinessInfoRepository, userRepo *repositories.UserRepository) *BusinessInfoService {
	return &BusinessInfoService{
		businessInfoRepo: businessInfoRepo,
		userRepo:         userRepo,
	}
}

// GetBusinessInfo retrieves business info for a user
func (s *BusinessInfoService) GetBusinessInfo(userID string) (*models.BusinessInfo, error) {
	ctx := context.Background()

	// First, check if the user exists and is a business
	user, err := s.userRepo.FindUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	if user.Role != "business" {
		return nil, errors.New("only business users have business information")
	}

	// Get business info
	businessInfo, err := s.businessInfoRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// If no business info exists yet, return a default empty one
	if businessInfo == nil {
		return &models.BusinessInfo{
			UserID:    userID,
			CreatedAt: time.Now(),
		}, nil
	}

	return businessInfo, nil
}

// UpdateBusinessInfo creates or updates business info
func (s *BusinessInfoService) UpdateBusinessInfo(userID, companyType, founded, companySize, website, aboutUs string) (*models.BusinessInfo, error) {
	ctx := context.Background()

	// First, check if the user exists and is a business
	user, err := s.userRepo.FindUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	if user.Role != "business" {
		return nil, errors.New("only business users can update business information")
	}

	// Get existing business info or create new one
	businessInfo, err := s.businessInfoRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if businessInfo == nil {
		businessInfo = &models.BusinessInfo{
			UserID:    userID,
			CreatedAt: time.Now(),
		}
	}

	// Update fields
	businessInfo.CompanyType = companyType
	businessInfo.Founded = founded
	businessInfo.CompanySize = companySize
	businessInfo.Website = website
	businessInfo.AboutUs = aboutUs
	businessInfo.UpdatedAt = time.Now()

	// Save to database
	err = s.businessInfoRepo.UpsertBusinessInfo(ctx, businessInfo)
	if err != nil {
		return nil, err
	}

	// Also update website in the user model if provided
	if website != "" {
		user.Website = website
		err = s.userRepo.UpdateUser(ctx, user)
		if err != nil {
			return nil, err
		}
	}

	return businessInfo, nil
}

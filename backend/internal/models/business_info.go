package models

import "time"

// BusinessInfo represents additional information about a business user
type BusinessInfo struct {
	// ID is the unique identifier (same as the user ID)
	UserID string `json:"user_id" bson:"_id"`

	// CompanyType describes the category of the company
	CompanyType string `json:"company_type" bson:"company_type"`

	// Founded is the year the company was established
	Founded string `json:"founded" bson:"founded"`

	// CompanySize represents the number of employees
	CompanySize string `json:"company_size" bson:"company_size"`

	// Website is the company's website URL
	Website string `json:"website" bson:"website"`

	// AboutUs is a description of the company
	AboutUs string `json:"about_us" bson:"about_us"`

	// CreatedAt is the time when the business info was created
	CreatedAt time.Time `json:"created_at" bson:"created_at"`

	// UpdatedAt is the time when the business info was last updated
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

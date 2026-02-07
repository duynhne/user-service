package domain

import "context"

// UserRepository defines the interface for user data access
type UserRepository interface {
	GetUser(ctx context.Context, id string) (*User, error)
	GetProfileByUserID(ctx context.Context, userID int) (*UserProfile, error)
	CreateUserProfile(ctx context.Context, userID int, firstName, lastName string) (int, error)
	UpdateUserProfile(ctx context.Context, userID int, firstName, lastName, phone string) (bool, error)
	CheckProfileExists(ctx context.Context, userID int) (bool, error)
	UpsertUserProfile(ctx context.Context, userID int, firstName, lastName, phone string) error
}

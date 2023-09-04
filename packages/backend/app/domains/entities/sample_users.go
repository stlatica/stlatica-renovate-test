// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entities

import (
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
)

// SampleUser is SampleUser entity object.
type SampleUser struct { // ulid
	UserID string `json:"user_id"`
	// Unix time
	CreatedAt types.UnixTime `json:"created_at"`
	// Unix time
	UpdatedAt types.UnixTime `json:"updated_at"`
}

// GetUserID is get user_id value, if receiver is nil, returns the specified value.
func (m *SampleUser) GetUserID() string {
	if m == nil {
		return ""
	}
	return m.UserID
}

// GetCreatedAt is get created_at value, if receiver is nil, returns the specified value.
func (m *SampleUser) GetCreatedAt() types.UnixTime {
	if m == nil {
		return types.UnixTime(0)
	}
	return m.CreatedAt
}

// GetUpdatedAt is get updated_at value, if receiver is nil, returns the specified value.
func (m *SampleUser) GetUpdatedAt() types.UnixTime {
	if m == nil {
		return types.UnixTime(0)
	}
	return m.UpdatedAt
}

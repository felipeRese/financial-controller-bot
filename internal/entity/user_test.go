package entity

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSetPlan_WithValidDuration(t *testing.T) {
	// Given a new user
	user := NewUser(uuid.New().String(), "John Doe")

	// When calling SetPlan with a valid duration
	err := user.SetPlan(30 * 24 * time.Hour) // 30 days

	// Then no error should be returned, and the user should be active with a valid expiration date
	assert.Nil(t, err)
	assert.True(t, user.IsActive)

	// Check if expiration date is set correctly
	expectedExpiration := time.Now().Add(30 * 24 * time.Hour)
	assert.WithinDuration(t, expectedExpiration, user.ExpirationDate, time.Minute)
}

func TestSetPlan_WithZeroDuration(t *testing.T) {
	// Given a new user
	user := NewUser(uuid.New().String(), "John Doe")

	// When calling SetPlan with zero duration
	err := user.SetPlan(0)

	// Then an error should be returned and the user should remain inactive
	assert.Error(t, err)
	assert.Equal(t, "invalid duration: must be greater than zero", err.Error())
	assert.False(t, user.IsActive)
	assert.True(t, user.ExpirationDate.IsZero()) // Expiration date should be unset
}

func TestSetPlan_WithNegativeDuration(t *testing.T) {
	// Given a new user
	user := NewUser(uuid.New().String(), "John Doe")

	// When calling SetPlan with a negative duration
	err := user.SetPlan(-5 * time.Hour)

	// Then an error should be returned and the user should remain inactive
	assert.Error(t, err)
	assert.Equal(t, "invalid duration: must be greater than zero", err.Error())
	assert.False(t, user.IsActive)
	assert.True(t, user.ExpirationDate.IsZero()) // Expiration date should be unset
}

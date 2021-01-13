package main

import (
	"encoding/json"
	"net/http"

	"github.com/gustvision/backend-interview/pkg/user"
	"github.com/gustvision/backend-interview/pkg/user/dto"
	"github.com/rs/zerolog/log"
	"testing"
	"https://github.com/stretchr/testify/assert"
)

func TestGetUserFetchUserDB(t *testing.T) {
	u, err := user.Fetch(ctx, user.Filter{ID: "testuid"})

	assert.Nil(t, err)
	assert.EqualValues(t, "testuid", u.ID)
	assert.EqualValues(t, "testname", u.Name)
}

func TestGetUserFetchUserDBError(t *testing.T) {
	u, err := user.Fetch(ctx, user.Filter{ID: "testinvaliduid"})

	assert.NotNil(t, err)
	assert.Nil(t, u.ID)
	assert.EqualValues(t, u.Name)
	assert.Equal
}

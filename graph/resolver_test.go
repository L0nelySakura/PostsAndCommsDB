package graph

import (
	"context"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
)

func TestCreatePost(t *testing.T) {

	connectionString := "postgres://postgres:123@localhost:5432/postgres"
	r := &Resolver{
		DB: NewDatabase(connectionString),
	}

	post, err := r.Mutation().CreatePost(
		context.Background(),
		"Test Title",
		"Test Content",
		"Test Author",
		nil,
	)

	assert.NoError(t, err)
	assert.NotEmpty(t, post.ID)
	assert.Equal(t, "Test Title", post.Title)
	assert.True(t, post.CommentsEnabled)

	comment1, err := r.Mutation().CreateComment(
		context.Background(),
		"Test Comment1",
		"Test Author",
		post.ID,
		"",
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, comment1.ID)
	assert.Equal(t, "Test Comment1", comment1.Content)
	assert.Equal(t, comment1.Parentid, uuid.Nil.String())

	comment2, err := r.Mutation().CreateComment(
		context.Background(),
		"Test Comment2",
		"Test Author",
		post.ID,
		comment1.ID,
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, comment2.ID)
	assert.Equal(t, comment1.ID, comment2.Parentid)
}
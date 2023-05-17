package service

import (
	"testing"

	"github.com/feliciacia/go-gin-framework/go-gin-framework/entity"
	"github.com/stretchr/testify/assert"
)

const (
	TITLE       = "Title"
	DESCRIPTION = "Description"
	URL         = "https://youtu.be/JgW-i2QjgHQ"
)

func getVideo() entity.Video {
	return entity.Video{
		Title:       TITLE,
		Description: DESCRIPTION,
		url:         URL,
	}
}

func TestFindAll(t *testing.T) {
	service := New()
	service.Save(getVideo())
	videos := service.FindAll()
	firstVideo := videos[0]
	assert.NotNil(t, videos)
	assert.Equal(t, TITLE, firstVideo.Title)
	assert.Equal(t, DESCRIPTION, firstVideo.Description)
	assert.Equal(t, URL, firstVideo.URL)
}

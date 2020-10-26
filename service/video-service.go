package service

import (
	"github.com/Aibar01/golang-test/entity"
)

type VideoService interface {
	Save(entity.Video) entity.Video
}

package dto

import (
	"time"
	"wallet-api/internal/model"
)

type Photo struct {
	ID           uint64    `json:"id" valid:"notnull"`
	AlbumID      uint64    `json:"album_id" valid:"notnull"`
	UrlImage     *string   `json:"url_image"`
	CreatedAt    time.Time `json:"created_at" valid:"-" `
	DeletedAt    time.Time `json:"deleted_at" valid:"-"`
	LikeCount    uint64    `valid:"-"`
	CommentCount uint64    `valid:"-"`
} // @name Photo

func (dto *Photo) ParseFromVO(photo *model.Photo) {
	dto.ID = photo.ID
	dto.AlbumID = photo.AlbumID
	dto.UrlImage = photo.UrlImage
	dto.CreatedAt = photo.CreatedAt
	dto.DeletedAt = photo.DeletedAt
	dto.LikeCount = photo.LikeCount
	dto.CommentCount = photo.CommentCount

}

func (dto *Photo) ConvertToArrayDTO(photos []model.Photo) []Photo {
	var photosDTO []Photo

	for _, value := range photos {
		photosDTO = append(photosDTO, Photo{
			ID:           value.ID,
			UrlImage:     value.UrlImage,
			AlbumID:      value.AlbumID,
			CreatedAt:    value.CreatedAt,
			DeletedAt:    value.DeletedAt,
			LikeCount:    value.LikeCount,
			CommentCount: value.CommentCount,
		})
	}
	return photosDTO

}

func (dto *Photo) ConvertToVO() *model.Photo {
	result := &model.Photo{}

	result.ID = dto.ID
	result.AlbumID = dto.AlbumID
	result.UrlImage = dto.UrlImage
	result.CreatedAt = dto.CreatedAt
	result.DeletedAt = dto.DeletedAt
	result.LikeCount = dto.LikeCount
	result.CommentCount = dto.CommentCount

	return result
}

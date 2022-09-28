package model

import "time"

type Photo struct {
	ID           uint64    `json:"id" gorm:"primaryKey;column:id;autoIncrement" valid:"notnull"`
	UrlImage     *string   `valid:"-"`
	AlbumID      uint64    `valid:"notnull"`
	Album        Album     `gorm:"foreignKey:AlbumID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" valid:"notnull"`
	CreatedAt    time.Time `valid:"-" gorm:"autoCreateTime"`
	DeletedAt    time.Time `valid:"-"`
	LikeCount    uint64    `valid:"-"`
	CommentCount uint64    `valid:"-"`
}

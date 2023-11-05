package model

import (
	"context"

	"gorm.io/gorm"
)

type BirdModel interface {
	Create(
		birdName string,
		birdType int64,
		description string,
		picUrl string ) error

}

type Bird struct {
	Id          int64  `gorm:"Id" json:"birdId"`
	BirdName    string `gorm:"birdName" json:"birdName"`
	BirdType    int64  `gorm:"birdType" json:"birdType"`
	Description string `gorm:"description" json:"description"`
	PicUrl      string `gorm:"picUrl" json:"picUrl"`
}

func (t *Bird) TableName() string {
	return "bird"
}

type defaultBirdModel struct {
	db  gorm.DB
	ctx context.Context
}

func NewBirdModel(db gorm.DB, ctx context.Context) BirdModel {
	return &defaultBirdModel{
		db:  db,
		ctx: ctx,
	}
}

func (m *defaultBirdModel) Create(
	birdName string,
	birdType int64,
	description string,
	picUrl string ) error {
	b := &Bird{
		BirdName: birdName,
		BirdType: birdType,
		Description: description,
		PicUrl: picUrl,
	}
	return m.db.Create(b).Error
}

func (m *defaultBirdModel) Update(
	birdId int64,
	birdName string,
	birdType int64,
	description string,
	picUrl string ) error {
	updates := map[string]any{
		"birdName": birdName,
		"birdType": birdType,
		"description": description,
		"picUrl": picUrl,
	}
	return m.db.Where("id = ?", birdId).Updates(updates).Error
}

func (m *defaultBirdModel) List(
	page int,
	pageSize int) ([]*Bird, error) {
	var birds []*Bird
	err := m.db.Limit(pageSize).Offset((page-1)*pageSize).Find(&birds).Error
	return birds, err
}

func (m *defaultBirdModel) Get(
	birdId int64,
	) (*Bird, error) {
	var b Bird
	err := m.db.Where("id = ?", birdId).Find(&b).Error
	return &b, err
}

func (m *defaultBirdModel) Delete(
	birdId int64,
	) error {
	err := m.db.Where("id = ?", birdId).Delete(&Bird{}).Error
	return err
}
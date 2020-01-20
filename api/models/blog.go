package models

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
)

// Blog ...
type Blog struct {
	ID         int             `json:"id"`
	UUID       string          `json:"uuid"`
	Title      string          `json:"title"`
	Content    string          `json:"content"`
	Slug       string          `json:"slug"`
	Image      json.RawMessage `json:"image"`
	Author     string          `json:"author"`
	CategoryID int             `json:"category_id"`
	CreatedAt  time.Time       `json:"created_at"`
}

// PagesListArray ...
type BlogListArray struct {
	Items []*Blog `json:"items"`
	Total int     `json:"count"`
}

// Create a new record
func (pg *Blog) Create(db *gorm.DB) error {
	return db.Create(pg).Error
}

// All pages ...
func (pg *Blog) All(db *gorm.DB, limit, offset int) (ss []*Blog, count int, err error) {
	err = db.
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&ss).
		Limit(-1).
		Offset(-1).
		Count(&count).Error
	return
}

// GET ...
func (pg *Blog) GETBYUUID(db *gorm.DB, id string) error {
	return db.First(pg, "uuid = ?", id).Error
}

// GET ...
func (pg *Blog) GET(db *gorm.DB, id int) error {
	return db.First(pg, "id = ?", id).Error
}

// GET ...
func (pg *Blog) GETBYSLUG(db *gorm.DB, slug string) error {
	return db.First(pg,"slug = ?", slug).Error
}

// DELETE ...
func (pg *Blog) DELETE(db *gorm.DB) error {
	return db.Delete(pg).Error
}

// UPDATE ...
func (pg *Blog) UPDATE(db *gorm.DB) error {
	return db.Model(pg).Updates(*pg).Error
}

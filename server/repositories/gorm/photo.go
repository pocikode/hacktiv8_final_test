package gorm

import (
	"MyGram/server/models"
	"MyGram/server/repositories"
	"gorm.io/gorm"
)

type photoRepo struct {
	db *gorm.DB
}

func NewPhotoRepo(db *gorm.DB) repositories.PhotoRepo {
	return &photoRepo{db: db}
}

func (g *photoRepo) CreatePhoto(photo *models.Photo) error {
	return g.db.Create(photo).Error
}

func (g *photoRepo) GetPhoto() (*[]models.Photos, error) {
	var photos []models.Photos
	err := g.db.Model(&models.Photo{}).Select("photos.id, photos.title, photos.caption, photos.photo_url, photos.user_id, photos.created_at, photos.updated_at, users.username, users.email").Joins("inner join users on users.id = photos.user_id").Find(&photos).Error
	if err != nil {
		return nil, err
	}
	return &photos, nil
}

func (g *photoRepo) UpdatePhoto(id int, req *models.Photo) error {
	photo := models.Photo{}
	err := g.db.Model(&photo).Where("id = ?", id).Updates(models.Photo{Title: req.Title, Caption: req.Caption, PhotoUrl: req.PhotoUrl}).Error
	if err != nil {
		return err
	}
	return nil
}

func (g *photoRepo) GetPhotoById(id int) (*models.Photo, error) {
	photo := models.Photo{}
	err := g.db.First(&photo, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &photo, nil
}

func (g *photoRepo) DeletePhoto(id int) error {
	photo := models.Photo{}
	err := g.db.Where("id=?", id).Delete(&photo).Error
	if err != nil {
		return err
	}
	return nil
}

func (g *photoRepo) GetPhotoByUserId(id int) (*models.Photo, error) {
	photo := models.Photo{}
	err := g.db.First(&photo, "user_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &photo, nil
}

func (g *photoRepo) DeletePhotoByUserId(userId int) error {
	photo := models.Photo{}
	err := g.db.Where("userId=?", userId).Delete(&photo).Error
	if err != nil {
		return err
	}
	return nil
}

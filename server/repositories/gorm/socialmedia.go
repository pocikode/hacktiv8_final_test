package gorm

import (
	"MyGram/server/models"
	"MyGram/server/repositories"
	"gorm.io/gorm"
)

type socialMediaRepo struct {
	db *gorm.DB
}

func NewSocialMediaRepo(db *gorm.DB) repositories.SocialMediaRepo {
	return &socialMediaRepo{db: db}
}

func (s *socialMediaRepo) CreateSocialMedia(socialMedia *models.SocialMedia) error {
	return s.db.Create(socialMedia).Error
}

func (s *socialMediaRepo) GetSocialMedia() (*[]models.SocialMedias, error) {
	var socMeds []models.SocialMedias
	err := s.db.Model(&models.SocialMedia{}).Select("social_media.id, social_media.name, social_media.social_media_url, social_media.user_id, social_media.created_at, social_media.updated_at, users.username, users.email").Joins("left join users on users.id = social_media.user_id").Find(&socMeds).Error
	if err != nil {
		return nil, err
	}

	return &socMeds, nil
}

func (s *socialMediaRepo) UpdateSocialMedia(id int, req *models.SocialMedia) error {
	socialMedia := models.SocialMedia{}
	err := s.db.Model(&socialMedia).Where("id = ?", id).Updates(models.SocialMedia{Name: req.Name, SocialMediaUrl: req.SocialMediaUrl}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *socialMediaRepo) DeleteSocMed(id int) error {
	socialMedia := models.SocialMedia{}
	err := s.db.Where("id=?", id).Delete(&socialMedia).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *socialMediaRepo) GetSocMedById(id int) (*models.SocialMedia, error) {
	socialMedia := models.SocialMedia{}
	err := s.db.First(&socialMedia, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &socialMedia, nil
}

func (s *socialMediaRepo) DeleteSocMedByUserId(userId int) error {
	socialMedia := models.SocialMedia{}
	err := s.db.Where("userId=?", userId).Delete(&socialMedia).Error
	if err != nil {
		return err
	}
	return nil
}

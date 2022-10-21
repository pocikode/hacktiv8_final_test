package gorm

import (
	"MyGram/server/models"
	"MyGram/server/repositories"
	"gorm.io/gorm"
)

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) repositories.CommentRepo {
	return &commentRepo{db: db}
}

func (c *commentRepo) CreateComment(comment *models.Comment) error {
	return c.db.Create(comment).Error
}

func (c *commentRepo) GetComment() (*[]models.Comments, error) {
	var comment []models.Comments
	err := c.db.Model(&models.Comment{}).Select("comments.id, comments.message, comments.photo_id, comments.user_id, comments.created_at, comments.updated_at, users.id as user_id_user ,users.username, users.email, photos.id as photo_id_photo, photos.title, photos.caption, photos.photo_url, photos.user_id as user_id_photo").Joins("left join users on users.id = comments.user_id").Joins("left join photos on photos.id = comments.photo_id").Find(&comment).Error
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (c *commentRepo) UpdateComment(id int, req *models.Comment) error {
	comment := models.Comment{}
	err := c.db.Model(&comment).Where("id = ?", id).Updates(models.Comment{Message: req.Message}).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *commentRepo) GetCommentById(id int) (*models.Comment, error) {
	comment := models.Comment{}
	err := c.db.First(&comment, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (c *commentRepo) DeleteComment(id int) error {
	comment := models.Comment{}
	err := c.db.Where("id=?", id).Delete(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *commentRepo) DeleteCommentByUserId(userId int) error {
	comment := models.Comment{}
	err := c.db.Where("userId=?", userId).Delete(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

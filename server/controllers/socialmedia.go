package controllers

import (
	"MyGram/server/helper"
	"MyGram/server/models"
	"MyGram/server/repositories"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type userSocMed struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	ProfileImageUrl string `json:"profile_image_url"`
}

type respSocMed struct {
	Id             int       `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"SocialMediaUrl"`
	UserId         int       `json:"user_id"`
	User           userSocMed
}

type SocialMediaController struct {
	repo      repositories.SocialMediaRepo
	repoPhoto repositories.PhotoRepo
}

func NewSocialMediaController(repo repositories.SocialMediaRepo, repoPhoto repositories.PhotoRepo) *SocialMediaController {
	return &SocialMediaController{repo: repo,
		repoPhoto: repoPhoto}
}

// Create Socmed
// @Summary    Socmed
// @Decription Socmed
// @Tags       socmed
// @Accept     json
// @Produce    json
// @Param Socmed body models.ReqSocMed true "Photo"
// @Router     /socialmedias [post]
// @Security BearerAuth
func (s *SocialMediaController) CreateSocialMedia(ctx *gin.Context) {
	getId := ctx.GetFloat64("id")
	var id int = int(getId)

	var req models.SocialMedia
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	req.UserId = id
	errs := s.repo.CreateSocialMedia(&req)
	if errs != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "ADD_SOCIAL_FAIL",
			Error:   errs.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &helper.Response{
		Status:  http.StatusCreated,
		Message: "CREATE_SOCIAL_SUCCESS",
		Payload: req,
	})
}

// Get Socmed
// @Summary    Socmed
// @Decription Socmed
// @Tags       socmed
// @Accept     json
// @Produce    json
// @Router     /socialmedias [get]
// @Security BearerAuth
func (s *SocialMediaController) GetSocialMedia(ctx *gin.Context) {
	socMeds, err := s.repo.GetSocialMedia()
	if err != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_PHOTO_FAIL",
			Error:   err.Error(),
		})
	}
	var data []respSocMed
	for _, socmed := range *socMeds {
		dataPhoto, _ := s.repoPhoto.GetPhotoByUserId(socmed.UserId)
		fmt.Println(dataPhoto)

		data = append(data, respSocMed{Id: socmed.Id, Name: socmed.Name, SocialMediaUrl: socmed.SocialMediaUrl, UserId: socmed.UserId, CreatedAt: socmed.CreatedAt, UpdatedAt: socmed.UpdatedAt, User: userSocMed{Email: socmed.Email, Username: socmed.Username, ProfileImageUrl: dataPhoto.PhotoUrl}})
	}
	WriteJsonResponse(ctx, &helper.Response{
		Status:  http.StatusOK,
		Message: "GET_PHOTO_SUCCESS",
		Payload: data,
	})
}

// Update Socmed
// @Summary    Socmed
// @Decription Socmed
// @Tags       socmed
// @Accept     json
// @Produce    json
// @Param Socmed body models.ReqSocMed true "Photo"
// @Param      id path int true "Socmed ID"
// @Router     /socialmedias/{id} [put]
// @Security BearerAuth
func (s *SocialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	getId := ctx.Params.ByName("socialMediaId")
	id, errId := strconv.Atoi(getId)
	if errId != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_PHOTO_FAIL",
			Error:   errId.Error(),
		})
		return
	}

	//check Author photo
	dataSocmed, errc := s.repo.GetSocMedById(id)
	if errc != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "Socmed tidak di temukan",
			Error:   errc.Error(),
		})
		return
	}

	getIdCtx := ctx.GetFloat64("id")
	var idUser int = int(getIdCtx)

	if dataSocmed.UserId != idUser {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "NO AUTHOR",
		})
		return
	}

	var req models.SocialMedia
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	errUp := s.repo.UpdateSocialMedia(id, &req)
	if errUp != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_SOCMED_FAIL",
			Error:   errUp.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &helper.Response{
		Status:  http.StatusOK,
		Message: "UPDATE_PHOTO_SUCCESS",
		Payload: req,
	})
}

// Delete Socmed
// @Summary    Socmed
// @Decription Socmed
// @Tags       socmed
// @Accept     json
// @Produce    json
// @Param      id path int true "Socmed ID"
// @Router     /socialmedias/{id} [delete]
// @Security BearerAuth
func (s *SocialMediaController) DeleteSocmed(ctx *gin.Context) {
	getId := ctx.Params.ByName("socialMediaId")
	id, err := strconv.Atoi(getId)
	if err != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_SOSMED_FAIL",
			Error:   err.Error(),
		})
		return
	}

	//check Author photo
	dataSocMed, err := s.repo.GetSocMedById(id)
	if err != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "Socmed tidak di temukan",
			Error:   err.Error(),
		})
		return
	}

	getIdCtx := ctx.GetFloat64("id")
	var idUser int = int(getIdCtx)

	if dataSocMed.UserId != idUser {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "NO AUTHOR",
		})
		return
	}

	err = s.repo.DeleteSocMed(id)
	if err != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_SOCMED_FAIL",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &helper.Response{
		Status:  http.StatusOK,
		Message: "Your Social media has been successfully deleted",
		Payload: id,
	})
}

package dto

import (
	"errors"

	"github.com/someguy609/be-proyek-fsi/entity"
)

const (
	// Failed
	MESSAGE_FAILED_CREATE_LOCATION   = "failed create location"
	MESSAGE_FAILED_GET_LIST_LOCATION = "failed get list location"
	MESSAGE_FAILED_GET_LOCATION      = "failed get location"
	MESSAGE_FAILED_UPDATE_LOCATION   = "failed update location"
	MESSAGE_FAILED_DELETE_LOCATION   = "failed delete location"

	// Success
	MESSAGE_SUCCESS_CREATE_LOCATION   = "success create location"
	MESSAGE_SUCCESS_GET_LIST_LOCATION = "success get list location"
	MESSAGE_SUCCESS_GET_LOCATION      = "success get location"
	MESSAGE_SUCCESS_UPDATE_LOCATION   = "success update location"
	MESSAGE_SUCCESS_DELETE_LOCATION   = "success delete location"
)

var (
	ErrCreateLocation   = errors.New("failed to create location")
	ErrGetLocationById  = errors.New("failed to get location by id")
	ErrUpdateLocation   = errors.New("failed to update location")
	ErrLocationNotFound = errors.New("location not found")
	ErrDeleteLocation   = errors.New("failed to delete location")
)

type (
	LocationCreateRequest struct {
		CameraID uint    `json:"camera_id" form:"camera_id" binding:"required"`
		Name     string  `json:"name" form:"name" binding:"required"`
		X1       float32 `json:"x1" form:"x1" binding:"required"`
		Y1       float32 `json:"y1" form:"y1" binding:"required"`
		X2       float32 `json:"x2" form:"x2" binding:"required"`
		Y2       float32 `json:"y2" form:"y2" binding:"required"`
	}

	LocationResponse struct {
		ID       string  `json:"id"`
		CameraID uint    `json:"camera_id"`
		Name     string  `json:"name"`
		X1       float32 `json:"x1"`
		Y1       float32 `json:"y1"`
		X2       float32 `json:"x2"`
		Y2       float32 `json:"y2"`
	}

	LocationPaginationResponse struct {
		Data []LocationResponse `json:"data"`
		PaginationResponse
	}

	GetAllLocationRepositoryResponse struct {
		Locations []entity.Location `json:"locations"`
		PaginationResponse
	}

	LocationUpdateRequest struct {
		Name string  `json:"name" form:"name" binding:"omitempty"`
		X1   float32 `json:"x1" form:"x1" binding:"omitempty"`
		Y1   float32 `json:"y1" form:"y1" binding:"omitempty"`
		X2   float32 `json:"x2" form:"x2" binding:"omitempty"`
		Y2   float32 `json:"y2" form:"y2" binding:"omitempty"`
	}

	LocationUpdateResponse struct {
		ID   string  `json:"id"`
		Name string  `json:"name"`
		X1   float32 `json:"x1"`
		Y1   float32 `json:"y1"`
		X2   float32 `json:"x2"`
		Y2   float32 `json:"y2"`
	}
)

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/someguy609/be-proyek-fsi/dto"
	"github.com/someguy609/be-proyek-fsi/service"
	"github.com/someguy609/be-proyek-fsi/utils"
)

type (
	LocationController interface {
		Create(ctx *gin.Context)
		GetAllLocation(ctx *gin.Context)
		GetLocationById(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	locationController struct {
		locationService service.LocationService
	}
)

func NewLocationController(us service.LocationService) LocationController {
	return &locationController{
		locationService: us,
	}
}

func (c *locationController) Create(ctx *gin.Context) {
	var location dto.LocationCreateRequest
	if err := ctx.ShouldBind(&location); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.locationService.Create(ctx.Request.Context(), location)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_LOCATION, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_LOCATION, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *locationController) GetAllLocation(ctx *gin.Context) {
	var req dto.PaginationRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.locationService.GetAllLocationWithPagination(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_LOCATION, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_LIST_LOCATION,
		Data:    result.Data,
		Meta:    result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *locationController) GetLocationById(ctx *gin.Context) {
	locationId := ctx.Param("id")

	result, err := c.locationService.GetLocationById(ctx.Request.Context(), locationId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LOCATION, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LOCATION, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *locationController) Update(ctx *gin.Context) {
	var req dto.LocationUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	locationId := ctx.Param("id")

	result, err := c.locationService.Update(ctx.Request.Context(), req, locationId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_LOCATION, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_LOCATION, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *locationController) Delete(ctx *gin.Context) {
	locationId := ctx.Param("id")

	if err := c.locationService.Delete(ctx.Request.Context(), locationId); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_LOCATION, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_LOCATION, nil)
	ctx.JSON(http.StatusOK, res)
}

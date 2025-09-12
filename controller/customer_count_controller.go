package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/someguy609/be-proyek-fsi/dto"
	"github.com/someguy609/be-proyek-fsi/service"
	"github.com/someguy609/be-proyek-fsi/utils"
)

type (
	CustomerCountController interface {
		Create(ctx *gin.Context)
		GetCustomerCountByLocation(ctx *gin.Context)
		Update(ctx *gin.Context)
		// Delete(ctx *gin.Context)
	}

	customerCountController struct {
		customerCountService service.CustomerCountService
	}
)

func NewCustomerCountController(us service.CustomerCountService) CustomerCountController {
	return &customerCountController{
		customerCountService: us,
	}
}

func (c *customerCountController) Create(ctx *gin.Context) {
	locationId := ctx.Param("id")

	var customerCount dto.CustomerCountCreateRequest
	if err := ctx.ShouldBind(&customerCount); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.customerCountService.Create(ctx.Request.Context(), customerCount, locationId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_CUSTOMER_COUNT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_CUSTOMER_COUNT, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *customerCountController) GetCustomerCountByLocation(ctx *gin.Context) {
	locationId := ctx.Param("id")

	var start, end *time.Time

	startString := ctx.Query("start")

	if startString != "" {
		startTime, err := time.Parse(time.RFC3339, startString)

		if err != nil {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_CUSTOMER_COUNT, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		start = &startTime
	}

	endString := ctx.Query("end")

	if endString != "" {
		endTime, err := time.Parse(time.RFC3339, endString)

		if err != nil {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_CUSTOMER_COUNT, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		end = &endTime
	}

	interval := ctx.DefaultQuery("interval", "minute")

	result, err := c.customerCountService.GetCustomerCountByLocation(ctx.Request.Context(), locationId, start, end, interval)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_CUSTOMER_COUNT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_CUSTOMER_COUNT, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *customerCountController) Update(ctx *gin.Context) {
	locationId := ctx.Param("id")

	var req []dto.CustomerCountUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.customerCountService.Update(ctx.Request.Context(), req, locationId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_CUSTOMER_COUNT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_CUSTOMER_COUNT, result)
	ctx.JSON(http.StatusOK, res)
}

// func (c *customerCountController) Delete(ctx *gin.Context) {
// 	locationId := ctx.Param("id")

// 	if err := c.customerCountService.Delete(ctx.Request.Context(), locationId, nil); err != nil {
// 		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_CUSTOMER_COUNT, err.Error(), nil)
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_CUSTOMER_COUNT, nil)
// 	ctx.JSON(http.StatusOK, res)
// }

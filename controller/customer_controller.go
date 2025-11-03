package controller

// import (
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/someguy609/be-proyek-fsi/dto"
// 	"github.com/someguy609/be-proyek-fsi/service"
// 	"github.com/someguy609/be-proyek-fsi/utils"
// )

// type (
// 	CustomerController interface {
// 		Create(ctx *gin.Context)
// 		GetCustomerById(ctx *gin.Context)
// 		GetCustomerByLocation(ctx *gin.Context)
// 		Update(ctx *gin.Context)
// 		Delete(ctx *gin.Context)
// 	}

// 	customerController struct {
// 		customerService service.CustomerService
// 	}
// )

// func NewCustomerController(us service.CustomerService) CustomerController {
// 	return &customerController{
// 		customerService: us,
// 	}
// }

// func (c *customerController) Create(ctx *gin.Context) {
// 	var customer dto.CustomerCreateRequest
// 	if err := ctx.ShouldBind(&customer); err != nil {
// 		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	result, err := c.customerService.Create(ctx.Request.Context(), customer)
// 	if err != nil {
// 		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_CUSTOMER, err.Error(), nil)
// 		ctx.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_CUSTOMER, result)
// 	ctx.JSON(http.StatusOK, res)
// }

// func (c *customerController) GetCustomerById(ctx *gin.Context) {
// 	customerId := ctx.Param("id")

// 	result, err := c.customerService.GetCustomerById(ctx.Request.Context(), customerId)
// 	if err != nil {
// 		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_CUSTOMER, err.Error(), nil)
// 		ctx.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_CUSTOMER, result)
// 	ctx.JSON(http.StatusOK, res)
// }

// func (c *customerController) GetCustomerByLocation(ctx *gin.Context) {
// 	locationId := ctx.Param("id")

// 	var start, end *time.Time

// 	startString := ctx.Query("start")

// 	if startString != "" {
// 		startTime, err := time.Parse(time.RFC3339, startString)

// 		if err != nil {
// 			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_CUSTOMER, err.Error(), nil)
// 			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
// 			return
// 		}

// 		start = &startTime
// 	}

// 	endString := ctx.Query("end")

// 	if endString != "" {
// 		endTime, err := time.Parse(time.RFC3339, endString)

// 		if err != nil {
// 			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_CUSTOMER, err.Error(), nil)
// 			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
// 			return
// 		}

// 		end = &endTime
// 	}

// 	interval := ctx.DefaultQuery("interval", "minute")

// 	result, err := c.customerService.GetCustomerByLocation(ctx.Request.Context(), locationId, start, end, interval)
// 	if err != nil {
// 		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_CUSTOMER, err.Error(), nil)
// 		ctx.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_CUSTOMER, result)
// 	ctx.JSON(http.StatusOK, res)
// }

// func (c *customerController) Update(ctx *gin.Context) {
// 	customerId := ctx.Param("id")

// 	var req []dto.CustomerUpdateRequest
// 	if err := ctx.ShouldBind(&req); err != nil {
// 		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	result, err := c.customerService.Update(ctx.Request.Context(), req, customerId)
// 	if err != nil {
// 		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_CUSTOMER, err.Error(), nil)
// 		ctx.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_CUSTOMER, result)
// 	ctx.JSON(http.StatusOK, res)
// }

// func (c *customerController) Delete(ctx *gin.Context) {
// 	customerId := ctx.Param("id")

// 	if err := c.customerService.Delete(ctx.Request.Context(), customerId); err != nil {
// 		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_CUSTOMER, err.Error(), nil)
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_CUSTOMER, nil)
// 	ctx.JSON(http.StatusOK, res)
// }

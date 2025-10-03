package dto

import (
	"errors"
	"time"

	"github.com/someguy609/be-proyek-fsi/entity"
)

const (
	// Failed
	MESSAGE_FAILED_CREATE_CUSTOMER_COUNT = "failed create customer count"
	MESSAGE_FAILED_GET_CUSTOMER_COUNT    = "failed get customer count"
	MESSAGE_FAILED_UPDATE_CUSTOMER_COUNT = "failed update customer count"
	MESSAGE_FAILED_DELETE_CUSTOMER_COUNT = "failed delete customer count"

	// Success
	MESSAGE_SUCCESS_CREATE_CUSTOMER_COUNT = "success create customer count"
	MESSAGE_SUCCESS_GET_CUSTOMER_COUNT    = "success get customer count"
	MESSAGE_SUCCESS_UPDATE_CUSTOMER_COUNT = "success update customer count"
	MESSAGE_SUCCESS_DELETE_CUSTOMER_COUNT = "success delete customer count"
)

var (
	ErrCreateCustomerCount   = errors.New("failed to create customer count")
	ErrGetCustomerCountById  = errors.New("failed to get customer count by id")
	ErrUpdateCustomerCount   = errors.New("failed to update customer count")
	ErrCustomerCountNotFound = errors.New("customer count not found")
	ErrDeleteCustomerCount   = errors.New("failed to delete customer count")
)

type (
	CustomerCountCreateRequest struct {
		Timestamp time.Time     `json:"timestamp" form:"timestamp" binding:"required"`
		Gender    entity.Gender `json:"gender" form:"gender" binding:"required"`
		Count     int64         `json:"count" form:"count" binding:"required"`
	}

	CustomerCountResponse struct {
		Timestamp  time.Time     `json:"timestamp"`
		LocationID string        `json:"location_id"`
		Gender     entity.Gender `json:"gender"`
		Count      int64         `json:"count"`
	}

	CustomerCountGetResponse struct {
		Data []entity.CustomerCount `json:"data"`
	}

	CustomerCountUpdateRequest struct {
		Timestamp time.Time     `json:"timestamp" form:"timestamp" binding:"required"`
		Gender    entity.Gender `json:"gender" form:"gender" binding:"omitempty"`
		Count     int64         `json:"count" form:"count" binding:"omitempty"`
	}

	CustomerCountUpdateResponse struct {
		Data []entity.CustomerCount `json:"data"`
	}
)

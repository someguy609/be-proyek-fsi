package dto

// import (
// 	"errors"
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/someguy609/be-proyek-fsi/entity"
// )

// const (
// 	// Failed
// 	MESSAGE_FAILED_CREATE_CUSTOMER = "failed create customer count"
// 	MESSAGE_FAILED_GET_CUSTOMER    = "failed get customer count"
// 	MESSAGE_FAILED_UPDATE_CUSTOMER = "failed update customer count"
// 	MESSAGE_FAILED_DELETE_CUSTOMER = "failed delete customer count"

// 	// Success
// 	MESSAGE_SUCCESS_CREATE_CUSTOMER = "success create customer count"
// 	MESSAGE_SUCCESS_GET_CUSTOMER    = "success get customer count"
// 	MESSAGE_SUCCESS_UPDATE_CUSTOMER = "success update customer count"
// 	MESSAGE_SUCCESS_DELETE_CUSTOMER = "success delete customer count"
// )

// var (
// 	ErrCreateCustomer   = errors.New("failed to create customer count")
// 	ErrGetCustomerById  = errors.New("failed to get customer count by id")
// 	ErrUpdateCustomer   = errors.New("failed to update customer count")
// 	ErrCustomerNotFound = errors.New("customer count not found")
// 	ErrDeleteCustomer   = errors.New("failed to delete customer count")
// )

// type (
// 	CustomerCreateRequest struct {
// 		Gender     entity.Gender `json:"gender" form:"gender" binding:"required"`
// 		EntryTime  time.Time     `json:"entry_time" form:"entry_time" binding:"required"`
// 		ExitTime   time.Time     `json:"exit_time" form:"exit_time" binding:"omitempty"`
// 		LocationID string        `json:"location_id" form:"location_id" binding:"required"`
// 	}

// 	CustomerResponse struct {
// 		ID        uuid.UUID     `json:"id"`
// 		Gender    entity.Gender `json:"gender"`
// 		EntryTime time.Time     `json:"entry_time"`
// 		ExitTime  time.Time     `json:"exit_time"`
// 	}

// 	CustomerPaginationResponse struct {
// 		Data []CustomerResponse `json:"data"`
// 		PaginationResponse
// 	}

// 	GetAllCustomerRepositoryResponse struct {
// 		Customers []entity.Customer `json:"customers"`
// 		PaginationResponse
// 	}

// 	CustomerUpdateRequest struct {
// 		ID        uuid.UUID     `json:"id" form:"id" binding:"required"`
// 		Gender    entity.Gender `json:"gender" form:"gender" binding:"omitempty"`
// 		EntryTime time.Time     `json:"entry_time" form:"entry_time" binding:"omitempty"`
// 		ExitTime  time.Time     `json:"exit_time" form:"exit_time" binding:"omitempty"`
// 	}

// 	CustomerUpdateResponse struct {
// 		ID   string  `json:"id"`
// 		Name string  `json:"name"`
// 		X1   float32 `json:"x1"`
// 		Y1   float32 `json:"y1"`
// 		X2   float32 `json:"x2"`
// 		Y2   float32 `json:"y2"`
// 	}
// )

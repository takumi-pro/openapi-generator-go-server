/*
 * todos
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service DefaultApiServicer
	errorHandler ErrorHandler
}

// DefaultApiOption for how the controller is set up.
type DefaultApiOption func(*DefaultApiController)

// WithDefaultApiErrorHandler inject ErrorHandler into controller
func WithDefaultApiErrorHandler(h ErrorHandler) DefaultApiOption {
	return func(c *DefaultApiController) {
		c.errorHandler = h
	}
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer, opts ...DefaultApiOption) Router {
	controller := &DefaultApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{ 
		{
			"GetTask",
			strings.ToUpper("Get"),
			"/task",
			c.GetTask,
		},
		{
			"GetTaskTaskId",
			strings.ToUpper("Get"),
			"/task/{taskId}",
			c.GetTaskTaskId,
		},
		{
			"PostTask",
			strings.ToUpper("Post"),
			"/task",
			c.PostTask,
		},
		{
			"PutTaskTaskId",
			strings.ToUpper("Put"),
			"/task/{taskId}",
			c.PutTaskTaskId,
		},
	}
}

// GetTask - Fetch All Task
func (c *DefaultApiController) GetTask(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetTask(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetTaskTaskId - Fetch Task
func (c *DefaultApiController) GetTaskTaskId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskIdParam := params["taskId"]
	result, err := c.service.GetTaskTaskId(r.Context(), taskIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// PostTask - Create New Task
func (c *DefaultApiController) PostTask(w http.ResponseWriter, r *http.Request) {
	taskParam := Task{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&taskParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertTaskRequired(taskParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PostTask(r.Context(), taskParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// PutTaskTaskId - Update Task
func (c *DefaultApiController) PutTaskTaskId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskIdParam := params["taskId"]
	taskParam := Task{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&taskParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertTaskRequired(taskParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PutTaskTaskId(r.Context(), taskIdParam, taskParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

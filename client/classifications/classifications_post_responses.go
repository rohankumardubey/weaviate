//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package classifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/entities/models"
)

// ClassificationsPostReader is a Reader for the ClassificationsPost structure.
type ClassificationsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClassificationsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewClassificationsPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewClassificationsPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewClassificationsPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewClassificationsPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewClassificationsPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewClassificationsPostCreated creates a ClassificationsPostCreated with default headers values
func NewClassificationsPostCreated() *ClassificationsPostCreated {
	return &ClassificationsPostCreated{}
}

/*
ClassificationsPostCreated describes a response with status code 201, with default header values.

Successfully started classification.
*/
type ClassificationsPostCreated struct {
	Payload *models.Classification
}

// IsSuccess returns true when this classifications post created response has a 2xx status code
func (o *ClassificationsPostCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this classifications post created response has a 3xx status code
func (o *ClassificationsPostCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this classifications post created response has a 4xx status code
func (o *ClassificationsPostCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this classifications post created response has a 5xx status code
func (o *ClassificationsPostCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this classifications post created response a status code equal to that given
func (o *ClassificationsPostCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the classifications post created response
func (o *ClassificationsPostCreated) Code() int {
	return 201
}

func (o *ClassificationsPostCreated) Error() string {
	return fmt.Sprintf("[POST /v1/classifications/][%d] classificationsPostCreated  %+v", 201, o.Payload)
}

func (o *ClassificationsPostCreated) String() string {
	return fmt.Sprintf("[POST /v1/classifications/][%d] classificationsPostCreated  %+v", 201, o.Payload)
}

func (o *ClassificationsPostCreated) GetPayload() *models.Classification {
	return o.Payload
}

func (o *ClassificationsPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Classification)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClassificationsPostBadRequest creates a ClassificationsPostBadRequest with default headers values
func NewClassificationsPostBadRequest() *ClassificationsPostBadRequest {
	return &ClassificationsPostBadRequest{}
}

/*
ClassificationsPostBadRequest describes a response with status code 400, with default header values.

Incorrect request
*/
type ClassificationsPostBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this classifications post bad request response has a 2xx status code
func (o *ClassificationsPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this classifications post bad request response has a 3xx status code
func (o *ClassificationsPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this classifications post bad request response has a 4xx status code
func (o *ClassificationsPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this classifications post bad request response has a 5xx status code
func (o *ClassificationsPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this classifications post bad request response a status code equal to that given
func (o *ClassificationsPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the classifications post bad request response
func (o *ClassificationsPostBadRequest) Code() int {
	return 400
}

func (o *ClassificationsPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/classifications/][%d] classificationsPostBadRequest  %+v", 400, o.Payload)
}

func (o *ClassificationsPostBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/classifications/][%d] classificationsPostBadRequest  %+v", 400, o.Payload)
}

func (o *ClassificationsPostBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClassificationsPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClassificationsPostUnauthorized creates a ClassificationsPostUnauthorized with default headers values
func NewClassificationsPostUnauthorized() *ClassificationsPostUnauthorized {
	return &ClassificationsPostUnauthorized{}
}

/*
ClassificationsPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type ClassificationsPostUnauthorized struct {
}

// IsSuccess returns true when this classifications post unauthorized response has a 2xx status code
func (o *ClassificationsPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this classifications post unauthorized response has a 3xx status code
func (o *ClassificationsPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this classifications post unauthorized response has a 4xx status code
func (o *ClassificationsPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this classifications post unauthorized response has a 5xx status code
func (o *ClassificationsPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this classifications post unauthorized response a status code equal to that given
func (o *ClassificationsPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the classifications post unauthorized response
func (o *ClassificationsPostUnauthorized) Code() int {
	return 401
}

func (o *ClassificationsPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/classifications/][%d] classificationsPostUnauthorized ", 401)
}

func (o *ClassificationsPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/classifications/][%d] classificationsPostUnauthorized ", 401)
}

func (o *ClassificationsPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClassificationsPostForbidden creates a ClassificationsPostForbidden with default headers values
func NewClassificationsPostForbidden() *ClassificationsPostForbidden {
	return &ClassificationsPostForbidden{}
}

/*
ClassificationsPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ClassificationsPostForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this classifications post forbidden response has a 2xx status code
func (o *ClassificationsPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this classifications post forbidden response has a 3xx status code
func (o *ClassificationsPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this classifications post forbidden response has a 4xx status code
func (o *ClassificationsPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this classifications post forbidden response has a 5xx status code
func (o *ClassificationsPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this classifications post forbidden response a status code equal to that given
func (o *ClassificationsPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the classifications post forbidden response
func (o *ClassificationsPostForbidden) Code() int {
	return 403
}

func (o *ClassificationsPostForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/classifications/][%d] classificationsPostForbidden  %+v", 403, o.Payload)
}

func (o *ClassificationsPostForbidden) String() string {
	return fmt.Sprintf("[POST /v1/classifications/][%d] classificationsPostForbidden  %+v", 403, o.Payload)
}

func (o *ClassificationsPostForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClassificationsPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClassificationsPostInternalServerError creates a ClassificationsPostInternalServerError with default headers values
func NewClassificationsPostInternalServerError() *ClassificationsPostInternalServerError {
	return &ClassificationsPostInternalServerError{}
}

/*
ClassificationsPostInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ClassificationsPostInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this classifications post internal server error response has a 2xx status code
func (o *ClassificationsPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this classifications post internal server error response has a 3xx status code
func (o *ClassificationsPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this classifications post internal server error response has a 4xx status code
func (o *ClassificationsPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this classifications post internal server error response has a 5xx status code
func (o *ClassificationsPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this classifications post internal server error response a status code equal to that given
func (o *ClassificationsPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the classifications post internal server error response
func (o *ClassificationsPostInternalServerError) Code() int {
	return 500
}

func (o *ClassificationsPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/classifications/][%d] classificationsPostInternalServerError  %+v", 500, o.Payload)
}

func (o *ClassificationsPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/classifications/][%d] classificationsPostInternalServerError  %+v", 500, o.Payload)
}

func (o *ClassificationsPostInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClassificationsPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/entities/models"
)

// ObjectsListReader is a Reader for the ObjectsList structure.
type ObjectsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ObjectsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewObjectsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewObjectsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewObjectsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewObjectsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewObjectsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewObjectsListUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewObjectsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewObjectsListOK creates a ObjectsListOK with default headers values
func NewObjectsListOK() *ObjectsListOK {
	return &ObjectsListOK{}
}

/*
ObjectsListOK describes a response with status code 200, with default header values.

Successful response.
*/
type ObjectsListOK struct {
	Payload *models.ObjectsListResponse
}

// IsSuccess returns true when this objects list o k response has a 2xx status code
func (o *ObjectsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this objects list o k response has a 3xx status code
func (o *ObjectsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects list o k response has a 4xx status code
func (o *ObjectsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this objects list o k response has a 5xx status code
func (o *ObjectsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this objects list o k response a status code equal to that given
func (o *ObjectsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the objects list o k response
func (o *ObjectsListOK) Code() int {
	return 200
}

func (o *ObjectsListOK) Error() string {
	return fmt.Sprintf("[GET /v1/objects][%d] objectsListOK  %+v", 200, o.Payload)
}

func (o *ObjectsListOK) String() string {
	return fmt.Sprintf("[GET /v1/objects][%d] objectsListOK  %+v", 200, o.Payload)
}

func (o *ObjectsListOK) GetPayload() *models.ObjectsListResponse {
	return o.Payload
}

func (o *ObjectsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsListBadRequest creates a ObjectsListBadRequest with default headers values
func NewObjectsListBadRequest() *ObjectsListBadRequest {
	return &ObjectsListBadRequest{}
}

/*
ObjectsListBadRequest describes a response with status code 400, with default header values.

Malformed request.
*/
type ObjectsListBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this objects list bad request response has a 2xx status code
func (o *ObjectsListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects list bad request response has a 3xx status code
func (o *ObjectsListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects list bad request response has a 4xx status code
func (o *ObjectsListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this objects list bad request response has a 5xx status code
func (o *ObjectsListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this objects list bad request response a status code equal to that given
func (o *ObjectsListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the objects list bad request response
func (o *ObjectsListBadRequest) Code() int {
	return 400
}

func (o *ObjectsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/objects][%d] objectsListBadRequest  %+v", 400, o.Payload)
}

func (o *ObjectsListBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/objects][%d] objectsListBadRequest  %+v", 400, o.Payload)
}

func (o *ObjectsListBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsListUnauthorized creates a ObjectsListUnauthorized with default headers values
func NewObjectsListUnauthorized() *ObjectsListUnauthorized {
	return &ObjectsListUnauthorized{}
}

/*
ObjectsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type ObjectsListUnauthorized struct {
}

// IsSuccess returns true when this objects list unauthorized response has a 2xx status code
func (o *ObjectsListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects list unauthorized response has a 3xx status code
func (o *ObjectsListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects list unauthorized response has a 4xx status code
func (o *ObjectsListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this objects list unauthorized response has a 5xx status code
func (o *ObjectsListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this objects list unauthorized response a status code equal to that given
func (o *ObjectsListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the objects list unauthorized response
func (o *ObjectsListUnauthorized) Code() int {
	return 401
}

func (o *ObjectsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/objects][%d] objectsListUnauthorized ", 401)
}

func (o *ObjectsListUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/objects][%d] objectsListUnauthorized ", 401)
}

func (o *ObjectsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsListForbidden creates a ObjectsListForbidden with default headers values
func NewObjectsListForbidden() *ObjectsListForbidden {
	return &ObjectsListForbidden{}
}

/*
ObjectsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ObjectsListForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this objects list forbidden response has a 2xx status code
func (o *ObjectsListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects list forbidden response has a 3xx status code
func (o *ObjectsListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects list forbidden response has a 4xx status code
func (o *ObjectsListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this objects list forbidden response has a 5xx status code
func (o *ObjectsListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this objects list forbidden response a status code equal to that given
func (o *ObjectsListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the objects list forbidden response
func (o *ObjectsListForbidden) Code() int {
	return 403
}

func (o *ObjectsListForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/objects][%d] objectsListForbidden  %+v", 403, o.Payload)
}

func (o *ObjectsListForbidden) String() string {
	return fmt.Sprintf("[GET /v1/objects][%d] objectsListForbidden  %+v", 403, o.Payload)
}

func (o *ObjectsListForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsListNotFound creates a ObjectsListNotFound with default headers values
func NewObjectsListNotFound() *ObjectsListNotFound {
	return &ObjectsListNotFound{}
}

/*
ObjectsListNotFound describes a response with status code 404, with default header values.

Successful query result but no resource was found.
*/
type ObjectsListNotFound struct {
}

// IsSuccess returns true when this objects list not found response has a 2xx status code
func (o *ObjectsListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects list not found response has a 3xx status code
func (o *ObjectsListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects list not found response has a 4xx status code
func (o *ObjectsListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this objects list not found response has a 5xx status code
func (o *ObjectsListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this objects list not found response a status code equal to that given
func (o *ObjectsListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the objects list not found response
func (o *ObjectsListNotFound) Code() int {
	return 404
}

func (o *ObjectsListNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/objects][%d] objectsListNotFound ", 404)
}

func (o *ObjectsListNotFound) String() string {
	return fmt.Sprintf("[GET /v1/objects][%d] objectsListNotFound ", 404)
}

func (o *ObjectsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsListUnprocessableEntity creates a ObjectsListUnprocessableEntity with default headers values
func NewObjectsListUnprocessableEntity() *ObjectsListUnprocessableEntity {
	return &ObjectsListUnprocessableEntity{}
}

/*
ObjectsListUnprocessableEntity describes a response with status code 422, with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?
*/
type ObjectsListUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this objects list unprocessable entity response has a 2xx status code
func (o *ObjectsListUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects list unprocessable entity response has a 3xx status code
func (o *ObjectsListUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects list unprocessable entity response has a 4xx status code
func (o *ObjectsListUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this objects list unprocessable entity response has a 5xx status code
func (o *ObjectsListUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this objects list unprocessable entity response a status code equal to that given
func (o *ObjectsListUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the objects list unprocessable entity response
func (o *ObjectsListUnprocessableEntity) Code() int {
	return 422
}

func (o *ObjectsListUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /v1/objects][%d] objectsListUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ObjectsListUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /v1/objects][%d] objectsListUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ObjectsListUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsListUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsListInternalServerError creates a ObjectsListInternalServerError with default headers values
func NewObjectsListInternalServerError() *ObjectsListInternalServerError {
	return &ObjectsListInternalServerError{}
}

/*
ObjectsListInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ObjectsListInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this objects list internal server error response has a 2xx status code
func (o *ObjectsListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects list internal server error response has a 3xx status code
func (o *ObjectsListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects list internal server error response has a 4xx status code
func (o *ObjectsListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this objects list internal server error response has a 5xx status code
func (o *ObjectsListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this objects list internal server error response a status code equal to that given
func (o *ObjectsListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the objects list internal server error response
func (o *ObjectsListInternalServerError) Code() int {
	return 500
}

func (o *ObjectsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/objects][%d] objectsListInternalServerError  %+v", 500, o.Payload)
}

func (o *ObjectsListInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/objects][%d] objectsListInternalServerError  %+v", 500, o.Payload)
}

func (o *ObjectsListInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

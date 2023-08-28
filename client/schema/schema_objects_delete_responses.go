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

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/entities/models"
)

// SchemaObjectsDeleteReader is a Reader for the SchemaObjectsDelete structure.
type SchemaObjectsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchemaObjectsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSchemaObjectsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSchemaObjectsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSchemaObjectsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSchemaObjectsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSchemaObjectsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSchemaObjectsDeleteOK creates a SchemaObjectsDeleteOK with default headers values
func NewSchemaObjectsDeleteOK() *SchemaObjectsDeleteOK {
	return &SchemaObjectsDeleteOK{}
}

/*
SchemaObjectsDeleteOK describes a response with status code 200, with default header values.

Removed the Object class from the schema.
*/
type SchemaObjectsDeleteOK struct {
}

// IsSuccess returns true when this schema objects delete o k response has a 2xx status code
func (o *SchemaObjectsDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this schema objects delete o k response has a 3xx status code
func (o *SchemaObjectsDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schema objects delete o k response has a 4xx status code
func (o *SchemaObjectsDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this schema objects delete o k response has a 5xx status code
func (o *SchemaObjectsDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this schema objects delete o k response a status code equal to that given
func (o *SchemaObjectsDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the schema objects delete o k response
func (o *SchemaObjectsDeleteOK) Code() int {
	return 200
}

func (o *SchemaObjectsDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/schema/{className}][%d] schemaObjectsDeleteOK ", 200)
}

func (o *SchemaObjectsDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /v1/schema/{className}][%d] schemaObjectsDeleteOK ", 200)
}

func (o *SchemaObjectsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSchemaObjectsDeleteBadRequest creates a SchemaObjectsDeleteBadRequest with default headers values
func NewSchemaObjectsDeleteBadRequest() *SchemaObjectsDeleteBadRequest {
	return &SchemaObjectsDeleteBadRequest{}
}

/*
SchemaObjectsDeleteBadRequest describes a response with status code 400, with default header values.

Could not delete the Object class.
*/
type SchemaObjectsDeleteBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this schema objects delete bad request response has a 2xx status code
func (o *SchemaObjectsDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schema objects delete bad request response has a 3xx status code
func (o *SchemaObjectsDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schema objects delete bad request response has a 4xx status code
func (o *SchemaObjectsDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this schema objects delete bad request response has a 5xx status code
func (o *SchemaObjectsDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this schema objects delete bad request response a status code equal to that given
func (o *SchemaObjectsDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the schema objects delete bad request response
func (o *SchemaObjectsDeleteBadRequest) Code() int {
	return 400
}

func (o *SchemaObjectsDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/schema/{className}][%d] schemaObjectsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *SchemaObjectsDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /v1/schema/{className}][%d] schemaObjectsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *SchemaObjectsDeleteBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaObjectsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaObjectsDeleteUnauthorized creates a SchemaObjectsDeleteUnauthorized with default headers values
func NewSchemaObjectsDeleteUnauthorized() *SchemaObjectsDeleteUnauthorized {
	return &SchemaObjectsDeleteUnauthorized{}
}

/*
SchemaObjectsDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type SchemaObjectsDeleteUnauthorized struct {
}

// IsSuccess returns true when this schema objects delete unauthorized response has a 2xx status code
func (o *SchemaObjectsDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schema objects delete unauthorized response has a 3xx status code
func (o *SchemaObjectsDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schema objects delete unauthorized response has a 4xx status code
func (o *SchemaObjectsDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this schema objects delete unauthorized response has a 5xx status code
func (o *SchemaObjectsDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this schema objects delete unauthorized response a status code equal to that given
func (o *SchemaObjectsDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the schema objects delete unauthorized response
func (o *SchemaObjectsDeleteUnauthorized) Code() int {
	return 401
}

func (o *SchemaObjectsDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/schema/{className}][%d] schemaObjectsDeleteUnauthorized ", 401)
}

func (o *SchemaObjectsDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/schema/{className}][%d] schemaObjectsDeleteUnauthorized ", 401)
}

func (o *SchemaObjectsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSchemaObjectsDeleteForbidden creates a SchemaObjectsDeleteForbidden with default headers values
func NewSchemaObjectsDeleteForbidden() *SchemaObjectsDeleteForbidden {
	return &SchemaObjectsDeleteForbidden{}
}

/*
SchemaObjectsDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SchemaObjectsDeleteForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this schema objects delete forbidden response has a 2xx status code
func (o *SchemaObjectsDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schema objects delete forbidden response has a 3xx status code
func (o *SchemaObjectsDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schema objects delete forbidden response has a 4xx status code
func (o *SchemaObjectsDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this schema objects delete forbidden response has a 5xx status code
func (o *SchemaObjectsDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this schema objects delete forbidden response a status code equal to that given
func (o *SchemaObjectsDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the schema objects delete forbidden response
func (o *SchemaObjectsDeleteForbidden) Code() int {
	return 403
}

func (o *SchemaObjectsDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/schema/{className}][%d] schemaObjectsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *SchemaObjectsDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/schema/{className}][%d] schemaObjectsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *SchemaObjectsDeleteForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaObjectsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaObjectsDeleteInternalServerError creates a SchemaObjectsDeleteInternalServerError with default headers values
func NewSchemaObjectsDeleteInternalServerError() *SchemaObjectsDeleteInternalServerError {
	return &SchemaObjectsDeleteInternalServerError{}
}

/*
SchemaObjectsDeleteInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type SchemaObjectsDeleteInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this schema objects delete internal server error response has a 2xx status code
func (o *SchemaObjectsDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schema objects delete internal server error response has a 3xx status code
func (o *SchemaObjectsDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schema objects delete internal server error response has a 4xx status code
func (o *SchemaObjectsDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this schema objects delete internal server error response has a 5xx status code
func (o *SchemaObjectsDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this schema objects delete internal server error response a status code equal to that given
func (o *SchemaObjectsDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the schema objects delete internal server error response
func (o *SchemaObjectsDeleteInternalServerError) Code() int {
	return 500
}

func (o *SchemaObjectsDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/schema/{className}][%d] schemaObjectsDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *SchemaObjectsDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/schema/{className}][%d] schemaObjectsDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *SchemaObjectsDeleteInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaObjectsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

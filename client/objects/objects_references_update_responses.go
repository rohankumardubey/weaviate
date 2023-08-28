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

// ObjectsReferencesUpdateReader is a Reader for the ObjectsReferencesUpdate structure.
type ObjectsReferencesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ObjectsReferencesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewObjectsReferencesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewObjectsReferencesUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewObjectsReferencesUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewObjectsReferencesUpdateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewObjectsReferencesUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewObjectsReferencesUpdateOK creates a ObjectsReferencesUpdateOK with default headers values
func NewObjectsReferencesUpdateOK() *ObjectsReferencesUpdateOK {
	return &ObjectsReferencesUpdateOK{}
}

/*
ObjectsReferencesUpdateOK describes a response with status code 200, with default header values.

Successfully replaced all the references.
*/
type ObjectsReferencesUpdateOK struct {
}

// IsSuccess returns true when this objects references update o k response has a 2xx status code
func (o *ObjectsReferencesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this objects references update o k response has a 3xx status code
func (o *ObjectsReferencesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects references update o k response has a 4xx status code
func (o *ObjectsReferencesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this objects references update o k response has a 5xx status code
func (o *ObjectsReferencesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this objects references update o k response a status code equal to that given
func (o *ObjectsReferencesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the objects references update o k response
func (o *ObjectsReferencesUpdateOK) Code() int {
	return 200
}

func (o *ObjectsReferencesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1/objects/{id}/references/{propertyName}][%d] objectsReferencesUpdateOK ", 200)
}

func (o *ObjectsReferencesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /v1/objects/{id}/references/{propertyName}][%d] objectsReferencesUpdateOK ", 200)
}

func (o *ObjectsReferencesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsReferencesUpdateUnauthorized creates a ObjectsReferencesUpdateUnauthorized with default headers values
func NewObjectsReferencesUpdateUnauthorized() *ObjectsReferencesUpdateUnauthorized {
	return &ObjectsReferencesUpdateUnauthorized{}
}

/*
ObjectsReferencesUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type ObjectsReferencesUpdateUnauthorized struct {
}

// IsSuccess returns true when this objects references update unauthorized response has a 2xx status code
func (o *ObjectsReferencesUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects references update unauthorized response has a 3xx status code
func (o *ObjectsReferencesUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects references update unauthorized response has a 4xx status code
func (o *ObjectsReferencesUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this objects references update unauthorized response has a 5xx status code
func (o *ObjectsReferencesUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this objects references update unauthorized response a status code equal to that given
func (o *ObjectsReferencesUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the objects references update unauthorized response
func (o *ObjectsReferencesUpdateUnauthorized) Code() int {
	return 401
}

func (o *ObjectsReferencesUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/objects/{id}/references/{propertyName}][%d] objectsReferencesUpdateUnauthorized ", 401)
}

func (o *ObjectsReferencesUpdateUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/objects/{id}/references/{propertyName}][%d] objectsReferencesUpdateUnauthorized ", 401)
}

func (o *ObjectsReferencesUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsReferencesUpdateForbidden creates a ObjectsReferencesUpdateForbidden with default headers values
func NewObjectsReferencesUpdateForbidden() *ObjectsReferencesUpdateForbidden {
	return &ObjectsReferencesUpdateForbidden{}
}

/*
ObjectsReferencesUpdateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ObjectsReferencesUpdateForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this objects references update forbidden response has a 2xx status code
func (o *ObjectsReferencesUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects references update forbidden response has a 3xx status code
func (o *ObjectsReferencesUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects references update forbidden response has a 4xx status code
func (o *ObjectsReferencesUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this objects references update forbidden response has a 5xx status code
func (o *ObjectsReferencesUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this objects references update forbidden response a status code equal to that given
func (o *ObjectsReferencesUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the objects references update forbidden response
func (o *ObjectsReferencesUpdateForbidden) Code() int {
	return 403
}

func (o *ObjectsReferencesUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/objects/{id}/references/{propertyName}][%d] objectsReferencesUpdateForbidden  %+v", 403, o.Payload)
}

func (o *ObjectsReferencesUpdateForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/objects/{id}/references/{propertyName}][%d] objectsReferencesUpdateForbidden  %+v", 403, o.Payload)
}

func (o *ObjectsReferencesUpdateForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsReferencesUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsReferencesUpdateUnprocessableEntity creates a ObjectsReferencesUpdateUnprocessableEntity with default headers values
func NewObjectsReferencesUpdateUnprocessableEntity() *ObjectsReferencesUpdateUnprocessableEntity {
	return &ObjectsReferencesUpdateUnprocessableEntity{}
}

/*
ObjectsReferencesUpdateUnprocessableEntity describes a response with status code 422, with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the property exists or that it is a class?
*/
type ObjectsReferencesUpdateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this objects references update unprocessable entity response has a 2xx status code
func (o *ObjectsReferencesUpdateUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects references update unprocessable entity response has a 3xx status code
func (o *ObjectsReferencesUpdateUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects references update unprocessable entity response has a 4xx status code
func (o *ObjectsReferencesUpdateUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this objects references update unprocessable entity response has a 5xx status code
func (o *ObjectsReferencesUpdateUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this objects references update unprocessable entity response a status code equal to that given
func (o *ObjectsReferencesUpdateUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the objects references update unprocessable entity response
func (o *ObjectsReferencesUpdateUnprocessableEntity) Code() int {
	return 422
}

func (o *ObjectsReferencesUpdateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /v1/objects/{id}/references/{propertyName}][%d] objectsReferencesUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ObjectsReferencesUpdateUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /v1/objects/{id}/references/{propertyName}][%d] objectsReferencesUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ObjectsReferencesUpdateUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsReferencesUpdateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsReferencesUpdateInternalServerError creates a ObjectsReferencesUpdateInternalServerError with default headers values
func NewObjectsReferencesUpdateInternalServerError() *ObjectsReferencesUpdateInternalServerError {
	return &ObjectsReferencesUpdateInternalServerError{}
}

/*
ObjectsReferencesUpdateInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ObjectsReferencesUpdateInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this objects references update internal server error response has a 2xx status code
func (o *ObjectsReferencesUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects references update internal server error response has a 3xx status code
func (o *ObjectsReferencesUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects references update internal server error response has a 4xx status code
func (o *ObjectsReferencesUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this objects references update internal server error response has a 5xx status code
func (o *ObjectsReferencesUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this objects references update internal server error response a status code equal to that given
func (o *ObjectsReferencesUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the objects references update internal server error response
func (o *ObjectsReferencesUpdateInternalServerError) Code() int {
	return 500
}

func (o *ObjectsReferencesUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/objects/{id}/references/{propertyName}][%d] objectsReferencesUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *ObjectsReferencesUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/objects/{id}/references/{propertyName}][%d] objectsReferencesUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *ObjectsReferencesUpdateInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsReferencesUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

package backups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/entities/models"
)

// BackupsCreateStatusReader is a Reader for the BackupsCreateStatus structure.
type BackupsCreateStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupsCreateStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBackupsCreateStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewBackupsCreateStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBackupsCreateStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBackupsCreateStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewBackupsCreateStatusUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBackupsCreateStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBackupsCreateStatusOK creates a BackupsCreateStatusOK with default headers values
func NewBackupsCreateStatusOK() *BackupsCreateStatusOK {
	return &BackupsCreateStatusOK{}
}

/*
BackupsCreateStatusOK describes a response with status code 200, with default header values.

Backup creation status successfully returned
*/
type BackupsCreateStatusOK struct {
	Payload *models.BackupCreateStatusResponse
}

// IsSuccess returns true when this backups create status o k response has a 2xx status code
func (o *BackupsCreateStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this backups create status o k response has a 3xx status code
func (o *BackupsCreateStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backups create status o k response has a 4xx status code
func (o *BackupsCreateStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this backups create status o k response has a 5xx status code
func (o *BackupsCreateStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this backups create status o k response a status code equal to that given
func (o *BackupsCreateStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the backups create status o k response
func (o *BackupsCreateStatusOK) Code() int {
	return 200
}

func (o *BackupsCreateStatusOK) Error() string {
	return fmt.Sprintf("[GET /v1/backups/{backend}/{id}][%d] backupsCreateStatusOK  %+v", 200, o.Payload)
}

func (o *BackupsCreateStatusOK) String() string {
	return fmt.Sprintf("[GET /v1/backups/{backend}/{id}][%d] backupsCreateStatusOK  %+v", 200, o.Payload)
}

func (o *BackupsCreateStatusOK) GetPayload() *models.BackupCreateStatusResponse {
	return o.Payload
}

func (o *BackupsCreateStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackupCreateStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupsCreateStatusUnauthorized creates a BackupsCreateStatusUnauthorized with default headers values
func NewBackupsCreateStatusUnauthorized() *BackupsCreateStatusUnauthorized {
	return &BackupsCreateStatusUnauthorized{}
}

/*
BackupsCreateStatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type BackupsCreateStatusUnauthorized struct {
}

// IsSuccess returns true when this backups create status unauthorized response has a 2xx status code
func (o *BackupsCreateStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backups create status unauthorized response has a 3xx status code
func (o *BackupsCreateStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backups create status unauthorized response has a 4xx status code
func (o *BackupsCreateStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this backups create status unauthorized response has a 5xx status code
func (o *BackupsCreateStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this backups create status unauthorized response a status code equal to that given
func (o *BackupsCreateStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the backups create status unauthorized response
func (o *BackupsCreateStatusUnauthorized) Code() int {
	return 401
}

func (o *BackupsCreateStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/backups/{backend}/{id}][%d] backupsCreateStatusUnauthorized ", 401)
}

func (o *BackupsCreateStatusUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/backups/{backend}/{id}][%d] backupsCreateStatusUnauthorized ", 401)
}

func (o *BackupsCreateStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBackupsCreateStatusForbidden creates a BackupsCreateStatusForbidden with default headers values
func NewBackupsCreateStatusForbidden() *BackupsCreateStatusForbidden {
	return &BackupsCreateStatusForbidden{}
}

/*
BackupsCreateStatusForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type BackupsCreateStatusForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this backups create status forbidden response has a 2xx status code
func (o *BackupsCreateStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backups create status forbidden response has a 3xx status code
func (o *BackupsCreateStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backups create status forbidden response has a 4xx status code
func (o *BackupsCreateStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this backups create status forbidden response has a 5xx status code
func (o *BackupsCreateStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this backups create status forbidden response a status code equal to that given
func (o *BackupsCreateStatusForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the backups create status forbidden response
func (o *BackupsCreateStatusForbidden) Code() int {
	return 403
}

func (o *BackupsCreateStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/backups/{backend}/{id}][%d] backupsCreateStatusForbidden  %+v", 403, o.Payload)
}

func (o *BackupsCreateStatusForbidden) String() string {
	return fmt.Sprintf("[GET /v1/backups/{backend}/{id}][%d] backupsCreateStatusForbidden  %+v", 403, o.Payload)
}

func (o *BackupsCreateStatusForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BackupsCreateStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupsCreateStatusNotFound creates a BackupsCreateStatusNotFound with default headers values
func NewBackupsCreateStatusNotFound() *BackupsCreateStatusNotFound {
	return &BackupsCreateStatusNotFound{}
}

/*
BackupsCreateStatusNotFound describes a response with status code 404, with default header values.

Not Found - Backup does not exist
*/
type BackupsCreateStatusNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this backups create status not found response has a 2xx status code
func (o *BackupsCreateStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backups create status not found response has a 3xx status code
func (o *BackupsCreateStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backups create status not found response has a 4xx status code
func (o *BackupsCreateStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this backups create status not found response has a 5xx status code
func (o *BackupsCreateStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this backups create status not found response a status code equal to that given
func (o *BackupsCreateStatusNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the backups create status not found response
func (o *BackupsCreateStatusNotFound) Code() int {
	return 404
}

func (o *BackupsCreateStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/backups/{backend}/{id}][%d] backupsCreateStatusNotFound  %+v", 404, o.Payload)
}

func (o *BackupsCreateStatusNotFound) String() string {
	return fmt.Sprintf("[GET /v1/backups/{backend}/{id}][%d] backupsCreateStatusNotFound  %+v", 404, o.Payload)
}

func (o *BackupsCreateStatusNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BackupsCreateStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupsCreateStatusUnprocessableEntity creates a BackupsCreateStatusUnprocessableEntity with default headers values
func NewBackupsCreateStatusUnprocessableEntity() *BackupsCreateStatusUnprocessableEntity {
	return &BackupsCreateStatusUnprocessableEntity{}
}

/*
BackupsCreateStatusUnprocessableEntity describes a response with status code 422, with default header values.

Invalid backup restoration status attempt.
*/
type BackupsCreateStatusUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this backups create status unprocessable entity response has a 2xx status code
func (o *BackupsCreateStatusUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backups create status unprocessable entity response has a 3xx status code
func (o *BackupsCreateStatusUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backups create status unprocessable entity response has a 4xx status code
func (o *BackupsCreateStatusUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this backups create status unprocessable entity response has a 5xx status code
func (o *BackupsCreateStatusUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this backups create status unprocessable entity response a status code equal to that given
func (o *BackupsCreateStatusUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the backups create status unprocessable entity response
func (o *BackupsCreateStatusUnprocessableEntity) Code() int {
	return 422
}

func (o *BackupsCreateStatusUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /v1/backups/{backend}/{id}][%d] backupsCreateStatusUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *BackupsCreateStatusUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /v1/backups/{backend}/{id}][%d] backupsCreateStatusUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *BackupsCreateStatusUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BackupsCreateStatusUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupsCreateStatusInternalServerError creates a BackupsCreateStatusInternalServerError with default headers values
func NewBackupsCreateStatusInternalServerError() *BackupsCreateStatusInternalServerError {
	return &BackupsCreateStatusInternalServerError{}
}

/*
BackupsCreateStatusInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type BackupsCreateStatusInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this backups create status internal server error response has a 2xx status code
func (o *BackupsCreateStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backups create status internal server error response has a 3xx status code
func (o *BackupsCreateStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backups create status internal server error response has a 4xx status code
func (o *BackupsCreateStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this backups create status internal server error response has a 5xx status code
func (o *BackupsCreateStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this backups create status internal server error response a status code equal to that given
func (o *BackupsCreateStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the backups create status internal server error response
func (o *BackupsCreateStatusInternalServerError) Code() int {
	return 500
}

func (o *BackupsCreateStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/backups/{backend}/{id}][%d] backupsCreateStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *BackupsCreateStatusInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/backups/{backend}/{id}][%d] backupsCreateStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *BackupsCreateStatusInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BackupsCreateStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

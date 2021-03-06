// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/gotify/go-api-client/v2/models"
)

// UpdateApplicationReader is a Reader for the UpdateApplication structure.
type UpdateApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateApplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateApplicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateApplicationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateApplicationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateApplicationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateApplicationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateApplicationOK creates a UpdateApplicationOK with default headers values
func NewUpdateApplicationOK() *UpdateApplicationOK {
	return &UpdateApplicationOK{}
}

/*UpdateApplicationOK handles this case with default header values.

Ok
*/
type UpdateApplicationOK struct {
	Payload *models.Application
}

func (o *UpdateApplicationOK) Error() string {
	return fmt.Sprintf("[PUT /application/{id}][%d] updateApplicationOK  %+v", 200, o.Payload)
}

func (o *UpdateApplicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Application)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateApplicationBadRequest creates a UpdateApplicationBadRequest with default headers values
func NewUpdateApplicationBadRequest() *UpdateApplicationBadRequest {
	return &UpdateApplicationBadRequest{}
}

/*UpdateApplicationBadRequest handles this case with default header values.

Bad Request
*/
type UpdateApplicationBadRequest struct {
	Payload *models.Error
}

func (o *UpdateApplicationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /application/{id}][%d] updateApplicationBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateApplicationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateApplicationUnauthorized creates a UpdateApplicationUnauthorized with default headers values
func NewUpdateApplicationUnauthorized() *UpdateApplicationUnauthorized {
	return &UpdateApplicationUnauthorized{}
}

/*UpdateApplicationUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateApplicationUnauthorized struct {
	Payload *models.Error
}

func (o *UpdateApplicationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /application/{id}][%d] updateApplicationUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateApplicationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateApplicationForbidden creates a UpdateApplicationForbidden with default headers values
func NewUpdateApplicationForbidden() *UpdateApplicationForbidden {
	return &UpdateApplicationForbidden{}
}

/*UpdateApplicationForbidden handles this case with default header values.

Forbidden
*/
type UpdateApplicationForbidden struct {
	Payload *models.Error
}

func (o *UpdateApplicationForbidden) Error() string {
	return fmt.Sprintf("[PUT /application/{id}][%d] updateApplicationForbidden  %+v", 403, o.Payload)
}

func (o *UpdateApplicationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateApplicationNotFound creates a UpdateApplicationNotFound with default headers values
func NewUpdateApplicationNotFound() *UpdateApplicationNotFound {
	return &UpdateApplicationNotFound{}
}

/*UpdateApplicationNotFound handles this case with default header values.

Not Found
*/
type UpdateApplicationNotFound struct {
	Payload *models.Error
}

func (o *UpdateApplicationNotFound) Error() string {
	return fmt.Sprintf("[PUT /application/{id}][%d] updateApplicationNotFound  %+v", 404, o.Payload)
}

func (o *UpdateApplicationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/gotify/server/model"
)

// CreateAppReader is a Reader for the CreateApp structure.
type CreateAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateAppOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateAppBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateAppUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateAppForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateAppOK creates a CreateAppOK with default headers values
func NewCreateAppOK() *CreateAppOK {
	return &CreateAppOK{}
}

/*CreateAppOK handles this case with default header values.

Ok
*/
type CreateAppOK struct {
	Payload *models.Application
}

func (o *CreateAppOK) Error() string {
	return fmt.Sprintf("[POST /application][%d] createAppOK  %+v", 200, o.Payload)
}

func (o *CreateAppOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Application)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppBadRequest creates a CreateAppBadRequest with default headers values
func NewCreateAppBadRequest() *CreateAppBadRequest {
	return &CreateAppBadRequest{}
}

/*CreateAppBadRequest handles this case with default header values.

Bad Request
*/
type CreateAppBadRequest struct {
	Payload *models.Error
}

func (o *CreateAppBadRequest) Error() string {
	return fmt.Sprintf("[POST /application][%d] createAppBadRequest  %+v", 400, o.Payload)
}

func (o *CreateAppBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppUnauthorized creates a CreateAppUnauthorized with default headers values
func NewCreateAppUnauthorized() *CreateAppUnauthorized {
	return &CreateAppUnauthorized{}
}

/*CreateAppUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateAppUnauthorized struct {
	Payload *models.Error
}

func (o *CreateAppUnauthorized) Error() string {
	return fmt.Sprintf("[POST /application][%d] createAppUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateAppUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppForbidden creates a CreateAppForbidden with default headers values
func NewCreateAppForbidden() *CreateAppForbidden {
	return &CreateAppForbidden{}
}

/*CreateAppForbidden handles this case with default header values.

Forbidden
*/
type CreateAppForbidden struct {
	Payload *models.Error
}

func (o *CreateAppForbidden) Error() string {
	return fmt.Sprintf("[POST /application][%d] createAppForbidden  %+v", 403, o.Payload)
}

func (o *CreateAppForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
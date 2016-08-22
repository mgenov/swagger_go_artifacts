package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mgenov/temp/models"
)

/*DeletePetNoContent pet deleted

swagger:response deletePetNoContent
*/
type DeletePetNoContent struct {
}

// NewDeletePetNoContent creates DeletePetNoContent with default headers values
func NewDeletePetNoContent() *DeletePetNoContent {
	return &DeletePetNoContent{}
}

// WriteResponse to the client
func (o *DeletePetNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

/*DeletePetDefault unexpected error

swagger:response deletePetDefault
*/
type DeletePetDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeletePetDefault creates DeletePetDefault with default headers values
func NewDeletePetDefault(code int) *DeletePetDefault {
	if code <= 0 {
		code = 500
	}

	return &DeletePetDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete pet default response
func (o *DeletePetDefault) WithStatusCode(code int) *DeletePetDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete pet default response
func (o *DeletePetDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete pet default response
func (o *DeletePetDefault) WithPayload(payload *models.Error) *DeletePetDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete pet default response
func (o *DeletePetDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePetDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
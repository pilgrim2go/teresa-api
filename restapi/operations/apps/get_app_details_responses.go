package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/luizalabs/tapi/models"
)

/*GetAppDetailsOK App details

swagger:response getAppDetailsOK
*/
type GetAppDetailsOK struct {

	// In: body
	Payload *models.App `json:"body,omitempty"`
}

// NewGetAppDetailsOK creates GetAppDetailsOK with default headers values
func NewGetAppDetailsOK() *GetAppDetailsOK {
	return &GetAppDetailsOK{}
}

// WithPayload adds the payload to the get app details o k response
func (o *GetAppDetailsOK) WithPayload(payload *models.App) *GetAppDetailsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get app details o k response
func (o *GetAppDetailsOK) SetPayload(payload *models.App) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppDetailsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetAppDetailsUnauthorized User not authorized

swagger:response getAppDetailsUnauthorized
*/
type GetAppDetailsUnauthorized struct {

	// In: body
	Payload *models.Unauthorized `json:"body,omitempty"`
}

// NewGetAppDetailsUnauthorized creates GetAppDetailsUnauthorized with default headers values
func NewGetAppDetailsUnauthorized() *GetAppDetailsUnauthorized {
	return &GetAppDetailsUnauthorized{}
}

// WithPayload adds the payload to the get app details unauthorized response
func (o *GetAppDetailsUnauthorized) WithPayload(payload *models.Unauthorized) *GetAppDetailsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get app details unauthorized response
func (o *GetAppDetailsUnauthorized) SetPayload(payload *models.Unauthorized) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppDetailsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetAppDetailsForbidden User does not have the credentials to access this resource


swagger:response getAppDetailsForbidden
*/
type GetAppDetailsForbidden struct {

	// In: body
	Payload *models.Unauthorized `json:"body,omitempty"`
}

// NewGetAppDetailsForbidden creates GetAppDetailsForbidden with default headers values
func NewGetAppDetailsForbidden() *GetAppDetailsForbidden {
	return &GetAppDetailsForbidden{}
}

// WithPayload adds the payload to the get app details forbidden response
func (o *GetAppDetailsForbidden) WithPayload(payload *models.Unauthorized) *GetAppDetailsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get app details forbidden response
func (o *GetAppDetailsForbidden) SetPayload(payload *models.Unauthorized) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppDetailsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetAppDetailsDefault Error

swagger:response getAppDetailsDefault
*/
type GetAppDetailsDefault struct {
	_statusCode int

	// In: body
	Payload *models.GenericError `json:"body,omitempty"`
}

// NewGetAppDetailsDefault creates GetAppDetailsDefault with default headers values
func NewGetAppDetailsDefault(code int) *GetAppDetailsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAppDetailsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get app details default response
func (o *GetAppDetailsDefault) WithStatusCode(code int) *GetAppDetailsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get app details default response
func (o *GetAppDetailsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get app details default response
func (o *GetAppDetailsDefault) WithPayload(payload *models.GenericError) *GetAppDetailsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get app details default response
func (o *GetAppDetailsDefault) SetPayload(payload *models.GenericError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppDetailsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

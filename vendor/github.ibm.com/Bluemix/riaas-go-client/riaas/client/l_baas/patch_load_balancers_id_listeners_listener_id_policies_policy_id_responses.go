// Code generated by go-swagger; DO NOT EDIT.

package l_baas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDReader is a Reader for the PatchLoadBalancersIDListenersListenerIDPoliciesPolicyID structure.
type PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDOK creates a PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDOK with default headers values
func NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDOK() *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDOK {
	return &PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDOK{}
}

/*PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDOK handles this case with default header values.

The policy update request was accepted.
*/
type PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDOK struct {
	Payload *models.ListenerPolicy
}

func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDOK) Error() string {
	return fmt.Sprintf("[PATCH /load_balancers/{id}/listeners/{listener_id}/policies/{policy_id}][%d] patchLoadBalancersIdListenersListenerIdPoliciesPolicyIdOK  %+v", 200, o.Payload)
}

func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListenerPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDBadRequest creates a PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDBadRequest with default headers values
func NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDBadRequest() *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDBadRequest {
	return &PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDBadRequest{}
}

/*PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDBadRequest handles this case with default header values.

An invalid policy template was provided.
*/
type PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /load_balancers/{id}/listeners/{listener_id}/policies/{policy_id}][%d] patchLoadBalancersIdListenersListenerIdPoliciesPolicyIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDNotFound creates a PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDNotFound with default headers values
func NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDNotFound() *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDNotFound {
	return &PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDNotFound{}
}

/*PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDNotFound handles this case with default header values.

A load balancer or listener with the specified identifier could not be found.
*/
type PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /load_balancers/{id}/listeners/{listener_id}/policies/{policy_id}][%d] patchLoadBalancersIdListenersListenerIdPoliciesPolicyIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
/*
Service Manager

Service Manager provides REST APIs that are responsible for the creation and consumption of service instances in any connected runtime environment.   Use the Service Manager APIs to perform various operations related to your platforms, service brokers, service instances, and service bindings.  Get service plans and service offerings associated with your environment.    #### Platforms   Platforms are OSBAPI-enabled software systems on which applications and services are hosted.   With the Service Manager, you can now register your platform and enable it to consume the SAP BTP services from your native environment.   This registration results in a returned set of credentials that are needed to deploy the Service Manager agent.     #### Service Brokers   Service brokers act as brokers between the Service Manager and a platform’s marketplace to advertise catalogues of service offerings and service plans.  They also receive and process the requests from the marketplace to provision, bind, unbind, and deprovision these offerings and plans.    #### Service Instances   Service instances are instantiations of service plans that make the functionality of those service plans available for consumption.    #### Service Bindings   Service bindings provide access details to existing service instances.  The access details are part of the service bindings' ‘credentials’ property, and typically include access URLs and credentials.    #### Service Plans   Service plans represent sets of capabilities provided by a service offering.  For example, database service offerings provide different plans for different database versions or sizes, while the Service Manager plans offer different data access levels.    #### Service Offerings   Service offerings are advertisements of the services that are supported by a service broker.  For example, software that you can consume in the subaccount.  Service offerings are related to one or more service plans.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type ServiceBrokersAPI interface {

	/*
	GetAllServiceBrokers Get all service brokers

	View the list of all service brokers in the subaccount. <br/><br/>Required scopes: <xsappname>.subaccount.broker.read

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetAllServiceBrokersRequest
	*/
	GetAllServiceBrokers(ctx context.Context) ApiGetAllServiceBrokersRequest

	// GetAllServiceBrokersExecute executes the request
	//  @return ServiceBrokerResponseList
	GetAllServiceBrokersExecute(r ApiGetAllServiceBrokersRequest) (*ServiceBrokerResponseList, *http.Response, error)

	/*
	GetServiceBrokerId Get service broker details.

	View details of a specific service broker registered in the subaccount. <br><br>Required scopes: <xsappname>.subaccount.broker.read

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceBrokerID The ID of the service broker for which to get details.
	@return ApiGetServiceBrokerIdRequest
	*/
	GetServiceBrokerId(ctx context.Context, serviceBrokerID string) ApiGetServiceBrokerIdRequest

	// GetServiceBrokerIdExecute executes the request
	//  @return ServiceBrokerResponseObject
	GetServiceBrokerIdExecute(r ApiGetServiceBrokerIdRequest) (*ServiceBrokerResponseObject, *http.Response, error)
}

// ServiceBrokersAPIService ServiceBrokersAPI service
type ServiceBrokersAPIService service

type ApiGetAllServiceBrokersRequest struct {
	ctx context.Context
	ApiService ServiceBrokersAPI
	fieldQuery *string
	labelQuery *string
	token *string
	maxItems *int32
}

// Filters the response based on the field query.  &lt;br/&gt; If used, must be a nonempty string.&lt;br/&gt;For example:&lt;br&gt; name eq &#39;my service broker&#39;.
func (r ApiGetAllServiceBrokersRequest) FieldQuery(fieldQuery string) ApiGetAllServiceBrokersRequest {
	r.fieldQuery = &fieldQuery
	return r
}

// Filters the response based on the label query. &lt;br&gt; If used, must be a nonempty string.&lt;br&gt;For example:&lt;br/&gt; environment eq &#39;dev&#39;.
func (r ApiGetAllServiceBrokersRequest) LabelQuery(labelQuery string) ApiGetAllServiceBrokersRequest {
	r.labelQuery = &labelQuery
	return r
}

// You get this parameter in the response list of the API if the total number of items to return (num_items) is larger than the number of items returned in a single API call (max_items).&lt;br/&gt; You get a different token in each response to be used in each consecutive call as long as there are more items to list.&lt;br/&gt; Use the returned tokens to get the full list of resources associated with your subaccount.&lt;br/&gt; Leave the field empty if this is the first time you are calling the API.
func (r ApiGetAllServiceBrokersRequest) Token(token string) ApiGetAllServiceBrokersRequest {
	r.token = &token
	return r
}

// The maximum number of service brokers to return in the response. 
func (r ApiGetAllServiceBrokersRequest) MaxItems(maxItems int32) ApiGetAllServiceBrokersRequest {
	r.maxItems = &maxItems
	return r
}

func (r ApiGetAllServiceBrokersRequest) Execute() (*ServiceBrokerResponseList, *http.Response, error) {
	return r.ApiService.GetAllServiceBrokersExecute(r)
}

/*
GetAllServiceBrokers Get all service brokers

View the list of all service brokers in the subaccount. <br/><br/>Required scopes: <xsappname>.subaccount.broker.read

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllServiceBrokersRequest
*/
func (a *ServiceBrokersAPIService) GetAllServiceBrokers(ctx context.Context) ApiGetAllServiceBrokersRequest {
	return ApiGetAllServiceBrokersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ServiceBrokerResponseList
func (a *ServiceBrokersAPIService) GetAllServiceBrokersExecute(r ApiGetAllServiceBrokersRequest) (*ServiceBrokerResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceBrokerResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceBrokersAPIService.GetAllServiceBrokers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/service_brokers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldQuery != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fieldQuery", r.fieldQuery, "form", "")
	}
	if r.labelQuery != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "labelQuery", r.labelQuery, "form", "")
	}
	if r.token != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "token", r.token, "form", "")
	}
	if r.maxItems != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "max_items", r.maxItems, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetServiceBrokerIdRequest struct {
	ctx context.Context
	ApiService ServiceBrokersAPI
	serviceBrokerID string
}

func (r ApiGetServiceBrokerIdRequest) Execute() (*ServiceBrokerResponseObject, *http.Response, error) {
	return r.ApiService.GetServiceBrokerIdExecute(r)
}

/*
GetServiceBrokerId Get service broker details.

View details of a specific service broker registered in the subaccount. <br><br>Required scopes: <xsappname>.subaccount.broker.read

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceBrokerID The ID of the service broker for which to get details.
 @return ApiGetServiceBrokerIdRequest
*/
func (a *ServiceBrokersAPIService) GetServiceBrokerId(ctx context.Context, serviceBrokerID string) ApiGetServiceBrokerIdRequest {
	return ApiGetServiceBrokerIdRequest{
		ApiService: a,
		ctx: ctx,
		serviceBrokerID: serviceBrokerID,
	}
}

// Execute executes the request
//  @return ServiceBrokerResponseObject
func (a *ServiceBrokersAPIService) GetServiceBrokerIdExecute(r ApiGetServiceBrokerIdRequest) (*ServiceBrokerResponseObject, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceBrokerResponseObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceBrokersAPIService.GetServiceBrokerId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/service_brokers/{serviceBrokerID}"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceBrokerID"+"}", url.PathEscape(parameterValueToString(r.serviceBrokerID, "serviceBrokerID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

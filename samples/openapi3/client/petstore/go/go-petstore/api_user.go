/*
OpenAPI Petstore

This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type UserApi interface {

	/*
		CreateUser Create user

		This can only be done by the logged in user.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiCreateUserRequest
	*/
	CreateUser(ctx context.Context) ApiCreateUserRequest

	// CreateUserExecute executes the request
	CreateUserExecute(r ApiCreateUserRequest) (*http.Response, error)

	/*
		CreateUsersWithArrayInput Creates list of users with given input array



		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiCreateUsersWithArrayInputRequest
	*/
	CreateUsersWithArrayInput(ctx context.Context) ApiCreateUsersWithArrayInputRequest

	// CreateUsersWithArrayInputExecute executes the request
	CreateUsersWithArrayInputExecute(r ApiCreateUsersWithArrayInputRequest) (*http.Response, error)

	/*
		CreateUsersWithListInput Creates list of users with given input array



		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiCreateUsersWithListInputRequest
	*/
	CreateUsersWithListInput(ctx context.Context) ApiCreateUsersWithListInputRequest

	// CreateUsersWithListInputExecute executes the request
	CreateUsersWithListInputExecute(r ApiCreateUsersWithListInputRequest) (*http.Response, error)

	/*
		DeleteUser Delete user

		This can only be done by the logged in user.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param username The name that needs to be deleted
		@return ApiDeleteUserRequest
	*/
	DeleteUser(ctx context.Context, username string) ApiDeleteUserRequest

	// DeleteUserExecute executes the request
	DeleteUserExecute(r ApiDeleteUserRequest) (*http.Response, error)

	/*
		GetUserByName Get user by user name



		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param username The name that needs to be fetched. Use user1 for testing.
		@return ApiGetUserByNameRequest
	*/
	GetUserByName(ctx context.Context, username string) ApiGetUserByNameRequest

	// GetUserByNameExecute executes the request
	//  @return User
	GetUserByNameExecute(r ApiGetUserByNameRequest) (*User, *http.Response, error)

	/*
		LoginUser Logs user into the system



		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiLoginUserRequest
	*/
	LoginUser(ctx context.Context) ApiLoginUserRequest

	// LoginUserExecute executes the request
	//  @return string
	LoginUserExecute(r ApiLoginUserRequest) (string, *http.Response, error)

	/*
		LogoutUser Logs out current logged in user session



		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiLogoutUserRequest
	*/
	LogoutUser(ctx context.Context) ApiLogoutUserRequest

	// LogoutUserExecute executes the request
	LogoutUserExecute(r ApiLogoutUserRequest) (*http.Response, error)

	/*
		UpdateUser Updated user

		This can only be done by the logged in user.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param username name that need to be deleted
		@return ApiUpdateUserRequest
	*/
	UpdateUser(ctx context.Context, username string) ApiUpdateUserRequest

	// UpdateUserExecute executes the request
	UpdateUserExecute(r ApiUpdateUserRequest) (*http.Response, error)
}

// UserApiService UserApi service
type UserApiService service

type ApiCreateUserRequest struct {
	ctx        context.Context
	ApiService UserApi
	user       *User
}

// Created user object
func (r ApiCreateUserRequest) User(user User) ApiCreateUserRequest {
	r.user = &user
	return r
}

func (r ApiCreateUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateUserExecute(r)
}

/*
CreateUser Create user

This can only be done by the logged in user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateUserRequest
*/
func (a *UserApiService) CreateUser(ctx context.Context) ApiCreateUserRequest {
	return ApiCreateUserRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *UserApiService) CreateUserExecute(r ApiCreateUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.CreateUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.user == nil {
		return nil, reportError("user is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.user
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiCreateUsersWithArrayInputRequest struct {
	ctx        context.Context
	ApiService UserApi
	user       *[]User
}

// List of user object
func (r ApiCreateUsersWithArrayInputRequest) User(user []User) ApiCreateUsersWithArrayInputRequest {
	r.user = &user
	return r
}

func (r ApiCreateUsersWithArrayInputRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateUsersWithArrayInputExecute(r)
}

/*
CreateUsersWithArrayInput Creates list of users with given input array

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateUsersWithArrayInputRequest
*/
func (a *UserApiService) CreateUsersWithArrayInput(ctx context.Context) ApiCreateUsersWithArrayInputRequest {
	return ApiCreateUsersWithArrayInputRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *UserApiService) CreateUsersWithArrayInputExecute(r ApiCreateUsersWithArrayInputRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.CreateUsersWithArrayInput")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/createWithArray"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.user == nil {
		return nil, reportError("user is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.user
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiCreateUsersWithListInputRequest struct {
	ctx        context.Context
	ApiService UserApi
	user       *[]User
}

// List of user object
func (r ApiCreateUsersWithListInputRequest) User(user []User) ApiCreateUsersWithListInputRequest {
	r.user = &user
	return r
}

func (r ApiCreateUsersWithListInputRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateUsersWithListInputExecute(r)
}

/*
CreateUsersWithListInput Creates list of users with given input array

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateUsersWithListInputRequest
*/
func (a *UserApiService) CreateUsersWithListInput(ctx context.Context) ApiCreateUsersWithListInputRequest {
	return ApiCreateUsersWithListInputRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *UserApiService) CreateUsersWithListInputExecute(r ApiCreateUsersWithListInputRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.CreateUsersWithListInput")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/createWithList"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.user == nil {
		return nil, reportError("user is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.user
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteUserRequest struct {
	ctx        context.Context
	ApiService UserApi
	username   string
}

func (r ApiDeleteUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteUserExecute(r)
}

/*
DeleteUser Delete user

This can only be done by the logged in user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param username The name that needs to be deleted
	@return ApiDeleteUserRequest
*/
func (a *UserApiService) DeleteUser(ctx context.Context, username string) ApiDeleteUserRequest {
	return ApiDeleteUserRequest{
		ApiService: a,
		ctx:        ctx,
		username:   username,
	}
}

// Execute executes the request
func (a *UserApiService) DeleteUserExecute(r ApiDeleteUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.DeleteUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/{username}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetUserByNameRequest struct {
	ctx        context.Context
	ApiService UserApi
	username   string
}

func (r ApiGetUserByNameRequest) Execute() (*User, *http.Response, error) {
	return r.ApiService.GetUserByNameExecute(r)
}

/*
GetUserByName Get user by user name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param username The name that needs to be fetched. Use user1 for testing.
	@return ApiGetUserByNameRequest
*/
func (a *UserApiService) GetUserByName(ctx context.Context, username string) ApiGetUserByNameRequest {
	return ApiGetUserByNameRequest{
		ApiService: a,
		ctx:        ctx,
		username:   username,
	}
}

// Execute executes the request
//
//	@return User
func (a *UserApiService) GetUserByNameExecute(r ApiGetUserByNameRequest) (*User, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *User
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.GetUserByName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/{username}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/xml", "application/json"}

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

type ApiLoginUserRequest struct {
	ctx        context.Context
	ApiService UserApi
	username   *string
	password   *string
}

// The user name for login
func (r ApiLoginUserRequest) Username(username string) ApiLoginUserRequest {
	r.username = &username
	return r
}

// The password for login in clear text
func (r ApiLoginUserRequest) Password(password string) ApiLoginUserRequest {
	r.password = &password
	return r
}

func (r ApiLoginUserRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.LoginUserExecute(r)
}

/*
LoginUser Logs user into the system

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiLoginUserRequest
*/
func (a *UserApiService) LoginUser(ctx context.Context) ApiLoginUserRequest {
	return ApiLoginUserRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return string
func (a *UserApiService) LoginUserExecute(r ApiLoginUserRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.LoginUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/login"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.username == nil {
		return localVarReturnValue, nil, reportError("username is required and must be specified")
	}
	if r.password == nil {
		return localVarReturnValue, nil, reportError("password is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "username", r.username, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "password", r.password, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/xml", "application/json"}

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

type ApiLogoutUserRequest struct {
	ctx        context.Context
	ApiService UserApi
}

func (r ApiLogoutUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.LogoutUserExecute(r)
}

/*
LogoutUser Logs out current logged in user session

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiLogoutUserRequest
*/
func (a *UserApiService) LogoutUser(ctx context.Context) ApiLogoutUserRequest {
	return ApiLogoutUserRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *UserApiService) LogoutUserExecute(r ApiLogoutUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.LogoutUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/logout"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdateUserRequest struct {
	ctx        context.Context
	ApiService UserApi
	username   string
	user       *User
}

// Updated user object
func (r ApiUpdateUserRequest) User(user User) ApiUpdateUserRequest {
	r.user = &user
	return r
}

func (r ApiUpdateUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateUserExecute(r)
}

/*
UpdateUser Updated user

This can only be done by the logged in user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param username name that need to be deleted
	@return ApiUpdateUserRequest
*/
func (a *UserApiService) UpdateUser(ctx context.Context, username string) ApiUpdateUserRequest {
	return ApiUpdateUserRequest{
		ApiService: a,
		ctx:        ctx,
		username:   username,
	}
}

// Execute executes the request
func (a *UserApiService) UpdateUserExecute(r ApiUpdateUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.UpdateUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/{username}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.user == nil {
		return nil, reportError("user is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.user
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

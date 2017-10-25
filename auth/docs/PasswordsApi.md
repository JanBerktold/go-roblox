# \PasswordsApi

All URIs are relative to *https://auth.roblox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**POSTV1PasswordsReset**](PasswordsApi.md#POSTV1PasswordsReset) | **Post** /v1/passwords/reset | Resets a password for a user that belongs to the password reset ticket.
[**POSTV1PasswordsResetSend**](PasswordsApi.md#POSTV1PasswordsResetSend) | **Post** /v1/passwords/reset/send | Sends a password reset email to the specified user.
[**POSTV1UserPasswordsChange**](PasswordsApi.md#POSTV1UserPasswordsChange) | **Post** /v1/user/passwords/change | Changes the password for the authenticated user.


# **POSTV1PasswordsReset**
> ApiEmptyResponseModel POSTV1PasswordsReset($request)

Resets a password for a user that belongs to the password reset ticket.

The user will be logged out after the password changes.  A call to login should be made after calling this endpoint, and getting a succesful response to log the user in.  This endpoint is not responsible for authenticating the user.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**PasswordResetModel**](PasswordResetModel.md)| The request model including the ticket, and new password. | 

### Return type

[**ApiEmptyResponseModel**](ApiEmptyResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTV1PasswordsResetSend**
> ApiEmptyResponseModel POSTV1PasswordsResetSend($request)

Sends a password reset email to the specified user.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**UsernameModel**](UsernameModel.md)| The request model containing the username of the user. | 

### Return type

[**ApiEmptyResponseModel**](ApiEmptyResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTV1UserPasswordsChange**
> ApiEmptyResponseModel POSTV1UserPasswordsChange($request)

Changes the password for the authenticated user.

The current password is needed for verification that the password can be changed.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**PasswordChangeModel**](PasswordChangeModel.md)| The request model including the current password, and the new password. | 

### Return type

[**ApiEmptyResponseModel**](ApiEmptyResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


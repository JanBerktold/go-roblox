# \AccountPinApi

All URIs are relative to *https://auth.roblox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEV1AccountPin**](AccountPinApi.md#DELETEV1AccountPin) | **Delete** /v1/account/pin | Request for deletes the account pin from the account.
[**GETV1AccountPin**](AccountPinApi.md#GETV1AccountPin) | **Get** /v1/account/pin | Gets the account pin status. If the account pin is valid, this returns the time in seconds until when the account pin is unlocked.
[**PATCHV1AccountPin**](AccountPinApi.md#PATCHV1AccountPin) | **Patch** /v1/account/pin | Request made to update the account pin on the account.
[**POSTV1AccountPin**](AccountPinApi.md#POSTV1AccountPin) | **Post** /v1/account/pin | Reuqest to create the account pin.
[**POSTV1AccountPinLock**](AccountPinApi.md#POSTV1AccountPinLock) | **Post** /v1/account/pin/lock | Request to locks the account which has an account pin enabled.
[**POSTV1AccountPinUnlock**](AccountPinApi.md#POSTV1AccountPinUnlock) | **Post** /v1/account/pin/unlock | Requests to unlock the account pin.


# **DELETEV1AccountPin**
> ApiSuccessResponse DELETEV1AccountPin()

Request for deletes the account pin from the account.


### Parameters
This endpoint does not need any parameter.

### Return type

[**ApiSuccessResponse**](ApiSuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETV1AccountPin**
> AccountPinStatusResponseModel GETV1AccountPin()

Gets the account pin status. If the account pin is valid, this returns the time in seconds until when the account pin is unlocked.


### Parameters
This endpoint does not need any parameter.

### Return type

[**AccountPinStatusResponseModel**](AccountPinStatusResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PATCHV1AccountPin**
> ApiSuccessResponse PATCHV1AccountPin($requestBody)

Request made to update the account pin on the account.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | [**AccountPinRequestModel**](AccountPinRequestModel.md)| The request body. | 

### Return type

[**ApiSuccessResponse**](ApiSuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTV1AccountPin**
> ApiSuccessResponse POSTV1AccountPin($requestBody)

Reuqest to create the account pin.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | [**AccountPinRequestModel**](AccountPinRequestModel.md)| The {Roblox.Api.Authentication.Models.AccountPinRequestModel}. | 

### Return type

[**ApiSuccessResponse**](ApiSuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTV1AccountPinLock**
> ApiSuccessResponse POSTV1AccountPinLock()

Request to locks the account which has an account pin enabled.


### Parameters
This endpoint does not need any parameter.

### Return type

[**ApiSuccessResponse**](ApiSuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTV1AccountPinUnlock**
> AccountPinResponseModel POSTV1AccountPinUnlock($requestBody)

Requests to unlock the account pin.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | [**AccountPinRequestModel**](AccountPinRequestModel.md)| The {Roblox.Api.Authentication.Models.AccountPinRequestModel} containing the entered pin. | 

### Return type

[**AccountPinResponseModel**](AccountPinResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


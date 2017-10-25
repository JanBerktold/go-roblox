# \ValidatorsApi

All URIs are relative to *https://auth.roblox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETV1ValidatorsEmail**](ValidatorsApi.md#GETV1ValidatorsEmail) | **Get** /v1/validators/email | 
[**GETV1ValidatorsUsername**](ValidatorsApi.md#GETV1ValidatorsUsername) | **Get** /v1/validators/username | Tries to get a valid username if the current username is taken


# **GETV1ValidatorsEmail**
> EmailValidationResponseModel GETV1ValidatorsEmail($requestBodyEmail)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBodyEmail** | **string**| Gets or sets the email to check for validation | [optional] 

### Return type

[**EmailValidationResponseModel**](EmailValidationResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETV1ValidatorsUsername**
> RecommendedUsernameResponseModel GETV1ValidatorsUsername($requestBodyUsername, $requestBodyBirthday)

Tries to get a valid username if the current username is taken


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBodyUsername** | **string**| Gets or sets the username to use as the base username provided by the user | [optional] 
 **requestBodyBirthday** | **time.Time**| Gets or sets the birth day. | [optional] 

### Return type

[**RecommendedUsernameResponseModel**](RecommendedUsernameResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


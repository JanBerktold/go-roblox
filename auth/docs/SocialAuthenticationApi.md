# \SocialAuthenticationApi

All URIs are relative to *https://auth.roblox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETV1SocialConnectedProviders**](SocialAuthenticationApi.md#GETV1SocialConnectedProviders) | **Get** /v1/social/connected-providers | Get social network user information if the given social auth method is connected to current user.
[**POSTV1SocialProviderDisconnect**](SocialAuthenticationApi.md#POSTV1SocialProviderDisconnect) | **Post** /v1/social/{provider}/disconnect | Remove the given social provider auth method from current Roblox user if it is connected.


# **GETV1SocialConnectedProviders**
> SocialProvidersModel GETV1SocialConnectedProviders()

Get social network user information if the given social auth method is connected to current user.


### Parameters
This endpoint does not need any parameter.

### Return type

[**SocialProvidersModel**](SocialProvidersModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTV1SocialProviderDisconnect**
> ApiSuccessResponse POSTV1SocialProviderDisconnect($provider)

Remove the given social provider auth method from current Roblox user if it is connected.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string**| The social network provider, e.g., Facebook. | 

### Return type

[**ApiSuccessResponse**](ApiSuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


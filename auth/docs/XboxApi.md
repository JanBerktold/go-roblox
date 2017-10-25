# \XboxApi

All URIs are relative to *https://auth.roblox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETV1XboxConnection**](XboxApi.md#GETV1XboxConnection) | **Get** /v1/xbox/connection | Check if the current user has an Xbox connected.
[**POSTV1XboxDisconnect**](XboxApi.md#POSTV1XboxDisconnect) | **Post** /v1/xbox/disconnect | Unlink the current ROBLOX account from the Xbox live account.


# **GETV1XboxConnection**
> XboxConnectionModel GETV1XboxConnection()

Check if the current user has an Xbox connected.


### Parameters
This endpoint does not need any parameter.

### Return type

[**XboxConnectionModel**](XboxConnectionModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTV1XboxDisconnect**
> ApiSuccessResponse POSTV1XboxDisconnect()

Unlink the current ROBLOX account from the Xbox live account.


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


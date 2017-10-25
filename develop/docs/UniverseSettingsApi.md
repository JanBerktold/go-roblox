# \UniverseSettingsApi

All URIs are relative to *https://develop.roblox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETV1UniversesUniverseIdConfiguration**](UniverseSettingsApi.md#GETV1UniversesUniverseIdConfiguration) | **Get** /v1/universes/{universeId}/configuration | Get settings for an owned universe.
[**GETV1UniversesUniverseIdConfigurationVipServers**](UniverseSettingsApi.md#GETV1UniversesUniverseIdConfigurationVipServers) | **Get** /v1/universes/{universeId}/configuration/vip-servers | Get settings for an owned universe&#39;s VIP servers.
[**PATCHV1UniversesUniverseIdConfiguration**](UniverseSettingsApi.md#PATCHV1UniversesUniverseIdConfiguration) | **Patch** /v1/universes/{universeId}/configuration | Update universe settings for an owned universe.


# **GETV1UniversesUniverseIdConfiguration**
> UniverseSettingsResponse GETV1UniversesUniverseIdConfiguration($universeId)

Get settings for an owned universe.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **universeId** | **int64**| The universe Id. | 

### Return type

[**UniverseSettingsResponse**](UniverseSettingsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETV1UniversesUniverseIdConfigurationVipServers**
> PrivateServerDetailsResponse GETV1UniversesUniverseIdConfigurationVipServers($universeId)

Get settings for an owned universe's VIP servers.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **universeId** | **int64**| The universe Id. | 

### Return type

[**PrivateServerDetailsResponse**](PrivateServerDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PATCHV1UniversesUniverseIdConfiguration**
> UniverseSettingsResponse PATCHV1UniversesUniverseIdConfiguration($universeId, $model)

Update universe settings for an owned universe.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **universeId** | **int64**| The universeId. | 
 **model** | [**UniverseSettingsRequest**](UniverseSettingsRequest.md)| The {Roblox.Api.Develop.Models.UniverseSettingsRequest} model. | 

### Return type

[**UniverseSettingsResponse**](UniverseSettingsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


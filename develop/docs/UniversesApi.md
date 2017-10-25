# \UniversesApi

All URIs are relative to *https://develop.roblox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETV1UniversesMultiget**](UniversesApi.md#GETV1UniversesMultiget) | **Get** /v1/universes/multiget | Gets a {System.Collections.Generic.List&#x60;1}.
[**GETV1UniversesMultigetPermissions**](UniversesApi.md#GETV1UniversesMultigetPermissions) | **Get** /v1/universes/multiget/permissions | Returns an array of granted and declined permissions related to the universes with the ids in {universeIds} for the authenticated user.
[**GETV1UniversesUniverseId**](UniversesApi.md#GETV1UniversesUniverseId) | **Get** /v1/universes/{universeId} | Gets a {Roblox.Api.Develop.Models.UniverseModel}.
[**GETV1UniversesUniverseIdPermissions**](UniversesApi.md#GETV1UniversesUniverseIdPermissions) | **Get** /v1/universes/{universeId}/permissions | Returns list of granted and declined permissions related to the universe with the id {universeId} for authenticated user
[**GETV1UniversesUniverseIdPlaces**](UniversesApi.md#GETV1UniversesUniverseIdPlaces) | **Get** /v1/universes/{universeId}/places | Gets a list of places for a universe.
[**POSTV1UniversesUniverseIdActivate**](UniversesApi.md#POSTV1UniversesUniverseIdActivate) | **Post** /v1/universes/{universeId}/activate | Activates a universes.
[**POSTV1UniversesUniverseIdDeactivate**](UniversesApi.md#POSTV1UniversesUniverseIdDeactivate) | **Post** /v1/universes/{universeId}/deactivate | Deactivates a universes.


# **GETV1UniversesMultiget**
> ApiMultiGetResponseUniverseModel GETV1UniversesMultiget($ids)

Gets a {System.Collections.Generic.List`1}.

If a universe can not be found for a given ID (such as -1) it will be skipped.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**[]int64**](int64.md)| The universe IDs to get. Limit 100. | 

### Return type

[**ApiMultiGetResponseUniverseModel**](ApiMultiGetResponse[UniverseModel].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETV1UniversesMultigetPermissions**
> ApiMultiGetResponseUniverseIdPermissionsModel GETV1UniversesMultigetPermissions($ids)

Returns an array of granted and declined permissions related to the universes with the ids in {universeIds} for the authenticated user.

If a universe can not be found for a given ID (such as -1) it will be skipped.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**[]int64**](int64.md)| The universe ids. | 

### Return type

[**ApiMultiGetResponseUniverseIdPermissionsModel**](ApiMultiGetResponse[UniverseIdPermissionsModel].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETV1UniversesUniverseId**
> UniverseModel GETV1UniversesUniverseId($universeId)

Gets a {Roblox.Api.Develop.Models.UniverseModel}.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **universeId** | **int64**| The Universe id. | 

### Return type

[**UniverseModel**](UniverseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETV1UniversesUniverseIdPermissions**
> UniversePermissionsModel GETV1UniversesUniverseIdPermissions($universeId)

Returns list of granted and declined permissions related to the universe with the id {universeId} for authenticated user


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **universeId** | **int64**| The universe id. | 

### Return type

[**UniversePermissionsModel**](UniversePermissionsModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETV1UniversesUniverseIdPlaces**
> ApiPageResponsePlaceModel GETV1UniversesUniverseIdPlaces($universeId, $sortOrder, $limit, $cursor)

Gets a list of places for a universe.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **universeId** | **int64**| The asset id. | 
 **sortOrder** | **string**| Sorted by placeId | [optional] [default to Asc]
 **limit** | **int32**| The amount of results per request. | [optional] [default to 10]
 **cursor** | **string**| The paging cursor for the previous or next page. | [optional] 

### Return type

[**ApiPageResponsePlaceModel**](ApiPageResponse[PlaceModel].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTV1UniversesUniverseIdActivate**
> ApiEmptyResponseModel POSTV1UniversesUniverseIdActivate($universeId)

Activates a universes.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **universeId** | **int64**| The universe id. | 

### Return type

[**ApiEmptyResponseModel**](ApiEmptyResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTV1UniversesUniverseIdDeactivate**
> ApiEmptyResponseModel POSTV1UniversesUniverseIdDeactivate($universeId)

Deactivates a universes.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **universeId** | **int64**| The universe id. | 

### Return type

[**ApiEmptyResponseModel**](ApiEmptyResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# \UserApi

All URIs are relative to *https://develop.roblox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETV1UserGroupsCanmanage**](UserApi.md#GETV1UserGroupsCanmanage) | **Get** /v1/user/groups/canmanage | Gets a list of Groups that a user can manage.
[**GETV1UserStudiodata**](UserApi.md#GETV1UserStudiodata) | **Get** /v1/user/studiodata | Retrieves a JSON object from persistant storage.
[**GETV1UserUniverses**](UserApi.md#GETV1UserUniverses) | **Get** /v1/user/universes | Gets a list of universes for the authenticated user.
[**GETV1UserUniversesInvitedforteamcreate**](UserApi.md#GETV1UserUniversesInvitedforteamcreate) | **Get** /v1/user/universes/invitedforteamcreate | Gets the invited team edit universes.
[**POSTV1UserStudiodata**](UserApi.md#POSTV1UserStudiodata) | **Post** /v1/user/studiodata | Saves a JSON object to persistent storage.


# **GETV1UserGroupsCanmanage**
> ApiArrayResponseGroupModel GETV1UserGroupsCanmanage()

Gets a list of Groups that a user can manage.


### Parameters
This endpoint does not need any parameter.

### Return type

[**ApiArrayResponseGroupModel**](ApiArrayResponse[GroupModel].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETV1UserStudiodata**
> interface{} GETV1UserStudiodata($clientKey)

Retrieves a JSON object from persistant storage.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientKey** | **string**| A Key to find this data under. Optional. | [optional] 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETV1UserUniverses**
> ApiPageResponseUniverseModel GETV1UserUniverses($sortOrder, $limit, $cursor)

Gets a list of universes for the authenticated user.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortOrder** | **string**| Sorted by universeId | [optional] [default to Asc]
 **limit** | **int32**| The amount of results per request. | [optional] [default to 10]
 **cursor** | **string**| The paging cursor for the previous or next page. | [optional] 

### Return type

[**ApiPageResponseUniverseModel**](ApiPageResponse[UniverseModel].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETV1UserUniversesInvitedforteamcreate**
> ApiPageResponseUniverseModel GETV1UserUniversesInvitedforteamcreate($sortOrder, $limit, $cursor)

Gets the invited team edit universes.

This endpoint is under development!


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortOrder** | **string**| The order the results are sorted in. | [optional] [default to Asc]
 **limit** | **int32**| The amount of results per request. | [optional] [default to 10]
 **cursor** | **string**| The paging cursor for the previous or next page. | [optional] 

### Return type

[**ApiPageResponseUniverseModel**](ApiPageResponse[UniverseModel].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTV1UserStudiodata**
> ApiSuccessResponse POSTV1UserStudiodata($data, $clientKey)

Saves a JSON object to persistent storage.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**interface{}**](interface{}.md)| The data. | 
 **clientKey** | **string**| A Key to save this data under. Optional. | [optional] 

### Return type

[**ApiSuccessResponse**](ApiSuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


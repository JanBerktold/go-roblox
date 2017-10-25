# \TeamCreateApi

All URIs are relative to *https://develop.roblox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEV1UniversesUniverseIdTeamcreateMemberships**](TeamCreateApi.md#DELETEV1UniversesUniverseIdTeamcreateMemberships) | **Delete** /v1/universes/{universeId}/teamcreate/memberships | Removes a user from a TeamCreate permissions list.
[**GETV1UniversesUniverseIdTeamcreate**](TeamCreateApi.md#GETV1UniversesUniverseIdTeamcreate) | **Get** /v1/universes/{universeId}/teamcreate | Gets TeamCreate settings for an {Roblox.Platform.Universes.IUniverse}.
[**GETV1UniversesUniverseIdTeamcreateMemberships**](TeamCreateApi.md#GETV1UniversesUniverseIdTeamcreateMemberships) | **Get** /v1/universes/{universeId}/teamcreate/memberships | List of users allowed to TeamCreate a universe.
[**GETV1UserTeamcreateMemberships**](TeamCreateApi.md#GETV1UserTeamcreateMemberships) | **Get** /v1/user/teamcreate/memberships | List of universes the authenticated user has permission to TeamCreate.
[**PATCHV1UniversesUniverseIdTeamcreate**](TeamCreateApi.md#PATCHV1UniversesUniverseIdTeamcreate) | **Patch** /v1/universes/{universeId}/teamcreate | Edit team create settings for a universe.
[**POSTV1UniversesUniverseIdTeamcreateMemberships**](TeamCreateApi.md#POSTV1UniversesUniverseIdTeamcreateMemberships) | **Post** /v1/universes/{universeId}/teamcreate/memberships | Adds a user to a TeamCreate permissions list.


# **DELETEV1UniversesUniverseIdTeamcreateMemberships**
> ApiEmptyResponseModel DELETEV1UniversesUniverseIdTeamcreateMemberships($universeId, $request)

Removes a user from a TeamCreate permissions list.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **universeId** | **int64**| The universe Id. | 
 **request** | [**TeamCreateMembershipRequest**](TeamCreateMembershipRequest.md)| The request body. | 

### Return type

[**ApiEmptyResponseModel**](ApiEmptyResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETV1UniversesUniverseIdTeamcreate**
> TeamCreateSettingsResponse GETV1UniversesUniverseIdTeamcreate($universeId)

Gets TeamCreate settings for an {Roblox.Platform.Universes.IUniverse}.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **universeId** | **int64**| The universe Id. | 

### Return type

[**TeamCreateSettingsResponse**](TeamCreateSettingsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETV1UniversesUniverseIdTeamcreateMemberships**
> ApiPageResponseUserResponse GETV1UniversesUniverseIdTeamcreateMemberships($universeId, $sortOrder, $limit, $cursor)

List of users allowed to TeamCreate a universe.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **universeId** | **int64**| The universe Id. | 
 **sortOrder** | **string**| TeamCreate membership grant date | [optional] [default to Asc]
 **limit** | **int32**| The amount of results per request. | [optional] [default to 10]
 **cursor** | **string**| The paging cursor for the previous or next page. | [optional] 

### Return type

[**ApiPageResponseUserResponse**](ApiPageResponse[UserResponse].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETV1UserTeamcreateMemberships**
> ApiPageResponseUniverseModel GETV1UserTeamcreateMemberships($sortOrder, $limit, $cursor)

List of universes the authenticated user has permission to TeamCreate.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortOrder** | **string**| TeamCreate membership grant date | [optional] [default to Asc]
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

# **PATCHV1UniversesUniverseIdTeamcreate**
> ApiEmptyResponseModel PATCHV1UniversesUniverseIdTeamcreate($universeId, $request)

Edit team create settings for a universe.

Enables, or disables team create for a universe.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **universeId** | **int64**| The universe Id. | 
 **request** | [**UpdateTeamCreateSettingsRequest**](UpdateTeamCreateSettingsRequest.md)| The request body containing the team create settings. | 

### Return type

[**ApiEmptyResponseModel**](ApiEmptyResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTV1UniversesUniverseIdTeamcreateMemberships**
> ApiEmptyResponseModel POSTV1UniversesUniverseIdTeamcreateMemberships($universeId, $request)

Adds a user to a TeamCreate permissions list.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **universeId** | **int64**| The universe Id. | 
 **request** | [**TeamCreateMembershipRequest**](TeamCreateMembershipRequest.md)| The request body. | 

### Return type

[**ApiEmptyResponseModel**](ApiEmptyResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


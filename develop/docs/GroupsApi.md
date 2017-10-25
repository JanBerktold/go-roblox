# \GroupsApi

All URIs are relative to *https://develop.roblox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETV1GroupsGroupIdUniverses**](GroupsApi.md#GETV1GroupsGroupIdUniverses) | **Get** /v1/groups/{groupId}/universes | Gets a list of universes for the given group.


# **GETV1GroupsGroupIdUniverses**
> ApiPageResponseUniverseModel GETV1GroupsGroupIdUniverses($groupId, $sortOrder, $limit, $cursor)

Gets a list of universes for the given group.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **int32**| The group id. | 
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


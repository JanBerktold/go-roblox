# \MembershipApi

All URIs are relative to *https://groups.roblox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETV1GroupsGroupIdRolesRoleSetIdUsers**](MembershipApi.md#GETV1GroupsGroupIdRolesRoleSetIdUsers) | **Get** /v1/groups/{groupId}/roles/{roleSetId}/users | Gets a list of users in a group for a specific roleset.


# **GETV1GroupsGroupIdRolesRoleSetIdUsers**
> ApiPageResponseUserModel GETV1GroupsGroupIdRolesRoleSetIdUsers($groupId, $roleSetId, $sortOrder, $limit, $cursor)

Gets a list of users in a group for a specific roleset.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **int32**| The group id. | 
 **roleSetId** | **int32**| The group&#39;s role set id. | 
 **sortOrder** | **string**| Sorted by user group join date | [optional] [default to Asc]
 **limit** | **int32**| The amount of results per request. | [optional] [default to 10]
 **cursor** | **string**| The paging cursor for the previous or next page. | [optional] 

### Return type

[**ApiPageResponseUserModel**](ApiPageResponse[UserModel].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


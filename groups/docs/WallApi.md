# \WallApi

All URIs are relative to *https://groups.roblox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEV1GroupsGroupIdWallPostsPostId**](WallApi.md#DELETEV1GroupsGroupIdWallPostsPostId) | **Delete** /v1/groups/{groupId}/wall/posts/{postId} | Deletes a group wall post.
[**GETV1GroupsGroupIdWallPosts**](WallApi.md#GETV1GroupsGroupIdWallPosts) | **Get** /v1/groups/{groupId}/wall/posts | Gets a list of group wall posts.


# **DELETEV1GroupsGroupIdWallPostsPostId**
> ApiEmptyResponseModel DELETEV1GroupsGroupIdWallPostsPostId($groupId, $postId)

Deletes a group wall post.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **int32**| The group id. | 
 **postId** | **int32**| The group wall post id. | 

### Return type

[**ApiEmptyResponseModel**](ApiEmptyResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETV1GroupsGroupIdWallPosts**
> ApiPageResponseGroupWallPostModel GETV1GroupsGroupIdWallPosts($groupId, $sortOrder, $limit, $cursor)

Gets a list of group wall posts.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **int32**| The group id. | 
 **sortOrder** | **string**| Sorted by group wall post Id | [optional] [default to Asc]
 **limit** | **int32**| The amount of results per request. | [optional] [default to 10]
 **cursor** | **string**| The paging cursor for the previous or next page. | [optional] 

### Return type

[**ApiPageResponseGroupWallPostModel**](ApiPageResponse[GroupWallPostModel].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


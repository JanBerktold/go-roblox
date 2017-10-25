# \InventoryApi

All URIs are relative to *https://inventory.roblox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETV1UsersUserIdAssetsCollectibles**](InventoryApi.md#GETV1UsersUserIdAssetsCollectibles) | **Get** /v1/users/{userId}/assets/collectibles | Gets all collectible assets owned by the specified user.
[**GETV1UsersUserIdInventoryAssetType**](InventoryApi.md#GETV1UsersUserIdInventoryAssetType) | **Get** /v1/users/{userId}/inventory/{assetType} | Gets a list of assets by type for the specified user.  Note that the &#39;Hat&#39; asset type may return accessories while we are migrating.


# **GETV1UsersUserIdAssetsCollectibles**
> ApiPageResponseCollectibleUserAssetModel GETV1UsersUserIdAssetsCollectibles($userId, $assetType, $sortOrder, $limit, $cursor)

Gets all collectible assets owned by the specified user.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The userid of the owner of the collectibles. | 
 **assetType** | **string**| The asset type for the collectibles you&#39;re trying to get. | 
 **sortOrder** | **string**| Sorted by userAssetId | [optional] [default to Asc]
 **limit** | **int32**| The amount of results per request. | [optional] [default to 10]
 **cursor** | **string**| The paging cursor for the previous or next page. | [optional] 

### Return type

[**ApiPageResponseCollectibleUserAssetModel**](ApiPageResponse[CollectibleUserAssetModel].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETV1UsersUserIdInventoryAssetType**
> InventoryPageResponse GETV1UsersUserIdInventoryAssetType($userId, $assetType, $pageNumber, $itemsPerPage, $keyword)

Gets a list of assets by type for the specified user.  Note that the 'Hat' asset type may return accessories while we are migrating.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The user identifier. | 
 **assetType** | **string**| The asset type. | 
 **pageNumber** | **int32**| The start index. | [optional] 
 **itemsPerPage** | **int32**| The max number of items that can be returned. | [optional] 
 **keyword** | **string**| Filter results for items containing this. | [optional] 

### Return type

[**InventoryPageResponse**](InventoryPageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


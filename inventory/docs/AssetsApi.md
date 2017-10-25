# \AssetsApi

All URIs are relative to *https://inventory.roblox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETV1AssetsAssetIdOwners**](AssetsApi.md#GETV1AssetsAssetIdOwners) | **Get** /v1/assets/{assetId}/owners | Gets a list of owners of an asset.
[**GETV1PackagesPackageIdAssets**](AssetsApi.md#GETV1PackagesPackageIdAssets) | **Get** /v1/packages/{packageID}/assets | Given a package ID, returns the list of asset IDs for that package
[**GETV1RecommendationsAssetTypeId**](AssetsApi.md#GETV1RecommendationsAssetTypeId) | **Get** /v1/recommendations/{assetTypeId} | Gets a list of recommended assets, given an asset type


# **GETV1AssetsAssetIdOwners**
> ApiPageResponseUserAssetModel GETV1AssetsAssetIdOwners($assetId, $sortOrder, $limit, $cursor)

Gets a list of owners of an asset.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetId** | **int64**| The asset id. | 
 **sortOrder** | **string**| Sorted by userAssetId | [optional] [default to Asc]
 **limit** | **int32**| The amount of results per request. | [optional] [default to 10]
 **cursor** | **string**| The paging cursor for the previous or next page. | [optional] 

### Return type

[**ApiPageResponseUserAssetModel**](ApiPageResponse[UserAssetModel].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETV1PackagesPackageIdAssets**
> AssetIdListModel GETV1PackagesPackageIdAssets($packageID)

Given a package ID, returns the list of asset IDs for that package


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **packageID** | **int64**| The asset ID of the package | 

### Return type

[**AssetIdListModel**](AssetIdListModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETV1RecommendationsAssetTypeId**
> ApiLegacyPageResponseRecommendationViewModel GETV1RecommendationsAssetTypeId($assetTypeId, $numItems, $contextAssetId, $thumbWidth, $thumbHeight)

Gets a list of recommended assets, given an asset type


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetTypeId** | **int32**| The asset type id | 
 **numItems** | **int32**| The number of recommended items to return. | [optional] 
 **contextAssetId** | **int64**|  | [optional] 
 **thumbWidth** | **int32**|  | [optional] 
 **thumbHeight** | **int32**|  | [optional] 

### Return type

[**ApiLegacyPageResponseRecommendationViewModel**](ApiLegacyPageResponse[RecommendationViewModel].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


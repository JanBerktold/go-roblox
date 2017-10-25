# \PlacesApi

All URIs are relative to *https://develop.roblox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PATCHV1PlacesPlaceId**](PlacesApi.md#PATCHV1PlacesPlaceId) | **Patch** /v1/places/{placeId} | Updates the place configuration for the place with the id {placeId}
[**POSTV1PlacesPlaceId**](PlacesApi.md#POSTV1PlacesPlaceId) | **Post** /v1/places/{placeId} | Updates the place configuration for the place with the id {placeId}


# **PATCHV1PlacesPlaceId**
> PlaceModel PATCHV1PlacesPlaceId($placeId, $configuration)

Updates the place configuration for the place with the id {placeId}

Currently the only supported functionality for updating the configuration is around Name, and Description.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **placeId** | **int64**| The place id for the place to be updated. | 
 **configuration** | [**PlaceConfigurationModel**](PlaceConfigurationModel.md)|  | 

### Return type

[**PlaceModel**](PlaceModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTV1PlacesPlaceId**
> PlaceModel POSTV1PlacesPlaceId($placeId, $configuration)

Updates the place configuration for the place with the id {placeId}

Currently the only supported functionality for updating the configuration is around Name, and Description.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **placeId** | **int64**| The place id for the place to be updated. | 
 **configuration** | [**PlaceConfigurationModel**](PlaceConfigurationModel.md)|  | 

### Return type

[**PlaceModel**](PlaceModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


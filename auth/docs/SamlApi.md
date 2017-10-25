# \SamlApi

All URIs are relative to *https://auth.roblox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETV1SamlMetadata**](SamlApi.md#GETV1SamlMetadata) | **Get** /v1/saml/metadata | Gets the SAML2 metadata XML.
[**POSTV1SamlLogin**](SamlApi.md#POSTV1SamlLogin) | **Post** /v1/saml/login | Authenticates a user for a service through SAML2.


# **GETV1SamlMetadata**
> SamlMetadataResult GETV1SamlMetadata()

Gets the SAML2 metadata XML.

This endpoint is exceptional for SAML which is XML-based.  Most Apis should not return XML.


### Parameters
This endpoint does not need any parameter.

### Return type

[**SamlMetadataResult**](SamlMetadataResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTV1SamlLogin**
> SamlAuthenticationResult POSTV1SamlLogin()

Authenticates a user for a service through SAML2.

This endpoint is very exceptional.  It can return a redirect response to www, or a SAML2 response.  Most Apis should not return anything except for Json.


### Parameters
This endpoint does not need any parameter.

### Return type

[**SamlAuthenticationResult**](SamlAuthenticationResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


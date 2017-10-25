/* 
 * Roblox.Api.Authentication V1
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: v1
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

// A class representing the status of a Recommended Username
type RecommendedUsernameResponseModel struct {

	// Gets or sets a value indicating whether a new username was generated
	DidGenerateNewUsername bool `json:"didGenerateNewUsername,omitempty"`

	// Gets or sets a value indicating the suggested username that is not taken
	SuggestedUsername string `json:"suggestedUsername,omitempty"`
}
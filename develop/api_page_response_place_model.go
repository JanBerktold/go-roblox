/* 
 * Roblox.Api.Develop V1
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: v1
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type ApiPageResponsePlaceModel struct {

	PreviousPageCursor string `json:"previousPageCursor,omitempty"`

	NextPageCursor string `json:"nextPageCursor,omitempty"`

	Data []PlaceModel `json:"data,omitempty"`
}
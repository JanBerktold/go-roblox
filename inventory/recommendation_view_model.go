/* 
 * Roblox.Api.Inventory V1
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: v1
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type RecommendationViewModel struct {

	Item ItemContainer `json:"Item,omitempty"`

	Creator CreatorContainer `json:"Creator,omitempty"`

	Thumbnail ThumbResultContainer `json:"Thumbnail,omitempty"`

	Product ProductContainer `json:"Product,omitempty"`
}

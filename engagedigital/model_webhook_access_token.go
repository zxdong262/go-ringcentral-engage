/*
 * Engage Digital API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package engagedigital

import (
	"time"
)

type WebhookAccessToken struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	Id        string    `json:"id"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

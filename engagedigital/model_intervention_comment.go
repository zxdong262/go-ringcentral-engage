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

type InterventionComment struct {
	Body           string    `json:"body,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	Id             string    `json:"id,omitempty"`
	InterventionId string    `json:"intervention_id,omitempty"`
	ThreadId       string    `json:"thread_id,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	UserId         string    `json:"user_id,omitempty"`
	VidentityId    string    `json:"videntity_id,omitempty"`
}

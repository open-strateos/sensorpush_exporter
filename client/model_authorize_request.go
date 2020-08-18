/*
 * SensorPush Public API
 *
 * This is a swagger definition for the SensorPush public REST API. Download the definition file [here](https://api.sensorpush.com/api/v1/support/swagger/swagger-v1.json).
 *
 * API version: v1.0.20200621
 * Contact: support@sensorpush.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// AuthorizeRequest Request object for an oAuth authorization code.
type AuthorizeRequest struct {
	// Email associated with a valid account.
	Email string `json:"email,omitempty"`
	// Password associated with the email.
	Password string `json:"password,omitempty"`
}
/*
 * SensorPush Public API
 *
 * This is a swagger definition for the SensorPush public REST API. Download the definition file [here](https://api.sensorpush.com/api/v1/support/swagger/swagger-v1.json).
 *
 * API version: v1.0.20200621
 * Contact: support@sensorpush.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package sensorpush

import "time"

// Samples Map of registered SensorPush sensors
type Samples struct {
	// Timestamp of the last sample returned. Use this as the start_ts to query for the next page of samples.
	LastTime time.Time `json:"last_time,omitempty"`
	// Map of sensors and the associated samples.
	Sensors map[string][]Sample `json:"sensors,omitempty"`
	// Message describing state of the api call.
	Status string `json:"status,omitempty"`
	// Total number of samples across all sensors
	TotalSamples float32 `json:"total_samples,omitempty"`
	// Total number of sensors returned
	TotalSensors float32 `json:"total_sensors,omitempty"`
	// The query returned too many results, causing the sample list to be truncated. Consider adjusting the limit or startTime parameters.
	Truncated bool `json:"truncated,omitempty"`
}

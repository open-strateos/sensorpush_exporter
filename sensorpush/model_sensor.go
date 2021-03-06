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
// Sensor struct for Sensor
type Sensor struct {
	// Is the sensor active?
	Active bool `json:"active,omitempty"`
	// MAC address
	Address string `json:"address,omitempty"`
	Alerts SensorAlerts `json:"alerts,omitempty"`
	// Current battery voltage
	BatteryVoltage float32 `json:"battery_voltage,omitempty"`
	Calibration SensorCalibration `json:"calibration,omitempty"`
	// Short device Id
	DeviceId string `json:"deviceId,omitempty"`
	// Long device Id
	Id string `json:"id,omitempty"`
	// Name of the sensor sensor
	Name string `json:"name,omitempty"`
}

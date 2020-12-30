SensorPush Prometheus Exporter
==============================

This is a tool for exporting [SensorPush](https://www.sensorpush.com/) temperature and humitidty measurements as [Prometheus](https://prometheus.io/) metrics.

SenorPush is a commercial IoT device that publishes its temeperature and humidity measurements to an hosted service that's [API-Accessible](https://www.sensorpush.com/gateway-cloud-api).  This exporter connects to the API and publishes the telemetry in Prometheus Format.

Metrics Exported
================

* `uptime_seconds`: Number of seconds the exporter instance has been running
* `reauth_count`: Number of times the exporter has re-authenticated with the SensorPush API
* `number_of_sensors`: Number of SensorPush devices being monitored
* `temperature_celcius`: Temperature reading, labeled by `device_name`
* `relative_humidity`: Relative humidty, labeled by `device_name`

Configuration
=============
The only required configuration are environment variables named `SENSORPUSH_USERNAME` and `SENSORPUSH_PASSWORD`, which you need to fill with the corresponding credentials for your sensorpush account.

The exporter will retrieve the list of all the sensor devices registered with your account, and periodically refresh that list. It will poll sensors in that list and expose Prometheus metrics on the configured TCP port (9687, by default).

There are a few runtime option flags, as follows:
```
Usage of sensorpush_exporter:
  -interval int
    	Polling interval, in seconds. (default 60)
  -name-refresh-interval int
    	How frequently to automatically refresh the sensor names table, in seconds (default 300)
  -port int
    	Port to publish metrics on. (default 9687)

```
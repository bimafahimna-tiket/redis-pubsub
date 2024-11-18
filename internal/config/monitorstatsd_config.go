package config

import "os"

type MonitorStatsDConfig struct {
	Domain      string
	Port        string
	ServiceName string
}

func initMonitorStatsDConfig() MonitorStatsDConfig {
	return MonitorStatsDConfig{
		Domain:      os.Getenv("MONITORSTATSD_DOMAIN"),
		Port:        os.Getenv("MONITORSTATSD_PORT"),
		ServiceName: os.Getenv("MONITORSTATSD_SERVICE_NAME"),
	}
}

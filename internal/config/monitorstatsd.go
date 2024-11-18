package config

import "os"

type MonitorStatsD struct {
	Domain      string
	Port        string
	ServiceName string
}

func initMonitorStatsD() MonitorStatsD {
	return MonitorStatsD{
		Domain:      os.Getenv("MONITORSTATSD_DOMAIN"),
		Port:        os.Getenv("MONITORSTATSD_PORT"),
		ServiceName: os.Getenv("MONITORSTATSD_SERVICE_NAME"),
	}
}

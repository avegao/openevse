package openevse

import (
	"time"
	"github.com/avegao/openevse_http/command/get/rtc_time"
	"github.com/avegao/openevse/command/get/version"
	"github.com/avegao/openevse/command/get/ammeter/settings"
)

func GetRtcTime(host string) (time.Time, error) {
	return getRtcTime.New().Run(host)
}

func GetVersion(host string) (firmwareVersion string, protocolVersion string, err error) {
	return getVersion.New().Run(host)
}

func GetAmmeterSettings(host string) (currentScaleFactor int, currentOffset int, err error) {
	return getAmmeterSettings.New().Run(host)
}

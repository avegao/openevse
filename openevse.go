package openevse

import (
	"time"
	"github.com/avegao/openevse/command/get/version"
	"github.com/avegao/openevse/command/get/ammeter/settings"
	"github.com/avegao/openevse/command/get/auth_lock_state"
	"github.com/avegao/openevse/command/set/rtc_time"
	"github.com/avegao/openevse/command/get/rtc_time"
)

func GetRtcTime(host string) (time.Time, error) {
	return getRtcTime.New(host).Run()
}

func GetVersion(host string) (firmwareVersion string, protocolVersion string, err error) {
	return getVersion.New(host).Run()
}

func GetAmmeterSettings(host string) (currentScaleFactor int, currentOffset int, err error) {
	return getAmmeterSettings.New(host).Run()
}

func GetAuthLockState(host string) (locked bool, err error) {
	return getAuthLockState.New(host).Run()
}

func SetRtcTime(host string, rtcTime time.Time) (err error) {
	return setRtcTime.New(host).Run(rtcTime)
}
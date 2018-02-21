package openevse

import (
	"time"
	"github.com/avegao/openevse/command/get/version"
	"github.com/avegao/openevse/command/get/ammeter/settings"
	"github.com/avegao/openevse/command/get/auth_lock_state"
	"github.com/avegao/openevse/command/set/rtc_time"
	"github.com/avegao/openevse/command/get/rtc_time"
	"github.com/avegao/openevse/command/get/charge_limit"
	"github.com/avegao/openevse/command/get/charging_current_and_voltage"
	"github.com/avegao/openevse/command/get/current_capacity_range_in_amps"
	"github.com/avegao/openevse/command/get/delay_timer"
	"github.com/avegao/openevse/command/get/energy_usage"
	"github.com/avegao/openevse/command/ev_connect_state"
	"github.com/avegao/openevse/command/get/ev_connect_state"
	"github.com/avegao/openevse/command/get/fault_counters"
	"github.com/avegao/openevse/command/get/over_temperature_thresholds"
	"github.com/avegao/openevse/command/get/settings"
	"github.com/avegao/openevse/command/get/time_limit"
	"github.com/avegao/openevse/command/get/voltmeter_settings"
)

func GetAmmeterSettings(host string) (currentScaleFactor int, currentOffset int, err error) {
	return getAmmeterSettings.New(host).Run()
}

func GetAuthLockState(host string) (locked bool, err error) {
	return getAuthLockState.New(host).Run()
}

func GetChargeLimit(host string) (kwh int, err error) {
	return getChargeLimit.New(host).Run()
}

func GetChargingCurrentAndVoltage(host string) (milliAmps int, milliVolts int, err error) {
	return getChargingCurrentAndVoltage.New(host).Run()
}

func GetCurrentCapacityRangeInAmps(host string) (minAmps int, maxAmps int, err error) {
	return getCurrentCapacityRangeInAmps.New(host).Run()
}

func GetDelayTimer(host string) (startTime string, endTime string, err error) {
	return getDelayTimer.New(host).Run()
}

func GetEnergyUsage(host string) (whInSession float32, whAccumulated float32, err error) {
	return getEnergyUsage.New(host).Run()
}

func GetEvConnectState(host string) (state evConnectState.EvConnectState, err error) {
	return getEvConnectState.New(host).Run()
}

func GetFaultCounters(host string) (gfdi int, noGround int, stuckRelay int, err error) {
	return getFaultCounters.New(host).Run()
}

func GetOverTemperatureThresholds(host string) (ambient float32, ir float32, err error) {
	return getOverTemperatureThresholds.New(host).Run()
}

func GetRtcTime(host string) (time.Time, error) {
	return getRtcTime.New(host).Run()
}

func GetSettings(host string) (amperes int, flags interface{}, err error) {
	return getSettings.New(host).Run()
}

func GetTimeLimit(host string) (limit int, err error) {
	return getTimeLimit.New(host).Run()
}

func GetVersion(host string) (firmwareVersion string, protocolVersion string, err error) {
	return getVersion.New(host).Run()
}

func GetVoltmeterSettings(host string) (calefactor int, offset int, err error) {
	return getVoltmeterSettings.New(host).Run()
}

func SetRtcTime(host string, rtcTime time.Time) (err error) {
	return setRtcTime.New(host).Run(rtcTime)
}
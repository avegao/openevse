package openevse

import (
	"time"
	"github.com/avegao/openevse_http/command/get/rtc_time"
	"fmt"
	"github.com/avegao/openevse/command/get/version"
)

const (
	host = "192.168.1.156"
)

func main() {
	rctTime, err := GetRtcTime(host)

	if err != nil {
		panic(err)
	}

	fmt.Println(rctTime)
}

func GetRtcTime(host string) (time.Time, error) {
	return getRtcTime.New().Run(host)
}

func GetVersion(host string) (firmwareVersion string, protocolVersion string, err error) {
	return getVersion.New().Run(host)
}

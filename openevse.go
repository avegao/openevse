package openevse

import (
	"time"
	"github.com/avegao/openevse_http/command/get/rtc_time"
	"fmt"
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

func GetRtcTime(host string) (rtcTime time.Time, err error) {
	return getRtcTime.New().Run(host)
}

package evConnectState

const (
	NotConnected EvConnectState = 0
	Connected    EvConnectState = 1
	Unknown      EvConnectState = 2
)

type EvConnectState int

func (s EvConnectState) String() (value string) {
	switch s {
	case NotConnected:
		value = "NOT_CONNECTED"
	case Connected:
		value = "CONNECTED"
	case Unknown:
	default:
		value = "UNKNOWN"
	}

	return
}

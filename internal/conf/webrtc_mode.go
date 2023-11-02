package conf

import (
	"encoding/json"
	"fmt"
)

// WebRTCMode is a WebRTC mode.
type WebRTCMode int

// values.
const (
	WebRTCModePassive WebRTCMode = iota
	WebRTCModeICE
)

// MarshalJSON implements json.Marshaler.
func (m WebRTCMode) MarshalJSON() ([]byte, error) {
	var out string

	switch m {
	case WebRTCModeICE:
		out = "ice"

	default:
		out = "passive"
	}

	return json.Marshal(out)
}

// UnmarshalJSON implements json.Unmarshaler.
func (m *WebRTCMode) UnmarshalJSON(b []byte) error {
	var in string
	if err := json.Unmarshal(b, &in); err != nil {
		return err
	}

	switch in {
	case "passive":
		*m = WebRTCModePassive

	case "ice":
		*m = WebRTCModeICE

	default:
		return fmt.Errorf("invalid WebRTC mode '%s'", in)
	}

	return nil
}

// UnmarshalEnv implements env.Unmarshaler.
func (m *WebRTCMode) UnmarshalEnv(_ string, v string) error {
	return m.UnmarshalJSON([]byte(`"` + v + `"`))
}

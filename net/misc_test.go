package net

import "testing"

func TestSplitHostPort(t *testing.T) {
	type addr struct {
		host string
		port int
	}
	table := map[string]addr{
		"host-name:132":  {host: "host-name", port: 132},
		"hostname:65535": {host: "hostname", port: 65535},
		"[::1]:321":      {host: "::1", port: 321},
		"::1:432":        {host: "::1", port: 432},
	}
	for input, want := range table {
		gotHost, gotPort, err := SplitHostPort(input)
		if err != nil {
			t.Errorf("SplitHostPort error: %v", err)
		}
		if gotHost != want.host || gotPort != want.port {
			t.Errorf("SplitHostPort(%#v) = (%v, %v), want (%v, %v)", input, gotHost, gotPort, want.host, want.port)
		}
	}
}

func TestSplitHostPortFail(t *testing.T) {
	// These cases should all fail to parse.
	inputs := []string{
		"host-name",
		"host-name:123abc",
	}
	for _, input := range inputs {
		_, _, err := SplitHostPort(input)
		if err == nil {
			t.Errorf("expected error from SplitHostPort(%q), but got none", input)
		}
	}
}

func TestJoinHostPort(t *testing.T) {
	type addr struct {
		host string
		port int32
	}
	table := map[string]addr{
		"host-name:132": {host: "host-name", port: 132},
		"[::1]:321":     {host: "::1", port: 321},
	}
	for want, input := range table {
		if got := JoinHostPort(input.host, input.port); got != want {
			t.Errorf("SplitHostPort(%v, %v) = %#v, want %#v", input.host, input.port, got, want)
		}
	}
}

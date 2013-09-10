package gopenexchangerates

import (
	"os"
	"testing"
)

var appid string

func init() {
	appid = os.Getenv("APP_ID")
	if appid == "" {
		panic("APP_ID env not set")
	}
}

func TestPopulate(t *testing.T) {
	xr := New(appid)
	if err := xr.Populate(); err != nil {
		t.Fatal(err)
	}

	if xr.Timestamp.IsZero() {
		t.Error("Timestamp expected to be populated, was zero.")
	}

	if xr.Base != "USD" {
		t.Error("Base currency isn't USD.")
	}

	rate, err := xr.Get("AUD")
	if err != nil {
		t.Fatal(err)
	}
	if rate == 0.0 {
		t.Error("AUD currency is zero, something went wrong")
	}

	// check we can repopulate
	if err := xr.Populate(); err != nil {
		t.Fatal(err)
	}
}

# gopenexchangerates

[Open Exchange Rates](http://openexchangerates.org) client for Go. Doesn't do much.

## Install

    go get github.com/jagregory/gopenexchangerates

## Usage

```go
import (
  oxr "github.com/jagregory/gopenexchangerates"
)

xr := oxr.New("your-app-id")
err := rx.Populate()
if err != nil {
  // err
}
v, err := rx.Get("AUD") // 1.543
```

`Populate` fetches the exchange rates from Open Exchange Rates. Call
this method whenever you want to update the cached copy of the rates.
Be kind to Open Exchange Rates and don't hammer them.

```go
// update exchange rates every 2 hours
go func() {
  for {
    rx.Populate()
    time.Sleep(2 * time.Hour)
  }
}()
```


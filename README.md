# slogf

`slogf` sets up a logger based on Go's slog package at 1.21 and enhances features with formatted print.

## Installation

```
go get -u github.com/keithshum/slogf
```

## Use

### 2 log levels

`slogf` supports 2 log levels, debug and info. Setting `debug` to `true` for debug level while `false` is info level.  
Debug level displays `DEBUG`, `INFO`, `WARN`, `ERROR`, `FATAL` logs.  
Info level displays `INFO`, `WARN`, `ERROR`, `FATAL` logs.  

### 2 log formats

`slogf` supports 2 log formats: `text` and `json`.  
```
time=2023-07-11T17:12:46.649Z level=INFO source=main.go:29 msg="Entered main."
```
```
{"time":"2023-07-11T17:05:15.924382Z","level":"INFO","source":{"function":"main.main","file":"main.go","line":29},"msg":"Entered main."}
```

### 2 log styles
  
`slogf` supports 2 log styles.  

#### Extra key value pairs

`Debug()`, `Info()`, `Warn()`, `Error()`, `Fatal()` supports extra arguments and attributes.  

Way to call: `Debug(message, key1, value1, key2, value2)` while message is string type value. For passing `err` directly, just append `.Error()` afterwards!  

`Debug("Hello world!", "Hello", "Peter Parker")`
```
{"time":"2023-07-11T17:05:15.924556Z","level":"DEBUG","source":{"function":"main.main","file":"main.go","line":32},"msg":"Hello world!", "Hello":"Peter Parker"}
```

#### One-liner print format

`Debugf()`, `Infof()`, `Warnf()`, `Errorf()`, `Fatalf()` supports print format style.

Way to call: `Debugf(format, substitue)`

`Debugf("Hello, %v!", "Peter Parker")`
```
{"time":"2023-07-11T17:05:15.924556Z","level":"DEBUG","source":{"function":"main.main","file":"main.go","line":32},"msg":"Hello, Peter Parker!"}
```

#### Complete example

```
package main

import (
    "errors"
    log "github.com/keithshum/slogf"
)

func main() {
    debug := true
    format := "text" // Can be "text" or "json"
    log.InitLogging(debug, format) // One-liner to initiate the logger.

    log.Debug("Dummy debug message", "pod", "MY-POD-1")
    log.Debugf("Dummy debug message. The pod is %v", "MY-POD-1")

    log.Info("Entered main.", "Hello", "Peter Parker!")
    log.Infof("Entered main. %v, %v", "Hello", "Peter Parker!")

    log.Warn("Fake warning.", "pod", "MY-POD-2")
    log.Warnf("Fake warning. The pod is %v", "MY-POD-2")

    err := errors.New("dial tcp: lookup __some_service__: no such host")
    log.Error(err.Error(), "pod", "MY-POD-3")
    log.Errorf("Failed: %v", err.Error())

    log.Fatal("Fatal and exit", "pod", "MY-POD-4")
    log.Fatalf("Fatal and exit, pod = %v", "MY-POD-4") // This line will never get called as program has already exited.
}
```

```
[2023-08-29 UTC 23:02:07] /tmp/slogf  $ go run main.go
time=2023-08-29T23:02:19.921Z level=DEBUG source=main.go:13 msg="Dummy debug message" pod=MY-POD-1
time=2023-08-29T23:02:19.921Z level=DEBUG source=main.go:14 msg="Dummy debug message. The pod is MY-POD-1"
time=2023-08-29T23:02:19.921Z level=INFO source=main.go:16 msg="Entered main." Hello="Peter Parker!"
time=2023-08-29T23:02:19.921Z level=INFO source=main.go:17 msg="Entered main. Hello, Peter Parker!"
time=2023-08-29T23:02:19.921Z level=WARN source=main.go:19 msg="Fake warning." pod=MY-POD-2
time=2023-08-29T23:02:19.921Z level=WARN source=main.go:20 msg="Fake warning. The pod is MY-POD-2"
time=2023-08-29T23:02:19.921Z level=ERROR source=main.go:23 msg="dial tcp: lookup __some_service__: no such host" pod=MY-POD-3
time=2023-08-29T23:02:19.921Z level=ERROR source=main.go:24 msg="Failed: dial tcp: lookup __some_service__: no such host"
time=2023-08-29T23:02:19.921Z level=FATAL source=main.go:26 msg="Fatal and exit" pod=MY-POD-4
exit status 1
```
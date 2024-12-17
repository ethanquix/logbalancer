# LogBalancer

**LogBalancer** is a powerful Go-based logging utility that enables easy redirection, filtering, and forwarding of logs. It serves as a logging server where logs can be sent and redirected to various endpoints, such as Slack, Telegram, or even custom-defined destinations.

---

## Features

✅ **Use HTTP like paths to redirect your logs to different clients** (see examples)

✅ **Support for multiple endpoints**: `/json`, `/proto`, `/connect`.

✅ **Customizable log forwarding**: Send logs to Slack, Telegram, or create your own destination.

✅ **Log filtering**: Route logs based on severity and other factors.

✅ **Flexible client support**: Use `connect-rpc` clients to send logs.

✅ **Extendable**: Create and integrate custom routes, etc....

---

## Installation

Install LogBalancer using `go get`:

```bash
go get github.com/ethanquix/logbalancer
```

You will find the proto definitions in `/gen` as well at the connect clients 
```bash
go proto path: "github.com/ethanquix/logbalancer/gen/go/pkg/model/pb_logs"
```
---

## Quick Start

Here is a basic example of using LogBalancer to forward and filter logs.

### 1. Import LogBalancer

```go
import (
	"context"
	"testing"
	"time"

	"github.com/ethanquix/logbalancer"
	"github.com/ethanquix/logbalancer/lbdestinations"
	"github.com/ethanquix/logbalancer/lbclients"
	"github.com/ethanquix/logbalancer/pb_logs"
	"github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/types/known/timestamppb"
)
```

### 2. Create a LogBalancer Instance

```go
func main() {
	// Initialize Slack Destination
	sl := lbdestinations.NewSlack("<your-slack-token>")

	// Initialize LogBalancer Server
	lb := logbalancer.New(
		logbalancer.WithPassword("password"),
		logbalancer.WithPort("8080"),
	)

	// Redirect logs to stdout or Slack
	lb.On("/", lbdestinations.StdoutSend)
	
	lb.On("/my_project/*", sl.SendTo("logs-channel"))
	// will accept any logs whose path is /my_project, /my_project/task1, ...

	// Filter logs by severity
	lb.On("/", lbdestinations.FilterBySeverity(lbdestinations.SeverityFilter{
		WARN:  sl.SendTo("warning-channel"),
		ERROR: sl.SendTo("error-channel"),
		DEBUG: lbdestinations.Join(sl.SendTo("debug-channel"), lbdestinations.StdoutSend),
	}))

	lb.Run()
}
```

### 3. Send Logs via Connect-RPC

```go
cc := lbclients.NewConnectLBClient("http://localhost:8080/connect", "password")
_, err := cc.Send(context.Background(), connect.NewRequest(&pb_logs.RuntimeLogs{
    LogDate:  timestamppb.New(time.Now()),
    Severity: pb_logs.Severity_SEVERITY_INFO,
    Source:   "project1",
    Message:  "Hello World",
    Path:     "/projects/1/",
}))
```

## Log Forwarding Destinations

LogBalancer supports multiple destinations to forward logs. You can use the provided helpers or create custom ones.

### 1. **Slack**

```go
sl := lbdestinations.NewSlack("<your-slack-token>")

// Send logs to a Slack channel
lb.On("/", sl.SendTo("channel-name"))
```

### 2. **Telegram**

```go
tg, err := lbdestinations.NewTelegram("<your-telegram-bot-token>")
require.NoError(t, err)

// Send logs to a Telegram chat
lb.On("/", tg.SendTo(<chat-id>))
```

### 3. **Stdout**

```go
lb.On("/", lbdestinations.StdoutSend)
```

### 4. **Custom Destination**

You can implement your own log destinations:

```go
func CustomLogDestination(log interface{}) {
	fmt.Println("Custom log destination:", log)
}

lb.On("/", CustomLogDestination)
```

---

## Log Filtering

LogBalancer allows you to filter logs based on their severity (e.g., DEBUG, WARN, ERROR).

Example:

```go
lb.On("/", lbdestinations.FilterBySeverity(lbdestinations.SeverityFilter{
	WARN:  sl.SendTo("warning"),
	ERROR: sl.SendTo("error"),
	DEBUG: lbdestinations.Join(sl.SendTo("debug"), lbdestinations.StdoutSend),
}))
```

---

## Endpoints

LogBalancer exposes the following endpoints:

| Endpoint   | Description                         |
|------------|-------------------------------------|
| `/json`    | Accepts logs in JSON format         |
| `/proto`   | Accepts logs in Protobuf format     |
| `/connect` | Accepts logs via connect-rpc client |

Example JSON payload for `/json`:

```json
{
  "logDate": "2024-12-17T09:09:50.466540Z",
  "severity": "SEVERITY_INFO",
  "source": "test",
  "message": "Hello from test",
  "context": {
    "client": "123"
  },
  "path": "/",
  "tags": {
    "version": "456"
  }
}
```

---

## Extending LogBalancer

LogBalancer is extendable, and you can define your own endpoints or destinations.

### Create a Custom Endpoint

```go
lb.On("/custom", func(incomingLog *pb_logs.RuntimeLog) {
	fmt.Println(incomingLog.Message)
	return nil
})
```

---

## License

This project is licensed under the MIT License.

---

## Contributing

Contributions are welcome! Feel free to fork the repository and submit a PR.

---

## Author

**Dimitri Wyzlic**  
GitHub: [ethanquix](https://github.com/ethanquix)

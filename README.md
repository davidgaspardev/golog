# Go Log

Color log system written in [Golang](https://go.dev).

## Log types

- Info (`information`)
- Io (`input/output`)
- Service
- System
- Error
- Warn (`warning`)

## Using


File: `server.go`

```go
import "github.com/davidgaspardev/golog"

func Run(port uint16) {
    golog.Info("Server", "running server at: "+port)
    // Code here
}
```

File: `email_service.go`

```go
import "github.com/davidgaspardev/golog"

var log = golog.NewLogTag("Email Service")

func SendEmail(message string) {
    log.Service("SendEmail", "sending email with this content: "+message)
    // Code here
}
```
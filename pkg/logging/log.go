package logging

import (
	"time"
)

type RedirectLog struct {
	Hash        string
	Requestid   string
	Requestdate time.Time
}

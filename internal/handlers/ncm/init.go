package ncm

import (
	"net/http"
	"time"
)

var Client *http.Client

func Init() {
	Client = &http.Client{Timeout: 10 * time.Second}
}

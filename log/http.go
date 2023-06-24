package log

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	HTTPKey = key("http")
)

type HTTP struct {
	TrackID  string       `json:"track_id"`
	Module   string       `json:"module"`
	Latency  float64      `json:"latency"`
	Level    logrus.Level `json:"-"`
	Error    string       `json:"error"`
	Request  *Request     `json:"request""`
	Response *Response    `json:"response"`
}

type Request struct {
	Host   string      `json:"host"`
	Route  string      `json:"route"`
	Header http.Header `json:"header"`
	Param  string      `json:"param"`
}

type Response struct {
	Header   http.Header `json:"header"`
	Status   int         `json:"status"`
	Body     interface{} `json:"body"`
	RemoteIP string      `json:"remote_ip"`
}

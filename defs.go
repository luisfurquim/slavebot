package slavebot


import (
   "errors"
   "github.com/luisfurquim/goose"
)

type ServiceT struct {
   // end point for monitoring
   ping    bool `method:"GET" path:"/ping" ok:"Test for service availability/operation"`

   // end point for monitoring
   stop    bool `method:"GET" path:"/stop" ok:"Ends service operation"`

   botId  string
   onStop func()
}


var config []byte
var ErrRecvConfig = errors.New("Error receiving configuration from Master Bot")
var Goose goose.Alert

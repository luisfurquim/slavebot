package slavebot


import (
   "errors"
   "github.com/luisfurquim/goose"
   "github.com/luisfurquim/stonelizard"
)

type ServiceT struct {
   // end point for monitoring
//   ping    bool `method:"GET" path:"/ping" ok:"Test for service availability/operation" ping:"test name"`
   ping    stonelizard.Void `method:"GET" path:"/ping" ok:"Test for service availability/operation" ping:"test name"`

   // end point for monitoring
   stop    stonelizard.Void `method:"GET" path:"/stop" ok:"Ends service operation"`

   botId  string
   onStop func()
}


var config []byte
var ErrRecvConfig = errors.New("Error receiving configuration from Master Bot")
var Goose goose.Alert

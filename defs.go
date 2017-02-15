package slavebot


import (
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

type SlavebotG struct {
   Stop goose.Alert
   Ping goose.Alert
}

var Goose SlavebotG

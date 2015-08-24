package slavebot


import (
   "errors"
   "github.com/luisfurquim/goose"
)


var config []byte
var ErrRecvConfig = errors.New("Error receiving configuration from Master Bot")
var Goose goose.Alert

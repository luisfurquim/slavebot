package slavebot


import (
   "net/http"
   "github.com/luisfurquim/stonelizard"
)

func Ping() stonelizard.Response {
   Goose.Logf(2,"Ping received")
   return stonelizard.Response{
      Status: http.StatusOK,
      Body: "OK",
   }

}


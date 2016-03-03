package slavebot


import (
   "net/http"
   "github.com/luisfurquim/stonelizard"
)

func (svc ServiceT) Ping() stonelizard.Response {
   Goose.Ping.Logf(2,"Ping received")
   return stonelizard.Response{
      Status: http.StatusNoContent,
   }

}


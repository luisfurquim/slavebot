package slavebot


import (
   "sync"
   "net/http"
   "github.com/luisfurquim/stonelizard"
)

func (svc ServiceT) Stop() stonelizard.Response {
   var wg sync.WaitGroup

   wg.Add(1)
   defer wg.Done()

   go (func () {
      wg.Wait()
      svc.onStop()
   })()

   Goose.Logf(2,"Stopping " + svc.botId + " bot")

   return stonelizard.Response{
      Status: http.StatusNoContent,
   }
}


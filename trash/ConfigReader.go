package slavebot


import (
   "io"
   "bytes"
   "github.com/luisfurquim/stonelizard"
)

func ConfigReader() stonelizard.Shaper {
   err = json.NewDecoder(bytes.NewReader(config)).Decode(&resp)

   if (err!=nil) && (err!=io.EOF) {
      Goose.Logf(1,"Failed parsing config file: %s", err)
      return nil, err
   }


   return
}


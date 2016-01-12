package slavebot


import (
   "io"
   "io/ioutil"
)

func ReceiveConfig(cfgReader io.Reader) error {
   var err error
   config, err = ioutil.ReadAll(cfgReader)
   if err != nil {
      Goose.Logf(1,"%s (%s)",ErrRecvConfig,err)
      return ErrRecvConfig
   }
//   os.Stdin.Close()
//   os.Stdout.Close()
//   os.Stderr.Close()
   return nil
}


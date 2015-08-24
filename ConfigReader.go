package slavebot


import (
   "io"
   "bytes"
)

func ConfigReader() io.Reader {
   return bytes.NewReader(config)
}


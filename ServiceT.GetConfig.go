package slavebot

import (
   "io"
)

func (svc ServiceT) GetConfig() (io.Reader, error) {
   return ConfigReader(), nil
}

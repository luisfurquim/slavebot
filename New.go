package slavebot

func New(botId string, onStop func()) *ServiceT {
   return &ServiceT{
      botId: botId,
      onStop: onStop,
   }
}


package main

import (
  def "config"
  "driver"
  "os"
  "os/signal"
)

func main() {
  var floor int
  var err error
  floor, error = driver.Init()
  if err != nil {
    def.Restart.Run()
    log.Fatal(err)
  }
}

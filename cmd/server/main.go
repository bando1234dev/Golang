package main

import "golang-BE/internal/routers"

func main() {
  r := routers.NewRouters()

  r.Run(":8002")
}

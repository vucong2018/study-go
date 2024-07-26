package main

import (
	"github.com/vucong2018/study-go/internal/routers"
)

func main() {
	r := routers.NewRouter()
  	r.Run()
}


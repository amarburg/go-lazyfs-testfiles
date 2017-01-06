package main

import "os"
import "os/signal"
import "syscall"
import "fmt"

import "github.com/amarburg/go-lazyfs-testfiles/http_server"


func main() {
  srv := lazyfs_testfiles.HttpServer(4567)

  stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT)

	fmt.Printf("Serving HTTP\n")
	select {
	case signal := <-stop:
		fmt.Printf("Got signal:%v\n", signal)
	}
	fmt.Printf("Stopping server\n")
  srv.Stop();
}

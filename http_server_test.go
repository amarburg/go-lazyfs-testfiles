package lazyfs_testfiles

import "os"
import "os/signal"
import "syscall"
import "fmt"

import "testing"

func TestServer( t *testing.T ) {
  srv := HttpServer(4567, ".")

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

package lazyfs_testfiles

import "net/http"
//import "net/url"
import "os"
import "io"
import "fmt"
import "sync"
import "net"
//import "log"

import "github.com/hydrogen18/stoppableListener"


type SLServer struct {
  wg  sync.WaitGroup
  sl  *stoppableListener.StoppableListener
  server *http.Server
}

func HttpServer( port int )  (*SLServer) {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    localPath := "." + r.URL.Path
    fmt.Printf("Looking for: %s\n", localPath )

    if info,err := os.Stat( localPath ); err == nil && (info.Mode() & os.ModeType)==0 {
      file,err := os.Open( localPath)
      if err == nil {
        io.Copy( w, file)
        file.Close()
        return
      }
    }
    http.Error(w, "File not found", 404 )
   } )


    srvIp := fmt.Sprintf("127.0.0.1:%d", port )
    originalListener, err := net.Listen("tcp", srvIp)
    if err != nil {
      panic(err)
    }

    sl, err := stoppableListener.New(originalListener)
    if err != nil {
      panic(err)
    }

    //var wg sync.WaitGroup
    srv := SLServer{ server: &http.Server{}, sl: sl, wg: sync.WaitGroup{} }
    srv.wg.Add(1)
    go func() {
      defer srv.wg.Done()
      srv.server.Serve(sl)
    }()

    return &srv
}

    //
    // srvUrl,_ := url.Parse( fmt.Sprintf("http://127.0.0.1:%d/", port ))
    //
    // log.Fatal(http.ListenAndServe(srvUrl.String(), nil))
    //}

func (srv *SLServer) Stop() {
  srv.sl.Stop()
  srv.wg.Wait()
}

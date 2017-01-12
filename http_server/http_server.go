package lazyfs_testfiles_http_server

import "net/http"
//import "net/url"
import "os"
import "io"
import "fmt"
import "sync"
import "net"
import "runtime"
//import "log"
import "path/filepath"

import "github.com/hydrogen18/stoppableListener"


func RepoRoot() string {
  _, file, _, _ := runtime.Caller(0)
  return filepath.Clean(file + "/../..")
}
var root string = RepoRoot()


type SLServer struct {
  wg  sync.WaitGroup
  sl  *stoppableListener.StoppableListener
  server *http.Server
  Url string
}

var once bool = true

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path == "/" { r.URL.Path = "/index.html"}

  localPath := root + r.URL.Path

  if info,err := os.Stat( localPath ); err == nil && (info.Mode() & os.ModeType)==0 {
    file,err := os.Open( localPath )
    if err == nil {

      // Need to handle Content-Range :-)
      contentRange,hasContentRange := r.Header["Range"]

      if hasContentRange {
        //fmt.Printf("Oy, has content range: %s\n", contentRange[0] )
        var start, end int
        n,_ := fmt.Sscanf( contentRange[0], "bytes=%d-%d", &start, &end )

        if n != 2 {
          http.Error(w, "Parse Error", 400 )
          return
        }

        sz := info.Size()
        if start > end { start = end }

        w.Header().Set("Content-Range",fmt.Sprintf( "%d-%d/%d", start, end, sz ) )
        //w.Header().Set("Content-Length",fmt.Sprintf( "%d", end-start ) )

        if start <= end {
          rdr := io.NewSectionReader( file, int64(start), int64(end-start) )
          io.Copy( w, rdr )
        }

      } else {
        io.Copy( w, file)
      }
      file.Close()

      // dt := time.Now().Sub( startTime )
      // fmt.Printf("Hander out %d\n", dt.Nanoseconds() )

      return
    }
  }

  http.Error(w, "File not found", 404 )
 }

type HttpConfig struct {
  host  string
  port  int
}

func HttpServer( configFuncs ...func( *HttpConfig ) )  (*SLServer) {

  config := HttpConfig{
    host: "127.0.0.1",
    port: 4567,
  }

  for _,f := range configFuncs { f( &config ) }

  if once {
    http.HandleFunc("/", HandlerFunc )
    once = false
  }


    srvIp := fmt.Sprintf("%s:%d", config.host, config.port )
    originalListener, err := net.Listen("tcp", srvIp)
    if err != nil {
      panic(err)
    }

    //fmt.Printf("Starting web server at %s\n", url)

    sl, err := stoppableListener.New(originalListener)
    if err != nil {
      panic(err)
    }

    //var wg sync.WaitGroup
    srv := SLServer{ server: &http.Server{},
                    sl: sl,
                    wg: sync.WaitGroup{},
                    Url: fmt.Sprintf("http://%s/", srvIp ) }

    srv.wg.Add(1)
    go func() {
      defer srv.wg.Done()
      srv.server.Serve(sl)
    }()

    return &srv
}


func (srv *SLServer) Stop() {
  //fmt.Printf("Stopping web server...")
  srv.sl.Stop()
  srv.wg.Wait()
  //fmt.Printf("done\n")

}

package lazyfs_testfiles

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
}

var once bool = true

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path == "/" { r.URL.Path = "/index.html"}

  localPath := root + r.URL.Path
  //fmt.Println(r.Header)

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
          w.Header().Set("Trailer", "Content-Range")

          if start <= end {
            rdr := io.NewSectionReader( file, int64(start), int64(end-start) )
            io.Copy( w, rdr )
          }

        w.Header().Set("Content-Range",fmt.Sprintf( "%d-%d/%d", start, end, sz ) )


      } else {
        io.Copy( w, file)
      }
      file.Close()
      return
    }
  } //else {
  //  fmt.Println( info, err )
  //}

  http.Error(w, "File not found", 404 )
 }

func HttpServer( port int )  (*SLServer) {


  //fmt.Println("Repo root is", root )

  if once {
  http.HandleFunc("/", HandlerFunc )
   once = false
 }


    srvIp := fmt.Sprintf("127.0.0.1:%d", port )
    url   := fmt.Sprintf("http://127.0.0.1:%d/", port )
    originalListener, err := net.Listen("tcp", srvIp)
    if err != nil {
      panic(err)
    }

    fmt.Printf("Starting web server at %s\n", url)

    sl, err := stoppableListener.New(originalListener)
    if err != nil {
      panic(err)
    }


    //var wg sync.WaitGroup
    srv := SLServer{ server: &http.Server{},
                    sl: sl,
                    wg: sync.WaitGroup{} }
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
  fmt.Printf("Stopping web server...")
  srv.sl.Stop()
  srv.wg.Wait()
  fmt.Printf("done\n")

}

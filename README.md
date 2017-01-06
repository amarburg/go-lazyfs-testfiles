# go-lazyfs-testfiles

Support files for testing go-lazyfs, go-quicktime, and go-prores-libav.   Also includes a primitive HTTP server which an be embedded as a goroutine for in situ testing w/o a dedicated webserver.

These resources can be accessed a couple of different ways: as a local file from the git checkout, by launching the local webserver (either standalone or within a Test), or from Github.

The cmd `http_server` launches a web server at port `localhost:4567`.   

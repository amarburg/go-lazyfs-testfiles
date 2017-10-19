package lazyfs_testfiles

import "runtime"
import "path"
import "path/filepath"


func RepoRoot() string {
  _, file, _, _ := runtime.Caller(0)
  return filepath.Clean(file + "/../")
}

var TenMegBinaryFile string = "ten_meg_random.dd"
var TenMegFileLength int = 10485760
var AlphabetFile string     = "alphabet.fs"

var TestMovFile = "CamHD_Vent_Short.mov"
var TestMovPath = path.Join( RepoRoot(), TestMovFile )
// Known apriori
var TestMovNumFrames = 31        // frames
var TestMovDuration = 1.034367   // seconds

var TestMovieWidth = 1920
var TestMovieHeight = 1080

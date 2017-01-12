package lazyfs_testfiles

import "runtime"
import "path/filepath"


func RepoRoot() string {
  _, file, _, _ := runtime.Caller(0)
  return filepath.Clean(file + "/../")
}

var TenMegBinaryFile string = "ten_meg_random.dd"
var TenMegFileLength int = 10485760
var AlphabetFile string     = "alphabet.fs"

var TestMovFile = "CamHD_Vent_Short.mov"

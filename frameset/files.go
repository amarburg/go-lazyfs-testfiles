package frameset_testfiles

import "runtime"
import "path"
import "path/filepath"

func RepoRoot() string {
  _, file, _, _ := runtime.Caller(0)
  return filepath.Clean(file + "/../")
}

var GoodMultiMovJson = path.Join( RepoRoot(), "good_frameset.json" )

const GoodMultiMovJsonChunks = 5
const GoodMultiMovJsonFrames = 31

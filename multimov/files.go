package multimov_testfiles

import "runtime"
import "path"
import "path/filepath"

func RepoRoot() string {
  _, file, _, _ := runtime.Caller(0)
  return filepath.Clean(file + "/../")
}

var EmptyMultiMovJson = path.Join( RepoRoot(), "empty_multimov.json" )
var ZeroLengthMultiMovJson = path.Join( RepoRoot(), "zero_length_multimov.json" )

var SingleMovMultiMovJson = path.Join( RepoRoot(), "single_mov.json" )
var FourMovMultiMovJson = path.Join( RepoRoot(), "four_mov.json" )

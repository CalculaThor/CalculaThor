package SVAnalyzer

// #cgo CFLAGS: -Wa,-mbig-obj -xc++ -lstdc++ -shared-libgcc
// #include "bridge.h"
import "C"

type RootRange struct {
	a, b, root      float64
	isRoot, isRange bool
	iterations      uint
}

type RootAnswer struct {
	root                        float64
	isRoot, isAlmostRoot, badIn bool
	iterations                  uint
}

func IncrementalSearch(x0, dx float64, maxIt uint) *RootRange {
	r := &RootRange{}
	cr := C.incrementalSearch(C.double(x0), C.double(dx), C.uint(maxIt))
	r.a = float64(cr.a)
	r.b = float64(cr.b)
	r.root = float64(cr.root)
	r.isRange = bool(cr.isRange)
	r.isRoot = bool(cr.isRoot)

	return r
}

func setF(s string) bool {
	return bool(C.setF(C.CString(s)))
}

package simplemath

import "fmt"

// func main() {
// 	sv := simplemath.NewSemanticVersion(1, 2, 3)
// 	sv.IncrementMajor()
// 	sv.IncrementMinor()
// 	sv.IncrementPatch()
// 	println(sv.String())
// }

// SemanticVersion ...
type SemanticVersion struct {
	major, minor, patch int
}

// NewSemanticVersion ...
func NewSemanticVersion(major, minor, patch int) SemanticVersion {
	return SemanticVersion{
		major: major,
		minor: minor,
		patch: patch,
	}
}

func (sv SemanticVersion) String() string { // Defining method in go.
	return fmt.Sprintf("%d,%d,%d", sv.major, sv.minor, sv.patch)
}

// IncrementMajor ...
func (sv *SemanticVersion) IncrementMajor() {
	sv.major++
}

// IncrementMinor ...
func (sv *SemanticVersion) IncrementMinor() {
	sv.minor++
}

// IncrementPatch ...
func (sv *SemanticVersion) IncrementPatch() {
	sv.patch++
}

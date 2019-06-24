//This is Lecture 35 Excerise
package goversion

import (
	"runtime"
)

func Version() string {
	return runtime.Version()
}

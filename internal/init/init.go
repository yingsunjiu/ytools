package init

import "os"

func init() {
	// All ytools operations must be under UTC.
	os.Setenv("TZ", "UTC")
}

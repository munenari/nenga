package query

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if err := Init(""); err != nil {
		panic(err)
	}
	defer Close()
	os.Exit(m.Run())
}

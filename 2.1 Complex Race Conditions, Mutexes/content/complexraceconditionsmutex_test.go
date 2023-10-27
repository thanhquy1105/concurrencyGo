package complexraceconditionsmutex

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_run(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	Run()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "$7920.00") {
		t.Error("wrong balance returned")
	} else {
		fmt.Println("ok")
	}
}

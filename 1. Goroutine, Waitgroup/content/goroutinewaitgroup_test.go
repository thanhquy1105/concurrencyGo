package goroutinewaitgroup

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)

	go updateMessage("epsilon")

	wg.Wait()

	if msg != "epsilon" {
		t.Errorf("Expected to find epsilon, but it is not there")
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "epsilon"
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, msg) {
		t.Errorf("Expected to find epsilon, but it is not there")
	}
}

func Test_run(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	Run()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("Expected to find Hello, universe!, but it is not there")
	}

	if !strings.Contains(output, "Hello, cosmo!") {
		t.Errorf("Expected to find Hello, cosmo!, but it is not there")
	}

	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("Expected to find Hello, world!, but it is not there")
	}

}

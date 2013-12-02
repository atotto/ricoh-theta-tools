package main_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"
)

const (
	test_cmd = "./cmd.exe"
)

var debug = false

func init() {
	if os.Getenv("DEBUG") == "true" {
		debug = true
	} else {
		debug = false
	}
}

func TestCmd(t *testing.T) {
	// go build -o test_cmd
	cmd := exec.Command("go", "build", "-o", test_cmd)
	run(cmd, t)

	// remove test_cmd
	defer os.Remove(test_cmd)

	// do exec
	cmd = exec.Command(test_cmd, "fixture/Ver1.10.JPG")
	out, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
	}
	// should be output json format
	if err = checkJsonFormat(out); err != nil {
		t.Fatal(err)
	}
}

func checkJsonFormat(message json.RawMessage) error {
	var obj map[string]*json.RawMessage
	err := json.Unmarshal(message, &obj)
	if err != nil {
		return err
	}
	return nil
}

func run(c *exec.Cmd, t *testing.T) {
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		t.Fatal(err)
	}
}

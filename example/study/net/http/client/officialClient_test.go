package client_test

import (
	"github.com/zhyhang/gofirst/example/study/net/http/client"
	"testing"
)

func TestRunGet(t *testing.T) {
	client.RunGet()
}

func TestGetRunPooled(t *testing.T) {
	client.GetRunPooled()
}

func TestRunGetNoZip(t *testing.T) {
	client.RunGetNoZip()
}

func TestPostPooled(t *testing.T) {
	client.PostPooled()
}

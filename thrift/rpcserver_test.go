package thrift

import (
	"context"
	"fmt"
	"testing"
)

func TestServer_Start(t *testing.T) {
	server := NewServer(true, "compact", true, false, "localhost:9090", false)
	err := server.Start(context.Background())
	fmt.Println(err)
}

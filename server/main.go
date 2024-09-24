package main

import (
	"context"
	"os"

	"github.com/sourcegraph/jsonrpc2"
	// "go.lsp.dev/jsonrpc2" // Removed unused import
	"go.lsp.dev/protocol"
)

type server struct{}

func (s *server) Initialize(ctx context.Context, params *protocol.InitializeParams) (*protocol.InitializeResult, error) {
	return &protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{
			HoverProvider: true,
		},
	}, nil
}

func (s *server) Hover(ctx context.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	return &protocol.Hover{
		Contents: protocol.MarkupContent{
			Kind:  protocol.PlainText,
			Value: "This is a hover message for AgentFlow",
		},
	}, nil
}

func main() {
	ctx := context.Background()

	// Replace NewHeaderStream with an alternative method if available
	stream := jsonrpc2.NewBufferedStream(stdrwc{}, jsonrpc2.VSCodeObjectCodec{})
	server := &server{}
	conn := jsonrpc2.NewConn(ctx, stream, server) // Update this line

	handler := protocol.ServerHandler(server, nil)

	conn.Go(ctx, handler)
	<-conn.DisconnectNotify()
}

type stdrwc struct{}

func (stdrwc) Read(p []byte) (int, error)  { return os.Stdin.Read(p) }
func (stdrwc) Write(p []byte) (int, error) { return os.Stdout.Write(p) }
func (stdrwc) Close() error                { return nil }

package factory

import (
	"github.com/DharitriOne/drt-chain-communication-go/websocket"
	"github.com/DharitriOne/drt-chain-communication-go/websocket/client"
	"github.com/DharitriOne/drt-chain-communication-go/websocket/data"
	"github.com/DharitriOne/drt-chain-communication-go/websocket/server"
	"github.com/DharitriOne/drt-chain-core-go/core"
	"github.com/DharitriOne/drt-chain-core-go/marshal"
)

// ArgsWebSocketHost holds all the arguments needed in order to create a FullDuplexHost
type ArgsWebSocketHost struct {
	WebSocketConfig data.WebSocketConfig
	Marshaller      marshal.Marshalizer
	Log             core.Logger
}

// CreateWebSocketHost will create and start a new instance of factory.FullDuplexHost
func CreateWebSocketHost(args ArgsWebSocketHost) (FullDuplexHost, error) {
	switch args.WebSocketConfig.Mode {
	case data.ModeServer:
		return createWebSocketServer(args)
	case data.ModeClient:
		return createWebSocketClient(args)
	default:
		return nil, data.ErrInvalidWebSocketHostMode
	}
}

func createWebSocketClient(args ArgsWebSocketHost) (FullDuplexHost, error) {
	payloadConverter, err := websocket.NewWebSocketPayloadConverter(args.Marshaller)
	if err != nil {
		return nil, err
	}

	return client.NewWebSocketClient(client.ArgsWebSocketClient{
		RetryDurationInSeconds:     args.WebSocketConfig.RetryDurationInSec,
		WithAcknowledge:            args.WebSocketConfig.WithAcknowledge,
		URL:                        args.WebSocketConfig.URL,
		PayloadConverter:           payloadConverter,
		Log:                        args.Log,
		BlockingAckOnError:         args.WebSocketConfig.BlockingAckOnError,
		DropMessagesIfNoConnection: args.WebSocketConfig.DropMessagesIfNoConnection,
		AckTimeoutInSeconds:        args.WebSocketConfig.AcknowledgeTimeoutInSec,
		PayloadVersion:             args.WebSocketConfig.Version,
	})
}

func createWebSocketServer(args ArgsWebSocketHost) (FullDuplexHost, error) {
	payloadConverter, err := websocket.NewWebSocketPayloadConverter(args.Marshaller)
	if err != nil {
		return nil, err
	}

	host, err := server.NewWebSocketServer(server.ArgsWebSocketServer{
		RetryDurationInSeconds:     args.WebSocketConfig.RetryDurationInSec,
		WithAcknowledge:            args.WebSocketConfig.WithAcknowledge,
		URL:                        args.WebSocketConfig.URL,
		PayloadConverter:           payloadConverter,
		Log:                        args.Log,
		BlockingAckOnError:         args.WebSocketConfig.BlockingAckOnError,
		DropMessagesIfNoConnection: args.WebSocketConfig.DropMessagesIfNoConnection,
		AckTimeoutInSeconds:        args.WebSocketConfig.AcknowledgeTimeoutInSec,
		PayloadVersion:             args.WebSocketConfig.Version,
	})
	if err != nil {
		return nil, err
	}

	return host, nil
}

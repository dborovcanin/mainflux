// Package nats contains NATS-specific message repository implementation.
package nats

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/mainflux/mainflux"
	log "github.com/mainflux/mainflux/logger"
	broker "github.com/nats-io/go-nats"
)

var _ mainflux.MessagePublisher = (*natsPublisher)(nil)

var logger log.Logger

const (
	maxFailedReqs   = 3
	maxFailureRatio = 0.6
	prefix          = "channel"
)

type natsPublisher struct {
	nc     *broker.Conn
	logger log.Logger
}

// New instantiates NATS message pubsub.
func New(nc *broker.Conn, l log.Logger) Service {
	return &natsPublisher{nc, l}
}

func (pubsub *natsPublisher) Publish(msg mainflux.RawMessage) error {
	data, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	return pubsub.nc.Publish(fmt.Sprintf("%s.%s", prefix, msg.Channel), data)
}

func (pubsub *natsPublisher) Subscribe(chanID string, channel Channel) error {
	sub, err := pubsub.nc.Subscribe(fmt.Sprintf("%s.%s", prefix, chanID), func(msg *broker.Msg) {
		if msg == nil {
			return
		}
		var rawMsg mainflux.RawMessage
		if err := proto.Unmarshal(msg.Data, &rawMsg); err != nil {
			return
		}
		channel.Messages <- rawMsg
	})

	go func() {
		<-channel.Closed
		for i := 0; i < 3; i++ {
			if err := sub.Unsubscribe(); err != nil {
				pubsub.logger.Error(fmt.Sprintf("Failed to unsubscribe from NATS %s", err.Error()))
				continue
			}
			break
		}
		channel.Close()
	}()
	return err
}

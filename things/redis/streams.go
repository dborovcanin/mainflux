// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package events

import (
	"context"

<<<<<<<< HEAD:things/clients/events/streams.go
	mfclients "github.com/mainflux/mainflux/pkg/clients"
	"github.com/mainflux/mainflux/pkg/events"
	"github.com/mainflux/mainflux/pkg/events/redis"
	"github.com/mainflux/mainflux/things/clients"
========
	"github.com/go-redis/redis/v8"
	"github.com/mainflux/mainflux"
	mfredis "github.com/mainflux/mainflux/internal/clients/redis"
	mfclients "github.com/mainflux/mainflux/pkg/clients"
	"github.com/mainflux/mainflux/things"
>>>>>>>> 9492132bb (Return Auth service):things/redis/streams.go
)

const streamID = "mainflux.things"

var _ things.Service = (*eventStore)(nil)

type eventStore struct {
<<<<<<<< HEAD:things/clients/events/streams.go
	events.Publisher
	svc clients.Service
========
	mfredis.Publisher
	svc    things.Service
	client *redis.Client
>>>>>>>> 9492132bb (Return Auth service):things/redis/streams.go
}

// NewEventStoreMiddleware returns wrapper around things service that sends
// events to event store.
<<<<<<<< HEAD:things/clients/events/streams.go
func NewEventStoreMiddleware(ctx context.Context, svc clients.Service, url string) (clients.Service, error) {
	publisher, err := redis.NewPublisher(ctx, url, streamID)
	if err != nil {
		return nil, err
========
func NewEventStoreMiddleware(ctx context.Context, svc things.Service, client *redis.Client) things.Service {
	es := &eventStore{
		svc:       svc,
		client:    client,
		Publisher: mfredis.NewEventStore(client, streamID, streamLen),
>>>>>>>> 9492132bb (Return Auth service):things/redis/streams.go
	}

	return &eventStore{
		svc:       svc,
		Publisher: publisher,
	}, nil
}

func (es *eventStore) CreateThings(ctx context.Context, token string, thing ...mfclients.Client) ([]mfclients.Client, error) {
	sths, err := es.svc.CreateThings(ctx, token, thing...)
	if err != nil {
		return sths, err
	}

	for _, th := range sths {
		event := createClientEvent{
			th,
		}
		if err := es.Publish(ctx, event); err != nil {
			return sths, err
		}
	}

	return sths, nil
}

func (es *eventStore) UpdateClient(ctx context.Context, token string, thing mfclients.Client) (mfclients.Client, error) {
	cli, err := es.svc.UpdateClient(ctx, token, thing)
	if err != nil {
		return cli, err
	}

	return es.update(ctx, "", cli)
}

func (es *eventStore) UpdateClientOwner(ctx context.Context, token string, thing mfclients.Client) (mfclients.Client, error) {
	cli, err := es.svc.UpdateClientOwner(ctx, token, thing)
	if err != nil {
		return cli, err
	}

	return es.update(ctx, "owner", cli)
}

func (es *eventStore) UpdateClientTags(ctx context.Context, token string, thing mfclients.Client) (mfclients.Client, error) {
	cli, err := es.svc.UpdateClientTags(ctx, token, thing)
	if err != nil {
		return cli, err
	}

	return es.update(ctx, "tags", cli)
}

func (es *eventStore) UpdateClientSecret(ctx context.Context, token, id, key string) (mfclients.Client, error) {
	cli, err := es.svc.UpdateClientSecret(ctx, token, id, key)
	if err != nil {
		return cli, err
	}

	return es.update(ctx, "secret", cli)
}

func (es *eventStore) update(ctx context.Context, operation string, thing mfclients.Client) (mfclients.Client, error) {
	event := updateClientEvent{
		thing, operation,
	}

	if err := es.Publish(ctx, event); err != nil {
		return thing, err
	}

	return thing, nil
}

func (es *eventStore) ViewClient(ctx context.Context, token, id string) (mfclients.Client, error) {
	cli, err := es.svc.ViewClient(ctx, token, id)
	if err != nil {
		return cli, err
	}

	event := viewClientEvent{
		cli,
	}
	if err := es.Publish(ctx, event); err != nil {
		return cli, err
	}

	return cli, nil
}

func (es *eventStore) ListClients(ctx context.Context, token string, pm mfclients.Page) (mfclients.ClientsPage, error) {
	cp, err := es.svc.ListClients(ctx, token, pm)
	if err != nil {
		return cp, err
	}
	event := listClientEvent{
		pm,
	}
	if err := es.Publish(ctx, event); err != nil {
		return cp, err
	}

	return cp, nil
}

func (es *eventStore) ListClientsByGroup(ctx context.Context, token, chID string, pm mfclients.Page) (mfclients.MembersPage, error) {
	mp, err := es.svc.ListClientsByGroup(ctx, token, chID, pm)
	if err != nil {
		return mp, err
	}
	event := listClientByGroupEvent{
		pm, chID,
	}
	if err := es.Publish(ctx, event); err != nil {
		return mp, err
	}

	return mp, nil
}

func (es *eventStore) EnableClient(ctx context.Context, token, id string) (mfclients.Client, error) {
	cli, err := es.svc.EnableClient(ctx, token, id)
	if err != nil {
		return cli, err
	}

	return es.delete(ctx, cli)
}

func (es *eventStore) DisableClient(ctx context.Context, token, id string) (mfclients.Client, error) {
	cli, err := es.svc.DisableClient(ctx, token, id)
	if err != nil {
		return cli, err
	}

	return es.delete(ctx, cli)
}

func (es *eventStore) delete(ctx context.Context, cli mfclients.Client) (mfclients.Client, error) {
	event := removeClientEvent{
		id:        cli.ID,
		updatedAt: cli.UpdatedAt,
		updatedBy: cli.UpdatedBy,
		status:    cli.Status.String(),
	}
	if err := es.Publish(ctx, event); err != nil {
		return cli, err
	}

	return cli, nil
}

func (es *eventStore) Identify(ctx context.Context, key string) (string, error) {
	thingID, err := es.svc.Identify(ctx, key)
	if err != nil {
		return thingID, err
	}
	event := identifyClientEvent{
		thingID: thingID,
	}

	if err := es.Publish(ctx, event); err != nil {
		return thingID, err
	}
	return thingID, nil
}

func (es *eventStore) Connect(ctx context.Context, token, thingID, channelID, permission string) error {
	return es.svc.Connect(ctx, token, thingID, channelID, permission)
}
func (es *eventStore) Disconnect(ctx context.Context, token, thingID, channelID, permission string) error {
	return es.svc.Disconnect(ctx, token, thingID, channelID, permission)
}

func (es *eventStore) Authorize(ctx context.Context, req *mainflux.AuthorizeReq) (string, error) {
	return es.svc.Authorize(ctx, req)
}

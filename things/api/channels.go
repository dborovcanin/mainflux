// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/mainflux/mainflux/internal/api"
	"github.com/mainflux/mainflux/internal/apiutil"
	gapi "github.com/mainflux/mainflux/internal/groups/api"
	"github.com/mainflux/mainflux/logger"
	"github.com/mainflux/mainflux/pkg/groups"
	"github.com/mainflux/mainflux/things"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func groupsHandler(svc groups.Service, tscv things.Service, r *chi.Mux, logger logger.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(apiutil.LoggingErrorEncoder(logger, api.EncodeError)),
	}
	r.Route("/channels", func(r chi.Router) {
		r.Post("/", otelhttp.NewHandler(kithttp.NewServer(
			gapi.CreateGroupEndpoint(svc),
			gapi.DecodeGroupCreate,
			api.EncodeResponse,
			opts...,
		), "create_channel").ServeHTTP)

		r.Get("/{groupID}", otelhttp.NewHandler(kithttp.NewServer(
			gapi.ViewGroupEndpoint(svc),
			gapi.DecodeGroupRequest,
			api.EncodeResponse,
			opts...,
		), "view_channel").ServeHTTP)

		r.Put("/{groupID}", otelhttp.NewHandler(kithttp.NewServer(
			gapi.UpdateGroupEndpoint(svc),
			gapi.DecodeGroupUpdate,
			api.EncodeResponse,
			opts...,
		), "update_channel").ServeHTTP)

		r.Get("/{groupID}/things", otelhttp.NewHandler(kithttp.NewServer(
			listMembersEndpoint(tscv),
			gapi.DecodeListMembershipRequest,
			api.EncodeResponse,
			opts...,
		), "list_things_by_channel").ServeHTTP)

		r.Get("/", otelhttp.NewHandler(kithttp.NewServer(
			gapi.ListGroupsEndpoint(svc, "things"),
			gapi.DecodeListGroupsRequest,
			api.EncodeResponse,
			opts...,
		), "list_channels").ServeHTTP)

		r.Post("/{groupID}/enable", otelhttp.NewHandler(kithttp.NewServer(
			gapi.EnableGroupEndpoint(svc),
			gapi.DecodeChangeGroupStatus,
			api.EncodeResponse,
			opts...,
		), "enable_channel").ServeHTTP)

		r.Post("/{groupID}/disable", otelhttp.NewHandler(kithttp.NewServer(
			gapi.DisableGroupEndpoint(svc),
			gapi.DecodeChangeGroupStatus,
			api.EncodeResponse,
			opts...,
		), "disable_channel").ServeHTTP)
	})

	r.Get("/things/{memberID}/channels", otelhttp.NewHandler(kithttp.NewServer(
		gapi.ListGroupsEndpoint(svc, "things"),
		gapi.DecodeListGroupsRequest,
		api.EncodeResponse,
		opts...,
	), "list_channel_by_things").ServeHTTP)

	return r
}

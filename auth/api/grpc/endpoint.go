// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package grpc

import (
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/mainflux/mainflux/auth"
	context "golang.org/x/net/context"
)

func issueEndpoint(svc auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(issueReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		now := time.Now().UTC()
		key := auth.Key{
			Type:     req.keyType,
			IssuedAt: now,
		}

		k, err := svc.Issue(ctx, req.issuer, key)
		if err != nil {
			return identityRes{}, err
		}
		return identityRes{k.Secret, nil}, nil
	}
}

func identifyEndpoint(svc auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(identityReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		id, err := svc.Identify(ctx, req.token, req.kind)
		if err != nil {
			return identityRes{}, err
		}
		return identityRes{id, nil}, nil
	}
}
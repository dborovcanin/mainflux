// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package tracing

import (
	"context"

	"github.com/mainflux/mainflux/things/policies"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var _ policies.Service = (*tracingMiddleware)(nil)

type tracingMiddleware struct {
	tracer trace.Tracer
	psvc   policies.Service
}

// New returns a new group service with tracing capabilities.
func New(psvc policies.Service, tracer trace.Tracer) policies.Service {
	return &tracingMiddleware{tracer, psvc}
}

// Authorize traces the "Authorize" operation of the wrapped policies.Service.
func (tm *tracingMiddleware) Authorize(ctx context.Context, ar policies.AccessRequest) (policies.Policy, error) {
	ctx, span := tm.tracer.Start(ctx, "svc_authorize", trace.WithAttributes(
		attribute.String("subject", ar.Subject),
		attribute.String("object", ar.Object),
		attribute.String("action", ar.Action),
		attribute.String("entity", ar.Entity),
	))
	defer span.End()

	return tm.psvc.Authorize(ctx, ar)
}

// AddPolicy traces the "AddPolicy" operation of the wrapped policies.Service.
func (tm *tracingMiddleware) AddPolicy(ctx context.Context, token, client string, p policies.Policy) (policies.Policy, error) {
	ctx, span := tm.tracer.Start(ctx, "svc_connect", trace.WithAttributes(
		attribute.String("client", client),
		attribute.String("subject", p.Subject),
		attribute.String("object", p.Object),
		attribute.StringSlice("actions", p.Actions),
	))
	defer span.End()

	return tm.psvc.AddPolicy(ctx, token, client, p)
}

// UpdatePolicy traces the "UpdatePolicy" operation of the wrapped policies.Service.
func (tm *tracingMiddleware) UpdatePolicy(ctx context.Context, token string, p policies.Policy) (policies.Policy, error) {
	ctx, span := tm.tracer.Start(ctx, "svc_update_policy", trace.WithAttributes(
		attribute.String("subject", p.Subject),
		attribute.String("object", p.Object),
		attribute.StringSlice("actions", p.Actions),
	))
	defer span.End()

	return tm.psvc.UpdatePolicy(ctx, token, p)
}

// ListPolicies traces the "ListPolicies" operation of the wrapped policies.Service.
func (tm *tracingMiddleware) ListPolicies(ctx context.Context, token string, p policies.Page) (policies.PolicyPage, error) {
	ctx, span := tm.tracer.Start(ctx, "svc_list_policies")
	defer span.End()

	return tm.psvc.ListPolicies(ctx, token, p)
}

// DeletePolicy traces the "DeletePolicy" operation of the wrapped policies.Service.
func (tm *tracingMiddleware) DeletePolicy(ctx context.Context, token string, p policies.Policy) error {
	ctx, span := tm.tracer.Start(ctx, "svc_disconnect", trace.WithAttributes(
		attribute.String("subject", p.Subject),
		attribute.String("object", p.Object),
	))
	defer span.End()

	return tm.psvc.DeletePolicy(ctx, token, p)
}

// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package groups_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/mainflux/mainflux/auth"
	httpapi "github.com/mainflux/mainflux/auth/api/http"
	"github.com/mainflux/mainflux/auth/jwt"
	"github.com/mainflux/mainflux/auth/mocks"
	"github.com/mainflux/mainflux/internal/apiutil"
	"github.com/mainflux/mainflux/logger"
	"github.com/mainflux/mainflux/pkg/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	contentType     = "application/json"
	email           = "user@example.com"
	secret          = "secret"
	id              = "testID"
	loginDuration   = 30 * time.Minute
	refreshDuration = 24 * time.Hour
)

type testRequest struct {
	client      *http.Client
	method      string
	url         string
	contentType string
	token       string
	body        io.Reader
}

func (tr testRequest) make() (*http.Response, error) {
	req, err := http.NewRequest(tr.method, tr.url, tr.body)
	if err != nil {
		return nil, err
	}
	if tr.token != "" {
		req.Header.Set("Authorization", apiutil.BearerPrefix+tr.token)
	}
	if tr.contentType != "" {
		req.Header.Set("Content-Type", tr.contentType)
	}
	return tr.client.Do(req)
}

func newService() auth.Service {
	krepo := new(mocks.Keys)
	grepo := new(mocks.Repository)
	prepo := new(mocks.PolicyAgent)
	idProvider := uuid.NewMock()

	t := jwt.New([]byte(secret))

	return auth.New(krepo, grepo, idProvider, t, prepo, loginDuration, refreshDuration)
}

func newServer(svc auth.Service) *httptest.Server {
	logger := logger.NewMock()
	mux := httpapi.MakeHandler(svc, logger, "")
	return httptest.NewServer(mux)
}

func toJSON(data interface{}) string {
	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}

func TestShareGroupAccess(t *testing.T) {
	svc := newService()
	ts := newServer(svc)
	defer ts.Close()

	token, err := svc.Issue(context.Background(), "", auth.Key{Type: auth.AccessKey, IssuedAt: time.Now(), SubjectID: id, Subject: email})
	assert.Nil(t, err, fmt.Sprintf("Issuing login key expected to succeed: %s", err))

	key := auth.Key{
		ID:        "id",
		Type:      auth.APIKey,
		SubjectID: id,
		Subject:   email,
		IssuedAt:  time.Now(),
	}

	apiToken, err := svc.Issue(context.Background(), token.AccessToken, key)
	assert.Nil(t, err, fmt.Sprintf("Issuing user's key expected to succeed: %s", err))

	type shareGroupAccessReq struct {
		token        string
		userGroupID  string
		ThingGroupID string `json:"thing_group_id"`
	}
	data := shareGroupAccessReq{token: apiToken.AccessToken, userGroupID: "ug", ThingGroupID: "tg"}
	invalidData := shareGroupAccessReq{token: apiToken.AccessToken, userGroupID: "ug", ThingGroupID: ""}

	cases := []struct {
		desc        string
		req         string
		contentType string
		auth        string
		userGroupID string
		status      int
	}{
		{
			desc:        "share a user group with thing group",
			req:         toJSON(data),
			contentType: contentType,
			auth:        apiToken.AccessToken,
			userGroupID: "ug",
			status:      http.StatusOK,
		},
		{
			desc:        "share a user group with invalid thing group",
			req:         toJSON(invalidData),
			contentType: contentType,
			auth:        apiToken.AccessToken,
			userGroupID: "ug",
			status:      http.StatusBadRequest,
		},
		{
			desc:        "share an invalid user group with thing group",
			req:         toJSON(data),
			contentType: contentType,
			auth:        apiToken.AccessToken,
			userGroupID: "",
			status:      http.StatusBadRequest,
		},
		{
			desc:        "share an invalid user group with invalid thing group",
			req:         toJSON(invalidData),
			contentType: contentType,
			auth:        apiToken.AccessToken,
			userGroupID: "",
			status:      http.StatusBadRequest,
		},
		{
			desc:        "share a user group with thing group with invalid content type",
			req:         toJSON(data),
			contentType: "",
			auth:        apiToken.AccessToken,
			userGroupID: "ug",
			status:      http.StatusUnsupportedMediaType,
		},
		{
			desc:        "share a user group with thing group with invalid token",
			req:         toJSON(data),
			contentType: contentType,
			auth:        "token",
			userGroupID: "ug",
			status:      http.StatusUnauthorized,
		},
	}

	for _, tc := range cases {
		req := testRequest{
			client:      ts.Client(),
			method:      http.MethodPost,
			url:         fmt.Sprintf("%s/groups/%s/share", ts.URL, tc.userGroupID),
			contentType: tc.contentType,
			token:       tc.auth,
			body:        strings.NewReader(tc.req),
		}
		res, err := req.make()
		assert.Nil(t, err, fmt.Sprintf("%s: unexpected error %s", tc.desc, err))
		assert.Equal(t, tc.status, res.StatusCode, fmt.Sprintf("%s: expected status code %d got %d", tc.desc, tc.status, res.StatusCode))
	}
}

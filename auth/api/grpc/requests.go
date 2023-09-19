// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package grpc

import (
	"github.com/mainflux/mainflux/auth"
	"github.com/mainflux/mainflux/internal/apiutil"
)

type identityReq struct {
	token string
	kind  uint32
}

func (req identityReq) validate() error {
	if req.token == "" {
		return apiutil.ErrBearerToken
	}
	if req.kind != auth.LoginKey &&
		req.kind != auth.APIKey &&
		req.kind != auth.RecoveryKey {
		return apiutil.ErrInvalidAuthKey
	}

	return nil
}

type issueReq struct {
	id      string
	email   string
	keyType uint32
}

func (req issueReq) validate() error {
	if req.email == "" {
		return apiutil.ErrMissingEmail
	}
	if req.keyType != auth.LoginKey &&
		req.keyType != auth.APIKey &&
		req.keyType != auth.RecoveryKey {
		return apiutil.ErrInvalidAuthKey
	}

	return nil
}

type assignReq struct {
	token     string
	groupID   string
	memberID  string
	groupType string
}

func (req assignReq) validate() error {
	if req.token == "" {
		return apiutil.ErrBearerToken
	}
	if req.groupID == "" || req.memberID == "" {
		return apiutil.ErrMissingID
	}
	return nil
}

type membersReq struct {
	token      string
	groupID    string
	offset     uint64
	limit      uint64
	memberType string
}

func (req membersReq) validate() error {
	if req.token == "" {
		return apiutil.ErrBearerToken
	}
	if req.groupID == "" {
		return apiutil.ErrMissingID
	}
	if req.memberType == "" {
		return apiutil.ErrMissingMemberType
	}
	return nil
}

// authReq represents authorization request. It contains:
// 1. subject - an action invoker
// 2. object - an entity over which action will be executed
// 3. action - type of action that will be executed (read/write)
type authReq struct {
	Namespace   string
	SubjectType string
	Subject     string
	Relation    string
	Permission  string
	ObjectType  string
	Object      string
}

func (req authReq) validate() error {

	if req.Subject == "" {
		return apiutil.ErrMissingPolicySub
	}

	if req.Object == "" {
		return apiutil.ErrMissingPolicyObj
	}

	if req.Permission == "" {
		return apiutil.ErrMissingPolicyAct
	}

	return nil
}

type policyReq struct {
	Namespace   string
	SubjectType string
	Subject     string
	Relation    string
	Permission  string
	ObjectType  string
	Object      string
}

func (req policyReq) validate() error {
	if req.Subject == "" {
		return apiutil.ErrMissingPolicySub
	}

	if req.Object == "" {
		return apiutil.ErrMissingPolicyObj
	}

	if req.Relation == "" {
		return apiutil.ErrMissingPolicyAct
	}

	return nil
}

type listObjectsReq struct {
	Namespace     string
	SubjectType   string
	Subject       string
	Relation      string
	Permission    string
	ObjectType    string
	Object        string
	NextPageToken string
	Limit         int32
}

type countObjectsReq struct {
	Namespace     string
	SubjectType   string
	Subject       string
	Relation      string
	Permission    string
	ObjectType    string
	Object        string
	NextPageToken string
}

type listSubjectsReq struct {
	Namespace     string
	SubjectType   string
	Subject       string
	Relation      string
	Permission    string
	ObjectType    string
	Object        string
	NextPageToken string
	Limit         int32
}

type countSubjectsReq struct {
	Namespace     string
	SubjectType   string
	Subject       string
	Relation      string
	Permission    string
	ObjectType    string
	Object        string
	NextPageToken string
}

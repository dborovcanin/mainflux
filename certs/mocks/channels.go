// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package mocks

// var _ mfgroups.Service = (*mainfluxChannels)(nil)

// type mainfluxChannels struct {
// 	mu       sync.Mutex
// 	counter  uint64
// 	channels map[string]mfgroups.Group
// 	auth     upolicies.AuthServiceClient
// }

// // NewChannelsService returns Mainflux Channels service mock.
// // Only methods used by SDK are mocked.
// func NewChannelsService(channels map[string]mfgroups.Group, auth upolicies.AuthServiceClient) mfgroups.Service {
// 	return &mainfluxChannels{
// 		channels: channels,
// 		auth:     auth,
// 	}
// }

// func (svc *mainfluxChannels) CreateGroups(ctx context.Context, token string, chs ...mfgroups.Group) ([]mfgroups.Group, error) {
// 	svc.mu.Lock()
// 	defer svc.mu.Unlock()

// 	userID, err := svc.auth.Identify(ctx, &upolicies.IdentifyReq{Token: token})
// 	if err != nil {
// 		return []mfgroups.Group{}, errors.ErrAuthentication
// 	}
// 	for i := range chs {
// 		svc.counter++
// 		chs[i].Owner = userID.GetId()
// 		chs[i].ID = strconv.FormatUint(svc.counter, 10)
// 		svc.channels[chs[i].ID] = chs[i]
// 	}

// 	return chs, nil
// }

// func (svc *mainfluxChannels) ViewGroup(_ context.Context, owner, id string) (mfgroups.Group, error) {
// 	if c, ok := svc.channels[id]; ok {
// 		return c, nil
// 	}
// 	return mfgroups.Group{}, errors.ErrNotFound
// }

// func (svc *mainfluxChannels) ListGroups(context.Context, string, mfgroups.Page) (mfgroups.Page, error) {
// 	panic("not implemented")
// }

// func (svc *mainfluxChannels) ListMemberships(context.Context, string, string, mfgroups.Page) (mfgroups.Memberships, error) {
// 	panic("not implemented")
// }

// func (svc *mainfluxChannels) UpdateGroup(context.Context, string, mfgroups.Group) (mfgroups.Group, error) {
// 	panic("not implemented")
// }

// func (svc *mainfluxChannels) EnableGroup(ctx context.Context, token, id string) (mfgroups.Group, error) {
// 	svc.mu.Lock()
// 	defer svc.mu.Unlock()

// 	userID, err := svc.auth.Identify(ctx, &upolicies.IdentifyReq{Token: token})
// 	if err != nil {
// 		return mfgroups.Group{}, errors.ErrAuthentication
// 	}

// 	if t, ok := svc.channels[id]; !ok || t.Owner != userID.GetId() {
// 		return mfgroups.Group{}, errors.ErrNotFound
// 	}
// 	if t, ok := svc.channels[id]; ok && t.Owner == userID.GetId() {
// 		t.Status = mfclients.EnabledStatus
// 		return t, nil
// 	}
// 	return mfgroups.Group{}, nil
// }

// func (svc *mainfluxChannels) DisableGroup(ctx context.Context, token, id string) (mfgroups.Group, error) {
// 	svc.mu.Lock()
// 	defer svc.mu.Unlock()

// 	userID, err := svc.auth.Identify(ctx, &upolicies.IdentifyReq{Token: token})
// 	if err != nil {
// 		return mfgroups.Group{}, errors.ErrAuthentication
// 	}

// 	if t, ok := svc.channels[id]; !ok || t.Owner != userID.GetId() {
// 		return mfgroups.Group{}, errors.ErrNotFound
// 	}
// 	if t, ok := svc.channels[id]; ok && t.Owner == userID.GetId() {
// 		t.Status = mfclients.DisabledStatus
// 		return t, nil
// 	}
// 	return mfgroups.Group{}, nil
// }

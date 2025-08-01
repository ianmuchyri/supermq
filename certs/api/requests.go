// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"time"

	apiutil "github.com/absmach/supermq/api/http/util"
	"github.com/absmach/supermq/certs"
)

const maxLimitSize = 100

type addCertsReq struct {
	token    string
	domainID string
	ClientID string `json:"client_id"`
	TTL      string `json:"ttl"`
}

func (req addCertsReq) validate() error {
	if req.token == "" {
		return apiutil.ErrBearerToken
	}

	if req.domainID == "" {
		return apiutil.ErrMissingDomainID
	}

	if req.ClientID == "" {
		return apiutil.ErrMissingID
	}

	if req.TTL == "" {
		return apiutil.ErrMissingCertData
	}

	if _, err := time.ParseDuration(req.TTL); err != nil {
		return apiutil.ErrInvalidCertData
	}

	return nil
}

type listReq struct {
	clientID string
	pm       certs.PageMetadata
}

func (req *listReq) validate() error {
	if req.pm.Limit > maxLimitSize {
		return apiutil.ErrLimitSize
	}

	return nil
}

type viewReq struct {
	serialID string
}

func (req *viewReq) validate() error {
	if req.serialID == "" {
		return apiutil.ErrMissingID
	}

	return nil
}

type revokeAllReq struct {
	token    string
	clientID string
	domainID string
}

func (req *revokeAllReq) validate() error {
	if req.token == "" {
		return apiutil.ErrBearerToken
	}

	if req.domainID == "" {
		return apiutil.ErrMissingDomainID
	}

	if req.clientID == "" {
		return apiutil.ErrMissingID
	}

	return nil
}

type revokeBySerialReq struct {
	serialID string
}

func (req *revokeBySerialReq) validate() error {
	if req.serialID == "" {
		return apiutil.ErrMissingID
	}

	return nil
}

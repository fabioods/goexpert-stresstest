package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/fabioods/goexpert-stresstest-cli-challenge/internal/pkg/httpclient"
)

type HttpClientMock struct {
	mock.Mock
}

func (m *HttpClientMock) Get(endpoint string) *httpclient.HttpClientResponse {
	ret := m.Called(endpoint)
	var r *httpclient.HttpClientResponse

	if ret.Get(0) != nil {
		r = ret.Get(0).(*httpclient.HttpClientResponse)
	}

	return r
}

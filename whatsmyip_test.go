package whatsmyip_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/darkraiden/whatsmyip"
	"github.com/darkraiden/whatsmyip/mocks"
	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetBaseURL(t *testing.T) {
	t.Run("test the GetBaseURL function and its content", func(t *testing.T) {
		expected := "http://ifconfig.me"
		actual := whatsmyip.GetBaseURL()

		require.NotEmpty(t, actual)
		assert.Equal(t, expected, actual)
	})
}

func TestGet(t *testing.T) {
	t.Run("test Get() with errors", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockDoer := mocks.NewMockDoer(ctrl)
		req, err := http.NewRequest(http.MethodGet, whatsmyip.GetBaseURL(), nil)
		require.Nil(t, err)
		require.NotNil(t, req)
		mockDoer.EXPECT().Do(req).Return(nil, errors.New("dummy error"))

		ip, err := whatsmyip.Get(mockDoer)
		assert.Empty(t, ip)
		assert.Error(t, err)
	})
}

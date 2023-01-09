package aozorarodokuweb

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func Test_GetAppMeta(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", AppMetaJsonURL,
		httpmock.NewStringResponder(200, `{
  "version":  "1.7.4",
  "url": "https://aozoraroudoku.jp/app/ar_app_list_20230101.csv"
}`),
	)
	meta, err := GetAppMeta()
	assert.NoError(t, err)
	assert.Equal(t, AppMeta{"1.7.4", "https://aozoraroudoku.jp/app/ar_app_list_20230101.csv"}, meta)
}
func Test_GetAppMeta_Error(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", AppMetaJsonURL,
		httpmock.NewStringResponder(404, ``),
	)
	meta, err := GetAppMeta()
	assert.Error(t, err)
	assert.Empty(t, meta)
}

func Test_GetContentsCSV_NotFound(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://google.com",
		httpmock.NewStringResponder(404, "NotFound"),
	)
	res, err := GetContentsCSV("https://google.com")
	assert.Error(t, err)
	assert.Empty(t, res)
}
func Test_GetContentsCSV(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://google.com",
		httpmock.NewStringResponder(200, ContentsFixture),
	)
	bytes, err := GetContentsCSV("https://google.com")
	assert.NoError(t, err)
	assert.Equal(t, ContentsFixture, string(bytes))
}

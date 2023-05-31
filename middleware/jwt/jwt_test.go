package jwt

import (
	"testing"

	"github.com/kataras/iris/v12/httptest"
	"github.com/weiheguang/iris_rest_framework"
)

func TestListView(t *testing.T) {
	app := iris_rest_framework.NewIrisApp("test_settings")

	e := httptest.New(t, app)
}

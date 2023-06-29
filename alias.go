package iris_rest_framework

import (
	"github.com/weiheguang/iris_rest_framework/alias"
	"github.com/weiheguang/iris_rest_framework/views"
)

type (
	// ListAPIViewConf 列表视图配置
	ListAPIViewConf = views.ListAPIViewConf
	List            = views.ListAPIView

	// RetrieveAPIViewConf 详情视图配置
	RetrieveAPIViewConf = views.RetrieveAPIViewConf
	Retrieve            = views.RetrieveAPIView

	Map = alias.Map
)

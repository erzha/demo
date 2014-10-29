package mwiki

import (
	"github.com/erzha/kernel"
	"code.google.com/p/go.net/context"
)

func appServerInit(ctx context.Context, pServer *kernel.Server) error {
	return nil
}

func newAppHook() (interface{}, error) {
	return nil, nil
}

func init() {
	info := kernel.PluginInfo{}
	info.Creater = newAppHook
	info.ServerInit = appServerInit
	kernel.RegisterPlugin("_app", info)
}


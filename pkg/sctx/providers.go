package sctx

import "context"

type DefaultContextProviderFunc func() context.Context

func DefaultContextProvider(mainCtx context.Context) func() context.Context {
	return func() context.Context {
		ctx, _ := context.WithCancel(mainCtx)
		return ctx
	}
}

package rest

import (
	"github.com/gorilla/mux"

	"github.com/mydexchain/chain-sdk/client/rest"

	"github.com/mydexchain/chain-sdk/client"
)

func RegisterHandlers(clientCtx client.Context, rtr *mux.Router) {
	r := rest.WithHTTPDeprecationHeaders(rtr)
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)
}

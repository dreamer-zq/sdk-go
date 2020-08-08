package admin

import (
	sdk "gitlab.bianjie.ai/cschain/sdk-go/common/types"
)

type AdminI interface {
	sdk.Module
	AddRoles(address string, roles []Role, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	RemoveRoles(address string, roles []Role, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	BlockAccount(address string, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	UnblockAccount(address string, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)

	QueryRoles(address string) ([]Role, sdk.Error)
	QueryBlacklist(page, limit int) ([]string, sdk.Error)
}

package esm

import (
	utils "github.com/comdex-official/comdex/types"
	assettypes "github.com/comdex-official/comdex/x/asset/types"
	"github.com/comdex-official/comdex/x/esm/keeper"
	"github.com/comdex-official/comdex/x/esm/types"
	markettypes "github.com/comdex-official/comdex/x/market/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func BeginBlocker(ctx sdk.Context, _ abci.RequestBeginBlock, k keeper.Keeper) {
	_ = utils.ApplyFuncIfNoError(ctx, func(ctx sdk.Context) error {
		apps, found := k.GetApps(ctx)
		if !found {
			return assettypes.AppIdsDoesntExist
		}
		for _, v := range apps {
			esmStatus, found := k.GetESMStatus(ctx, v.Id)
			if !found {
				return types.ErrESMParamsNotFound
			}
			if ctx.BlockTime().After(esmStatus.EndTime) && esmStatus.Status && !esmStatus.VaultRedemptionStatus {
				err := k.SetUpCollateralRedemptionForVault(ctx, esmStatus.AppId)
				if err != nil {
					return err
				}
			}
			if ctx.BlockTime().After(esmStatus.EndTime) && esmStatus.Status && !esmStatus.StableVaultRedemptionStatus {
				err := k.SetUpCollateralRedemptionForStableVault(ctx, esmStatus.AppId)
				if err != nil {
					return err
				}
			}
			esmMarket, found := k.GetESMMarketForAsset(ctx, v.Id)
			if found {
				return types.ErrMarketDataNotFound
			}
			if !esmMarket.IsPriceSet && esmStatus.Status {
				assets := k.GetAssetsForOracle(ctx)
				var markets []types.Market
				for _, a := range assets {
					price, found := k.GetPriceForAsset(ctx, a.Id)
					if !found {
						return markettypes.ErrorMarketForAssetDoesNotExist
					}
					market := types.Market{
						AssetID: a.Id,
						Rates:   price,
					}
					markets = append(markets, market)
				}
				em := types.ESMMarketPrice{
					AppId:      v.Id,
					IsPriceSet: true,
					Market:     markets,
				}
				k.SetESMMarketForAsset(ctx, em)
			}
		}
		return nil
	})
}
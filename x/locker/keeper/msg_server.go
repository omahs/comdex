package keeper

import (
	"context"
	"time"
	"strconv"

	"github.com/comdex-official/comdex/x/locker/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ types.MsgServiceServer = (*msgServer)(nil)
)

type msgServer struct {
	Keeper
}

func NewMsgServiceServer(keeper Keeper) types.MsgServiceServer {
	return &msgServer{
		Keeper: keeper,
	}
}

func (k *msgServer) MsgCreateLocker(c context.Context, msg *types.MsgCreateLockerRequest) (*types.MsgCreateLockerResponse, error) {

	ctx := sdk.UnwrapSDKContext(c)
	asset, found := k.GetAsset(ctx, msg.AssetID)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}
	app_mapping, found := k.GetApp(ctx, msg.AppMappingId)
	if !found {
		return nil, types.ErrorAppMappingDoesNotExist
	}
	//Checking if user mapping exists
	//if does then check app to asset mapping has any locker key
	//if it does throw error
	user_locker_asset_mapping, user_exists := k.GetUserLockerAssetMapping(ctx, msg.Depositor)

	if user_exists {
		_, already_exists := k.CheckUserAppToAssetMapping(ctx, user_locker_asset_mapping, asset.Id, app_mapping.Id)
		if already_exists {
			return nil, types.ErrorUserLockerAlreadyExists

		}

	}

	locker_product_asset_mapping, found := k.GetLockerProductAssetMapping(ctx, app_mapping.Id)
	if !found {
		return nil, types.ErrorAppMappingDoesNotExist
	}
	isfound := k.CheckLockerProductAssetMapping(ctx, asset.Id, locker_product_asset_mapping)
	if isfound {
		//This asset is accepted by the app
		//Create a new instance of locker

		//call Lookup table to get relevant data
		lookup_table_data, exists := k.GetLockerLookupTable(ctx, locker_product_asset_mapping.AppMappingId)
		if !exists {
			return nil, types.ErrorAppMappingDoesNotExist

		} else {

			//Transferring amount from user to module
			depositor, err := sdk.AccAddressFromBech32(msg.Depositor)
			if err != nil {
				return nil, err
			}
			if err := k.SendCoinFromAccountToModule(ctx, depositor, types.ModuleName, sdk.NewCoin(asset.Denom, msg.Amount)); err != nil {
				return nil, err
			}
			//Creating locker instance
			var userLocker types.Locker
			counter := lookup_table_data.Counter + 1
			userLocker.LockerId = app_mapping.ShortName + strconv.FormatUint(asset.Id,10)
			userLocker.Depositor = msg.Depositor
			userLocker.AssetDepositId = asset.Id
			userLocker.CreatedAt = time.Now()
			userLocker.IsLocked = false
			userLocker.NetBalance = msg.Amount
			userLocker.ReturnsAccumulated = sdk.ZeroInt()
			userLocker.AppMappingId = app_mapping.Id
			k.SetLocker(ctx, userLocker)

			//Creating user mapping data
			var user_mapping_data types.UserLockerAssetMapping
			var user_app_data types.LockerToAppMapping
			var user_asset_data types.AssetToLockerMapping

			user_asset_data.AssetId = asset.Id
			user_asset_data.LockerId = userLocker.LockerId
			user_app_data.AppMappingId = app_mapping.Id
			user_app_data.UserAssetLocker = append(user_app_data.UserAssetLocker, &user_asset_data)
			user_mapping_data.Owner = msg.Depositor
			user_mapping_data.LockerAppMapping = append(user_mapping_data.LockerAppMapping, &user_app_data)

			k.SetUserLockerAssetMapping(ctx, user_mapping_data)

			//Update LockerMapping Values

			k.UpdateTokenLockerMapping(ctx, lookup_table_data, counter, userLocker)

			// lookup_table_data.Counter=counter
			// var token_locker_mapping types.TokenToLockerMapping
			// token_locker_mapping.AssetId=asset.Id
			// token_locker_mapping.DepositedAmount=token_locker_mapping.DepositedAmount.Add(userLocker.DepositedAmount)
			// token_locker_mapping.LockerIds = append(token_locker_mapping.LockerIds,user_asset_data.LockerId )

		}

	} else {
		//Not a whitelisted asset , return errr
		return nil, types.ErrorLockerProductAssetMappingDoesNotExists

	}

	return &types.MsgCreateLockerResponse{}, nil

}

//Remove asset id from Deposit & Withdraw- redundant
func (k *msgServer) MsgDepositAsset(c context.Context, msg *types.MsgDepositAssetRequest) (*types.MsgDepositAssetResponse, error) {

	//Update Locker Data

	ctx := sdk.UnwrapSDKContext(c)
	asset, found := k.GetAsset(ctx, msg.AssetID)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}
	app_mapping, found := k.GetApp(ctx, msg.AppMappingId)
	if !found {
		return nil, types.ErrorAppMappingDoesNotExist
	}

	lockerData, found := k.GetLocker(ctx, msg.LockerID)

	if !found {
		return nil, types.ErrorLockerDoesNotExists
	}
	if lockerData.AssetDepositId != asset.Id {

		return nil, types.ErrorInvalidAssetID

	}
	if msg.Depositor != lockerData.Depositor {
		return nil, types.ErrorUnauthorized

	}
	if app_mapping.Id != lockerData.AppMappingId {
		return nil, types.ErrorAppMappingDoesNotExist
	}

	lookup_table_data, exists := k.GetLockerLookupTable(ctx, app_mapping.Id)
	if !exists {
		return nil, types.ErrorAppMappingDoesNotExist
	}
	depositor, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		return nil, err
	}
	if err := k.SendCoinFromAccountToModule(ctx, depositor, types.ModuleName, sdk.NewCoin(asset.Denom, msg.Amount)); err != nil {
		return nil, err
	}

	lockerData.NetBalance = lockerData.NetBalance.Add(msg.Amount)
	k.SetLocker(ctx, lockerData)

	//Update  Amount in Locker Mapping
	k.UpdateAmountLockerMapping(ctx, lookup_table_data, asset.Id, msg.Amount, true)

	return &types.MsgDepositAssetResponse{}, nil

}

//Remove asset id from Deposit & Withdraw-redundant
func (k *msgServer) MsgWithdrawAsset(c context.Context, msg *types.MsgWithdrawAssetRequest) (*types.MsgWithdrawAssetResponse, error) {

	ctx := sdk.UnwrapSDKContext(c)
	asset, found := k.GetAsset(ctx, msg.AssetID)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}
	app_mapping, found := k.GetApp(ctx, msg.AppMappingId)
	if !found {
		return nil, types.ErrorAppMappingDoesNotExist
	}

	lockerData, found := k.GetLocker(ctx, msg.LockerID)

	if !found {
		return nil, types.ErrorLockerDoesNotExists
	}
	if lockerData.AssetDepositId != asset.Id {

		return nil, types.ErrorInvalidAssetID

	}
	if msg.Depositor != lockerData.Depositor {
		return nil, types.ErrorUnauthorized

	}
	if app_mapping.Id != lockerData.AppMappingId {
		return nil, types.ErrorAppMappingDoesNotExist
	}

	lookup_table_data, exists := k.GetLockerLookupTable(ctx, app_mapping.Id)
	if !exists {
		return nil, types.ErrorAppMappingDoesNotExist
	}
	depositor, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		return nil, err
	}
	if lockerData.NetBalance.LT(msg.Amount) {
		return nil, types.ErrorRequestedAmountExceedsDepositAmount
	}

	lockerData.NetBalance = lockerData.NetBalance.Sub(msg.Amount)

	if err := k.SendCoinFromModuleToAccount(ctx, types.ModuleName, depositor, sdk.NewCoin(asset.Denom, msg.Amount)); err != nil {
		return nil, err
	}

	k.SetLocker(ctx, lockerData)

	//Update  Amount in Locker Mapping
	k.UpdateAmountLockerMapping(ctx, lookup_table_data, asset.Id, msg.Amount, false)

	return &types.MsgWithdrawAssetResponse{}, nil

}

func (k *msgServer) MsgAddWhiteListedAsset(c context.Context, msg *types.MsgAddWhiteListedAssetRequest) (*types.MsgAddWhiteListedAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	app_mapping, found := k.GetApp(ctx, msg.AppMappingId)
	if !found {
		return nil, types.ErrorAppMappingDoesNotExist
	}
	asset, found := k.GetAsset(ctx, msg.AssetID)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}
	locker_product_asset_mapping, found := k.GetLockerProductAssetMapping(ctx, msg.AppMappingId)

	if !found {
		//Set a new instance of Locker Product Asset  Mapping

		var locker types.LockerProductAssetMapping
		locker.AppMappingId = app_mapping.Id
		locker.AssetIds = append(locker.AssetIds, asset.Id)
		k.SetLockerProductAssetMapping(ctx, locker)

		//Also Create a LockerLookup table Instance and set it with the new asset id
		var lockerLookupdata types.LockerLookupTable
		var lockerAssetData types.TokenToLockerMapping

		lockerAssetData.AssetId = asset.Id
		lockerLookupdata.Counter = 0
		lockerLookupdata.AppMappingId = app_mapping.Id
		lockerLookupdata.Lockers = append(lockerLookupdata.Lockers, &lockerAssetData)
		k.SetLockerLookupTable(ctx, lockerLookupdata)

		return &types.MsgAddWhiteListedAssetResponse{}, nil

	} else {

		// Check if the asset from msg exists or not ,
		found := k.CheckLockerProductAssetMapping(ctx, msg.AssetID, locker_product_asset_mapping)

		if found {

			return nil, types.ErrorLockerProductAssetMappingExists
		}

		// Since it does not exists , push the asset id to the LockerProductAssetMapping

		locker_product_asset_mapping.AssetIds = append(locker_product_asset_mapping.AssetIds, asset.Id)
		k.SetLockerProductAssetMapping(ctx, locker_product_asset_mapping)

		//append  the asset in LockerLookup table and set it

		lockerLookupTableData, _ := k.GetLockerLookupTable(ctx, app_mapping.Id)
		var lockerAssetData types.TokenToLockerMapping

		lockerAssetData.AssetId = asset.Id
		lockerLookupTableData.Lockers = append(lockerLookupTableData.Lockers, &lockerAssetData)
		k.SetLockerLookupTable(ctx, lockerLookupTableData)

		return &types.MsgAddWhiteListedAssetResponse{}, nil

	}

}
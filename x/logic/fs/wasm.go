package fs

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/fs"
	"net/url"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/axone-protocol/axoned/v8/x/logic/types"
)

const (
	queryKey        = "query"
	base64DecodeKey = "base64Decode"
	scheme          = "cosmwasm"
)

type WasmHandler struct {
	wasmKeeper types.WasmKeeper
}

var _ URIHandler = (*WasmHandler)(nil)

func NewWasmHandler(keeper types.WasmKeeper) WasmHandler {
	return WasmHandler{wasmKeeper: keeper}
}

func (w WasmHandler) Scheme() string {
	return scheme
}

func (w WasmHandler) Open(ctx context.Context, uri *url.URL) (fs.File, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	if uri.Scheme != scheme {
		return nil, fmt.Errorf("invalid scheme")
	}

	paths := strings.SplitAfter(uri.Opaque, ":")
	pathsLen := len(paths)
	if pathsLen < 1 || paths[pathsLen-1] == "" {
		return nil, fmt.Errorf("emtpy path given, should be '%s:{contractName}:{contractAddr}?query={query}'",
			scheme)
	}

	contractAddr, err := sdk.AccAddressFromBech32(paths[pathsLen-1])
	if err != nil {
		return nil, fmt.Errorf("failed convert path '%s' to contract address: %w", paths[pathsLen-1], err)
	}

	if !uri.Query().Has(queryKey) {
		return nil, fmt.Errorf("uri should contains `query` params")
	}
	query := uri.Query().Get(queryKey)

	base64Decode := true
	if uri.Query().Has(base64DecodeKey) {
		if base64Decode, err = strconv.ParseBool(uri.Query().Get(base64DecodeKey)); err != nil {
			return nil, fmt.Errorf("failed convert 'base64Decode' query value to boolean: %w", err)
		}
	}

	data, err := w.wasmKeeper.QuerySmart(sdkCtx, contractAddr, []byte(query))
	if err != nil {
		return nil, fmt.Errorf("failed query wasm keeper: %w", err)
	}

	if base64Decode {
		var program string
		err = json.Unmarshal(data, &program)
		if err != nil {
			return nil, fmt.Errorf("failed unmarshal json wasm response to string: %w", err)
		}

		decoded, err := base64.StdEncoding.DecodeString(program)
		if err != nil {
			return nil, fmt.Errorf("failed decode wasm base64 respone: %w", err)
		}

		return NewVirtualFile(decoded, uri, sdkCtx.BlockTime()), nil
	}

	return NewVirtualFile(data, uri, sdkCtx.BlockTime()), nil
}

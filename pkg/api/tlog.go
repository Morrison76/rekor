//
// Copyright 2021 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"context"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/google/trillian/types"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"

	"github.com/Morrison76/rekor/pkg/generated/models"
	"github.com/Morrison76/rekor/pkg/generated/restapi/operations/tlog"
	"github.com/Morrison76/rekor/pkg/log"
	"github.com/Morrison76/rekor/pkg/trillianclient"
	"github.com/Morrison76/rekor/pkg/util"
)

// GetLogInfoHandler returns the current size of the tree and the STH
func GetLogInfoHandler(params tlog.GetLogInfoParams) middleware.Responder {
	tc := trillianclient.NewTrillianClient(params.HTTPRequest.Context(), api.logClient, api.treeID)

	// for each inactive shard, get the loginfo
	var inactiveShards []*models.InactiveShardLogInfo
	for _, shard := range api.logRanges.GetInactive() {
		// Get details for this inactive shard
		is, err := inactiveShardLogInfo(params.HTTPRequest.Context(), shard.TreeID, api.cachedCheckpoints)
		if err != nil {
			return handleRekorAPIError(params, http.StatusInternalServerError, fmt.Errorf("inactive shard error: %w", err), unexpectedInactiveShardError)
		}
		inactiveShards = append(inactiveShards, is)
	}

	if swag.BoolValue(params.Stable) && redisClient != nil {
		// key is treeID/latest
		key := fmt.Sprintf("%d/latest", api.logRanges.GetActive().TreeID)
		redisResult, err := redisClient.Get(params.HTTPRequest.Context(), key).Result()
		if err != nil {
			return handleRekorAPIError(params, http.StatusInternalServerError,
				fmt.Errorf("error getting checkpoint from redis: %w", err), "error getting checkpoint from redis")
		}
		// should not occur, a checkpoint should always be present
		if redisResult == "" {
			return handleRekorAPIError(params, http.StatusInternalServerError,
				fmt.Errorf("no checkpoint found in redis: %w", err), "no checkpoint found in redis")
		}
		decoded, err := hex.DecodeString(redisResult)
		if err != nil {
			return handleRekorAPIError(params, http.StatusInternalServerError,
				fmt.Errorf("error decoding checkpoint from redis: %w", err), "error decoding checkpoint from redis")
		}
		checkpoint := util.SignedCheckpoint{}
		if err := checkpoint.UnmarshalText(decoded); err != nil {
			return handleRekorAPIError(params, http.StatusInternalServerError, fmt.Errorf("invalid checkpoint: %w", err), "invalid checkpoint")
		}
		logInfo := models.LogInfo{
			RootHash:       stringPointer(hex.EncodeToString(checkpoint.Hash)),
			TreeSize:       swag.Int64(int64(checkpoint.Size)),
			SignedTreeHead: stringPointer(string(decoded)),
			TreeID:         stringPointer(fmt.Sprintf("%d", api.treeID)),
			InactiveShards: inactiveShards,
		}
		return tlog.NewGetLogInfoOK().WithPayload(&logInfo)
	}

	resp := tc.GetLatest(0)
	if resp.Status != codes.OK {
		return handleRekorAPIError(params, http.StatusInternalServerError, fmt.Errorf("grpc error: %w", resp.Err), trillianCommunicationError)
	}
	result := resp.GetLatestResult

	root := &types.LogRootV1{}
	if err := root.UnmarshalBinary(result.SignedLogRoot.LogRoot); err != nil {
		return handleRekorAPIError(params, http.StatusInternalServerError, err, trillianUnexpectedResult)
	}

	hashString := hex.EncodeToString(root.RootHash)
	treeSize := int64(root.TreeSize)

	scBytes, err := util.CreateAndSignCheckpoint(params.HTTPRequest.Context(),
		viper.GetString("rekor_server.hostname"), api.logRanges.GetActive().TreeID, root.TreeSize, root.RootHash, api.logRanges.GetActive().Signer)
	if err != nil {
		return handleRekorAPIError(params, http.StatusInternalServerError, err, sthGenerateError)
	}

	logInfo := models.LogInfo{
		RootHash:       &hashString,
		TreeSize:       &treeSize,
		SignedTreeHead: stringPointer(string(scBytes)),
		TreeID:         stringPointer(fmt.Sprintf("%d", api.treeID)),
		InactiveShards: inactiveShards,
	}

	return tlog.NewGetLogInfoOK().WithPayload(&logInfo)
}

func stringPointer(s string) *string {
	return &s
}

// GetLogProofHandler returns information required to compute a consistency proof between two snapshots of log
func GetLogProofHandler(params tlog.GetLogProofParams) middleware.Responder {
	if *params.FirstSize > params.LastSize {
		errMsg := fmt.Sprintf(firstSizeLessThanLastSize, *params.FirstSize, params.LastSize)
		return handleRekorAPIError(params, http.StatusBadRequest, fmt.Errorf("consistency proof: %s", errMsg), errMsg)
	}
	tc := trillianclient.NewTrillianClient(params.HTTPRequest.Context(), api.logClient, api.treeID)
	if treeID := swag.StringValue(params.TreeID); treeID != "" {
		id, err := strconv.Atoi(treeID)
		if err != nil {
			log.Logger.Infof("Unable to convert %s to string, skipping initializing client with Tree ID: %v", treeID, err)
		} else {
			tc = trillianclient.NewTrillianClient(params.HTTPRequest.Context(), api.logClient, int64(id))
		}
	}

	resp := tc.GetConsistencyProof(*params.FirstSize, params.LastSize)
	if resp.Status != codes.OK {
		return handleRekorAPIError(params, http.StatusInternalServerError, fmt.Errorf("grpc error: %w", resp.Err), trillianCommunicationError)
	}
	result := resp.GetConsistencyProofResult

	var root types.LogRootV1
	if err := root.UnmarshalBinary(result.SignedLogRoot.LogRoot); err != nil {
		return handleRekorAPIError(params, http.StatusInternalServerError, err, trillianUnexpectedResult)
	}

	hashString := hex.EncodeToString(root.RootHash)
	proofHashes := []string{}

	if proof := result.GetProof(); proof != nil {
		for _, hash := range proof.Hashes {
			proofHashes = append(proofHashes, hex.EncodeToString(hash))
		}
	} else {
		// The proof field may be empty if the requested tree_size was larger than that available at the server
		// (e.g. because there is skew between server instances, and an earlier client request was processed by
		// a more up-to-date instance). root.TreeSize is the maximum size currently observed
		err := fmt.Errorf(lastSizeGreaterThanKnown, params.LastSize, root.TreeSize)
		return handleRekorAPIError(params, http.StatusBadRequest, err, err.Error())
	}

	consistencyProof := models.ConsistencyProof{
		RootHash: &hashString,
		Hashes:   proofHashes,
	}

	return tlog.NewGetLogProofOK().WithPayload(&consistencyProof)
}

func inactiveShardLogInfo(ctx context.Context, tid int64, cachedCheckpoints map[int64]string) (*models.InactiveShardLogInfo, error) {
	tc := trillianclient.NewTrillianClient(ctx, api.logClient, tid)
	resp := tc.GetLatest(0)
	if resp.Status != codes.OK {
		return nil, fmt.Errorf("resp code is %d", resp.Status)
	}
	result := resp.GetLatestResult

	root := &types.LogRootV1{}
	if err := root.UnmarshalBinary(result.SignedLogRoot.LogRoot); err != nil {
		return nil, err
	}

	hashString := hex.EncodeToString(root.RootHash)
	treeSize := int64(root.TreeSize)

	m := models.InactiveShardLogInfo{
		RootHash:       &hashString,
		TreeSize:       &treeSize,
		TreeID:         stringPointer(fmt.Sprintf("%d", tid)),
		SignedTreeHead: stringPointer(cachedCheckpoints[tid]),
	}
	return &m, nil
}

// handlers for APIs that may be disabled in a given instance

func GetLogInfoNotImplementedHandler(_ tlog.GetLogInfoParams) middleware.Responder {
	err := &models.Error{
		Code:    http.StatusNotImplemented,
		Message: "Get Log Info API not enabled in this Rekor instance",
	}

	return tlog.NewGetLogInfoDefault(http.StatusNotImplemented).WithPayload(err)
}

func GetLogProofNotImplementedHandler(_ tlog.GetLogProofParams) middleware.Responder {
	err := &models.Error{
		Code:    http.StatusNotImplemented,
		Message: "Get Log Proof API not enabled in this Rekor instance",
	}

	return tlog.NewGetLogProofDefault(http.StatusNotImplemented).WithPayload(err)
}

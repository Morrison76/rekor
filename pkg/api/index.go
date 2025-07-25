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
	"log"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"github.com/Morrison76/rekor/pkg/generated/models"
	"github.com/Morrison76/rekor/pkg/generated/restapi/operations/index"
	"github.com/Morrison76/rekor/pkg/pki"
	"github.com/Morrison76/rekor/pkg/util"
)

func SearchIndexHandler(params index.SearchIndexParams) middleware.Responder {
	httpReqCtx := params.HTTPRequest.Context()

	queryOperator := params.Query.Operator
	// default to "or" if no operator is specified
	if params.Query.Operator == "" {
		queryOperator = "or"
	}
	var result = NewCollection(queryOperator)

	var lookupKeys []string

	if params.Query.Hash != "" {
		// This must be a valid hash
		sha := strings.ToLower(util.PrefixSHA(params.Query.Hash))
		if queryOperator == "or" {
			lookupKeys = append(lookupKeys, sha)
		} else {
			resultUUIDs, err := indexStorageClient.LookupIndices(httpReqCtx, []string{sha})
			if err != nil {
				return handleRekorAPIError(params, http.StatusInternalServerError, fmt.Errorf("index storage error: %w", err), indexStorageUnexpectedResult)
			}
			result.Add(resultUUIDs)
		}
	}
	if params.Query.PublicKey != nil {
		af, err := pki.NewArtifactFactory(pki.Format(swag.StringValue(params.Query.PublicKey.Format)))
		if err != nil {
			return handleRekorAPIError(params, http.StatusBadRequest, err, unsupportedPKIFormat)
		}
		keyReader, err := util.FileOrURLReadCloser(httpReqCtx, params.Query.PublicKey.URL.String(), params.Query.PublicKey.Content)
		if err != nil {
			return handleRekorAPIError(params, http.StatusBadRequest, err, malformedPublicKey)
		}
		defer keyReader.Close()

		key, err := af.NewPublicKey(keyReader)
		if err != nil {
			return handleRekorAPIError(params, http.StatusBadRequest, err, malformedPublicKey)
		}
		canonicalKey, err := key.CanonicalValue()
		if err != nil {
			return handleRekorAPIError(params, http.StatusInternalServerError, err, failedToGenerateCanonicalKey)
		}

		keyHash := sha256.Sum256(canonicalKey)
		keyHashStr := strings.ToLower(hex.EncodeToString(keyHash[:]))
		if queryOperator == "or" {
			lookupKeys = append(lookupKeys, keyHashStr)
		} else {
			resultUUIDs, err := indexStorageClient.LookupIndices(httpReqCtx, []string{keyHashStr})
			if err != nil {
				return handleRekorAPIError(params, http.StatusInternalServerError, fmt.Errorf("index storage error: %w", err), indexStorageUnexpectedResult)
			}
			result.Add(resultUUIDs)
		}
	}
	if params.Query.Email != "" {
		emailStr := strings.ToLower(params.Query.Email.String())
		if queryOperator == "or" {
			lookupKeys = append(lookupKeys, emailStr)
		} else {
			resultUUIDs, err := indexStorageClient.LookupIndices(httpReqCtx, []string{emailStr})
			if err != nil {
				return handleRekorAPIError(params, http.StatusInternalServerError, fmt.Errorf("index storage error: %w", err), indexStorageUnexpectedResult)
			}
			result.Add(resultUUIDs)
		}
	}

	if params.Query.ParticipantID != "" {
		participantKey := fmt.Sprintf("participantid:%s", strings.ToLower(params.Query.ParticipantID))
		log.Printf("[DEBUG] Search request by participantID: %s", participantKey)
	
		if queryOperator == "or" {
			log.Printf("[DEBUG] Appending participantKey to lookupKeys: %s", participantKey)
			lookupKeys = append(lookupKeys, participantKey)
		} else 
		{
			log.Printf("[DEBUG] Querying indexStorageClient.LookupIndices with key: %s", participantKey)
			resultUUIDs, err := indexStorageClient.LookupIndices(httpReqCtx, []string{participantKey})
			
			if err != nil {
				log.Printf("[ERROR] indexStorageClient.LookupIndices failed: %v", err)
				return handleRekorAPIError(params, http.StatusInternalServerError, fmt.Errorf("index storage error: %w", err), indexStorageUnexpectedResult)
			}

			log.Printf("[DEBUG] Retrieved UUIDs: %v", resultUUIDs)
			result.Add(resultUUIDs)
		}
	}
	
	if len(lookupKeys) > 0 {
		resultUUIDs, err := indexStorageClient.LookupIndices(httpReqCtx, lookupKeys)
		if err != nil {
			return handleRekorAPIError(params, http.StatusInternalServerError, fmt.Errorf("index storage error: %w", err), indexStorageUnexpectedResult)
		}
		result.Add(resultUUIDs)
	}

	return index.NewSearchIndexOK().WithPayload(result.Values())
}

func SearchIndexNotImplementedHandler(_ index.SearchIndexParams) middleware.Responder {
	err := models.Error{
		Code:    http.StatusNotImplemented,
		Message: "Search Index API not enabled in this Rekor instance",
	}

	return index.NewSearchIndexDefault(http.StatusNotImplemented).WithPayload(&err)

}

func addToIndex(ctx context.Context, keys []string, value string) error {
	err := indexStorageClient.WriteIndex(ctx, keys, value)
	if err != nil {
		return fmt.Errorf("redis client: %w", err)
	}
	return nil
}

func storeAttestation(ctx context.Context, uuid string, attestation []byte) error {
	return attestationStorageClient.StoreAttestation(ctx, uuid, attestation)
}

// Uniq is a collection of unique elements.
type Uniq map[string]struct{}

func NewUniq() Uniq {
	return make(Uniq)
}

func (u Uniq) Add(elements ...string) {
	for _, e := range elements {
		u[e] = struct{}{}
	}
}

func (u Uniq) Values() []string {
	var result []string
	for k := range u {
		result = append(result, k)
	}
	return result
}

// Intersect returns the intersection of two collections.
func (u Uniq) Intersect(other Uniq) Uniq {
	result := make(Uniq)
	for k := range u {
		if _, ok := other[k]; ok {
			result[k] = struct{}{}
		}
	}
	return result
}

// Union returns the union of two collections.
func (u Uniq) Union(other Uniq) Uniq {
	result := make(Uniq)
	for k := range u {
		result[k] = struct{}{}
	}
	for k := range other {
		result[k] = struct{}{}
	}
	return result
}

// Collection is a collection of sets.
//
// its resulting values is a union or intersection of all the sets, depending on the operator.
type Collection struct {
	subsets  []Uniq
	operator string
}

// NewCollection creates a new collection.
func NewCollection(operator string) *Collection {
	return &Collection{
		subsets:  []Uniq{},
		operator: operator,
	}
}

// Add adds the elements into a new subset in the collection.
func (u *Collection) Add(elements []string) {
	subset := Uniq{}
	subset.Add(elements...)
	u.subsets = append(u.subsets, subset)
}

// Values flattens the subsets using the operator, and returns the collection as a slice of strings.
func (u *Collection) Values() []string {
	if len(u.subsets) == 0 {
		return []string{}
	}
	subset := u.subsets[0]
	for i := 1; i < len(u.subsets); i++ {
		if strings.EqualFold(u.operator, "and") {
			subset = subset.Intersect(u.subsets[i])
		} else {
			subset = subset.Union(u.subsets[i])
		}
	}
	return subset.Values()
}

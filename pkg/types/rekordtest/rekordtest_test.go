package rekordtest

import (
    "testing"
    "github.com/Morrison76/rekor/pkg/types"
)

func TestRekordTestRegistration(t *testing.T) {
    if _, ok := types.TypeMap.Load("rekordtest"); !ok {
        t.Fatal("rekordtest not registered in TypeMap")
    }
}
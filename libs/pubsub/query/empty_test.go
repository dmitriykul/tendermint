package query_test

import (
	"testing"

	abci "github.com/number571/tendermint/abci/types"
	"github.com/number571/tendermint/libs/pubsub/query"
	"github.com/stretchr/testify/require"
)

func TestEmptyQueryMatchesAnything(t *testing.T) {
	q := query.Empty{}

	testCases := []struct {
		events []abci.Event
	}{
		{
			[]abci.Event{},
		},
		{
			[]abci.Event{
				{
					Type:       "Asher",
					Attributes: []abci.EventAttribute{{Key: "Roth"}},
				},
			},
		},
		{
			[]abci.Event{
				{
					Type:       "Route",
					Attributes: []abci.EventAttribute{{Key: "66"}},
				},
			},
		},
		{
			[]abci.Event{
				{
					Type:       "Route",
					Attributes: []abci.EventAttribute{{Key: "66"}},
				},
				{
					Type:       "Billy",
					Attributes: []abci.EventAttribute{{Key: "Blue"}},
				},
			},
		},
	}

	for _, tc := range testCases {
		match, err := q.Matches(tc.events)
		require.Nil(t, err)
		require.True(t, match)
	}
}

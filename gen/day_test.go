package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func timeMustParse(t *testing.T, iso8601 string) time.Time {
	parsed, err := time.Parse(time.RFC3339, iso8601)
	require.NoError(t, err)

	return parsed
}

func TestDayLimit(t *testing.T) {
	t.Run("Converts to UTC", func(t *testing.T) {
		require.Equal(t, 1, dayLimit(timeMustParse(t, "2020-12-01T00:00:01-05:00")))
	})

	t.Run("Starts on 2020-12-01 at 0500 UTC", func(t *testing.T) {
		assert.PanicsWithValue(t, "It's not december yet!", func() {
			dayLimit(timeMustParse(t, "2020-11-30T00:00:00-05:00"))
		})
		assert.Equal(t, 0, dayLimit(timeMustParse(t, "2020-11-30T23:59:59-05:00")))
		assert.Equal(t, 1, dayLimit(timeMustParse(t, "2020-12-01T00:00:00-05:00")))
		assert.Equal(t, 1, dayLimit(timeMustParse(t, "2020-12-01T00:01:00-05:00")))
	})

	t.Run("Unlocks at Midnight UTC", func(t *testing.T) {
		for i := 1; i <= 25; i++ {
			assert.Equal(t, i, dayLimit(timeMustParse(t, fmt.Sprintf("2020-12-%02dT00:00:00-05:00", i))))
		}
	})

	t.Run("Unlocks all after 2020-12-25", func(t *testing.T) {
		assert.Equal(t, 25, dayLimit(timeMustParse(t, "2020-12-26T00:00:00-05:00")))
		assert.Equal(t, 25, dayLimit(timeMustParse(t, "2021-12-26T00:00:00-05:00")))
	})
}

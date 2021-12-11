package models

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/rapando/hash-engine/utils"
)

// GetNoOfCombinations
// after much deliberation with myself, I arrived at this formula.
// For 5 character string, the number of combinations is:
// 5.1 + 5.4 + 5.4.3 + 5.4.3.2 + 5.4.3.2.1
func GetNoOfCombinations(length int64) *big.Int {
	var combinations = new(big.Int)
	var i, j int64

	for i = 0; i < length; i++ {
		var c = big.NewInt(1)
		for j = 1; j <= i; j++ {
			diff := big.NewInt(length - j)
			c.Mul(c, diff)
		}
		combinations.Add(combinations, c)
	}
	return combinations.Mul(combinations, big.NewInt(length))
}

func GetCombinations(wg *sync.WaitGroup, ctx context.Context, chars string, cache *redis.Client) {
	defer wg.Done()
	var ml = strings.Split(chars, "")

	for _, c := range ml {
		utils.Publish(ctx, c, cache)
	}

	for z := 0; z < len(chars)-1; z++ {
		var tmp []string
		for _, i := range chars {
			for _, k := range ml {
				if !strings.Contains(k, string(i)) {
					x := fmt.Sprintf("%s%c", k, i)
					utils.Publish(ctx, x, cache)
					tmp = append(tmp, x)
				}
			}
		}
		ml = tmp
	}
}

package models

import (
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/rapando/hash-engine/utils"
	"github.com/streadway/amqp"
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

func GetCombinations(chars string, totalCount *big.Int, qChan *amqp.Channel) {
	var ml = strings.Split(chars, "")
	var exchange = os.Getenv("Q_EXCHANGE")
	var totalCountFlt = new(big.Float).SetInt(totalCount)
	var counter = big.NewFloat(1)
	var percentage = new(big.Float)
	for _, c := range ml {
		_ = utils.QPublish(qChan, exchange, c)
		counter.Add(counter, big.NewFloat(1))
		percentage = new(big.Float).Quo(counter, totalCountFlt)
		fmt.Printf("[%.2f%%]\r", percentage.Mul(percentage, big.NewFloat(100)))

	}

	for z := 0; z < len(chars)-1; z++ {
		var tmp []string
		for _, i := range chars {
			for _, k := range ml {
				if !strings.Contains(k, string(i)) {
					x := fmt.Sprintf("%s%c", k, i)
					_ = utils.QPublish(qChan, exchange, x)
					tmp = append(tmp, x)
					counter.Add(counter, big.NewFloat(1))
					percentage = new(big.Float).Quo(counter, totalCountFlt)
					fmt.Printf("[%.2f%%]\r", percentage.Mul(percentage, big.NewFloat(100)))
				}
			}
		}
		ml = tmp
	}
}

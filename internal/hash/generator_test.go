package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GenerateUniqueCustomerID(t *testing.T) {
	unixTime := int64(1597726137)
	hash, _ := GenerateUniqueCustomerID("Misha", "1234567890", unixTime)
	assert.Equal(t, "09b843b24f5c966771ce2029a173c9ad", hash)
}

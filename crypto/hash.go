package crypto

import (
	ghash "bitbucket.org/number571/go-cryptopro/gost_r_34_11_2012"
)

func HashSum(bytes []byte) []byte {
	return ghash.Sum(ghash.H256, bytes)
}

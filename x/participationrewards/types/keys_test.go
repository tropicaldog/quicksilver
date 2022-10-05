package types

import (
	"testing"

	"github.com/ingenuity-build/quicksilver/utils"
	"github.com/stretchr/testify/require"
)

func TestKeys(t *testing.T) {
	address := utils.GenerateAccAddressForTest()
	keyClaim := GetKeyClaim("testzone-1", address.String(), ClaimTypeOsmosisPool, "testzone-2")
	prefixClaim := GetPrefixClaim("testzone-1")
	prefixClaimForUser := GetPrefixUserClaim("testzone-1", address.String())

	require.Equal(t, append([]byte{0x01}, []byte("testzone-1")...), prefixClaim)
	require.Equal(t, append([]byte{0x01}, []byte("testzone-1"+address.String())...), prefixClaimForUser)
	require.Equal(t, append([]byte{0x01}, append(append([]byte("testzone-1"+address.String()), []byte{0x00, 0x00, 0x00, 0x02}...), []byte("testzone-2")...)...), keyClaim)
}

func TestLastEpochKeys(t *testing.T) {
	address := utils.GenerateAccAddressForTest()
	keyClaim := GetKeyLastEpochClaim("testzone-1", address.String(), ClaimTypeOsmosisPool, "testzone-2")
	prefixClaim := GetPrefixLastEpochClaim("testzone-1")
	prefixClaimForUser := GetPrefixLastEpochUserClaim("testzone-1", address.String())

	require.Equal(t, append([]byte{0x02}, []byte("testzone-1")...), prefixClaim)
	require.Equal(t, append([]byte{0x02}, []byte("testzone-1"+address.String())...), prefixClaimForUser)
	require.Equal(t, append([]byte{0x02}, append(append([]byte("testzone-1"+address.String()), []byte{0x00, 0x00, 0x00, 0x02}...), []byte("testzone-2")...)...), keyClaim)
}
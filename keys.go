package lorawan

import (
	"crypto/aes"

	"github.com/tkiraly/lorawan/joinaccept"
	"github.com/tkiraly/lorawan/joinrequest"
)

//GenerateNwkSKey generates network session key based on Join Request and Join Accept
func GenerateNwkSKey(jr, ja, appkey []byte) ([]byte, error) {
	return generatekey(0x01, jr, ja, appkey)
}

//GenerateAppSKey generates application session key based on Join Request and Join Accept
func GenerateAppSKey(jr, ja, appkey []byte) ([]byte, error) {
	return generatekey(0x02, jr, ja, appkey)
}

func generatekey(lead byte, jr, ja, appkey []byte) ([]byte, error) {
	pja, err := joinaccept.Parse(ja, appkey)
	if err != nil {
		return nil, err
	}
	pjr, err := joinrequest.Parse(jr)
	if err != nil {
		return nil, err
	}

	B := make([]byte, 16)
	B[0] = lead
	an := pja.AppNonce()
	B[1] = an[2]
	B[2] = an[1]
	B[3] = an[0]
	nid := pja.NetID()
	B[4] = nid[2]
	B[5] = nid[1]
	B[6] = nid[0]
	dn := pjr.DevNonce()
	B[7] = dn[1]
	B[8] = dn[0]

	cipher, _ := aes.NewCipher(appkey)
	r := make([]byte, 16)
	cipher.Encrypt(r, B)
	return r, nil
}

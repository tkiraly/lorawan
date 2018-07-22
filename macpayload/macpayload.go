package macpayload

import "crypto/aes"

func encrypt(payload []byte, updir bool, devaddr []byte, cnt uint32, key []byte) []byte {
	var dir byte = 0x01
	if updir {
		dir = 0x00
	}

	/*
	   this is an example how Block a is formed according to Lorawan specs
	   Block_A[0]  = 0x01;
	   Block_A[1]  = 0x00;
	   Block_A[2]  = 0x00;
	   Block_A[3]  = 0x00;
	   Block_A[4]  = 0x00;
	   Block_A[5]  = 0;        // 0 for uplink frames 1 for downlink frames;
	   Block_A[6]  = dev_addr[3]; // LSB devAddr 4 bytes
	   Block_A[7]  = dev_addr[2];  // ..
	   Block_A[8]  = dev_addr[1];  // ..
	   Block_A[9]  = dev_addr[0];  // MSB
	   Block_A[10] = sequence_counter & 0xff;  // LSB framecounter
	   Block_A[11] = (sequence_counter >> 8) & 0xff;  // MSB framecounter
	   Block_A[12] = (sequence_counter >> 16) & 0xff;     // Frame counter upper Bytes
	   Block_A[13] = (sequence_counter >> 24) & 0xff;
	   Block_A[14] = 0x00;
	*/

	block := []byte{0x01, 0x00, 0x00, 0x00,
		0x00, dir, devaddr[3], devaddr[2],
		devaddr[1], devaddr[0], byte(cnt & 0xff), byte((cnt >> 8) & 0xff),
		byte((cnt >> 16) & 0xff), byte((cnt >> 24) & 0xff), 0x00, 0x00}

	cipher, _ := aes.NewCipher(key)
	r := make([]byte, len(payload))
	bufferIndex := 0
	size := len(payload)
	ctr := 1
	S := make([]byte, 16)

	for size >= 16 {
		block[15] = byte(ctr & 0xFF)
		ctr += 1
		cipher.Encrypt(S, block)

		for i := 0; i < 16; i++ {
			r[bufferIndex+i] = payload[bufferIndex+i] ^ S[i]
		}
		size -= 16
		bufferIndex += 16
	}
	if size > 0 {
		block[15] = byte(ctr & 0xFF)
		ctr += 1

		cipher.Encrypt(S, block)
		for i := 0; i < size; i++ {
			r[bufferIndex+i] = payload[bufferIndex+i] ^ S[i]
		}
	}
	return r
}

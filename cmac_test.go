package cmac

import (
	"bytes"
	"crypto/aes"
	"testing"
)

func TestGenSubkeys(t *testing.T) {
	c, err := aes.NewCipher([]byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c})
	if err != nil {
		t.Fatal(err)
	}
	z := make([]byte, c.BlockSize())
	c.Encrypt(z, z)

	k1, k2 := gensubkeys(c)

	if !bytes.Equal(k1, []byte{0xfb, 0xee, 0xd6, 0x18, 0x35, 0x71, 0x33, 0x66, 0x7c, 0x85, 0xe0, 0x8f, 0x72, 0x36, 0xa8, 0xde}) {
		t.Errorf("unexpected subkey k1")
	}

	if !bytes.Equal(k2, []byte{0xf7, 0xdd, 0xac, 0x30, 0x6a, 0xe2, 0x66, 0xcc, 0xf9, 0x0b, 0xc1, 0x1e, 0xe4, 0x6d, 0x51, 0x3b}) {
		t.Errorf("unexpected subkey k2")
	}
}

func TestAESCMAC(t *testing.T) {
	key := []byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c}
	msg := []byte{0x6b, 0xc1, 0xbe, 0xe2, 0x2e, 0x40, 0x9f, 0x96, 0xe9, 0x3d, 0x7e, 0x11, 0x73, 0x93, 0x17, 0x2a}
	mac := []byte{0x07, 0x0a, 0x16, 0xb4, 0x6b, 0x4d, 0x41, 0x44, 0xf7, 0x9b, 0xdd, 0x9d, 0xd0, 0x4a, 0x28, 0x7c}

	m := New(key)
	m.Write(msg)
	tmac := m.Sum(nil)

	if !bytes.Equal(mac, tmac) {
		t.Errorf("expected %x got %x\n", mac, tmac)
	}
}
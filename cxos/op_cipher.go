// +build os

package cxos

import (
	"github.com/skycoin/skycoin/src/cipher"

	"github.com/skycoin/cx/cx"
)

func init() {
	cipherPkg := cxcore.MakePackage("cipher")
	pubkeyStrct := cxcore.MakeStruct("PubKey")
	seckeyStrct := cxcore.MakeStruct("SecKey")

	// PubKey
	pubkeyFld := cxcore.MakeArgument("PubKey", "", -1).AddType(cxcore.TypeNames[cxcore.TYPE_UI8]).AddPackage(cipherPkg)
	pubkeyFld.DeclarationSpecifiers = append(pubkeyFld.DeclarationSpecifiers, cxcore.DECL_ARRAY)
	pubkeyFld.IsArray = true
	pubkeyFld.Lengths = []int{33} // Yes, PubKey is 33 bytes long.
	pubkeyFld.TotalSize = 33      // 33 * 1 byte (ui8)

	// SecKey
	seckeyFld := cxcore.MakeArgument("SecKey", "", -1).AddType(cxcore.TypeNames[cxcore.TYPE_UI8]).AddPackage(cipherPkg)
	seckeyFld.DeclarationSpecifiers = append(seckeyFld.DeclarationSpecifiers, cxcore.DECL_ARRAY)
	seckeyFld.IsArray = true
	seckeyFld.Lengths = []int{32} // Yes, SecKey is 32 bytes long.
	seckeyFld.TotalSize = 33      // 33 * 1 byte (ui8)

	pubkeyStrct.AddField(pubkeyFld)
	seckeyStrct.AddField(seckeyFld)

	cipherPkg.AddStruct(pubkeyStrct)
	cipherPkg.AddStruct(seckeyStrct)

	cxcore.PROGRAM.AddPackage(cipherPkg)
}

// opCipherGenerateKeyPair generates a PubKey and a SecKey.
func opCipherGenerateKeyPair(inputs []cxcore.CXValue, outputs []cxcore.CXValue) {
	pubKey, secKey := cipher.GenerateKeyPair()

	bPubKey := make([]byte, len(pubKey))
	bSecKey := make([]byte, len(secKey))

	// Copying bytes
	for i, byt := range pubKey {
		bPubKey[i] = byt
	}
	for i, byt := range secKey {
		bSecKey[i] = byt
	}

    outputs[0].Set_bytes(bPubKey)
    outputs[1].Set_bytes(bSecKey)
}

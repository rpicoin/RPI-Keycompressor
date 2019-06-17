package main

import (
	//"encoding/hex"
	"errors"

	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

type Network struct {
	name        string
	symbol      string
	xpubkey     byte
	xprivatekey byte
}

var network = map[string]Network{
	"rdd": {name: "reddcoin", symbol: "rdd", xpubkey: 0x3d, xprivatekey: 0xbd},
	"dgb": {name: "digibyte", symbol: "dgb", xpubkey: 0x1e, xprivatekey: 0x80},
	"btc": {name: "bitcoin", symbol: "btc", xpubkey: 0x00, xprivatekey: 0x80},
	"ltc": {name: "litecoin", symbol: "ltc", xpubkey: 0x30, xprivatekey: 0xb0},
	"wsp": {name: "wispr", symbol: "wsp", xpubkey: 0x49, xprivatekey: 0x91},
	"rpi": {name: "rpicoin", symbol: "rpi", xpubkey: 0x3C, xprivatekey: 0x91},
}

func (network Network) GetNetworkParams() *chaincfg.Params {
	networkParams := &chaincfg.MainNetParams
	networkParams.PubKeyHashAddrID = network.xpubkey
	networkParams.PrivateKeyID = network.xprivatekey
	return networkParams
}

func (network Network) CreatePrivateKey() (*btcutil.WIF, error) {
	secret, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return nil, err
	}
	return btcutil.NewWIF(secret, network.GetNetworkParams(), true)
}
func (network Network) GetDecompressedPrivateKey(privKey *btcec.PrivateKey) (*btcutil.WIF, error) {
	return btcutil.NewWIF(privKey, network.GetNetworkParams(), false)
}

//func (network Network) ImportPrivateKey(secretHex string) (*btcutil.WIF, error) {
//	secret, err := btcec.SetPrivateKey(secretHex)
//	if err != nil {
//		return nil, err
//	}
//	return btcutil.NewWIF(secret, network.GetNetworkParams(), true)
//}

func (network Network) ImportWIF(wifStr string) (*btcutil.WIF, error) {
	wif, err := btcutil.DecodeWIF(wifStr)
	if err != nil {
		return nil, err
	}
	if !wif.IsForNet(network.GetNetworkParams()) {
		return nil, errors.New("The WIF string is not valid for the `" + network.name + "` network")
	}
	return wif, nil
}

func (network Network) GetAddress(wif *btcutil.WIF) (*btcutil.AddressPubKey, error) {
	return btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeCompressed(), network.GetNetworkParams())
}
func (network Network) GetDecompressedAddress(wif *btcutil.WIF) (*btcutil.AddressPubKey, error) {
	return btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeUncompressed(), network.GetNetworkParams())
}

func main() {
	fmt.Println("Starting the application...")
	wifC, _ := network["rpi"].CreatePrivateKey()
	wifD, _ := network["rpi"].GetDecompressedPrivateKey(wifC.PrivKey)
	addressC, _ := network["rpi"].GetAddress(wifC)
	addressD, _ := network["rpi"].GetDecompressedAddress(wifC)
	fmt.Printf("Compressed: %s - %s Decompressed: %s - %s", wifC.String(), addressC.EncodeAddress(), wifD.String(), addressD.EncodeAddress())
}

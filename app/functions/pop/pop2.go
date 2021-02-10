package pop

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

// Used by all functions to identify
func PopulateWallet2(wallet *gateway.Wallet) error {
	credPath := filepath.Join(
		"/home",
		"rtorrealba",
		"ubiot-network",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org2.example.com",
		"users",
		"User1@org2.example.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore2")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return errors.New("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org2MSP", string(cert), string(key))

	err = wallet.Put("appUser2", identity)
	if err != nil {
		return err
	}
	return nil
}

package config

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func generate_ecdsa_keypair(certificate_path, private_key_path string) (err error) {
	certificate_duration := 365 * 24 * time.Hour
	invalid_before := time.Now()
	invalid_after := invalid_before.Add(certificate_duration)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return fmt.Errorf("failed to generate serial number: %s", err)
	}
	rootKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return err
	}
	if err = ec_private_key_to_file(private_key_path, rootKey); err != nil {
		return
	}

	rootTemplate := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Gophercraft"},
			CommonName:   "GC",
		},
		NotBefore:             invalid_before,
		NotAfter:              invalid_after,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	certificate_bytes, err := x509.CreateCertificate(rand.Reader, &rootTemplate, &rootTemplate, &rootKey.PublicKey, rootKey)
	if err != nil {
		return err
	}
	err = ec_certificate_to_file(certificate_path, certificate_bytes)
	return
}

func ec_private_key_to_file(filename string, key *ecdsa.PrivateKey) (err error) {
	var file *os.File
	var ec_key_data []byte
	file, err = os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()
	ec_key_data, err = x509.MarshalECPrivateKey(key)
	if err != nil {
		return
	}
	if err = pem.Encode(file, &pem.Block{Type: "EC PRIVATE KEY", Bytes: ec_key_data}); err != nil {
		return
	}
	return
}

func ec_certificate_to_file(filename string, ec_certificate_bytes []byte) (err error) {
	var file *os.File
	file, err = os.Create(filename)
	if err != nil {
		return
	}
	if err = pem.Encode(file, &pem.Block{Type: "CERTIFICATE", Bytes: ec_certificate_bytes}); err != nil {
		return
	}
	return file.Close()
}

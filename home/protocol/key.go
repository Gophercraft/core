package protocol

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"time"
)

func ec_private_key_to_pem(key *ecdsa.PrivateKey) (data []byte, err error) {
	var ec_key_data []byte
	ec_key_data, err = x509.MarshalECPrivateKey(key)
	if err != nil {
		return
	}
	var buffer bytes.Buffer
	if err = pem.Encode(&buffer, &pem.Block{Type: "EC PRIVATE KEY", Bytes: ec_key_data}); err != nil {
		return
	}
	data = buffer.Bytes()
	return
}

func ec_certificate_to_pem(ec_certificate_bytes []byte) (data []byte, err error) {
	var buffer bytes.Buffer
	if err = pem.Encode(&buffer, &pem.Block{Type: "CERTIFICATE", Bytes: ec_certificate_bytes}); err != nil {
		return
	}
	data = buffer.Bytes()
	return
}

func generate_ephemeral_ecdsa_keypair() (certificate *tls.Certificate, err error) {
	certificate_duration := 365 * 24 * time.Hour
	invalid_before := time.Now()
	invalid_after := invalid_before.Add(certificate_duration)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		err = fmt.Errorf("failed to generate serial number: %s", err)
		return
	}
	var rootKey *ecdsa.PrivateKey
	rootKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
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
	var (
		certificate_bytes []byte
		certificate_pem   []byte
		private_key_pem   []byte
	)

	private_key_pem, err = ec_private_key_to_pem(rootKey)
	if err != nil {
		return
	}

	certificate_bytes, err = x509.CreateCertificate(rand.Reader, &rootTemplate, &rootTemplate, &rootKey.PublicKey, rootKey)
	if err != nil {
		return
	}

	certificate_pem, err = ec_certificate_to_pem(certificate_bytes)
	if err != nil {
		return
	}

	new_certificate, err := tls.X509KeyPair(certificate_pem, private_key_pem)
	return &new_certificate, err
}

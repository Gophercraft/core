package protocol

import (
	context "context"
	"crypto/sha512"
	"crypto/subtle"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	fmt "fmt"
	"os"
	"strings"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
)

type Fingerprint [64]byte

func (print Fingerprint) String() string {
	parts := make([]string, 64)

	for i := range parts {
		parts[i] = hex.EncodeToString(print[i : (i)+1])
	}

	return strings.Join(parts, ":")
}

func (print *Fingerprint) DecodeWord(data string) (err error) {
	parts := strings.Split(data, ":")
	if len(parts) != 64 {
		err = fmt.Errorf("invalid fingerprint: invalid number of parts")
		return
	}
	var bytes []byte
	bytes, err = hex.DecodeString(strings.Join(parts, ""))
	if err != nil {
		return
	}
	if len(bytes) != 64 {
		err = fmt.Errorf("invalid fingerprint: invalid number of bytes")
		return
	}
	copy(print[:], bytes)
	return
}

func (print *Fingerprint) EncodeWord() (word string, err error) {
	word = print.String()
	return
}

func GetCertFileFingerprint(certificate_path string) (print Fingerprint, err error) {
	var (
		cert_file []byte
		cert      *x509.Certificate
		pem_block *pem.Block
	)
	cert_file, err = os.ReadFile(certificate_path)
	if err != nil {
		return
	}

	pem_block, _ = pem.Decode(cert_file)

	cert, err = x509.ParseCertificate(pem_block.Bytes)
	if err != nil {
		return
	}

	return GetCertFingerprint(cert)
}

func GetCertFingerprint(cert *x509.Certificate) (print Fingerprint, err error) {
	var pkix_der []byte
	pkix_der, err = x509.MarshalPKIXPublicKey(cert.PublicKey)
	if err != nil {
		return
	}

	print = Fingerprint(sha512.Sum512(pkix_der))
	return
}

func GetPeerFingerprint(ctx context.Context) (print Fingerprint, err error) {
	peer, ok := peer.FromContext(ctx)
	if !ok {
		err = fmt.Errorf("could not extract peer information")
		return
	}

	tlsInfo, ok := peer.AuthInfo.(credentials.TLSInfo)
	if !ok {
		err = fmt.Errorf("peer does not have tls info")
		return
	}

	certCount := len(tlsInfo.State.PeerCertificates)

	if certCount != 1 {
		err = fmt.Errorf("invalid certificate count (%d)", certCount)
		return
	}

	cert := tlsInfo.State.PeerCertificates[0]

	return GetCertFingerprint(cert)
}

func FingerprintsEqual(f1, f2 Fingerprint) bool {
	return subtle.ConstantTimeCompare(f1[:], f2[:]) == 1
}

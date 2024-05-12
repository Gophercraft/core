package protocol

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func FingerprintServer(address string) (Fingerprint, error) {
	conn, err := tls.Dial("tcp", address, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return Fingerprint{}, err
	}

	cert := conn.ConnectionState().PeerCertificates[0]

	// Waltuh...
	finger, err := GetCertFingerprint(cert)
	if err != nil {
		return Fingerprint{}, err
	}

	conn.Close()

	// log.Println(address, "fingerprint is", finger)

	return finger, nil
}

// open a GRPC clientConn. the server's fingerprint must match the provided fingerprint, or the function will fail
func DialConn(address string, trusted_fingerprint Fingerprint, cert *tls.Certificate) (*grpc.ClientConn, error) {
	gc, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
			config := &tls.Config{
				InsecureSkipVerify: true,
				MinVersion:         tls.VersionTLS12,
			}

			if cert != nil {
				config.Certificates = []tls.Certificate{
					*cert,
				}
			} else {
				cert, err := generate_ephemeral_ecdsa_keypair()
				if err != nil {
					return nil, err
				}
				config.Certificates = []tls.Certificate{
					*cert,
				}
			}

			conn, err := tls.Dial("tcp", address, config)
			if err != nil {
				return nil, err
			}

			cert := conn.ConnectionState().PeerCertificates[0]

			peer_certificate_fingerprint, err := GetCertFingerprint(cert)
			if err != nil {
				return nil, err
			}

			if !FingerprintsEqual(peer_certificate_fingerprint, trusted_fingerprint) {
				conn.Close()
				return nil, fmt.Errorf("home/protocol: Server connection fingerprint does not match the one on record!\nServer certificate: %s\nTrusted certificate: %s", peer_certificate_fingerprint, trusted_fingerprint)
			}

			return conn, nil
		}),
	)
	if err != nil {
		return nil, err
	}

	return gc, nil
}

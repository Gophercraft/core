package rpcnet

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func FingerprintServer(address string) (string, error) {
	conn, err := tls.Dial("tcp", address, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return "", err
	}

	cert := conn.ConnectionState().PeerCertificates[0]

	// Waltuh...
	finger, err := GetCertFingerprint(cert)
	if err != nil {
		return "", err
	}

	conn.Close()

	log.Println(address, "fingerprint is", finger)

	return finger, nil
}

// open a GRPC clientConn. the server's fingerprint must match the provided fingerprint, or the function will fail
func DialConn(address string, fingerprint string, cert *tls.Certificate) (*grpc.ClientConn, error) {
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
			}

			conn, err := tls.Dial("tcp", address, config)
			if err != nil {
				return nil, err
			}

			cert := conn.ConnectionState().PeerCertificates[0]

			finger, err := GetCertFingerprint(cert)
			if err != nil {
				return nil, err
			}

			if !FingerprintsEqual(finger, fingerprint) {
				conn.Close()
				return nil, fmt.Errorf("sys: Peer fingerprint mismatch! should be %s, received %s", fingerprint, finger)
			}

			return conn, nil
		}),
	)
	if err != nil {
		return nil, err
	}

	return gc, nil
}

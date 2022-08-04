package protocol

import (
	"crypto/tls"
	"fmt"
	"google.golang.org/grpc/keepalive"
	"time"

	"github.com/anima-protocol/anima-go/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var client AnimaClient

type Config struct {
	Secure bool
}

// Init - Initialize New Client
func Init(config *Config, protocol *models.Protocol) error {
	if client == nil {
		fmt.Printf("-> Anima Client")
		creds := credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})

		opts := []grpc.DialOption{
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:                5 * time.Minute,
				Timeout:             20 * time.Second,
				PermitWithoutStream: true,
			}),
		}

		if config.Secure {
			opts = append(opts, grpc.WithTransportCredentials(creds))
		} else {
			opts = append(opts, grpc.WithInsecure())
		}

		fmt.Printf("-> network: %v\n", protocol.Network)
		cc, err := grpc.Dial(protocol.Network, opts...)
		if err != nil {
			return fmt.Errorf("could not connect to GRPC Server")
		}

		client = NewAnimaClient(cc)
	}

	return nil
}

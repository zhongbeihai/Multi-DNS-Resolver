package pkg

import (
	"context"
	"time"

	"github.com/miekg/dns"
)

type DNSResolver interface {
	Resolve(ctx context.Context, dnsServer, qname string, qtype uint16) (*dns.Msg, time.Duration, error)
}

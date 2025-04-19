package pkg

import (
	"context"
	"github.com/miekg/dns"
)

type DNSResolver interface {
	Resolve(ctx context.Context, dnsServer, qname string, qtype uint16) (*dns.Msg, error)
}

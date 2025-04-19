package pkg

import (
	"time"
	"context"
	"github.com/miekg/dns"
)
type UDPResolver struct{
	Timeout time.Duration
}

func (u *UDPResolver) Resolve(ctx context.Context, dnsServer, qname string, qtype uint16) (*dns.Msg, error){
	client := dns.Client{Net:"udp", Timeout: u.Timeout}
	msg := new(dns.Msg)
	msg.SetQuestion(dns.Fqdn(qname), qtype)
	res, _ , err := client.ExchangeContext(ctx, msg, dnsServer)

	return res, err
}

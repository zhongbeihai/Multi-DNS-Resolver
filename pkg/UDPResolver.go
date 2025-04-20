package pkg

import (
	"time"
	"context"
	"github.com/miekg/dns"
)
type UDPResolver struct{
	Timeout time.Duration
}

func (u *UDPResolver) Resolve(ctx context.Context, dnsServer, qname string, qtype uint16) (*dns.Msg, time.Duration, error){
	client := dns.Client{Net:"udp", Timeout: u.Timeout}
	msg := new(dns.Msg)
	msg.SetQuestion(dns.Fqdn(qname), qtype)
	res, rtt , err := client.ExchangeContext(ctx, msg, dnsServer)

	return res, rtt,err
}

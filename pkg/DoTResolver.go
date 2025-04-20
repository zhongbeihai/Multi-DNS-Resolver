package pkg

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	"github.com/miekg/dns"
)

type DoTResolver struct {
	Timeout time.Duration
}

func (r *DoTResolver) Resolve(ctx context.Context, dnsServer, qname string, 
		qtype uint16) (*dns.Msg, time.Duration, error){
		host, _, err := net.SplitHostPort(dnsServer)
		if err != nil {
			host = dnsServer
		}
		config := &tls.Config{ServerName: host}
		client := dns.Client{Net: "tcp-tls", Timeout: r.Timeout, TLSConfig: config}

		msg := new(dns.Msg)
		msg.SetQuestion(dns.Fqdn(qname), qtype)

		res, rtt, err := client.ExchangeContext(ctx, msg, dnsServer)

		return res, rtt, err
	}
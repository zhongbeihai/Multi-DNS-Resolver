package multidnsresolver

import (
	"context"
	"multi-dns-resolver/pkg"
	"time"

	"github.com/miekg/dns"
)

type DNSServer struct{
	Addr string // e.g. "8.8.8.8:53", "https://dns.google/dns-query", "1.1.1.1:853"
	Protocol string // "udp", "doh", or "dot"
}

type DNSClient struct{
	servers []DNSServer
	resolvers map[string]pkg.DNSResolver
}

func NewClient(servers []DNSServer) *DNSClient{
	return &DNSClient{
		servers: servers,
		resolvers: map[string]pkg.DNSResolver{
			"udp": &pkg.UDPResolver{Timeout: 5 * time.Second},
			//"doh": &DoHResolver{Client: &http.Client{Timeout: 10 * time.Second}},
			//"dot": &DoTResolver{Timeout: 5 * time.Second},
		},
	}
}

func (c *DNSClient) Query(qname string, qtype uint16) (*dns.Msg, error){
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	type result struct {
		msg *dns.Msg
		err error
	}
	resCh := make(chan result, len(c.servers))


	for _, server := range c.servers{
		resolver, ok := c.resolvers[server.Protocol]
		if !ok {
			continue
		}

		go func (s DNSServer, r pkg.DNSResolver)  {
			msg, err := r.Resolve(ctx, s.Addr, qname, qtype)
			select{
			case resCh <- result{msg: msg, err: err}:
			case <- ctx.Done():
			}
		}(server, resolver)
	}

	var firstErr error
	for i := 0; i < len(c.servers); i++{
		r := <-resCh
		if r.err == nil && r.msg != nil && len(r.msg.Answer) > 0 {
			cancel()
			return r.msg, nil
		}
		if firstErr == nil {
			firstErr = r.err
		}
	}

	return nil, firstErr
}
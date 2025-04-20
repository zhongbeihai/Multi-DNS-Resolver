package pkg

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/miekg/dns"
)

type DoHResolver struct {
	Client *http.Client
}

func (r *DoHResolver) Resolve(ctx context.Context, dnsServer, qname string, 
	qtype uint16) (*dns.Msg, time.Duration, error){
		endpoint, err := url.Parse(dnsServer)
		if err != nil {
			return nil, time.Duration(0), err
		}

		q := endpoint.Query()
		q.Set("name", qname)
		q.Set("type", fmt.Sprintf("%d", qtype))
		endpoint.RawQuery = q.Encode()

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), nil)
		if err != nil{
			return nil, time.Duration(0), err
		}
		req.Header.Set("Accept", "application/dns-message")

		resp, err := r.Client.Do(req)
		if err != nil {
			return nil, time.Duration(0), err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, time.Duration(0), err
		}

		msg := new(dns.Msg)
		if err = msg.Unpack(body); err != nil{
			return nil, time.Duration(0), err
		}
		return msg, 0, nil
	}
package dnsc

import (
	"github.com/miekg/dns"
	"net"
	"log"
	"fmt"
	"errors"
)

type DNSClient struct {
	MasterDNS	string
	MasterDNSPort	string
	Debug		bool
}

func Init(sp *DNSClient) DNSClient {
	var s DNSClient
	if sp != nil {
		s = *sp
	} else {
		s = DNSClient{
			Debug: false,
		}
	}
	return s
}

func (s *DNSClient) SendMsg(m *dns.Msg) (msgAnswer []dns.RR, err error) {
	c := new(dns.Client)
	r, _, errE := c.Exchange(m, net.JoinHostPort(s.MasterDNS, s.MasterDNSPort))
	if r == nil {
		return []dns.RR{}, errE
	}
	var sMsgAnswer string
	for _, a := range r.Answer {
		sMsgAnswer += fmt.Sprintf("%v\n", a)
	}
	if s.Debug {
		log.Println(fmt.Sprintf("The result code is %v, r.Answer %v", r.Rcode, sMsgAnswer))
	}
	if (r.Rcode != dns.RcodeSuccess) {
		return []dns.RR{}, errors.New(fmt.Sprintf("The result code is %v, r.Answer %v", r.Rcode, sMsgAnswer))
	}
	return r.Answer, nil
}

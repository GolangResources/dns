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

func (s *DNSClient) AddDNS(zone string, fqdn string, ip string, ttl uint32) error {
	msg := new(dns.Msg)
	msg.SetUpdate(zone)
	msgRR := []dns.RR{
		&dns.A{
			Hdr: dns.RR_Header{Name: dns.Fqdn(fqdn), Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: ttl},
			A:   net.ParseIP(ip),
		},
	}
	msg.Insert(msgRR)
	_, err := s.SendMsg(msg)
	return err
}

func (s *DNSClient) DelDNS(zone string, fqdn string, ip string, ttl uint32) error {
	msg := new(dns.Msg)
	msg.SetUpdate(zone)
	msgRR := []dns.RR{
		&dns.A{
			Hdr: dns.RR_Header{Name: dns.Fqdn(fqdn), Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: ttl},
			A:   net.ParseIP(ip),
		},
	}
	msg.Remove(msgRR)
	_, err := s.SendMsg(msg)
	return err
}

func (s *DNSClient) DelAllDNS(zone string, fqdn string) error {
	msg := new(dns.Msg)
	msg.SetUpdate(zone)
	msgRR := []dns.RR{
		&dns.RR_Header{Name: dns.Fqdn(fqdn)},
	}
	msg.RemoveName(msgRR)
	_, err := s.SendMsg(msg)
	return err
}

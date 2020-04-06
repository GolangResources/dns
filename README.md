# DNS

## Examples
```
package main

import (
        "github.com/miekg/dns"
        "github.com/GolangResources/dns/v1"
        "net"
        "log"
)

//BASED ON -> https://github.com/SpComb/go-nsupdate/blob/master/update.go

func main() {
        dnsConf := dnsc.DNSClient{
                MasterDNS: "127.0.0.1",
                MasterDNSPort: "53",
                Debug: false,
        }
        d := dnsc.Init(&dnsConf)
        msg := new(dns.Msg)
        msg.SetUpdate("companydomain.com.")
        //DEL ALL RECORDS
        msgDNSRemoveALLNames := []dns.RR{
                        &dns.RR_Header{Name: "lol.companydomain.com."},
        }
        msg.RemoveName(msgDNSRemoveALLNames)
        //DEL SPECIFIC RECORD
        msgDNSRemoveNames := []dns.RR{
                        &dns.A{
                                Hdr: dns.RR_Header{Name: "lol.companydomain.com.", Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: uint32(30)},
                                A:   net.IPv4(127, 0, 0, 1),
                        },
        }
        msg.Remove(msgDNSRemoveNames)
        //ADD RECORD
        msgDNSAddNames := []dns.RR{
                        &dns.A{
                                Hdr: dns.RR_Header{Name: "lol.companydomain.com.", Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: uint32(30)},
                                A:   net.IPv4(127, 0, 0, 1),
                        },
        }
        msg.Insert(msgDNSAddNames)
        log.Println(d.SendMsg(msg))
        //QUERY A RECORD
        msgc := new(dns.Msg)
        msgc.SetQuestion(dns.Fqdn("lol.companydomain.com"), dns.TypeA)
        msgc.RecursionDesired = true
        log.Println(d.SendMsg(msgc))
}
```

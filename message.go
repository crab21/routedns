package rdns

import (
	"strconv"

	"github.com/miekg/dns"
)

// Return the query name from a DNS query.
func qName(q *dns.Msg) string {
	if len(q.Question) == 0 {
		return ""
	}
	return q.Question[0].Name
}

// Return the result code name from a DNS response.
func rCode(r *dns.Msg) string {
	if result, ok := dns.RcodeToString[r.Rcode]; ok {
		return result
	}
	return strconv.Itoa(r.Rcode)
}

// Returns a NXDOMAIN answer for a query.
func nxdomain(q *dns.Msg) *dns.Msg {
	return responseWithCode(q, dns.RcodeNameError)
}

// Returns a SERVFAIL answer for a query.
func servfail(q *dns.Msg) *dns.Msg {
	return responseWithCode(q, dns.RcodeServerFailure)
}

// Returns a REFUSED answer for a query.
func refused(q *dns.Msg) *dns.Msg {
	return responseWithCode(q, dns.RcodeRefused)
}

// Build a response for a query with the given responce code.
func responseWithCode(q *dns.Msg, rcode int) *dns.Msg {
	a := new(dns.Msg)
	a.SetRcode(q, rcode)
	return a
}

// Answers a PTR query with a name
func ptr(q *dns.Msg, name string) *dns.Msg {
	a := new(dns.Msg)
	a.SetReply(q)
	a.Answer = []dns.RR{
		&dns.PTR{
			Hdr: dns.RR_Header{
				Name:   q.Question[0].Name,
				Rrtype: dns.TypePTR,
				Class:  dns.ClassINET,
				Ttl:    3600,
			},
			Ptr: dns.Fqdn(name),
		},
	}
	return a
}

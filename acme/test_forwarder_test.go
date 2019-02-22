package acme

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/miekg/dns"
)

const forwardingAddress = "8.8.8.8:53"

// testForwarder is a DNS forwarder designed to test the custom
// propagation nameserver functionality.
type testForwarder struct {
	called bool
	conn   net.PacketConn
	server *dns.Server
}

// newTestForwarder creates a new testForwarder and starts it.
// Call LocalAddr after creation to get the address after creation.
func newTestForwarder() (*testForwarder, error) {
	f := new(testForwarder)
	var err error
	f.conn, err = net.ListenPacket("udp", ":0")
	if err != nil {
		return nil, err
	}

	mux := dns.NewServeMux()
	mux.HandleFunc(".", f.forward)
	f.server = &dns.Server{
		Handler:    mux,
		PacketConn: f.conn,
	}

	log.Printf("[DEBUG] new test forwarding DNS server registered at %s", f.conn.LocalAddr())
	f.start()
	return f, nil
}

// LocalAddr returns the local address of the test forwarder as a
// string.
func (f *testForwarder) LocalAddr() string {
	return f.conn.LocalAddr().String()
}

// Called returns true if the testForwarder has seen traffic.
func (f *testForwarder) Called() bool {
	return f.called
}

// start starts the DNS server in a new goroutine. Errors operating
// the server result in a panic.
//
// Make sure to call Shutdown when finished as to not leak the
// goroutine started by this function.
func (f *testForwarder) start() {
	go func() {
		log.Printf("[DEBUG] starting test forwarding DNS server on %s", f.conn.LocalAddr())
		if err := f.server.ActivateAndServe(); err != nil {
			panic(err)
		}

		log.Printf("[DEBUG] test forwarding DNS server on %s shutting down", f.conn.LocalAddr())
	}()
}

// Shutdown shuts down the DNS server and should be called when the
// test is complete. It panics on errors.
func (f *testForwarder) Shutdown() {
	if err := f.server.Shutdown(); err != nil {
		panic(err)
	}
}

// forward sends DNS traffic to the forwarding address
func (f *testForwarder) forward(r dns.ResponseWriter, msg *dns.Msg) {
	f.called = true
	defer r.Close()
	log.Printf("[DEBUG] query from %s: \n  %s", r.RemoteAddr(), f.questions(msg.Question))

	c := new(dns.Client)
	resp, rtt, err := c.Exchange(msg, forwardingAddress)
	if err != nil {
		log.Printf("[DEBUG] error forwarding test DNS request: %s", err)
	}

	log.Printf("[DEBUG] reply from %s (RTT=%d): \n  %s", forwardingAddress, rtt, f.answers(resp.Answer))
	if err := r.WriteMsg(resp); err != nil {
		log.Printf("[DEBUG] error relaying back forwarded DNS response: %s", err)
	}
}

func (f *testForwarder) Check() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if !f.Called() {
			return fmt.Errorf("no DNS records were sent through %s", f.conn.LocalAddr())
		}

		return nil
	}
}

func (f *testForwarder) questions(qs []dns.Question) string {
	var result []string
	for _, q := range qs {
		result = append(result, q.String())
	}

	return strings.Join(result, "\n  ")
}

func (f *testForwarder) answers(rrs []dns.RR) string {
	var result []string
	for _, rr := range rrs {
		result = append(result, rr.String())
	}

	return strings.Join(result, "\n  ")
}

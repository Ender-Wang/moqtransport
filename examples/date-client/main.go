package main

import (
	"context"
	"crypto/tls"
	"flag"
	"io"
	"log"

	"github.com/mengelbart/moqtransport"
	"github.com/mengelbart/moqtransport/quicmoq"
	"github.com/mengelbart/moqtransport/webtransportmoq"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/webtransport-go"
)

func main() {
	addr := flag.String("addr", "localhost:8080", "address to connect to")
	wt := flag.Bool("webtransport", false, "Use webtransport instead of QUIC")
	namespace := flag.String("namespace", "clock", "Namespace to subscribe to")
	trackname := flag.String("trackname", "second", "Track to subscribe to")
	flag.Parse()

	if err := run(context.Background(), *addr, *wt, *namespace, *trackname); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, addr string, wt bool, namespace, trackname string) error {
	var conn moqtransport.Connection
	var err error
	if wt {
		conn, err = dialWebTransport(ctx, addr)
	} else {
		conn, err = dialQUIC(ctx, addr)
	}
	if err != nil {
		return err
	}
	announcementWaitCh := make(chan *moqtransport.Announcement)
	session := &moqtransport.Session{
		Conn:            conn,
		EnableDatagrams: true,
		LocalRole:       moqtransport.RoleSubscriber,
		AnnouncementHandler: moqtransport.AnnouncementHandlerFunc(func(s *moqtransport.Session, a *moqtransport.Announcement, arw moqtransport.AnnouncementResponseWriter) {
			if a.Namespace() == "clock" {
				arw.Accept()
				announcementWaitCh <- a
				return
			}
			arw.Reject(0, "invalid namespace")
		}),
	}
	if err = session.RunClient(); err != nil {
		return err
	}
	defer session.Close()

	a := <-announcementWaitCh
	log.Printf("got Announcement: %v\n", a)
	log.Println("subscribing")
	rs, err := session.Subscribe(context.Background(), 0, 0, namespace, trackname, "")
	if err != nil {
		return err
	}
	log.Println("got subscription")
	for {
		o, err := rs.ReadObject(ctx)
		if err != nil {
			if err == io.EOF {
				log.Printf("got last object")
				return nil
			}
			return err
		}
		log.Printf("got object: %v\n", string(o.Payload))
	}
}

func dialQUIC(ctx context.Context, addr string) (moqtransport.Connection, error) {
	conn, err := quic.DialAddr(ctx, addr, &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"moq-00"},
	}, &quic.Config{
		EnableDatagrams: true,
	})
	if err != nil {
		return nil, err
	}
	return quicmoq.New(conn), nil
}

func dialWebTransport(ctx context.Context, addr string) (moqtransport.Connection, error) {
	dialer := webtransport.Dialer{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	_, session, err := dialer.Dial(ctx, addr, nil)
	if err != nil {
		return nil, err
	}
	return webtransportmoq.New(session), nil
}

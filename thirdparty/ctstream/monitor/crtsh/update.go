package crtsh

import (
	"github.com/anonymous-author-00000/ctstream/core"
)

func null() *CrtshCTClient { return nil }

func SelectByDomain(
	domain string,
	clients *core.CTClients[*CrtshCTClient],
) (*CrtshCTClient, int, error) {
	client, i, err := core.SelectByDomain(domain, clients, null)
	return client, i, err
}

func AddByDomain(
	domain string,
	clients *core.CTClients[*CrtshCTClient],
) (*CrtshCTClient, int, error) {
	client, i, err := core.AddByDomain(domain, NewCTClient, clients, null)
	return client, i, err
}

func DelByDomain(
	domain string,
	clients *core.CTClients[*CrtshCTClient],
) (*CrtshCTClient, int, error) {
	client, i, err := core.DelByDomain(domain, clients, null)
	return client, i, err
}

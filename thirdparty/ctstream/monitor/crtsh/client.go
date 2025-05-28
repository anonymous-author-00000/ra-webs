package crtsh

import (
	"github.com/anonymous-author-00000/crtsh"
	"github.com/anonymous-author-00000/ctstream/core"
	"github.com/anonymous-author-00000/ctstream/utils"
)

type CrtshCTClient struct {
	ID     int
	Domain string
}

type CrtshCTParams struct {
	ID     int
	Client *CrtshCTClient
}

func NewCTClient(domain string) (*CrtshCTClient, error) {
	return &CrtshCTClient{
		ID:     0,
		Domain: domain,
	}, nil
}

func (client *CrtshCTClient) Init() error {
	client.ID = 0
	return nil
}

func (client *CrtshCTClient) Next(callback core.Callback) {
	entries, err := crtsh.Fetch(client.Domain, crtsh.EXCLUDE_EXPIRED)
	if err != nil {
		callback(nil, 0, &CrtshCTParams{}, err)
		return
	}

	for i, entry := range entries {
		if entry.ID <= client.ID {
			continue
		}

		c, err := utils.ReformatCertificate(entry.Certificate)
		callback(c, i, &CrtshCTParams{
			ID:     entries[i].ID,
			Client: client,
		}, err)

		client.ID = entry.ID
	}
}

func (client *CrtshCTClient) GetDomain() string {
	return client.Domain
}

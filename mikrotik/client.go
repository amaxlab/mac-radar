package mikrotik

import (
	"github.com/amaxlab/go-lib/log"
	"github.com/go-routeros/routeros"
)

type Configuration struct {
	Address  string
	Username string
	Password string
}

type Client struct {
	Configuration Configuration
}

func NewClient(c Configuration) *Client {
	return &Client{Configuration: c}
}

func (c *Client) GetMacAddress() []string {
	routerOsClient, err := routeros.Dial(c.Configuration.Address, c.Configuration.Username, c.Configuration.Password)
	if err != nil {
		log.Error.Printf("Error on connect mikrotik router -> %s", err)
		return nil
	}
	defer routerOsClient.Close()

	r, err := routerOsClient.Run("/ip/arp/print", "?complete=true", "=.proplist=mac-address")
	if err != nil {
		log.Error.Printf("Error on run get arp mikrotik command -> %s", err)
		return nil
	}

	m := make([]string, 0)
	for _, re := range r.Re {
		m = append(m, re.Map["mac-address"])
	}
	return m
}

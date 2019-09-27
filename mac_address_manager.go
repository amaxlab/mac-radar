package main

import (
	"github.com/amaxlab/go-lib/log"
	"github.com/amaxlab/mac-radar/mikrotik"
	"strings"
	"time"
)

type MacAddress struct {
	Mac         string    `json:"mac"`
	Online      bool      `json:"online"`
	OnlineTime  time.Time `json:"online_time"`
	OfflineTime time.Time `json:"offline_time"`
}

type MacAddressManager struct {
	MacAddress []MacAddress `json:"mac_address"`
	Client *mikrotik.Client
}

func NewMacAddressManager(c *mikrotik.Client) *MacAddressManager {
	return &MacAddressManager{Client: c}
}

func (m *MacAddressManager) Watch() {
	for {
		m.Update()
		time.Sleep(time.Minute * 1)
	}
}

func (m *MacAddressManager) Update() {
	log.Debug.Printf("Update mac-address")
	address := m.Client.GetMacAddress()
	for _, item := range address {
		m.SetOnline(getValueByMacAddress(item))
	}
	m.UpdateOffline(address)
}

func (m *MacAddressManager) SetOnline(mac string) {
	for id, item := range m.MacAddress {
		if item.Mac == mac {
			if item.Online == false {
				log.Info.Printf("Mac-address %s is online", mac)
				m.MacAddress[id].Online = true
				m.MacAddress[id].OnlineTime = time.Now()
			}
			return
		}
	}

	log.Info.Printf("Mac-address %s is online", mac)
	m.MacAddress = append(m.MacAddress, MacAddress{Mac: mac, Online: true, OnlineTime: time.Now()})
}

func (m MacAddressManager) GetByMac(mac string) *MacAddress {
	for id := range m.MacAddress {
		if m.MacAddress[id].Mac == mac {
			return &m.MacAddress[id]
		}
	}
	return nil
}

func (m *MacAddressManager) UpdateOffline(address []string) {
	log.Debug.Printf("Update offline mac-address")
	for id, item := range m.MacAddress {
		found := false
		for _, mac := range address {
			if item.Mac == getValueByMacAddress(mac) {
				found = true
				continue
			}
		}

		if !found {
			if item.Online == true {
				log.Info.Printf("Mac-address %s is offline", item.Mac)
				m.MacAddress[id].Online = false
				m.MacAddress[id].OfflineTime = time.Now()
			}
		}
	}
}

func getValueByMacAddress(mac string) string {
	return strings.Replace(strings.ToLower(mac), ":", "", -1)
}

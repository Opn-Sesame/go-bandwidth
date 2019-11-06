package bandwidth

import (
	"encoding/xml"
	"time"
)

// AssociatedSipPeer returns an associated sip-peer (aka location) with the app.
type AssociatedSipPeer struct {
	// SiteId is the ID of the site.
	SiteID string `xml:"SiteId"`
	// SiteName is the name of the site.
	SiteName string
	// PeerId is the ID of the peer. This is used alongside side-id to retrieve the actual numbers.
	PeerID string `xml:"PeerId"`
}

// AssociatedSipPeers is a list of associated sip peers.
type AssociatedSipPeers struct {
	Associated []AssociatedSipPeer `xml:"AssociatedSipPeer"`
}

// AssociatedSipPeersResponse struct
type AssociatedSipPeersResponse struct {
	Peers AssociatedSipPeers `xml:"AssociatedSipPeers"`
}

// TelephoneNumber is the phone number.
type TelephoneNumber struct {
	FullNumber string
}

// SipPeerTelephoneNumbers is a collection of phone numbers.
type SipPeerTelephoneNumbers struct {
	Numbers []TelephoneNumber `xml:"SipPeerTelephoneNumber"`
}

// SipPeerTelephoneNumbersResponse is the response to fetching sip-peers.
type SipPeerTelephoneNumbersResponse struct {
	Peers SipPeerTelephoneNumbers `xml:"SipPeerTelephoneNumbers"`
}

type OrderRequest struct {
	Name           string
	SiteID         string `xml:"SiteId"`
	PeerID         string `xml:"PeerId"`
	PartialAllowed bool
}

type AreaCodeOrder struct {
	Name                       string `xml:",omitempty"`
	SiteID                     string `xml:"SiteId"`
	PeerID                     string `xml:"PeerId"`
	PartialAllowed             bool
	CustomerOrderID            string `xml:"CustomerOrderId,omitempty"`
	AreaCodeSearchAndOrderType AreaCodeSearchAndOrderType
}

/*
<OrderResponse><Order><OrderCreateDate>2019-11-05T13:48:43.238Z</OrderCreateDate><PeerId>596135</PeerId><BackOrderRequested>false</BackOrderRequested><id>983d1b3a-0698-4c47-bb87-c455b3fbf4ca</id><AreaCodeSearchAndOrderType><AreaCode>734</AreaCode><Quantity>1</Quantity></AreaCodeSearchAndOrderType><PartialAllowed>true</PartialAllowed><SiteId>27161</SiteId></Order><OrderStatus>RECEIVED</OrderStatus></OrderResponse>

*/
type OrderResponseOrder struct {
	OrderCreated               time.Time
	PeerID                     string `xml:"PeerId"`
	BackOrderRequested         bool
	ID                         string `xml:"id"`
	AreaCodeSearchAndOrderType AreaCodeSearchAndOrderType
	PartiallyAllowed           bool
	SiteID                     string `xml:"SiteId"`
}

type OrderResponse struct {
	Order       OrderResponseOrder `xml:"Order"`
	OrderStatus string
	// Only relevant if OrderStatus == COMPLETE
	CompletedNumbers CompletedNumbers
}

type CompletedNumbers struct {
	TelephoneNumbers []TelephoneNumber `xml:"TelephoneNumber"`
}

type AreaCodeSearchAndOrderType struct {
	AreaCode string
	Quantity int
}

type AreaCodeRequest struct {
	XMLName xml.Name `xml:"Order"`
	AreaCodeOrder
}

type TollFreeWildCharSearchAndOrderType struct {
	TollFreeWildCardPattern string
	Quantity                int
}

type TollFreeOrder struct {
	Name                               string `xml:",omitempty"`
	SiteID                             string `xml:"SiteId"`
	PeerID                             string `xml:"PeerId"`
	PartialAllowed                     bool
	CustomerOrderID                    string `xml:"CustomerOrderId,omitempty"`
	TollFreeWildCharSearchAndOrderType TollFreeWildCharSearchAndOrderType
}
type TollFreeOrderRequest struct {
	XMLName xml.Name `xml:"Order"`
	TollFreeOrder
}

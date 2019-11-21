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

type SipPeer struct {
	PeerName      string
	Description   string
	IsDefaultPeer bool
}

type HttpSettings struct {
	ProxyPeerId int `xml:",omitempty"`
}

type SipPeerSmsFeatureSettings struct {
	TollFree    bool
	ShortCode   bool
	A2pLongCode string // DefaultOff?
	Protocol    string // HTTP
	Zone1       bool
	Zone2       bool
	Zone3       bool
	Zone4       bool
	Zone5       bool
}
type SipPeerSmsFeature struct {
	SipPeerSmsFeatureSettings SipPeerSmsFeatureSettings
	HttpSettings              HttpSettings
}

type SipPeerSmsFeatureResponse struct {
	SipPeerSmsFeature SipPeerSmsFeature
}

type MmsSettings struct {
	Protocol string
}
type HTTPProtocol struct {
	HttpSettings HttpSettings
}
type Protocols struct {
	HTTP HTTPProtocol
}
type MmsFeature struct {
	MmsSettings MmsSettings
	Protocols   Protocols
}

type MmsFeatureResponse struct {
	MmsFeature MmsFeature
}

type ApplicationsSettings struct {
	HttpMessagingV2AppId string
}

type ApplicationsSettingsResponse struct {
	ApplicationsSettings ApplicationsSettings
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

type TelephoneNumberList struct {
	TelephoneNumber []string
}
type SearchResult struct {
	ResultCount         int
	TelephoneNumberList TelephoneNumberList
}
type DisconnectTelephoneNumberOrder struct {
	Name                               string `xml:"name,omitempty"`
	DisconnectTelephoneNumberOrderType DisconnectTelephoneNumberOrderType
}
type DisconnectTelephoneNumberOrderType struct {
	TelephoneNumberList TelephoneNumberList
}
type DisconnectTelephoneNumberOrderResponse struct {
	OrderRequest DisconnectOrderRequest
	OrderStatus  string
}

type DisconnectOrderRequest struct {
	XMLName                            xml.Name `xml:"orderRequest"`
	OrderCreateDate                    time.Time
	ID                                 string `xml:"id"`
	DisconnectTelephoneNumberOrderType DisconnectTelephoneNumberOrderType
	DisconnectMode                     string
}

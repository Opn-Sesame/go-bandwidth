package bandwidth

import (
	"net/http"
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

// GetAssociatedPeers returns the associated sippeers (aka locations) for the application.
func (c *Client) GetAssociatedPeers(applicationID string) (*AssociatedSipPeersResponse, error) {
	path := c.AccountsEndpoint + "/applications/" + applicationID + "/associatedsippeers"
	result, _, err := c.makeAccountsRequest(http.MethodGet, path, &AssociatedSipPeersResponse{})
	if err != nil {
		return nil, err
	}
	return result.(*AssociatedSipPeersResponse), nil
}

// SipPeerTelephoneNumber is the phone number.
type SipPeerTelephoneNumber struct {
	FullNumber string
}

// SipPeerTelephoneNumbers is a collection of phone numbers.
type SipPeerTelephoneNumbers struct {
	Numbers []SipPeerTelephoneNumber `xml:"SipPeerTelephoneNumber"`
}

// SipPeerTelephoneNumbersResponse is the response to fetching sip-peers.
type SipPeerTelephoneNumbersResponse struct {
	Peers SipPeerTelephoneNumbers `xml:"SipPeerTelephoneNumbers"`
}

// GetTollFreeNumbers returns the toll-free numbers associated with the site.
func (c *Client) GetTollFreeNumbers(siteID, peerID string) (*SipPeerTelephoneNumbersResponse, error) {
	path := c.AccountsEndpoint + "/sites/" + siteID + "/sippeers/" + peerID + "/tns"
	result, _, err := c.makeAccountsRequest(http.MethodGet, path, &SipPeerTelephoneNumbersResponse{})
	if err != nil {
		return nil, err
	}
	return result.(*SipPeerTelephoneNumbersResponse), nil
}

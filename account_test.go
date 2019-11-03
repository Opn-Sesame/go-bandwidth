package bandwidth

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetAssociatedPeers(t *testing.T) {
	siteID := "12345"
	peerID := "123123"
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: fmt.Sprintf("%s%s/applications/%s/associatedsippeers", accountsPath, testAccountID, testApplicationID),
		Method:       http.MethodGet,
		ContentToSend: fmt.Sprintf(`
		<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
			<AssociatedSipPeersResponse>
				<AssociatedSipPeers>
					<AssociatedSipPeer>
						<SiteId>%s</SiteId>
						<SiteName>Test-Subaccount-1</SiteName>
						<PeerId>%s</PeerId>
						<PeerName>default</PeerName>
					</AssociatedSipPeer>
				</AssociatedSipPeers>
			</AssociatedSipPeersResponse>`, siteID, peerID)}})
	defer server.Close()
	result, err := api.GetAssociatedPeers(context.Background(), testApplicationID)
	if err != nil {
		t.Errorf("Failed call of GetAssociatedPeers(): %v", err)
		return
	}
	expect(t, len(result.Peers.Associated), 1)
	expect(t, result.Peers.Associated[0].SiteID, siteID)
	expect(t, result.Peers.Associated[0].PeerID, peerID)
}

func TestGetAssociatedPeersFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     fmt.Sprintf("%s%s/applications/%s/associatedsippeers", accountsPath, testAccountID, testApplicationID),
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetAssociatedPeers(context.Background(), testApplicationID) })
}

func TestGetTollFreeNumbers(t *testing.T) {
	siteID := "12345"
	peerID := "123123"
	number := "1234567890"
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: fmt.Sprintf("%s%s/sites/%s/sippeers/%s/tns", accountsPath, testAccountID, siteID, peerID),
		Method:       http.MethodGet,
		ContentToSend: fmt.Sprintf(`
		<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
			<SipPeerTelephoneNumbersResponse>
				<Links><first>Link=&lt;https://dashboard.bandwidth.com:443/v1.0/accounts/%v/sites/%v/sippeers/%v/tns?page=1&amp;size=50000&gt;;rel="first";</first>
				</Links>
				<SipPeerTelephoneNumbers>
					<SipPeerTelephoneNumber>
						<FullNumber>%s</FullNumber>
					</SipPeerTelephoneNumber>
				</SipPeerTelephoneNumbers>
			</SipPeerTelephoneNumbersResponse>`, testAccountID, siteID, peerID, number)}})
	defer server.Close()
	result, err := api.GetTollFreeNumbers(context.Background(), siteID, peerID)
	if err != nil {
		t.Errorf("Failed call of GetAssociatedPeers(): %v", err)
		return
	}
	expect(t, len(result.Peers.Numbers), 1)
	expect(t, result.Peers.Numbers[0].FullNumber, number)
}

func TestGetTollFreeNumbersFail(t *testing.T) {
	siteID := "12345"
	peerID := "123123"
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     fmt.Sprintf("%s%s/sites/%s/sippeers/%s/tns", accountsPath, testAccountID, siteID, peerID),
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetTollFreeNumbers(context.Background(), siteID, peerID) })
}

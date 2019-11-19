package bandwidth

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"testing"
)

func TestEnableMMS(t *testing.T) {
	siteID := "12345"
	peerID := "123123"
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     fmt.Sprintf("%s%s/sites/%s/sippeers/%s/products/messaging/features/mms", accountsPath, testAccountID, siteID, peerID),
		Method:           http.MethodPost,
		EstimatedContent: fmt.Sprintf(`<MmsFeature><MmsSettings><Protocol>HTTP</Protocol></MmsSettings><Protocols><HTTP><HttpSettings></HttpSettings></HTTP></Protocols></MmsFeature>`),
		ContentToSend: fmt.Sprintf(`
		<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
		<MmsFeatureResponse>
			<MmsFeature>
				<MmsSettings>
					<Protocol>HTTP</Protocol>
				</MmsSettings>
				<Protocols>
					<HTTP>
						<HttpSettings>
							<ProxyPeerId>1234</ProxyPeerId>
						</HttpSettings>
					</HTTP>
				</Protocols>
			</MmsFeature>
		</MmsFeatureResponse>`)}})
	defer server.Close()
	result, err := api.EnableMMS(context.Background(), siteID, peerID)
	if err != nil {
		t.Errorf("Failed call of EnableMMS(): %v", err)
		return
	}
	expect(t, result.MmsFeature.Protocols.HTTP.HttpSettings.ProxyPeerId, 1234)
}

func TestEnableSMS(t *testing.T) {
	siteID := "12345"
	peerID := "123123"
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     fmt.Sprintf("%s%s/sites/%s/sippeers/%s/products/messaging/features/sms", accountsPath, testAccountID, siteID, peerID),
		Method:           http.MethodPost,
		EstimatedContent: fmt.Sprintf(`<SipPeerSmsFeature><SipPeerSmsFeatureSettings><TollFree>true</TollFree><ShortCode>true</ShortCode><A2pLongCode>DefaultOff</A2pLongCode><Protocol>HTTP</Protocol><Zone1>true</Zone1><Zone2>false</Zone2><Zone3>false</Zone3><Zone4>false</Zone4><Zone5>false</Zone5></SipPeerSmsFeatureSettings><HttpSettings></HttpSettings></SipPeerSmsFeature>`),
		ContentToSend: fmt.Sprintf(`
		<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
		<SipPeerSmsFeatureResponse>
			<SipPeerSmsFeature>
				<SipPeerSmsFeatureSettings>
					<TollFree>true</TollFree>
					<ShortCode>true</ShortCode>
					<A2pLongCode>DefaultOff</A2pLongCode>
					<Protocol>HTTP</Protocol>
					<Zone1>true</Zone1>
					<Zone2>false</Zone2>
					<Zone3>false</Zone3>
					<Zone4>false</Zone4>
					<Zone5>false</Zone5>
				</SipPeerSmsFeatureSettings>
				<HttpSettings>
					<ProxyPeerId>1234</ProxyPeerId>
				</HttpSettings>
			</SipPeerSmsFeature>
		</SipPeerSmsFeatureResponse>`)}})
	defer server.Close()
	result, err := api.EnableSMS(context.Background(), siteID, peerID)
	if err != nil {
		t.Errorf("Failed call of EnableSMS(): %v", err)
		return
	}
	expect(t, result.SipPeerSmsFeature.SipPeerSmsFeatureSettings.TollFree, true)
}

func TestAssociateApplication(t *testing.T) {
	siteID := "12345"
	peerID := "123123"
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     fmt.Sprintf("%s%s/sites/%s/sippeers/%s/products/messaging/applicationSettings", accountsPath, testAccountID, siteID, peerID),
		Method:           http.MethodPut,
		EstimatedContent: fmt.Sprintf(`<ApplicationsSettings><HttpMessagingV2AppId>%s</HttpMessagingV2AppId></ApplicationsSettings>`, testApplicationID),
		ContentToSend: fmt.Sprintf(`
		<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
		<ApplicationsSettingsResponse>
			<ApplicationsSettings>
				<HttpMessagingV2AppId>%s</HttpMessagingV2AppId>
			</ApplicationsSettings>
		</ApplicationsSettingsResponse>`, testApplicationID)}})
	defer server.Close()
	result, err := api.AssociateApplication(context.Background(), siteID, peerID, testApplicationID)
	if err != nil {
		t.Errorf("Failed call of AssociateApplication(): %v", err)
		return
	}
	expect(t, result.ApplicationsSettings.HttpMessagingV2AppId, testApplicationID)
}

func TestCreatePeer(t *testing.T) {
	siteID := "12345"
	peerID := "678"
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: fmt.Sprintf("%s%s/sites/%s/sippeers", accountsPath, testAccountID, siteID),
		Method:       http.MethodPost,
		HeadersToSend: map[string]string{
			"Location": fmt.Sprintf("https://dashboard.bandwidth.com:443/v1.0/accounts/%s/sites/%s/sippeers/%s",
				testAccountID,
				siteID,
				peerID),
		},
		ContentToSend: "",
	}})
	defer server.Close()
	result, err := api.CreatePeer(context.Background(), testAccountID, siteID, "test peer", false)
	if err != nil {
		t.Errorf("Failed call of CreatePeer(): %v", err)
		return
	}
	expect(t, result, peerID)
}

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

func TestGetNumbers(t *testing.T) {
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
	result, err := api.GetNumbers(context.Background(), siteID, peerID)
	if err != nil {
		t.Errorf("Failed call of GetNumbers(): %v", err)
		return
	}
	expect(t, len(result.Peers.Numbers), 1)
	expect(t, result.Peers.Numbers[0].FullNumber, number)
}

func TestGetNumbersFail(t *testing.T) {
	siteID := "12345"
	peerID := "123123"
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     fmt.Sprintf("%s%s/sites/%s/sippeers/%s/tns", accountsPath, testAccountID, siteID, peerID),
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetNumbers(context.Background(), siteID, peerID) })
}

func TestOrderNumbersByAreaCode(t *testing.T) {
	siteID := "12345"
	peerID := "123123"
	id := "1-2-3-4"
	areaCode := "734"
	quantity := 1
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     fmt.Sprintf("%s%s/orders", accountsPath, testAccountID),
		Method:           http.MethodPost,
		EstimatedContent: fmt.Sprintf("<Order><SiteId>%s</SiteId><PeerId>%s</PeerId><PartialAllowed>false</PartialAllowed><AreaCodeSearchAndOrderType><AreaCode>%s</AreaCode><Quantity>%d</Quantity></AreaCodeSearchAndOrderType></Order>", siteID, peerID, areaCode, quantity),
		ContentToSend: fmt.Sprintf(`
		<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
		<OrderResponse>
			<Order>
				<OrderCreateDate>2019-11-05T13:48:43.238Z</OrderCreateDate>
				<PeerId>%s</PeerId>
				<BackOrderRequested>false</BackOrderRequested>
				<id>%s</id>
				<AreaCodeSearchAndOrderType>
					<AreaCode>%s</AreaCode>
					<Quantity>%d</Quantity>
				</AreaCodeSearchAndOrderType>
				<PartialAllowed>true</PartialAllowed>
				<SiteId>%s</SiteId>
			</Order>
			<OrderStatus>RECEIVED</OrderStatus>
		</OrderResponse>`, peerID, id, areaCode, quantity, siteID)}})
	defer server.Close()
	result, err := api.OrderNumbersByAreaCode(context.Background(), siteID, peerID, areaCode, quantity)
	if err != nil {
		t.Errorf("Failed call of OrderNumbersByAreaCode(): %v", err)
		return
	}
	expect(t, result.OrderStatus, "RECEIVED")
}

func TestOrderTollFreeNumbers(t *testing.T) {
	siteID := "12345"
	peerID := "123123"
	id := "1-2-3-4"
	mask := "8**"
	quantity := 1
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     fmt.Sprintf("%s%s/orders", accountsPath, testAccountID),
		Method:           http.MethodPost,
		EstimatedContent: fmt.Sprintf("<Order><SiteId>%s</SiteId><PeerId>%s</PeerId><PartialAllowed>false</PartialAllowed><TollFreeWildCharSearchAndOrderType><TollFreeWildCardPattern>%s</TollFreeWildCardPattern><Quantity>%d</Quantity></TollFreeWildCharSearchAndOrderType></Order>", siteID, peerID, mask, quantity),
		ContentToSend: fmt.Sprintf(`
		<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
		<OrderResponse>
			<Order>
				<OrderCreateDate>2019-11-05T13:48:43.238Z</OrderCreateDate>
				<PeerId>%s</PeerId>
				<BackOrderRequested>false</BackOrderRequested>
				<id>%s</id>
				<TollFreeWildCharSearchAndOrderType>
					<TollFreeWildCardPattern>%s</TollFreeWildCardPattern>
					<Quantity>%d</Quantity>
				</TollFreeWildCharSearchAndOrderType>
				<PartialAllowed>true</PartialAllowed>
				<SiteId>%s</SiteId>
			</Order>
			<OrderStatus>RECEIVED</OrderStatus>
		</OrderResponse>`, peerID, id, mask, quantity, siteID)}})
	defer server.Close()
	result, err := api.OrderTollFreeNumbers(context.Background(), siteID, peerID, mask, quantity)
	if err != nil {
		t.Errorf("Failed call of OrderTollFreeNumbers(): %v", err)
		return
	}
	expect(t, result.OrderStatus, "RECEIVED")
}

func TestGetTollFreeOrder(t *testing.T) {
	siteID := "12345"
	peerID := "123123"
	id := "1-2-3-4"
	mask := "8**"
	number := "8441231234"
	quantity := 1
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: fmt.Sprintf("%s%s/orders/%s", accountsPath, testAccountID, id),
		Method:       http.MethodGet,
		ContentToSend: fmt.Sprintf(`
		<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
		<OrderResponse>
			<CompletedQuantity>1</CompletedQuantity>
			<CreatedByUser>foo@bar.com</CreatedByUser>
			<LastModifiedDate>2019-11-05T15:04:50.531Z</LastModifiedDate>
			<OrderCompleteDate>2019-11-05T15:04:50.531Z</OrderCompleteDate>
			<Order>
				<OrderCreateDate>2019-11-05T15:04:50.327Z</OrderCreateDate>
				<PeerId>%s</PeerId>
				<BackOrderRequested>false</BackOrderRequested>
				<TollFreeWildCharSearchAndOrderType>
					<Quantity>%d</Quantity>
					<TollFreeWildCardPattern>%s</TollFreeWildCardPattern>
				</TollFreeWildCharSearchAndOrderType>
				<PartialAllowed>true</PartialAllowed>
				<SiteId>%s</SiteId>
			</Order>
			<OrderStatus>COMPLETE</OrderStatus>
			<CompletedNumbers>
				<TelephoneNumber>
					<FullNumber>%s</FullNumber>
				</TelephoneNumber>
			</CompletedNumbers>
			<Summary>1 number ordered in (844)</Summary>
			<FailedQuantity>0</FailedQuantity>
		</OrderResponse>`, peerID, quantity, mask, siteID, number)}})
	defer server.Close()
	result, err := api.GetOrder(context.Background(), id)
	if err != nil {
		t.Errorf("Failed call of GetOrder(): %v", err)
		return
	}
	expect(t, result.OrderStatus, "COMPLETE")
	expect(t, result.Order.PeerID, peerID)
	expect(t, result.CompletedNumbers.TelephoneNumbers[0].FullNumber, number)

}

func TestGetOrderByAreaCode(t *testing.T) {
	siteID := "12345"
	peerID := "123123"
	id := "1-2-3-4"
	areaCode := "734"
	quantity := 1
	number := "7341231234"
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: fmt.Sprintf("%s%s/orders/%s", accountsPath, testAccountID, id),
		Method:       http.MethodGet,
		ContentToSend: fmt.Sprintf(`
		<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
		<OrderResponse>
			<CompletedQuantity>1</CompletedQuantity>
			<CreatedByUser>foo@bar.com</CreatedByUser>
			<LastModifiedDate>2019-11-05T13:48:43.379Z</LastModifiedDate>
			<OrderCompleteDate>2019-11-05T13:48:43.379Z</OrderCompleteDate>
			<Order>
				<OrderCreateDate>2019-11-05T13:48:43.238Z</OrderCreateDate>
				<PeerId>%s</PeerId>
				<BackOrderRequested>false</BackOrderRequested>
				<AreaCodeSearchAndOrderType>
					<AreaCode>%s</AreaCode>
					<Quantity>%d</Quantity>
				</AreaCodeSearchAndOrderType>
				<PartialAllowed>true</PartialAllowed>
				<SiteId>%s</SiteId>
			</Order>
			<OrderStatus>COMPLETE</OrderStatus>
			<CompletedNumbers>
				<TelephoneNumber>
					<FullNumber>%s</FullNumber>
				</TelephoneNumber>
			</CompletedNumbers>
			<Summary>1 number ordered in (734)</Summary>
			<FailedQuantity>0</FailedQuantity>
		</OrderResponse>`, peerID, areaCode, quantity, siteID, number)}})
	defer server.Close()
	result, err := api.GetOrder(context.Background(), id)
	if err != nil {
		t.Errorf("Failed call of GetOrder(): %v", err)
		return
	}
	expect(t, result.OrderStatus, "COMPLETE")
	expect(t, result.Order.PeerID, peerID)
	expect(t, result.CompletedNumbers.TelephoneNumbers[0].FullNumber, number)
}

func TestSearchNumbersByAreaCode(t *testing.T) {
	areaCode := "510"
	numbers := []string{"5101231234", "5101234567"}
	sort.Strings(numbers)
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: fmt.Sprintf("%s%s/availableNumbers?areaCode=%s&quantity=%d", accountsPath, testAccountID, areaCode, len(numbers)),
		Method:       http.MethodGet,
		ContentToSend: fmt.Sprintf(`
		<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
		<SearchResult>
			<ResultCount>%d</ResultCount>
			<TelephoneNumberList>
				<TelephoneNumber>%s</TelephoneNumber>
				<TelephoneNumber>%s</TelephoneNumber>
			</TelephoneNumberList>
		</SearchResult>`, len(numbers), numbers[0], numbers[1])}})
	defer server.Close()
	result, err := api.SearchNumbersByAreaCode(context.Background(), areaCode, len(numbers))
	if err != nil {
		t.Errorf("Failed call of SearchNumbersByAreaCode(): %v", err)
		return
	}
	expect(t, result.ResultCount, len(numbers))
	expect(t, len(result.TelephoneNumberList.TelephoneNumber), len(numbers))
	var found []string
	for _, el := range result.TelephoneNumberList.TelephoneNumber {
		found = append(found, el)
	}
	sort.Strings(found)
	expect(t, found, numbers)
}

func TestSearchTollFreeNumbers(t *testing.T) {
	mask := "8**"
	numbers := []string{"8441231234", "8441234567"}
	sort.Strings(numbers)
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: fmt.Sprintf("%s%s/availableNumbers?quantity=%d&tollFreeWildCardPattern=%s", accountsPath, testAccountID, len(numbers), url.QueryEscape(mask)),
		Method:       http.MethodGet,
		ContentToSend: fmt.Sprintf(`
		<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
		<SearchResult>
			<ResultCount>%d</ResultCount>
			<TelephoneNumberList>
				<TelephoneNumber>%s</TelephoneNumber>
				<TelephoneNumber>%s</TelephoneNumber>
			</TelephoneNumberList>
		</SearchResult>`, len(numbers), numbers[0], numbers[1])}})
	defer server.Close()
	result, err := api.SearchTollFreeNumbers(context.Background(), mask, len(numbers))
	if err != nil {
		t.Errorf("Failed call of SearchTollFreeNumbers(): %v", err)
		return
	}
	expect(t, result.ResultCount, len(numbers))
	expect(t, len(result.TelephoneNumberList.TelephoneNumber), len(numbers))
	var found []string
	for _, el := range result.TelephoneNumberList.TelephoneNumber {
		found = append(found, el)
	}
	sort.Strings(found)
	expect(t, found, numbers)
}

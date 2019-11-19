package bandwidth

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// CreatePeer creates the Sip peer and returns its ID.
func (c *Client) CreatePeer(ctx context.Context, applicationID, siteID, peerName string, isDefault bool) (string, error) {
	path := c.AccountsEndpoint + "/sites/" + siteID + "/sippeers"
	sipPeer := SipPeer{
		PeerName:      peerName,
		IsDefaultPeer: isDefault,
	}
	_, headers, err := c.makeAccountsRequest(ctx, http.MethodPost, path, nil, &sipPeer)
	if err != nil {
		return "", err
	}
	splitted := strings.Split(headers.Get("Location"), "/sites/"+siteID+"/sippeers/")
	if len(splitted) != 2 {
		return "", fmt.Errorf("unknown peer ID: %v", headers.Get("Location"))
	}
	return splitted[1], nil
}

// EnableSMS enables SMS
func (c *Client) EnableSMS(ctx context.Context, siteID, peerID string) (*SipPeerSmsFeatureResponse, error) {
	path := c.AccountsEndpoint + "/sites/" + siteID + "/sippeers/" + peerID + "/products/messaging/features/sms"
	feature := SipPeerSmsFeature{
		SipPeerSmsFeatureSettings: SipPeerSmsFeatureSettings{
			TollFree:    true,
			ShortCode:   false,
			Protocol:    "HTTP",
			Zone1:       true,
			A2pLongCode: "DefaultOff",
		},
	}

	result, _, err := c.makeAccountsRequest(ctx, http.MethodPost, path, &SipPeerSmsFeatureResponse{}, &feature)
	if err != nil {
		return nil, err
	}
	return result.(*SipPeerSmsFeatureResponse), nil
}

// EnableMMS enables MMS
func (c *Client) EnableMMS(ctx context.Context, siteID, peerID string) (*MmsFeatureResponse, error) {
	path := c.AccountsEndpoint + "/sites/" + siteID + "/sippeers/" + peerID + "/products/messaging/features/mms"
	feature := MmsFeature{
		MmsSettings: MmsSettings{
			Protocol: "HTTP",
		},
	}

	result, _, err := c.makeAccountsRequest(ctx, http.MethodPost, path, &MmsFeatureResponse{}, &feature)
	if err != nil {
		return nil, err
	}
	return result.(*MmsFeatureResponse), nil
}

// AssociateApplication associates the peer with the application.
func (c *Client) AssociateApplication(ctx context.Context, siteID, peerID, applicationID string) (*ApplicationsSettingsResponse, error) {
	path := c.AccountsEndpoint + "/sites/" + siteID + "/sippeers/" + peerID + "/products/messaging/applicationSettings"
	req := ApplicationsSettings{
		HttpMessagingV2AppId: applicationID,
	}

	result, _, err := c.makeAccountsRequest(ctx, http.MethodPut, path, &ApplicationsSettingsResponse{}, &req)
	if err != nil {
		return nil, err
	}
	return result.(*ApplicationsSettingsResponse), nil
}

// GetAssociatedPeers returns the associated sippeers (aka locations) for the application.
func (c *Client) GetAssociatedPeers(ctx context.Context, applicationID string) (*AssociatedSipPeersResponse, error) {
	path := c.AccountsEndpoint + "/applications/" + applicationID + "/associatedsippeers"
	result, _, err := c.makeAccountsRequest(ctx, http.MethodGet, path, &AssociatedSipPeersResponse{})
	if err != nil {
		return nil, err
	}
	return result.(*AssociatedSipPeersResponse), nil
}

// GetNumbers returns the toll-free numbers associated with the site.
func (c *Client) GetNumbers(ctx context.Context, siteID, peerID string) (*SipPeerTelephoneNumbersResponse, error) {
	path := c.AccountsEndpoint + "/sites/" + siteID + "/sippeers/" + peerID + "/tns"
	result, _, err := c.makeAccountsRequest(ctx, http.MethodGet, path, &SipPeerTelephoneNumbersResponse{})
	if err != nil {
		return nil, err
	}
	return result.(*SipPeerTelephoneNumbersResponse), nil
}

// OrderNumbersByAreaCode purchases n numbers given area-code.
func (c *Client) OrderNumbersByAreaCode(ctx context.Context, siteID, peerID, areaCode string, n int) (*OrderResponse, error) {
	path := c.AccountsEndpoint + "/orders"
	req := AreaCodeRequest{
		AreaCodeOrder: AreaCodeOrder{
			SiteID: siteID,
			PeerID: peerID,
			AreaCodeSearchAndOrderType: AreaCodeSearchAndOrderType{
				Quantity: n,
				AreaCode: areaCode,
			},
		},
	}
	result, _, err := c.makeAccountsRequest(ctx, http.MethodPost, path, &OrderResponse{}, &req)
	if err != nil {
		return nil, err
	}
	return result.(*OrderResponse), nil
}

// SearchNumbersByAreaCode finds n numbers given area-code.
func (c *Client) SearchNumbersByAreaCode(ctx context.Context, areaCode string, n int) (*SearchResult, error) {
	path := c.AccountsEndpoint + "/availableNumbers"
	params := map[string]string{
		"areaCode": areaCode,
		"quantity": strconv.Itoa(n),
	}
	result, _, err := c.makeAccountsRequest(ctx, http.MethodGet, path, &SearchResult{}, params)
	if err != nil {
		return nil, err
	}
	return result.(*SearchResult), nil
}

// OrderTollFreeNumbers purchases n numbers given toll-free mask.
func (c *Client) OrderTollFreeNumbers(ctx context.Context, siteID, peerID, mask string, n int) (*OrderResponse, error) {
	path := c.AccountsEndpoint + "/orders"
	req := TollFreeOrderRequest{
		TollFreeOrder: TollFreeOrder{
			SiteID: siteID,
			PeerID: peerID,
			TollFreeWildCharSearchAndOrderType: TollFreeWildCharSearchAndOrderType{
				Quantity:                n,
				TollFreeWildCardPattern: mask,
			},
		},
	}
	result, _, err := c.makeAccountsRequest(ctx, http.MethodPost, path, &OrderResponse{}, &req)
	if err != nil {
		return nil, err
	}
	return result.(*OrderResponse), nil
}

// SearchTollFreeNumbers finds n numbers given tollfree mask.
func (c *Client) SearchTollFreeNumbers(ctx context.Context, mask string, n int) (*SearchResult, error) {
	path := c.AccountsEndpoint + "/availableNumbers"
	params := map[string]string{
		"tollFreeWildCardPattern": mask,
		"quantity":                strconv.Itoa(n),
	}
	result, _, err := c.makeAccountsRequest(ctx, http.MethodGet, path, &SearchResult{}, params)
	if err != nil {
		return nil, err
	}
	return result.(*SearchResult), nil
}

// GetOrder returns information regarding the given order.
func (c *Client) GetOrder(ctx context.Context, id string) (*OrderResponse, error) {
	path := c.AccountsEndpoint + "/orders/" + id
	result, _, err := c.makeAccountsRequest(ctx, http.MethodGet, path, &OrderResponse{})
	if err != nil {
		return nil, err
	}
	return result.(*OrderResponse), nil
}

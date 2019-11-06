package bandwidth

import (
	"context"
	"net/http"
)

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
		AreaCodeOrder: AreaCodeOrder {
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

// GetOrder returns information regarding the given order.
func (c *Client) GetOrder(ctx context.Context, id string) (*OrderResponse, error) {
	path := c.AccountsEndpoint + "/orders/" + id
	result, _, err := c.makeAccountsRequest(ctx, http.MethodGet, path, &OrderResponse{})
	if err != nil {
		return nil, err
	}
	return result.(*OrderResponse), nil
}
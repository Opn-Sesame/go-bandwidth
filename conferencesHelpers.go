package bandwidth

import "strconv"

// TerminateConference terminates a  conference
func (api *Client) TerminateConference(id string) error{
	return api.UpdateConference(id, map[string]interface{}{"state": "completed"})
}

// MuteConference mutes/unmutes a  conference
func (api *Client) MuteConference(id string, mute bool) error{
	return api.UpdateConference(id, map[string]interface{}{"mute": strconv.FormatBool(mute)})
}

// DeleteConferenceMember removes the member from the conference
func (api *Client) DeleteConferenceMember(id string, memberID string) error{
	return api.UpdateConferenceMember(id, memberID, map[string]interface{}{"state": "completed"})
}

// MuteConferenceMember mute/unmute the conference member
func (api *Client) MuteConferenceMember(id string, memberID string, mute bool) error{
	return api.UpdateConferenceMember(id, memberID, map[string]interface{}{"mute": strconv.FormatBool(mute)})
}

// HoldConferenceMember hold/unhold the conference member
func (api *Client) HoldConferenceMember(id string, memberID string, hold bool) error{
	return api.UpdateConferenceMember(id, memberID, map[string]interface{}{"hold": strconv.FormatBool(hold)})
}

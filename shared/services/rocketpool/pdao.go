package rocketpool

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/goccy/go-json"

	"github.com/rocket-pool/rocketpool-go/types"
	"github.com/rocket-pool/smartnode/shared/types/api"
)

// Get protocol DAO proposals
func (c *Client) PDAOProposals() (api.PDAOProposalsResponse, error) {
	responseBytes, err := c.callAPI("pdao proposals")
	if err != nil {
		return api.PDAOProposalsResponse{}, fmt.Errorf("Could not get protocol DAO proposals: %w", err)
	}
	var response api.PDAOProposalsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.PDAOProposalsResponse{}, fmt.Errorf("Could not decode protocol DAO proposals response: %w", err)
	}
	if response.Error != "" {
		return api.PDAOProposalsResponse{}, fmt.Errorf("Could not get protocol DAO proposals: %s", response.Error)
	}
	return response, nil
}

// Get protocol DAO proposal details
func (c *Client) PDAOProposalDetails(proposalID uint64) (api.PDAOProposalResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("pdao proposal-details %d", proposalID))
	if err != nil {
		return api.PDAOProposalResponse{}, fmt.Errorf("Could not get protocol DAO proposal: %w", err)
	}
	var response api.PDAOProposalResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.PDAOProposalResponse{}, fmt.Errorf("Could not decode protocol DAO proposal response: %w", err)
	}
	if response.Error != "" {
		return api.PDAOProposalResponse{}, fmt.Errorf("Could not get protocol DAO proposal: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can vote on a proposal
func (c *Client) PDAOCanVoteProposal(proposalID uint64, voteDirection types.VoteDirection) (api.CanVoteOnPDAOProposalResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("pdao can-vote-proposal %d %s", proposalID, types.VoteDirections[voteDirection]))
	if err != nil {
		return api.CanVoteOnPDAOProposalResponse{}, fmt.Errorf("Could not get protocol DAO can-vote-proposal: %w", err)
	}
	var response api.CanVoteOnPDAOProposalResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanVoteOnPDAOProposalResponse{}, fmt.Errorf("Could not decode protocol DAO can-vote-proposal response: %w", err)
	}
	if response.Error != "" {
		return api.CanVoteOnPDAOProposalResponse{}, fmt.Errorf("Could not get protocol DAO can-vote-proposal: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can vote on a proposal
func (c *Client) PDAOVoteProposal(proposalID uint64, voteDirection types.VoteDirection) (api.VoteOnPDAOProposalResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("pdao vote-proposal %d %s", proposalID, types.VoteDirections[voteDirection]))
	if err != nil {
		return api.VoteOnPDAOProposalResponse{}, fmt.Errorf("Could not get protocol DAO vote-proposal: %w", err)
	}
	var response api.VoteOnPDAOProposalResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.VoteOnPDAOProposalResponse{}, fmt.Errorf("Could not decode protocol DAO vote-proposal response: %w", err)
	}
	if response.Error != "" {
		return api.VoteOnPDAOProposalResponse{}, fmt.Errorf("Could not get protocol DAO vote-proposal: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can execute a proposal
func (c *Client) PDAOCanExecuteProposal(proposalID uint64) (api.CanExecutePDAOProposalResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("pdao can-execute-proposal %d", proposalID))
	if err != nil {
		return api.CanExecutePDAOProposalResponse{}, fmt.Errorf("Could not get protocol DAO can-execute-proposal: %w", err)
	}
	var response api.CanExecutePDAOProposalResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanExecutePDAOProposalResponse{}, fmt.Errorf("Could not decode protocol DAO can-execute-proposal response: %w", err)
	}
	if response.Error != "" {
		return api.CanExecutePDAOProposalResponse{}, fmt.Errorf("Could not get protocol DAO can-execute-proposal: %s", response.Error)
	}
	return response, nil
}

// Execute a proposal
func (c *Client) PDAOExecuteProposal(proposalID uint64) (api.ExecutePDAOProposalResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("pdao execute-proposal %d", proposalID))
	if err != nil {
		return api.ExecutePDAOProposalResponse{}, fmt.Errorf("Could not get protocol DAO execute-proposal: %w", err)
	}
	var response api.ExecutePDAOProposalResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.ExecutePDAOProposalResponse{}, fmt.Errorf("Could not decode protocol DAO execute-proposal response: %w", err)
	}
	if response.Error != "" {
		return api.ExecutePDAOProposalResponse{}, fmt.Errorf("Could not get protocol DAO execute-proposal: %s", response.Error)
	}
	return response, nil
}

// Get protocol DAO settings
func (c *Client) PDAOGetSettings() (api.GetPDAOSettingsResponse, error) {
	responseBytes, err := c.callAPI("pdao get-settings")
	if err != nil {
		return api.GetPDAOSettingsResponse{}, fmt.Errorf("Could not get protocol DAO get-settings: %w", err)
	}
	var response api.GetPDAOSettingsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.GetPDAOSettingsResponse{}, fmt.Errorf("Could not decode protocol DAO get-settings response: %w", err)
	}
	if response.Error != "" {
		return api.GetPDAOSettingsResponse{}, fmt.Errorf("Could not get protocol DAO get-settings: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can vote on a proposal
func (c *Client) PDAOCanProposeSetting(setting string, value string) (api.CanProposePDAOSettingResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("pdao can-propose-setting %s %s", setting, value))
	if err != nil {
		return api.CanProposePDAOSettingResponse{}, fmt.Errorf("Could not get protocol DAO can-propose-setting: %w", err)
	}
	var response api.CanProposePDAOSettingResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanProposePDAOSettingResponse{}, fmt.Errorf("Could not decode protocol DAO can-propose-setting response: %w", err)
	}
	if response.Error != "" {
		return api.CanProposePDAOSettingResponse{}, fmt.Errorf("Could not get protocol DAO can-propose-setting: %s", response.Error)
	}
	return response, nil
}

// Propose updating a PDAO setting (use can-propose-setting to get the pollard)
func (c *Client) PDAOProposeSetting(setting string, value string, blockNumber uint32) (api.ProposePDAOSettingResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("pdao propose-setting %s %s %d", setting, value, blockNumber))
	if err != nil {
		return api.ProposePDAOSettingResponse{}, fmt.Errorf("Could not get protocol DAO propose-setting: %w", err)
	}
	var response api.ProposePDAOSettingResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.ProposePDAOSettingResponse{}, fmt.Errorf("Could not decode protocol DAO propose-setting response: %w", err)
	}
	if response.Error != "" {
		return api.ProposePDAOSettingResponse{}, fmt.Errorf("Could not get protocol DAO propose-setting: %s", response.Error)
	}
	return response, nil
}

// Get the allocation percentages of RPL rewards for the Oracle DAO, the Protocol DAO, and the node operators
func (c *Client) PDAOGetRewardsPercentages() (api.PDAOGetRewardsPercentagesResponse, error) {
	responseBytes, err := c.callAPI("pdao get-rewards-percentages")
	if err != nil {
		return api.PDAOGetRewardsPercentagesResponse{}, fmt.Errorf("Could not get protocol DAO get-rewards-percentages: %w", err)
	}
	var response api.PDAOGetRewardsPercentagesResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.PDAOGetRewardsPercentagesResponse{}, fmt.Errorf("Could not decode protocol DAO get-rewards-percentages response: %w", err)
	}
	if response.Error != "" {
		return api.PDAOGetRewardsPercentagesResponse{}, fmt.Errorf("Could not get protocol DAO get-rewards-percentages: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can propose new RPL rewards allocation percentages for the Oracle DAO, the Protocol DAO, and the node operators
func (c *Client) PDAOCanProposeRewardsPercentages(node *big.Int, odao *big.Int, pdao *big.Int) (api.PDAOCanProposeRewardsPercentagesResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("pdao can-propose-rewards-percentages %s %s %s", node.String(), odao.String(), pdao.String()))
	if err != nil {
		return api.PDAOCanProposeRewardsPercentagesResponse{}, fmt.Errorf("Could not get protocol DAO can-propose-rewards-percentages: %w", err)
	}
	var response api.PDAOCanProposeRewardsPercentagesResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.PDAOCanProposeRewardsPercentagesResponse{}, fmt.Errorf("Could not decode protocol DAO can-propose-rewards-percentages response: %w", err)
	}
	if response.Error != "" {
		return api.PDAOCanProposeRewardsPercentagesResponse{}, fmt.Errorf("Could not get protocol DAO can-propose-rewards-percentages: %s", response.Error)
	}
	return response, nil
}

// Propose new RPL rewards allocation percentages for the Oracle DAO, the Protocol DAO, and the node operators
func (c *Client) PDAOProposeRewardsPercentages(node *big.Int, odao *big.Int, pdao *big.Int, blockNumber uint32) (api.ProposePDAOSettingResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("pdao propose-rewards-percentages %s %s %s %d", node, odao, pdao, blockNumber))
	if err != nil {
		return api.ProposePDAOSettingResponse{}, fmt.Errorf("Could not get protocol DAO propose-rewards-percentages: %w", err)
	}
	var response api.ProposePDAOSettingResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.ProposePDAOSettingResponse{}, fmt.Errorf("Could not decode protocol DAO propose-rewards-percentages response: %w", err)
	}
	if response.Error != "" {
		return api.ProposePDAOSettingResponse{}, fmt.Errorf("Could not get protocol DAO propose-rewards-percentages: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can propose a one-time spend of the Protocol DAO's treasury
func (c *Client) PDAOCanProposeOneTimeSpend(invoiceID string, recipient common.Address, amount *big.Int) (api.PDAOCanProposeOneTimeSpendResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("pdao can-propose-one-time-spend %s %s %s", invoiceID, recipient.Hex(), amount.String()))
	if err != nil {
		return api.PDAOCanProposeOneTimeSpendResponse{}, fmt.Errorf("Could not get protocol DAO can-propose-one-time-spend: %w", err)
	}
	var response api.PDAOCanProposeOneTimeSpendResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.PDAOCanProposeOneTimeSpendResponse{}, fmt.Errorf("Could not decode protocol DAO can-propose-one-time-spend response: %w", err)
	}
	if response.Error != "" {
		return api.PDAOCanProposeOneTimeSpendResponse{}, fmt.Errorf("Could not get protocol DAO can-propose-one-time-spend: %s", response.Error)
	}
	return response, nil
}

// Propose a one-time spend of the Protocol DAO's treasury
func (c *Client) PDAOProposeOneTimeSpend(invoiceID string, recipient common.Address, amount *big.Int, blockNumber uint32) (api.PDAOProposeOneTimeSpendResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("pdao propose-one-time-spend %s %s %s %d", invoiceID, recipient.Hex(), amount.String(), blockNumber))
	if err != nil {
		return api.PDAOProposeOneTimeSpendResponse{}, fmt.Errorf("Could not get protocol DAO propose-one-time-spend: %w", err)
	}
	var response api.PDAOProposeOneTimeSpendResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.PDAOProposeOneTimeSpendResponse{}, fmt.Errorf("Could not decode protocol DAO propose-one-time-spend response: %w", err)
	}
	if response.Error != "" {
		return api.PDAOProposeOneTimeSpendResponse{}, fmt.Errorf("Could not get protocol DAO propose-one-time-spend: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can propose a recurring spend of the Protocol DAO's treasury
func (c *Client) PDAOCanProposeRecurringSpend(contractName string, recipient common.Address, amountPerPeriod *big.Int, periodLength time.Duration, startTime time.Time, numberOfPeriods uint64) (api.PDAOCanProposeRecurringSpendResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("pdao can-propose-recurring-spend %s %s %s %s %d %d", contractName, recipient.Hex(), amountPerPeriod.String(), periodLength.String(), startTime.Unix(), numberOfPeriods))
	if err != nil {
		return api.PDAOCanProposeRecurringSpendResponse{}, fmt.Errorf("Could not get protocol DAO can-propose-recurring-spend: %w", err)
	}
	var response api.PDAOCanProposeRecurringSpendResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.PDAOCanProposeRecurringSpendResponse{}, fmt.Errorf("Could not decode protocol DAO can-propose-recurring-spend response: %w", err)
	}
	if response.Error != "" {
		return api.PDAOCanProposeRecurringSpendResponse{}, fmt.Errorf("Could not get protocol DAO can-propose-recurring-spend: %s", response.Error)
	}
	return response, nil
}

// Propose a recurring spend of the Protocol DAO's treasury
func (c *Client) PDAOProposeRecurringSpend(contractName string, recipient common.Address, amountPerPeriod *big.Int, periodLength time.Duration, startTime time.Time, numberOfPeriods uint64, blockNumber uint32) (api.PDAOProposeRecurringSpendResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("pdao propose-recurring-spend %s %s %s %s %d %d %d", contractName, recipient.Hex(), amountPerPeriod.String(), periodLength.String(), startTime.Unix(), numberOfPeriods, blockNumber))
	if err != nil {
		return api.PDAOProposeRecurringSpendResponse{}, fmt.Errorf("Could not get protocol DAO propose-recurring-spend: %w", err)
	}
	var response api.PDAOProposeRecurringSpendResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.PDAOProposeRecurringSpendResponse{}, fmt.Errorf("Could not decode protocol DAO propose-recurring-spend response: %w", err)
	}
	if response.Error != "" {
		return api.PDAOProposeRecurringSpendResponse{}, fmt.Errorf("Could not get protocol DAO propose-recurring-spend: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can propose an update to an existing recurring spend plan
func (c *Client) PDAOCanProposeRecurringSpendUpdate(contractName string, recipient common.Address, amountPerPeriod *big.Int, periodLength time.Duration, numberOfPeriods uint64) (api.PDAOCanProposeRecurringSpendUpdateResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("pdao can-propose-recurring-spend-update %s %s %s %s %d", contractName, recipient.Hex(), amountPerPeriod.String(), periodLength.String(), numberOfPeriods))
	if err != nil {
		return api.PDAOCanProposeRecurringSpendUpdateResponse{}, fmt.Errorf("Could not get protocol DAO can-propose-recurring-spend-update: %w", err)
	}
	var response api.PDAOCanProposeRecurringSpendUpdateResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.PDAOCanProposeRecurringSpendUpdateResponse{}, fmt.Errorf("Could not decode protocol DAO can-propose-recurring-spend-update response: %w", err)
	}
	if response.Error != "" {
		return api.PDAOCanProposeRecurringSpendUpdateResponse{}, fmt.Errorf("Could not get protocol DAO can-propose-recurring-spend-update: %s", response.Error)
	}
	return response, nil
}

// Propose an update to an existing recurring spend plan
func (c *Client) PDAOProposeRecurringSpendUpdate(contractName string, recipient common.Address, amountPerPeriod *big.Int, periodLength time.Duration, numberOfPeriods uint64, blockNumber uint32) (api.PDAOProposeRecurringSpendUpdateResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("pdao propose-recurring-spend-update %s %s %s %s %d %d", contractName, recipient.Hex(), amountPerPeriod.String(), periodLength.String(), numberOfPeriods, blockNumber))
	if err != nil {
		return api.PDAOProposeRecurringSpendUpdateResponse{}, fmt.Errorf("Could not get protocol DAO propose-recurring-spend-update: %w", err)
	}
	var response api.PDAOProposeRecurringSpendUpdateResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.PDAOProposeRecurringSpendUpdateResponse{}, fmt.Errorf("Could not decode protocol DAO propose-recurring-spend-update response: %w", err)
	}
	if response.Error != "" {
		return api.PDAOProposeRecurringSpendUpdateResponse{}, fmt.Errorf("Could not get protocol DAO propose-recurring-spend-update: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can invite someone to the security council
func (c *Client) PDAOCanProposeInviteToSecurityCouncil(id string, address common.Address) (api.PDAOCanProposeInviteToSecurityCouncilResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("pdao can-propose-invite-to-security-council %s %s", id, address.Hex()))
	if err != nil {
		return api.PDAOCanProposeInviteToSecurityCouncilResponse{}, fmt.Errorf("Could not get protocol DAO can-propose-invite-to-security-council: %w", err)
	}
	var response api.PDAOCanProposeInviteToSecurityCouncilResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.PDAOCanProposeInviteToSecurityCouncilResponse{}, fmt.Errorf("Could not decode protocol DAO can-propose-invite-to-security-council response: %w", err)
	}
	if response.Error != "" {
		return api.PDAOCanProposeInviteToSecurityCouncilResponse{}, fmt.Errorf("Could not get protocol DAO can-propose-invite-to-security-council: %s", response.Error)
	}
	return response, nil
}

// Propose inviting someone to the security council
func (c *Client) PDAOProposeInviteToSecurityCouncil(id string, address common.Address, blockNumber uint32) (api.PDAOProposeInviteToSecurityCouncilResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("pdao propose-invite-to-security-council %s %s %d", id, address.Hex(), blockNumber))
	if err != nil {
		return api.PDAOProposeInviteToSecurityCouncilResponse{}, fmt.Errorf("Could not get protocol DAO propose-invite-to-security-council: %w", err)
	}
	var response api.PDAOProposeInviteToSecurityCouncilResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.PDAOProposeInviteToSecurityCouncilResponse{}, fmt.Errorf("Could not decode protocol DAO propose-invite-to-security-council response: %w", err)
	}
	if response.Error != "" {
		return api.PDAOProposeInviteToSecurityCouncilResponse{}, fmt.Errorf("Could not get protocol DAO propose-invite-to-security-council: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can kick someone from the security council
func (c *Client) PDAOCanProposeKickFromSecurityCouncil(address common.Address) (api.PDAOCanProposeKickFromSecurityCouncilResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("pdao can-propose-kick-from-security-council %s", address.Hex()))
	if err != nil {
		return api.PDAOCanProposeKickFromSecurityCouncilResponse{}, fmt.Errorf("Could not get protocol DAO can-propose-kick-from-security-council: %w", err)
	}
	var response api.PDAOCanProposeKickFromSecurityCouncilResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.PDAOCanProposeKickFromSecurityCouncilResponse{}, fmt.Errorf("Could not decode protocol DAO can-propose-kick-from-security-council response: %w", err)
	}
	if response.Error != "" {
		return api.PDAOCanProposeKickFromSecurityCouncilResponse{}, fmt.Errorf("Could not get protocol DAO can-propose-kick-from-security-council: %s", response.Error)
	}
	return response, nil
}

// Propose kicking someone from the security council
func (c *Client) PDAOProposeKickFromSecurityCouncil(address common.Address, blockNumber uint32) (api.PDAOProposeKickFromSecurityCouncilResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("pdao propose-kick-from-security-council %s %d", address.Hex(), blockNumber))
	if err != nil {
		return api.PDAOProposeKickFromSecurityCouncilResponse{}, fmt.Errorf("Could not get protocol DAO propose-kick-from-security-council: %w", err)
	}
	var response api.PDAOProposeKickFromSecurityCouncilResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.PDAOProposeKickFromSecurityCouncilResponse{}, fmt.Errorf("Could not decode protocol DAO propose-kick-from-security-council response: %w", err)
	}
	if response.Error != "" {
		return api.PDAOProposeKickFromSecurityCouncilResponse{}, fmt.Errorf("Could not get protocol DAO propose-kick-from-security-council: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can kick multiple members from the security council
func (c *Client) PDAOCanProposeKickMutliFromSecurityCouncil(addresses []common.Address) (api.PDAOCanProposeKickMultiFromSecurityCouncilResponse, error) {
	addressStrings := make([]string, len(addresses))
	for i, address := range addresses {
		addressStrings[i] = address.Hex()
	}

	responseBytes, err := c.callAPI(fmt.Sprintf("pdao can-propose-kick-multi-from-security-council %s", strings.Join(addressStrings, ",")))
	if err != nil {
		return api.PDAOCanProposeKickMultiFromSecurityCouncilResponse{}, fmt.Errorf("Could not get protocol DAO can-propose-kick-multi-from-security-council: %w", err)
	}
	var response api.PDAOCanProposeKickMultiFromSecurityCouncilResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.PDAOCanProposeKickMultiFromSecurityCouncilResponse{}, fmt.Errorf("Could not decode protocol DAO can-propose-kick-multi-from-security-council response: %w", err)
	}
	if response.Error != "" {
		return api.PDAOCanProposeKickMultiFromSecurityCouncilResponse{}, fmt.Errorf("Could not get protocol DAO can-propose-kick-multi-from-security-council: %s", response.Error)
	}
	return response, nil
}

// Propose kicking multiple members from the security council
func (c *Client) PDAOProposeKickMultiFromSecurityCouncil(addresses []common.Address, blockNumber uint32) (api.PDAOProposeKickMultiFromSecurityCouncilResponse, error) {
	addressStrings := make([]string, len(addresses))
	for i, address := range addresses {
		addressStrings[i] = address.Hex()
	}

	responseBytes, err := c.callAPI(fmt.Sprintf("pdao propose-kick-multi-from-security-council %s %d", strings.Join(addressStrings, ","), blockNumber))
	if err != nil {
		return api.PDAOProposeKickMultiFromSecurityCouncilResponse{}, fmt.Errorf("Could not get protocol DAO propose-kick-multi-from-security-council: %w", err)
	}
	var response api.PDAOProposeKickMultiFromSecurityCouncilResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.PDAOProposeKickMultiFromSecurityCouncilResponse{}, fmt.Errorf("Could not decode protocol DAO propose-kick-multi-from-security-council response: %w", err)
	}
	if response.Error != "" {
		return api.PDAOProposeKickMultiFromSecurityCouncilResponse{}, fmt.Errorf("Could not get protocol DAO propose-kick-multi-from-security-council: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can propose replacing someone on the security council with another member
func (c *Client) PDAOCanProposeReplaceMemberOfSecurityCouncil(existingAddress common.Address, newID string, newAddress common.Address) (api.PDAOCanProposeReplaceMemberOfSecurityCouncilResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("pdao can-propose-replace-member-of-security-council %s %s %s", existingAddress.Hex(), newID, newAddress.Hex()))
	if err != nil {
		return api.PDAOCanProposeReplaceMemberOfSecurityCouncilResponse{}, fmt.Errorf("Could not get protocol DAO can-propose-replace-member-of-security-council: %w", err)
	}
	var response api.PDAOCanProposeReplaceMemberOfSecurityCouncilResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.PDAOCanProposeReplaceMemberOfSecurityCouncilResponse{}, fmt.Errorf("Could not decode protocol DAO can-propose-replace-member-of-security-council response: %w", err)
	}
	if response.Error != "" {
		return api.PDAOCanProposeReplaceMemberOfSecurityCouncilResponse{}, fmt.Errorf("Could not get protocol DAO can-propose-replace-member-of-security-council: %s", response.Error)
	}
	return response, nil
}

// Propose replacing someone on the security council with another member
func (c *Client) PDAOProposeReplaceMemberOfSecurityCouncil(existingAddress common.Address, newID string, newAddress common.Address, blockNumber uint32) (api.PDAOProposeReplaceMemberOfSecurityCouncilResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("pdao propose-replace-member-of-security-council %s %s %s %d", existingAddress.Hex(), newID, newAddress.Hex(), blockNumber))
	if err != nil {
		return api.PDAOProposeReplaceMemberOfSecurityCouncilResponse{}, fmt.Errorf("Could not get protocol DAO propose-replace-member-of-security-council: %w", err)
	}
	var response api.PDAOProposeReplaceMemberOfSecurityCouncilResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.PDAOProposeReplaceMemberOfSecurityCouncilResponse{}, fmt.Errorf("Could not decode protocol DAO propose-replace-member-of-security-council response: %w", err)
	}
	if response.Error != "" {
		return api.PDAOProposeReplaceMemberOfSecurityCouncilResponse{}, fmt.Errorf("Could not get protocol DAO propose-replace-member-of-security-council: %s", response.Error)
	}
	return response, nil
}

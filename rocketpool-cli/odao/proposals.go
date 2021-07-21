package odao

import (
    "encoding/hex"
    "fmt"

    "github.com/rocket-pool/rocketpool-go/dao"
    "github.com/rocket-pool/rocketpool-go/types"
    "github.com/urfave/cli"

    "github.com/rocket-pool/smartnode/shared/services/rocketpool"
)


func getProposals(c *cli.Context) error {

    // Get RP client
    rp, err := rocketpool.NewClientFromCtx(c)
    if err != nil { return err }
    defer rp.Close()

    // Get oracle DAO proposals
    allProposals, err := rp.TNDAOProposals()
    if err != nil {
        return err
    }

    // Get proposals by state
    stateProposals := map[string][]dao.ProposalDetails{}
    for _, proposal := range allProposals.Proposals {
        stateName := proposal.State.String()
        if _, ok := stateProposals[stateName]; !ok {
            stateProposals[stateName] = []dao.ProposalDetails{}
        }
        stateProposals[stateName] = append(stateProposals[stateName], proposal)
    }

    // Proposal states print order
    proposalStates := []string{"Pending", "Active", "Succeeded", "Executed", "Cancelled", "Defeated", "Expired"}

    // Print & return
    if len(allProposals.Proposals) == 0 {
        fmt.Println("There are no oracle DAO proposals yet.")
    }
    for _, stateName := range proposalStates {
        proposals, ok := stateProposals[stateName]
        if !ok { continue }

        // Proposal state count
        fmt.Printf("%d %s proposal(s):\n", len(proposals), stateName)
        fmt.Println("")

        // Proposals
        for _, proposal := range proposals {
            fmt.Printf("--------------------\n")
            fmt.Printf("\n")

            // Main details
            fmt.Printf("Proposal ID:          %d\n", proposal.ID)
            fmt.Printf("Message:              %s\n", proposal.Message)
            fmt.Printf("Payload:              %s\n", proposal.PayloadStr)
            fmt.Printf("Payload (bytes):      %s\n", hex.EncodeToString(proposal.Payload))
            fmt.Printf("Proposed by:          %s\n", proposal.ProposerAddress.Hex())
            fmt.Printf("Created time:         %d\n", proposal.CreatedTime)

            // Start block - pending proposals
            if proposal.State == types.Pending {
            fmt.Printf("Start time:           %d\n", proposal.StartTime)
            }

            // End block - active proposals
            if proposal.State == types.Active {
            fmt.Printf("Ends time:            %d\n", proposal.EndTime)
            }

            // Expiry block - succeeded proposals
            if proposal.State == types.Succeeded {
            fmt.Printf("Expiry time:          %d\n", proposal.ExpiryTime)
            }

            // Vote details
            fmt.Printf("Votes required:       %.2f\n", proposal.VotesRequired)
            fmt.Printf("Votes for:            %.2f\n", proposal.VotesFor)
            fmt.Printf("Votes against:        %.2f\n", proposal.VotesAgainst)
            if proposal.MemberVoted {
                if proposal.MemberSupported {
            fmt.Printf("Node has voted:       for\n")
                } else {
            fmt.Printf("Node has voted:       against\n")
                }
            } else {
            fmt.Printf("Node has voted:       no\n")
            }

            fmt.Printf("\n")
        }

    }
    return nil

}


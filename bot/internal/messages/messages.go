package messages

var (
	NotEnoughPermissionsToCreateBounty = `
You do not have sufficient permissions to create a bounty. Please check organization membership status or contact the organization's owner or administrator for assistance.

Bounty label will be removed.
`
	NotEnoughPermissionsToRunCommands = `
You are not allowed to run commands on this repository. 
`
	InvalidValueInput = `The entered reward amount is not valid. Please check and try again.`
	InvalidUserInput  = `The user you've mentioned doesn't seem to exist. Please double-check the username and try again.`
)

var (
	UserHasNoWalletTmpl = `
Oops! Looks like this user forgot to set his wallet address. 

%s, please set your wallet address to receive the payment - [Documentation](https://deed-labs.gitbook.io/gittips/product-guides/commands#set-wallet-address)
`

	PaymentSentTmpl = `
Payment successfully sent! %s, check your wallet!
`
)

package parser

type Result struct {
	Commands      []string
	WalletAddress string
	Reward        uint64
}

func ParseBody(body string) Result {
	res := Result{}

	// TODO

	return res
}

func SearchLabel(target LabelText, labels []string) bool {
	for _, v := range labels {
		if v == string(target) {
			return true
		}
	}

	return false
}

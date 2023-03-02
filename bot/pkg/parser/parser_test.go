package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseBody(t *testing.T) {
	testCases := []struct {
		text string
		res  Result
	}{
		{
			text: `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod 
				tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, 
				quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.

				;; wallet: EQA_HlQjfA0l94sUZJ4dPPodOOZxdOTMy_68077awDIAua-D
			`,
			res: Result{
				WalletAddress: "EQA_HlQjfA0l94sUZJ4dPPodOOZxdOTMy_68077awDIAua-D",
			},
		},
		{
			text: `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod 
				tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, 
				quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.

				;; address EQA_HlQjfA0l94sUZJ4dPPodOOZxdOTMy_68077awDIAua-D
			`,
			res: Result{
				WalletAddress: "EQA_HlQjfA0l94sUZJ4dPPodOOZxdOTMy_68077awDIAua-D",
			},
		},
		{
			text: `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod 
				tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, 
				quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.

				;; reward: 10.0
			`,
			res: Result{
				Reward: "10.0",
			},
		},
		{
			text: `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod 
				tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, 
				quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.

				;; wallet: EQA_HlQjfA0l94sUZJ4dPPodOOZxdOTMy_68077awDIAua-D
				;; reward: 10,0
			`,
			res: Result{
				WalletAddress: "EQA_HlQjfA0l94sUZJ4dPPodOOZxdOTMy_68077awDIAua-D",
				Reward:        "10,0",
			},
		},
		{
			text: `;; gt pay
				;; gt set reward 10,0`,
			res: Result{
				Commands: []string{"pay", "set reward 10,0"},
			},
		},
		{
			text: `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod 
				tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, 
				quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
			`,
			res: Result{},
		},
		{
			text: ``,
			res:  Result{},
		},
	}

	for _, tc := range testCases {
		result := Parse(tc.text)
		require.Equal(t, tc.res, result)
	}
}

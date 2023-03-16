export type Bounty = {
	ownerId: number;
	owner: string;
	ownerUrl: string;
	ownerAvatarUrl: string;
	ownerType: string;
	title: string;
	url: string;
	reward: string;
	rewardUSD: string;
};

export type OwnerInfo = {
	name: string;
	avatarUrl: string;
	totalBudget: string;
	availableBudget: string;
	totalBounties: number;
	availableBounties: number;
	bounties: Bounty[];
};

export type InstallationInfo = {
	installed: boolean;
	name: string;
	ownerId: number;
};

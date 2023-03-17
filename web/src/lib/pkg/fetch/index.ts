import type { Bounty, InstallationInfo, OwnerInfo } from '../../types';
import { PUBLIC_API_URL } from '$env/static/public';
import axios from 'axios';

interface BountyJSON {
	owner_id: number;
	owner: string;
	owner_url: string;
	owner_avatar_url: string;
	owner_type: string;
	title: string;
	url: string;
	reward: string;
}

interface BountiesResponse {
	bounties: BountyJSON[];
}

interface InstallationInfoResponse {
	installed: boolean;
	owner_name: string;
	owner_id: number;
}

interface OwnerInfoResponse {
	name: string;
	avatar_url: string;
	total_budget: string;
	available_budget: string;
	total_bounties: number;
	available_bounties: number;
	bounties: BountyJSON[];
}

const axiosAPI = axios.create({
	baseURL: PUBLIC_API_URL
});

export const fetchBounties = async (): Promise<Bounty[]> => {
	try {
		const resp = await axiosAPI.get<BountiesResponse>('/api/bounties', {
			headers: { Accept: 'application/json' }
		});

		let bounties: Bounty[] = [];
		resp.data.bounties.forEach((b) => {
			bounties.push({
				ownerId: b.owner_id,
				owner: b.owner,
				ownerUrl: b.owner_url,
				ownerAvatarUrl: b.owner_avatar_url,
				ownerType: b.owner_type,
				title: b.title,
				url: b.url,
				reward: b.reward,
				rewardUSD: 0
			});
		});

		return bounties;
	} catch (error) {
		if (axios.isAxiosError(error)) {
			console.log('error message: ', error.message);
		} else {
			console.log('unexpected error: ', error);
		}

		return [];
	}
};

export const setupInstallation = async (
	walletAddress: string,
	installationId: number
): Promise<InstallationInfo> => {
	let info: InstallationInfo = {} as InstallationInfo;

	const body = JSON.stringify({
		installation_id: installationId,
		wallet_address: walletAddress
	});

	try {
		const resp = await axiosAPI.post('/setup', body, {
			headers: { Accept: 'application/json' }
		});

		if (resp.status != 200) throw new Error('failed to setup installation');

		info.installed = resp.data.installed;
		info.name = resp.data.owner_name;
		info.ownerId = resp.data.owner_id;

		return info;
	} catch (error) {
		if (axios.isAxiosError(error)) {
			console.log('error message: ', error.message);
		} else {
			console.log('unexpected error: ', error);
		}

		return info;
	}
};

export const fetchInstallationInfo = async (address: string): Promise<InstallationInfo> => {
	let info: InstallationInfo = {} as InstallationInfo;

	try {
		const resp = await axiosAPI.get<InstallationInfoResponse>(`/api/installation/${address}`);

		info.installed = resp.data.installed;
		info.name = resp.data.owner_name;
		info.ownerId = resp.data.owner_id;

		return info;
	} catch (error) {
		if (axios.isAxiosError(error)) {
			console.log('error message: ', error.message);
		} else {
			console.log('unexpected error: ', error);
		}

		return info;
	}
};

export const fetchOwnerInfo = async (id: number): Promise<OwnerInfo> => {
	let info: OwnerInfo = { bounties: new Array<Bounty>() } as OwnerInfo;

	try {
		const resp = await axiosAPI.get<OwnerInfoResponse>('/api/owner/' + id, {
			headers: { Accept: 'application/json' }
		});

		info.name = resp.data.name;
		info.avatarUrl = resp.data.avatar_url;
		info.totalBudget = resp.data.total_budget;
		info.availableBudget = resp.data.available_budget;
		info.totalBounties = resp.data.total_bounties;
		info.availableBounties = resp.data.available_bounties;

		resp.data.bounties.forEach((b) => {
			info.bounties.push({
				ownerId: b.owner_id,
				owner: b.owner,
				ownerUrl: b.owner_url,
				ownerAvatarUrl: b.owner_avatar_url,
				ownerType: b.owner_type,
				title: b.title,
				url: b.url,
				reward: b.reward,
				rewardUSD: 0
			});
		});

		return info;
	} catch (error) {
		if (axios.isAxiosError(error)) {
			console.log('error message: ', error.message);
		} else {
			console.log('unexpected error: ', error);
		}

		return info;
	}
};

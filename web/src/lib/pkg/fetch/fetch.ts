import type { Bounty } from '../../types';
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

interface BountiesRequest {
	bounties: BountyJSON[];
}

const axiosAPI = axios.create({
	baseURL: PUBLIC_API_URL
});

export const fetchBounties = async (): Promise<Bounty[]> => {
	try {
		const resp = await axiosAPI.get<BountiesRequest>('/api/bounties', {
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
				rewardUSD: ''
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
): Promise<void> => {
	const body = JSON.stringify({
		installation_id: installationId,
		wallet_address: walletAddress
	});

	try {
		const resp = await axiosAPI.post('/setup', body, {
			headers: { Accept: 'application/json' }
		});

		if (resp.status != 200) throw new Error('failed to setup installation');
	} catch (error) {
		if (axios.isAxiosError(error)) {
			console.log('error message: ', error.message);
		} else {
			console.log('unexpected error: ', error);
		}
	}
};

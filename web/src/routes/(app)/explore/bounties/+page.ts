import { fetchBounties } from '$lib/pkg/fetch';
import { bigIntToFloat } from '$lib/utils';
import CoinGecko from 'coingecko-api';

/** @type {import('./$types').PageLoad} */
export async function load({}) {
	const bounties = await fetchBounties();

	const cgClient = new CoinGecko();
	const tonPriceData = (
		await cgClient.simple.price({ ids: 'the-open-network', vs_currencies: 'usd' })
	).data;

	bounties.forEach((b) => {
		b.rewardUSD = Number(bigIntToFloat(b.reward, 9, 2)) * tonPriceData['the-open-network']['usd'];
	});

	return { bounties };
}

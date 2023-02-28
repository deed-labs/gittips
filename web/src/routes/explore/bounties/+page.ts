import { fetchBounties } from '$lib/pkg/fetch/fetch';

/** @type {import('./$types').PageLoad} */
export async function load({}) {
	const bounties = await fetchBounties();

	return { bounties };
}

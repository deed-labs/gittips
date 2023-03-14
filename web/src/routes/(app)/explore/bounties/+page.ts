import { fetchBounties } from '$lib/pkg/fetch';

/** @type {import('./$types').PageLoad} */
export async function load({}) {
	const bounties = await fetchBounties();

	return { bounties };
}

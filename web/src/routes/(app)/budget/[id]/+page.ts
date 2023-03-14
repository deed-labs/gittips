import { fetchOwnerInfo } from '$lib/pkg/fetch';

/** @type {import('./$types').PageLoad} */
export async function load({ params }) {
	const owner = await fetchOwnerInfo(params.id);

	return owner;
}

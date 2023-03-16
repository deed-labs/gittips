<script>
	import Header from '$lib/components/Header.svelte';
	import TONDiamondBlueLogo from '$lib/images/ton_diamond_blue.png';
	import { bigIntToFloat } from '$lib/utils';

	/** @type {import('./$types').PageData} */
	export let data;
</script>

<div>
	<Header
		breadcrumbs={[
			{ name: 'explore', href: '' },
			{ name: 'bounties', href: '' }
		]}
	/>

	<div class="p-10 sm:w-full md:w-2/3 mx-auto">
		<h1 class="text-4xl font-bold mb-5">Available bounties</h1>
		<div class="overflow-x-auto w-full border border-secondary rounded rounded-xl">
			<table class="table w-full">
				<!-- head -->
				<thead>
					<tr>
						<th>Title</th>
						<th>Owner</th>
						<th>Reward</th>
						<th />
					</tr>
				</thead>
				<tbody>
					<!-- rows -->
					{#each data.bounties as bounty}
						<tr class="hover">
							<td>
								<div class="flex items-center space-x-3">
									<div class="avatar">
										<div class="w-10 h-10 rounded">
											<img src={bounty.ownerAvatarUrl} alt="Avatar Tailwind CSS Component" />
										</div>
									</div>
									<div>
										<div class="font-bold">{bounty.title}</div>
									</div>
								</div>
							</td>
							<td>
								<a class="link link-primary" href={bounty.ownerUrl} target="_blank" rel="noreferrer"
									>{bounty.owner}</a
								>
								<br />
								<span class="badge badge-ghost badge-sm">{bounty.ownerType}</span>
							</td>
							<td>
								<div class="flex flex-row items-center gap-1">
									<img src={TONDiamondBlueLogo} alt="ethereum logo" width={17} />
									<p class="text-lg">{bigIntToFloat(bounty.reward, 9, 2)}</p>
								</div>

								<!--
									TODO
									<div class="text-sm opacity-50">~ ${bounty.rewardUSD}</div>
								-->
							</td>
							<th>
								<a class="link link-primary" target="_blank" rel="noreferrer" href={bounty.url}
									><svg
										class="h-6 w-6"
										xmlns="http://www.w3.org/2000/svg"
										xmlns:xlink="http://www.w3.org/1999/xlink"
										viewBox="0 0 24 24"
										><g
											fill="none"
											stroke="currentColor"
											stroke-width="2"
											stroke-linecap="round"
											stroke-linejoin="round"
											><path d="M11 7H6a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-5" /><path
												d="M10 14L20 4"
											/><path d="M15 4h5v5" /></g
										></svg
									></a
								>
							</th>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>
	</div>
</div>

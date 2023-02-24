<script>
	import Header from '$lib/components/Header.svelte';
	import TONLogo from '$lib/images/ton_logo.png';

	// TODO: replace with real data
	let data = [
		{
			id: '4',
			title: 'Fix minor bug in interface',
			owner: 'deed-labs',
			ownerType: 'organization',
			ownerURL: 'https://github.com/deed-labs',
			link: 'https://github.com/deed-labs/gittips/issues',
			reward: '10.0',
			rewardUSD: '23.4'
		},
		{
			id: '3',
			title: 'Write tests for parser',
			owner: 'deed-labs',
			ownerType: 'organization',
			ownerURL: 'https://github.com/deed-labs',
			link: 'https://github.com/deed-labs/gittips/issues',
			reward: '50.0',
			rewardUSD: '23.4'
		},
		{
			id: '2',
			title: 'Update fonts',
			owner: 'deed-labs',
			ownerType: 'organization',
			ownerURL: 'https://github.com/deed-labs',
			link: 'https://github.com/deed-labs/gittips/issues',
			reward: '5.0',
			rewardUSD: '23.4'
		},
		{
			id: '1',
			title: 'Implement OAuth authorization',
			owner: 'deed-labs',
			ownerType: 'organization',
			ownerURL: 'https://github.com/deed-labs',
			link: 'https://github.com/deed-labs/gittips/issues',
			reward: '250.0',
			rewardUSD: '23.4'
		}
	];
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
		<div class="overflow-x-auto w-full">
			<table class="table w-full">
			  <!-- head -->
			  <thead>
				<tr>
				  <th>Title</th>
				  <th>Owner</th>
				  <th>Reward</th>
				  <th></th>
				</tr>
			  </thead>
			  <tbody>
				<!-- rows -->
				{#each data as bounty }
				<tr class="hover">
					<td>
					  <div class="flex items-center space-x-3">
						<div class="avatar">
						  <div class="w-10 h-10 rounded">
							<img src="https://avatars.githubusercontent.com/u/98539501?s=200&v=4" alt="Avatar Tailwind CSS Component" />
						  </div>
						</div>
						<div>
						  <div class="font-bold">{bounty.title}</div>
						</div>
					  </div>
					</td>
					<td>
					  <a class="link link-primary" href={bounty.ownerURL} target="_blank" rel="noreferrer" >{bounty.owner}</a>
					  <br/>
					  <span class="badge badge-ghost badge-sm">{bounty.ownerType}</span>
					</td>
					<td>
						<div class="flex flex-row items-center gap-1">
							<img src={TONLogo} alt="ethereum logo" width={25} />
							<p class="text-lg">{bounty.reward}</p>
						</div>
						<div class="text-sm opacity-50">~ ${bounty.rewardUSD}</div>
					</td>
					<th>
						<a class="link link-primary" target="_blank" rel="noreferrer" href={bounty.link}
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

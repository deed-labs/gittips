<script>
	import Header from '$lib/components/Header.svelte';
	import TONLogo from '$lib/images/ton_logo.png';

	// TODO: replace with real data
	let data = [
		{
			id: '4',
			title: 'Fix minor bug in interface',
			owner: 'deed-labs',
			link: 'https://github.com/deed-labs/openroll/issues',
			reward: '10.0'
		},
		{
			id: '3',
			title: 'Write tests for parser',
			owner: 'deed-labs',
			link: 'https://github.com/deed-labs/openroll/issues',
			reward: '50.0'
		},
		{
			id: '2',
			title: 'Update fonts',
			link: 'https://github.com/deed-labs/openroll/issues',
			owner: 'deed-labs',
			reward: '5.0'
		},
		{
			id: '1',
			title: 'Implement OAuth authorization',
			owner: 'deed-labs',
			link: 'https://github.com/deed-labs/openroll/issues',
			reward: '250.0'
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

	<div class="p-10 xs:w-full md:w-8/10 mx-auto">
		<div class="grid xs:grid-cols-1 md:grid-cols-3 gap-5 my-10">
			{#each data as bounty}
				<div class="card bg-base-300 shadow-md">
					<div class="card-body">
						<h2 class="card-title">{bounty.title}</h2>
						<p class="py-2">{bounty.owner}</p>
						<div class="card-actions justify-between items-center">
							<div class="flex flex-row items-center gap-3">
								<img src={TONLogo} alt="ethereum logo" width={30} />
								<p class="text-xl">{bounty.reward}</p>
							</div>
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
						</div>
					</div>
				</div>
			{/each}
		</div>
	</div>
</div>

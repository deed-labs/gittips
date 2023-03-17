<script lang="ts">
	import Header from '$lib/components/Header.svelte';
	import TONDiamondBlueLogo from '$lib/images/ton_diamond_blue.png';
	import { fetchOwnerInfo } from '$lib/pkg/fetch';
	import { addBudgetMessage, withdrawBudgetMessage } from '$lib/pkg/txs/ton';
	import { TON } from '$lib/stores/network';
	import { storage } from '$lib/stores/storage';
	import type { OwnerInfo } from '$lib/types';
	import { bigIntToFloat } from '$lib/utils';
	import { onMount } from 'svelte';
	import { fade } from 'svelte/transition';
	import CoinGecko from 'coingecko-api';

	const { connected } = $TON;

	let data: OwnerInfo = {
		name: '',
		avatarUrl: '',
		totalBudget: '',
		availableBudget: '',
		totalBounties: 0,
		availableBounties: 0,
		bounties: []
	};
	let donePercentage = '0';

	onMount(async () => {
		let info = await fetchOwnerInfo($storage.ownerId);

		const cgClient = new CoinGecko();
		const tonPriceData = (
			await cgClient.simple.price({ ids: 'the-open-network', vs_currencies: 'usd' })
		).data;

		info.bounties.forEach((b) => {
			b.rewardUSD = +(
				Number(bigIntToFloat(b.reward, 9, 2)) * tonPriceData['the-open-network']['usd']
			).toFixed(2);
			console.log(b.rewardUSD);
		});

		const remainingNumber = info.totalBounties - info.availableBounties;
		donePercentage =
			data.totalBounties > 0 ? ((remainingNumber * 100) / info.totalBounties).toFixed(0) : '00';

		data = info;
	});

	let errorMessage = '';
	const showError = (msg: string) => {
		errorMessage = msg;
		setTimeout(() => {
			errorMessage = '';
		}, 2000);
	};

	let tonAmount = 0;

	const addBudget = async () => {
		let wallet = TON.getConnectedWallet();
		if (!wallet) return;

		try {
			await wallet.sendTransaction(addBudgetMessage(tonAmount.toString()));

			(document.getElementById('add-funds-modal') as HTMLInputElement).checked = false;
		} catch (e) {
			console.error(e);
			showError('Failed to send transaction.');
		}
	};

	const withdrawBudget = async () => {
		let wallet = TON.getConnectedWallet();
		if (!wallet) return;

		try {
			await wallet.sendTransaction(withdrawBudgetMessage(tonAmount.toString()));

			(document.getElementById('withdrawal-modal') as HTMLInputElement).checked = false;
		} catch (e) {
			console.error(e);
			showError('Failed to send transaction.');
		}
	};
</script>

{#if errorMessage !== ''}
	<div class="toast toast-top toast-center z-20" out:fade>
		<div class="alert alert-error w-96">
			<div>
				<span>{errorMessage}</span>
			</div>
		</div>
	</div>
{/if}

<div>
	<Header breadcrumbs={[{ name: 'budget', href: '' }]} hideGitHubButton={true} />

	<div class="p-10 sm:w-full md:w-2/3 mx-auto max-w-screen-2xl">
		{#if !$connected}
			<div class="f-full text-center mt-24">
				<div class="flex justify-center">
					<svg class="text-center stroke-error h-10 w-10" fill="none" viewBox="0 0 24 24"
						><path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
						/></svg
					>
				</div>
				<h1 class="text-4xl font-bold py-6">Connect your wallet to see this page!</h1>
			</div>
		{:else}
			<div class="flex flex-col items-center w-full mb-12">
				<h1 class="text-4xl font-bold my-5">{data.name}</h1>

				<div class="flex flex-col gap-5 w-full items-start">
					<div
						class="stats stats-vertical md:stats-horizontal shadow border border-secondary w-full bg-base-200"
					>
						<div class="stat">
							<div class="stat-title">Total budget</div>
							<div class="flex flex-row items-center gap-1">
								<img src={TONDiamondBlueLogo} alt="ton logo" width={25} />
								<p class="stat-value text-secondary ">{bigIntToFloat(data.totalBudget, 9, 2)}</p>
							</div>
						</div>

						<div class="stat">
							<div class="stat-title">Available budget</div>
							<div class="flex flex-row items-center gap-1">
								<img src={TONDiamondBlueLogo} alt="ton logo" width={25} />
								<p class="stat-value text-secondary ">
									{bigIntToFloat(data.availableBudget, 9, 2)}
								</p>
							</div>
						</div>

						<div class="stat">
							<div class="stat-figure text-secondary">
								<div class="avatar">
									<div class="w-16 rounded-full">
										<img src={data.avatarUrl} alt="Org logo" />
									</div>
								</div>
							</div>
							<div class="stat-value">{donePercentage}%</div>
							<div class="stat-title">Bounties done</div>
							<div class="stat-desc text-secondary">{data.availableBounties} tasks remaining</div>
						</div>
					</div>
					<div class="flex flex-row gap-2 items-center justify-center">
						<label for="add-funds-modal" class="btn btn-primary btn-sm">Add funds</label>
						<label for="withdrawal-modal" class="btn btn-secondary btn-sm">Withdraw</label>
					</div>
				</div>
			</div>

			<h1 class="text-3xl font-bold my-5">Bounties</h1>
			<div class="overflow-x-auto w-full border border-secondary rounded rounded-2xl">
				<table class="table w-full">
					<!-- head -->
					<thead>
						<tr>
							<th>Title</th>
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
										<div class="font-bold">{bounty.title}</div>
									</div>
								</td>
								<td>
									<div class="flex flex-row items-center gap-1">
										<img src={TONDiamondBlueLogo} alt="ton logo" width={17} />
										<p class="text-lg">{bigIntToFloat(bounty.reward, 9, 2)}</p>
									</div>
									<div class="text-sm opacity-50">~ ${bounty.rewardUSD}</div>
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
		{/if}
	</div>
</div>

<!-- Add funds modal -->

<input type="checkbox" id="add-funds-modal" class="modal-toggle" />
<div class="modal modal-bottom sm:modal-middle">
	<div class="modal-box">
		<h3 class="font-bold text-2xl">Add funds</h3>
		<label
			for="add-funds-modal"
			class="btn btn-primary btn-outline btn-sm btn-circle absolute right-2 top-2">✕</label
		>
		<p class="my-5 text-gray-400">Enter the amount to be transferred to the contract balance.</p>
		<div class="flex flex-row items-center gap-5 w-full">
			<div>
				<img src={TONDiamondBlueLogo} alt="ton logo" width={35} />
			</div>
			<div class="w-full">
				<input
					type="number"
					placeholder="10.000"
					class="input input-bordered border w-full text-right"
					min="0.5"
					step="0.001"
					bind:value={tonAmount}
				/>
			</div>
			<button class="btn btn-primary" on:click={addBudget}>Confirm</button>
		</div>
		<div class="modal-action" />
	</div>
</div>

<!-- Withdrawal modal -->

<input type="checkbox" id="withdrawal-modal" class="modal-toggle" />
<div class="modal modal-bottom sm:modal-middle">
	<div class="modal-box">
		<h3 class="font-bold text-2xl">Withdrawal</h3>
		<label
			for="withdrawal-modal"
			class="btn btn-primary btn-outline btn-sm btn-circle absolute right-2 top-2">✕</label
		>
		<p class="mt-5 text-gray-400">Enter the amount to withdraw from the contract balance.</p>
		<p class="mb-5 mt-2">Maximum - {bigIntToFloat(data.availableBudget, 9, 4)}</p>
		<div class="flex flex-row items-center gap-5 w-full">
			<div>
				<img src={TONDiamondBlueLogo} alt="ton logo" width={35} />
			</div>
			<div class="w-full">
				<input
					type="number"
					placeholder="10.000"
					class="input input-bordered border w-full text-right"
					min="0.000"
					step="0.001"
					bind:value={tonAmount}
				/>
			</div>
			<button class="btn btn-primary" on:click={withdrawBudget}>Confirm</button>
		</div>
		<div class="modal-action" />
	</div>
</div>

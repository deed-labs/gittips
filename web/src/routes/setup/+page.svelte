<script lang="ts">
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { setupInstallation } from '$lib/pkg/fetch';
	import { storage, type LocalStorageData } from '$lib/stores/storage';

	let data: LocalStorageData = JSON.parse(localStorage['data']);
	let installationId = Number($page.url.searchParams.get('installation_id')) ?? 0;
	let action = $page.url.searchParams.get('setup_action');

	let failed: boolean;

	onMount(async () => {
		if (!action || action !== 'install') return;

		try {
			let installationInfo = await setupInstallation(data.wallet_address, installationId);

			storage.set({
				...$storage,
				bot_installation_done: true,
				installed: installationInfo.installed,
				name: installationInfo.name,
				ownerId: installationInfo.ownerId
			});
		} catch (e) {
			failed = true;
			console.error(e);
		}
	});
</script>

<div class="hero min-h-screen bg-base-200">
	<div class="hero-content text-center">
		<div class="max-w-md">
			{#if failed}
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
				<h1 class="text-4xl font-bold py-6">Something went wrong</h1>
			{:else}
				<div class="flex justify-center">
					<svg class="text-center stroke-success h-10 w-10" fill="none" viewBox="0 0 24 24"
						><path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
						/></svg
					>
				</div>
				<h1 class="text-4xl font-bold py-6">You are all set!</h1>
			{/if}
			<button
				class="btn btn-primary"
				on:click={() => {
					window.close();
				}}>Close</button
			>
		</div>
	</div>
</div>

<script>
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { setupInstallation } from '$lib/pkg/fetch';
	import { storage } from '$lib/stores/storage';

	let address = $storage.wallet_address;
	let installationId = Number($page.url.searchParams.get('installation_id')) ?? 0;

	// TODO: show error of address or installation id is null.
	// This means that the installation process was corrupted.

	onMount(async () => {
		await setupInstallation(address, installationId);
		storage.set({ ...$storage, bot_installation_done: true });
	});
</script>

<div class="hero min-h-screen bg-base-200">
	<div class="hero-content text-center">
		<div class="max-w-md">
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
			<button
				class="btn btn-primary"
				on:click={() => {
					window.close();
				}}>Close</button
			>
		</div>
	</div>
</div>

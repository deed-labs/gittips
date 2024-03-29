<script lang="ts">
	import GittipsLogo from '$lib/images/gittips_logo.png';
	import TONDiamondWhiteLogo from '$lib/images/ton_diamond_white.png';
	import TONDiamondBlueLogo from '$lib/images/ton_diamond_blue.png';
	import * as QROptions from './qr.json';

	import QRCodeStyling from 'qr-code-styling';

	import GitHubLogo from '$lib/images/github_logo.png';
	import { TON, type WalletStore } from '$lib/stores/network';
	import { fade } from 'svelte/transition';
	import { base } from '$app/paths';
	import type { InstallationInfo } from '$lib/types';
	import { fetchInstallationInfo } from '$lib/pkg/fetch';
	import { defaultData, storage, type LocalStorageData } from '$lib/stores/storage';
	import { shortAccountString } from '$lib/utils';
	import type { Unsubscriber } from 'svelte/store';

	const { connected, address, wallets } = $TON;

	const qrCode = new QRCodeStyling(QROptions as object);

	type Breadcrumb = {
		name: string;
		href: string;
	};
	export let breadcrumbs: Breadcrumb[] = [];

	let errorMessage = '';
	const showError = (msg: string) => {
		errorMessage = msg;
		setTimeout(() => {
			errorMessage = '';
		}, 2000);
	};

	let installationInfo: InstallationInfo = {
		installed: $storage.installed,
		name: $storage.name,
		ownerId: $storage.ownerId
	};
	let unsubscribe: Unsubscriber | null;

	const onInstalled = async () => {
		if (unsubscribe) unsubscribe();

		(document.getElementById('install-modal') as HTMLInputElement).checked = false;
	};

	const startInstallation = async () => {
		$storage.bot_installation_done = false;

		window.addEventListener('storage', () => {
			let data: LocalStorageData = JSON.parse(localStorage['data']);
			if (data.bot_installation_done) {
				installationInfo = {
					installed: data.installed,
					name: data.name,
					ownerId: data.ownerId
				};
				data.bot_installation_done = false;
				storage.set(data);
				onInstalled();
			}
		});
	};

	const onConnected = async () => {
		(document.getElementById('qr-modal') as HTMLInputElement).checked = false;
		installationInfo = await fetchInstallationInfo($address);
		storage.set({
			...$storage,
			wallet_address: $address,
			installed: installationInfo.installed,
			name: installationInfo.name,
			ownerId: installationInfo.ownerId
		});
	};

	const connect = async (wallet: WalletStore) => {
		if (!wallet.available) {
			showError('Wallet is not installed.');
			return;
		}

		if (wallet.injected) {
			await wallet.connectInjected(onConnected);
		} else {
			let connectionLink = await wallet.connectExternal(onConnected);

			qrCode.update({
				data: connectionLink,
				image: GittipsLogo
			});
			qrCode.append(document.getElementById('qr-code')!);

			(document.getElementById('qr-modal') as HTMLInputElement).checked = true;
		}
	};

	let isDisconnectingModalOpen = false;
	const disconnect = async () => {
		await TON.disconnect();
		isDisconnectingModalOpen = false;

		storage.set(defaultData);
	};

	let classProps = '';

	export let hideGitHubButton: boolean = false;
	export { classProps as class };
</script>

<div class={'navbar max-w-screen-2xl m-auto bg-base-100 p-5 ' + classProps}>
	<div class="flex-1">
		<a class="btn btn-ghost normal-case text-xl text-white" href={base + '/'}
			><img src={GittipsLogo} width={35} class="mr-2" alt="logo" />Gittips</a
		>
		<div class="text-sm breadcrumbs text-info mx-4 hidden md:block">
			<ul>
				{#each breadcrumbs as { name, href }}
					{#if href}
						<li>
							<a {href}>{name}</a>
						</li>
					{:else}
						<li>{name}</li>
					{/if}
				{/each}
			</ul>
		</div>
	</div>
	<div class="flex flex-col-reverse md:flex-row items-end gap-5">
		{#if $connected && !hideGitHubButton}
			<div>
				{#if !installationInfo.installed}
					<label
						for="install-modal"
						class="btn btn-github btn-outline text-white rounded-full capitalize"
					>
						<img class="mr-2" src={GitHubLogo} alt="github logo" width={25} />Add to GitHub
					</label>
				{:else}
					<a
						class="btn btn-github btn-outline text-white rounded-full capitalize"
						href={base + '/budget'}
					>
						<img
							class="mr-2"
							src={GitHubLogo}
							alt="github logo"
							width={25}
						/>{installationInfo.name}
					</a>
				{/if}
			</div>
		{/if}
		<div>
			{#if !$connected}
				<div class="dropdown dropdown-bottom dropdown-end">
					<label tabindex="0" class="btn btn-primary text-white rounded-full capitalize"
						><img class="mr-2" src={TONDiamondWhiteLogo} alt="ton logo" width={18} />Connect TON</label
					>
					<ul
						tabindex="0"
						class="dropdown-content menu p-2 mt-2 shadow bg-base-100 rounded-box w-52 border"
					>
						{#each Object.entries(wallets) as [name, wallet]}
							{#if wallet.supported}
								<li>
									<!-- svelte-ignore a11y-click-events-have-key-events -->
									<label
										for=""
										on:click={() => {
											connect(wallet);
										}}
									>
										{name}
									</label>
								</li>
							{/if}
						{/each}
					</ul>
				</div>
			{:else}
				<label for="disconnect-modal" class="btn btn-outline btn-primary rounded-full"
					><img
						class="mr-2"
						src={TONDiamondBlueLogo}
						alt="ton logo"
						width={18}
					/>{shortAccountString(10, 5, $address ?? '')}</label
				>
			{/if}
		</div>
	</div>
</div>

{#if errorMessage !== ''}
	<div class="toast toast-top toast-center z-20" out:fade>
		<div class="alert alert-error w-96">
			<div>
				<span>{errorMessage}</span>
			</div>
		</div>
	</div>
{/if}

<input
	type="checkbox"
	id="disconnect-modal"
	class="modal-toggle"
	bind:checked={isDisconnectingModalOpen}
/>
<div class="modal modal-bottom sm:modal-middle">
	<div class="modal-box">
		<h3 class="font-bold text-lg">You are going to disconnect</h3>
		<label for="disconnect-modal" class="btn btn-sm btn-circle absolute right-2 top-2">✕</label>
		<div class="modal-action">
			<button class="btn btn-sm" on:click={disconnect}>Disconnect</button>
		</div>
	</div>
</div>

<!--- Installation modal -->

<input type="checkbox" id="install-modal" class="modal-toggle" />
<div class="modal modal-bottom sm:modal-middle">
	<div class="modal-box">
		<h3 class="font-bold text-lg">Install app</h3>
		<label for="install-modal" class="btn btn-sm btn-circle absolute right-2 top-2">✕</label>
		<p class="py-4">
			Your GitHub organization will be linked to {shortAccountString(10, 5, $address ?? '')} address
		</p>
		<div class="modal-action">
			<a
				href="https://github.com/apps/gittips-bot"
				target="_blank"
				rel="noreferrer"
				class="btn btn-sm"
				on:click={startInstallation}>Install</a
			>
		</div>
	</div>
</div>

<!--- QR modal -->

<input type="checkbox" id="qr-modal" class="modal-toggle" />
<label for="qr-modal" class="modal cursor-pointer">
	<label
		class="modal-box relative w-fit text-center flex flex-col gap-2 items-center py-4 px-6"
		for=""
	>
		<h3 class="text-lg font-bold">Connect TON</h3>
		<p class="text-gray-400 text-sm w-56">Scan the QR code with your phone's camera or wallet.</p>
		<div id="qr-code" />
		<p class="text-gray-400 text-sm w-56">
			We do not store your wallet credentials, so your TON is safe.
		</p>
	</label>
</label>

<script lang="ts">
	import TONLogo from '$lib/images/ton_logo.png';
	import TONConnectLogo from '$lib/images/ton_connect_logo.png';
	import * as QROptions from './qr.json';

	import QRCodeStyling from 'qr-code-styling';

	import GitHubLogo from '$lib/images/github_logo.png';
	import { shortAccountString } from '$lib/utils';
	import { TON } from '$lib/stores';

	const { connected, address, wallets } = $TON;

	const qrCode = new QRCodeStyling(QROptions as object);

	type Breadcrumb = {
		name: string;
		href: string;
	};
	export let breadcrumbs: Breadcrumb[] = [];

	let isDisconnectingModalOpen = false;

	const onConnected = () => {
		(document.getElementById('qr-modal') as HTMLInputElement).checked = false;
	};

	const connect = async () => {
		let connectionLink = await wallets.TonKeeper.connectExternal(onConnected);

		qrCode.update({
			data: connectionLink
		});
		qrCode.append(document.getElementById('qr-code')!);

		(document.getElementById('qr-modal') as HTMLInputElement).checked = true;
	};

	const disconnect = () => {
		TON.disconnect();
		isDisconnectingModalOpen = false;
	};

	let classProps = '';

	export { classProps as class };
</script>

<div class={'navbar bg-base-100 ' + classProps}>
	<div class="flex-1">
		<a class="btn btn-ghost normal-case text-xl" href="/"
			><img src={TONLogo} width={35} class="mr-3" alt="logo" />Gittips</a
		>
		<div class="text-sm breadcrumbs text-info mx-4">
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
	<div class="flex-none">
		<div>
			<a
				href="https://github.com/apps/gittips-test"
				target="_blank"
				rel="noreferrer"
				class="btn btn-github mr-4 text-white rounded-full capitalize"
				><img class="mr-2" src={GitHubLogo} alt="github logo" width={25} />Add to GitHub</a
			>
		</div>
		<div>
			{#if !$connected}
				<button class="btn btn-primary mr-4 text-white rounded-full capitalize"
					><img src={TONConnectLogo} alt="ton logo" width={40} on:click={connect} /> Connect TON</button
				>
			{:else}
				<label for="disconnect-modal" class="btn btn-outline btn-primary mr-4"
					>{shortAccountString(10, 5, $address ?? '')}</label
				>
			{/if}
		</div>
	</div>
</div>

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

<!--- QR modal -->

<input type="checkbox" id="qr-modal" class="modal-toggle" />
<label for="qr-modal" class="modal cursor-pointer">
	<label
		class="modal-box relative w-fit text-center flex flex-col gap-2 items-center py-4 px-6"
		for=""
	>
		<h3 class="text-lg font-bold">Connect TON</h3>
		<p class="text-neutral text-sm w-56">Scan the QR code with your phone's camera or Tonkeeper.</p>
		<div id="qr-code" />
		<p class="text-neutral text-sm w-56">
			We do not store your wallet credentials, so your TON is safe.
		</p>
	</label>
</label>

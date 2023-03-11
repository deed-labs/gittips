<script lang="ts">
	import IssueScreenshot from '$lib/images/issue_screenshot.png';
	import CommentScreenshot from '$lib/images/comment_screenshot.png';
	import PaymentScreenshot from '$lib/images/payment_screenshot.png';

	import Header from '$lib/components/Header.svelte';
	import { fade, fly, type FlyParams } from 'svelte/transition';
	import { inview } from 'svelte-inview';
	import { onMount } from 'svelte';

	let ready = false;
	onMount(() => (ready = true));

	// array for storing blocks animation statuses
	let isInView: boolean[] = new Array(6);

	if (window.innerWidth <= 500) {
		isInView[1] = true;
		isInView[2] = true;
		isInView[3] = true;
	}

	const animateFeature = (node: Element, args?: FlyParams) => {
		return window.innerWidth > 500 ? fly(node, args) : {};
	};
</script>

<Header class="bg-base-100" />

<div class="hero min-h-screen">
	<div class="hero-content text-center">
		{#if ready}
			<div class="max-w-4xl" in:fly={{ y: 200, duration: 1500 }}>
				<h1 class="text-5xl md:text-7xl font-bold text-white">
					Reward the community, fuel innovation
				</h1>
				<p class="text-xl py-10 text-gray-400">
					Gittips is a bot for GitHub that enables owners of open-source projects to reward
					contributors with cryptocurrency
				</p>
				<div class="flex flex-col md:flex-row items-center justify-center gap-4">
					<a
						class="btn btn-wide btn-primary btn-outline text-white font-bold rounded-full capitalize"
						target="_blank"
						rel="noreferrer"
						href="https://deed-labs.gitbook.io/gittips/">Docs</a
					>
					<a
						class="btn btn-wide btn-primary text-white font-bold rounded-full capitalize"
						href="/explore/bounties">Explore</a
					>
				</div>
			</div>
		{/if}
	</div>
</div>

<div
	class="flex flex-col items-center text-center justify-center p-5 md:p-28 h-64"
	use:inview={{ unobserveOnEnter: true, rootMargin: '-20%' }}
	on:change={({ detail }) => {
		isInView[0] = detail.inView;
	}}
>
	{#if isInView[0]}
		<div in:fade={{ duration: 1500 }}>
			<h2 class="text-4xl text-white font-bold">
				Our solution does not interfere with your <span class="text-primary">usual workflow</span>
			</h2>
			<p class="text-2xl text-gray-400 mt-5">
				Create tasks and send payments without leaving GitHub
			</p>
		</div>
	{/if}
</div>

<div class="flex flex-row items-stretch py-12 w-full max-w-screen-2xl m-auto text-white">
	<div class="flex flex-col justify-center items-center w-2/12 md:w-1/12">
		<ul class="steps steps-vertical h-full">
			<li data-content="✨" class={'step ' + (isInView[1] ? 'step-primary' : '')} />
			<li data-content="💰" class={'step ' + (isInView[2] ? 'step-primary' : '')} />
			<li data-content="💎" class={'step ' + (isInView[3] ? 'step-primary' : '')} />
		</ul>
	</div>
	<div class="w-10/12 md:w-11/12 p-4 md:p-0 flex flex-col gap-12 md:gap-5">
		<div
			class="flex flex-col md:flex-row items-center gap-5 md:p-12 md:h-96"
			use:inview={{ unobserveOnEnter: true, rootMargin: '-20%' }}
			on:change={({ detail }) => {
				if (window.innerWidth > 500) isInView[1] = detail.inView;
			}}
		>
			{#if isInView[1]}
				<div class="w-full md:w-1/2" in:animateFeature={{ x: -200, duration: 1000 }}>
					<h1 class="text-4xl font-bold">Automated Bounty Creation</h1>
					<p class="text-xl text-gray-400 mt-5">
						We automate the process of creating bounties from issues on GitHub. <br /> Find developers
						for tasks without breaking your usual flow.
					</p>
					<p class="text-xl text-gray-400">
						Just add <span class="badge badge-primary badge-outline">bounty</span> label to an issue.
					</p>
				</div>
				<div
					class="flex justify-center w-full md:w-1/2"
					in:animateFeature={{ x: 200, duration: 1000 }}
				>
					<img
						class="card border border-primary image-full w-full md:w-5/6"
						src={IssueScreenshot}
						alt="issue example"
					/>
				</div>
			{/if}
		</div>

		<div
			class="flex flex-col-reverse md:flex-row items-center gap-5 md:p-12 md:h-96"
			use:inview={{ unobserveOnEnter: true, rootMargin: '-20%' }}
			on:change={({ detail }) => {
				if (window.innerWidth > 500) isInView[2] = detail.inView;
			}}
		>
			{#if isInView[2]}
				<div class="w-full md:w-1/2" in:animateFeature={{ x: -200, duration: 1000 }}>
					<img
						class="card border border-primary w-full md:w-5/6"
						src={CommentScreenshot}
						alt="issue example"
					/>
				</div>
				<div class="w-full md:w-1/2" in:animateFeature={{ x: 200, duration: 1000 }}>
					<h1 class="text-4xl font-bold">Budget Control</h1>
					<p class="text-xl text-gray-400 mt-5">
						Each organization and user has its own balance reserved on the smart contract. So they
						can't accidentally spend more than planned.
					</p>
				</div>
			{/if}
		</div>

		<div
			class="flex flex-col md:flex-row items-center gap-5 md:p-12 h-full md:h-96"
			use:inview={{ unobserveOnEnter: true, rootMargin: '-20%' }}
			on:change={({ detail }) => {
				if (window.innerWidth > 500) isInView[3] = detail.inView;
			}}
		>
			{#if isInView[3]}
				<div class="w-full md:w-1/2" in:animateFeature={{ x: -200, duration: 1000 }}>
					<h1 class="text-4xl font-bold">Wallet Integration</h1>
					<p class="text-xl text-gray-400 mt-5">
						Gittips integrates with blockchain, allowing contributors to receive their rewards
						directly in their wallets. <br /> This makes it easy for contributors to use rewards as they
						see fit.
					</p>
				</div>
				<div
					class="flex justify-center w-full md:w-1/2"
					in:animateFeature={{ x: 200, duration: 1000 }}
				>
					<img
						class="card border border-primary image-full w-full md:w-5/6"
						src={PaymentScreenshot}
						alt="issue example"
					/>
				</div>
			{/if}
		</div>
	</div>
</div>

<div
	class="py-24 text-center"
	use:inview={{ unobserveOnEnter: true, rootMargin: '-20%' }}
	on:change={({ detail }) => {
		isInView[4] = detail.inView;
	}}
>
	{#if isInView[4]}
		<div in:fade={{ duration: 1500 }}>
			<h1 class="text-white text-4xl">Empower your developer community</h1>
			<a
				class="btn btn-wide btn-primary text-white font-bold rounded-full mt-12"
				target="_blank"
				rel="noreferrer"
				href="https://deed-labs.gitbook.io/gittips/">Try it out</a
			>
		</div>
	{/if}
</div>

<div
	class="flex items-center justify-center max-w-screen-2xl w-full md:h-96 m-auto p-5 py-12"
	use:inview={{ unobserveOnEnter: true, rootMargin: '-20%' }}
	on:change={({ detail }) => {
		isInView[5] = detail.inView;
	}}
>
	{#if isInView[5]}
		<div class="card bg-gray-700 w-full md:w-2/3 text-center p-12" in:fade={{ duration: 1500 }}>
			<h1 class="text-white text-4xl">Roadmap 2023</h1>
			<ul class="steps steps-vertical md:steps-horizontal mt-12 text-white">
				<li data-content="Q1" class="step step-primary">Launch</li>
				<li data-content="Q1" class="step">NFT Rewards Support</li>
				<li data-content="Q2" class="step">Deeper GitHub integration</li>
			</ul>
		</div>
	{/if}
</div>

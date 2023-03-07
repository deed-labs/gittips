<script lang="ts">
	import IssueScreenshot from '$lib/images/issue_screenshot.png'
	import CommentScreenshot from '$lib/images/comment_screenshot.png'
	import PaymentScreenshot from '$lib/images/payment_screenshot.png'


	import Header from '$lib/components/Header.svelte';
	import { fade, fly } from 'svelte/transition';
	import { inview } from 'svelte-inview';
	import { onMount } from 'svelte';

	let ready = false;
	onMount(() => (ready = true));

	// array for storing block animation status
	let isInView: boolean[] = new Array(4);
</script>

<Header class="bg-base-200"/>

<div class="hero min-h-screen bg-base-200">
	<div class="hero-content text-center">
		{#if ready}
			<div class="max-w-4xl" in:fly={{ y: 200, duration: 1500 }}>
				<h1 class="text-7xl font-bold text-white">Reward the community, fuel innovation</h1>
				<p class="text-lg py-10 text-gray-400">
					Gittips is a bot for GitHub that enables owners of open-source projects to reward contributors with cryptocurrency
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

<div class="flex flex-col items-center text-center justify-center p-28"
	use:inview={{ unobserveOnEnter: true, rootMargin: '-20%' }}
	on:change={({ detail }) => {
		isInView[0] = detail.inView;
	}}
>
	{#if isInView[0]}
		<div in:fly={{ y: 100, duration: 1500 }}>
			<h2 class="text-4xl text-white font-bold">Our solution does not interfere with your <span class="text-primary">usual workflow</span></h2>
			<p class="text-2xl text-gray-400 mt-5">
				Create tasks and send payments without leaving GitHub
			</p>
		</div>
	{/if}
</div>

<div class="flex flex-row items-stretch py-12 max-w-screen-2xl m-auto bg-base-200 text-white">
	<div class="flex flex-col justify-center items-center w-2/12 md:w-1/12">
		<ul class="steps steps-vertical h-full">
			<li data-content="✨" class={"step " + (isInView[1] ? "step-primary" : "")}></li>
			<li data-content="💰" class={"step " + (isInView[2] ? "step-primary" : "")}></li>
			<li data-content="💎" class={"step " + (isInView[3] ? "step-primary" : "")}></li>
		  </ul>
	</div>
	<div class="w-10/12 md:w-11/12 h-full">
		<div class="flex flex-col md:flex-row items-center gap-5 p-12 md:h-96"
			use:inview={{ unobserveOnEnter: true, rootMargin: '-20%' }}
			on:change={({ detail }) => {
				isInView[1] = detail.inView;
			}}
		>
			{#if isInView[1]}
				<div class="w-full md:w-1/2"
					in:fly={{ x: -200, duration: 1000 }}
				>
					<h1 class="text-4xl font-bold">Automated Bounty Creation</h1>
					<p  class="text-xl text-gray-400 mt-5">We automate the process of creating bounties from issues on GitHub. <br/> Find developers for tasks without breaking your usual flow.</p>
					<p class="text-xl text-gray-400">Just add <span class="badge badge-primary badge-outline">bounty</span> label to an issue.</p>
				</div>
				<div class="flex justify-center w-full md:w-1/2"
					in:fly={{ x: 200, duration: 1000 }}
				>
					<img class="card border border-primary image-full w-full md:w-5/6" src={IssueScreenshot} alt="issue example"/>
				</div>
			{/if}
		</div>

		<div class="flex flex-col-reverse md:flex-row items-center gap-5 p-12 md:h-96"
			use:inview={{ unobserveOnEnter: true, rootMargin: '-20%' }}
			on:change={({ detail }) => {
				isInView[2] = detail.inView;
			}}
		>
			{#if isInView[2]}
				<div class="w-full md:w-1/2"
					in:fly={{ x: -200, duration: 1000 }}
				>
					<img class="card border border-primary w-full md:w-5/6" src={CommentScreenshot} alt="issue example"/>	
				</div>
				<div class="w-full md:w-1/2"
					in:fly={{ x: 200, duration: 1000 }}
				>
					<h1 class="text-4xl font-bold">Budget Control</h1>
					<p  class="text-xl text-gray-400 mt-5">Each organization and user has its own balance reserved on the smart contract. So they can't accidentally spend more than planned.</p>
				</div>
			{/if}
		</div>

		<div class="flex flex-col md:flex-row items-center gap-5 p-12 h-full md:h-96"
			use:inview={{ unobserveOnEnter: true, rootMargin: '-20%' }}
			on:change={({ detail }) => {
				isInView[3] = detail.inView;
			}}
		>
			{#if isInView[3]}
				<div class="w-full md:w-1/2"
					in:fly={{ x: -200, duration: 1000 }}
				>
					<h1 class="text-4xl font-bold">Wallet Integration</h1>
					<p  class="text-xl text-gray-400 mt-5">Gittips integrates with blockchain, allowing contributors to receive their rewards directly in their wallets. <br/> This makes it easy for contributors to use rewards as they see fit.</p>
				</div>
				<div class="flex justify-center w-full md:w-1/2"
					in:fly={{ x: 200, duration: 1000 }}
				>
					<img class="card border border-primary image-full w-full md:w-5/6" src={PaymentScreenshot} alt="issue example"/>
				</div>
			{/if}
		</div>
	</div>
</div>

<div class="bg-accent py-12">
	<div class="text-center my-12">
		<h1 class="text-white text-4xl">Empower your developer community</h1>
		<a
			class="btn btn-wide btn-primary text-white font-bold rounded-full mt-12"
			target="_blank"
			rel="noreferrer"
			href="https://deed-labs.gitbook.io/gittips/">Try it out</a
		>
	</div>
</div>
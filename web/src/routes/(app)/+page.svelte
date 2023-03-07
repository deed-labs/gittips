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

	let isInView: {[key: number]: boolean} = {
		0: false,
		1: false,
		2: false,
	};
</script>

<Header class="bg-base-200"/>

<div class="hero min-h-screen bg-base-200">
	<div class="hero-content text-center">
		{#if ready}
			<div class="max-w-xl" in:fly={{ y: 200, duration: 1500 }}>
				<h1 class="text-5xl font-bold">Welcome to Gittips 👋</h1>
				<p class="py-8">
					Encourage developers to contribute to your projects by giving a crypto reward for each
					completed task and valuable issue
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

<div class="flex flex-row items-stretch py-12 max-w-screen-2xl m-auto bg-base-200">
	<div class="flex flex-col justify-center items-center w-2/12 md:w-1/12">
		<ul class="steps steps-vertical h-full">
			<li data-content="✨" class={"step " + (isInView[0] ? "step-primary" : "")}></li>
			<li data-content="💰" class={"step " + (isInView[1] ? "step-primary" : "")}></li>
			<li data-content="💎" class={"step " + (isInView[2] ? "step-primary" : "")}></li>
		  </ul>
	</div>
	<div class="w-10/12 md:w-11/12 h-full">
		<div class="flex flex-col md:flex-row items-center gap-5 p-12 md:h-96"
			use:inview={{ unobserveOnEnter: true, rootMargin: '-20%' }}
			on:change={({ detail }) => {
				isInView[0] = detail.inView;
			}}
		>
			{#if isInView[0]}
				<div class="w-full md:w-1/2"
					in:fly={{ x: -200, duration: 1000 }}
				>
					<h1 class="text-4xl font-bold">Automated Bounty Creation</h1>
					<p  class="text-xl pt-5">We automate the process of creating bounties from issues on GitHub. <br/> Find developers for tasks without breaking your usual flow.</p>
					<p class="text-xl">Just add <span class="badge badge-primary badge-outline">bounty</span> label to an issue.</p>
				</div>
				<div class="flex justify-center w-full md:w-1/2"
					in:fly={{ x: 200, duration: 1000 }}
				>
					<img class="card shadow-lg shadow-primary image-full w-full md:w-5/6" src={IssueScreenshot} alt="issue example"/>
				</div>
			{/if}
		</div>

		<div class="flex flex-col-reverse md:flex-row items-center gap-5 p-12 md:h-96"
			use:inview={{ unobserveOnEnter: true, rootMargin: '-20%' }}
			on:change={({ detail }) => {
				isInView[1] = detail.inView;
			}}
		>
			{#if isInView[1]}
				<div class="w-full md:w-1/2"
					in:fly={{ x: -200, duration: 1000 }}
				>
					<img class="card shadow-lg shadow-primary w-full md:w-5/6" src={CommentScreenshot} alt="issue example"/>	
				</div>
				<div class="w-full md:w-1/2"
					in:fly={{ x: 200, duration: 1000 }}
				>
					<h1 class="text-4xl font-bold">Budget Control</h1>
					<p  class="text-xl pt-5">Each organization and user has its own balance reserved on the smart contract. So they can't accidentally spend more than planned.</p>
				</div>
			{/if}
		</div>

		<div class="flex flex-col md:flex-row items-center gap-5 p-12 h-full md:h-96"
			use:inview={{ unobserveOnEnter: true, rootMargin: '-20%' }}
			on:change={({ detail }) => {
				isInView[2] = detail.inView;
			}}
		>
			{#if isInView[2]}
				<div class="w-full md:w-1/2"
					in:fly={{ x: -200, duration: 1000 }}
				>
					<h1 class="text-4xl font-bold">Wallet Integration</h1>
					<p  class="text-xl pt-5">Gittips integrates with blockchain, allowing contributors to receive their rewards directly in their wallets. <br/> This makes it easy for contributors to use rewards as they see fit.</p>
				</div>
				<div class="flex justify-center w-full md:w-1/2"
					in:fly={{ x: 200, duration: 1000 }}
				>
					<img class="card shadow-lg shadow-primary  image-full w-full md:w-5/6" src={PaymentScreenshot} alt="issue example"/>
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
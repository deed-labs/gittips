<script>
	import Header from '$lib/components/Header.svelte';
	import { fly } from 'svelte/transition';
	import { onMount } from 'svelte';
	let ready = false;
	onMount(() => (ready = true));
</script>

<Header class="bg-base-200" />

<div class="hero h-full bg-base-200">
	<div class="hero-content text-center">
		{#if ready}
			<div class="max-w-xl" in:fly={{ y: 200, duration: 1500 }}>
				<h1 class="text-5xl font-bold">Welcome to Gittips 👋</h1>
				<p class="py-8">
					Encourage developers to contribute to your projects by giving a crypto reward for each
					completed task and opened issue.
				</p>
				<div class="flex justify-center gap-4">
					<a
						class="btn btn-wide btn-primary btn-outline text-white font-bold rounded-full capitalize"
						href="/">Docs</a
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

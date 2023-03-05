<div class="flex flex-col h-screen justify-between">
	<slot />

	<footer class="footer footer-center p-5 bg-base-200 gap-5 text-base-content rounded">
		<div class="grid grid-flow-col gap-4">
			<a
				class="link link-hover"
				target="_blank"
				rel="noreferrer"
				href="https://github.com/deed-labs">GitHub</a
			>
			<a class="link link-hover" href="#">Telegram</a>
		</div>
		<div>
			<p>&#169; {new Date().getFullYear()} DEED Labs</p>
		</div>
	</footer>
</div>

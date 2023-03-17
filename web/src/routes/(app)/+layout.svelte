<div class="flex flex-col min-h-screen justify-between bg-base-100">
	<slot />

	<footer class="footer footer-center p-5 gap-5 text-white">
		<div class="grid grid-flow-col gap-4">
			<a
				class="link link-hover"
				target="_blank"
				rel="noreferrer"
				href="https://deed-labs.gitbook.io/gittips/">Docs</a
			>
			<a
				class="link link-hover"
				target="_blank"
				rel="noreferrer"
				href="https://github.com/deed-labs">GitHub</a
			>
			<a
				class="link link-hover"
				target="_blank"
				rel="noreferrer"
				href="https://t.me/+3B7rHMU0zro3NDli">Telegram</a
			>
		</div>
		<div>
			<p>&#169; {new Date().getFullYear()} DEED Labs</p>
		</div>
	</footer>
</div>

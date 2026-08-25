<script lang="ts">
	import '../app.css';

	let { children } = $props();
</script>

<div class="shell">
	<header>
		<a class="mark" href="/">
			<span class="dot" aria-hidden="true"></span>
			snakra
		</a>
	</header>

	<main>
		{@render children()}
	</main>

	<footer>voice notes here are disposable</footer>
</div>

<style>
	.shell {
		position: relative;
		z-index: 1;
		min-height: 100svh;
		display: grid;
		grid-template-rows: auto 1fr auto;
		padding: clamp(1.25rem, 4vw, 2.25rem);
		gap: 2rem;
	}

	header {
		display: flex;
		justify-content: center;
	}

	.mark {
		display: inline-flex;
		align-items: center;
		gap: 0.55em;
		color: var(--ink-dim);
		text-decoration: none;
		font-size: 0.8125rem;
		font-weight: 500;
		letter-spacing: 0.18em;
		text-transform: lowercase;
		transition: color 0.3s var(--ease);
	}

	.mark:hover {
		color: var(--ink);
	}

	.dot {
		width: 7px;
		height: 7px;
		border-radius: 50%;
		background: var(--accent);
		box-shadow: 0 0 12px rgb(255 90 71 / 0.55);
	}

	main {
		display: grid;
		place-items: center;
		width: 100%;
	}

	footer {
		text-align: center;
		color: var(--ink-faint);
		font-size: 0.75rem;
		letter-spacing: 0.04em;
	}
</style>

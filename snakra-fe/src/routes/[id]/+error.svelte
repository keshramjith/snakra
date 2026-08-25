<script lang="ts">
	import { page } from '$app/state';
</script>

<svelte:head>
	<title>snakra · gone</title>
</svelte:head>

<section class="stage">
	<h1>It's gone.</h1>
	<p class="sub">
		{page.status === 404
			? "Notes here don't stick around. This one has already faded — or the link was never a note at all."
			: 'Something went wrong reaching that note.'}
	</p>
	<a class="btn" href="/">Record a new one</a>
</section>

<style>
	.stage {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 1.35rem;
		width: min(100%, 27rem);
		text-align: center;
		animation: rise 0.6s var(--ease) both;
	}

	h1 {
		margin: 0;
		font-size: clamp(1.6rem, 6vw, 2.1rem);
		font-weight: 600;
		letter-spacing: -0.035em;
	}

	.sub {
		max-width: 30ch;
		color: var(--ink-dim);
		font-size: 0.9375rem;
		text-wrap: balance;
	}

	.btn {
		padding: 0.7rem 1.35rem;
		border-radius: 999px;
		background: var(--accent);
		color: #1a0b08;
		font-size: 0.9375rem;
		font-weight: 500;
		text-decoration: none;
		transition: background-color 0.25s var(--ease);
	}

	.btn:hover {
		background: var(--accent-soft);
	}
</style>

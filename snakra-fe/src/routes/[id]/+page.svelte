<script lang="ts">
	import { onDestroy } from 'svelte';
	import { page } from '$app/state';
	import Player from '$lib/Player.svelte';

	const id = $derived(page.params.id);

	let objectUrl: string | undefined;
	let copied = $state(false);
	let copyTimer: ReturnType<typeof setTimeout> | undefined;

	const convertURIToBinary = (dataURI: string) => {
		const BASE64_MARKER = ';base64,';
		const base64Index = dataURI.indexOf(BASE64_MARKER) + BASE64_MARKER.length;
		const raw = window.atob(dataURI.substring(base64Index));
		const arr = new Uint8Array(new ArrayBuffer(raw.length));
		for (let i = 0; i < raw.length; i++) {
			arr[i] = raw.charCodeAt(i);
		}
		return arr;
	};

	const setAudio = (binary: Uint8Array<ArrayBuffer>) => {
		if (objectUrl) URL.revokeObjectURL(objectUrl);
		objectUrl = URL.createObjectURL(new Blob([binary], { type: 'audio/ogg; codecs=opus' }));
		return objectUrl;
	};

	const getVn = async (id: string) => {
		const response = await fetch(`${import.meta.env.VITE_API_URL}/api/v1/vn/${id}`);
		if (!response.ok) throw new Error(`note unavailable (${response.status})`);
		const json = await response.json();
		return setAudio(convertURIToBinary(json.audio));
	};

	const copyLink = async () => {
		try {
			await navigator.clipboard.writeText(window.location.href);
			copied = true;
			clearTimeout(copyTimer);
			copyTimer = setTimeout(() => (copied = false), 2000);
		} catch {
			// clipboard unavailable — the URL bar still works
		}
	};

	onDestroy(() => {
		clearTimeout(copyTimer);
		if (objectUrl) URL.revokeObjectURL(objectUrl);
	});
</script>

<svelte:head>
	<title>snakra · a note for you</title>
</svelte:head>

<section class="stage">
	{#if id}
		{#await getVn(id)}
			<p class="eyebrow">unwrapping</p>
			<div class="skeleton" aria-hidden="true">
				<span class="knob"></span>
				<span class="bars">
					{#each Array(28) as _, i (i)}
						<span style="--d:{i * 40}ms"></span>
					{/each}
				</span>
			</div>
		{:then src}
			<p class="eyebrow">someone left you this</p>
			<Player {src} />
			<div class="actions">
				<button class="btn ghost" onclick={copyLink}>{copied ? 'Link copied' : 'Copy link'}</button>
				<a class="btn ghost" href="/">Record your own</a>
			</div>
		{:catch}
			<h1>It's gone.</h1>
			<p class="sub">
				Notes here don't stick around. This one has already faded — or the link was never a note at
				all.
			</p>
			<a class="btn primary" href="/">Record a new one</a>
		{/await}
	{/if}
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

	.eyebrow {
		color: var(--ink-faint);
		font-size: 0.6875rem;
		letter-spacing: 0.2em;
		text-transform: uppercase;
	}

	.actions {
		display: flex;
		flex-wrap: wrap;
		justify-content: center;
		gap: 0.6rem;
	}

	.btn {
		padding: 0.7rem 1.35rem;
		border-radius: 999px;
		font-size: 0.9375rem;
		font-weight: 500;
		text-decoration: none;
		transition:
			background-color 0.25s var(--ease),
			border-color 0.25s var(--ease),
			color 0.25s var(--ease);
	}

	.btn.primary {
		background: var(--accent);
		color: #1a0b08;
	}

	.btn.primary:hover {
		background: var(--accent-soft);
	}

	.btn.ghost {
		border: 1px solid var(--edge-strong);
		color: var(--ink-dim);
	}

	.btn.ghost:hover {
		border-color: var(--ink-faint);
		color: var(--ink);
	}

	/* ── loading ─────────────────────────────────────────────────────── */

	.skeleton {
		display: flex;
		align-items: center;
		gap: 0.9rem;
		width: 100%;
		padding: 0.7rem 0.9rem 0.7rem 0.7rem;
		background: var(--bg-raised);
		border: 1px solid var(--edge);
		border-radius: var(--radius);
	}

	.knob {
		flex: none;
		width: 42px;
		height: 42px;
		border-radius: 50%;
		background: var(--edge);
	}

	.bars {
		flex: 1;
		display: flex;
		align-items: center;
		gap: 3px;
		height: 40px;
	}

	.bars span {
		flex: 1;
		height: 40%;
		border-radius: 2px;
		background: var(--edge-strong);
		animation: breathe 1.6s var(--ease) infinite;
		animation-delay: var(--d);
	}

	@keyframes breathe {
		0%,
		100% {
			height: 22%;
			opacity: 0.35;
		}
		50% {
			height: 70%;
			opacity: 0.8;
		}
	}
</style>

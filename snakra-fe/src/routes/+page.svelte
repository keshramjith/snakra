<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import LiveWave from '$lib/LiveWave.svelte';
	import Player from '$lib/Player.svelte';
	import { formatTime } from '$lib/formatTime';

	type Stage = 'arming' | 'blocked' | 'ready' | 'recording' | 'review' | 'sending';

	let stage = $state<Stage>('arming');
	let analyser = $state<AnalyserNode | null>(null);
	let elapsed = $state(0);
	let audioSrc = $state<string | undefined>(undefined);
	let error = $state<string | undefined>(undefined);

	let recorder: MediaRecorder | undefined;
	let stream: MediaStream | undefined;
	let audioCtx: AudioContext | undefined;
	let chunks: Blob[] = [];
	let blob: Blob | null = null;
	let ticker: ReturnType<typeof setInterval> | undefined;
	let startedAt = 0;

	const arm = async () => {
		stage = 'arming';
		try {
			stream = await navigator.mediaDevices.getUserMedia({ audio: true });

			audioCtx = new AudioContext();
			const node = audioCtx.createAnalyser();
			node.fftSize = 1024;
			node.smoothingTimeConstant = 0.75;
			audioCtx.createMediaStreamSource(stream).connect(node);
			analyser = node;

			recorder = new MediaRecorder(stream);
			recorder.ondataavailable = (e) => chunks.push(e.data);
			recorder.onstop = () => {
				blob = new Blob(chunks, { type: 'audio/ogg; codecs=opus' });
				chunks = [];
				audioSrc = URL.createObjectURL(blob);
				stage = 'review';
			};

			stage = 'ready';
		} catch {
			stage = 'blocked';
		}
	};

	onMount(arm);

	const capture = () => {
		if (stage === 'ready') {
			chunks = [];
			startedAt = Date.now();
			elapsed = 0;
			recorder?.start();
			stage = 'recording';
			ticker = setInterval(() => (elapsed = (Date.now() - startedAt) / 1000), 100);
		} else if (stage === 'recording') {
			clearInterval(ticker);
			recorder?.stop(); // `stage` moves to 'review' once the blob is assembled
		}
	};

	const startOver = () => {
		if (audioSrc) URL.revokeObjectURL(audioSrc);
		audioSrc = undefined;
		blob = null;
		elapsed = 0;
		error = undefined;
		stage = 'ready';
	};

	const share = async () => {
		if (!blob) return;
		stage = 'sending';
		error = undefined;
		try {
			const vnString = await new Promise<string>((resolve, reject) => {
				const reader = new FileReader();
				reader.onerror = () => reject(reader.error);
				reader.onloadend = () => resolve(reader.result as string);
				reader.readAsDataURL(blob!);
			});

			const resp = await fetch(`${import.meta.env.VITE_API_URL}/api/v1/vn`, {
				method: 'POST',
				body: JSON.stringify({ vnString })
			});
			if (!resp.ok) throw new Error(String(resp.status));

			const json = await resp.json();
			await goto(`/${json.url_short_form}`);
		} catch {
			error = "That didn't send. Give it another go?";
			stage = 'review';
		}
	};

	onDestroy(() => {
		clearInterval(ticker);
		if (audioSrc) URL.revokeObjectURL(audioSrc);
		stream?.getTracks().forEach((t) => t.stop());
		void audioCtx?.close();
	});
</script>

<svelte:head>
	<title>snakra · say it once</title>
	<meta name="description" content="Record a voice note, share the link. It isn't meant to last." />
</svelte:head>

<section class="stage">
	{#if stage === 'arming'}
		<p class="waiting">waiting for your microphone<span class="ellipsis"></span></p>
	{:else if stage === 'blocked'}
		<h1>No ears.</h1>
		<p class="sub">
			snakra needs your microphone to hear a note. Allow it in your browser's address bar, then try
			again.
		</p>
		<button class="btn ghost" onclick={arm}>Try again</button>
	{:else if stage === 'review' || stage === 'sending'}
		<p class="eyebrow">your note</p>
		{#if audioSrc}
			<div class="review" class:busy={stage === 'sending'}>
				<Player src={audioSrc} />
			</div>
		{/if}
		<div class="actions">
			<button class="btn primary" onclick={share} disabled={stage === 'sending'}>
				{stage === 'sending' ? 'Sending…' : 'Make shareable'}
			</button>
			<button class="btn ghost" onclick={startOver} disabled={stage === 'sending'}>
				Start over
			</button>
		</div>
		{#if error}
			<p class="error" role="alert">{error}</p>
		{/if}
	{:else}
		{#if stage === 'recording'}
			<div class="monitor">
				<LiveWave {analyser} />
				<p class="timer">{formatTime(elapsed)}</p>
			</div>
		{:else}
			<h1>Say it once.</h1>
			<p class="sub">Record a note, pass on the link. It isn't meant to last.</p>
		{/if}

		<button
			class="capture"
			class:live={stage === 'recording'}
			onclick={capture}
			aria-label={stage === 'recording' ? 'Stop recording' : 'Start recording'}
		>
			<span class="inner"></span>
		</button>

		<p class="hint">{stage === 'recording' ? 'tap to stop' : 'tap to record'}</p>
	{/if}
</section>

<style>
	.stage {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 1.5rem;
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

	.hint {
		color: var(--ink-faint);
		font-size: 0.8125rem;
		letter-spacing: 0.06em;
	}

	/* ── the one control ─────────────────────────────────────────────── */

	.capture {
		position: relative;
		display: grid;
		place-items: center;
		width: 84px;
		height: 84px;
		border: 2px solid var(--edge-strong);
		border-radius: 50%;
		transition:
			border-color 0.4s var(--ease),
			box-shadow 0.4s var(--ease);
	}

	.capture:hover {
		border-color: var(--accent-soft);
	}

	.capture.live {
		border-color: var(--accent);
		box-shadow: 0 0 0 7px rgb(255 90 71 / 0.09);
	}

	.capture .inner {
		width: 58px;
		height: 58px;
		border-radius: 50%;
		background: var(--accent);
		box-shadow: 0 0 30px rgb(255 90 71 / 0.45);
		transition:
			width 0.4s var(--ease),
			height 0.4s var(--ease),
			border-radius 0.4s var(--ease);
	}

	.capture:active .inner {
		transform: scale(0.94);
	}

	/* the dot becomes a stop square — one button, two meanings */
	.capture.live .inner {
		width: 28px;
		height: 28px;
		border-radius: 9px;
	}

	.capture.live::after {
		content: '';
		position: absolute;
		inset: -2px;
		border: 2px solid var(--accent);
		border-radius: 50%;
		animation: ripple 2s var(--ease) infinite;
	}

	@keyframes ripple {
		from {
			transform: scale(1);
			opacity: 0.5;
		}
		to {
			transform: scale(1.45);
			opacity: 0;
		}
	}

	/* ── recording ───────────────────────────────────────────────────── */

	.monitor {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 0.75rem;
		width: 100%;
	}

	.timer {
		font-family: var(--font-mono);
		font-size: 1.75rem;
		font-variant-numeric: tabular-nums;
		letter-spacing: -0.02em;
		color: var(--ink);
	}

	/* ── review ──────────────────────────────────────────────────────── */

	.review {
		width: 100%;
		transition: opacity 0.3s var(--ease);
	}

	.review.busy {
		opacity: 0.45;
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
		transition:
			background-color 0.25s var(--ease),
			border-color 0.25s var(--ease),
			color 0.25s var(--ease),
			transform 0.2s var(--ease);
	}

	.btn:active:not(:disabled) {
		transform: translateY(1px);
	}

	.btn:disabled {
		opacity: 0.55;
		cursor: default;
	}

	.btn.primary {
		background: var(--accent);
		color: #1a0b08;
	}

	.btn.primary:hover:not(:disabled) {
		background: var(--accent-soft);
	}

	.btn.ghost {
		border: 1px solid var(--edge-strong);
		color: var(--ink-dim);
	}

	.btn.ghost:hover:not(:disabled) {
		border-color: var(--ink-faint);
		color: var(--ink);
	}

	.error {
		color: var(--accent-soft);
		font-size: 0.875rem;
	}

	/* ── waiting ─────────────────────────────────────────────────────── */

	.waiting {
		color: var(--ink-faint);
		font-size: 0.875rem;
		letter-spacing: 0.06em;
	}

	.ellipsis::after {
		content: '…';
		animation: blink 1.4s steps(1) infinite;
	}

	@keyframes blink {
		0%,
		100% {
			opacity: 1;
		}
		50% {
			opacity: 0.25;
		}
	}
</style>

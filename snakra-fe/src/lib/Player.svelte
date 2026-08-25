<script lang="ts">
	import { formatTime } from '$lib/formatTime';

	interface Props {
		src: string;
	}

	let { src }: Props = $props();

	const BARS = 64;

	let paused = $state(true);
	let currentTime = $state(0);
	/** MediaRecorder blobs report `Infinity` here until they're fully decoded. */
	let mediaDuration = $state(0);
	let decodedDuration = $state(0);
	let peaks = $state<number[]>([]);

	const duration = $derived(
		Number.isFinite(mediaDuration) && mediaDuration > 0 ? mediaDuration : decodedDuration
	);
	const progress = $derived(duration > 0 ? Math.min(1, currentTime / duration) : 0);

	// Decode once to draw the real shape of the note — and to get a duration we
	// can trust when the container doesn't carry one.
	$effect(() => {
		const url = src;
		let cancelled = false;

		(async () => {
			try {
				const bytes = await (await fetch(url)).arrayBuffer();
				const ctx = new AudioContext();
				const decoded = await ctx.decodeAudioData(bytes);
				void ctx.close();
				if (cancelled) return;

				const data = decoded.getChannelData(0);
				const block = Math.floor(data.length / BARS) || 1;
				const next: number[] = [];
				for (let i = 0; i < BARS; i++) {
					let peak = 0;
					for (let j = 0; j < block; j++) {
						const v = Math.abs(data[i * block + j] ?? 0);
						if (v > peak) peak = v;
					}
					next.push(peak);
				}
				const ceiling = Math.max(...next, 0.01);
				peaks = next.map((p) => Math.max(0.07, p / ceiling));
				decodedDuration = decoded.duration;
			} catch {
				// Browser can't decode this codec — still playable, just show a
				// neutral shape instead of a broken one.
				if (!cancelled) {
					peaks = Array.from({ length: BARS }, (_, i) => 0.3 + Math.abs(Math.sin(i * 1.7)) * 0.5);
				}
			}
		})();

		return () => (cancelled = true);
	});

	const seekTo = (ratio: number) => {
		if (duration > 0) currentTime = Math.min(duration, Math.max(0, ratio * duration));
	};

	const onScrub = (e: PointerEvent) => {
		const rect = (e.currentTarget as HTMLElement).getBoundingClientRect();
		seekTo((e.clientX - rect.left) / rect.width);
	};

	const onScrubKey = (e: KeyboardEvent) => {
		const step = e.key === 'ArrowLeft' ? -2 : e.key === 'ArrowRight' ? 2 : 0;
		if (!step) return;
		e.preventDefault();
		currentTime = Math.min(duration, Math.max(0, currentTime + step));
	};
</script>

<div class="player" class:playing={!paused}>
	<audio bind:paused bind:currentTime bind:duration={mediaDuration} {src} preload="metadata">
		<track kind="captions" />
	</audio>

	<button
		class="play"
		type="button"
		aria-label={paused ? 'Play' : 'Pause'}
		onclick={() => (paused = !paused)}
	>
		{#if paused}
			<svg viewBox="0 0 24 24" aria-hidden="true"><path d="M8 5.2v13.6L19 12z" /></svg>
		{:else}
			<svg viewBox="0 0 24 24" aria-hidden="true"
				><rect x="7" y="5" width="3.4" height="14" rx="1.4" /><rect
					x="13.6"
					y="5"
					width="3.4"
					height="14"
					rx="1.4"
				/></svg
			>
		{/if}
	</button>

	<div
		class="wave"
		role="slider"
		tabindex="0"
		aria-label="Seek"
		aria-valuemin={0}
		aria-valuemax={Math.round(duration) || 0}
		aria-valuenow={Math.round(currentTime)}
		aria-valuetext={formatTime(currentTime)}
		onpointerdown={onScrub}
		onkeydown={onScrubKey}
	>
		{#each peaks as peak, i (i)}
			<span class="bar" class:played={i / peaks.length < progress} style="--h:{peak}"></span>
		{/each}
	</div>

	<span class="time">{formatTime(currentTime)}<span class="sep">/</span>{formatTime(duration)}</span>
</div>

<style>
	.player {
		display: flex;
		align-items: center;
		gap: 0.9rem;
		width: 100%;
		padding: 0.7rem 0.9rem 0.7rem 0.7rem;
		background: var(--bg-raised);
		border: 1px solid var(--edge);
		border-radius: var(--radius);
		transition: border-color 0.4s var(--ease);
	}

	.player.playing {
		border-color: var(--edge-strong);
	}

	.play {
		flex: none;
		display: grid;
		place-items: center;
		width: 42px;
		height: 42px;
		border-radius: 50%;
		background: var(--accent);
		color: var(--bg);
		transition:
			transform 0.25s var(--ease),
			box-shadow 0.35s var(--ease);
	}

	.play:hover {
		transform: scale(1.06);
	}

	.play:active {
		transform: scale(0.96);
	}

	.playing .play {
		box-shadow: 0 0 0 6px rgb(255 90 71 / 0.13);
	}

	.play svg {
		width: 19px;
		height: 19px;
		fill: currentColor;
	}

	.wave {
		flex: 1;
		display: flex;
		align-items: center;
		gap: 2px;
		height: 40px;
		min-width: 0;
		cursor: pointer;
		touch-action: none;
	}

	.bar {
		flex: 1;
		min-width: 2px;
		height: max(3px, calc(var(--h) * 100%));
		border-radius: 2px;
		background: var(--ink-faint);
		transition:
			background-color 0.18s linear,
			height 0.3s var(--ease);
	}

	.bar.played {
		background: var(--accent);
	}

	.time {
		flex: none;
		font-family: var(--font-mono);
		font-size: 0.75rem;
		font-variant-numeric: tabular-nums;
		color: var(--ink-dim);
	}

	.sep {
		opacity: 0.4;
		margin: 0 0.35em;
	}

	@media (max-width: 380px) {
		.time {
			display: none;
		}
	}
</style>

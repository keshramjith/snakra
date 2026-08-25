<script lang="ts">
	/**
	 * Scrolling amplitude history of the live mic input — newest sample on the
	 * right, older ones fading out to the left. Drawn on canvas so a 60fps
	 * redraw never touches the reactivity graph.
	 */
	interface Props {
		analyser: AnalyserNode | null;
		bars?: number;
	}

	let { analyser, bars = 56 }: Props = $props();

	let canvas: HTMLCanvasElement | undefined = $state();

	$effect(() => {
		const node = analyser;
		const el = canvas;
		if (!node || !el) return;

		const ctx = el.getContext('2d');
		if (!ctx) return;

		const samples = new Uint8Array(node.fftSize);
		const history = new Float32Array(bars);
		let frame = 0;

		const resize = () => {
			const dpr = Math.min(window.devicePixelRatio || 1, 2);
			const rect = el.getBoundingClientRect();
			el.width = Math.max(1, Math.round(rect.width * dpr));
			el.height = Math.max(1, Math.round(rect.height * dpr));
			ctx.setTransform(dpr, 0, 0, dpr, 0, 0);
			return rect;
		};

		let box = resize();
		const observer = new ResizeObserver(() => (box = resize()));
		observer.observe(el);

		const draw = () => {
			frame = requestAnimationFrame(draw);

			node.getByteTimeDomainData(samples);
			let sum = 0;
			for (let i = 0; i < samples.length; i++) {
				const v = (samples[i] - 128) / 128;
				sum += v * v;
			}
			const rms = Math.sqrt(sum / samples.length);

			history.copyWithin(0, 1);
			history[bars - 1] = Math.min(1, rms * 3.4);

			const { width, height } = box;
			ctx.clearRect(0, 0, width, height);

			const gap = 3;
			const barWidth = Math.max(1.5, (width - gap * (bars - 1)) / bars);
			const mid = height / 2;

			for (let i = 0; i < bars; i++) {
				const age = i / (bars - 1);
				const h = Math.max(barWidth * 0.85, history[i] * height * 0.92);
				ctx.fillStyle = `rgba(255, 122, 104, ${0.1 + age * 0.9})`;
				ctx.beginPath();
				ctx.roundRect(i * (barWidth + gap), mid - h / 2, barWidth, h, barWidth / 2);
				ctx.fill();
			}
		};

		frame = requestAnimationFrame(draw);

		return () => {
			cancelAnimationFrame(frame);
			observer.disconnect();
		};
	});
</script>

<canvas bind:this={canvas} aria-hidden="true"></canvas>

<style>
	canvas {
		display: block;
		width: 100%;
		height: 84px;
	}
</style>

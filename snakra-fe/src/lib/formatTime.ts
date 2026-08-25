/** `0:07`, and never `NaN:aN` for the Infinity durations MediaRecorder blobs report. */
export const formatTime = (seconds: number) => {
	if (!Number.isFinite(seconds) || seconds < 0) return '0:00';
	const total = Math.floor(seconds);
	return `${Math.floor(total / 60)}:${String(total % 60).padStart(2, '0')}`;
};

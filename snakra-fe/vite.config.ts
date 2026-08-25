import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import fs from 'node:fs';

export default defineConfig({
	plugins: [sveltekit()],
	server:
		process.env.NODE_ENV === 'dev'
			? {
					https: {
						key: fs.readFileSync(`${import.meta.dirname}/tls/key.pem`),
						cert: fs.readFileSync(`${import.meta.dirname}/tls/cert.pem`)
					},
					proxy: {}
				}
			: {}
});

import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import fs from 'fs';

export default defineConfig({
	plugins: [sveltekit()],
	server: function() {
		if (process.env.NODE_ENV === 'dev') {
			return {
				https: {
					key: fs.readFileSync(`${__dirname}/tls/key.pem`),
					cert: fs.readFileSync(`${__dirname}/tls/cert.pem`),
				},
				proxy: {}
			}
		} else {
			return {}
		}
	}()
});

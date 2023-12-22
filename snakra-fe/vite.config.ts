import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import fs from 'fs'

console.log('dirname: ', `./../snakra-be/tls/key.pem`)

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  server: {
    https: {
      key: fs.readFileSync(`./../tls/key.pem`),
      cert: fs.readFileSync(`./../tls/cert.pem`),
    },
    proxy: {}
  }
})

import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import fs from 'fs'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  appType: "spa",
  server: function() {
    if (process.env.NODE_ENV === 'Development') {
      return {
        https: {
          key: fs.readFileSync(`./../tls/key.pem`),
          cert: fs.readFileSync(`./../tls/cert.pem`),
        },
        proxy: {}
      }
    }
  }()
})

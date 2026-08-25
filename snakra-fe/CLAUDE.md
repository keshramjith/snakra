# snakra-fe

SvelteKit frontend for snakra.com — a web app for creating disposable voice notes.
Record a note, get a shareable short link, play it back at `/<url_short_form>`.

The Go backend lives in `../snakra-be` and is reached via `VITE_API_URL`
(`POST /api/v1/vn` to upload, `GET /api/v1/vn/:id` to fetch).

## Tooling

Use Bun as the package manager and script runner — the lockfile is `bun.lock`.

- `bun install` instead of npm/yarn/pnpm install
- `bun run <script>` instead of `npm run <script>`
- `bunx <package>` instead of `npx`
- Bun automatically loads `.env`, so don't add `dotenv`

Bun is *not* the bundler or dev server here. This is a Vite + SvelteKit app:
don't reach for `Bun.serve()` or HTML imports, and don't replace Vite.

## Scripts

- `bun run dev` — Vite dev server
- `bun run build` — production build (`@sveltejs/adapter-auto`)
- `bun run preview` — serve the production build
- `bun run check` — `svelte-kit sync` + `svelte-check` (must stay at 0 errors)

For local HTTPS, set `NODE_ENV=dev` and provide `tls/key.pem` + `tls/cert.pem`;
see `vite.config.ts`.

## Conventions

- Svelte 5 with runes: `$state`, `$derived`, `$props`. Use `onclick`, not `on:click`.
- Import page/navigation state from `$app/state`, not the deprecated `$app/stores`.
- Routes live in `src/routes`; shared code goes in `src/lib` (`$lib` alias).
- Both routes opt out of SSR (`export const ssr = false`) because recording and
  playback depend on browser-only APIs (`MediaRecorder`, `Blob`, `URL.createObjectURL`).

import './app.css'
import App from './App.svelte'
// import { createAuth0Client } from '@auth0/auth0-spa-js'

// const auth0 = await createAuth0Client({
//   domain: '1234',
//   clientId: "1234",
//   authorizationParams: {
//     audience: "12344"
//   }
// })

const app = new App({
  target: document.getElementById('app'),
})

export default app

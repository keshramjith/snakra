<script lang="ts">
import { Router, Route } from "svelte-navigator"
import Index from "./pages/Index.svelte";
  import Player from "./pages/Player.svelte";
</script>

<Router>
  <main>
    <Route path="/">
      <Index />
    </Route>
    <Route path="/:id" let:params>
      <Player id="{params.id}"/>
    </Route>
  </main>
</Router>

<script lang="ts">
import { Router, Route } from "svelte-navigator"
import Index from "./pages/Index.svelte";
  import Player from "./pages/Player.svelte";
  import NotFound from "./pages/NotFound.svelte";
</script>

<Router>
  <main>
    <Route path="*">
      <Route path="/">
        <Index />
      </Route>
      <Route path="/:id" let:params>
        <Player id="{params.id}"/>
      </Route>
      <Route path="/notfound">
        <NotFound />
      </Route>
    </Route>
  </main>
</Router>

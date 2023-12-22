<script lang="ts">
import { Router, Route } from "svelte-navigator"
import Index from "./pages/Index.svelte";
  import Player from "./pages/Player.svelte";
  import NotFound from "./pages/NotFound.svelte";
</script>

<Router>
  <main>
    <Route path="/">
      <Index />
    </Route>
    <Route path="/:id/*" let:params>
      <Router>
        <Route path="/">
          <Player id="{params.id}"/>
        </Route>
      </Router>
    </Route>
    <Route path="/notfound">
      <NotFound />
    </Route>
  </main>
</Router>

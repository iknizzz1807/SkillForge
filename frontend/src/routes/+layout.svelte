<script lang="ts">
  import "../app.css";
  let { children } = $props();
  import Nav from "../components/Nav.svelte";
</script>

<Nav />
{@render children()}

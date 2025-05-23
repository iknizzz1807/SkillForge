import tailwindcss from "@tailwindcss/vite";
import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vite";

export default defineConfig({
  plugins: [tailwindcss(), sveltekit()],
  server: {
    // ...existing code...
    allowedHosts: [
      // ...existing allowed hosts...
      "skillforge.ikniz.id.vn",
    ],
    hmr: false,
    // Additional server settings
    watch: {
      usePolling: false,
    },
  },
  ssr: {
    noExternal: ["sortablejs", "socket.io-client"],
  },
});

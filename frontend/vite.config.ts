import tailwindcss from "@tailwindcss/vite";
import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vite";

export default defineConfig({
  plugins: [tailwindcss(), sveltekit()],
  server: {
    allowedHosts: [
      "skillforge.ikniz.id.vn",
    ],
    hmr: false,
    watch: {
      usePolling: false,
    },
    proxy: {
      "/api": "http://backend:8080",
      "/avatars": "http://backend:8080",
      "/portfolios": "http://backend:8080",
      "/storage": "http://backend:8080",
      "/ws": {
        target: "ws://backend:8080",
        ws: true,
      },
    },
  },
  ssr: {
    noExternal: ["sortablejs", "socket.io-client"],
  },
});

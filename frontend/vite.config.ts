import tailwindcss from "@tailwindcss/vite";
import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vitest/config";

export default defineConfig({
  plugins: [tailwindcss(), sveltekit()],
  resolve: {
    conditions: ["browser"],
  },
  test: {
    environment: "jsdom",
    include: ["src/**/*.{test,spec}.{js,ts}"],
    setupFiles: ["src/lib/__tests__/setup.ts"],
    server: {
      deps: {
        inline: ["svelte"],
      },
    },
  },
  server: {
    allowedHosts: [
      "skillforge.ikniz.id.vn",
    ],
    hmr: false,
    watch: {
      usePolling: false,
    },
    proxy: {
      "/avatars": process.env.BACKEND_URL || "http://backend:8080",
      "/portfolios": process.env.BACKEND_URL || "http://backend:8080",
      "/storage": process.env.BACKEND_URL || "http://backend:8080",
      "/ws": {
        target: (process.env.BACKEND_URL || "http://backend:8080").replace(/^http/, "ws"),
        ws: true,
      },
    },
  },
  ssr: {
    noExternal: ["sortablejs", "socket.io-client"],
  },
});

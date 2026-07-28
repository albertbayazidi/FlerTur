import react from "@vitejs/plugin-react-swc";
import { defineConfig } from "vite";

export default defineConfig({
  plugins: [react()],
  server: {
    host: true,
    port: 1234,
    proxy: {
      "/api": {
        target: "http://127.0.0.1:3005",
        changeOrigin: true,
      },
    },
  },
});

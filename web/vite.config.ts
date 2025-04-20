import { fileURLToPath, URL } from "node:url";

import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import vueDevTools from "vite-plugin-vue-devtools";

// https://vite.dev/config/
export default defineConfig({
    plugins: [vue(), vueDevTools()],
    resolve: {
        alias: {
            "@": fileURLToPath(new URL("./src", import.meta.url)),
        },
    },
    server: {
        port: 5174,
        proxy: {
            '/api': {
                target: 'http://127.0.0.1:8083',
                changeOrigin: true,
                rewrite: (path) => path.replace(/^\/api/, '')
            },
            '/oauth2' :{
                target: 'http://127.0.0.1:8084',
                changeOrigin: true,
                rewrite: (path) => path.replace(/^\/oauth2/, '')
            }
        },
    },
});

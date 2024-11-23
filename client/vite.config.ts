import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react-swc';
import fs from 'node:fs';
import path from 'node:path';

// https://vite.dev/config/
export default defineConfig({
    plugins: [react()],
    server: {
        https: {
            key: fs.readFileSync(path.resolve(__dirname, 'cert', 'localhost-key.pem')),
            cert: fs.readFileSync(path.resolve(__dirname, 'cert', 'localhost.pem')),
        },
    },
});

import { defineConfig } from 'vite';

export default defineConfig({
    build:{
        minify:'terser',
        //minify:false,
        terserOptions:{
            keep_classnames:true,
            keep_fnames:true,
            mangle:false,
        },
        modulePreload:{
            polyfill:false,
        },
        sourcemap: 'inline',
    },
})
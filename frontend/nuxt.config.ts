// https://nuxt.com/docs/api/configuration/nuxt-config
import tailwindcss from "@tailwindcss/vite";

export default defineNuxtConfig({
  $development: {
    logLevel: 'verbose',
  },
  $production: {
    logLevel: 'silent',
  },
  runtimeConfig: {
    public: {
      apiBase: 'https://api.example.com' //TODO
    }
  },
  app: {
    head: {
      htmlAttrs: { 
        'data-theme': 'dracula'
      }
    },
  },
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  css:['~/assets/css/main.css'],
  vite: {
    plugins: [
      tailwindcss()
    ],
    server: {
      allowedHosts: true  
    },    
  },
  modules: [
    '@nuxt/eslint',
    '@nuxt/fonts',
    '@nuxt/icon',
    '@nuxt/image',
    '@nuxt/ui',
    '@pinia/nuxt',
  ]
})
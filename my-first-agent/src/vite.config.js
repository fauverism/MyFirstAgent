// vite.config.js
export default {
  server: {
    proxy: {
      '/chat': 'http://localhost:8080'
    }
  }
}
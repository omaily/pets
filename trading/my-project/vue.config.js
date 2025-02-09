const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  devServer: {
    port: 3000, // Устанавливаем нужный порт
  },
  transpileDependencies: true
})

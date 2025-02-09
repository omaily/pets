cd ./web
vue --clear-cache                   #1Очистка кэша Vue CLI
npm-check-updates                   #2обновить все пакеты до последних версий
npm install -g npm-check-updates    #3обновить все пакеты до последних версий
ncu -u                              #3обновить все пакеты до последних версий
npx @vue/cli create my-project --preset vue-next #4устанавливает последние версии зависимостей, независимо от того, что установлено глобально.
devServer: {
    port: 3000, // Устанавливаем нужный порт
},
# последние версии зависимостей на 10.02.2025
#   "dependencies": {
#     "core-js": "^3.30.0",
#     "vue": "^3.3.0"
#   },
#   "devDependencies": {
#     "@babel/core": "^7.21.0",
#     "@babel/eslint-parser": "^7.21.0",
#     "@vue/cli-plugin-babel": "^5.0.0",
#     "@vue/cli-plugin-eslint": "^5.0.0",
#     "@vue/cli-service": "^5.0.0",
#     "eslint": "^8.0.0",
#     "eslint-plugin-vue": "^9.0.0"
#   },
npm install lightweight-charts
npm run serve
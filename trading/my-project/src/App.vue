<template>
  <div id="app">
    <StockChart />
  </div>
</template>

<script>
import StockChart from './components/StockChart.vue'

export default {
  name: 'App',
  components: {
    StockChart
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>

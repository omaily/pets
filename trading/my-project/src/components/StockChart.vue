<template>
  <div ref="chartContainer" style="width: 100%; height: 500px;"></div>
</template>

<script>
import { createChart, CandlestickSeries } from 'lightweight-charts'; // Импортируем CandlestickSeries
import axios from 'axios';

export default {
  data() {
    return {
      chart: null,
      candleSeries: null
    };
  },
  mounted() {
    this.initChart();
    this.fetchCandles();
  },
  methods: {
    initChart() {
      // Инициализируем график
      this.chart = createChart(this.$refs.chartContainer, {
        width: this.$refs.chartContainer.clientWidth,
        height: this.$refs.chartContainer.clientHeight,
        layout: {
          backgroundColor: '#fff',
        },
        grid: {
          vertLines: { color: '#eee' },
          horzLines: { color: '#eee' },
        },
        crosshair: {
          mode: 0,
        },
        priceScale: {
          borderColor: '#ccc',
        },
        timeScale: {
          borderColor: '#ccc',
        },
      });

      // Используем addSeries для добавления CandlestickSeries
      this.candleSeries = this.chart.addSeries(CandlestickSeries, {}); // Передаем класс CandlestickSeries
    },
    async fetchCandles() {
      try {
        const response = await axios.get('http://localhost:8080/candles');
        const formattedData = response.data.map((candle) => ({
          time: candle.time, // unix timestamp
          open: candle.open,
          high: candle.high,
          low: candle.low,
          close: candle.close,
        }));

        // Отображаем данные на графике
        if (this.candleSeries) {
          this.candleSeries.setData(formattedData);
        }
      } catch (error) {
        console.error('Error fetching candles:', error);
      }
    }
  },
  beforeUnmount() {
    if (this.chart) {
      this.chart.remove();
    }
  }
};
</script>

<style scoped>
/* Стили для графика */
</style>
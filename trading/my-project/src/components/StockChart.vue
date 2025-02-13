<template>
  <div>
    <div ref="mainChart" style="width: 100%; height: 350px;"></div>
    <div ref="impulseChart" style="width: 100%; height: 150px; margin-top: 5px;"></div>
  </div>
</template>

<script>
import { createChart, CandlestickSeries, HistogramSeries } from 'lightweight-charts';
import axios from 'axios';

export default {
  data() {
    return {
      mainChart: null,
      impulseChart: null,
      candleSeries: null,
      impulseSeries: null,
    };
  },
  mounted() {
    this.initCharts();
    this.fetchCandles();
  },
  methods: {
    async fetchCandles() {
      try {
        const response = await axios.get('http://localhost:8080/candles');

        // Формируем данные для свечного графика
        const candleData = response.data.map(c => ({
          time: c.time,
          open: c.open,
          high: c.high,
          low: c.low,
          close: c.close,
        }));

        // Формируем данные для импульса
        const impulseData = response.data.map(c => ({
          time: c.time,
          value: c.impulse ?? 0,
          color: this.getImpulseColor(c.impulse),
        }));

        // Устанавливаем данные
        this.candleSeries.setData(candleData);
        this.impulseSeries.setData(impulseData);

        // Синхронизация временных шкал
        const mainTimeScale = this.mainChart.timeScale();
        const impulseTimeScale = this.impulseChart.timeScale();
        mainTimeScale.subscribeVisibleTimeRangeChange((range) => {
          impulseTimeScale.setVisibleRange(range);
        });

        mainTimeScale.fitContent();
      } catch (error) {
        console.error('Ошибка загрузки данных:', error);
      }
    },

    initCharts() {
      // Инициализация основного графика
      this.mainChart = createChart(this.$refs.mainChart, {
        layout: { backgroundColor: '#fff', textColor: '#000' },
        grid: { vertLines: { color: '#eee' }, horzLines: { color: '#eee' } },
        timeScale: { borderColor: '#ccc' },
      });
      this.candleSeries = this.mainChart.addSeries(CandlestickSeries, {
        upColor: '#26a69a',
        downColor: '#ef5350',
        borderVisible: false,
        wickUpColor: '#26a69a',
        wickDownColor: '#ef5350',
      });

      // Инициализация графика импульсов
      this.impulseChart = createChart(this.$refs.impulseChart, {
        layout: { backgroundColor: '#fff', textColor: '#000' },
        grid: { vertLines: { color: '#eee' }, horzLines: { color: '#eee' } },
        timeScale: { borderColor: '#ccc' },
      });
      this.impulseSeries = this.impulseChart.addSeries(HistogramSeries, {
        color: 'blue',
        lineWidth: 2,
      });
    },

    getImpulseColor(impulse) {
      if (impulse === 2) return '#15ff00';
      if (impulse === 1) return '#005A00';
      if (impulse === 0) return '#FFFFFF';
      if (impulse === -1) return '#5A0000';
      if (impulse === -2) return '#FF0000';
      return '#000000';
    }
  },

  beforeUnmount() {
    if (this.mainChart) this.mainChart.remove();
    if (this.impulseChart) this.impulseChart.remove();
  }
};
</script>
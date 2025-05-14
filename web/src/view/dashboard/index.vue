<template>
  <div
    class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-7 py-2 gap-4 md:gap-2 gva-container2"
  >
  <gva-card custom-class="col-span-1 lg:col-span-2 ">
      <gva-chart :type="1" title="访问人数" />
    </gva-card>
  </div>
</template>

<script setup>
  import {

    GvaCard,
  } from './components'
  defineOptions({
    name: 'Dashboard'
  })
</script>

<style lang="scss" scoped></style>

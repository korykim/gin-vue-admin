<template>
    <div class="flex justify-center items-center min-h-screen">
        <div class="text-center">
            <h1 class="text-3xl font-bold mb-4">Eye</h1>
            <!-- <p>当前路由: {{ route.path }}</p>
            <p>当前路由: {{ route.query }}</p> -->
            <p>当前路由: {{ route.params.id }}</p>
            <!-- <p>当前路由: {{ route.hash }}</p>
            <p>当前路由: {{ route.fullPath }}</p>
            <p>当前路由: {{ route.name }}</p>
            <p>当前路由: {{ route.meta }}</p>
            <p>当前路由: {{ route.redirectedFrom }}</p> -->
        </div>
    </div>

</template>

<script setup>

import { useRoute } from 'vue-router';
const route = useRoute();



</script>
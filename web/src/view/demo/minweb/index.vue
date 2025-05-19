<template>
    <div class="min-web-container">
        <div class="bg-gray-50 text-slate-700 dark:text-slate-500 dark:bg-slate-800 w-full h-screen">
            <iframe id="custom-iframe" class="w-full h-screen border-t border-gray-200 dark:border-slate-700"
                src="https://www.wikipedia.org"></iframe>
        </div>
    </div>
</template>

<script setup>
defineOptions({
    name: 'MinWeb'
})
</script>

<style scoped>
.min-web-container {
    width: 100%;
    height: 100vh;
}
</style>

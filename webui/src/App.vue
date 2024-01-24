<script setup>
	import { RouterView } from 'vue-router'
</script>

<script>
	export default {}
</script>

<template>
	<main>
		<RouterView :key="$route.fullPath"/>
	</main>
</template>

<style>
</style>

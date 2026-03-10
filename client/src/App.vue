<script setup lang="ts">
import { RouterView } from 'vue-router';
import Header from './components/Header.vue';
import Footer from './components/Footer.vue';
import { useAuthStore } from './stores/auth';
import { onMounted } from 'vue';

const auth = useAuthStore();
onMounted(() => {
	auth.fetchMe();
});
</script>

<template>
	<Header />
	<RouterView />
</template>
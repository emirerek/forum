<template>
	<div class="subforum-list">
		<Subforum v-for="subforum in subforumList" :key="subforum.id" :subforum="subforum" />
	</div>
</template>

<script setup lang="ts">
import Subforum from "./Subforum.vue";
import type { SubforumModel } from "@/models/subforum";

const props = defineProps<{ subforumList: SubforumModel[] }>();
</script>

<style scoped>
.subforum-list {
	width: 100%;
	display: flex;
	flex-direction: column;
	justify-content: center;
	align-items: flex-start;
	gap: 1rem;
	list-style: none;
	padding: 0;
}
</style>

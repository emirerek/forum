<template>
	<div class="category">
		<h2>{{ category.name }}</h2>
		<SubforumList :subforumList="category.subforums" />
	</div>
</template>

<script setup lang="ts">
import SubforumList from "./SubforumList.vue";
import type { CategoryModel } from "@/models/category";

const props = defineProps<{ category: CategoryModel }>();
</script>

<style scoped>
.category {
	margin-bottom: 2rem;
}
</style>

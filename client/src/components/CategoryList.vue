<template>
	<div class="category-list">
		<Category v-for="category in categoryList" :key="category.id" :category="category" />
	</div>
</template>

<script setup lang="ts">
import Category from "./Category.vue";
import type { CategoryModel } from "@/models/category";

const props = defineProps<{ categoryList: CategoryModel[] }>();
</script>

<style scoped>
.category-list {
	width: 100%;
	display: flex;
	flex-direction: column;
	gap: 2rem;
}
</style>

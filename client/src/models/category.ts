import type { SubforumModel } from "./subforum";

export interface CategoryModel {
	id: number;
	createdAt: string;
	updatedAt: string;
	name: string;
	subforums: SubforumModel[];
}

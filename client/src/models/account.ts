import type { ReplyModel } from "./reply";
import type { ThreadModel } from "./thread";

export interface AccountModel {
	id: number;
	username: string;
	email: string;
	createdAt: string;
	updatedAt: string;
	threads: ThreadModel[] | null;
	replies: ReplyModel[] | null;
}

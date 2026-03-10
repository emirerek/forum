import type { AccountModel } from "./account";
import type { ThreadModel } from "./thread";

export interface ReplyModel {
	id: number;
	accountId: number;
	threadId: number;
	content: string;
	createdAt: string;
	updatedAt: string;
	account: AccountModel;
	thread: ThreadModel | null;
}

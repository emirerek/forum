import type { AccountModel } from "./account";
import type { SubforumModel } from "./subforum";
import type { ReplyModel } from "./reply";

export interface ThreadModel {
	id: number;
	accountId: number;
	subforumId: number;
	title: string;
	content: string;
	createdAt: string;
	updatedAt: string;
	account: AccountModel;
	subforum: SubforumModel[] | null;
	replies: ReplyModel[] | null;
}

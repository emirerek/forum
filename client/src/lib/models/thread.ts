import { AccountModel } from "./account";
import { SubforumModel } from "./subforum";
import { ReplyModel } from "./reply";

export interface ThreadModel {
    id: number
    accountId: number
    subforumId: number
    title: string
    content: string
    createdAt: string
    updatedAt: string
    account: AccountModel
    subforum: SubforumModel[] | null
    replies: ReplyModel[] | null
}
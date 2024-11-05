import { ThreadModel } from "./thread";

export interface SubforumModel {
    id: number
    title: string
    description: string
    createdAt: string
    updatedAt: string
    threads: ThreadModel[] | null
}
import type { PageLoad } from './$types.js'
import type { SubforumModel } from '$lib/models/subforum.js';
import type { ThreadModel } from '$lib/models/thread';

export const load: PageLoad = async ({ fetch }) => {
    const subforumList: SubforumModel[] = await (await fetch("/api/subforum/")).json();
    const threadPreviewList: ThreadModel[] = await (await fetch("/api/thread/")).json();
    return {
        subforumList,
        threadPreviewList
    };
}
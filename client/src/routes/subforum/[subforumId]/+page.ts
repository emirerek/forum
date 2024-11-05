import type { PageLoad } from './$types.js'
import type { ThreadModel } from '$lib/models/thread';

export const load: PageLoad = async ({ fetch, params }) => {
    const url = "/api/thread/?subforumId=" + params.subforumId;
    const response = await fetch(url);
    const threadPreviewList: ThreadModel[] = await response.json();
    
    return {
        threadPreviewList
    };
}
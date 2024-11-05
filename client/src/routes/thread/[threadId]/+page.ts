import type { PageLoad } from './$types.js'
import type { ThreadModel } from '$lib/models/thread';

export const load: PageLoad = async ({ fetch, params }) => {
    const url = "/api/thread/" + params.threadId;
    const response = await fetch(url);
    const threadPreview: ThreadModel = await response.json();
    
    return {
        threadPreview
    };
}
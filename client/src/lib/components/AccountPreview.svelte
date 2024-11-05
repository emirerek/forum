<script lang="ts">
    import type { AccountModel } from "$lib/models/account";

    const { accountPreview }: { accountPreview: AccountModel } = $props();
    const accountURL = $derived("/account/" + accountPreview.id);
</script>

<div>
    <p>
        <a href={accountURL}>
            {accountPreview.username}
        </a>
    </p>
    <p>{accountPreview.createdAt}</p>
</div>
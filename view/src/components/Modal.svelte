<script>
    import { closeModal } from "svelte-modals";
    export let title;
    export let isOpen;
    export let redirect;
    export let goly;
    export let random;
    export let send;

    let data = {
        redirect: redirect,
        goly: goly,
        random: random,
    };
</script>

{#if isOpen}
    <div role="dialog" class="modal">
        <div class="contents">
            <h2>{title}</h2>
            <label for="redirect">Redirect To:</label>
            <input
                type="text"
                name="redirect"
                id="redirect"
                bind:value={redirect}
            />

            <label for="goly">Goly</label>
            <input
                type="text"
                name="goly"
                id="goly"
                bind:value={data.goly}
                disabled={data.random}
                class={data.random ? "disabled" : ""}
            />

            <label for="isRandom">Randomly Generated Goly</label>
            <input
                type="checkbox"
                name="isRandom"
                id="isRandom"
                bind:checked={data.random}
            />
            <div class="actions">
                <button on:click={closeModal}>Cancel</button>
                <button on:click={send(data)}>Send</button>
            </div>
        </div>
    </div>
{/if}

<style>
    .modal {
        position: fixed;
        top: 0;
        bottom: 0;
        right: 0;
        left: 0;
        display: flex;
        justify-content: center;
        align-items: center;

        pointer-events: none;
    }

    .contents {
        min-width: 500px;
        border-radius: 0.5rem;
        padding: 1rem;
        background: white;
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        pointer-events: auto;
    }

    h2 {
        text-align: center;
        font-size: 1.5rem;
    }

    .actions {
        margin-top: 2rem;
        display: flex;
        justify-content: space-between;
        gap: 0.5rem;
    }

    .disabled {
        background: slategray;
    }
</style>

<script>
    import Modal from "./Modal.svelte";
    import Card from "./Card.svelte";
    import { Modals, closeModal, openModal } from "svelte-modals";

    export let goly;
    let isShow = true;

    async function update(data) {
        const json = {
            redirect: data.redirect,
            goly: data.goly,
            random: data.random,
            id: goly.id,
        };

        await fetch("http://localhost:3000/goly", {
            method: "PATCH",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(json),
        }).then((res) => {
            console.log(res);
        });
    }

    function handleOpen(goly) {
        openModal(Modal, {
            title: "Update Goly Link",
            send: update,
            goly: goly.goly,
            redirect: goly.redirect,
            random: goly.random,
        });
    }

    async function deleteGoly() {
        if (
            confirm(
                "Are you sure you wish to delete this Goly link (" +
                    goly.goly +
                    ") ?"
            )
        ) {
            await fetch("http://localhost:3000/goly/" + goly.id, {
                method: "DELETE",
            }).then((response) => {
                ifShow = false;
                console.log(response);
            });
        }
    }
</script>

{#if isShow}
    <Card>
        <p>Goly: http://localhost:3000/r/{goly.goly}</p>
        <p>Redirect: {goly.redirect}</p>
        <p>Clicled: {goly.clicked}</p>
        <button class="update" on:click={handleOpen(goly)}>Update</button>
        <button class="delete" on:click={deleteGoly}>Delete</button>
    </Card>
{/if}
<Modals>
    <div slot="backdrop" class="backdrop" on:click={closeModal} />
</Modals>

<style>
    button {
        color: white;
        font-weight: bolder;
        border: none;
        padding: 0.75rem;
        border-radius: 0.25rem;
    }

    .update {
        background-color: yellowgreen;
    }

    .delete {
        color: red;
    }

    button:hover {
        cursor: pointer;
    }

    .backdrop {
        position: fixed;
        top: 0;
        bottom: 0;
        right: 0;
        left: 0;
        background: yellowgreen;
    }
</style>

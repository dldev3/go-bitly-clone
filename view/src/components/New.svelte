<script>
    import { Modals, closeModal, openModal } from "svelte-modals";
    import Modal from "./Modal.svelte";

    async function createGoly(data) {
        const json = {
            redirect: data.redirect,
            goly: data.goly,
            random: data.random,
        };

        await fetch("http://localhost:3000/goly", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(json),
        }).then((res) => {
            console.log(res);
        });
    }

    function handleOpen() {
        openModal(Modal, {
            title: "Create New Goly Link",
            send: createGoly,
            redirect: "",
            goly: "",
            random: false,
        });
    }
</script>

<button on:click={handleOpen}>Create a New Goly Link</button>

<style>
    button {
        background-color: green;
        color: white;
        font-weight: bold;
        border: none;
        padding: 0.5rem;
        border-radius: 0.25rem;
    }

    button:hover {
        cursor: pointer;
    }
</style>

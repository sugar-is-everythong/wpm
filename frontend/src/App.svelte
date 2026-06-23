<script>
    import {routes} from "./routes";
    import {Router} from "@mateothegreat/svelte5-router";

    const ocm = (e) => {
        e.preventDefault()
    }
</script>

<!--<svelte:body oncontextmenu={ocm}/>-->

<main class="WINDOW-MAIN">
    <Router {routes}/>
</main>

<style>
    ::-webkit-scrollbar {
        display: none;
    }

    .WINDOW-MAIN {
        ::-webkit-scrollbar {
            display: none;
        }
    }
</style>
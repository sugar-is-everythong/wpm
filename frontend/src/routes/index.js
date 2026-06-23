import Main from "@/views/main/index.svelte";
import Tip from "@/views/tip/index.svelte";
import Config from "@/views/config/index.svelte";

export const routes= [
    {
        component: Main,
    },
    {
        path: "/tip",
        component: Tip,
    },
    {
        path: "/config",
        component: Config,
    },
];
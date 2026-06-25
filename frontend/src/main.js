import { mount } from 'svelte'
import App from './App.svelte'
import { loadConfig } from './lib/config.svelte.js'

loadConfig().then(() => {
    mount(App, { target: document.getElementById("app") })
})

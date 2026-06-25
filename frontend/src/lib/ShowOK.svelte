<script>
    import {fly} from 'svelte/transition'

    let {ok, cancel, show} = $props()

    const okSet = () => {
        ok?.()
    }

    const cancelfunc = () => {
        cancel?.()
    }
</script>

{#if show}
    <div class="show-form-b" transition:fly={{y: 40, duration: 300}}>
        <button onclick={cancelfunc}>取消</button>
        <button class="ok" onclick={okSet}>确定</button>
    </div>
{/if}

<style lang="scss">
    .show-form-b {
        position: fixed;
        bottom: 30px;
        left: 50%;
        transform: translate(-50%);
        border-radius: 9999px;
        background-color: rgb(255 255 255 / 0.85);
        backdrop-filter: blur(12px);
        padding: 6px 8px;
        box-shadow: 0 4px 20px rgba(0, 0, 0, 0.12), 0 1px 4px rgba(0, 0, 0, 0.08);
        display: flex;
        gap: 4px;

        button {
            padding: 6px 22px;
            border-radius: 9999px;
            font-size: 13px;
            border: none;
            cursor: pointer;
            transition: background-color .15s ease, opacity .15s ease;
            white-space: nowrap;

            &:hover {
                opacity: 0.85;
            }

            &:active {
                opacity: 0.7;
            }
        }

        .ok {
            background-color: var(--main-3);
            color: #fff;
        }

        button:not(.ok) {
            background-color: transparent;
            color: #333;

            &:hover {
                background-color: rgba(0, 0, 0, 0.06);
                opacity: 1;
            }
        }
    }
</style>
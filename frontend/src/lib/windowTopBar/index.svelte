<script>
    import {button, close, hidden, icon, minSize, search, ToggleFullscreen} from "./ys.svelte";

    const {
        configs = {}, ll, lr, rl, rr, cl, cr,
    } = $props();

    let a = {
        left: [],// 代表左侧的配置，组件展示按照数组顺序展示
        center: [ // 代表中间的配置，组件展示按照数组顺序展示
            {
                name: "icon",//组件名称，同一个组件可以出现在多个位置,就是说可以次插入，每次插入相互独立
                mouseup: () => {
                    console.log("点击了icon2")
                },
                style: {},// 组件样式
                props: {},// 组件其他属性
            },
            {
                name: "search",
                mouseup: () => {
                    console.log("点击了search")
                },
                style: {},
                props: {},
            }
        ],
        right: [],// 代表右侧的配置，组件展示按照数组顺序展示
    }


    function init() {
        if (Object.keys(configs).length === 0) Object.assign(configs, {
            right: [
                {
                    name: "minSize"
                },
                {
                    name: "ToggleFullscreen",
                },
                {
                    name: "close",
                }
            ],
        })
    }

    init()

</script>

<div class="window-toolbar">
<!-- 为啥在这写css呢？无奈之举 -->
    <style>
        .top-bar-func-button {
            aspect-ratio: 1;
            display: flex;
            align-items: center;
            justify-content: center;
            height: 25px;
            --wails-draggable: no-drag;
            border-radius: 6px;
            transition: .1s ease;
            user-select: none;
            -webkit-user-drag: none;
            cursor: default;
        }

        .top-bar-func-button:hover {
            background-color: #d2d2d2;
        }

        .close:hover {
            background-color: red;

            svg {
                fill: #fff;
            }
        }
    </style>

    {#snippet preset(config = {})}
        {#if config.name === 'icon'}
            {@render icon(config)}
        {:else if config.name === 'div'}
            {@render div(config)}
        {:else if config.name === 'search'}
            {@render search(config)}
        {:else if config.name === 'button'}
            {@render button(config)}
        {:else if config.name === 'minSize'}
            {@render minSize(config)}
        {:else if config.name === 'hidden'}
            {@render hidden(config)}
        {:else if config.name === 'ToggleFullscreen'}
            {@render ToggleFullscreen(config)}
        {:else if config.name === 'close'}
            {@render close(config)}
        {/if}
    {/snippet}

    <div class="window-toolbar-left window-toolbar-item">
        {@render ll?.()}
        {#if configs.left?.length > 0}
            {#each configs.left as config }
                {@render preset(config)}
            {/each}
        {/if}
        {@render lr?.()}
    </div>

    <div class="window-toolbar-center window-toolbar-item">
        {@render cl?.()}
        {#if configs.center?.length > 0}
            {#each configs.center as config }
                {@render preset(config)}
            {/each}
        {/if}
        {@render cr?.()}
    </div>

    <div class="window-toolbar-right window-toolbar-item">
        {@render rl?.()}
        {#if configs.right?.length > 0}
            {#each configs.right as config }
                {@render preset(config)}
                <!--{@render hidden(config)}-->
            {/each}
        {/if}
        {@render rr?.()}
    </div>

</div>

<style lang="scss">

    .window-toolbar {
        position: sticky;
        top: 0;
        left: 0;
        width: 100%;
        height: 40px;
        display: flex;
        justify-content: space-between;
        align-items: center;
        user-select: none;
        -webkit-user-drag: none;
        -moz-user-select: none;
        -ms-user-select: none;
        z-index: 999;
        --wails-draggable: drag;
        padding: 0 6px;
    }

    .window-toolbar-item {
        display: flex;
        align-items: center;
        --wails-draggable: no-drag;
    }

    .window-toolbar-right {
        gap: 3px;
        justify-content: flex-end;
    }


</style>
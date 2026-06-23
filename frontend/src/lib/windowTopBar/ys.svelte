<script module>
    import {Window} from "@wailsio/runtime";

    export {hidden, ToggleFullscreen, close, button, icon, search, div, minSize}

    let isMax = $state(false)
    const toggleFull = () => {
        isMax = !isMax
        Window.ToggleMaximise()
    }

    const funcBtnHeight = '20px', funcBtnWidth = '15px'
</script>

<script>


</script>
<!-- 最小化按钮【预设】 -->
{#snippet minSize(config = {})}
    <button onmouseup={config.mouseup??Window.Minimise}
            style={config.style}
            class={["hidden" ,"top-bar-func-button", "ys", config.class]}>
        {#if config.inner}
            {config.inner}
        {:else}
            <svg viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" height={funcBtnHeight}
                 width={funcBtnWidth}>
                <path d="M929.8 528.1H93.5c-15.5 0-28-12.5-28-28s12.5-28 28-28h836.3c15.5 0 28 12.5 28 28s-12.5 28-28 28z"></path>
            </svg>
        {/if}
    </button>
{/snippet}

<!-- 隐藏按钮【预设】 -->
{#snippet hidden(config = {})}
    <button onmouseup={config.mouseup??Window.Hide}
            style={config.style}
            class={["hidden" ,"top-bar-func-button", "ys", config.class]}>
        {#if config.inner}
            {config.inner}
        {:else}
            <svg viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" height={funcBtnHeight}
                 width={funcBtnWidth}>
                <path d="M929.8 528.1H93.5c-15.5 0-28-12.5-28-28s12.5-28 28-28h836.3c15.5 0 28 12.5 28 28s-12.5 28-28 28z"></path>
            </svg>
        {/if}
    </button>
{/snippet}

<!-- 全屏切换按钮【预设】 -->
{#snippet ToggleFullscreen(config = {})}
    <button onmouseup={config.mouseup??toggleFull}
            style={config.style}
            class={["ToggleFullscreen" ,"top-bar-func-button", "ys", config.class]}>
        {#if config.inner}
            {config.inner}
        {:else}
            {#if isMax}
                <svg viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" height={funcBtnHeight}
                     width={funcBtnWidth}>
                    <path d="M812.2 65H351.6c-78.3 0-142.5 61.1-147.7 138.1-77 5.1-138.1 69.4-138.1 147.7v460.6c0 81.6 66.4 148 148 148h460.6c78.3 0 142.5-61.1 147.7-138.1 77-5.1 138.1-69.4 138.1-147.7V213c0-81.6-66.4-148-148-148z m-45.8 746.3c0 50.7-41.3 92-92 92H213.8c-50.7 0-92-41.3-92-92V350.7c0-50.7 41.3-92 92-92h460.6c50.7 0 92 41.3 92 92v460.6z m137.8-137.7c0 47.3-35.8 86.3-81.8 91.4V350.7c0-81.6-66.4-148-148-148H260.2c5.1-45.9 44.2-81.8 91.4-81.8h460.6c50.7 0 92 41.3 92 92v460.7z"
                    ></path>
                </svg>
            {:else}
                <svg viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" height={funcBtnHeight}
                     width={funcBtnWidth}>
                    <path d="M812.3 959.4H213.7c-81.6 0-148-66.4-148-148V212.9c0-81.6 66.4-148 148-148h598.5c81.6 0 148 66.4 148 148v598.5C960.3 893 893.9 959.4 812.3 959.4zM213.7 120.9c-50.7 0-92 41.3-92 92v598.5c0 50.7 41.3 92 92 92h598.5c50.7 0 92-41.3 92-92V212.9c0-50.7-41.3-92-92-92H213.7z"
                    ></path>
                </svg>
            {/if}
        {/if}
    </button>
{/snippet}

<!-- 关闭按钮【预设】 -->
{#snippet close(config = {})}
    <button onmouseup={config.mouseup??Window.Hide}
            style={config.style}
            class={["close" ,"top-bar-func-button" ,"ys",config.class]}>
        {#if config.inner}
            {config.inner}
        {:else}
            <svg viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" height={funcBtnHeight}
                 width={funcBtnWidth}>
                <path d="M96.6 915c-10.9-10.9-10.9-28.7 0-39.6L887 85c10.9-10.9 28.7-10.9 39.6 0 10.9 10.9 10.9 28.7 0 39.6L136.2 915c-10.9 10.9-28.6 10.9-39.6 0z"
                ></path>
                <path d="M887 915L96.6 124.6c-10.9-10.9-10.9-28.7 0-39.6 10.9-10.9 28.7-10.9 39.6 0l790.4 790.4c10.9 10.9 10.9 28.7 0 39.6-10.9 10.9-28.6 10.9-39.6 0z"
                ></path>
            </svg>
        {/if}
    </button>
{/snippet}

<!-- 普通div -->
{#snippet div(config = {})}
    <div
            onmouseup={config.mouseup}
            onclick={config.click}
            class={[config.class,'pt', 'top-bar-func-button']}
            style={config.style}
    >
        {@html config.inner}
    </div>
{/snippet}

<!-- 普通按钮 -->
{#snippet button(config = {})}
    <button
            onmouseup={config.mouseup}
            onclick={config.click}
            class={[config.class,'pt', 'top-bar-func-button']}
            style={config.style}
    >
        {@html config.inner}
    </button>
{/snippet}

<!-- 普通图标 -->
{#snippet icon(config = {})}
    <img src="{config.src}" draggable="false"
         onclick={config.click}
         alt=""
         class={["pt" , config.class]}
         onmouseup={config.mouseup} style={config.style}/>
{/snippet}

<!-- 普通搜索 -->
{#snippet search(config = {})}
    <input type="text"
           class={[config.class,'pt']}
           onclick={config.click}
           onmouseup={config.mouseup} style={config.style}/>
{/snippet}


<style lang="scss">
</style>
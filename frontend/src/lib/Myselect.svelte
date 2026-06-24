<script>
    let {options, value = $bindable()} = $props()
    let showText = $state('')
    let inputFocused = $state(false)
    let dropUp = $state(false)
    let boxEl = $state()

    $effect(() => {
        showText = options.find(o => o.value === value)?.text
    })

    const onfocus = () => {
        if (boxEl) {
            const rect = boxEl.getBoundingClientRect()
            // 估算下拉高度：每项约27px + padding 10px
            const estimatedHeight = options.length * 27 + 10
            const spaceBelow = window.innerHeight - rect.bottom
            dropUp = spaceBelow < estimatedHeight
        }
        inputFocused = true
    }

    const onblur = () => {
        setTimeout(() => {
            inputFocused = false
        }, 100)
    }

</script>

<div class="select-box" bind:this={boxEl}>
    <div class="input-container">
        <input class="input" type="text" placeholder="" bind:value={showText}
               {onfocus}
               {onblur}
        >
        <svg t="1782310007533" class="icon" class:icon-open={inputFocused} viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="8671" width="10"
             height="10">
            <path d="M0 749.136842L512 237.136842 1024 746.442105l-37.726316 37.726316L512 309.894737l-474.27368399 476.968421L0 749.136842zM0 749.136842"
                  p-id="8672"></path>
        </svg>
    </div>

    <div class="options" class:show-options={inputFocused} class:drop-up={dropUp}>
        <div class="aa">
            {#each options as {text, value: ov} ,index (ov)}
                <p onclick={()=> {value=ov}}>{text}</p>
            {/each}
        </div>
    </div>
</div>

<style lang="scss">
    .select-box {
        --bw: 220px;

        position: relative;
        width: var(--box-width, var(--bw));


        .input {
        }
    }

    .input-container {
        position: relative;
        display: flex;
        align-items: center;

        .icon {
            position: absolute;
            right: 8px;
            transform: rotate(180deg);
            transition: transform .2s ease;
            pointer-events: none;
            fill: #00000080;

            &.icon-open {
                transform: rotate(0deg);
            }
        }
    }

    .options {
        position: absolute;
        transform: scaleY(0);
        transform-origin: top;
        transition: opacity .15s ease, transform .15s ease;
        overflow: hidden;
        background-color: #fff;
        border-radius: 5px;
        z-index: 9;
        padding: 5px;
        width: var(--box-width, var(--bw));
        opacity: 0;

        .aa {
        }

        p {
            cursor: pointer;
            padding: 1px 10px;
            border-radius: 4px;
            transition: .1s ease;
            text-align: center;
            font-weight: 350;
        }

        p:hover {
            background-color: var(--main-6);
        }
    }

    .show-options {
        box-shadow: 0 0 5px rgba(117, 117, 117, 0.5);
        margin-top: 10px;
        transform: scaleY(1);
        opacity: 1;
        font-size: 15px;

        P {
            height: 25px;
            line-height: 25px;
        }
    }

    .drop-up {
        top: auto;
        bottom: 100%;
        margin-top: 0;
        transform-origin: bottom;

        &.show-options {
            margin-bottom: 10px;
        }
    }
</style>
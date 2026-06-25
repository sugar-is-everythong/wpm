<script>
    let {checked = $bindable(), onchange} = $props()
</script>

<label class="toggle">
    <input type="checkbox" bind:checked {onchange}/>
    <span class="toggle-slider"></span>
</label>

<style lang="scss">
    .toggle {
        position: relative;
        display: inline-block;
        width: 38px;
        height: 20px;
        flex-shrink: 0;

        input {
            opacity: 0;
            width: 0;
            height: 0;
        }

        .toggle-slider {
            position: absolute;
            cursor: pointer;
            top: 0;
            left: 0;
            right: 0;
            bottom: 0;
            background-color: #ccc;
            border-radius: 22px;
            transition: .2s;

            &::before {
                content: "";
                position: absolute;
                height: 14px;
                aspect-ratio: 1;
                left: 3px;
                bottom: 3px;
                background-color: white;
                border-radius: 50%;
                transition: .2s;
            }
        }

        input:checked + .toggle-slider {
            background-color: var(--main-5);

            &::before {
                transform: translateX(18px);
            }
        }

        input:focus + .toggle-slider {
            box-shadow: 0 0 0 3px var(--main-7);
        }
    }
</style>

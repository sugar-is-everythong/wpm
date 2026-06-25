<script>
    import Toggle from "./Toggle.svelte";
    import Myselect from "./Myselect.svelte";

    /**
     * items 大对象，key 为字段名，value 为配置项：
     * {
     *   fieldA: {
     *     label: string,           // 显示名称
     *     desc: string,            // 描述（可选）
     *     mode: 'toggle' | 'select' | 'text' | 'number',  // 默认 text
     *     value: any,              // 当前值
     *     options: [{text, value}], // select 专用
     *     hidden: boolean,         // 是否隐藏
     *     disabled: boolean,
     *     placeholder: string,
     *     onchange: (newVal) => void,
     *   },
     *   fieldB: { ... }
     * }
     *
     * 配合 format.js 的 ft() 可将此对象一键转为提交格式 {fieldA: value, fieldB: value}
     */
    let {items = {}} = $props()

    // 大对象转为有序数组供渲染使用
    const entries = $derived(Object.entries(items))
</script>

<div class="fast-form">
    {#each entries as [key, item]}
        {#if !item.hidden}
            <div class="form-item">
                <div class="item-info">
                    <div class="item-label">{item.label}</div>
                    {#if item.desc}
                        <div class="item-desc">{item.desc}</div>
                    {/if}
                </div>

                <div class="item-control">
                    {#if item.mode === 'toggle'}
                        <Toggle
                            checked={item.value}
                            onchange={(e) => item.onchange?.(e.target.checked)}
                        />

                    {:else if item.mode === 'select'}
                        <Myselect
                            options={item.options ?? []}
                            value={item.value}
                            onchange={item.onchange}
                        />

                    {:else if item.mode === 'number'}
                        <input type="number"
                               value={item.value}
                               placeholder={item.placeholder ?? item.label}
                               disabled={item.disabled}
                               onchange={(e) => item.onchange?.(+e.target.value)}
                        />

                    {:else}
                        <input type="text"
                               value={item.value}
                               placeholder={item.placeholder ?? item.label}
                               disabled={item.disabled}
                               onchange={(e) => item.onchange?.(e.target.value)}
                        />
                    {/if}
                </div>
            </div>
        {/if}
    {/each}
</div>

<style lang="scss">
    .fast-form {
        display: flex;
        flex-direction: column;
        gap: 20px;
    }

    .form-item {
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: 16px;

        .item-info {
            flex: 1;
            min-width: 0;

            .item-label {
                font-size: 14px;
                font-weight: 600;
            }

            .item-desc {
                font-size: 12px;
                color: var(--text-3);
                margin-top: 2px;
            }
        }

        .item-control {
            flex-shrink: 0;

            input {
                width: 120px;
            }
        }
    }
</style>

<script>
    import WindowTopBar from "@/lib/windowTopBar/index.svelte";
    import {ConfigService} from "~/configservice.js";

    let cfg = $state({
        autoStart: true,
        switchAllMonitor: false,
        recursiveDir: true,
        strategy: 2,
        registerContextMenu: true,
    });
    let change = $state(false);
    let saving = $state(false);

    const toggleItems = [
        {key: "autoStart", label: "启动后立即执行", desc: "程序启动时是否立即开始自动切换壁纸"},
        {key: "switchAllMonitor", label: "切换所有显示器", desc: "自动切换时是否同时切换所有显示器壁纸"},
        {key: "recursiveDir", label: "目录递归扫描", desc: "扫描壁纸目录时是否递归包含子目录中的图片"},
        {key: "registerContextMenu", label: "注册右键菜单", desc: "是否将程序注册到系统右键菜单"},
    ];

    const strategyOptions = [
        {value: 0, label: "随机切换"},
        {value: 1, label: "顺序切换"},
        {value: 2, label: "鼠标所在屏幕"},
    ];

    // 加载配置
    (() => {
        ConfigService.GetConfig().then(res => { Object.assign(cfg, res); });
    })();

    const markChange = () => { change = true; };

    const save = async () => {
        saving = true;
        try {
            await ConfigService.SaveAppConfig(cfg);
            change = false;
        } catch (e) {
            console.error("保存配置失败:", e);
        } finally {
            saving = false;
        }
    };

    const cancel = () => {
        ConfigService.GetConfig().then(res => { Object.assign(cfg, res); change = false; });
    };
</script>

<main class="config-box">
    <div class="top">
        <WindowTopBar --hover-color='var(--main-9)'/>
    </div>

    <div class="config-content">
        <div class="config-section">
            <div class="section-title">通用设置</div>
            {#each toggleItems as item}
                <div class="config-item">
                    <div class="item-info">
                        <div class="item-label">{item.label}</div>
                        <div class="item-desc">{item.desc}</div>
                    </div>
                    <label class="toggle">
                        <input type="checkbox" bind:checked={cfg[item.key]} onchange={markChange}/>
                        <span class="toggle-slider"></span>
                    </label>
                </div>
            {/each}
        </div>

        <div class="config-section">
            <div class="section-title">切换策略</div>
            <div class="strategy-group">
                {#each strategyOptions as opt}
                    <label class="strategy-item" class:active={cfg.strategy === opt.value}>
                        <input type="radio" name="strategy" value={opt.value} bind:group={cfg.strategy} onchange={markChange}/>
                        <span class="strategy-radio"></span>
                        <span class="strategy-label">{opt.label}</span>
                    </label>
                {/each}
            </div>
        </div>
    </div>

    {#if change}
        <div class="bottom-box">
            <button class="ok" onclick={save} disabled={saving}>{saving ? '保存中...' : '确定'}</button>
            <button onclick={cancel}>取消</button>
        </div>
    {/if}
</main>

<style lang="scss">
    .config-box {
        display: grid;
        height: 100vh;
        grid-template-rows: 40px 1fr auto;
        background-color: rgba(243, 243, 243, 0.8);
        border-radius: 4px;
        font-family: "Segoe UI Variable", "Segoe UI", system-ui, sans-serif;
    }

    .config-content {
        padding: 16px 20px;
        overflow-y: auto;
        display: flex;
        flex-direction: column;
        gap: 16px;
    }

    .config-section {
        background: rgba(255, 255, 255, 0.75);
        border: 1px solid rgba(0, 0, 0, 0.06);
        border-radius: 8px;
        padding: 16px 20px;
        display: flex;
        flex-direction: column;
        gap: 14px;

        .section-title {
            font-size: 14px;
            font-weight: 600;
            color: var(--main-1);
            padding-bottom: 6px;
            border-bottom: 1px solid var(--text-7);
        }
    }

    .config-item {
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: 16px;

        .item-info {
            flex: 1;
            min-width: 0;

            .item-label {
                font-size: 14px;
                font-weight: 500;
                color: var(--text-1);
            }

            .item-desc {
                font-size: 12px;
                color: var(--text-3);
                margin-top: 2px;
            }
        }
    }

    /* Toggle 开关样式 */
    .toggle {
        position: relative;
        display: inline-block;
        width: 40px;
        height: 22px;
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
                height: 16px;
                width: 16px;
                left: 3px;
                bottom: 3px;
                background-color: white;
                border-radius: 50%;
                transition: .2s;
            }
        }

        input:checked + .toggle-slider {
            background-color: var(--main-1);

            &::before {
                transform: translateX(18px);
            }
        }

        input:focus + .toggle-slider {
            box-shadow: 0 0 0 3px var(--main-7);
        }
    }

    /* 单选策略组 */
    .strategy-group {
        display: flex;
        flex-direction: column;
        gap: 8px;
    }

    .strategy-item {
        display: flex;
        align-items: center;
        gap: 10px;
        padding: 10px 14px;
        border-radius: 6px;
        cursor: pointer;
        transition: background .15s;
        user-select: none;

        &:hover {
            background: rgba(0, 0, 0, 0.03);
        }

        &.active {
            background: var(--main-7);
        }

        input {
            display: none;
        }

        .strategy-radio {
            width: 16px;
            height: 16px;
            border: 2px solid var(--text-4);
            border-radius: 50%;
            position: relative;
            flex-shrink: 0;
            transition: border-color .15s;

            &::after {
                content: "";
                position: absolute;
                top: 50%;
                left: 50%;
                transform: translate(-50%, -50%) scale(0);
                width: 8px;
                height: 8px;
                border-radius: 50%;
                background: var(--main-1);
                transition: transform .15s;
            }
        }

        &.active .strategy-radio {
            border-color: var(--main-1);

            &::after {
                transform: translate(-50%, -50%) scale(1);
            }
        }

        .strategy-label {
            font-size: 14px;
            color: var(--text-1);
        }
    }

    .bottom-box {
        padding: 5px 10px;
        padding-bottom: 10px;
        display: flex;
        flex-direction: row-reverse;
        gap: 10px;

        button {
            padding: 5px 20px;
            border-radius: 5px;
            transition: background-color .1s ease;
            border: 1px solid #ccc;

            &:disabled {
                opacity: 0.6;
                cursor: not-allowed;
            }
        }

        .ok {
            background-color: var(--main-2);
            border: 1px solid transparent;
            color: #fff;
        }
    }
</style>

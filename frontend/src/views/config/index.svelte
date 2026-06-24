<script>
    import WindowTopBar from "@/lib/windowTopBar/index.svelte";
    import {ConfigService} from "~/index.js";
    import {route} from "@mateothegreat/svelte5-router";
    import Myselect from "@/lib/Myselect.svelte";

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
        {key: "autoStart", label: "启动后立即执行", desc: "程序启动时是否立即切换一次壁纸"},
        {key: "switchAllMonitor", label: "协同切换", desc: "自动切换时是否同时切换所有显示器壁纸"},
        {key: "recursiveDir", label: "目录递归扫描", desc: "扫描壁纸目录时会一直向下找到所有图片随机设置壁纸"},
        {key: "registerContextMenu", label: "注册右键菜单", desc: "是否将程序注册到系统右键菜单"},
    ];

    const strategyOptions = [
        {value: 0, text: "随机切换"},
        {value: 1, text: "顺序切换"},
        {value: 2, text: "鼠标所在屏幕"},
    ];

    // 加载配置
    (() => {
        ConfigService.GetConfig().then(res => {
            Object.assign(cfg, res);
        });
    })();

    const markChange = () => {
        change = true;
    };

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
        ConfigService.GetConfig().then(res => {
            Object.assign(cfg, res);
            change = false;
        });
    };
</script>

<main class="config-box">
    <div class="top">
        <WindowTopBar>
            {#snippet ll()}
                <div style="display: flex; align-items: flex-start; gap:3px;">
                    <a href="/" use:route>
                        <svg t="1782308980340" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="8478"
                             width="20" height="20">
                            <path d="M224.32 505.6a31.936 31.936 0 0 1 10.88-19.84l222.08-222.08a32 32 0 0 1 45.12 45.12l-169.28 169.28H768a32 32 0 0 1 0 64H333.12l169.28 169.28a32 32 0 1 1-45.12 45.44l-224-224a31.968 31.968 0 0 1-8.96-27.2z"
                                  fill="#202425" p-id="8479"></path>
                        </svg>
                    </a>
                    壁纸切换策略
                </div>
            {/snippet}
        </WindowTopBar>
    </div>

    <div class="config-content">
        <div class="config-section">
            <div class="section-title">切换策略</div>
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

            <div class="config-item">
                <div class="item-info">
                    <div class="item-label">切换策略</div>
                    <div class="item-desc">选择壁纸切换的方式</div>
                </div>
                <Myselect options={strategyOptions} bind:value={cfg.strategy} />
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
    .top {
        padding-left: 5px;
        font-size: 14px;
    }

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
            color: var(--main-4);
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
            background-color: var(--main-4);

            &::before {
                transform: translateX(18px);
            }
        }

        input:focus + .toggle-slider {
            box-shadow: 0 0 0 3px var(--main-7);
        }
    }



    .bottom-box {
        padding: 5px 10px;
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

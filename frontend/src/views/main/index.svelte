<script>
    import WindowTopBar from "@/lib/windowTopBar/index.svelte";
    import {Dialogs, Screens, Window} from "@wailsio/runtime";
    import {route} from "@mateothegreat/svelte5-router";
    import ShowOK from "@/lib/ShowOK.svelte";
    import {appConfig, saveConfig} from "@/lib/config.svelte.js";

    let display = $state(0);
    let change = $state(false);

    Window.SetAlwaysOnTop(true);
    (() => {
        Screens.GetAll().then((screens) => {
            display = screens;
            for (let i = 0; i < screens.length; i++) {
                if (!appConfig.monitorDirs[i]) appConfig.monitorDirs[i] = [];
            }
        });
    })()

    const openDir = async (monitorIndex) => {
        let res = await Dialogs.OpenFile({
            CanChooseDirectories: true,
            AllowsMultipleSelection: true,
            CanChooseFiles: false,
        })
        if (res && res.length > 0) {
            appConfig.monitorDirs[monitorIndex] = [...(appConfig.monitorDirs[monitorIndex] || []), ...res];
        }
        change = true
    }

    const removeDir = (monitorIndex, dirIndex) => {
        appConfig.monitorDirs[monitorIndex] = appConfig.monitorDirs[monitorIndex].filter((_, i) => i !== dirIndex);
        change = true
    }

    const okSet = async () => {
        try {
            await saveConfig();
            change = false;
        } catch (e) {
            console.error("保存失败:", e);
        }
    }

    const cancel = async () => {
        const {ConfigService} = await import('~/index.js')
        const res = await ConfigService.GetConfig()
        Object.assign(appConfig, res)
        change = false;
    }

    const viewSm = () => {}
</script>

<main class="main-box">
    <div class="top">
        <WindowTopBar>
            {#snippet ll()}
                <div style="padding-left: 15px;font-size: 14px;font-family: 'hxb', sans-serif;">塞壁纸</div>
            {/snippet}
        </WindowTopBar>
    </div>

    <div class="set-box">
        {#each display as d, i}
            <div class="d-s">
                <div class="d-s-title" style:color={d.IsPrimary?'var(--main-4)':''}>
                    <img src="/dmcz.png" alt="" height="20px">
                    <div>显示器 {i + 1}</div>
                    <span>
                        宽：{d.Size.Width } 高：{d.Size.Height} | {d.Size.Width > d.Size.Height ? '横' : '纵'}
                        向显示器
                    </span>
                </div>

                <div class="select">
                    <input type="number" placeholder="切换间隔(分钟)" bind:value={appConfig.interval} onchange={()=> {change=true}}>
                    <span class="input-tip" title="设置壁纸自动切换的时间间隔，单位为分钟">⚠</span>
                    <button class="dir-add" onclick={() => openDir(i)}>
                        <img src="wfjmjw.png" alt="" height="19px">
                        选择
                    </button>
                </div>

                <div class="d-s-dirs">
                    {#each (appConfig.monitorDirs[i] || []) as dir, j}
                        <div class="dir-item">
                            <span class="dir-path" title={dir}> {dir} </span>
                            <button class="dir-del" onclick={() => removeDir(i, j)}>✕</button>
                        </div>
                    {/each}
                </div>
            </div>
        {/each}
    </div>

    <ShowOK show={change} ok={okSet} cancel={cancel}></ShowOK>

    <div class="bottom-box">
        <button onclick={viewSm} title="使用说明">🔖</button>
        <a title="设置页面" href="/#/config" use:route>⚙</a>
    </div>
</main>


<style lang="scss">
    .main-box {
        display: grid;
        height: 100vh;
        grid-template-rows: 40px 1fr auto auto;
        background-color: rgba(243, 243, 243, 0.92);
        border-radius: 4px;
    }

    .set-box {
        padding: 16px 20px;
        overflow-y: auto;
        display: flex;
        flex-direction: column;
        gap: 16px;

        .d-s {
            background: rgba(255, 255, 255, 0.75);
            border-radius: 8px;
            padding: 16px 20px;
            display: flex;
            flex-direction: column;
            gap: 12px;

            // 显示器标题
            .d-s-title {
                user-select: none;
                display: flex;
                gap: 10px;
                align-items: flex-end;
                font-family: 'hxb', sans-serif;

                span {
                    font-size: 13px;
                    color: var(--text-5);
                    transform: translateY(-1px);
                }
            }

            .select {
                font-size: 14px;
                font-weight: 600;
                display: flex;
                align-items: center;
                gap: 12px;

                input {
                    flex-shrink: 0;
                    width: 110px;
                    height: 28px;
                    padding: 0 8px;
                    font-size: 13px;
                    border-radius: 4px;
                    background: rgba(255, 255, 255, 0.6);
                }

                .input-tip {
                    font-size: 12px;
                    color: var(--text-4);
                    line-height: 1;
                    cursor: default;

                    &:hover {
                        color: var(--main-1);
                    }
                }

                .dir-add {
                    position: sticky;
                    left: 9999px;
                    border-radius: 3px;
                    border: 1px #ccc solid;
                    padding: 3px 10px;
                    transition: 0.1s ease;
                    display: flex;
                    align-items: center;
                    gap: 5px;
                }
            }

            // 目录列表
            .d-s-dirs {
                display: flex;
                flex-direction: column;
                gap: 2px;
                padding-left: 12px;
                border-left: 2px solid var(--text-7);

                .dir-item {
                    display: flex;
                    align-items: center;
                    gap: 8px;
                    padding: 5px 8px;
                    border-radius: 4px;
                    font-size: 13px;
                    transition: background 0.1s;

                    &:hover {
                        background: rgba(0, 0, 0, 0.04);
                    }

                    .dir-path {
                        flex: 1;
                        overflow: hidden;
                        text-overflow: ellipsis;
                        white-space: nowrap;
                        color: var(--text-2);
                    }

                    .dir-del {
                        flex-shrink: 0;
                        width: 22px;
                        height: 22px;
                        display: flex;
                        align-items: center;
                        justify-content: center;
                        font-size: 10px;
                        color: var(--text-4);
                        border-radius: 3px;
                        transition: background 0.1s;

                        &:hover {
                            background: rgba(196, 43, 28, 0.08);
                            color: #c42b1c;
                        }
                    }
                }

            }

        }
    }

    .bottom-box {
        padding: 5px 10px;
        padding-bottom: 10px;
        display: flex;
        justify-content: space-between;
        gap: 10px;

        button {
            font-size: 20px;
        }

        button:first-of-type {
            font-size: 15px;
        }
    }
</style>

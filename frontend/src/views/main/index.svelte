<script>
    import WindowTopBar from "@/lib/windowTopBar/index.svelte";
    import {Dialogs, Screens, Window} from "@wailsio/runtime";
    import {route} from "@mateothegreat/svelte5-router";

    let display = $state(0);
    let dirMap = $state({}); // { 0: ["C:/Images", "D:/Photos"], 1: [] }
    let change = $state(false); // 修改标志用于显示确认按钮

    Window.SetAlwaysOnTop(true);
    (() => {
        Screens.GetAll().then((num) => {
            display = num
            // 初始化每个显示器的目录列表
            for (let i = 0; i < num; i++) {
                if (!dirMap[i]) dirMap[i] = [];
            }
        })
    })()

    const openDir = async (monitorIndex) => {
        let res = await Dialogs.OpenFile({
            CanChooseDirectories: true,
            AllowsMultipleSelection: true,
            CanChooseFiles: false,
        })
        if (res && res.length > 0) {
            dirMap[monitorIndex] = [...(dirMap[monitorIndex] || []), ...res];
        }

        change = true
    }

    // 移除目录
    const removeDir = (monitorIndex, dirIndex) => {
        dirMap[monitorIndex] = dirMap[monitorIndex].filter((_, i) => i !== dirIndex);
        change = true
    }

    // 确定设置
    const okSet = () => {

    }

    // 查看说明
    const viewSm = () => {

    }
</script>

<main class="main-box">
    <div class="top">
        <WindowTopBar>
            {#snippet ll()}
                <div style="padding-left: 15px;font-size: 14px">塞壁纸</div>
            {/snippet}
        </WindowTopBar>
    </div>

    <div class="set-box">
        {#each display as d, i}
            <div class="d-s">
                <div class="d-s-title" style:color={d.IsPrimary?'var(--main-4)':''}>
                    🖥️ 显示器 {i + 1}
                    <span> 宽：{d.Size.Width } 高：{d.Size.Height} | {d.Size.Width > d.Size.Height ? '横' : '纵'}
                        向显示器</span>
                </div>
                <div class="select">
                    <input type="number" placeholder="切换间隔">
                    <span class="input-tip" title="设置壁纸自动切换的时间间隔，单位为分钟">⚠</span>
                    <button class="dir-add" onclick={() => openDir(i)}>选择目录</button>
                </div>

                <div class="d-s-dirs">
                    {#each (dirMap[i] || []) as dir, j}
                        <div class="dir-item">
                            <span class="dir-path" title={dir}> {dir} </span>
                            <button class="dir-del" onclick={() => removeDir(i, j)}>✕</button>
                        </div>
                    {/each}
                </div>
            </div>
        {/each}
    </div>

    {#if change}
        <div class="o-c-box">
            <button class="ok" onclick={okSet}>确定</button>
            <button>取消</button>
        </div>
    {/if}


    <div class="bottom-box">
        <button onclick={viewSm} title="使用说明">🔖</button>
        <a title="设置页面" href="/config" use:route>⚙</a>
    </div>
</main>


<style lang="scss">
    .main-box {
        display: grid;
        height: 100vh;
        grid-template-rows: 40px 1fr auto auto;
        background-color: rgba(243, 243, 243, 0.8);
        border-radius: 4px;
        font-family: "Segoe UI Variable", "Segoe UI", system-ui, sans-serif;
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

                span {
                    font-size: 13px;
                    color: var(--text-5);
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
                    width: 80px;
                    height: 28px;
                    padding: 0 8px;
                    font-size: 13px;
                    border-radius: 4px;
                    background: rgba(255, 255, 255, 0.6);
                    //border-bottom: 2px solid rgba(0,0,0,0);
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

    .o-c-box {
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
        }

        button:active {
        }

        .ok {
            background-color: var(--main-3);
            border: 1px solid transparent;
            color: #fff;
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
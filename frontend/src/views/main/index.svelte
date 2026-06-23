<script>
    import WindowTopBar from "@/lib/windowTopBar/index.svelte";
    import {Dialogs, Screens} from "@wailsio/runtime";

    let display = $state(0);
    let dirMap = $state({}); // { 0: ["C:/Images", "D:/Photos"], 1: [] }

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
    }

    const removeDir = (monitorIndex, dirIndex) => {
        dirMap[monitorIndex] = dirMap[monitorIndex].filter((_, i) => i !== dirIndex);
    }
</script>

<main class="main-box">
    <div class="top">
        <WindowTopBar --hover-color='var(--main-9)'/>
    </div>

    <div class="set-box">
        {#each Array(display) as _, i}
            <div class="d-s">
                <div class="d-s-title">🖥️ 显示器 {i + 1}
                    <input type="number" placeholder="切换间隔">
                    <span class="input-tip" title="设置壁纸自动切换的时间间隔，单位为分钟">⚠</span>
                    <button class="dir-add" onclick={() => openDir(i)}>＋ 选择目录</button>
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
</main>


<style lang="scss">
    .main-box {
        height: 100vh;
        display: grid;
        grid-template-rows: 40px 1fr;
        background: #f3f3f3;
        border-radius: 8px;
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
            border: 1px solid rgba(0, 0, 0, 0.06);
            border-radius: 8px;
            padding: 16px 20px;
            display: flex;
            flex-direction: column;
            gap: 12px;

            .d-s-title {
                font-size: 14px;
                font-weight: 600;
                color: var(--text-1);
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
                    border: 1px solid var(--text-6);
                    background: rgba(255, 255, 255, 0.6);

                    &:hover {
                        border-color: var(--text-4);
                    }

                    &:focus {
                        border-color: var(--main-1);
                        border-bottom: 2px solid var(--main-1);
                        background: #fff;
                    }
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
            }

            .d-s-dirs {
                display: flex;
                flex-direction: column;
                gap: 2px;
                padding-left: 12px;
                border-left: 2px solid var(--text-7);
            }

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

            .dir-add {
                flex-shrink: 0;
                margin-left: auto;
                padding: 4px 14px;
                height: 28px;
                font-size: 12px;
                border-radius: 4px;
                transition: background 0.1s;

                &:hover {
                    //background: var(--main-2);
                    background: #e3e3e3;
                }
            }
        }
    }
</style>
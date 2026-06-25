<script>
    import WindowTopBar from "@/lib/windowTopBar/index.svelte";
    import {route} from "@mateothegreat/svelte5-router";
    import FastForm from "@/lib/fast-form.svelte";
    import ShowOK from "@/lib/ShowOK.svelte";
    import {appConfig, saveConfig} from "@/lib/config.svelte.js";

    let change = $state(false);
    const markChange = () => {
        change = true;
    };

    const formItems = $derived({
        autoStart: {
            label: "启动后立即执行", desc: "程序启动时是否立即切换一次壁纸",
            mode: 'toggle', value: appConfig.autoStart,
            onchange: (v) => {
                appConfig.autoStart = v;
                markChange();
            }
        },
        strategy: {
            label: "切换策略", desc: "屏幕切换的方式",
            mode: 'select', value: appConfig.strategy,
            options: [
                {value: 0, text: "协同切换"}, {value: 1, text: "顺序屏幕切换"},
                {value: 2, text: "随机屏幕切换"}, {value: 3, text: "鼠标所在屏幕"}
            ],
            onchange: (v) => {
                appConfig.strategy = v;
                markChange();
            }
        },
        recursiveDir: {
            label: "目录递归扫描", desc: "扫描壁纸目录时会一直向下找到所有图片",
            mode: 'toggle', value: appConfig.recursiveDir,
            onchange: (v) => {
                appConfig.recursiveDir = v;
                markChange();
            }
        },
        registerContextMenu: {
            label: "注册右键菜单", desc: "是否将程序注册到系统右键菜单",
            mode: 'toggle', value: appConfig.registerContextMenu,
            onchange: (v) => {
                appConfig.registerContextMenu = v;
                markChange();
            }
        },
        autoMatchRatio: {
            label: "自动匹配显示器比例", desc: "切换壁纸时自动过滤与显示器比例相近的图片",
            mode: 'toggle', value: appConfig.autoMatchRatio,
            onchange: (v) => {
                appConfig.autoMatchRatio = v;
                markChange();
            }
        },
        wallpaperFit: {
            label: "壁纸应用模式", desc: "壁纸在显示器上的裁剪拉伸方式",
            mode: 'select', value: appConfig.wallpaperFit,
            hidden: appConfig.autoMatchRatio,
            options: [
                {value: 0, text: "居中"}, {value: 1, text: "平铺"}, {value: 2, text: "拉伸"},
                {value: 3, text: "适应"}, {value: 4, text: "填充"}, {value: 5, text: "跨屏"}
            ],
            onchange: (v) => {
                appConfig.wallpaperFit = v;
                markChange();
            }
        },
    });

    const save = async () => {
        try {
            await saveConfig();
            change = false;
        } catch (e) {
            console.error("保存配置失败:", e);
        }
    };

    // 取消：重新从 Go 加载覆盖本地修改
    const cancel = async () => {
        const {ConfigService} = await import('~/index.js')
        const res = await ConfigService.GetConfig()
        Object.assign(appConfig, res)
        change = false;
    };
</script>

<main class="config-box">
    <div class="top">
        <WindowTopBar>
            {#snippet ll()}

                <div style="display: flex; align-items:center; gap:3px; font-family: 'hxb', sans-serif;">
                    <a href="/#/" use:route style="display: flex;align-items: center">
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
            <FastForm items={formItems}/>
        </div>
    </div>

    <ShowOK show={change} ok={save} cancel={cancel}></ShowOK>
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
        background-color: rgb(255 255 255/0.9);
        border-radius: 4px;
        font-family: "Segoe UI Variable", "Segoe UI", system-ui, sans-serif;
        user-select: none;
    }

    .config-content {
        padding: 16px 20px;
        overflow-y: auto;
        display: flex;
        flex-direction: column;
        gap: 16px;
    }

    .config-section {
        display: flex;
        flex-direction: column;
        gap: 20px;
    }

</style>

import {ConfigService} from "~/index.js";

// 全局响应式配置，启动时加载一次，各页面直接 import 使用
export const appConfig = $state({
    autoStart: true,
    recursiveDir: true,
    strategy: 2,
    registerContextMenu: true,
    autoMatchRatio: true,
    wallpaperFit: 4,
    monitorDirs: {}, // {0:{dirs:[], d:0, Interval:0}} 屏幕索引：{文件目录、优先级（未实现）、切换间隔}
    autoRun: false,
})

let loaded = false

/**
 * 从 Go 加载配置到内存，启动时调用一次
 */
export async function loadConfig() {
    if (loaded) return
    loaded = true
    const cfg = await ConfigService.GetConfig()
    Object.assign(appConfig, structuredClone(cfg))
}

/**
 * 保存当前内存配置到 Go / 文件
 */
export async function saveConfig() {

    for (let k in appConfig.monitorDirs) {
        const m = appConfig.monitorDirs[k]
        if (m.Dirs.length === 0) return "壁纸目录不能为空"

        m.Interval = Math.ceil(m.Interval)
        if (m.Interval <= 0) return "切换间隔必须大于0"
    }

    await ConfigService.Save({...appConfig})
}
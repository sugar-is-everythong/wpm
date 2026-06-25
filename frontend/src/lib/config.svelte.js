import {ConfigService} from "~/index.js";

// 全局响应式配置，启动时加载一次，各页面直接 import 使用
export const appConfig = $state({
    autoStart: true,
    recursiveDir: true,
    strategy: 2,
    registerContextMenu: true,
    autoMatchRatio: true,
    wallpaperFit: 4,
    monitorDirs: {},
    interval: 30,
})

let loaded = false

/**
 * 从 Go 加载配置到内存，启动时调用一次
 */
export async function loadConfig() {
    if (loaded) return
    loaded = true
    const cfg = await ConfigService.GetConfig()
    Object.assign(appConfig, cfg)
}

/**
 * 保存当前内存配置到 Go / 文件
 */
export async function saveConfig() {
    await ConfigService.Save({...appConfig})
}

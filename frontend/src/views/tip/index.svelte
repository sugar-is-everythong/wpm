<script>
	import {GetDisplay, ManualSwitch} from '~/greetservice.js'
	import {Screens, Window,Application} from "@wailsio/runtime";

	const opts = $state([
		[{
			label: '切换所有', icon: "🗘", color: 'var(--main-1)',
			click: () => {
				GetDisplay().then((data) => {
					for (const di in data) {
						ManualSwitch(di | 0)
					}
				})
			}
		}], [],
		[
			{
				label: '退出', icon: "×", color: "red",
				click() {
					console.log(111)
					Application.Quit()
				}
			},
		]
	]);

	(() => {
		Screens.GetAll().then((obj) => {
			const num = obj.length
			for (let i = 0; i < num; i++) {
				opts[1].push({
					label: `切换屏幕${ i + 1 }`, icon: "🗘", click: (di) => {
						ManualSwitch(di)
					}
				})
			}

			// 设置窗口大小
			Window.SetSize(150, 127 + (num - 1) * 33)
		})
	})()
</script>

<main class="tip-view">
    {#each opts as group,gi}
        <div class="ctx-menu-group" onclick={Window.Hide}>
            {#each group as item,di}
                <div class="ctx-menu-item" style:color={item.color} onclick={()=> item.click?.(di)}>
                    {#if item.icon}
                        <div class="icon">{item.icon}</div>
                    {/if}

                    <span>{item.label}</span>
                </div>
            {/each}
        </div>
    {/each}
</main>

<style lang="scss">

    .tip-view {
        background-color: rgb(255 255 255 / 0.95);
        backdrop-filter: blur(10px);
        border-radius: 4px;
        font-size: 13px;
        padding: 4px 10px;
        cursor: pointer;
        user-select: none;
        margin: 10px;
        box-shadow: 0 0 5px rgb(166 166 166 / 0.8);

        ::-webkit-scrollbar {
            display: none;
        }
    }

    .ctx-menu-group {
        position: relative;
        padding: 5px 0;
    }

    .ctx-menu-group:not(:last-of-type)::after {
        position: absolute;
        bottom: 0;
        left: 50%;
        transform: translateX(-50%);
        content: "";
        height: 1px;
        width: 100%;
        background-color: rgba(204, 204, 204, 0.65);
    }

    .ctx-menu-item {
        position: relative;
        padding: 2px 5px;
        white-space: nowrap;
        //cursor: pointer;
        display: grid;
        grid-template-columns: 18px 1fr;
        grid-gap: 5px;
        align-items: center;
        border-radius: 4px;

        img {
            aspect-ratio: 1;
        }
    }

    .icon {
        text-align: center;
    }

    .ctx-menu-item:hover {
        background-color: rgba(48, 88, 98, 0.15);
    }
</style>
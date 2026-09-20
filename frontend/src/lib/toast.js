// 默认图标配置
const defaultIcons = {
    success: '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path fill="currentColor" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/></svg>',
    warning: '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path fill="currentColor" d="M1 21h22L12 2 1 21zm12-3h-2v-2h2v2zm0-4h-2v-4h2v4z"/></svg>',
    error: '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path fill="currentColor" d="M12 2C6.47 2 2 6.47 2 12s4.47 10 10 10 10-4.47 10-10S17.53 2 12 2zm5 13.59L15.59 17 12 13.41 8.41 17 7 15.59 10.59 12 7 8.41 8.41 7 12 10.59 15.59 7 17 8.41 13.41 12 17 15.59z"/></svg>',
};

// 根据类型设置背景色
const typeColors = {
    success: 'rgba(24, 160, 88, 1)',
    warning: 'rgba(240, 160, 32, 1)',
    error: 'rgba(208, 48, 80, 1)',
};


// Toast容器管理
const toastContainer = {
    element: null,
    toasts: [],

    init() {
        if (!this.element) {
            this.element = document.createElement('div');
            this.element.id = 'toast-container';
            this.element.style.position = 'fixed';
            this.element.style.zIndex = '9999';
            this.element.style.pointerEvents = 'none';
            document.body.appendChild(this.element);
        }
    },

    addToast(toast) {
        this.toasts.push(toast);
        this.updateContainerPosition(toast.options.position);
    },

    removeToast(toast) {
        this.toasts = this.toasts.filter(t => t !== toast);
        if (this.toasts.length === 0) {
            this.element.style.top = '';
            this.element.style.bottom = '';
            this.element.style.left = '';
            this.element.style.right = '';
        }
    },

    updateContainerPosition(position) {
        const positions = {
            'top-center': { top: '24px', left: '50%', transform: 'translateX(-50%)' },
            'top-left': { top: '24px', left: '24px' },
            'top-right': { top: '24px', right: '24px' },
            'bottom-center': { bottom: '24px', left: '50%', transform: 'translateX(-50%)' },
            'bottom-left': { bottom: '24px', left: '24px' },
            'bottom-right': { bottom: '24px', right: '24px' },
        };

        const pos = positions[position] || positions['top-center'];
        Object.assign(this.element.style, pos);
    },
};

// Toast类
class Toast {
    constructor(options = {}) {
        this.options = {
            type: options.type || 'success', // success, warning, error
            message: options.message || '',
            icon: options.icon || null,
            duration: options.duration !== undefined ? options.duration : 2000,
            position: options.position || 'top-center',
            onClose: options.onClose || null,
        };

        this.element = null;
        this.timer = null;
        this.init();
    }

    init() {
        // 确保容器存在
        toastContainer.init();

        // 创建toast元素
        this.element = document.createElement('div');
        this.element.className = `toast toast-${this.options.type}`;

        // 设置样式
        this.applyStyles();

        // 创建内容
        this.createContent();

        // 添加到容器
        toastContainer.addToast(this);

        // 计算位置偏移
        this.calculateOffset();

        // 显示动画
        setTimeout(() => {
            this.element.style.opacity = '1';
            this.element.style.transform = 'translateY(0)';
        }, 10);

        // 设置自动关闭
        if (this.options.duration > 0) {
            this.timer = setTimeout(() => {
                this.close();
            }, this.options.duration);
        }
    }

    applyStyles() {
        // 基础样式
        this.element.style.display = 'flex';
        this.element.style.alignItems = 'center';
        this.element.style.padding = '10px 16px';
        this.element.style.marginBottom = '12px';
        this.element.style.borderRadius = '6px';
        this.element.style.boxShadow = '0 3px 6px -4px rgba(0, 0, 0, 0.12), 0 6px 16px 0 rgba(0, 0, 0, 0.08), 0 9px 28px 8px rgba(0, 0, 0, 0.05)';
        this.element.style.pointerEvents = 'auto';
        this.element.style.transition = 'all 0.3s ease';
        this.element.style.opacity = '0';
        this.element.style.transform = 'translateY(-20px)';
        this.element.style.maxWidth = '400px';
        this.element.style.backgroundColor = '#fff';

    }

    createContent() {
        // 创建图标容器
        const iconContainer = document.createElement('div');
        iconContainer.style.marginRight = '12px';
        iconContainer.style.display = 'flex';
        iconContainer.style.alignItems = 'center';
        iconContainer.style.color = typeColors[this.options.type] || typeColors.success

        iconContainer.style.alignItems = 'center';

        // 添加图标
        if (this.options.icon) {
            // 自定义图标
            if (typeof this.options.icon === 'string') {
                // 如果是URL，创建img元素
                if (this.options.icon.startsWith('http') || this.options.icon.startsWith('data:image')) {
                    const img = document.createElement('img');
                    img.src = this.options.icon;
                    img.style.width = '20px';
                    img.style.height = '20px';
                    iconContainer.appendChild(img);
                } else {
                    // 如果是SVG字符串，直接插入
                    iconContainer.innerHTML = this.options.icon;
                }
            } else if (this.options.icon instanceof HTMLElement) {
                // 如果是DOM元素，直接添加
                iconContainer.appendChild(this.options.icon);
            }
        } else {
            // 使用默认图标
            iconContainer.innerHTML = defaultIcons[this.options.type] || defaultIcons.success;
            iconContainer.querySelector('svg').style.width = '20px';
            iconContainer.querySelector('svg').style.height = '20px';
        }

        // 创建消息文本
        const messageElement = document.createElement('div');
        messageElement.textContent = this.options.message;
        messageElement.style.fontSize = '14px';
        messageElement.style.lineHeight = '1.5';

        // 添加到toast元素
        this.element.appendChild(iconContainer);
        this.element.appendChild(messageElement);

        // 添加到容器
        toastContainer.element.appendChild(this.element);
    }

    calculateOffset() {
        // 计算当前toast的偏移量，避免重叠
        const index = toastContainer.toasts.indexOf(this);
        if (index > 0) {
            const previousToast = toastContainer.toasts[index - 1];
            const previousHeight = previousToast.element.offsetHeight;
            // const previousMarginBottom = parseInt(window.getComputedStyle(previousToast.element).marginBottom); // 加上底边距好像有点远了

            const offset = previousHeight ;

            // 根据位置设置偏移
            if (this.options.position.startsWith('top')) {
                this.element.style.marginTop = `${offset}px`;
            } else {
                this.element.style.marginBottom = `${offset}px`;
            }
        }
    }

    close() {
        if (this.timer) {
            clearTimeout(this.timer);
        }

        // 关闭动画
        this.element.style.opacity = '0';
        this.element.style.transform = this.options.position.startsWith('top') ? 'translateY(-20px)' : 'translateY(20px)';

        // 动画结束后移除元素
        setTimeout(() => {
            if (this.element.parentNode) {
                this.element.parentNode.removeChild(this.element);
            }

            // 从容器中移除
            toastContainer.removeToast(this);

            // 调用回调
            if (this.options.onClose) {
                this.options.onClose();
            }
        }, 300);
    }
}

// 创建并显示toast的函数
function showToast(options) {
    return new Toast(options);
}

// 提供便捷方法
export const toast = {
    success(message, options = {}) {
        return showToast({ ...options, type: 'success', message });
    },
    warning(message, options = {}) {
        return showToast({ ...options, type: 'warning', message });
    },
    error(message, options = {}) {
        return showToast({ ...options, type: 'error', message });
    },
    show(options) {
        return showToast(options);
    },
};

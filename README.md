## 开发启动

 ```sh
 wails3 dev
 ```

## 构建项目
 ```sh
 wails3 build
 ```

# 注意事项
> 1.TMD wails 默认的把**cgo** 给关上了，害我找了半天 build/windows/Taskfile.yml 下的**CGO_ENABLED**  
> 2.html背景使用rgba全透明的时候无法使用**backdrop-filter**背景模糊  
> 3.flex 布局加 input type=number 时会导致input的点击范围隐形增大

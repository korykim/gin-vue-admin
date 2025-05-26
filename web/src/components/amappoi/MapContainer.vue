<script setup>
import { onMounted, onUnmounted, ref, defineComponent } from "vue";
import AMapLoader from "@amap/amap-jsapi-loader";

defineComponent({
    name: 'MapContainer'
})

let map = null;

// 用于存储搜索结果的POI点
const searchResults = ref([]);
// 存储当前点击位置的经纬度
const currentLngLat = ref({ lng: '', lat: '' });
// 复制成功提示
const copySuccess = ref(false);
// 已保存的POI数量
const savedPoisCount = ref(0);
// POI类型编码
const poiTypeCode = ref('071100');

onMounted(() => {
    window._AMapSecurityConfig = {
        securityJsCode: "",
    };
    AMapLoader.load({
        key: "", // 申请好的Web端开发者Key，首次调用 load 时必填
        version: "2.0", // 指定要加载的 JSAPI 的版本，缺省时默认为 1.4.15
        plugins: ["AMap.Scale", "AMap.PlaceSearch", "AMap.AutoComplete"], // 需要使用的的插件列表
    })
        .then((AMap) => {
            map = new AMap.Map("container", {
                viewMode: "3D", // 是否为3D地图模式
                zoom: 15, // 初始化地图级别
                center: [116.466485, 39.995197], // 初始化地图中心点位置
                resizeEnable: true,
            });

            // 初始化搜索框自动完成功能
            initAutoComplete(AMap);

            // 添加地图点击事件
            map.on('click', (e) => {
                getNearbyPois(e.lnglat, AMap);
            });
            
            // 加载已保存的POI数据
            loadSavedPois(AMap);
        })
        .catch((e) => {
            console.log(e);
        });
});

// 初始化搜索框自动完成功能
function initAutoComplete(AMap) {
    const autoOptions = {
        input: "tipinput"
    };

    const auto = new AMap.AutoComplete(autoOptions);
    const placeSearch = new AMap.PlaceSearch({
        map: map,
        pageSize: 25,
        extensions: 'all',
        panel: "panel",
        type: poiTypeCode.value // 使用响应式变量
    });

    // 注册监听，当选中某条记录时会触发
    auto.on("select", (e) => {
        document.getElementById('panel').style.display = 'block';
        placeSearch.setCity(e.poi.adcode);
        placeSearch.search(e.poi.name); // 关键字查询
    });
}

// 根据经纬度获取周边POI
function getNearbyPois(lnglat, AMap) {
    // 清除地图上已有的标记
    map.clearMap();

    // 格式化经纬度，保留6位小数
    const lng = lnglat.getLng().toFixed(6);
    const lat = lnglat.getLat().toFixed(6);

    // 更新当前经纬度
    currentLngLat.value = { lng, lat };

    // 在点击位置添加标记，显示经纬度信息
    new AMap.Marker({
        position: lnglat,
        map: map,
        content: `<div style="background:#fff;padding:2px 5px;border:1px solid #f00;">
      <div>点击位置</div>
      <div>经度: ${lng}</div>
      <div>纬度: ${lat}</div>
      <div>类型: ${poiTypeCode.value}</div>
    </div>`
    });

    // 使用PlaceSearch服务搜索周边POI
    const placeSearch = new AMap.PlaceSearch({
        pageSize: 25,
        pageIndex: 1,
        map: map,
        extensions: 'all',
        panel: "panel",
        keywords: '',
        type: poiTypeCode.value // 使用响应式变量
    });

    // 显示结果面板
    document.getElementById('panel').style.display = 'block';

    // 周边搜索，200米范围内
    placeSearch.searchNearBy('', lnglat, 200, (status, result) => {
        if (status === 'complete' && result.info === 'OK') {
            console.log('搜索结果:', result.poiList.pois);
            searchResults.value = result.poiList.pois;

            // 在地图上标记所有POI点
            result.poiList.pois.forEach((poi) => {
                new AMap.Marker({
                    position: poi.location,
                    map: map,
                    content: `<div class="poi-marker">${poi.name}</div>`
                });
            });
            
            // 保存POI信息到localStorage，确保根据id去重
            savePoisToLocalStorage(result.poiList.pois);
        } else {
            console.error('POI搜索失败:', result);
        }
    });
}

// 加载已保存的POI数据
function loadSavedPois(AMap) {
    try {
        const savedPoisStr = localStorage.getItem('amapPois');
        if (savedPoisStr) {
            const savedPois = JSON.parse(savedPoisStr);
            console.log(`从localStorage加载了${savedPois.length}个POI点`);
            
            // 将已保存的POI数据添加到searchResults
            searchResults.value = savedPois;
            
            // 更新保存的POI计数
            savedPoisCount.value = savedPois.length;
            
            // 可以选择是否在地图上显示这些点
            // 这里我们只显示最近的10个点，避免地图过于拥挤
            const recentPois = savedPois.slice(0, 10);
            
            recentPois.forEach(poi => {
                try {
                    // 处理不同格式的location数据
                    let position;
                    if (poi.location) {
                        // 如果location是数组 [lng, lat]
                        if (Array.isArray(poi.location)) {
                            position = new AMap.LngLat(poi.location[0], poi.location[1]);
                        } 
                        // 如果location是对象 {lng, lat}
                        else if (typeof poi.location === 'object' && poi.location.lng !== undefined && poi.location.lat !== undefined) {
                            position = new AMap.LngLat(poi.location.lng, poi.location.lat);
                        }
                        // 其他情况，跳过这个POI点
                        else {
                            console.warn('无效的location格式:', poi.location);
                            return;
                        }
                        
                        new AMap.Marker({
                            position: position,
                            map: map,
                            content: `<div class="poi-marker">${poi.name}</div>`
                        });
                    }
                } catch (err) {
                    console.warn(`标记POI点"${poi.name}"失败:`, err);
                }
            });
        }
    } catch (error) {
        console.error('加载已保存的POI数据失败:', error);
    }
}

// 保存POI信息到localStorage，确保根据id去重
function savePoisToLocalStorage(newPois) {
    // 从localStorage获取已存储的POI数据
    let savedPois = [];
    try {
        const savedPoisStr = localStorage.getItem('amapPois');
        if (savedPoisStr) {
            savedPois = JSON.parse(savedPoisStr);
        }
    } catch (error) {
        console.error('解析localStorage中的POI数据失败:', error);
        savedPois = [];
    }

    // 创建一个Map用于快速查找已有的POI（通过id）
    const poisMap = new Map();
    savedPois.forEach(poi => {
        poisMap.set(poi.id, poi);
    });

    // 合并新的POI，确保不重复，并统一location格式
    newPois.forEach(poi => {
        if (!poisMap.has(poi.id)) {
            // 确保location格式统一为数组形式 [lng, lat]
            if (poi.location && typeof poi.location === 'object' && !Array.isArray(poi.location)) {
                if (poi.location.lng !== undefined && poi.location.lat !== undefined) {
                    poi.location = [poi.location.lng, poi.location.lat];
                }
            }
            
            // 添加POI类型编码
            poi.poiType = poiTypeCode.value;

            // 添加POI类型编码
            poi.BigCategory=poi.type.split(";")[0]
            poi.MidCategory=poi.type.split(";")[1]
            poi.SubCategory=poi.type.split(";")[2]
            
            savedPois.push(poi);
            poisMap.set(poi.id, poi);
        }
    });

    // 保存回localStorage
    try {
        localStorage.setItem('amapPois', JSON.stringify(savedPois));
        // 更新保存的POI计数
        savedPoisCount.value = savedPois.length;
        console.log(`成功保存${savedPois.length}个POI点到localStorage，类型编码: ${poiTypeCode.value}`);
    } catch (error) {
        console.error('保存POI数据到localStorage失败:', error);
    }
}

// 复制经纬度到剪贴板
function copyLngLat() {
    const { lng, lat } = currentLngLat.value;
    if (lng && lat) {
        const text = `${lng},${lat}`;
        navigator.clipboard.writeText(text).then(() => {
            copySuccess.value = true;
            setTimeout(() => {
                copySuccess.value = false;
            }, 2000);
        }).catch(err => {
            console.error('复制失败: ', err);
        });
    }
}

// 清除已保存的POI数据
function clearSavedPois() {
    try {
        localStorage.removeItem('amapPois');
        searchResults.value = [];
        savedPoisCount.value = 0;
        if (map) {
            map.clearMap();
        }
        console.log('已清除所有保存的POI数据');
    } catch (error) {
        console.error('清除POI数据失败:', error);
    }
}

onUnmounted(() => {
    map?.destroy();
});
</script>

<template>
    <div class="w-full h-full relative bg-gray-50 dark:bg-slate-800">
        <div id="container" class="w-full h-full absolute top-0 left-0"></div>
        
        <!-- 搜索框 -->
        <div id="myPageTop" class="absolute top-3 right-3 bg-white dark:bg-slate-700 p-3 rounded-md shadow-md z-10 mt-35">
            <div class="flex flex-col">
                <label class="text-slate-700 dark:text-slate-300 mb-1">请输入关键字：</label>
                <input id="tipinput" class="w-48 h-8 px-2 border border-gray-300 dark:border-slate-600 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-slate-800 dark:text-slate-300" />
            </div>
            
            <div class="flex flex-col mt-2">
                <label class="text-slate-700 dark:text-slate-300 mb-1">POI类型编码：</label>
                <input v-model="poiTypeCode" class="w-48 h-8 px-2 border border-gray-300 dark:border-slate-600 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-slate-800 dark:text-slate-300" placeholder="例如：071100（美容美发店）" />
                <div class="text-xs text-gray-500 mt-1">默认：071100（美容美发店）</div>
            </div>
            
            <div class="flex justify-between mt-2">
                <button @click="clearSavedPois" class="text-xs text-red-500 hover:text-red-700">
                    清除历史记录
                </button>
                <span class="text-xs text-gray-500">已保存 {{ savedPoisCount }} 个位置</span>
            </div>
        </div>

        <!-- 经纬度信息显示区域 -->
        <div v-if="currentLngLat.lng && currentLngLat.lat" 
             class="absolute left-3 bottom-[210px] bg-white dark:bg-slate-700 p-3 rounded-md shadow-md z-10 min-w-[230px]">
            <div class="flex flex-col space-y-2">
                <div class="font-bold text-sm text-slate-700 dark:text-slate-300">位置坐标</div>
                <div class="font-mono text-sm bg-gray-100 dark:bg-slate-800 p-1 rounded border border-gray-200 dark:border-slate-600 text-slate-700 dark:text-slate-300">
                    {{ currentLngLat.lng }},{{ currentLngLat.lat }}
                </div>
                <button @click="copyLngLat" 
                        class="py-1.5 px-3 bg-blue-500 hover:bg-blue-600 transition-colors text-white rounded text-sm">
                    复制坐标
                </button>
                <span v-if="copySuccess" class="text-green-500 text-sm animate-fadeIn">复制成功!</span>
            </div>
        </div>

        <!-- 结果面板 -->
        <div id="panel" class="absolute bottom-0 left-0 w-full h-[200px] overflow-y-auto bg-white/90 dark:bg-slate-800/90 z-20 hidden"></div>
    </div>
</template>

<style scoped>
.poi-marker {
    @apply bg-white dark:bg-slate-700 py-0.5 px-1.5 border border-gray-300 dark:border-slate-600 rounded text-xs;
}

@keyframes fadeIn {
    from {
        opacity: 0;
    }
    to {
        opacity: 1;
    }
}

.animate-fadeIn {
    animation: fadeIn 0.3s;
}
</style>

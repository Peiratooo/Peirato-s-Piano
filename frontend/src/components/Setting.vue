<template>
    <div class="settings">
        <div class="setting-card">
            <div class="device-selector">
                <div class="selector in">
                    <div class="title">选择输入Midi设备</div>
                    <n-select size="small" v-model:value="store.devices.selectedInDevice"
                              :options="devices.inDevice" @update:value="changeDevice('in',$event)"
                    />
                </div>
                <div class="selector out">
                    <div class="title">选择输出Midi设备</div>
                    <n-select size="small" v-model:value="store.devices.selectedOutDevice"
                              :options="devices.outDevice" @update:value="changeDevice('out',$event)"
                    />

                </div>
                <div class="other-btns" style="margin-top: auto">
                    <div class="info" @click="store.showAuthor = true;store.settingMenu = false">
                        <n-icon size="24">
                            <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 16 16"><g fill="none"><path d="M8 2a6 6 0 1 1 0 12A6 6 0 0 1 8 2zm0 1a5 5 0 1 0 0 10A5 5 0 0 0 8 3zm0 7.5A.75.75 0 1 1 8 12a.75.75 0 0 1 0-1.5zm0-6a2 2 0 0 1 2 2c0 .73-.212 1.14-.754 1.708l-.264.27C8.605 8.87 8.5 9.082 8.5 9.5a.5.5 0 0 1-1 0c0-.73.212-1.14.754-1.708l.264-.27C8.895 7.13 9 6.918 9 6.5a1 1 0 0 0-2 0a.5.5 0 0 1-1 0a2 2 0 0 1 2-2z" fill="currentColor"></path></g></svg>
                        </n-icon>
                    </div>
                    <n-button type="info" size="small" @click="resetConfig">重置设置</n-button>
                </div>
            </div>
            <div class="color-settings">
                <div class="picker" v-for="(item,index) in store.config.colors" @click="changeColorIndex(index)">
                    <div class="label">
                        {{item.label}}
                    </div>
                    <n-color-picker
                        v-model:value="store.config.colors[index].color"
                        :modes="['hex']"
                        :show-alpha="false"
                        :show="false"
                    />
                </div>
            </div>
            <div class="example">
                <ExampleKeyboard/>
                <ExamplePedal class="example-pedal"/>
            </div>
        </div>
    </div>
    <n-drawer v-model:show="showColorPicker" placement="right" width="225px" class="color-picker" show-mask="transparent">
        <div class="color-picker" style="height: 100%;overflow: hidden">
            <Chrome
                :model-value="store.config.colors[colorIndex].color"
                :disable-alpha="true"
                v-if="colorIndex"
                @update:model-value="updateKeyColor"
            />
        </div>

    </n-drawer>

</template>

<script setup>
import {inject, onBeforeUnmount, onMounted, ref, watch} from "vue";
import {NSelect,NColorPicker,NButton,NDrawer,NIcon,NModal} from "naive-ui"
import ExampleKeyboard from "./ExampleKeyboard.vue";
import {Chrome} from "@ckpack/vue-color";
import ExamplePedal from "./ExamplePedal.vue";

const changeConfig = inject("changeConfig")
const resetConfig = inject("resetConfig")
const store = inject("store")
const changeDevice = inject("changeDevice")
const setKeyColor = inject("setKeyColor")
const colorIndex = ref("")
const showColorPicker = ref(false)
const devices = ref({
    inDevice:[

    ],
    outDevice:[

    ],
    loaded:false
})

function changeColorIndex(index) {
    index = index.replaceAll(" ","")
    console.log(index)
    if (index !== "") {
        colorIndex.value = index
        showColorPicker.value = true
    }

}


watch(showColorPicker,()=>{
    if (!showColorPicker.value) {
        colorIndex.value = ''
    }
})

function updateKeyColor(e) {
    store.config.colors[colorIndex.value].color = e.hex
    setKeyColor()
}

function constructDeviceList(){
    console.log(store.devices)
    for (let i in store.devices.outMidiPool) {
        devices.value.outDevice.push({
            label:store.devices.outMidiPool[i].name,
            value:store.devices.outMidiPool[i].value
        })
    }
    for (let i in store.devices.inMidiPool) {
        devices.value.inDevice.push({
            label:store.devices.inMidiPool[i].name,
            value:store.devices.inMidiPool[i].value
        })
    }
    devices.value.loaded = true
}

onMounted(()=>{
    constructDeviceList()
})

onBeforeUnmount(()=>{
    changeConfig()
})
</script>

<style lang="scss" scoped>
.settings {
    display: flex;
    gap: 32px;
    height: 100vh;
    justify-content: space-between;
    user-select: none;
    padding: 22px 32px;
    box-sizing: border-box;
}

.color-picker {
    height: 100vh;
    display: flex;

}

.setting-card {
    transition: 200ms;
    display: flex;
    gap: 32px;
    font-size: 12px;
    .device-selector {
        width: 256px;
        display: flex;
        flex-direction: column;
        gap: 18px;

        .selector {
            display: flex;
            flex-direction: column;
            gap: 6px;
        }
    }
    .color-settings {
        display: grid;

        grid-template-columns: repeat(2,1fr);
        gap: 12px;
        .picker {
            width: 96px;
            font-size: 12px;
            display: flex;
            flex-direction: column;
            gap: 6px;
            cursor: pointer;
        }
    }

}
.vc-chrome {
    box-sizing: border-box !important;
    border-radius: 8px !important;
    height: 100% !important;
}
.vc-chrome-body {
    box-sizing: border-box;
    padding: 6px !important;
}
@keyframes blurFadeIN {
    0% {
        opacity: 0;
        filter: blur(10px);

        transform: translateX(30%);

    }
    100% {
        opacity: 1;
        filter: blur(0px);
        transform: translateX(0%);

    }
}
.blurFadeIN {
    animation: blurFadeIN 0.3s ease;
    position: absolute;
}
@keyframes blurFadeOUT {
    0% {
        opacity: 1;
        filter: blur(0px);
        transform: translateX(0%);

    }
    100% {
        opacity: 0;
        filter: blur(10px);
        transform: translateX(30%);


    }
}
.blurFadeOUT {
    animation: blurFadeOUT 0.5s ease;
    position: absolute;
}
.example {
    position: relative;
    z-index: 1000;
}
.example-pedal {
    position: absolute;
    bottom: 12px;
    right: 8px;
    z-index: 9999;
}
.other-btns {
    display: flex;
    align-items: center;
    gap: 8px;
    .info {
        cursor: pointer;
        transition: 300ms;
        padding: 4px;
        border-radius: 12px;
        &:hover {
            color: #0050c2;
            background-color: #eaf1ff;
        }
    }
}
</style>
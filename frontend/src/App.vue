<template>
    <n-config-provider
        :date-locale="dateZhCN"
        :locale="zhCN"
        :style="{opacity: store.config.opacity/100}"
    >
        <n-modal-provider>
            <n-loading-bar-provider>
                <n-message-provider>
                    <div class="container">
                        <div class="drag-area">
                            <transition enter-active-class="blurFadeIN" leave-active-class="blurFadeOUT">
                                <div class="panel" v-if="store.showSetting"
                                     :style="{left: store.menuBar ? '0':'-553px'}">
                                    <div class="setting quit" @click="Keyboard.Quit()">
                                        <n-icon>
                                            <svg xmlns="http://www.w3.org/2000/svg"
                                                 xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 24 24">
                                                <g fill="none" stroke="currentColor" stroke-width="2"
                                                   stroke-linecap="round" stroke-linejoin="round">
                                                    <path d="M7 6a7.75 7.75 0 1 0 10 0"></path>
                                                    <path d="M12 4v8"></path>
                                                </g>
                                            </svg>
                                        </n-icon>
                                    </div>
                                    <div class="setting" @click="store.settingMenu=true">
                                        <n-icon>
                                            <svg xmlns="http://www.w3.org/2000/svg"
                                                 xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 20 20">
                                                <g fill="none">
                                                    <path
                                                        d="M1.91 7.383a8.491 8.491 0 0 1 1.78-3.08a.5.5 0 0 1 .54-.135l1.918.686a1 1 0 0 0 1.32-.762l.366-2.006a.5.5 0 0 1 .388-.4a8.532 8.532 0 0 1 3.555 0a.5.5 0 0 1 .387.4l.367 2.006a1 1 0 0 0 1.32.762l1.918-.686a.5.5 0 0 1 .54.136a8.491 8.491 0 0 1 1.78 3.079a.5.5 0 0 1-.152.535l-1.555 1.32a1 1 0 0 0 0 1.524l1.555 1.32a.5.5 0 0 1 .152.535a8.491 8.491 0 0 1-1.78 3.08a.5.5 0 0 1-.54.135l-1.918-.686a1 1 0 0 0-1.32.762l-.367 2.007a.5.5 0 0 1-.387.399a8.53 8.53 0 0 1-3.555 0a.5.5 0 0 1-.388-.4l-.365-2.006a1 1 0 0 0-1.32-.762l-1.919.686a.5.5 0 0 1-.54-.136a8.49 8.49 0 0 1-1.78-3.079a.5.5 0 0 1 .153-.535l1.554-1.32a1 1 0 0 0 0-1.524l-1.554-1.32a.5.5 0 0 1-.153-.535zM8 10a2 2 0 1 0 4 0a2 2 0 0 0-4 0z"
                                                        fill="currentColor"></path>
                                                </g>
                                            </svg>
                                        </n-icon>
                                    </div>
                                    <div class="slider volume">
                                        <n-icon>
                                            <svg xmlns="http://www.w3.org/2000/svg"
                                                 xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 512 512">
                                                <path
                                                    d="M264 416.19a23.92 23.92 0 0 1-14.21-4.69l-.66-.51l-91.46-75H88a24 24 0 0 1-24-24V200a24 24 0 0 1 24-24h69.65l91.46-75l.66-.51A24 24 0 0 1 288 119.83v272.34a24 24 0 0 1-24 24z"
                                                    fill="currentColor"></path>
                                                <path
                                                    d="M352 336a16 16 0 0 1-14.29-23.18c9.49-18.9 14.3-38 14.3-56.82c0-19.36-4.66-37.92-14.25-56.73a16 16 0 0 1 28.5-14.54C378.2 208.16 384 231.47 384 256c0 23.83-6 47.78-17.7 71.18A16 16 0 0 1 352 336z"
                                                    fill="currentColor"></path>
                                                <path
                                                    d="M400 384a16 16 0 0 1-13.87-24C405 327.05 416 299.45 416 256c0-44.12-10.94-71.52-29.83-103.95A16 16 0 0 1 413.83 136C434.92 172.16 448 204.88 448 256c0 50.36-13.06 83.24-34.12 120a16 16 0 0 1-13.88 8z"
                                                    fill="currentColor"></path>
                                            </svg>
                                        </n-icon>
                                        <n-slider :tooltip="false" v-model:value="store.config.volume" :min="0"
                                                  :max="127" @dragend="changeConfig"/>
                                    </div>
                                    <div class="slider velocity">
                                        <n-icon>
                                            <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 24 24"><path d="M19 3H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zm-5 11.5h.25V19h-4.5v-4.5H10c.55 0 1-.45 1-1V5h2v8.5c0 .55.45 1 1 1zM5 5h2v8.5c0 .55.45 1 1 1h.25V19H5V5zm14 14h-3.25v-4.5H16c.55 0 1-.45 1-1V5h2v14z" fill="currentColor"></path></svg>
                                        </n-icon>
                                        <n-slider :tooltip="false" v-model:value="store.config.velocity" :min="0"
                                                  :max="127" @dragend="changeConfig"/>
                                    </div>
                                    <div class="slider opacity">

                                        <n-icon>
                                            <svg t="1742717463560" class="icon" viewBox="0 0 1024 1024" version="1.1"
                                                 xmlns="http://www.w3.org/2000/svg" p-id="5111" width="512"
                                                 height="512">
                                                <path
                                                    d="M768 199.111111h-512a56.888889 56.888889 0 0 0-56.888889 56.888889v512a56.888889 56.888889 0 0 0 56.888889 56.888889h512a56.888889 56.888889 0 0 0 56.888889-56.888889v-512a56.888889 56.888889 0 0 0-56.888889-56.888889z m0 227.555556h-170.666667v170.666666h170.666667v170.666667h-170.666667v-170.666667h-170.666666v170.666667h-170.666667v-170.666667h170.666667v-170.666666h-170.666667v-170.666667h170.666667v170.666667h170.666666v-170.666667h170.666667z"
                                                    p-id="5112"></path>
                                            </svg>
                                        </n-icon>
                                        <n-slider :tooltip="false" v-model:value="store.config.opacity" :min="20"
                                                  :max="100" @dragend="changeConfig"/>
                                    </div>
                                    <div class="switcher" @click="store.menuBar = !store.menuBar"
                                         :style="{transform: store.menuBar ? 'rotate(0deg)':'rotate(180deg)'}">
                                        <n-icon depth="3">
                                            <svg xmlns="http://www.w3.org/2000/svg"
                                                 xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 24 24">
                                                <path
                                                    d="M16.88 2.88a1.25 1.25 0 0 0-1.77 0L6.7 11.29a.996.996 0 0 0 0 1.41l8.41 8.41c.49.49 1.28.49 1.77 0s.49-1.28 0-1.77L9.54 12l7.35-7.35c.48-.49.48-1.28-.01-1.77z"
                                                    fill="currentColor"></path>
                                            </svg>
                                        </n-icon>
                                    </div>
                                </div>
                            </transition>
                            <transition enter-active-class="blurFadeIN" leave-active-class="blurFadeOUT">
                                <div v-if="store.showSetting" class="keyboard-setting panel" :style="{right: store.keyboardMenu ? '0':'-510px'}">
                                    <div class="switcher" @click="store.keyboardMenu = !store.keyboardMenu" :style="{transform: store.keyboardMenu ? 'rotate(180deg)':'rotate(0deg)'}">
                                        <n-icon depth="3">
                                            <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 24 24"><path d="M16.88 2.88a1.25 1.25 0 0 0-1.77 0L6.7 11.29a.996.996 0 0 0 0 1.41l8.41 8.41c.49.49 1.28.49 1.77 0s.49-1.28 0-1.77L9.54 12l7.35-7.35c.48-.49.48-1.28-.01-1.77z" fill="currentColor"></path></svg>
                                        </n-icon>
                                    </div>
                                    <div class="key-count">
                                        <n-icon style="opacity: 0.6;">
                                            <svg t="1741731333941" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="9592" width="512" height="512"><path d="M493.714286 566.857143h-18.285715a18.285714 18.285714 0 0 1-18.285714-18.285714V128h-91.428571v420.571429a18.285714 18.285714 0 0 1-18.285715 18.285714h-18.285714v292.571428a36.571429 36.571429 0 0 0 36.571429 36.571429h91.428571a36.571429 36.571429 0 0 0 36.571429-36.571429V566.857143z m36.571428 0v292.571428a36.571429 36.571429 0 0 0 36.571429 36.571429h91.428571a36.571429 36.571429 0 0 0 36.571429-36.571429V566.857143h-18.285714a18.285714 18.285714 0 0 1-18.285715-18.285714V128h-91.428571v420.571429a18.285714 18.285714 0 0 1-18.285714 18.285714h-18.285715z m-237.714285 0h-18.285715a18.285714 18.285714 0 0 1-18.285714-18.285714V128H164.571429a36.571429 36.571429 0 0 0-36.571429 36.571429v694.857142a36.571429 36.571429 0 0 0 36.571429 36.571429h91.428571a36.571429 36.571429 0 0 0 36.571429-36.571429V566.857143zM768 128v420.571429a18.285714 18.285714 0 0 1-18.285714 18.285714h-18.285715v292.571428a36.571429 36.571429 0 0 0 36.571429 36.571429h91.428571a36.571429 36.571429 0 0 0 36.571429-36.571429V164.571429a36.571429 36.571429 0 0 0-36.571429-36.571429h-91.428571zM128 54.857143h768a73.142857 73.142857 0 0 1 73.142857 73.142857v768a73.142857 73.142857 0 0 1-73.142857 73.142857H128a73.142857 73.142857 0 0 1-73.142857-73.142857V128a73.142857 73.142857 0 0 1 73.142857-73.142857z" p-id="9593"></path></svg>
                                        </n-icon>
                                        <n-radio-group v-model:value="store.config.keyboardType" @update:value="changeKeyboardType" :disabled="!store.keyboardMenu">
                                            <n-radio-button
                                                v-for="(item,index) in store.keybordType"
                                                :key="index"
                                                :value="item.value"
                                                :label="item.label"
                                                :size="'small'"
                                            />
                                        </n-radio-group>
                                    </div>
                                    <div class="key-tips">
                                        <n-select
                                            v-model:value="store.config.keyLabel"
                                            :options="store.labelMap"
                                            :size="'small'"
                                            @update:value="changeConfig"
                                            :disabled="!store.keyboardMenu"
                                        />
                                    </div>
                                    <div class="show-pedal">
                                        <n-radio-group v-model:value="store.config.showPedal" @update:value="changeKeyboardType" :disabled="!store.keyboardMenu">
                                            <n-radio-button
                                                key="show"
                                                :value="true"
                                                label="显示踏板"
                                                :size="'small'"
                                            />
                                            <n-radio-button
                                                key="show"
                                                :value="false"
                                                label="隐藏踏板"
                                                :size="'small'"
                                            />
                                        </n-radio-group>
                                    </div>
                                </div>
                            </transition>
                        </div>
                        <ClassicKeyboard v-if="loaded"/>

                        <Pedal v-if="store.loaded && store.config.showPedal"
                               class="pedal-card"
                               @mousemove="dragPedal"
                               :style="{left:pedalDrag.x+'px',top:pedalDrag.y+'px'}"
                        />

                        <Chord v-if="store.loaded"
                               class="pedal-card"
                               @mousemove="dragChord"
                               :style="{left:chordPos.x+'px',top:chordPos.y+'px'}"

                        />

                        <n-drawer v-model:show="store.settingMenu" width="850" placement="left">
                            <div style="transform-origin: center;transition: 200ms;overflow: hidden"
                                 :style="{transform: `scale(${store.scale})`}">
                                <Setting/>
                            </div>
                        </n-drawer>
                        <n-modal v-model:show="store.showAuthor" transform-origin="center">
                            <Author/>
                        </n-modal>
                    </div>
                </n-message-provider>
            </n-loading-bar-provider>
        </n-modal-provider>
    </n-config-provider>
</template>

<script setup>
import {computed, h, onMounted, provide, ref, watch} from 'vue';
import {WML, Events} from "@wailsio/runtime";
import axios from "axios";
import {Midi} from '@tonejs/midi'
import {Keyboard} from "../bindings/main/service"
import {data} from "./store";
import ClassicKeyboard from "./components/ClassicKeyboard.vue";
import {
    dateZhCN,
    NConfigProvider,
    NLoadingBarProvider,
    NMessageProvider,
    NModalProvider,
    zhCN,
    NIcon,
    NSlider,
    NModal,
    NDrawer, NRadioGroup, NSelect, NRadioButton
} from "naive-ui";
import Pedal from "./components/Pedal.vue";
import Setting from "./components/Setting.vue";
import Author from "./components/Author.vue";
import Chord from "./components/Chord.vue";


const loaded = ref(false)

const request = axios.create({
    baseURL: "/",
    headers: {
        Authorization: localStorage.getItem('token'),
    },

});

const size = ref({
    wWidth:0,
    bWidth:0,
    gap:0,
    ratio:1.6
})


const pedalDrag = ref({
    x: window.innerWidth * 0.92,
    y: window.innerHeight * 0.7,
    xr: 0.92,
    yr: 0.7
})

const chordPos = ref({
    x: window.innerWidth * 0.02,
    y: window.innerHeight * 0.65,
    xr: 0.02,
    yr: 0.65
})

function dragPedal(e) {
    if (e.buttons === 1) {
        pedalDrag.value.x += e.movementX
        pedalDrag.value.y += e.movementY
        pedalDrag.value.xr = pedalDrag.value.x / window.innerWidth
        pedalDrag.value.yr = pedalDrag.value.y / window.innerHeight
    }
}

function dragChord(e) {
    if (e.buttons === 1) {
        chordPos.value.x += e.movementX
        chordPos.value.y += e.movementY
        chordPos.value.xr = chordPos.value.x / window.innerWidth
        chordPos.value.yr = chordPos.value.y / window.innerHeight
    }
}


async function initKeyboardConfig() {
    const skipKey = ["k","K","l","L",";","'","\\","|","`","~","[","{","]","}","p","P"]

    let keyboardData = await (await fetch('/keyboard_config.json')).json()
    let keyMappingData = await (await fetch('/case-1.json')).json()
    let chordsname = await (await fetch('/ChordNames.json')).json()
    for (let i in keyMappingData) {
       keyMappingData[i]-=12
    }
    store.keyMapping['case-1'] = keyMappingData
    let reverseMapping = {}
    for (let i in keyMappingData) {
        if (skipKey.includes(i)) {
            continue
        }
        if (!(keyMappingData[i] in reverseMapping)) {
            reverseMapping[keyMappingData[i]] =i
        }
    }
    for (let i of keyboardData) {
        if (i.note === 'C') {
            i['octave_key'] = i.note + i.octave
        } else {
            i['octave_key'] = ""
        }
        if (i.index in reverseMapping) {
            i['keyboard'] = reverseMapping[i.index]
        } else {
            i['keyboard'] = ""
        }
        if (i.pitch === 0) {
            i.pitch = null
        }
    }
    store.keyboardConfig = keyboardData
    store.chordsname = chordsname
    loaded.value = true
}

const store = data()


async function getMidiDevices() {
    await Keyboard.GetMidiDevices().then((data) => {
        for (let i in store.devices) {
            store.devices[i] = data[i]
        }

        store.loaded = true
        // console.log(store.devices)
    })

}

function startListen() {
    Keyboard.MidiListenerStart()
}

function stopListen() {
    Keyboard.MidiListenerStop()
}

function getConfig() {
    Keyboard.SendConfig().then((data) => {
        store.config = data
        setKeyColor()
    })
}

function changeConfig() {
    Keyboard.ReceiveConfig(store.config)
}

function darkenHexColor(hex, factor = 0.7) {
    hex = hex.replace('#', '');
    let r = parseInt(hex.slice(0, 2), 16);
    let g = parseInt(hex.slice(2, 4), 16);
    let b = parseInt(hex.slice(4, 6), 16);
    r = Math.round(r * factor);
    g = Math.round(g * factor);
    b = Math.round(b * factor);
    return `#${(1 << 24 | r << 16 | g << 8 | b).toString(16).slice(1)}`;
}

async function changeDevice(deviceType, device) {
    await Keyboard.MidiListenerStop()
    await Keyboard.ChangeDevice(deviceType, device).then((res) => {
        if (res) {
            getMidiDevices().then(() => {
                Keyboard.MidiListenerStart()
            })
        } else {
            alert("设备切换失败")
        }
    })
}

function checkStore() {
    console.log(store)
}

function setKeyColor() {
    for (let i in store.config.colors) {
        document.documentElement.style.setProperty('--' + i, store.config.colors[i].color);
        document.documentElement.style.setProperty('--' + i + '-o', darkenHexColor(store.config.colors[i].color, 0.1) + '66');
    }
}

function listenMouse() {
    document.addEventListener('mouseenter', () => {
        store.showSetting = true
    });
    document.addEventListener('mouseleave', () => {
        store.showSetting = false
    });
}

function resetConfig() {
    Keyboard.ResetConfig().then(res => {
        store.config = res
        setKeyColor()
    })
}

const pressedKeys = new Set(); // 记录已经触发的按键

function keyboardListener() {
    window.addEventListener("keydown", (e)=>{
        if (!pressedKeys.has(e.key) && e.key in store.keyMapping['case-1']) {
            pressedKeys.add(e.key)
            Keyboard.KeyboardPlay(store.keyMapping['case-1'][e.key])
            store.activeKey[store.keyMapping['case-1'][e.key]] = true
        }


    });

    window.addEventListener("keyup", (e)=>{
        if (e.key in store.keyMapping['case-1']) {
            pressedKeys.delete(e.key)
            Keyboard.KeyboardStop(store.keyMapping['case-1'][e.key])
            store.activeKey[store.keyMapping['case-1'][e.key]] = false
        }

    });
}

function getPressedKeys() {
    let res = []
    for (let i in store.activeKey) {
        const val = (i-9)%12
        if (store.activeKey[i] && !res.includes(val)) {
            res.push(val)
        }
    }
    return res.sort((a, b)=>{return a - b})
}

function getRootNote() {
    let activeNotes = []
    for (let i in store.activeKey) {
        if (store.activeKey[i] && !activeNotes.includes(i)) {
            activeNotes.push(i)
        }
    }
    return store.noteName[(activeNotes[0]-9)%12]
}

function changeKeyboardType() {
    resize()
    changeConfig()
}


onMounted(() => {
    store.activeNotes = computed(getPressedKeys)
    store.rootNote = computed(getRootNote)
    WML.Reload()
    listenMouse()
    initKeyboardConfig()
    Events.On("down", (data) => {
        store.activeKey[data.data[0].value] = true
    })
    Events.On("up", (data) => {
        store.activeKey[data.data[0].value] = false
    })
    Events.On("pedal", (data) => {
        if (store.devices.selectedInDevice === -1) {
            return
        }
        for (let i in data.data[0]) {
            store.devices.pedalStatus[store.devices.selectedInDevice][i] = data.data[0][i]
        }
    })
    Events.On("devices", (data) => {
        for (let i in store.devices) {
            store.devices[i] = data.data[0][i]
        }
    })
    getConfig()
    getMidiDevices()
    startListen()
    keyboardListener()
    window.addEventListener('resize', (e) => {
        pedalDrag.value.x = window.innerWidth * pedalDrag.value.xr
        pedalDrag.value.y = window.innerHeight * pedalDrag.value.yr
        chordPos.value.x = window.innerWidth * chordPos.value.xr
        chordPos.value.y = window.innerHeight * chordPos.value.yr
        if (window.innerHeight < 220) {
            store.scale = 0.8
        } else if (window.innerHeight < 280) {
            store.scale = 0.9
        } else {
            store.scale = 1
        }
    })
    store.menuBar=false
    store.keyboardMenu = false
    store.showSetting = false
    // loadMidi("/みちしるべ 「路标」.mid")
})

async function keyBoardPlayMidi(key, duration) {
    await Keyboard.KeyboardPlay(key)
    store.activeKey[key] = true
    await delay(duration)
    // await Keyboard.KeyboardStop(key)
    store.activeKey[key] = false
}

function delay(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

async function loadMidi(path) {
    const midiFile = await Midi.fromUrl(path)
    midiFile.tracks.forEach(async (track) => {
        track.notes.forEach(async (note) => {
            await delay(note.time * 1000)
            await keyBoardPlayMidi(note.midi, note.duration * 1000);
        })
    });
}

function resize() {
    let whitekey = 0
    for (let i of store.keyboardConfig.slice(store.keyboardRange[store.config.keyboardType][0],store.keyboardRange[store.config.keyboardType][1])) {
        if (i.color === 'white') {
            whitekey++
        }
    }
    if (window.innerWidth / window.innerHeight < 8) {
        size.value.ratio = 1.7
    } else {
        size.value.ratio = 1.6
    }
    size.value.wWidth = window.innerWidth / whitekey
    size.value.bWidth = size.value.wWidth / size.value.ratio
    size.value.gap = size.value.bWidth / size.value.ratio
    document.documentElement.style.setProperty('--black-key-width', size.value.bWidth+'px');
    document.documentElement.style.setProperty('--white-key-width', size.value.wWidth+'px');
    document.documentElement.style.setProperty('--white-key-offset', -size.value.bWidth+'px');
    document.documentElement.style.setProperty('--black-key-offset', -size.value.gap*0.7+'px');

}

provide("request", request);
provide("checkStore", checkStore)
provide('store', store)
provide('changeDevice', changeDevice)
provide("setKeyColor", setKeyColor)
provide("changeConfig", changeConfig)
provide("resetConfig", resetConfig)
provide("Keyboard",Keyboard)
provide("resize",resize)

</script>

<style lang="scss">
.container {
    position: relative;
    border-radius: 6px;
    overflow: hidden;

    .drag-area {
        position: absolute;
        left: 0;
        right: 0;
        height: 25vh;
        background-color: #1f1f1f00;
        z-index: 999;
        --wails-draggable: drag;
        cursor: move;
        transition: 500ms;
        display: flex;
    }
}

.panel {
    --wails-draggable: no-drag;
    height: 44px;
    align-items: center;
    font-size: 24px;
    gap: 16px;
    display: flex;
    position: relative;
    padding-left: 16px;
    padding-right: 8px;
    background-color: rgba(180, 180, 180, 0.29);
    border-radius: 0 0 16px 0;
    backdrop-filter: blur(8px);
    border: 1px solid #99999933;
    box-shadow: 0 0 8px #00000033;
    cursor: auto;
    transition: 200ms;

    .setting, .slider {
        display: flex;
        align-items: center;
        justify-content: center;
        cursor: pointer;
        opacity: 0.6;
        color: #333;
        transition: 500ms;

        &:hover {
            opacity: 1;
            color: #111;
        }
    }

    .slider {
        width: 138px;
        gap: 6px;
        font-size: 26px;
    }

    .setting {
        font-size: 24px;
    }

    .quit {
        background-color: rgba(194, 50, 50, 0.6);
        border-radius: 16px;
        color: rgba(255, 255, 255, 0.6);
        font-size: 18px;
        padding: 3px;

        &:hover {
            background-color: rgba(155, 58, 58, 0.8);
            color: #ff5252;
        }
    }

}

.pedal-card {
    position: absolute;
    z-index: 9999;
    user-select: none;
    cursor: grab;
}

.pedal-card:active {
    cursor: grabbing;
}

.n-modal-mask {
    background-color: rgba(0, 0, 0, 0.2);
    backdrop-filter: blur(8px) !important;
}

@keyframes blurFadeIN {
    0% {
        opacity: 0;


    }
    100% {
        opacity: 1;


    }
}

.blurFadeIN {
    animation: blurFadeIN 0.3s ease;
    position: absolute;
}

@keyframes blurFadeOUT {
    0% {
        opacity: 1;


    }
    100% {
        opacity: 0;


    }
}

.blurFadeOUT {
    animation: blurFadeOUT 0.3s ease;
    position: absolute;
}

.n-slider {
    --n-dot-height: 2px !important;
    --n-dot-width: 2px !important;
    --n-fill-color: #3058f8 !important;
    --n-fill-color-hover: #3063e7 !important;
    --n-font-size: 12px !important;
    --n-handle-size: 12px !important;
    --n-rail-height: 3px !important;
    --n-mark-font-size: 12px !important;
}

.n-select-menu {
    --n-height: 50vh !important;
    --n-option-font-size: 12px !important;
    --n-option-check-color: #1c38de !important;
    --n-option-text-color-active: #1c38de !important;
    --n-option-text-color-disabled: rgba(194, 194, 194, 1);
    --n-option-text-color-pressed: #0c387a !important;
    --n-loading-color: #1833a0 !important;
    --n-option-height: 30px !important;
}

.n-base-selection {
    --n-border-active: 1px solid #2951b0 !important;
    --n-border-focus: 1px solid #1e57e8 !important;
    --n-border-hover: 1px solid #2d68ff !important;
    --n-box-shadow-active: 0 0 0 2px rgba(28, 75, 224, 0.2) !important;
    --n-box-shadow-focus: 0 0 0 2px rgba(24, 67, 160, 0.2) !important;
    --n-caret-color: #1d60f1 !important;
    --n-loading-color: #2458dc !important;
    --n-color: rgba(255, 255, 255, 0.5) !important;
    --n-border: 1px solid rgba(224, 224, 230, 0.39) !important;
}

.n-modal {
    box-shadow: none;
    filter: drop-shadow(0 0 8px rgba(0, 0, 0, 0.2));
}

.vc-chrome-toggle-btn, #input__label__hex__828 {
    display: none;
}

.n-radio-group {
    --n-font-size: 12px !important;
    --n-bezier: cubic-bezier(.4, 0, .2, 1);
    --n-button-border-color: rgba(218, 218, 218, 0.25) !important;
    --n-button-border-color-active: #1d60f1 !important;
    --n-button-border-radius: 3px;
    --n-button-box-shadow: inset 0 0 0 1px #0000;
    --n-button-box-shadow-focus: inset 0 0 0 1px #1449bb, 0 0 0 2px rgba(24, 92, 160, 0.3) !important;
    --n-button-box-shadow-hover: inset 0 0 0 1px #0000;
    --n-button-color: rgba(197, 197, 197, 0.4) !important;
    --n-button-color-active: rgba(21, 88, 213, 0.71) !important;
    --n-button-text-color: rgb(51, 54, 57);
    --n-button-text-color-hover: #5275ff !important;
    --n-button-text-color-active: #ffffff !important;
    --n-height: 28px !important;
    --n-opacity-disabled: 0.5;

}

.n-radio__label {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
}

.n-radio-button {
    padding: 0 8px !important;
}

.switcher {
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 6px;
    border-radius: 6px;
    font-size: 16px;
    transform: rotate(180deg);
    transition: 300ms;
}

.switcher:hover {
    background-color: rgba(128, 128, 128, 0.1);
}
.keyboard-setting {
    position: absolute;
    right: 0;
    z-index: 999;
    display: flex;
    height: 44px;
    border-radius: 0 0 0 16px;
    padding-left: 8px;
}

.key-count {
    display: flex;
    align-items: center;
    gap: 8px;

}
.key-tips {
    width: 7rem;
}

</style>
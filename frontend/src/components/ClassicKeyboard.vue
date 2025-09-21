<template>
    <div class="classic-keyboard" @mousedown="mouse.down = true" @mouseup="mouse.down=false;" @mouseleave="mouse.down=false;mouse.keyIndex=-1">
        <div class="key" v-for="(item, index) in store.keyboardConfig.slice(store.keyboardRange[store.config.keyboardType][0],store.keyboardRange[store.config.keyboardType][1])"
             :class="[
                item.color,
                item.note,
                (mouse.down && mouse.keyIndex===item.index) || store.activeKey[item.index] ? keyColorMap[item.color]:'',
                ]"
             :key="index" @mouseenter="mouse.keyIndex=item.index" @mouseleave="store.activeKey[mouse.keyIndex] = false" @mouseup="store.activeKey[mouse.keyIndex] = false">
            <div class="label" v-if="store.config.keyLabel !== '' && item[store.config.keyLabel]">
                {{item[store.config.keyLabel]}}
            </div>
        </div>

    </div>
</template>

<script setup>
import {inject, onMounted,ref,onBeforeUnmount,watch} from "vue";
const store = inject("store")
const Keyboard = inject("Keyboard")
const keyColorMap = {
    'black':"b-active",
    'white':"w-active"
}

const resize = inject("resize")
const mouse = ref({
    down:false,
    keyIndex:-1,
})

watch(mouse.value,()=>{
    if (mouse.value.keyIndex === -1) {
        return
    }
    if (mouse.value.down) {
        store.activeKey[mouse.value.keyIndex] = true
        Keyboard.KeyboardPlay(mouse.value.keyIndex)
    } else {
        Keyboard.KeyboardStop(mouse.value.keyIndex)
    }
})

onMounted(()=>{
    resize()
    window.addEventListener("resize",resize)
})

onBeforeUnmount(() => {
    window.removeEventListener("resize",resize)
})
</script>

<style lang="scss" scoped>

.classic-keyboard {
    display: flex;
    flex-direction: row;
    height: 100vh;
    width: 100vw;

}

.key {
    position: relative;
    transition: 100ms;
    height: 100%;
    user-select: none;
}
.label {
    position: absolute;
    left: 0;
    right: 0;
    bottom: 16px;
    display: flex;
    justify-content: center;
    opacity: 0.5;
}
.black .label {
    opacity: 0.8;
}
.white {
    height:100%;
    width:var(--white-key-width);
    z-index:1;
    border-left:1px solid #bbb;
    border-bottom:1px solid #bbb;
    border-radius:0 0 5px 5px;
    box-shadow:-1px 0 0 rgba(255,255,255,0.5) inset,0 0 5px #ccc inset,0 0 3px rgba(0,0,0,0.1);
    background-color: #f6f6f6;
    .label {
        font-size: 18px;
    }
}
.w-active {
    border-left:1px solid var(--whiteKey-o);
    border-bottom:1px solid var(--whiteKey-o);
    box-shadow:3px 0 3px rgba(0,0,0,0.1) inset,-3px 0 8px rgba(0,0,0,0.1) inset,0 0 3px rgba(0,0,0,0.2);
    background-color: var(--whiteKey);
}
.black {
    height:63%;
    width: var(--black-key-width);
    transform: translateX(var(--black-key-offset));
    z-index:3;
    border-bottom:1px solid #000;

    border-left:1px solid #000;
    border-radius:0 0 3px 3px;
    box-shadow:-1px -1px 2px rgba(255,255,255,0.2) inset,0 -5px 2px 3px rgba(0,0,0,0.6) inset,0 2px 4px rgba(0,0,0,0.5);
    background-color:#333;
    .label {
        font-size: 13px;
        color: #eee;
    }
}

.b-active {
    box-shadow:-1px -5px 2px rgba(131, 131, 131, 0.2) inset,0 -10px 10px 5px rgba(0,0,0,0.6) inset,0 1px 2px rgba(0,0,0,0.5);
    background-color: var(--blackKey);

    border-bottom:1px solid var(--blackKey-o);

    border-left:1px solid var(--blackKey-o);
}


.A:first-child {
    margin: 0;
}
.B,.D,.E,.A,.G {
    margin-left: var(--white-key-offset);
}


</style>
<template>
    <transition enter-active-class="blurFadeIN" leave-active-class="blurFadeOUT">
        <div class="chord" v-if="Object.keys(chord).length > 0 && chord.type==='chord'">

            <div class="top">
                <div class="left">
                    {{ chord.chord }}
                </div>
                <div class="right">
                    <div class="symbols" v-if="chord.alternativeSymbols && chord.alternativeSymbols.length > 0">
                        <div class="symbol" style="white-space: nowrap"
                             v-for="(item,index) in chord.alternativeSymbols.slice(0,3)">
                            {{ item }}
                        </div>
                    </div>

                    <div class="synonym-notes" v-else>
                        <div class="notes">
                            {{ chord.notes?.join(" ") }}
                        </div>

                    </div>

                    <div class="cn">
                        {{ chord.chinese }}
                    </div>
                </div>
            </div>


            <div class="bottom">
                <div class="en">
                    {{ chord.name }}
                </div>
            </div>
        </div>
        <div class="note" v-else-if="Object.keys(chord).length > 0 && chord.type==='note'">
            {{chord.note}}
        </div>
    </transition>
</template>

<script setup>
import {computed, inject} from "vue";

const store = inject("store")

function getSuitableChord(noteList = []) {

    if (noteList.length === 1) {
        return noteList[0]
    }
    const chordNoteList = [store.rootNote, "C", "D", "E", "F", "G", "A", "B"]
    let rest = []
    for (let i of noteList) {
        if (i.length === 1) {
            for (let n of chordNoteList) {
                if (i.startsWith(n)) {
                    return i
                }
            }
        } else {
            rest.push(i)
        }
    }
    for (let i of rest) {
        for (let n of chordNoteList) {
            if (i.startsWith(n)) {
                return i
            }
        }
    }
}

const chord = computed(() => {
    if (store.activeNotes.length === 1) {
        return {
            type:"note",
            note:store.noteName[store.activeNotes[0]]
        }
    } else {
        const activeKeys = store.activeNotes.join(" ")
        if (activeKeys in store.chordsname) {
            return {type:"chord",...store.chordsname[activeKeys][getSuitableChord(Object.keys(store.chordsname[activeKeys] || {}))]}
        } else {
            return {}
        }
    }

})
</script>

<style lang="scss" scoped>

.chord,.note {
    background-color: #eeeeee55;
    border-radius: 4px;
    display: flex;
    flex-direction: column;
    padding: 8px;
    backdrop-filter: blur(4px);
    box-shadow: 0 0 6px 0 rgba(0, 0, 0, 0.25);
    gap: 6px;

    .top {
        display: flex;
        height: 80%;
        gap: 8px;

        .left {
            width: 100%;
            font-size: 36px;
        }

        .right {
            padding: 3px 0;
            justify-content: space-between;
            display: flex;
            align-items: flex-end;
            flex-direction: column;

        }
    }
}
.note {
    font-size: 32px;
    padding: 12px 16px;
    color: #333333;
}
.synonym-notes {
    font-size: 12px;
    display: flex;
    justify-content: space-between;
}

.synonym {
    color: #666;
}

.symbols {
    display: flex;
    font-size: 12px;
    gap: 8px;
    color: #333333;
}

.en {
    font-size: 10px;
}

.en, .cn {
    font-size: 12px;
    white-space: nowrap;
}
</style>
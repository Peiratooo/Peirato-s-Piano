import {defineStore} from 'pinia'

export const data = defineStore('data', {
    state: () => {
        return {
            keyboard: [],
            activeKey: {},
            activeNotes:[],
            rootNote:"",
            keyMapping:{},
            chordsname:{},
            loaded: false,
            devices: {
                inMidiPool: [],
                outMidiPool: [],
                selectedInDevice: -1,
                selectedOutDevice: -1,
                pedalStatus: {}
            },
            noteName:{
                0:"A",
                1:"Bb",
                2:"B",
                3:"C",
                4:"Db",
                5:"D",
                6:"Eb",
                7:"E",
                8:"F",
                9:"Gb",
                10:"G",
                11:"Ab"
            },
            labelMap: [
                {
                    label: "八度",
                    value: "octave_key",
                },
                {
                    label:"音符名",
                    value: "note"
                },
                {
                    label:"数字唱名法",
                    value: "pitch"
                },
                {
                    label:"音调唱名法",
                    value: "tone"
                },
                {
                    label:"键盘映射",
                    value: "keyboard"
                }
            ],
            scale:1,
            keyboardRange:[
                [0, 88],
                [3, 87],
                [6, 83],
                [8, 81],
                [27,88],
                [27,64],
            ],
            keybordType:[
                {
                    label:"88",
                    value: 0
                },
                {
                    label:"84",
                    value: 1
                },
                {
                    label:"76",
                    value: 2
                },
                {
                    label:"72",
                    value: 3
                },
                {
                    label:"61",
                    value: 4
                },
                {
                    label:"37",
                    value: 5
                },
            ],
            config: {},
            settingMenu: false,
            keyboardMenu:true,
            menuBar:true,
            showSetting:true,
            showAuthor:false
        }
    },
    getters: () => {
    },
    actions: {},
})

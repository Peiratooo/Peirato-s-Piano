<template>
    <div class="author-card">
        <div class="left">
            <div class="author">
                <div class="avatar">
                    <n-avatar src="/avatar.jpg" :size="72"/>
                </div>
                <div class="info">
                    <div class="name">Peirato</div>
                    <div class="social">
                        <div class="item" v-for="(item,index) in platforms" :key="index" @click="openPlatform(index)">
                            <n-image :preview-disabled="true" :src="item.icon" :width="30" :height="30"/>
                        </div>
                    </div>
                </div>
            </div>
            <div class="desc">
                <div class="text">
                    {{desc[0]}}
                </div>
                <div class="version" >
                    <n-tag :bordered="false" strong size="small">版本: {{ store.config.version }}</n-tag>
                </div>


            </div>
        </div>
        <div class="right">
            <div class="title">💰️支持作者</div>
            <div class="qrcodes">
                <div class="qrcode" v-for="(item,index) in donate">
                    <n-qr-code :padding="0"  style="box-sizing: content-box" :size="108" :value="item.url" :icon-src="item.icon" :icon-size="24"/>
                </div>
            </div>
        </div>
        <n-modal v-model:show="showSocial.show">
            <div class="platform">
                <n-qr-code :padding="0" style="box-sizing: content-box" :size="168" :value="platforms[showSocial.index].url" :icon-src="platforms[showSocial.index].icon" :icon-size="32" />
            </div>
        </n-modal>
    </div>
</template>

<script setup>
import {NImage,NQrCode,NAvatar,NTag,NModal} from "naive-ui";
import {inject, ref} from "vue";

const Keyboard = inject("Keyboard")

const store = inject("store")

const showSocial = ref({
    show:false,
    index:-1,
})

function openPlatform(index) {
    showSocial.value = {
        show:true,
        index:index,
    }
    if (platforms[index].method === "url") {
        Keyboard.OpenUrl(platforms[index].url)
    }

}

const platforms = [
    {
        name: "Bilibili",
        icon: "/bilibili.png",
        url: "https://space.bilibili.com/7277347",
        method: "url"
    },
    {
        name: "抖音",
        icon: "/tiktok.png",
        url: "https://www.douyin.com/user/MS4wLjABAAAAENe7s0M7uUpWwOqCXhoRiBD85CeYolcPcFltgjYW-hw",
        method: "url"
    },
    {
        name: "微信",
        icon: "/wechat.png",
        url: "https://u.wechat.com/MBX8JgLr7FANfm5l5RYByAg",
        method: "qrcode"
    },
    {
        name: "GitHub",
        icon: "/github.png",
        url: "https://github.com/Peiratooo",
        method: "url"
    },
]

const desc = [
    "一个简单、免费、无广告的钢琴键盘挂件"
]

const donate = [
    {
        name: "支付宝",
        icon: "/alipay.png",
        url: "https://qr.alipay.com/fkx17082m1wjekpkpwhif58",
        method: "qrcode"
    },
    {
        name: "微信",
        icon: "/wechat.png",
        url: "wxp://f2f0umI-AZys25kpxmHZYNmRW4SbywxsqwZmYYGVLWb_MfU",
        method: "qrcode"
    },
]

</script>

<style lang="scss" scoped>
.author-card {
    user-select: none;
    display: flex;
    background-color: #fff;
    padding: 22px;
    border-radius: 6px;
    box-sizing: border-box;
    align-items: center;
    gap: 58px;
    .left,.right {
        display: flex;
        flex-direction: column;
    }
    .author {
        display: flex;
        align-items: center;
        gap: 16px;
    }
    .info {
        display: flex;
        flex-direction: column;
        gap: 8px;
        .name {
            font-size: 32px;
            font-weight: bold;
            user-select: text;
        }
        .social {
            display: flex;
            gap: 8px;
            .item {
                display: flex;

                cursor: pointer;
                border-radius: 4px;
                overflow: hidden;
            }
        }
    }
    .social,.qrcodes {
        display: flex;
    }
    .desc {
        margin-top:32px;
        display: flex;
        flex-direction: column;
        gap: 8px;
    }
    .qrcodes {
        display: flex;
        gap: 16px;
    }
    .title {
        margin-bottom: 16px;
    }
}
.platform {
    background-color: #f6f6f6;
    border-radius: 8px;
    padding: 16px;
}
</style>
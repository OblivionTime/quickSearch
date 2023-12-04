<template>
    <nav class="more-mune contextmenu" :style="position" v-show="visible1">
        <ul class="more-mune-items">
            <li class="more-mune-item" v-for="item in RightMenu" :key="item.name" @click="item.action">
                <img :src="item.icon" alt="">
                {{ item.name }}
            </li>

        </ul>
    </nav>
</template>

<script setup>
const props = defineProps({
    RightMenu: {
        type: Array,
        default: () => [],
    },

});
/**
 * 右键菜单
 */
const position = ref("");
let visible1 = ref(false);
let active_item = ref(-1);
let rightClickDetector = ref("display:none");
const openMenu = (e) => {
    rightClickDetector.value = "display:none";
    setTimeout(() => {
        visible1.value = true;
        let current = document.elementFromPoint(e.clientX, e.clientY);
        let parentWithClass = current.closest(".item");
        if (!parentWithClass) {
            closeMenu();
            return;
        }
        let index = parentWithClass.getAttribute("index");
        active_item.value = +index;
        rightClickDetector.value = "display:block";
        if (window.innerHeight - e.pageY > 150) {
            position.value = `left: ${e.pageX}px;top: ${e.pageY}px`;
        } else {
            position.value = `left: ${e.pageX}px; bottom:${window.innerHeight - e.pageY - 20
                }px`;
        }
    }, 50);
};
const closeMenu = () => {
    rightClickDetector.value = "display:none";
    active_item.value = -1;
    visible1.value = false;
};
watch(visible1, () => {
    if (visible1.value) {
        document.body.addEventListener("click", closeMenu);
    } else {
        document.body.removeEventListener("click", closeMenu);
    }
});
//暴露子组件方法、属性
defineExpose({
    openMenu,
    closeMenu,
    active_item
});
</script>

<style lang="scss" scoped>
.more-mune {
    .more-mune-items {
        display: flex;
        flex-direction: column;
        gap: 10px;

        .more-mune-item {
            box-sizing: border-box;
            padding: 10px;
            padding-left: 20px;
            border-radius: 10px;
            display: flex;
            gap: 10px;
            width: 160px;

            align-items: center;
            font-size: 14px;
            cursor: pointer;

            img {
                width: 20px;
            }

            &:hover {
                background: #eee;
            }
        }
    }
}

.contextmenu {
    margin: 0;
    background: #fff;
    position: absolute;
    list-style-type: none;
    padding: 10px 0;
    border-radius: 4px;
    font-size: 16px;
    font-weight: 400;
    color: #333;
    box-shadow: 2px 2px 3px 2px rgba(0, 0, 0, 0.3);
    z-index: 1001;

    .item {
        padding: 0 15px;
        height: 35px;
        width: 100%;
        line-height: 35px;
        color: rgb(29, 33, 41);
        cursor: pointer;
    }

    .item:hover {
        background: rgb(229, 230, 235);
    }
}
</style>
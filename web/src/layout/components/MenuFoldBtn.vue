<template>
  <div class="gi_hover_btn menu-fold-btn" @click="onClick">
    <icon-menu-fold v-if="!appStore.menuCollapse" :size="20" :stroke-width="3" />
    <icon-menu-unfold v-else :size="20" :stroke-width="3" />
  </div>

  <div class="drawer" :class="{ 'app-menu-dark': appStore.menuDark }"
    :style="appStore.menuDark ? appStore.themeCSSVar : undefined">
    <a-drawer v-model:visible="visible" placement="left" :header="false" :footer="false" :render-to-body="false"
      :drawer-style="{
    'box-sizing': 'border-box',
    'background-color': 'var(--color-bg-1)',
  }">
      <Logo :collapsed="false"></Logo>
      <Menu class="menu w-full" @menu-item-click-after="visible = false"></Menu>
    </a-drawer>
  </div>
</template>

<script lang="ts" setup>
import { useAppStore } from '@/stores'
import Logo from '@/layout/components/Logo.vue'
import Menu from '@/layout/components/Menu/index.vue'
import { useDevice } from '@/hooks'

defineOptions({ name: 'MenuFoldBtn' })
const appStore = useAppStore()
const { isDesktop } = useDevice()
const visible = ref(false)

const onClick = () => {
  if (isDesktop.value) {
    appStore.setMenuCollapse(!appStore.menuCollapse)
  } else {
    visible.value = !visible.value
  }
}
</script>

<style lang="scss" scoped>
.menu-fold-btn {
  flex-shrink: 0;
  border-radius: 4px;
  width: 30px;
  height: 30px;
  display: flex;
  // 垂直剧中
  align-items: center;
  // 水平剧中
  justify-content: center;
  cursor: pointer;
  transition: all 0.3s;
  margin-right: 10px;
}

.drawer {
  .menu {
    flex: 1;
    overflow: hidden;
    background-color: inherit;
  }

  :deep(.arco-drawer-body) {
    padding: 0;
    overflow: hidden;
    display: flex;
    flex-direction: column;
  }
}
</style>

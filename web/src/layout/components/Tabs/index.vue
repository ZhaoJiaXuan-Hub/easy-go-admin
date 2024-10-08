<template>
  <div v-if="appStore.tab" class="tabs">
    <a-tabs :class="{ 'tabs-type-custom': appStore.tabMode === 'custom' }" editable hide-content size="medium"
      :type="calcTabsType" :active-key="route.path" @tab-click="(key) => handleTabClick(key as string)"
      @delete="tabsStore.closeCurrent">
      <a-tab-pane v-for="item of tabsStore.tagList" :key="item.path" :title="(item.meta?.title as string)"
        :closable="Boolean(!item.meta?.affix)">
      </a-tab-pane>
      <template #extra>
        <a-space size="medium">
          <ReloadIcon></ReloadIcon>
          <a-dropdown trigger="hover">
            <MagicIcon class="gi_mr"></MagicIcon>
            <template #content>
              <a-doption @click="tabsStore.closeCurrent(route.path)">
                <template #icon><icon-close /></template>
                <template #default>关闭当前</template>
              </a-doption>
              <a-doption @click="tabsStore.closeOther(route.path)">
                <template #icon><icon-eraser /></template>
                <template #default>关闭其他</template>
              </a-doption>
              <a-doption @click="tabsStore.closeAll">
                <template #icon><icon-minus /></template>
                <template #default>关闭全部</template>
              </a-doption>
            </template>
          </a-dropdown>
        </a-space>
      </template>
    </a-tabs>
  </div>
</template>

<script setup lang="ts">
import type { RouteRecordRaw } from 'vue-router'
import MagicIcon from './MagicIcon.vue'
import ReloadIcon from './ReloadIcon.vue'
import { useAppStore, useTabsStore } from '@/stores'

defineOptions({ name: 'Tabs' })
const route = useRoute()
const router = useRouter()
const appStore = useAppStore()
const tabsStore = useTabsStore()
const calcTabsType = computed(() => appStore.tabMode === 'custom' ? 'card' : appStore.tabMode)

// 重置, 同时把 affix: true 的路由筛选出来
tabsStore.reset()

// 监听路由变化
watch(
  () => route.path,
  () => {
    handleRouteChange()
  }
)

// 路由发生改变触发
function handleRouteChange() {
  const item = { ...route } as unknown as RouteRecordRaw
  tabsStore.addTagItem(toRaw(item))
  tabsStore.addCacheItem(toRaw(item))
  // console.log('路由对象', toRaw(item))
  // console.log('tagList', toRaw(tabsStore.tagList))
  // console.log('cacheList', toRaw(tabsStore.cacheList))
}
handleRouteChange()

// 点击页签
const handleTabClick = (key: string) => {
  router.push({ path: key })
}
</script>

<style lang="scss" scoped>
:deep(.arco-tabs-nav-tab) {
  .arco-tabs-tab {
    border-bottom-color: transparent !important;

    svg {
      width: 0;
      transition: all 0.15s;
    }

    &:hover {
      svg {
        width: 1em;
      }
    }
  }

  &:not(.arco-tabs-nav-tab-scroll) {
    .arco-tabs-tab:first-child {
      border-left: 0;
    }
  }
}

// 自定义tabs样式
:deep(.tabs-type-custom) {
  .arco-tabs-nav-tab {
    .arco-tabs-tab {
      padding: 5px 10px;
      border-radius: 7px;
      margin-top: 5px;
      margin-left: 10px;
      margin-bottom: 5px;
      border: none;
    }
  }

  .arco-tabs-nav-tab-list {
    .arco-tabs-tab-active {
      background-color: rgba(var(--primary-6), 0.1);
      padding: 5px 10px;

      &:hover {
        padding: 5px 10px;
        background-color: rgba(var(--primary-6), 0.1);
      }
    }
  }
}

:deep(.arco-dropdown-option-icon) {
  color: var(--color-text-3);
}

.tabs {
  background-color: var(--color-bg-1);
  box-shadow: 1px 0 2px 0 rgba(0, 0, 0, 0.08);
}

</style>

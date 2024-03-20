<template>
  <a-layout class="main">
    <a-layout>
      <a-layout-sider class="sider">
        <h1>DataX Tool</h1>
        <br />
        <a-menu
          :style="{ width: '100%' }"
          :default-selected-keys="['home']"
          :selected-keys="selected"
          @menu-item-click="handleSelect"
        >
          <router-link to="/home"
            ><a-menu-item key="home" data-obj="1"
              >首页</a-menu-item
            ></router-link
          >
          <router-link to="/datasource"
            ><a-menu-item key="datasource" data-obj="1"
              >数据源配置</a-menu-item
            ></router-link
          >
          <router-link to="/joblist"
            ><a-menu-item key="joblist" data-obj="1"
              >Job 配置列表</a-menu-item
            ></router-link
          >
          <router-link to="/setting"
            ><a-menu-item key="setting" data-obj="1"
              >系统设置</a-menu-item
            ></router-link
          >
          <router-link to="/about"
            ><a-menu-item key="about" data-obj="1"
              >关于</a-menu-item
            ></router-link
          >
        </a-menu>
        <br />
        <a-layout-footer><a-tag>©2024 DataX Tool</a-tag></a-layout-footer>
      </a-layout-sider>
      <a-layout-content><router-view></router-view></a-layout-content>
    </a-layout>
  </a-layout>
</template>
<script>
import { ref } from "vue";

export default {
  setup() {
    const selected = ref();
    return {
      selected,
    };
  },
  watch: {
    $route(to, from) {
      this.selected = [to.path.replace("/", "")];
    },
  },
  methods: {
    handleSelect(selectedMenu) {
      // 处理菜单项点击事件
      this.selected = [selectedMenu];
    },
  },
};
</script>
<style>
.main {
  height: 100%;
}
.arco-layout-content {
  height: 100%;
  overflow-y: auto;
}
.n-layout-sider {
  background-color: #f4f4f4;
}
.n-layout-sider-scroll-container {
  padding: 0 !important;
}
</style>

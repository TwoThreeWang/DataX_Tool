
<template>
  <div :style="{ background: 'var(--color-fill-2)' }">
    <a-page-header
      :style="{ background: 'var(--color-bg-2)' }"
      title="首页"
      subtitle="DataX 可视化配置工具"
      :show-back="false"
    >
    </a-page-header>
  </div>
  <div class="content">
    <a-card title="DataX Tool 怎么用？">
      <template #extra>
        <router-link to="/about"><a-link>关于</a-link></router-link>
      </template>
      <div class="frame-bg">
        <div class="frame-body">
          <div class="frame-aside">
            <a-steps :current="current" direction="vertical">
              <a-step>系统设置 </a-step>
              <a-step>数据源管理</a-step>
              <a-step>Job配置管理</a-step>
              <a-step>运行DataX</a-step>
            </a-steps>
          </div>
          <div class="frame-main">
            <div class="main-content" v-if="current == 1">
              <p>首先需要下载 DataX，并在系统设置中配置 DataX 的路径。</p>
              <p>
                DataX Github：<a-link
                  >https://github.com/alibaba/DataX/releases</a-link
                >
              </p>
              <p>
                官方 DataX 默认是基于 Python2，如果你本地是 Python3
                版本，可以根据以下教程将 DataX 改为 Python3 版本，<a-link
                  >https://github.com/TwoThreeWang/DataX_Python3</a-link
                >
              </p>
            </div>
            <div class="main-content" v-if="current == 2">
              <p>数据源管理相当于一个数据源链接的收藏夹，</p>
              <p>
                数据源在这里保存后，创建 Job
                配置时可以一键选择已经保存的数据源链接。
              </p>
            </div>
            <div class="main-content" v-if="current == 3">
              <p>
                Job 配置管理目录可以管理已有的 Job 配置文件，以及新增 Job 配置，
              </p>
              <p>同时也可以一键运行 DataX 任务。</p>
            </div>
            <div class="main-content" v-if="current == 4">
              <p>
                在 Job 配置管理目录下选择已有的 Job 配置，点击运行即可开始运行
                DataX 任务，
              </p>
              <p>运行中的任务可以点击查看运行日志以及停止任务</p>
            </div>
            <div class="main-bottom">
              <a-button :disabled="current === 1" @click="onPrev">
                <icon-left />
                Back
              </a-button>
              <a-button :disabled="current === 4" @click="onNext">
                Next
                <icon-right />
              </a-button>
            </div>
          </div>
        </div>
      </div>
    </a-card>
  </div>
</template>
<script>
import { ref } from "vue";

export default {
  setup() {
    const current = ref(1);
    const connect = ref("");

    const onPrev = () => {
      current.value = Math.max(1, current.value - 1);
    };

    const onNext = () => {
      current.value = Math.min(4, current.value + 1);
    };

    return {
      current,
      connect,
      onPrev,
      onNext,
    };
  },
};
</script>
<style scoped>
.content {
  padding: 25px;
  text-align: left;
}

.frame-body {
  display: flex;
  background: var(--color-bg-2);
}

.frame-aside {
  padding: 24px;
  border-right: 1px solid var(--color-border);
}

.frame-main {
  width: 100%;
}

.main-content {
  text-align: left;
  line-height: 20px;
  margin: 20px;
}

.main-bottom {
  margin-top: 40px;
  display: flex;
  justify-content: center;
}
</style>

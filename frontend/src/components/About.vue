<template>
  <div :style="{ background: 'var(--color-fill-2)' }">
    <a-page-header
      :style="{ background: 'var(--color-bg-2)' }"
      title="关于"
      subtitle="DataX Tool 可视化配置工具"
      :show-back="false"
    >
    </a-page-header>
  </div>
  <div class="content">
    <a-typography>
      <a-typography-title :heading="5">DataX Tool</a-typography-title>
      <a-typography-paragraph type="secondary">
        DataX Tool 是一个可视化配置管理工具，用于简化 DataX 的配置和运行过程。
      </a-typography-paragraph>
      <a-typography-title :heading="5">软件信息</a-typography-title>
      <a-typography-paragraph type="secondary">
        <a-tag>名称：</a-tag>DataX Tool<br />
        <a-tag>版本：</a-tag>1.0.2<br />
        <a-tag>上线日期：</a-tag>2024年03月15日<br />
        <a-tag>最新发布日期：</a-tag>2024年03月29日<br />
        <a-tag>反馈：</a-tag
        ><a
          href="https://github.com/TwoThreeWang/DataX_Tool/issues"
          target="_blank"
          ><a-link>https://github.com/TwoThreeWang/DataX_Tool/issues</a-link></a
        ><br />
      </a-typography-paragraph>
      <a-typography-title :heading="5">开发团队</a-typography-title>
      <a-typography-paragraph type="secondary" spacing="close">
        <a-space size="large">
          <a-avatar
            :size="64"
            shape="square"
            imageUrl="https://cdn.wangtwothree.com/imgur/RYc5CLI.jpeg"
          >
          </a-avatar>
        </a-space>
        <br />
        <a-tag>作者：</a-tag>王图思睿<br />
        <a-tag>邮箱：</a-tag>twothreewang@gmail.com<br />
        <a-tag>博客：</a-tag
        ><a target="_blank" href="https://wangtwothree.com"
          ><a-link>https://wangtwothree.com</a-link></a
        ><br />
        <a-tag>GitHub：</a-tag
        ><a target="_blank" href="https://github.com/TwoThreeWang"
          ><a-link>https://github.com/TwoThreeWang</a-link></a
        ><br />
      </a-typography-paragraph>
    </a-typography>
  </div>
</template>
<style>
.content {
  padding: 25px;
  text-align: left;
}
</style>
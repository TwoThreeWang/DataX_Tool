
<template>
  <a-page-header
    :style="{ background: 'var(--color-bg-2)' }"
    title="系统设置"
    subtitle="系统相关配置设置"
    :show-back="false"
  >
  </a-page-header>
  <div class="content">
    <a-form :model="form" auto-label-width @submit="save">
      <a-form-item field="DataxPath" label="datax 目录">
        <a-input v-model="form.DataxPath" placeholder="本地 datax 解压目录" />
      </a-form-item>
      <a-form-item>
        <a-button html-type="submit" type="outline">保存</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>
<script>
import { reactive } from "vue";
import { SaveSetting, GetSetting } from "../../wailsjs/go/main/App";

export default {
  setup() {
    const form = reactive({
      DataxPath: "",
    });

    return {
      form,
    };
  },
  methods: {
    init() {
      GetSetting().then((result) => {
        var res = JSON.parse(result);
        this.form.DataxPath = res.DataxPath;
      });
    },
    save() {
      SaveSetting(JSON.stringify(this.form)).then((result) => {
        if (result.includes("成功")) {
          this.$message.info(result);
        } else {
          this.$message.error(result);
        }
        this.init();
      });
    },
  },
  mounted() {
    this.init();
  },
};
</script>
<style scoped>
.content {
  padding: 25px;
  text-align: left;
}
</style>

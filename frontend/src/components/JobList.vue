
<template>
  <a-page-header
    :style="{ background: 'var(--color-bg-2)' }"
    title="job 配置列表"
    subtitle="配置管理及运行控制台"
    :show-back="false"
  >
    <template #extra>
      <a-button @click="handleClick('log/out.log')"> 查看日志 </a-button>
      <router-link to="/jobconfig"
        ><a-button type="primary">新增job配置</a-button></router-link
      >
    </template>
  </a-page-header>
  <div class="content">
    <a-list :max-height="600" :scrollbar="true">
      <template #header> Job 配置列表 </template>
      <a-list-item v-for="item in jobList">
        {{ item }}
        <template #actions v-if="nowRunFile == item">
          <a-button-group size="small" type="outline" status="success">
            <a-button> 运行中 </a-button>
            <a-button @click="handleClick('log/out.log')"> 查看日志 </a-button>
            <a-button @click="stop()"> 停止 </a-button>
          </a-button-group>
        </template>
        <template #actions v-else>
          <a-button-group size="small" type="outline">
            <a-button @click="run(item)"> 运行 </a-button>
            <a-button @click="handleClick(item)"> 编辑 </a-button>
            <a-button @click="del(item)"> 删除 </a-button>
          </a-button-group>
        </template>
      </a-list-item>
    </a-list>
    <a-modal
      v-model:visible="visible"
      @ok="handleOk"
      @cancel="handleCancel"
      fullscreen
      ok-text="保存"
    >
      <template #title> 查看内容【{{ handleOkFile }}】 </template>
      <a-textarea
        ref="textarea"
        v-model="fileConn"
        auto-size
        id="myTextarea"
        @input="scrollToBottom"
      />
    </a-modal>
  </div>
</template>
<script>
import { ref } from "vue";
import {
  DelJsonFile,
  GetJsonFileList,
  ReadFile,
  SaveFile,
  RunScript,
  KillScript,
  IsProcessRunning,
} from "../../wailsjs/go/main/App";

export default {
  setup() {
    const jobList = ref([]);
    const visible = ref(false);
    const fileConn = ref("");
    const handleOkFile = ref("");
    const handleCancel = () => {
      visible.value = false;
    };
    const nowPid = ref(0);
    const nowRunFile = ref("");
    const runLog = ref("");
    return {
      jobList,
      visible,
      fileConn,
      handleCancel,
      handleOkFile,
      nowPid,
      nowRunFile,
      runLog,
    };
  },
  methods: {
    init() {
      console.log("init");
      GetJsonFileList().then((result) => {
        this.jobList = result;
      });
    },
    del(fileName) {
      DelJsonFile(fileName).then((result) => {
        if (result.includes("报错")) {
          this.$message.error(result);
        } else {
          this.$message.info(result);
        }
        this.init();
      });
    },
    getFile(fileName) {
      ReadFile(fileName).then((result) => {
        if (result.includes("报错")) {
          this.$message.error(result);
          this.fileConn = result;
        } else {
          try {
            if (fileName != "log/out.log") {
              const obj = JSON.parse(result);
              this.fileConn = JSON.stringify(obj, null, 2);
            } else {
              this.fileConn = result;
            }
          } catch (e) {
            this.fileConn = result;
          }
        }
      });
    },
    handleClick(fileName) {
      this.handleOkFile = fileName;
      this.getFile(fileName);
      this.visible = true;
    },
    handleOk() {
      SaveFile(this.handleOkFile, this.fileConn).then((result) => {
        if (result.includes("报错")) {
          this.$message.error(result);
        } else {
          this.$message.info(result);
        }
      });
      this.visible = false;
    },
    run(fileName) {
      if (this.nowPid !== 0) {
        this.$message.error("当前有进程正在运行中！");
        return;
      }
      RunScript(fileName).then((result) => {
        var res = JSON.parse(result);
        this.nowPid = res.pid;
        if (res.info.includes("报错")) {
          this.$message.error(res.info);
        } else {
          this.$message.info(res.info);
          this.nowRunFile = fileName;
          this.poll();
        }
      });
    },
    stop() {
      KillScript(this.nowPid).then((result) => {
        if (result.includes("报错")) {
          this.$message.error(result);
        } else {
          this.$message.info(result);
          this.nowRunFile = "";
          this.nowPid = 0;
        }
      });
    },
    getLog() {
      KillScript(this.nowPid).then((result) => {
        if (result.includes("报错")) {
          this.$message.error(result);
        } else {
          this.$message.info(result);
          this.nowRunFile = "";
          this.nowPid = 0;
        }
      });
    },
    scrollToBottom() {
      var divElement = document.querySelector(".arco-modal-body");
      if (divElement) {
        divElement.scrollTop = divElement.scrollHeight;
      }
    },
    // 执行接口调用的函数
    async poll() {
      if (!this.nowPid) {
        console.log("退出循环");
        // 如果 flag 为 false，则停止轮询
        return;
      }
      this.getFile("log/out.log");
      this.scrollToBottom();
      IsProcessRunning(this.nowPid).then((result) => {
        if (!result) {
          console.log("进程已结束");
          this.nowRunFile = "";
          this.nowPid = 0;
          this.init();
        }
      });
      // 等待指定的时间间隔后再次调用 poll 函数
      setTimeout(this.poll, 1000);
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

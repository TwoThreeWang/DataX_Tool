<template>
  <a-page-header
    :style="{ background: 'var(--color-bg-2)' }"
    title="数据源配置"
    subtitle="数据源收藏夹"
    :show-back="false"
  >
    <template #extra>
      <a-button @click="handleClick" type="primary">新增数据源</a-button>
    </template>
  </a-page-header>
  <div class="content">
    <a-table :columns="columns" :data="data">
      <template #optional="{ record }">
        <a-button @click="editmodal(record)">编辑</a-button>
        <a-button @click="del(record.key)" status="danger">删除</a-button>
      </template>
    </a-table>
    <a-typography-text type="secondary">
      示例：<br />
      mysql数据源配置示例：
      <a-typography-text code>
        jdbc:mysql://127.0.0.1:3306/%database%?autoReconnect=true&failOverReadOnly=false&characterEncoding=utf-8&serverTimezone=GMT</a-typography-text
      ><br />
      配置mysql数据源时请将数据库名写为 “%database%”，配置 job
      时会自动替换为选择的数据库。
    </a-typography-text>
    <a-modal
      v-model:visible="visible"
      :title="modal_title"
      @cancel="handleCancel"
      @before-ok="save"
    >
      <a-form :model="modalData">
        <a-form-item field="alias" label="数据源别名">
          <a-input v-model="modalData.alias" />
        </a-form-item>
        <a-form-item field="db_type" label="数据源类型">
          <a-select v-model="modalData.db_type">
            <a-option value="txt">TXT</a-option>
            <a-option value="csv">CSV</a-option>
            <a-option value="mysql">Mysql</a-option>
            <a-option value="oracle">Oracle</a-option>
          </a-select>
        </a-form-item>
        <a-form-item field="link" label="数据源链接">
          <a-input v-model="modalData.link" />
        </a-form-item>
        <a-form-item field="username" label="用户名">
          <a-input v-model="modalData.username" />
        </a-form-item>
        <a-form-item field="password" label="密码">
          <a-input v-model="modalData.password" />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>
<script>
import { ref } from "vue";
import {
  GetDataSource,
  UpdateDataSource,
  SaveDataSource,
  DeleteDataSource,
} from "../../wailsjs/go/main/App";

export default {
  setup() {
    const show = ref(true);
    const modal_title = ref("新增数据源");

    const columns = [
      {
        title: "别名",
        dataIndex: "alias",
      },
      {
        title: "类型",
        dataIndex: "db_type",
      },
      {
        title: "链接",
        dataIndex: "link",
      },
      {
        title: "用户名",
        dataIndex: "username",
      },
      {
        title: "密码",
        dataIndex: "password",
      },
      {
        title: "操作",
        slotName: "optional",
      },
    ];
    const data = ref([]);
    const visible = ref(false);
    const modalData = ref({
      alias: "",
      db_type: "",
      link: "",
      username: "",
      password: "",
    });
    const handleClick = () => {
      modal_title.value = "新增数据源";
      modalData.value = {
        alias: "",
        db_type: "",
        link: "",
        username: "",
        password: "",
      };
      visible.value = true;
    };
    const handleCancel = () => {
      visible.value = false;
    };
    return {
      modal_title,
      columns,
      data,
      show,
      visible,
      modalData,
      handleClick,
      handleCancel,
    };
  },
  methods: {
    init() {
      GetDataSource().then((result) => {
        if (result.includes("错误")) {
          this.$message.error(result);
          return;
        }
        var res = JSON.parse(result);
        this.data = res;
      });
    },
    insertsave() {
      console.log(this.modalData);
      SaveDataSource(JSON.stringify(this.modalData)).then((result) => {
        if (result.includes("成功")) {
          this.$message.info(result);
        } else {
          this.$message.error(result);
        }
        this.init();
      });
    },
    editmodal(row) {
      this.modalData = row;
      this.modal_title = "编辑数据源";
      this.visible = true;
    },
    editsave() {
      console.log(this.modalData);
      UpdateDataSource(JSON.stringify(this.modalData)).then((result) => {
        if (result.includes("成功")) {
          this.$message.info(result);
        } else {
          this.$message.error(result);
        }
        this.init();
      });
    },
    save(done) {
      if (this.modal_title == "编辑数据源") {
        this.editsave();
      } else {
        this.insertsave();
      }
      this.visible = false;
      done();
    },
    del(key) {
      DeleteDataSource(key).then((result) => {
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

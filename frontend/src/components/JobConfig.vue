<template>
  <a-page-header
    :style="{ background: 'var(--color-bg-2)' }"
    title="Job作业配置"
    subtitle="新增配置"
    @back="back"
  >
  </a-page-header>
  <div class="content">
    <a-collapse
      default-active-key="1"
      accordion
      v-model:selected-keys="collapseKey"
      @change="changeCollapse"
    >
      <a-collapse-item header="一、配置数据源 Reader" key="1">
        <a-tabs v-model="readerType" @change="readerChange">
          <a-tab-pane key="txt" title="TXT">
            <!-- txt reader 配置 -->
            <a-form :model="txt_reader_form" auto-label-width>
              <a-form-item label="文件路径">
                <a-input
                  v-model="txt_reader_form.path"
                  placeholder="文件绝对路径，多个文件用英文逗号分隔"
                />
              </a-form-item>
              <a-form-item label="编码类型">
                <a-select
                  v-model="txt_reader_form.encoding"
                  placeholder="选择文件内容编码"
                >
                  <a-option value="UTF-8">UTF-8</a-option>
                  <a-option value="GBK">GBK</a-option>
                </a-select>
              </a-form-item>
              <a-form-item label="字段个数">
                <a-input
                  v-model="txt_reader_form.columnNum"
                  placeholder="文本文件有几个字段"
                />
              </a-form-item>
              <a-form-item label="字段分隔符">
                <a-input
                  v-model="txt_reader_form.fieldDelimiter"
                  placeholder="字段使用什么符号分隔"
                />
              </a-form-item>
              <a-form-item
                tooltip="填写文件路径和编码后点击预览，可以在这里预览文件第一行内容"
                label="预览"
              >
                <a-button
                  @click="getTxtFirstLine(txt_reader_form.encoding)"
                  :style="{ marginLeft: '10px' }"
                  >获取预览</a-button
                >
                <a-textarea
                  placeholder="点击预览文件第一行"
                  v-model="txtFirstLine"
                  auto-size
                />
              </a-form-item>
              <a-form-item>
                <a-checkbox v-model="txt_reader_form.skipHeader">
                  类 CSV
                  格式文件可能存在表头为标题情况，是否需要跳过。默认不跳过
                </a-checkbox>
              </a-form-item>
            </a-form>
          </a-tab-pane>
          <a-tab-pane key="mysql" title="Mysql">
            <!-- mysql reader 配置 -->
            <a-form :model="mysql_reader_form" auto-label-width>
              <a-form-item tooltip="选择已保存的数据源" label="选择数据源">
                <a-select
                  placeholder="选择一个数据源"
                  allow-clear
                  allow-search
                  @change="changeMysqlLink"
                >
                  <a-option v-for="(link, index) in dataSource" :value="index"
                    >{{ link.db_type }} - {{ link.alias }}</a-option
                  >
                </a-select>
              </a-form-item>
              <a-form-item tooltip="数据库的JDBC连接信息" label="jdbcUrl">
                <a-input
                  v-model="mysql_reader_form.connection.jdbcUrl"
                  placeholder="jdbcUrl"
                />
              </a-form-item>
              <a-form-item tooltip="数据源的用户名" label="username">
                <a-input
                  v-model="mysql_reader_form.username"
                  placeholder="数据库用户名"
                />
              </a-form-item>
              <a-form-item tooltip="数据源的用户密码" label="password">
                <a-input
                  type="password"
                  v-model="mysql_reader_form.password"
                  placeholder="数据库密码"
                />
              </a-form-item>
              <a-form-item
                tooltip="选取需要同步的表所在的库"
                label="选择数据库"
              >
                <a-select
                  v-model="mysql_reader_form.database"
                  placeholder="选择数据库"
                  allow-clear
                  allow-search
                  allow-create
                  @change="getMysqlTable"
                >
                  <a-option v-for="item in mysqlDbs" :value="item">{{
                    item
                  }}</a-option>
                </a-select>
              </a-form-item>
              <a-form-item tooltip="选取的需要同步的表" label="选择数据表">
                <a-select
                  v-model="mysql_reader_form.connection.table"
                  placeholder="选择数据表"
                  allow-clear
                  allow-search
                  allow-create
                  @change="getMysqlColumn"
                >
                  <a-option v-for="item in mysqlTables" :value="item">{{
                    item
                  }}</a-option>
                </a-select>
              </a-form-item>
              <a-form-item
                tooltip="所配置的表中需要同步的列名集合"
                label="选择字段"
              >
                <a-select
                  placeholder="选择字段 ..."
                  multiple
                  allow-clear
                  :default-value="mysql_reader_form.column"
                  :scrollbar="scrollbar"
                  @change="changeMysqlColumn"
                >
                  <a-option v-for="item in mysqlColumns">{{
                    item.value
                  }}</a-option>
                </a-select>
              </a-form-item>
              <a-form-item
                tooltip="如果指定splitPk，表示用户希望使用splitPk代表的字段进行数据分片，DataX因此会启动并发任务进行数据同步，这样可以大大提供数据同步的效能"
                label="splitPk"
              >
                <a-select
                  v-model="mysql_reader_form.splitPk"
                  placeholder="数据分片的字段，建议用主键（非必填）"
                  allow-clear
                  allow-search
                  allow-create
                >
                  <a-option v-for="item in mysqlColumns" :value="item.value">{{
                    item.value
                  }}</a-option>
                </a-select>
              </a-form-item>
              <a-form-item
                tooltip="where条件可以有效地进行业务增量同步。如果不填写where语句，包括不提供where的key或者value，DataX均视作同步全量数据"
                label="where"
              >
                <a-input
                  v-model="mysql_reader_form.where"
                  placeholder="where条件（非必填）"
                />
              </a-form-item>
              <a-form-item
                tooltip="当用户配置querySql时，MysqlReader直接忽略table、column、where条件的配置，querySql优先级大于table、column、where选项"
                label="querySql"
              >
                <a-input
                  v-model="mysql_reader_form.connection.querySql"
                  placeholder="自定义SQL（非必填）"
                />
              </a-form-item>
            </a-form>
          </a-tab-pane>
          <a-tab-pane key="oracle" title="Oracle">
            <!-- oracle reader 配置 -->
            开发中。
          </a-tab-pane>
        </a-tabs>
      </a-collapse-item>
      <a-collapse-item header="二、配置数据源 Writer" key="2">
        <a-tabs v-model="writerType" @change="writerChange">
          <a-tab-pane key="txt" title="TXT">
            <!-- txt writer 配置 -->
            <a-form :model="txt_writer_form" auto-label-width>
              <a-form-item field="path" label="文件路径">
                <a-input
                  v-model="txt_writer_form.path"
                  placeholder="文件保存路径"
                />
              </a-form-item>
              <a-form-item field="fileName" label="文件名称">
                <a-input
                  v-model="txt_writer_form.fileName"
                  placeholder="文件名称"
                />
              </a-form-item>
              <a-form-item field="writeMode" label="处理模式">
                <a-select
                  v-model="txt_writer_form.writeMode"
                  placeholder="处理模式"
                >
                  <a-option value="truncate">truncate</a-option>
                  <a-option value="append">append</a-option>
                  <a-option value="nonConflict">nonConflict</a-option>
                </a-select>
              </a-form-item>
              <a-form-item field="fieldDelimiter" label="字段分隔符">
                <a-input
                  v-model="txt_writer_form.fieldDelimiter"
                  placeholder="字段使用什么符号分隔"
                />
              </a-form-item>
              <a-form-item field="encoding" label="编码类型">
                <a-select
                  v-model="txt_writer_form.encoding"
                  placeholder="选择文件内容编码"
                >
                  <a-option value="UTF-8">UTF-8</a-option>
                  <a-option value="GBK">GBK</a-option>
                </a-select>
              </a-form-item>
              <a-form-item field="fileFormat" label="文件格式">
                <a-select
                  v-model="txt_writer_form.fileFormat"
                  placeholder="文件格式"
                >
                  <a-option value="csv">csv</a-option>
                  <a-option value="text">text</a-option>
                </a-select>
              </a-form-item>
            </a-form>
          </a-tab-pane>
          <a-tab-pane key="mysql" title="Mysql">
            <!-- mysql writer 配置 -->
            <a-form :model="mysql_writer_form" auto-label-width>
              <a-form-item tooltip="选择已保存的数据源" label="选择数据源">
                <a-select
                  placeholder="选择一个数据源"
                  allow-clear
                  allow-search
                  @change="changeMysqlLinkWriter"
                >
                  <a-option v-for="(item, index) in dataSource" :value="index"
                    >{{ item.db_type }} - {{ item.alias }}</a-option
                  >
                </a-select>
              </a-form-item>
              <a-form-item tooltip="数据库的JDBC连接信息" label="jdbcUrl">
                <a-input
                  v-model="mysql_writer_form.connection.jdbcUrl"
                  placeholder="jdbcUrl"
                />
              </a-form-item>
              <a-form-item tooltip="数据源的用户名" label="username">
                <a-input
                  v-model="mysql_writer_form.username"
                  placeholder="数据库用户名"
                />
              </a-form-item>
              <a-form-item tooltip="数据源的用户密码" label="password">
                <a-input
                  type="password"
                  v-model="mysql_writer_form.password"
                  placeholder="数据库密码"
                />
              </a-form-item>
              <a-form-item
                tooltip="选取需要同步的表所在的库"
                label="选择数据库"
              >
                <a-select
                  v-model="mysql_writer_form.database"
                  placeholder="选择数据库"
                  allow-clear
                  allow-search
                  allow-create
                  @change="getMysqlTable"
                >
                  <a-option v-for="item in mysqlDbs" :value="item">{{
                    item
                  }}</a-option>
                </a-select>
              </a-form-item>
              <a-form-item tooltip="选取的需要同步的表" label="选择数据表">
                <a-select
                  v-model="mysql_writer_form.connection.table"
                  placeholder="选择数据表"
                  allow-clear
                  allow-search
                  allow-create
                  @change="getMysqlColumn"
                >
                  <a-option v-for="item in mysqlTables" :value="item">{{
                    item
                  }}</a-option>
                </a-select>
              </a-form-item>
              <a-form-item
                tooltip="所配置的表中需要同步的列名集合"
                label="选择字段"
              >
                <a-select
                  placeholder="选择字段 ..."
                  multiple
                  allow-clear
                  :default-value="mysql_writer_form.column"
                  :scrollbar="scrollbar"
                  @change="changeMysqlColumnWriter"
                >
                  <a-option v-for="item in mysqlColumns">{{
                    item.value
                  }}</a-option>
                </a-select>
              </a-form-item>
              <a-form-item
                tooltip="DataX在获取Mysql连接时，执行session指定的SQL语句，修改当前connection session属性"
                label="session"
              >
                <a-input
                  v-model="mysql_writer_form.session"
                  placeholder="修改当前connection session属性（非必填）"
                />
              </a-form-item>
              <a-form-item
                tooltip="写入数据到目的表前，会先执行这里的标准语句"
                label="preSql"
              >
                <a-input
                  v-model="mysql_writer_form.preSql"
                  placeholder="写入前执行语句（非必填）"
                />
              </a-form-item>
              <a-form-item
                tooltip="写入数据到目的表后，会执行这里的标准语句"
                label="postSql"
              >
                <a-input
                  v-model="mysql_writer_form.postSql"
                  placeholder="写入后执行语句（非必填）"
                />
              </a-form-item>
              <a-form-item
                tooltip="控制写入数据到目标表采用 insert into 或者 replace into 或者 ON DUPLICATE KEY UPDATE 语句"
                label="writeMode"
              >
                <a-select
                  v-model="mysql_writer_form.writeMode"
                  placeholder="写入方式"
                >
                  <a-option value="insert">insert</a-option>
                  <a-option value="replace">replace</a-option>
                  <a-option value="update">update</a-option>
                </a-select>
              </a-form-item>
              <a-form-item
                tooltip="一次性批量提交的记录数大小，该值可以极大减少DataX与Mysql的网络交互次数，并提升整体吞吐量。但是该值设置过大可能会造成DataX运行进程OOM情况"
                label="batchSize"
              >
                <a-input
                  v-model="mysql_writer_form.batchSize"
                  placeholder="一次性批量提交的记录数大小（非必填）"
                />
              </a-form-item>
            </a-form>
          </a-tab-pane>
          <a-tab-pane key="oracle" title="Oracle">
            <!-- oracle writer 配置 -->
            开发中。
          </a-tab-pane>
        </a-tabs>
      </a-collapse-item>
      <a-collapse-item header="三、其他配置" key="3">
        <a-form-item field="jsonPath" label="文件名">
          <a-input
            v-model="jsonPath"
            placeholder="Json 配置文件保存的文件名，不需要加尾缀"
          />
        </a-form-item>
        <a-typography-paragraph
          >文件会自动添加 .json 尾缀，保存在 DataX 目录下的 job
          文件夹内，为空则不保存。</a-typography-paragraph
        >
      </a-collapse-item>
      <a-collapse-item header="四、Json 配置预览" key="4">
        <a-textarea
          placeholder="Json 配置预览（保存后查看）"
          v-model="jsonRes"
          auto-size
        />
      </a-collapse-item>
    </a-collapse>
    <br />
    <a-button type="primary" @click="save()">保存</a-button>
  </div>
</template>
<script>
import { ref } from "vue";
import {
  GetJobJson,
  GetDataSource,
  GetMysqlDb,
  GetMysqlTable,
  GetMysqlColumn,
  GetTxtFirstLine,
} from "../../wailsjs/go/main/App";

export default {
  setup() {
    const ORDER = "push";
    const dataSource = ref([]);
    const mysqlDbs = ref([]);
    const mysqlTables = ref([]);
    const mysqlColumns = ref([]);
    const txt_reader_form = ref({
      path: "",
      encoding: "UTF-8",
      columnNum: "2",
      fieldDelimiter: ",",
      skipHeader: false,
    });
    const txt_writer_form = ref({
      path: "",
      fileName: "",
      writeMode: "truncate",
      fileFormat: "csv",
      encoding: "UTF-8",
      fieldDelimiter: ",",
    });
    const mysql_jdbc_reader = ref("");
    const mysql_jdbc_writer = ref("");
    const mysql_reader_form = ref({
      username: "",
      password: "",
      database: "",
      column: [],
      connection: {
        table: "",
        jdbcUrl: "",
        querySql: "",
      },
      splitPk: "",
      where: "",
    });
    const mysql_writer_form = ref({
      writeMode: "insert",
      username: "",
      password: "",
      column: [],
      session: "",
      preSql: "",
      postSql: "",
      batchSize: "1024",
      connection: {
        table: "",
        jdbcUrl: "",
      },
    });
    const readerType = ref("txt");
    const writerType = ref("txt");
    const jsonRes = ref("");
    const collapseKey = ref(["1"]);
    const jsonPath = ref("");
    const txtFirstLine = ref("");
    return {
      ORDER,
      dataSource,
      mysqlDbs,
      mysqlTables,
      mysqlColumns,
      txt_reader_form,
      txt_writer_form,
      mysql_jdbc_reader,
      mysql_jdbc_writer,
      mysql_reader_form,
      mysql_writer_form,
      readerType,
      writerType,
      jsonRes,
      collapseKey,
      jsonPath,
      txtFirstLine,
    };
  },
  methods: {
    back() {
      this.$router.go(-1);
    },
    readerChange(activeTab) {
      this.readerType = activeTab;
    },
    writerChange(activeTab) {
      this.writerType = activeTab;
    },
    save() {
      try {
        var reader = "";
        if (this.readerType == "txt") {
          reader = this.txt_reader_form;
          reader.skipHeader = reader.skipHeader.toString();
        } else if (this.readerType == "mysql") {
          reader = this.mysql_reader_form;
          reader.connection.jdbcUrl = reader.connection.jdbcUrl.replace(
            "%database%",
            reader.database
          );
          reader.column = Object.values(reader.column);
          reader.connection.table = [reader.connection.table];
          reader.connection.jdbcUrl = [reader.connection.jdbcUrl];
          reader.connection.querySql == ""
            ? delete reader.connection.querySql
            : (reader.connection.querySql = [reader.connection.querySql]);
          reader.connection = [reader.connection];
          delete reader.database;
        } else {
          this.$message.error("不支持的reader类型");
          return;
        }
        var writer = "";
        if (this.writerType == "txt") {
          writer = this.txt_writer_form;
        } else if (this.writerType == "mysql") {
          writer = this.mysql_writer_form;
          writer.connection.jdbcUrl = writer.connection.jdbcUrl.replace(
            "%database%",
            writer.database
          );
          writer.column = Object.values(writer.column);
          writer.session == ""
            ? delete writer.session
            : (writer.session = [writer.session]);
          writer.preSql == ""
            ? delete writer.preSql
            : (writer.preSql = [writer.preSql]);
          writer.postSql == ""
            ? delete writer.postSql
            : (writer.postSql = [writer.postSql]);
          writer.connection.table = [writer.connection.table];
          writer.connection = [writer.connection];
          delete writer.database;
        } else {
          this.$message.error("不支持的writer类型");
          return;
        }
        GetJobJson(
          JSON.stringify(reader),
          this.readerType,
          JSON.stringify(writer),
          this.writerType,
          this.jsonPath
        ).then((result) => {
          if (result.includes("错误")) {
            this.$message.error(result);
            return;
          }
          const obj = JSON.parse(result);
          this.jsonRes = JSON.stringify(obj, null, 2);
          // this.jsonRes = result;
          this.collapseKey[0] = "4";
          this.$message.info("配置生成成功！");
        });
      } catch (e) {
        this.$message.error("保存错误");
        console.log(e);
      }
    },
    changeMysqlLink(index) {
      this.mysql_reader_form.connection.jdbcUrl =
        this.dataSource[index]["link"];
      this.mysql_reader_form.username = this.dataSource[index]["username"];
      this.mysql_reader_form.password = this.dataSource[index]["password"];
      this.mysql_jdbc_reader =
        this.mysql_reader_form.connection.jdbcUrl +
        "&username=" +
        this.mysql_reader_form.username +
        "&password=" +
        this.mysql_reader_form.password;
      this.getMysqlDb(this.mysql_jdbc_reader);
    },
    changeMysqlLinkWriter(index) {
      this.mysql_writer_form.connection.jdbcUrl =
        this.dataSource[index]["link"];
      this.mysql_writer_form.username = this.dataSource[index]["username"];
      this.mysql_writer_form.password = this.dataSource[index]["password"];
      this.mysql_jdbc_writer =
        this.mysql_writer_form.connection.jdbcUrl +
        "&username=" +
        this.mysql_writer_form.username +
        "&password=" +
        this.mysql_writer_form.password;
      this.getMysqlDb(this.mysql_jdbc_writer);
    },
    initDataSource() {
      GetDataSource().then((result) => {
        if (result.includes("错误")) {
          this.$message.error(result);
          return;
        }
        var res = JSON.parse(result);
        this.dataSource = res;
        console.log(this.dataSource);
      });
    },
    getMysqlDb(jdbcUrl) {
      GetMysqlDb(jdbcUrl).then((result) => {
        if (result.includes("错误")) {
          this.$message.error(result);
          return;
        }
        this.mysqlDbs = result.split(",");
        console.log(this.mysqlDbs);
      });
    },
    getMysqlTable(database) {
      var jdbcUrl = this.mysql_jdbc_reader;
      if (this.collapseKey[0] == 2) {
        jdbcUrl = this.mysql_jdbc_writer;
      }
      GetMysqlTable(jdbcUrl, database).then((result) => {
        if (result.includes("错误")) {
          this.$message.error(result);
          return;
        }
        this.mysqlTables = result.split(",");
        console.log(this.mysqlTables);
      });
    },
    getMysqlColumn(table) {
      var jdbcUrl = this.mysql_jdbc_reader;
      var database = this.mysql_reader_form.database;
      if (this.collapseKey[0] == 2) {
        jdbcUrl = this.mysql_jdbc_writer;
        database = this.mysql_writer_form.database;
      }
      GetMysqlColumn(jdbcUrl, database, table).then((result) => {
        if (result.includes("错误")) {
          this.$message.error(result);
          return;
        }
        var Column = result.split(",");
        console.log(Column);
        this.mysqlColumns = Column.map(function (item) {
          return {
            value: item,
            label: item,
          };
        });
        console.log(this.mysqlColumns);
      });
    },
    changeMysqlColumn(checkData) {
      this.mysql_reader_form.column = checkData;
      console.log(this.mysql_reader_form.column);
    },
    changeMysqlColumnWriter(checkData) {
      this.mysql_writer_form.column = checkData;
    },
    changeCollapse(key) {
      this.collapseKey = key;
      console.log(this.collapseKey);
    },
    getTxtFirstLine(encoding) {
      console.log(this.txt_reader_form.path, encoding);
      GetTxtFirstLine(this.txt_reader_form.path, encoding).then((result) => {
        if (result.includes("错误")) {
          this.$message.error(result);
          return;
        }
        this.txtFirstLine = result;
        console.log(result);
      });
    },
  },
  mounted() {
    this.initDataSource();
  },
};
</script>
<style scoped>
.content {
  padding: 25px;
  text-align: left;
}
</style>

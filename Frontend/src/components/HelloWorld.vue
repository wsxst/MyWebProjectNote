<template>
  <div class="hello">
    <el-form ref="form" :model="sizeForm" label-width="80px" size="mini">
      <el-form-item label="username">
        <el-input v-model="form.username"></el-input>
      </el-form-item>
      <el-form-item label="password">
        <el-input type="password" v-model="form.userpasswd"></el-input>
      </el-form-item>
      <el-form-item size="large">
        <el-button type="primary" @click="onSubmit">登陆</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
  import axios from "axios";
  const instance =axios.create({
    baseURL: "http://127.0.0.1:8081"
  });
  export default {
    name: 'HelloWorld',
    props: {
      msg: String
    },
    data() {
      return {
        form: {
          username: "",
          userpasswd: "",
        }
      };
    },
    methods: {
      onSubmit() {
          instance.post('/login', this.form)
                  .then((response) => {
                    alert(response.data.msg)
                  })
                  .catch ((error)=>{
                    alert(error);
                  });
      }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>

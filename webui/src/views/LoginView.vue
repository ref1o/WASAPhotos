<template>
  <div class="container-fluid h-100 m-0 p-0 login">

    <div class="row ">
      <div class="col">
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
      </div>
    </div>

    <div class="row h-100 w-100 m-0">

      <form @submit.prevent="login" class="d-flex flex-column align-items-center justify-content-center p-0 login-form mx-auto">

        <div class="row mt-2 mb-3">
          <div class="col">
            <h2 class="login-title">WASAPhotos</h2>
          </div>
        </div>

        <div class="row mt-2 mb-3">
          <div class="col">
            <input 
              type="text" 
              class="form-control" 
              v-model="identifier" 
              maxlength="16"
              minlength="3"
              placeholder="Your identifier" />
          </div>
        </div>

        <div class="row mt-2 mb-5 ">
          <div class="col ">
            <button class="btn btn-dark" :disabled="identifier == null || identifier.length >16 || identifier.length <3 || identifier.trim().length<3"> 
              Register/Login 
            </button>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  data: function() {
    return {
      errormsg: null,
      identifier: "",
      disabled: true,
    }
  },
  methods: {
    async login() {
      this.errormsg = null;
      try {
        // Login (POST): "/session"
        let response = await this.$axios.post("/session",{
          user_id: this.identifier.trim()
        });

        localStorage.setItem('token', response.data.user_id);
        this.$router.replace("/home");
        this.$emit('updatedLoggedChild', true);
      } catch (e) {
        this.errormsg = e.toString();
      }
    },
  },
  mounted(){
    if (localStorage.getItem('token')){
      this.$router.replace("/home");
    }
  },
}
</script>

// https://coolors.co/9bdde8-8392c0-e87386-f09a93-f7c19f

<style>
.login {
  background-image: url('../assets/images/background.jpg');
  background-size: cover;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}

.login-form {
  width: 100%;
  max-width: 400px;
}

.login-title {
  color: black;
}

.form-control {
	border-radius: 50px;
}

.btn-dark {
  border-radius: 75px; 
}
</style>
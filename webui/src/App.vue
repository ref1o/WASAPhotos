<template>
  <div class="app-container">
    <div class="row">
      <div class="col p-0">
        <main>
          <Navbar
            v-if="logged"
            @logoutNavbar="logout"
            @requestUpdateView="updateView"
            @searchNavbar="search"
          />

          <RouterView
            @updatedLoggedChild="updateLogged"
            @requestUpdateView="updateView"
            :searchValue="searchValue"
          />
        </main>
      </div>
    </div>
  </div>
</template>

<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<script>
export default {
  data() {
    return {
      logged: false,
      searchValue: "",
    };
  },
  methods: {
    logout(newValue) {
      this.logged = newValue;
      this.$router.replace("/login");
    },
    updateLogged(newLogged) {
      this.logged = newLogged;
    },
    updateView(newRoute) {
      this.$router.replace(newRoute);
    },
    search(queryParam) {
      this.searchValue = queryParam;
      this.$router.replace("/search");
    },
  },
  created() {
    if (!localStorage.getItem('notFirstStart')) {
      localStorage.clear();
      localStorage.setItem('notFirstStart', true);
    }
  },
  mounted() {
    if (!localStorage.getItem('token')) {
      this.$router.replace("/login");
    } else {
      this.logged = true;
    }
  },
};
</script>

<style>
.app-container {
  background-color: #f8f8f8; /* Colore di sfondo pulito */
  min-height: 100vh;
}
</style>

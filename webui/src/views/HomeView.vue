<script>
export default {
  data: function () {
    return {
      errormsg: null,
      photos: [],
    };
  },

  methods: {
    async loadStream() {
      try {
        this.errormsg = null;
        // Home get: "/users/:id/home"
        let response = await this.$axios.get(
          "/users/" + localStorage.getItem("token") + "/home"
        );

        if (response.data != null) {
          this.photos = response.data;
        }
      } catch (e) {
        this.errormsg = e.toString();
      }
    },
  },

  async mounted() {
    await this.loadStream();
  },
};
</script>

<template>
  <link href='https://fonts.googleapis.com/css?family=Courgette' rel='stylesheet'>
  <div class="container-fluid home-view">
    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

    <div class="row">
      <Photo
        v-for="(photo,index) in photos"
        :key="index"
        :owner="photo.owner"
        :photo_id="photo.photo_id"
        :comments="photo.comments != null ? photo.comments : []"
        :likes="photo.likes != null ? photo.likes : []"
        :upload_date="photo.date"
      />
    </div>

    <div v-if="photos.length === 0" class="row no-content-message">
	  <div class="vertical-dots">
        <div class="dot"></div>
        <div class="dot"></div>
        <div class="dot"></div>
      </div>

      <h1 class="d-flex justify-content-center mt-5">
        Explore the beauty captured by others and share your own moments!
      </h1>

	  <div class="vertical-dots">
        <div class="dot"></div>
        <div class="dot"></div>
        <div class="dot"></div>
      </div>

    </div>
  </div>
</template>

<style>
h1{
	margin-top: 0;
}

.home-view {
  background-color: #f8f8f8;
  padding: 20px;
}

.no-content-message {
  font-family: 'Courgette';
  color: #F09A93;
  text-align: center;
  padding: 80px;
}

.vertical-dots {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.dot {
  width: 10px;
  height: 10px;
  background-color: #F09A93;
  border-radius: 50%;
  margin: 10px 0;
}
</style>
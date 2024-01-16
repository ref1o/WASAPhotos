<script>
export default {
  data() {
    return {
      textVar: "",
      iconProfile: "fa-regular",
    };
  },
  methods: {
    logout() {
      localStorage.removeItem("token");
      this.$emit("logoutNavbar", false);
    },
    goBackHome() {
      this.$emit("requestUpdateView", "/home");
    },
    searchFunc() {
      this.$emit("searchNavbar", this.textVar);
      this.textVar = "";
    },
    myProfile() {
      this.$emit("requestUpdateView", "/users/" + localStorage.getItem("token"));
    },
    profileIconInactive() {
      this.iconProfile = "fa-regular";
    },
    profileIconActive() {
      this.iconProfile = "fa-solid";
    },
    async uploadFile() {
      let fileInput = document.getElementById("fileUploader");

      const file = fileInput.files[0];
      const reader = new FileReader();

      reader.readAsArrayBuffer(file);

      reader.onload = async () => {
        // Post photo: /users/:id/photos
        try {
          let response = await this.$axios.post(
            "/users/" + this.$route.params.id + "/photos",
            reader.result,
            {
              headers: {
                "Content-Type": file.type,
              },
            }
          );
          this.$emit("photoUploaded", response.data); // Emit an event to notify the parent component about the uploaded photo
        } catch (error) {
          console.error("Error uploading photo:", error);
          // Puoi gestire l'errore come preferisci, ad esempio mostrando un messaggio all'utente
        }
      };
    },

    // Aggiunto il metodo per emettere l'evento quando si clicca su "+"
    addPhoto() {
      document.getElementById("fileUploader").click();
    },
  },
};
</script>

<template>
  <nav class="navbar navbar-expand-lg bg-light d-flex justify-content-between sticky-top mb-3 my-nav">
    <div class="col-4">
      <a class="navbar-brand ms-2 d-flex" @click="goBackHome">
        <div class="brand-text">WASAPhoto</div>
      </a>
    </div>

    <div class="col-4">
      <form class="form-inline my-2 my-lg-0 d-flex justify-content-center m-auto">
        <input
          class="form-control mr-sm-2 w-50"
          v-model="textVar"
          type="search"
          placeholder="Search users"
        />
        <button
          class="btn btn-light ms-2"
          type="submit"
          @click.prevent="searchFunc"
          style="display: none;"
        >
          Search
        </button>
      </form>
    </div>

    <div class="col-4 d-flex justify-content-end">
      <button @click="addPhoto" class="my-trnsp-btn me-2" type="button">
        <i class="my-nav-icon-plus w-100 h-100 fa-solid fa-plus-circle" style="color: green;"></i>
      </button>

      <button @click="myProfile" class="my-trnsp-btn me-2" type="button">
        <i :class="'my-nav-icon-profile me-1 w-100 h-100 '+iconProfile+ ' fa-user'" @mouseover="profileIconActive" @mouseout="profileIconInactive"></i>
      </button>

      <button @click="logout" class="my-trnsp-btn me-2" type="button">
        <i class="my-nav-icon-quit me-1 w-100 h-100 fa-solid fa-right-from-bracket"></i>
      </button>
    </div>

    <!-- Input nascosto per il caricamento del file -->
    <input id="fileUploader" type="file" style="display: none" @change="uploadFile" accept=".jpg, .png" />
  </nav>
</template>

<style>
body {
  overflow: hidden;
}

.my-nav {
  background-color: #f8f8f8; /* Colore di sfondo pulito */
  padding: 20px; /* Aggiunta di spaziatura interna per una migliore visualizzazione */
}

.brand-text {
  color: black; /* Colore del testo */
  font-size: 1.5rem; /* Dimensione del testo */
}

.my-trnsp-btn {
  background: transparent;
  border: none;
  cursor: pointer;
}

.my-nav-icon-plus,
.my-nav-icon-profile,
.my-nav-icon-quit {
  font-size: 1.5rem; /* Dimensione delle icone */
  transition: color 0.3s, transform 0.3s;
}

.my-nav-icon-quit:hover {
  color: var(--color-red-danger);
  transform: scale(1.2);
}

.my-nav-icon-plus:hover {
  color: green;
  transform: scale(1.2);
}
</style>
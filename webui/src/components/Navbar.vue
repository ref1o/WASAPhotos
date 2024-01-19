<script>
export default {
    data() {
        return {
            textVar: "",
            iconProfile: "fa-regular",
        };
    },
    computed: {
        currentPath() {
            return this.$route.params.id
        },
        sameUser() {
            return this.$route.params.id === localStorage.getItem('token')
        }
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
            try {
                let fileInput = document.getElementById('fileUploader');
                const file = fileInput.files[0];
                const reader = new FileReader();

                reader.readAsArrayBuffer(file);

                reader.onload = async () => {
                    const token = localStorage.getItem('token');
                    if (token) {
                        // Post photo: /users/:id/photos
                        let response = await this.$axios.post(`/users/${token}/photos`, reader.result, {
                            headers: {
                                'Content-Type': file.type,
                                'Authorization': `Bearer ${token}`,
                            },
                        });

                        this.photos.unshift(response.data);
                        this.postCnt += 1;
                    } else {
                        console.error('Token di autenticazione mancante.');
                    }
                };
            } catch (error) {
                console.error('Errore durante il caricamento della foto:', error);
            }
        },

        async loadInfo() {
            if (this.$route.params.id === undefined) {
                return
            }

            try {
                // Get user profile: /users/:id
                let response = await this.$axios.get("/users/" + this.$route.params.id);

                this.banStatus = false
                this.userExists = true
                this.currentIsBanned = false

                if (response.status === 206) {
                    this.banStatus = true
                    return
                }

                if (response.status === 204) {
                    this.userExists = false
                }

                this.nickname = response.data.nickname
                this.followerCnt = response.data.followers != null ? response.data.followers.length : 0
                this.followingCnt = response.data.following != null ? response.data.following.length : 0
                this.postCnt = response.data.posts != null ? response.data.posts.length : 0
                this.followStatus = response.data.followers != null ? response.data.followers.find(obj => obj.user_id === localStorage.getItem('token')) : false
                this.photos = response.data.posts != null ? response.data.posts : []
                this.followers = response.data.followers != null ? response.data.followers : []
                this.following = response.data.following != null ? response.data.following : []

            } catch (e) {
                this.currentIsBanned = true
            }
        },

    },
    async mounted() {
        await this.loadInfo()
    },
};
</script>

<template>
    <nav class="navbar navbar-expand-lg bg-light d-flex justify-content-between sticky-top mb-3 my-nav">
        <div class="col-4">
            <a class="navbar-brand ms-2 d-flex" @click="goBackHome">
            <img src="@/assets/images/wasaLogo.png" alt="Logo" class="logo-img">
            <img src="@/assets/images/wasaLogoScritta.png" alt="LogoS" class="logo-scritta-img">
          </a>
        </div>
    
        <div class="col-4">
            <form class="form-inline my-2 my-lg-0 d-flex justify-content-center m-auto">
                <input class="form-control mr-sm-2 w-50" v-model="textVar" type="search" placeholder="Search users" />
                <button class="btn btn-light ms-2" type="submit" @click.prevent="searchFunc" style="display: none;">
              Search
            </button>
            </form>
        </div>
    
        <div class="col-4 d-flex justify-content-end">
            <label for="fileUploader" class="my-btn-add-photo ms-2">
            <i class="fas fa-plus"></i>
          </label>
            <input id="fileUploader" type="file" class="profile-file-upload" @change="uploadFile" accept=".jpg, .png" />
    
            <button @click="myProfile" class="my-trnsp-btn me-2" type="button">
            <i
              :class="'my-nav-icon-profile me-1 w-100 h-100 ' + iconProfile + ' fa-user'"
              @mouseover="profileIconActive"
              @mouseout="profileIconInactive"
            ></i>
          </button>
    
            <button @click="logout" class="my-trnsp-btn me-2" type="button">
            <i class="my-nav-icon-quit me-1 w-100 h-100 fa-solid fa-right-from-bracket"></i>
          </button>
        </div>
    </nav>
</template>

<style>
body {
    overflow: scroll;
}

.my-nav {
    background-color: #f8f8f8;
    /* Colore di sfondo pulito */
    padding: 20px;
    /* Aggiunta di spaziatura interna per una migliore visualizzazione */
}

.brand-text {
    color: black;
    /* Colore del testo */
    font-size: 1.5rem;
    /* Dimensione del testo */
}

.my-trnsp-btn {
    background: transparent;
    border: none;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: transform 0.3s;
}

.my-trnsp-btn:hover {
    transform: scale(1.2);
}

.my-trnsp-btn i {
    font-size: 1.5rem;
    /* Adjust the icon size */
}

.my-nav-icon-plus,
.my-nav-icon-profile,
.my-nav-icon-quit {
    font-size: 1.5rem;
    /* Dimensione delle icone */
    transition: color 0.3s, transform 0.3s;
}

.my-nav-icon-quit:hover {
    color: var(--color-red-danger);
    transform: scale(1.2);
}

.my-btn-add-photo {
    background-color: #4caf50;
    /* Green */
    border: none;
    padding: 0.4rem;
    cursor: pointer;
    display: flex;
    margin-right: 1.0rem;
    align-items: center;
    justify-content: center;
    border-radius: 50%;
    transition: background-color 0.3s, transform 0.3s;
}

.my-btn-add-photo:hover {
    background-color: #45a049;
    /* Darker green on hover */
    transform: scale(1.2);
}

.my-btn-add-photo i {
    color: white;
    font-size: 1.0rem;
    /* Adjust the icon size */
}

#fileUploader {
    display: none;
}

.logo-img {
    max-width: 50px;
    height: auto;
    display: block;
    margin-right: 10px;
    margin-left: 10px;
}

.logo-scritta-img {
    max-width: 200px;
    height: auto;
    display: block;
    margin-left: 0;
}
</style>
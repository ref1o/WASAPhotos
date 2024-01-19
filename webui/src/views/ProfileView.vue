<script>
export default {
    data: function() {
        return {
            errormsg: null,

            userExists: false,
            banStatus: false,

            nickname: "",


            followStatus: false,
            currentIsBanned: false,

            followerCnt: 0,
            followingCnt: 0,
            postCnt: 0,

            photos: [],
            following: [],
            followers: [],
        }
    },

    watch: {
        currentPath(newid, oldid) {
            if (newid !== oldid) {
                this.loadInfo()
            }
        },
    },

    computed: {

        currentPath() {
            return this.$route.params.id
        },

        sameUser() {
            return this.$route.params.id === localStorage.getItem('token')
        },
    },

    methods: {
        async uploadFile() {
            let fileInput = document.getElementById('fileUploader')

            const file = fileInput.files[0];
            const reader = new FileReader();

            reader.readAsArrayBuffer(file);

            reader.onload = async () => {
                // Post photo: /users/:id/photos
                let response = await this.$axios.post("/users/" + this.$route.params.id + "/photos", reader.result, {
                    headers: {
                        'Content-Type': file.type
                    },
                })
                //console.log(response)
                /*
                this.photos.unshift({
                    owner: response.data.owner,
                    date: response.data.date,
                    photo_id: response.data.photo_id,
                    likes: response.data.likes,
                    comments: response.data.comments,
                })
                */
                this.photos.unshift(response.data)
                this.postCnt += 1
            };
            window.location.reload()
        },

        async followClick() {
            try {
                if (this.followStatus) {
                    // Delete like: /users/:id/followers/:follower_id
                    await this.$axios.delete("/users/" + this.$route.params.id + "/followers/" + localStorage.getItem('token'));
                    this.followerCnt -= 1
                } else {
                    // Put like: /users/:id/followers/:follower_id
                    await this.$axios.put("/users/" + this.$route.params.id + "/followers/" + localStorage.getItem('token'));
                    this.followerCnt += 1
                }
                this.followStatus = !this.followStatus
            } catch (e) {
                this.errormsg = e.toString();
            }
        },

        async banClick() {
            try {
                if (this.banStatus) {
                    // Delete ban: /users/:id/banned_users/:banned_id
                    await this.$axios.delete("/users/" + localStorage.getItem('token') + "/banned_users/" + this.$route.params.id);
                    this.loadInfo()
                } else {
                    // Put ban: /users/:id/banned_users/:banned_id
                    await this.$axios.put("/users/" + localStorage.getItem('token') + "/banned_users/" + this.$route.params.id);
                    this.followStatus = false
                }
                this.banStatus = !this.banStatus
            } catch (e) {
                this.errormsg = e.toString();
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

        goToSettings() {
            this.$router.push(this.$route.params.id + '/settings')
        },

        removePhotoFromList(photo_id) {
            this.photos = this.photos.filter(item => item.photo_id !== photo_id)
        },
    },

    async mounted() {
        await this.loadInfo()
    },

}
</script>

<template>
    <div class="profile-page">
        <div class="profile-header">
            <div class="profile-card">
                <div class="profile-info">
                    <h5 class="profile-username">{{ nickname }} @{{ this.$route.params.id }}</h5>
                </div>
                <div class="profile-actions">
                    <button v-if="!sameUser && !banStatus" @click="followClick" class="btn btn-primary">
                      {{ followStatus ? "Unfollow" : "Follow" }}
                    </button>
                    <button v-if="!sameUser" @click="banClick" class="btn btn-danger">
                      {{ banStatus ? "Unban" : "Ban" }}
                    </button>
                    <button v-else class="btn btn-outline-secondary" @click="goToSettings">
                      <i class="fas fa-cog"></i> Settings
                    </button>
                </div>
                <div class="profile-stats">
                    <div class="profile-stat">
                        <strong>{{ postCnt }}</strong>
                        <span>Posts</span>
                    </div>
                    <div class="profile-stat">
                        <strong>{{ followerCnt }}</strong>
                        <span>Followers</span>
                    </div>
                    <div class="profile-stat">
                        <strong>{{ followingCnt }}</strong>
                        <span>Following</span>
                    </div>
                </div>
            </div>
        </div>
        <div class="profile-posts">
            <h2>Posts</h2>
            <hr>
            <div v-if="!banStatus && postCnt > 0" class="post-grid">
                <div v-for="(photo, index) in photos" :key="index" class="post-item">
                    <Photo :owner="this.$route.params.id" :photo_id="photo.photo_id" :comments="photo.comments" :likes="photo.likes" :upload_date="photo.date" :isOwner="sameUser" @removePhoto="removePhotoFromList" />
                </div>
            </div>
            <div v-else class="no-posts">
                <h2>No posts yet</h2>
            </div>
        </div>
        <ErrorMsg v-if="errormsg" :msg="errormsg" />
    </div>
</template>

<style>
.profile-page {
    background-color: #f5f5f5;
    padding: 20px;
}

.profile-header {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-bottom: 20px;
}

.profile-card {
    background-color: #fff;
    border-radius: 10px;
    padding: 20px;
    display: flex;
    flex-direction: column;
    align-items: center;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
}

.profile-info {
    text-align: center;
}

.profile-username {
    margin-top: 10px;
}

.profile-actions {
    margin-top: 20px;
    display: flex;
    gap: 10px;
}

.profile-actions button {
    flex: 1;
}

.profile-actions i {
    margin-right: 5px;
}

.profile-stats {
    display: flex;
    justify-content: center;
    margin-top: 20px;
    gap: 20px;
}

.profile-stat {
    text-align: center;
}

.profile-stat strong {
    display: block;
    font-size: 18px;
}

.profile-stat span {
    font-size: 14px;
}

.profile-posts {
    text-align: center;
}

.profile-posts h2 {
    margin-bottom: 10px;
}

.post-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    /* 3 colonne per riga */
    gap: 10px;
}

.no-posts h2 {
    color: #fff;
}
</style>
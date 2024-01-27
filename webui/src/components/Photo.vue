<script>
export default {
	data() {
		return {
			photoURL: "",
			liked: false,
			allComments: [],
			allLikes: [],
		};
	},

	props: ["owner", "likes", "comments", "upload_date", "photo_id", "isOwner"],

	computed: {
		formattedUploadDate() {
			const options = { year: "numeric", month: "long", day: "numeric", hour: "numeric", minute: "numeric", second: "numeric" };
			return new Date(this.upload_date).toLocaleDateString(undefined, options);
		},
	},

	methods: {
		loadPhoto() {
			// Get photo : "/users/:id/photos/:photo_id"
			this.photoURL = __API_URL__ + "/users/" + this.owner + "/photos/" + this.photo_id;
		},

		async deletePhoto() {
			try {
				// Delete photo: /users/:id/photos/:photo_id
				await this.$axios.delete("/users/" + this.owner + "/photos/" + this.photo_id);
				this.$emit("removePhoto", this.photo_id);
				window.location.reload();
			} catch (e) {
				//
			}
		},

		photoOwnerClick: function() {
			this.$router.replace("/users/" + this.owner);
		},

		async toggleLike() {
			if (this.isOwner) {
				return;
			}

			const bearer = localStorage.getItem("token");

			try {
				if (!this.liked) {
					// Put like: /users/:id/photos/:photo_id/likes/:like_id"
					await this.$axios.put("/users/" + this.owner + "/photos/" + this.photo_id + "/likes/" + bearer);
					this.allLikes.push({
						user_id: bearer,
						nickname: bearer,
					});
				} else {
					// Delete like: /users/:id/photos/:photo_id/likes/:like_id"
					await this.$axios.delete("/users/" + this.owner + "/photos/" + this.photo_id + "/likes/" + bearer);
					this.allLikes.pop();
				}

				this.liked = !this.liked;
			} catch (e) {
				//
			}
		},

		removeCommentFromList(value) {
			this.allComments = this.allComments.filter((item) => item.comment_id !== value);
		},

		addCommentToList(comment) {
			this.allComments.push(comment);
		},

		openLikeModal() {
			this.$emit("openLikeModal");
		},
	},

	async mounted() {
		await this.loadPhoto();

		if (this.likes != null) {
			this.allLikes = this.likes;
		}

		if (this.likes != null) {
			this.liked = this.allLikes.some((obj) => obj.user_id === localStorage.getItem("token"));
		}

		if (this.comments != null) {
			this.allComments = this.comments;
		}
	},
};
</script>

<template>
	<div class="photo-card">
		<LikeModal :modal_id="'like_modal' + photo_id" :likes="allLikes" @openLikeModal="openLikeModal" />
		<CommentModal :modal_id="'comment_modal' + photo_id" :comments_list="allComments" :photo_owner="owner" :photo_id="photo_id" @eliminateComment="removeCommentFromList" @addComment="addCommentToList" />
	
		<div class="photo-card-content">
			<div class="photo-card-image">
				<img :src="photoURL" alt="Photo" class="img-fluid" />
			</div>
	
			<div class="photo-card-info">
				<div class="photo-card-left">
					<button class="btn btn-light btn-sm me-2" @click="photoOwnerClick">
					  <b class="no-border">{{ owner }}</b>
					</button>
				</div>
	
				<div class="photo-card-right">
					<button class="btn btn-light btn-sm me-2 no-border" @click="toggleLike">
					  <i :class="'me-1 fas ' + (liked ? 'fa-heart text-danger' : 'fa-heart text-muted')"></i>
					  {{ allLikes.length }}
					</button>
					<button class="btn btn-light btn-sm no-border" data-bs-toggle="modal" :data-bs-target="'#comment_modal' + photo_id">
					  <i class="far fa-comment me-1 text-muted"></i>
					  {{ allComments != null ? allComments.length : 0 }}
					</button>
				</div>
			</div>
	
			<div class="photo-card-details" style="text-align: left; margin: 0; padding: 0">
				<p>Uploaded on {{ formattedUploadDate }}</p>
            	<div v-if="isOwner" class="delete-icon-container">
                	<i class="fa fa-trash delete-icon" @click="deletePhoto"></i>
            	</div>
			</div>
		</div>
	</div>
</template>
  
<style scoped>
.photo-card {
	background-color: #fff;
	border-radius: 10px;
	box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
	margin-bottom: 20px;
}

.photo-card-content {
	display: flex;
	flex-direction: column;
	padding: 20px;
}

.photo-card-image {
	width: 100%;
	max-height: 400px;
	overflow: hidden;
}

.photo-card-image img {
	width: 100%;
	height: auto;
}

.photo-card-info {
	display: flex;
	justify-content: space-between;
	align-items: center;
	margin-top: 10px;
}

.photo-card-left {
	display: flex;
	align-items: center;
}

.photo-card-left b {
	cursor: pointer;
	text-decoration: underline;
}

.photo-card-right {
	display: flex;
	align-items: center;
}

.photo-card-right button:hover .fa-heart {
	color: red;
}

.photo-card-actions button {
	font-size: 14px;
	padding: 6px 12px;
}

.photo-card-meta button {
	font-size: 14px;
	padding: 6px 12px;
	display: flex;
	align-items: center;
	gap: 5px;
}

.photo-card-meta i {
	font-size: 16px;
}

.photo-card-details {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.photo-card-details p {
	font-size: 14px;
	color: #888;
	margin: 0;
	padding: 0;
}

.no-border {
	border: none;
	background-color: transparent;
	padding: 0;
}

.delete-icon-container {
    display: flex;
    align-items: center;
}

.delete-icon {
    cursor: pointer;
    color: #000000;
    font-size: 20px;
}

.delete-icon:hover{
	color: red;

}
</style>
  
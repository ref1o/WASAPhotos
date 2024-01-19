<script>
export default {
	data() {
		return {
			commentValue: "",
		};
	},
	props: ["modal_id", "comments_list", "photo_owner", "photo_id"],
	methods: {
		async addComment() {
			try {
				// Comment post: /users/:id/photos/:photo_id/comments
				let response = await this.$axios.post(
					"/users/" + this.photo_owner + "/photos/" + this.photo_id + "/comments", {
						user_id: localStorage.getItem("token"),
						comment: this.commentValue,
					}, {
						headers: {
							"Content-Type": "application/json",
						},
					}
				);

				this.$emit("addComment", {
					comment_id: response.data.comment_id,
					photo_id: this.photo_id,
					user_id: localStorage.getItem("token"),
					comment: this.commentValue,
					nickname: localStorage.getItem("nickname"),
				});
				this.commentValue = "";
			} catch (e) {
				console.log(e.toString());
			}
		},

		eliminateCommentToParent(value) {
			this.$emit("eliminateComment", value);
		},

		addCommentToParent(newCommentJSON) {
			this.$emit("addComment", newCommentJSON);
		},
	},
};
</script>

<template>
    <div class="modal fade custom-modal" :id="modal_id" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered modal-dialog-scrollable">
            <div class="modal-content">
                <div class="modal-header">
                    <h1 class="modal-title" :id="`heading-${modal_id}`">Comments</h1>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>

                <div class="modal-body">
                    <div v-if="comments_list.length === 0" class="text-center text-muted">
                        No comments yet.
                    </div>

                    <div v-else class="comments-container">
                        <div class="comment-box" v-for="(comm, index) in comments_list" :key="index">
                            <PhotoComment :author="comm.user_id" :nickname="comm.nickname" :comment_id="comm.comment_id" :photo_id="comm.photo_id" :content="comm.comment" :photo_owner="photo_owner" @eliminateComment="eliminateCommentToParent"/>
                        </div>
                    </div>
                </div>

                <div class="modal-footer">
                    <div class="row w-100">
                        <div class="col-10">
                            <div class="mb-3 me-auto">
                                <textarea class="form-control comment-input" placeholder="Add a comment..." rows="3" maxLength="30" v-model="commentValue"></textarea>
                            </div>
                        </div>
                        <div class="col-2 d-flex align-items-center">
                            <button type="button" class="btn btn-send" @click.prevent="addComment" :disabled="commentValue.length < 1 || commentValue.length > 30">Send</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
  
<style>
.custom-modal .modal-content {
    background-color: #f8f9fa;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.custom-modal .modal-header {
    border-bottom: 1px solid #dee2e6;
    font-size: 1.25rem;
}

.custom-modal .modal-body {
    padding: 1rem;
}

.comments-container {
    padding: 15px;
}

.comment-box {
    background-color: #fff;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    margin-bottom: 15px;
    overflow: hidden;
}

.custom-modal .comment-input {
    min-height: 80px;
    border-radius: 5px;
}

.custom-modal .btn-send {
    background-color: #007bff;
    border-radius: 5px;
    color: white;
}

.custom-modal .btn-close {
    color: #6c757d;
    font-size: 1.2rem;
}
</style>
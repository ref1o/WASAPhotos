<script>
export default {
    data(){
        return {
            user: "",
        }
    },
	props: ['content','author','photo_owner','comment_id','photo_id','nickname'],

    methods:{
        async deleteComment(){
            try{
                // Delete comment: "/users/:id/photos/:photo_id/comments/:comment_id"
                await this.$axios.delete("/users/"+this.photo_owner+"/photos/"+this.photo_id+"/comments/"+this.comment_id)

                this.$emit('eliminateComment',this.comment_id)

            }catch (e){
                console.log(e.toString())
            }
        },
    },

    mounted() {
        this.user = localStorage.getItem('token');
        // Debugging: stampa delle props per verificare i valori passati
        console.log('PhotoComment Mounted:', {
            content: this.content,
            author: this.author,
            photo_owner: this.photo_owner,
            comment_id: this.comment_id,
            photo_id: this.photo_id,
            nickname: this.nickname,
        });
    }
}
</script>

<template>
    <div class="comment-container">
        <div class="comment-header">
            <h5 class="comment-author">{{nickname}} <span class="author-id">@{{author}}</span></h5>
            <button v-if="user === author || user === photo_owner" class="btn delete-btn" @click="deleteComment">
                <i class="fa-regular fa-trash-can"></i>
            </button>
        </div>
        <p class="comment-content">{{content}}</p>
    </div>
</template>

<style>
.comment-container {
    padding: 10px;
    margin-bottom: 10px;
}

.comment-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.comment-author {
    margin: 0;
    font-weight: bold;
    color: #007bff; /* O scegli un colore che si adatti al tuo tema */
}

.author-id {
    font-weight: normal;
    color: #6c757d;
}

.delete-btn {
    border: none;
    background: none;
    color: #dc3545; /* Colore rosso per il cestino */
}

.delete-btn:hover {
    color: #bd2130; /* Colore al passaggio del mouse */
    transform: scale(1.1);
}

.comment-content {
    margin-top: 5px;
    line-height: 1.4;
}
</style>

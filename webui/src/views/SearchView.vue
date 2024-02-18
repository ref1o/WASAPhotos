<script>
export default {
	data: function() {
		return {
			users: [],
			errormsg: null,
		}
	},

	props:['searchValue'],

	watch:{
		searchValue: function(){
			this.loadSearchedUsers()
		},
	},

	methods:{
		async loadSearchedUsers(){
			this.errormsg = null;
			if (
				this.searchValue === undefined ||
				this.searchValue === "" || 
				this.searchValue.includes("?") ||
				this.searchValue.includes("_")){
				this.users = []
				return 
			}
			try {
				let response = await this.$axios.get("/users",{
						params: {
						id: this.searchValue,
					},
				});
				this.users = response.data

			} catch (e) {
				this.errormsg = e.toString();
			}
		},

		goToProfile(profileId){
			this.$router.replace("/users/"+profileId)
		}
	},

	async mounted(){
		// Check if the user is logged
		if (!localStorage.getItem('token')){
			this.$router.replace("/login")
		}
		await this.loadSearchedUsers()
		
	},
}
</script>

<template>
    <div class="container-fluid">
        <div v-for="(user, index) in users" :key="index" class="user-card" @click="goToProfile(user.user_id)">
            <UserMiniCard :identifier="user.user_id" :nickname="user.nickname" />
        </div>

        <p v-if="users.length === 0" class="no-result-text d-flex justify-content-center"> Nessun utente trovato.</p>

        <div v-if="errormsg" class="error-msg">
            <ErrorMsg :msg="errormsg" />
        </div>
    </div>
</template>


<style>
.container-fluid {
    background-color: #f8f9fa; /* Sfondo chiaro */
    padding: 2rem; /* Più spazio intorno agli elementi */
}

.user-card {
    background-color: white;
    border-radius: 10px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    padding: 1rem;
    margin: 1rem auto; /* Centra la card e aggiunge margine sopra e sotto */
    max-width: 600px; /* Limita la larghezza massima della card */
    cursor: pointer;
    transition: transform 0.2s;
}

.user-card:hover {
    transform: scale(1.05);
}

.no-result-text {
    color: #6c757d; /* Colore del testo */
    font-style: italic;
    font-size: 1.2rem; /* Dimensione del testo */
    margin-top: 2rem; /* Spazio superiore */
}

.error-msg {
    color: var(--color-red-danger); /* Colore per i messaggi di errore */
    text-align: center; /* Allineamento del testo */
    margin-top: 2rem; /* Spazio superiore */
}
</style>
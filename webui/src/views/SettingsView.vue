<script>
export default {
	data: function() {
		return {
			errormsg: null,
			nickname: "",
		}
	},

	methods: {
		async modifyNickname() {
			try {
				// Nickname put: /users/:id
				let resp = await this.$axios.put("/users/" + this.$route.params.id, {
					nickname: this.nickname,
				})

				localStorage.setItem('nickname', this.nickname);

				this.nickname = ""
			} catch (e) {
				this.errormsg = e.toString();
			}
		},
	},

}
</script>

<template>
	<div class="container-fluid settings-container">
		<!-- Intestazione della Pagina -->
		<div class="row">
			<div class="col d-flex justify-content-center mb-4">
				<h1 class="settings-title">Impostazioni dell'Account</h1>
			</div>
		</div>
	
		<!-- Sezione Informativa -->
		<div class="row info-section">
			<div class="col-12 text-center">
				<p class="info-text">Qui puoi personalizzare il tuo <strong>nickname</strong>. Ricorda, il tuo nickname rappresenta la tua identità unica all'interno della nostra community!</p>
				<p class="info-text">Il tuo <strong>identificatore</strong> (parte dopo la @) rimarrà invariato.</p>
			</div>
		</div>
	
		<!-- Form per la Modifica del Nickname -->
		<div class="row">
			<div class="col d-flex justify-content-center">
				<div class="input-group nickname-input-group">
					<input type="text" class="form-control" placeholder="Nuovo nickname..." maxlength="16" minlength="3" v-model="nickname" />
					<div class="input-group-append">
						<button class="btn btn-outline-secondary" @click="modifyNickname" :disabled="nickname === null || nickname.length >16 || nickname.length <3 || nickname.trim().length===0">
							Modifica</button>
					</div>
				</div>
			</div>
		</div>
	
		<!-- Anteprima del Nickname Modificato -->
		<div class="row">
			<div v-if="nickname" class="col d-flex justify-content-center">
				<p class="preview">Anteprima: {{nickname}} @{{ this.$route.params.id }}</p>
			</div>
		</div>
	
		<!-- Messaggio di Errore -->
		<div class="row">
			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		</div>
	</div>
</template>

<style>
.settings-container {
	background-color: #f8f9fa;
	/* Sfondo chiaro */
	padding: 2rem;
	/* Più spazio intorno agli elementi */
}

.settings-title {
	color: #007bff;
	/* Colore del titolo */
	font-size: 2.5rem;
	/* Dimensione del titolo */
}

.info-section .info-text {
	color: #6c757d;
	/* Colore del testo informativo */
	font-size: 1rem;
	/* Dimensione del testo informativo */
	margin-bottom: 0.5rem;
	/* Spazio tra le righe informative */
}

.nickname-input-group {
	max-width: 400px;
	/* Limita la larghezza per un aspetto più nitido */
}

.form-control {
	border-radius: 20px 20px 20px 20px;
	/* Bordi arrotondati per un look moderno */
	border: 1px solid #ced4da;
	/* Bordo sottile */
}

.btn-outline-secondary {
	border-radius: 20px;
	/* Bordi arrotondati per il pulsante */
	border-color: #007bff;
	/* Colore del bordo del pulsante */
	color: #007bff;
	/* Colore del testo del pulsante */
	transition: all 0.2s;
	/* Transizione fluida al passaggio del mouse */
}

.btn-outline-secondary:hover {
	background-color: #007bff;
	/* Cambia sfondo al passaggio del mouse */
	color: white;
	/* Cambia colore del testo al passaggio del mouse */
}

.preview {
	color: #6c757d;
	/* Colore del testo per la preview */
	font-style: italic;
	/* Stile del testo per la preview */
}
</style>

<style>
.container-fluid {
	background-color: #f8f9fa;
	/* Sfondo chiaro */
	padding: 2rem;
	/* Più spazio intorno agli elementi */
}

.row {
	margin-bottom: 1rem;
	/* Spazio tra le righe */
}

.input-group {
	max-width: 400px;
	/* Limita la larghezza per un aspetto più nitido */
}

.form-control {
	border-radius: 20px;
	/* Bordi arrotondati per un look moderno */
	border: 1px solid #ced4da;
	/* Bordo sottile */
}

.btn-outline-secondary {
	border-radius: 20px;
	/* Bordi arrotondati per il pulsante */
	border-color: #007bff;
	/* Colore del bordo del pulsante */
	color: #007bff;
	/* Colore del testo del pulsante */
	transition: all 0.2s;
	/* Transizione fluida al passaggio del mouse */
}

.btn-outline-secondary:hover {
	background-color: #007bff;
	/* Cambia sfondo al passaggio del mouse */
	color: white;
	/* Cambia colore del testo al passaggio del mouse */
}

.preview {
	color: #6c757d;
	/* Colore del testo per la preview */
	font-style: italic;
	/* Stile del testo per la preview */
}
</style>

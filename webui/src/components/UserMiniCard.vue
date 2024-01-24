<script>
export default {
	props: ['identifier','nickname'],

    methods:{

        clickedUser(){
            this.$emit('clickedUser',this.identifier)
        }
    },
}
</script>

<template>
    <div class="user-mini-card card" @click="clickedUser">
        <div class="card-body">
            <h5 class="user-mini-card-title">{{ nickname }} @{{ identifier }}</h5>
        </div>
    </div>
</template>


<style>
.user-mini-card {
    color: #343a40; /* Colore del testo */
    width: 100%; /* Usa tutta la larghezza disponibile */
    border: none;
    background-color: transparent; /* Sfondo trasparente */
    transition: background-color 0.2s; /* Transizione per hover */
}

.user-mini-card:hover {
    background-color: #f8f9fa; /* Cambia colore di sfondo all'hover */
}

.user-mini-card-title {
    cursor: pointer;
    font-weight: 500; /* Peso del font */
    font-size: 1.2rem; /* Dimensione del testo */
    text-align: center; /* Allinea il testo al centro */
}

.card-body {
    padding: 1rem; /* Spazio interno */
}
</style>

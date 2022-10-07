<script setup>
import { RouterLink, RouterView } from "vue-router";
import Navigation from "./components/Navigation.vue";
</script>

<template>
<div>
  <navigation></navigation>
  <section>
    <div class="container"><RouterView /></div>
    
  </section>
</div>
</template>

<style scoped>

</style>

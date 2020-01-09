<template>
  <div>
    <div class="header-container">
      <main-header></main-header>
      <nav-bar></nav-bar>
    </div>
    <section
      class="hero serv-bg-image service-children-image is-danger is-large"
      :style="[{ 'background-image': `url(${servicesBanner.image.data})`},
      {'background-position': 'center'},
      {'background-size': 'cover'},
      {'background-repeat': 'no-repeat'}
      ]"
    >
      <div class="hero-body">
        <div class="container"></div>
      </div>
    </section>
    <section class="web-pages">
      <nuxt />
    </section>
    <section class="footer">
      <nav-footer></nav-footer>
    </section>
    <section class="footer2"></section>
  </div>
</template>

<script>
import MainHeader from "@/components/header";
import NavBar from "@/components/NavBar";
import NavFooter from "@/components/Footer";
import { mapState } from "vuex";

export default {
  components: {
    MainHeader,
    NavBar,
    NavFooter
  },

  computed: {
    ...mapState(["servicesBanner"])
  }
};
</script>

<style lang="" scoped></style>

<template>
  <div>
    <div class="header-container">
      <main-header></main-header>
      <nav-bar></nav-bar>
    </div>
    <section
      class="hero hero-bg is-primary is-fullheight"
      :style="{ 'background-image': `url(/img/banner/${Banner})` }"
    >
      <div class="hero-body">
        <div class="container homepage-container has-text-centered">
          <img class src="/img/banner/white-logo.png" alt />
          <div class="container-content">
            <p class="main-title has-text-white is-gbold">Providing the Quality home care you need</p>
            <p class="main-subtitle">Coverage areas include Bexley and Royal Borough of Greenwich</p>
            <div class="buttons-container">
              <div class="columns">
                <div class="column"></div>
                <div class="column">
                  <div class="level">
                    <div class="level-item">
                      <nuxt-link to="/process">
                        <b-button
                          class="has-fullwidth"
                          type="is-white is-size-5"
                          outlined
                          rounded
                        >Learn More</b-button>
                      </nuxt-link>
                    </div>
                    <div class="level-item">
                      <nuxt-link to="/process">
                        <b-button type="is-white is-size-5" outlined rounded>Request a call back</b-button>
                      </nuxt-link>
                    </div>
                  </div>
                </div>
                <div class="column"></div>
              </div>
            </div>
          </div>
        </div>
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
    ...mapState(["Banner"])
  }
};
</script>

<style lang="" scoped></style>

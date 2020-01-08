<template>
  <div class="process">
    <div class="pg-title is-blue-bold">Our Care Process</div>
    <div class="pg-process-content">
      <div class="columns">
        <div class="column is-2">
          <img class="process-image" src="/img/banner/process/01.png" alt />
        </div>
        <div class="column">
          
          <div
            class="process-title is-blue is-gmedium has-text-weight-semibold"
          >GET IN TOUCH WITH US</div>
          <div class="proceess-content">
            FIll in Our call back form or give us a call
            to find out how we can help you
          </div>
        </div>
      </div>
      <div class="columns">
        <div class="column is-2">
          <img class="process-image" src="/img/banner/process/02.png" alt />
        </div>
        <div class="column">
          <div class="process-title is-blue iis-gmedium has-text-weight-semibold">ASSESSMENT</div>
          <div class="proceess-content">
            We'll come out to you to find out what you or your loved ones need
            to help stay independent at home
          </div>
        </div>
      </div>
      <div class="columns">
        <div class="column is-2">
          <img class="process-image" src="/img/banner/process/03.png" alt />
        </div>
        <div class="column">
          <div
            class="process-title is-blue is-gmedium has-text-weight-semibold"
          >CARE ASSISTANT CHOSEN &amp; CARE STARTS</div>
          <div class="proceess-content">
            You'll be cared for by one of our specially trained team to support you
            to remain at home for as long as possible
          </div>
        </div>
      </div>
      <nuxt-link to="/contact">
        <b-button type="is-bg-red has-text-white is-size-5" outlined rounded>GET IN TOUCH WITH US</b-button>
      </nuxt-link>
    </div>

    <!-- <b-button outlined rounded type="has-text-white is-bg-red">GET IN TOUCH WITH US</b-button> -->
  </div>
</template>

<script>
export default {
  head() {
    return {
      title: "Process"
    };
  },
  async asyncData({ store }) {
    store.commit("setbanner", "Our-Process-Page-Image.jpg");
  }
};
</script>

<style lang="scss" scoped></style>

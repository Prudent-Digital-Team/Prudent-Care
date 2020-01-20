<template>
  <div class="services">
    <div class="pg-title is-blue-bold">Our Services</div>
    <div class="pg-services-content" v-html="page.attributes.content.content"></div>
    <div class="service-dropdown">
      <b-collapse class="card faq-card" :open="isOpen">
        <div slot="trigger" class="card-header" role="button" aria-controls="contentIdForA11y3">
          <p class="card-header-title is-blue-bold is-size-5">Services List</p>
          <a class="card-header-icon">
            <b-icon :icon="isOpen ? 'menu-down' : 'menu-up'"></b-icon>
          </a>
        </div>
        <div class="card-content">
          <div
            class="content is-blue is-glight is-size-5"
            v-html="page.attributes.content.serviceList"
          ></div>
        </div>
      </b-collapse>
    </div>
    <nuxt-link to="/contact">
      <b-button type="is-bg-red has-text-white is-size-6" outlined rounded>GET IN TOUCH WITH US</b-button>
    </nuxt-link>
  </div>
</template>

<script>
import PageMixin from "@/mixins/index";
export default {
  mixins: [PageMixin],
  data() {
    return {
      isOpen: true
    };
  },
  head() {
    return {
      title: "Our Services",
      meta: [
        // hid is used as unique identifier. Do not use `vmid` for it as it will not work
        {
          hid: "title",
          name: "title",
          content:
            "Prudent Domiciliary Our Services. Care Company with coverage in Coverage areas include Bexley and Royal Borough of Greenwich"
        },
        {
          hid: "description",
          name: "description",
          content:
            "Bexely Foremost Care Company, Uk. We Providing the Quality home care you need Coverage areas include Bexely and Royal Borough of Greenwich. Contact us today!"
        },
        {
          hid: "keywords",
          name: "keywords",
          content:
            "Bexely, Care, BexleyCare, Bexely Dementia Care, Home Care United Kingdom, Elder Care "
        }
      ]
    };
  }
  // async asyncData({ store }) {
  //   store.commit("setbanner", "Our-Services-Page-Image.jpg");
  // }
};
</script>

<style lang="scss" scoped></style>

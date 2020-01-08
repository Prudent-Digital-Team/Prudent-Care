<template>
  <div class="main-header is-hidden-mobile">
    <nav class="level">
      <!-- Left side -->
      <div class="level-left">
        <div class="round is-bg-red">
          <a class="mdi mdi-phone has-text-white" href></a>
        </div>
        <div class="level-item" v-for="mobile in Header.contactNumbers" :key="mobile.id">
          <div class="mobile-numbers">
            <a href class>{{ mobile }}</a>
          </div>
        </div>
      </div>
      <div class="level-right">
        <div class="round is-bg-red">
          <a class="mdi mdi-map-marker has-text-white" href></a>
        </div>
        <div class="address">
          <div class="level-item">{{ Header.address }}</div>
        </div>
        <div class="socials">
          <div class="level-item">
            <!-- <div class="round is-bg-grey">
              <a class="mdi mdi-twitter has-text-white" href></a>
            </div>-->
            <div class="round is-bg-grey">
              <a
                class="mdi mdi-linkedin has-text-white"
                target="_blank"
                :href="Header.socialMedia.linkedin"
              ></a>
            </div>

            <div class="round is-bg-grey">
              <a
                class="mdi mdi-instagram has-text-white"
                target="_blank"
                :href="Header.socialMedia.instagram"
              ></a>
            </div>
            <div class="round is-bg-grey">
              <a
                class="mdi mdi-facebook has-text-white"
                target="_blank"
                :href="Header.socialMedia.facebook"
              ></a>
            </div>
          </div>
        </div>
      </div>
    </nav>
    <!-- <div class="mobile-nav is-hidden-desktop">
      <div class="columns is-mobile">
        <div class="column is-2">
          <div class="round is-bg-red">
            <a class="mdi mdi-phone has-text-white" href></a>
          </div>
        </div>
        <div class="column" v-for="mobile in Header.contactNumbers" :key="mobile.id">{{mobile}}</div>
      </div>
      <div class="columns is-mobile">
        <div class="column is-2">
          <div class="round is-bg-red">
            <a class="mdi mdi-map-marker has-text-white" href></a>
          </div>
        </div>
        <div class="column">{{ Header.address }}</div>
      </div>
      <div class="columns is-mobile">
        <div class="column">
          <div class="round is-bg-grey">
            <a class="mdi mdi-twitter has-text-white" href></a>
          </div>
        </div>
        <div class="column">
          <div class="round is-bg-grey">
            <a class="mdi mdi-linkedin has-text-white" href></a>
          </div>
        </div>
        <div class="column">
          <div class="round is-bg-grey">
            <a class="mdi mdi-instagram has-text-white" href></a>
          </div>
        </div>
        <div class="column">
          <div class="round is-bg-grey">
            <a class="mdi mdi-facebook has-text-white" href></a>
          </div>
        </div>
      </div>
    </div>-->
  </div>
</template>

<script>
export default {
  data() {
    return {
      Header: {
        contactNumbers: ["01322686765", "+447440713708"],
        address: "130 Old Street, London EC1V 9BD",
        socialMedia: {
          linkedin: "https://www.linkedin.com/company/prudent-domiciliary-care",
          twitter: "",
          facebook: "https://web.facebook.com/PrudentDomCare/",
          instagram:
            "https://instagram.com/prudentdomiciliarycare?igshid=7acpi27gh9fg",
          whatsapp: ""
        }
      }
    };
  }
};
</script>

<style lang="scss" scoped></style>

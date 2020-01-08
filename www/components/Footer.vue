<template>
  <div class="footer-content">
    <div class="columns">
      <div class="column">
        <div class="footer-content__footer-title">EXPLORE PBG DOMICILIARY CARE</div>
        <div class="footer-content__footer-link width30" v-for="a in Navigation" :key="a.id">
          <span class="nav-title">></span>
          <nuxt-link :to="a.link" class="has-text-white">
            {{
            a.name
            }}
          </nuxt-link>
        </div>
      </div>
      <div class="column">
        <div class="footer-content__footer-title">Our Services</div>
        <div v-for="a in Services" :key="a.id" class="footer-content__footer-link">
          <span class="nav-title">></span>
          <nuxt-link :to="'/services/'+a.link" class="has-text-white">
            {{
            a.name
            }}
          </nuxt-link>
          <div class="footer-underline"></div>
        </div>
      </div>
      <div class="column">
        <div class="footer-content__footer-title">Office Address</div>
        <div class="footer-content__footer-link">
          <div class="office-title">PRUDENT DOMICILIARY CARE LIMITED</div>
          <div class="office-reg">
            <div class="office-title">Registered Office</div>
            <div class="office-address">130 Old Street, London EC1V 9BD</div>
          </div>
          <a
            class="office-email has-text-white is-lowercase"
            href="mailto:info@pbgcare.co.uk"
          >info@pbgcare.co.uk</a>
        </div>
        <div class="box contact-box">
          <form @submit.prevent="submit">
            <b-field
              label="Fullname"
              custom-class="is-small"
              :type="errors.has('Fullname') ? 'is-danger': ''"
              :message="errors.first('Fullname')"
            >
              <b-input
                v-model="form.name"
                expanded
                name="Fullname"
                placeholder="Fullname"
                v-validate="'required'"
              />
            </b-field>
            <b-field
              label="Mpbile Number"
              custom-class="is-small"
              :type="errors.has('Mpbile Number') ? 'is-danger': ''"
              :message="errors.first('Mpbile Number')"
            >
              <b-input
                v-model="form.mobile_number"
                expanded
                name="Mpbile Number"
                placeholder="Mpbile Number"
                v-validate="'required'"
              />
            </b-field>
            <b-field
              label="Email"
              custom-class="is-small"
              :type="errors.has('Email') ? 'is-danger': ''"
              :message="errors.first('Email')"
            >
              <b-input
                v-model="form.email"
                expanded
                name="Email"
                placeholder="Email"
                v-validate="'required'"
              />
            </b-field>
            <b-field
              label="Message"
              custom-class="is-small"
              :type="errors.has('Message') ? 'is-danger': ''"
              :message="errors.first('Message')"
            >
              <b-input
                v-model="form.message"
                expanded
                type="textarea"
                name="Message"
                placeholder="Message"
                v-validate="'required'"
              />
            </b-field>
            <button class="button is-bg-red has-text-white">Login</button>
            <!-- <b-button class="button is-bg-red has-text-white">Submit</b-button> -->
          </form>
        </div>
      </div>
    </div>
    <div class="columns">
      <div class="column has-text-centered">
        <img class src="/img/banner/white-logo.png" alt />
        <p>Copyright &#9400; 2019. Prudent Domiciliary Care Ltd. All Rights Reserved.</p>
      </div>
    </div>
  </div>
</template>
<script>
import { Notification } from "@/utils/helpers";
export default {
  data() {
    return {
      form: {},
      Navigation: [
        { name: "Abouts Us", link: "about" },
        { name: "Services", link: "services" },
        { name: "Process", link: "process" },
        { name: "FaQs", link: "faqs" },
        { name: "Contact Us", link: "contact" },
        { name: "Join Us", link: "join-us" }
      ],

      Services: [
        { name: "Companionship Care", link: "Companionship-care" },
        { name: "Palliative Care", link: "Palliative-care" },
        {
          name: "Emphatic End of Life Care",
          link: "emphatic-end-of-life-care"
        },
        { name: "Long Term Home Care", link: "long-term-home-care" },
        { name: "Dementia Care at Home", link: "dementia-care-at-home" },
        { name: "Home from Hospital", link: "home-from-hospital" },
        { name: "Take a break", link: "take-a-break" },
        { name: "Live in Care", link: "live-in-care" },
        {
          name: "Physical Disability Support at Home",
          link: "Physical-disability-Support-at-home"
        },
        {
          name: "Supporting Independent Living",
          link: "supporting-independent-living"
        },
        { name: "Learning Disabilities", link: "learning-disabilities" }
      ]
    };
  },
  methods: {
    async submit() {
      let result = await this.$validator.validateAll();
      if (result) {
        try {
          let param = {
            url: "contacts",
            data: this.form
          };
          let req = await this.$store.dispatch("Postdata", param);
          if (req) {
            this.isLoading = false;
            Notification(
              this,
              "Thank you for filling the form",
              "is-success",
              "is-top",
              1000
            );
            this.form = {};
          } else {
            this.isLoading = false;
            Notification(
              this,
              "An error occured. Please try again ",
              "is-danger",
              "is-top",
              1000
            );
          }
        } catch (error) {
          this.isLoading = false;

          Notification(
            this,
            "An error occured. Please try again ",
            "is-danger",
            "is-top",
            1000
          );
        }
      }
    }
  }
};
</script>

<style lang="scss" scoped></style>

<template>
  <div class="contact">
    <div class="pg-title is-blue-bold is-size-3">We Care, Let's talk</div>
    <div class="contact-content">
      <div class="columns">
        <div class="column">
          <div class="contact-container">
            <div class="head-office is-blue">Head Office</div>
            <p class="address">130 Old Street, London EC1V 9BD</p>
            <a
              class="email is-blue-bold is-lowercase"
              href="mailto:info@pbgcare.co.uk"
            >info@pbgcare.co.uk</a>
          </div>
        </div>
        <div class="column is-7">
          <div class="head-office is-blue">Write Us</div>

          <div class="pg-form">
            <form @submit.prevent="submit">
              <b-field
                label="Name"
                custom-class="is-small"
                :type="errors.has('name') ? 'is-danger': ''"
                :message="errors.first('name')"
              >
                <b-input
                  v-model="form.name"
                  expanded
                  name="name"
                  placeholder="Title"
                  v-validate="'required'"
                />
              </b-field>
              <b-field
                label="Email"
                custom-class="is-small"
                :type="errors.has('email') ? 'is-danger': ''"
                :message="errors.first('email')"
              >
                <b-input
                  v-model="form.email"
                  expanded
                  name="email"
                  placeholder="Email"
                  v-validate="'required'"
                />
              </b-field>
              <b-field
                label="Mobile Number"
                custom-class="is-small"
                :type="errors.has('Mobile Number') ? 'is-danger': ''"
                :message="errors.first('Mobile Number')"
              >
                <b-input
                  v-model="form.mobile_number"
                  expanded
                  name="Mobile Number"
                  placeholder="Mobile Number"
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
              <div class="field">
                <b-checkbox
                  v-validate="'required'"
                  v-model="checkbox"
                  name="checkbox"
                  class="is-grey"
                >{{ checkboxmsg }}</b-checkbox>
              </div>
              <button class="button is-bg-red has-text-white is-size-6 is-rounded">Send Message</button>

              <!-- <b-button type="is-bg-red has-text-white is-size-6" outlined rounded>Send Message</b-button> -->
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { Notification } from "@/utils/helpers";

export default {
  head() {
    return {
      title: "Contact"
    };
  },
  async asyncData({ store }) {
    store.commit("setbanner", "contact-us-image.jpg");
  },
  data() {
    return {
      checkbox: false,
      checkboxmsg:
        "By using this form you agree with the storage and handling of your data on this website",
      form: {}
    };
  },
  methods: {
    async submit() {
      let result = await this.$validator.validateAll();

      if (result) {
        if (!this.checkbox) {
          Notification(
            this,
            "Please select the checkbox ",
            "is-danger",
            "is-top",
            1000
          );
          return;
        }
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

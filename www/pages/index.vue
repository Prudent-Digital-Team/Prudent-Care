<template>
  <div class="homepage">
    <nav class="level">
      <div class="level-left"></div>
      <!-- Left side -->
      <div class="level-right">
        <div class="level-item">
          <img src="/img/banner/Dementia-Care-logo.jpg" alt />
        </div>
      </div>
    </nav>
    <div class="columns is-8 is-variable">
      <div class="column">
        <div class="home-img-box" :style="{'background-image':`url(${image})`}">
          <div class="box-content">
            <p class="box-title has-text-centered has-text-white">About Us</p>
            <nuxt-link to="/about">
              <b-button type="is-white is-size-5" outlined rounded>Learn More</b-button>
            </nuxt-link>
          </div>
        </div>
      </div>
      <div class="column">
        <div class="home-img-box" :style="{'background-image':`url(${image})`}">
          <div class="box-content">
            <p class="box-title has-text-centered has-text-white">Care Process</p>
            <nuxt-link to="/process">
              <b-button type="is-white is-size-5" outlined rounded>Learn More</b-button>
            </nuxt-link>
          </div>
        </div>
      </div>
      <div class="column">
        <div class="home-img-box" :style="{'background-image':`url(${image})`}">
          <div class="box-content">
            <p class="box-title has-text-centered has-text-white">Our Services</p>
            <nuxt-link to="/services">
              <b-button type="is-white is-size-5" outlined rounded>Learn More</b-button>
            </nuxt-link>
          </div>
        </div>
      </div>
    </div>
    <section class="testimonials">
      <div class="testimonial-title is-red">Clients Testimonials</div>
      <div class="testimonial-container">
        <div class="columns">
          <div class="column" v-for="testimony in Testimonials" :key="testimony.id">
            <div class="box testimonial-box">
              <div class="testimonial-text is-capitalized">{{testimony.content}}</div>
              <div class="testimonial-author is-red">{{testimony.author}}</div>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  layout: "home",
  head() {
    return {
      title: "Homepage"
    };
  },
  async asyncData({ store }) {
    store.commit("setbanner", "home.jpg");
  },
  data() {
    return {
      Testimonials: [
        {
          content: `from when i made contact with prudent domiciliary care for
                my mother, the manager came out to assist her needs, the care
                started the very next day. the team were fantastic at caring for my
                mother, i dont have any worry any more.`,
          author: "Mrs. E.Roberts"
        },
        {
          content: `The staff at prudent Domiciliary care are caring and kingand completely comfortable with them. The team and my
            carer does a wonderful job`,
          author: "Mr. M.Hards"
        },
        {
          content: `My carers are always on time and very professional. They also make sure my care is right for me by doing 
          my care needs assesment. They are all so good. i wouldn't be without them`,
          author: "Mrs. D.Smith"
        }
      ],
      image: "/img/banner/about.jpg"
    };
  }
};
</script>

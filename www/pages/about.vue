<template>
  <div class="about">
    <div class="pg-title is-blue-bold">About Us</div>
    <div class="pg-content" v-html="content">
      <!-- <p>
        At Prudent Domiciliary Care, we provide elderly citizens with quality
        care that enables them to live happier, healthier lives at home. Our
        services are distinguished by the calibre of our caregivers, the
        responsiveness of our staff and our expertise in daily care. We embrace
        a positive, balanced approach to aging centred on the evolving needs of
        older adults.
      </p>-->
      <!-- <div class="about-List" v-for="about in AboutList" :key="about.id">
        <div class="about-title is-red">{{ about.title }}</div>
        <div class="about-content">{{ about.content }}</div>
      </div>-->
    </div>
  </div>
</template>

<script>
export default {
  head() {
    return {
      title: "About"
    };
  },

  async asyncData({ store }) {
    store.commit("setbanner", "about.jpg");
  },
  data() {
    return {
      content: `<p>At Prudent Domiciliary Care, we provide elderly citizens with quality care that enables them to live happier, healthier lives at home. Our services are distinguished by the calibre of our caregivers, the responsiveness of our staff and our expertise in daily care. We embrace a positive, balanced approach to aging centred on the evolving needs of older adults.</p>
<p><span style="color: #ff0000;">Affordable</span></p>
<p>Prudent Domiciliary Care are designed to meet all your needs in a very special and personalized approach. We believe that Quality Home Care should be available to everyone. You or your loved ones deserve the home care which is meeting all your needs and being cost effective at the same time</p>
<p><span style="color: #ff0000;">Accessible</span></p>
<p>Prudent Domiciliary Care is a driven organization. Our Home Care team is very proactive and involved in all aspects of your care needs. We are always accessible whichever is convenient to you. Your concerns, likes and wishes are a priority for us. Contact details for Home Care Team:</p>
<p><span style="color: #ff0000;">Accountable</span></p>
<p>Prudent Domiciliary Care is a driven organization. Our Home Care team is very proactive and involved in all aspects of your care needs. We are always accessible whichever is convenient to you. Your concerns, likes and wishes are a priority for us.Contact details for Home Care Team</p>`,
      AboutList: [
        {
          title: "Affordable",
          content:
            "Prudent Domiciliary Care are designed to meet all your needs in a very special and personalized approach. We believe that Quality Home Care should be available to everyone. You or your loved ones deserve the home care which is meeting all your needs and being cost effective at the same time"
        },
        {
          title: "Accessible",
          content:
            "Prudent Domiciliary Care is a driven organization. Our Home Care team is very proactive and involved in all aspects of your care needs. We are always accessible whichever is convenient to you. Your concerns, likes and wishes are a priority for us. Contact details for Home Care Team:"
        },
        {
          title: "Accountable",
          content:
            "Prudent Domiciliary Care is a driven organization. Our Home Care team is very proactive and involved in all aspects of your care needs. We are always accessible whichever is convenient to you. Your concerns, likes and wishes are a priority for us.Contact details for Home Care Team"
        }
      ]
    };
  }
};
</script>

<style lang="scss" scoped></style>

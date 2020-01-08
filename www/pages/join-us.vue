<template>
  <div class="join-us">
    <div class="pg-title is-blue-bold">Join Us</div>
    <div class="join-content" v-html="Joinpg"></div>

    <div class="join-us-email is-blue">
      All Applications should be sent to
      <a
        class="email is-blue-bold is-lowercase"
        href="mailto:info@pbgcare.co.uk"
      >info@pbgcare.co.uk</a>
    </div>
  </div>
</template>

<script>
export default {
  head() {
    return {
      title: "Join Us"
    };
  },
  async asyncData({ store }) {
    store.commit("setbanner", "join-us.jpg");
  },
  data() {
    return {
      Joinpg: `<p><strong>If you are seeking a role with real job satisfaction, join us now and benefit from our excellent rates of pay of &pound;9.03 (flat rate) - &pound;9.45 (unsociable hour rate) per hour.</strong></p>
<br>
<p>Prudent Domiciliary Care is the leading provider of home care in Kent and Royal Borough of Greenwich, providing services throughout Charlton, Plumstead and Abbey Wood, Belvedere, Sidcup. we provide high quality, professional care to people in their own homes including the following:</p>
<p>&nbsp;</p>
<p><strong>Responsibilities</strong></p>
<ul>
<li>Personal care and helping clients to wash and dress</li>
<li>Assistance with medication</li>
<li>Shopping and support with everyday tasks</li>
<li>Light household chores</li>
<li>Meal preparation</li>
</ul>
<p>&nbsp;</p>
<p><strong>Requirements</strong></p>
<p>Whilst previous care experience is helpful, it is not essential, as we will provide you with the training you require to carry out your role with confidence and a support team who you can access at any time. We are seeking people who are:</p>
<ul>
<li>Kind, caring and wish to make a difference</li>
<li>Reliable and take pride in their work</li>
<li>Have a full UK driving licence and own a vehicle for use at work</li>
<li>Flexible in their approach and able to work alternate weekends</li>
<li>Available to take part in our introductory 4 day training programme</li>
</ul>`
    };
  }
};
</script>

<style lang="scss" scoped>
</style>
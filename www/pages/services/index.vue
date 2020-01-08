<template>
  <div class="services">
    <div class="pg-title is-blue-bold">Our Services</div>
    <div class="pg-services-content" v-html="main">
      <!-- <p v-for="service in ServicesList" :key="service.id" class="servicesList">
        <span class="mdi mdi-circle-medium is-red is-glight"></span>
        <span class="content">{{ service.content }}</span>
      </p>-->
    </div>
    <nuxt-link to="/contact">
      <b-button type="is-bg-red has-text-white is-size-6" outlined rounded>GET IN TOUCH WITH US</b-button>
    </nuxt-link>
  </div>
</template>

<script>
export default {
  data() {
    return {
      main: `
      <ul style="list-style-type: disc;">
<li>&nbsp;Assist with a &lsquo;good morning&rsquo; start to the day, helping you to get up, wash, shower or bathe, dress and have breakfast</li>
<li>Remind or assist you to take your medicines and collect or return medication from your pharmacy or dispensing GP surgery</li>
<li>Prepare meals with or for you and assist you at mealtimes</li>
<li>Collect your pension for you or with you</li>
<li>Shop with you or help you to make a shopping list, go to the shops, come back and put it all away</li>
<li>Help with your household tasks such as laundry or ironing and keeping your home clean and tidy</li>
<li>Support you with social activities such as going out for a walk, attending a day centre, visiting friends or family, going to your church or club etc</li>
<li>Give a little pampering when you need it &ndash; the occasional visit to organise bathing, clean clothes, fresh bedding and a thoughtfully prepared meal before bedtime</li>
<li>At the end of each day, some help with getting ready for bed</li>
</ul>
<p><strong>Our Services also includes</strong></p>
<ul style="list-style-type: disc;">
<li><a href="https://pbgcare.co.uk/services/Companionship-care">Companionship Care</a></li>
<li><a href="https://pbgcare.co.uk/services/Palliative-care">Palliative Care</a></li>
<li><a href="https://pbgcare.co.uk/services/emphatic-end-of-life-care">Emphatic End Of Life Care</a></li>
<li><a href="https://pbgcare.co.uk/services/long-term-home-care">Long Term Home Care</a></li>
<li><a href="https://pbgcare.co.uk/services/dementia-care-at-home">Dementia Care At Home</a></li>
<li><a href="https://pbgcare.co.uk/services/home-from-hospital">Home From Hospital</a></li>
<li><a href="https://pbgcare.co.uk/services/take-a-break">Take A Break</a></li>
<li><a href="https://pbgcare.co.uk/services/live-in-care">Live In Care</a></li>
<li><a href="https://pbgcare.co.uk/services/Physical-disability-Support-at-home">Physical Disability Support At Home</a></li>
<li><a href="https://pbgcare.co.uk/services/supporting-independent-living">Supporting Independent Living</a></li>
<li><a href="https://pbgcare.co.uk/services/learning-disabilities">Learning Disabilities</a></li>
</ul>`,
      serve: `<p><strong>Our Services also includes</strong></p>
<ul>
<li><a href="https://pbgcare.co.uk/services/Companionship-care">Companionship Care</a></li>
<li><a href="https://pbgcare.co.uk/services/Palliative-care">Palliative Care</a></li>
<li><a href="https://pbgcare.co.uk/services/emphatic-end-of-life-care">Emphatic End Of Life Care</a></li>
<li><a href="https://pbgcare.co.uk/services/long-term-home-care">Long Term Home Care</a></li>
<li><a href="https://pbgcare.co.uk/services/dementia-care-at-home">Dementia Care At Home</a></li>
<li><a href="https://pbgcare.co.uk/services/home-from-hospital">Home From Hospital</a></li>
<li><a href="https://pbgcare.co.uk/services/take-a-break">Take A Break</a></li>
<li><a href="https://pbgcare.co.uk/services/live-in-care">Live In Care</a></li>
<li><a href="https://pbgcare.co.uk/services/Physical-disability-Support-at-home">Physical Disability Support At Home</a></li>
<li><a href="https://pbgcare.co.uk/services/supporting-independent-living">Supporting Independent Living</a></li>
<li><a href="https://pbgcare.co.uk/services/learning-disabilities">Learning Disabilities</a></li>
</ul>`,
      ServicesList: [
        {
          content:
            "Assist with a ‘good morning’ start to the day, helping you to get up, wash, shower or bathe, dress and have breakfast"
        },
        {
          content:
            "Remind or assist you to take your medicines and collect or return medication from your pharmacy or dispensing GP surgery"
        },
        {
          content: "Prepare meals with or for you and assist you at mealtimes"
        },
        {
          content: "Collect your pension for you or with you"
        },
        {
          content:
            "Shop with you or help you to make a shopping list, go to the shops, come back and put it all away"
        },
        {
          content:
            "Help with your household tasks such as laundry or ironing and keeping your home clean and tidy"
        },
        {
          content:
            "Support you with social activities such as going out for a walk, attending a day centre, visiting friends or family, going to your church or club etc"
        },
        {
          content:
            "Give a little pampering when you need it – the occasional visit to organise bathing, clean clothes, fresh bedding and a thoughtfully prepared meal before bedtime"
        },
        {
          content:
            "At the end of each day, some help with getting ready for bed"
        }
      ]
    };
  },
  head() {
    return {
      title: "Our Services"
    };
  },
  async asyncData({ store }) {
    store.commit("setbanner", "Our-Services-Page-Image.jpg");
  }
};
</script>

<style lang="scss" scoped></style>

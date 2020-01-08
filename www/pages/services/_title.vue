<template>
  <div>
    <div class="services-children">
      <div class="pg-title is-blue-bold">{{ pg.name }}</div>
      <div class="service-content" v-html="pg.content"></div>
      <b-button
        class="service-button"
        type="is-bg-red has-text-white is-size-6"
        outlined
        rounded
        >GET IN TOUCH WITH US</b-button
      >
    </div>
    <!-- <div class="services-children">
      <div class="pg-title is-blue-bold">Palliative Care</div>
      <div class="service-content">
        We understand the emotional pain and practical pressures when someone in
        your family has a life-threatening illness. Our fully trained care
        assistants offer sensitive care and support to help you and your loved
        ones through the difficult times. Your Prudent domiciliary Care
        palliative care team can provide the emotional and practical support you
        need at any time, from when you receive a diagnosis until the final
        months. We agree a care plan that is tailored to your personal needs and
        choices. Our dedicated care assistants support our customers to be as
        comfortable as possible in the familiar surroundings of their own home.
      </div>
      <b-button
        class="service-button"
        type="is-bg-red has-text-white is-size-6"
        outlined
        rounded
        >GET IN TOUCH WITH US</b-button
      >
    </div> -->
  </div>
</template>

<script>
export default {
  layout: "services",

  async asyncData({ store, params }) {
    let serviceList = store.state.ServiceList;
    let title = params.title.toLowerCase();
    let pg = _.find(serviceList, s => s.link == title);
    store.commit("setbanner", pg.image);
    console.log(pg);
    return { pg };
  },
  head() {
    return {
      title: this.pg.name
    };
  }
};
</script>

<style lang="scss" scoped></style>

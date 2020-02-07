<template>
  <div class="meal">
    <div class="pg-title is-blue-bold">{{page.attributes.content.title}}</div>

    <div class="meal-text">
      <div class="columns">
        <div class="column">
          <div v-html="page.attributes.content.content"></div>
        </div>
        <div
          class="column is-3 meal-img"
          :style="{ 'background-image': `url(${page.attributes.content.image.data})` }"
        >
          <div></div>
        </div>
      </div>
      <p>We offer two main categories:</p>
    </div>
    <div class="option-section">
      <div class="menu" v-for="(page,idx) in page.attributes.content.List" :key="idx">
        <div class="menu-contain">
          <div class="menu-title">OPTION {{idx+1}}</div>
          <div class="menu-content" v-html="page.content"></div>
        </div>
      </div>
    </div>
    <div class="main-menu">
      <div class="main-menu__title">MENU</div>
      <div class="main-menu__subtitle">(GUIDELINES ONLY)</div>
      <div class="main-menu__menu-tags">
        <div class="main-menu__list-horizontal">
          <ul>
            <li v-for="list in page.attributes.content.menu.List" :key="list.id">
              <span class="bull">&bull;</span>
              {{list}}
            </li>
          </ul>
        </div>
        <!-- <b-tag
          class="is-blue"
          rounded
          multilined
          v-for="tag in page.attributes.content.menu.List"
          :key="tag.id"
        >{{getTag(tag)}}</b-tag>-->
      </div>
    </div>
  </div>
</template>


<script>
import PageMixin from "@/mixins/index";

export default {
  mixins: [PageMixin],
  methods: {
    getTag(val) {
      if (val.length > 50) {
        let first = val.split(/(.{100}\S*)\s/).filter(function(e) {
          return e;
        });
        let newString = first.map(a => `${a} \n`).join(" ");
        return newString;
      } else {
        return val;
      }
    }
  },

  head() {
    return {
      title: "Prep Meal",
      meta: [
        // hid is used as unique identifier. Do not use `vmid` for it as it will not work
        {
          hid: "title",
          name: "title",
          content:
            "Prudent Domiciliary Care Our Meal. Offering great and nutricious food to our clients"
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
};
</script>


<style lang="scss" scoped></style>

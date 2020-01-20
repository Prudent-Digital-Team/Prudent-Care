<template>
  <div class="blog">
    <div class="caption-blog">
      <div class="bcontent">
        <nuxt-link :to="`/blog/${featured_blog.slug}`" href>
          <div class="columns">
            <div
              class="column left is-paddingless is-8"
              :style="{ 'background-image': `url(${featured_blog.image.data})` }"
            ></div>
            <div class="column">
              <div class="right">
                <div class="sub-title">{{getCategory(featured_blog.category_id)}}</div>
                <div class="title">{{featured_blog.title}}</div>
                <div class="content" v-html="featured_blog.content"></div>
                <div class="date">{{ featured_blog.created_at | moment('MMMM DD, YYYY') }}</div>
              </div>
            </div>
          </div>
        </nuxt-link>
      </div>
    </div>
    <section class="blog-content" v-if="blogList.length != 0">
      <div class="columns is-variable is-4-desktop is-multiline">
        <div class="column is-4" v-for="blog in blogList" :key="blog.id">
          <div class="box blog-container">
            <p class="blog-author">
              <span class="category">{{getCategory(blog.category_id)}}</span>
              <span class="divider">&nbsp;|&nbsp;</span>
              <span class="datee">{{ blog.created_at | moment('MMMM DD, YYYY') }}</span>
            </p>
            <div
              class="blog-image"
              :style="{ 'background-image': `url(/static/uploads/bobnoiq23akqr63riar0.jpg)` }"
            ></div>
            <div class="blog-content">
              <h1 class="title">{{blog.title}}</h1>
              <div class="bg-text" v-html="blog.content"></div>
            </div>
            <div class="right">
              <nuxt-link class="blog-link" :to="'/blog/'+blog.slug">READ MORE</nuxt-link>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import PageMixin from "@/mixins/blog.index";
export default {
  mixins: [PageMixin],
  layout: "blog",
  methods: {
    getCategory(id) {
      let category = this.services.find(c => c.id == id);

      return category.name;
    }
  }
};
</script>

<style lang="scss" scoped>
</style>
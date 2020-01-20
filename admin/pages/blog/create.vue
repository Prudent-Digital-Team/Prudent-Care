<template>
  <div>
    <custom-container :hide="false" :hide1="true">
      <custom-box :title="'Create Service'">
        <div class="columns">
          <div class="column is-9">
            <b-field
              label="Title"
              custom-class="is-small"
              :type="errors.has('title') ? 'is-danger' : ''"
              :message="errors.first('title')"
            >
              <b-input
                expanded
                v-model="form.title"
                name="title"
                placeholder="Title"
                v-validate="'required'"
              />
            </b-field>
          </div>
          <div class="column">
            <b-field
              label="Author"
              custom-class="is-small"
              :type="errors.has('author') ? 'is-danger' : ''"
              :message="errors.first('author')"
            >
              <b-input
                expanded
                v-model="form.author"
                name="author"
                placeholder="Author"
                v-validate="'required'"
              />
            </b-field>
          </div>
        </div>
        <div class="columns">
          <div class="column is-9">
            <b-field
              label="Slug Name"
              custom-class="is-small"
              :type="errors.has('Slug Name') ? 'is-danger' : ''"
              :message="errors.first('Slug Name')"
            >
              <b-input
                v-model="slugUrl"
                expanded
                name="Slug Name"
                placeholder="Slug Name"
                v-validate="'required'"
              />
            </b-field>
          </div>
          <div class="column">
            <b-field
              label="Category"
              custom-class="is-small"
              :type="errors.has('Category') ? 'is-danger': ''"
              :message="errors.first('Category')"
            >
              <b-select
                v-model="form.category_id"
                expanded
                name="Category"
                placeholder="select a category"
                v-validate="'required'"
              >
                <option
                  :value="option.id"
                  v-for="(option, index) in services.data"
                  :key="index"
                >{{ option.name }}</option>
              </b-select>
            </b-field>
          </div>
        </div>
        <image-upload label="Header Image" v-model="form.image" @image="isImageValid($event)" />
        <b-field
          label="Content"
          custom-class="is-small"
          :type="errors.has('content') ? 'is-danger' : ''"
          :message="errors.first('content')"
        >
          <custom-editor name="content" v-validate="'required'" v-model="form.content"></custom-editor>
        </b-field>
      </custom-box>
      <div class="columns">
        <div class="column is-5">
          <nuxt-link to="/blog">
            <button class="button has-text-centered is-danger is-fullwidth">cancel</button>
          </nuxt-link>
        </div>
        <div class="column">
          <button
            @click="submit"
            class="button has-text-centered is-bg-blue has-text-white is-fullwidth"
          >Submit</button>
        </div>
      </div>
    </custom-container>
  </div>
</template>

<script>
import CustomDropdown from '@/components/widgets/Dropdown'
import PagesMix from '@/mixins/page.create'
import { Slugify } from '@/utils/helpers'

export default {
  mixins: [PagesMix],
  components: {
    CustomDropdown
  },
  data() {
    return {
      slugUrl: ''
    }
  },
  watch: {
    'form.title': function(val, oldval) {
      this.slugUrl = Slugify(val)
    },
    slugUrl(newval, oldval) {
      let slug = Slugify(newval)
      this.slugUrl = slug
      this.form.slug = slug
    }
  }
}
</script>

<style lang="scss" scoped></style>

<template>
  <div>
    <section class="custom-dropdown">
      <b-collapse
        class="card"
        :class="displayError ? 'custom-help' : ''"
        :key="index"
        :open="isOpen == index"
        @open="isOpen = index; displayError = false"
        @close="displayError = false"
      >
        <div slot="trigger" slot-scope="props" class="card-header" role="button">
          <p class="card-header-title">Item: {{ index + 1 }}</p>
          <a class="card-header-icon">
            <b-icon :icon="props.open ? 'menu-down' : 'menu-up'"></b-icon>
          </a>
        </div>
        <div class="card-content">
          <div class v-if="!EditMode">
            <div class="level">
              <div class="level-right level-item">
                <button
                  class="button is-danger is-small has-text-right is-outlined print-invoice-button"
                  @click="delSection(form)"
                >Delete</button>
              </div>
            </div>
          </div>
          <div class="columns">
            <div class="column">
              <b-field
                label="Content"
                custom-class="is-small"
                :type="errors.has('content') ? 'is-danger' : ''"
                :message="errors.first('content')"
              >
                <custom-editor
                  name="content"
                  :disabled="EditMode"
                  v-validate="'required'"
                  v-model="form.content"
                ></custom-editor>
              </b-field>
            </div>
          </div>
        </div>
      </b-collapse>
      <span
        class="displayDanger help is-danger"
        v-if="displayError"
      >There is pending data thats needs to save.</span>
    </section>
  </div>
</template>

<script>
import ImageUpload from '@/components/widgets/ImageUpload.vue'
import CustomEditor from '@/components/widgets/CustomEditor'

export default {
  components: { ImageUpload, CustomEditor },
  props: {
    total: {
      type: Number,
      default: () => 0
    },
    index: {
      type: Number,
      default: () => 0
    },
    itemIdx: {
      type: Number,
      default: () => 0
    },
    EditMode: {
      type: Boolean,
      default: () => false
    },
    form: {
      type: Object,
      default: () => {}
    },
    pagesList: {
      type: Array,
      default: () => []
    }
  },
  data() {
    return {
      isOpen: false,
      displayError: false,
      isImage: false
    }
  },
  methods: {
    delSection(obj) {
      this.$emit('deleteSection', obj)
    },
    // delSection(obj) {
    //   let catalog = this.form
    //   let delobj = _.find(catalog, obj)
    //   catalog.splice(catalog.indexOf(delobj), 1)
    // },
    isImageValid(validity) {
      this.isImage = validity
    },
    async isValid() {
      let result = await this.$validator.validateAll()
      if (!result) {
        return
      }
      this.$emit('input', copy(this.form))
      return true
    },
    emitForm() {
      this.$emit(this.form)
    },
    async validate() {
      this.displayError = false

      let result = await this.$validator.validateAll()
      if (!result) {
        this.displayError = true
        return false
      }
      return true
    }
  }
}
</script>

<style lang="scss" scoped>
</style>
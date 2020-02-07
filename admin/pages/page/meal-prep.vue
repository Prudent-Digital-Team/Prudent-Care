<template>
  <div class="about">
    <custom-container @save="save" @edit="toggleEdit" @cancel="cancel" :hideCreate="true">
      <custom-box :title="'Section: Page Header'">
        <image-upload
          label="Header Image"
          :readonly="EditMode"
          v-model="pageData.attributes.header.image"
          @image="isImageValid($event)"
        />
      </custom-box>
      <custom-box :title="`Section: ${pageData.attributes.content.title}`">
        <b-field
          label="Title"
          custom-class="is-small"
          :type="errors.has('title') ? 'is-danger' : ''"
          :message="errors.first('title')"
        >
          <b-input
            expanded
            v-model="pageData.attributes.content.title"
            name="title"
            :disabled="EditMode"
            placeholder="Title"
            v-validate="'required'"
          />
        </b-field>
        <image-upload
          label="Image"
          :readonly="EditMode"
          v-model="pageData.attributes.content.image"
          @image="isImageValid($event)"
        />

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
            v-model="pageData.attributes.content.content"
          ></custom-editor>
        </b-field>
      </custom-box>

      <custom-box :title="`Section: 3 (Options)`">
        <template v-slot:button>
          <button class="button is-blue" :disabled="EditMode" @click="addBlock">add new block</button>
        </template>

        <custom-dropdown
          ref="custom_dropdown"
          v-for="(form, idx) in pageData.attributes.content.List"
          :key="idx"
          :index="idx"
          @deleteSection="delSection($event)"
          :form="form"
          :EditMode="EditMode"
        ></custom-dropdown>
      </custom-box>

      <custom-box :title="`Section: 4 (Menu)`">
        <template v-slot:button>
          <button class="button is-blue" :disabled="EditMode" @click="addTag">add new menu</button>
        </template>

        <!-- <b-field
          label="Description"
          custom-class="is-small"
          :type="errors.has('description') ? 'is-danger' : ''"
          :message="errors.first('description')"
        >
          <custom-editor
            name="content"
            :disabled="EditMode"
            v-validate="'required'"
            v-model="pageData.attributes.content.menu.description"
          ></custom-editor>
        </b-field> -->

        <b-collapse class="card">
          <div
            slot="trigger"
            slot-scope="props"
            class="card-header"
            role="button"
            aria-controls="contentIdForA11y3"
          >
            <p class="card-header-title">Menu Items</p>
            <a class="card-header-icon">
              <b-icon :icon="props.open ? 'menu-down' : 'menu-up'"></b-icon>
            </a>
          </div>
          <div class="card-content">
            <section>
              <b-tag
                rounded
                :closable="!EditMode"
                v-for="tag in pageData.attributes.content.menu.List"
                @close="rmvTag(tag)"
                :key="tag.id"
              >{{tag}}</b-tag>
            </section>
          </div>
        </b-collapse>
      </custom-box>
    </custom-container>
  </div>
</template>

<script>
import PagesMix from '@/mixins/admin_pages'
import CustomEditor from '@/components/widgets/CustomEditor'
import CustomDropdown from '@/components/widgets/mealprep/MealDropdown'
import TagModal from '@/components/widgets/mealprep/MealTagModal'

import ServiceEditor from '@/components/widgets/services/EditorBlock'
import { copy, clone, Notification } from '@/utils/helpers'

export default {
  mixins: [PagesMix],
  components: {
    CustomEditor,
    ServiceEditor,
    CustomDropdown
  },
  methods: {
    delSection(obj) {
      let catalog = this.pageData.attributes.content.List
      let delobj = _.find(catalog, obj)
      catalog.splice(catalog.indexOf(delobj), 1)
    },
    addBlock() {
      if (!this.EditMode) {
        let form = this.pageData.attributes.content.List
        let block = { title: '', content: '' }
        form.push(block)
      }
    },
    addTag() {
      let that = this
      let form = this.pageData.attributes.content.menu.List

      this.$buefy.modal.open({
        parent: this,
        component: TagModal,
        hasModalCard: true,
        trapFocus: true,
        events: {
          input(val) {
            form.push(val)
          }
        }
      })
    },
    rmvTag(tag) {
      let catalog = this.pageData.attributes.content.menu.List
      let delobj = _.find(catalog, tag)
      catalog.splice(catalog.indexOf(delobj), 1)
    },
    cancel() {
      this.$router.push('/page')
    },
    editService() {
      let that = this
      this.$buefy.modal.open({
        parent: this,
        component: ServiceEditor,
        hasModalCard: true,
        trapFocus: true,
        events: {
          input(val) {}
        }
      })
    },

    async save() {
      let ref = this.$refs.custom_dropdown
      let validated = true

      for (let a = 0; a < ref.length; a++) {
        let x = await ref[a].validate()
        validated = x && validated
      }

      let result = await this.$validator.validateAll()
      if (!this.isImage) {
        Notification(this, 'Please upload image.', 'is-danger')
        return
      }

      if (result) {
        if (validated) {
          let form = copy(this.pageData)
          let param = {
            url: `pages/${this.uuid}`,
            data: form
          }
          let req = await this.$store.dispatch('Postdata', param)
          if (req.done) {
            Notification(this, 'Changes saved successfully', 'is-success')
            this.$router.push('/page')
          } else {
            Notification(this, 'an error occured', 'is-danger')
          }
        }
      }
    }
  }
}
</script>

<style lang="scss" scoped></style>

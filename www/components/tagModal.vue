<template>
  <div>
    <div class="modal-card" style="width: auto">
      <section class="modal-card-body">
        <div class="meal-tag-form">
          <b-field
            label="Menu Name"
            custom-class="is-small"
            :type="errors.has('Menu Name') ? 'is-danger' : ''"
            :message="errors.first('Menu Name')"
          >
            <b-input
              expanded
              v-model="title"
              name="Menu Name"
              placeholder="Menu Name"
              v-validate="'required'"
            />
          </b-field>
        </div>
      </section>
      <footer class="modal-card-foot">
        <div class="modal-level-right">
          <button class="button" type="button" @click="$parent.close()">Cancel</button>
          <button class="button is-danger" @click="save">save</button>
        </div>
      </footer>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      title: ""
    };
  },
  methods: {
    async save() {
      let result = await this.$validator.validateAll();
      if (result) {
        this.$emit("input", this.title);
        this.$parent.close();
      }
    }
  }
};
</script>

<style lang="scss" scoped>
.meal-tag-form {
  width: 50rem;
  padding: 1rem;
}
</style>
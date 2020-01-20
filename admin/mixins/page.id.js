import { copy, Notification } from '@/utils/helpers'
import CustomContainer from '@/components/widgets/CustomContainer'
import CustomBox from '@/components/widgets/CustomBox'
import ImageUpload from '@/components/widgets/ImageUpload.vue'
import { mapState } from 'vuex'

export default {
  components: {
    CustomContainer,
    CustomBox,
    ImageUpload
  },
  async asyncData({ store, route, params }) {
    let back = route
    let Result = []

    let apiRoute = route.path.split('/')[1]
    let pgname = params.id
    store.commit('commitpageTitle', `View ${apiRoute}`)
    let asyncReq = [
      { name: 'form', url: `${apiRoute}/${pgname}` },
      { name: 'serviceList', url: 'services/all' }
    ]
    for (let a = 0; a < asyncReq.length; a++) {
      Result[asyncReq[a].name] = await store.dispatch('Get', {
        url: asyncReq[a].url
      })
      Result[asyncReq[a].name] = Result[asyncReq[a].name].store
    }
    return {
      form: Result.form,
      services: Result.serviceList,
      apiRoute: apiRoute,
      pgname: pgname
    }
  },
  data() {
    return {}
  },
  computed: {
    ...mapState(['EditMode', 'pagesList'])
  },
  methods: {
    async save() {
      let result = await this.$validator.validateAll()
      if (result) {
        let param = {
          url: `${this.apiRoute}/${this.form.uuid}`,
          data: this.form
        }
        let req = await this.$store.dispatch('Postdata', param)
        if (req.done) {
          Notification(this, 'Successfully Edited Post.', 'is-success')
          setTimeout(() => this.$router.push(`/${this.apiRoute}`), 1000)
        }
      }
    },
    toggleEdit() {
      this.$store.commit('commitEdit', !this.EditMode)
    },
    cancel() {
      this.$router.push(`/${this.apiRoute}`)
    },

    toggleEdit() {
      this.$store.commit('commitEdit', !this.EditMode)
    }
  }
}

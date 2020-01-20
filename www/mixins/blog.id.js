// import { Notification, Notify } from '@/utils/helpers'
import { copy } from "@/utils/helpers";
import { mapState } from "vuex";
import _ from "lodash";

export default {
  async asyncData({ store, params }) {
    let title = params.title;
    let Result = [];
    let asyncReq = [
      { name: "blogItem", url: `blog/${title}` },
      { name: "blogList", url: "blog?page=1&perPage=6" },
      { name: "serviceList", url: "services/all" }
    ];
    for (let a = 0; a < asyncReq.length; a++) {
      Result[asyncReq[a].name] = await store.dispatch("Get", {
        url: asyncReq[a].url
      });
      Result[asyncReq[a].name] = Result[asyncReq[a].name].store;
    }
    store.commit("commitServicesList", Result.serviceList.data);

    let blogItem = Result.blogItem;
    let blogList = Result.blogList.data;
    return {
      blogItem,
      services: Result.serviceList.data,
      blogList: blogList.items,
      total: blogList.count
    };
  },
  computed: {},
  data() {
    return {
      perPage: 6,
      current: 1
    };
  },
  methods: {
    async setPage(val) {
      let param = { page: val, offset: this.perPage };
      let params = {
        url: `blog?page=${param.page}&perPage=${param.offset}`
      };
      let page = await this.$store.dispatch("Get", params);
      this.blogList = page.store.data.items;
    }
  }
};

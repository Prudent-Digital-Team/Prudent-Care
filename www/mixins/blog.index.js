// import { Notification, Notify } from '@/utils/helpers'
import { copy } from "@/utils/helpers";
import { mapState } from "vuex";
import _ from "lodash";

export default {
  async asyncData({ store, route }) {
    let Result = [];
    let asyncReq = [
      { name: "blogList", url: "blog?page=1&perPage=10" },
      { name: "serviceList", url: "services/all" }
    ];
    for (let a = 0; a < asyncReq.length; a++) {
      Result[asyncReq[a].name] = await store.dispatch("Get", {
        url: asyncReq[a].url
      });
      Result[asyncReq[a].name] = Result[asyncReq[a].name].store;
    }
    store.commit("commitServicesList", Result.serviceList.data);

    let blogList = Result.blogList.data.items;
    let featured_blog = blogList[0];
    // blogList = blogList.splice(1);
    return { blogList, featured_blog, services: Result.serviceList.data };
  },
  methods: {
    isEmpty(value) {
      let empty = _.isEmpty(value);
      return empty ? false : true;
    }
  }
};

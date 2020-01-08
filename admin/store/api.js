export const state = () => ({
  Test: 'test',
  Pages: [
    {
      name: 'Hompage',
      route: 'home',
      attributes: {
        header: {
          image: {},
          pageName: 'Homepage',
          content: {
            title: '',
            subtitle: '',
            image: {}
          }
        },
        content: {
          title: '',
          List: [
            { title: '', link: '', image: {} },
            { title: '', link: '', image: {} },
            { title: '', link: '', image: {} }
          ]
        }
      }
    },
    {
      name: 'About',
      route: 'about',
      attributes: {
        header: {
          image: {},
          pageName: 'About'
        },
        content: {
          content: '',
          title: ''
        }
      }
    },
    {
      name: 'Services',
      route: 'services',
      attributes: {
        header: {
          image: {},
          pageName: 'services'
        },
        content: {
          title: '',
          content: ''
        }
      }
    },
    {
      name: 'Process',
      route: 'process',
      attributes: {
        header: {
          image: {},
          title: 'process'
        },
        content: {
          title: 'Our Care Process',
          List: [{ title: '', image: {}, content: '' }]
        }
      }
    },
    {
      name: 'Faq',
      route: 'faq',
      attributes: {
        header: {
          title: 'Frequently Asked Questions (FAQ)'
        },
        content: {
          title: '',
          content: ''
        }
      }
    },
    {
      name: 'Contacts',
      route: 'contacts',
      attributes: {
        header: {
          image: {}
        },
        content: {
          List: [
            {
              address: '130 Old Street, London EC1V 9BD',
              email: 'info@pbgcare.co.uk',
              mobile: '01322686765',
              title: 'Head Office'
            }
          ]
        }
      }
    }
  ]
})

export const actions = {
  async Initialize({ commit, state }) {
    console.log(state.Pages)
    for (let a of state.Pages) {
      // let req = await this.$axios.$post(`api/pages`, a)
    }
  }
}

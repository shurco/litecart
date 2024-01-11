const { createApp, ref } = Vue

import FormButton from '/assets/js/form/button.js'

const App = {
  components: {
    FormButton
  },

  data() {
    return {
      currentSlide: ref(0),
      
      // global error
      error: ref(),

      // settings
      loaded: false,
      currency: sessionStorage.getItem('currency') || '',
      pages: JSON.parse(sessionStorage.getItem('pages')) || ref([]),
      socials: JSON.parse(sessionStorage.getItem('socials')) || ref([]),
      payments: ref([]),
      title: sessionStorage.getItem('title') || 'litecart',

      // cart
      cart: JSON.parse(localStorage.getItem('cart')) || ref([]),
      email: localStorage.getItem('email') || ref(''),
      provider: localStorage.getItem('provider') || ref(''),

      // products
      load: false,
      products: ref([]),

      // pages
      content: ref([]),

      socialUrl: {
        facebook: 'https://facebook.com/',
        instagram: 'https://instagram.com/',
        twitter: 'https://twitter.com/@',
        dribbble: 'https://dribbble.com/',
        github: 'https://github.com/',
        youtube: 'https://youtube.com/',
        other: '',
      }
    }
  },

  created() {
    if (
      !sessionStorage.getItem('currency') ||
      !sessionStorage.getItem('pages') ||
      !sessionStorage.getItem('socials') ||
      !sessionStorage.getItem('title') ||
      !sessionStorage.getItem('timestamp')
    ) {
      this.settings()
    }

    if (Date.now() > sessionStorage.getItem('timestamp')) {
      this.settings()
    }
  },

  beforeMount() {
    const currentPathname = window.location.pathname
    switch (true) {
      case currentPathname.startsWith('/cart'):
        if (currentPathname.startsWith('/cart/payment/cancel') || currentPathname.startsWith('/cart/payment/success')) {
          console.log('clear cart')
          localStorage.removeItem('cart')
          this.cart = ref([])
          break
        }
        this.listPayments()
        break
      case currentPathname.startsWith('/products'):
        this.getProduct(currentPathname.replace('/products/', ''))
        break
      case currentPathname.length > 1:
        this.getPage(currentPathname.substring(1))
        break
      default:
        this.listProducts()
        break
    }
  },

  mounted() {
    this.$nextTick(function () {
      this.loaded = true
    })

    // init meta tags
    document.title = this.title
    document.querySelector('meta[name="title"]').setAttribute('content', this.title)
    document.querySelector('meta[property="og:title"]').setAttribute('content', this.title)

    // init cart events
    window.addEventListener('addProduct', (event) => {
      this.cart.push(event.detail)
      localStorage.setItem('cart', JSON.stringify(this.cart))
    })
    window.addEventListener('delProduct', () => {
      this.cart = JSON.parse(localStorage.getItem('cart'))
    })
  },

  methods: {
    async settings() {
      const response = await fetch(`/api/settings`, {
        credentials: 'include',
        method: 'GET'
      })
      const resp = await response.json()
      if (resp.success) {
        var timestamp = new Date(Date.now())
        timestamp.setTime(timestamp.getTime() + 5 * 60 * 1000)
        sessionStorage.setItem('timestamp', timestamp.getTime())

        this.currency = resp.result.main.currency
        sessionStorage.setItem('currency', this.currency)

        this.title = resp.result.main.site_name
        sessionStorage.setItem('title', this.title)

        this.pages = resp.result.pages
        sessionStorage.setItem('pages', JSON.stringify(resp.result.pages))

        this.socials = resp.result.socials
        sessionStorage.setItem('socials', JSON.stringify(resp.result.socials))
      }
    },

    async listPayments() {
      const response = await fetch(`/api/cart/payment`, {
        credentials: 'include',
        method: 'GET'
      })
      const resp = await response.json()
      if (resp.success) {
        this.payments = resp.result
      }
    },

    // cart functions
    addCart(id) {
      let product = {}
      if (this.products.length > 0) {
        product = this.products.find((item) => item.id === id)
      } else {
        product = this.product
      }

      if (!this.inCart(product.id)) {
        product.inCart = true
        const image = product.images
          ? {
            name: product.images[0].name,
            ext: product.images[0].ext
          }
          : null

        window.dispatchEvent(
          new CustomEvent('addProduct', {
            detail: {
              id: product.id,
              name: product.name,
              slug: product.slug,
              amount: product.amount,
              image: image
            }
          })
        )
      }
    },

    removeCart(id) {
      const index = this.cart.findIndex((item) => item.id === id)
      if (index !== -1) {
        this.cart.splice(index, 1)
        localStorage.setItem('cart', JSON.stringify(this.cart))
        window.dispatchEvent(new CustomEvent('delProduct'))

        if (this.product) {
          this.product.inCart = false
        }

        if (this.products.length > 0) {
          console.log(this.products)
          this.products.find((item) => item.id === id).inCart = false
        }
      }
    },

    inCart(id) {
      if (localStorage.getItem('cart')) {
        const cart = JSON.parse(localStorage.getItem('cart'))
        let obj = cart.find((o) => o.id === id)
        if (obj) {
          return true
        }
        return false
      }
    },

    totalCartAmount() {
      let total = 0
      for (const item of this.cart) {
        total = total + item.amount
      }
      return this.costFormat(total)
    },

    async checkOut() {
      localStorage.setItem('email', this.email)
      localStorage.setItem('provider', this.provider)

      this.showOverlay()

      var cart = {
        email: this.email,
        provider: this.provider,
        products: this.cart.map((item) => ({ id: item.id, quantity: 1 }))
      }

      const response = await fetch(`/cart/payment`, {
        credentials: 'include',
        method: 'POST',
        body: JSON.stringify(cart),
        headers: {
          'Content-Type': 'application/json'
        }
      })
      const resp = await response.json()
      if (resp.success) {
        window.location.href = resp.result
      } 
      
      this.error = resp.message;
    },

    showPayments() {
      if (!this.payments['stripe'] && !this.payments['paypal'] && !this.payments['spectrocoin']) {
        localStorage.removeItem('provider')
        return false
      }
      return true
    },

    showSelectPayments() {
      if (this.payments['stripe'] && !this.payments['paypal'] && !this.payments['spectrocoin']) {
        localStorage.setItem('provider', 'stripe')
        return false
      }
      if (!this.payments['stripe'] && this.payments['paypal'] && !this.payments['spectrocoin']) {
        localStorage.setItem('provider', 'paypal')
        return false
      }

      if (!this.payments['stripe'] && !this.payments['paypal'] && this.payments['spectrocoin']) {
        localStorage.setItem('provider', 'spectrocoin')
        return false
      }
      return true
    },

    // product functions
    async listProducts() {
      const response = await fetch(`/api/products`, {
        credentials: 'include',
        method: 'GET'
      })
      const resp = await response.json()
      if (resp.success) {
        this.currency = sessionStorage.getItem('currency')
        this.products = resp.result.products
        this.load = true
      }
    },

    async getProduct(slug) {
      const response = await fetch(`/api/products/${slug}`, {
        credentials: 'include',
        method: 'GET'
      })
      this.resp = await response.json()
      if (this.resp.success) {
        this.currency = sessionStorage.getItem('currency')
        this.product = this.resp.result
        this.product.inCart = this.inCart(this.product.id)
        this.load = true

        if (this.product.seo.title) {
          document.title = this.product.seo.title
          document.querySelector('meta[name="title"]').setAttribute('content', this.product.seo.title)
        }
        if (this.product.seo.keywords) {
          document.querySelector('meta[name="keywords"]').setAttribute('content', this.product.seo.keywords)
        }
        if (this.product.seo.description) {
          document.querySelector('meta[name="description"]').setAttribute('content', this.product.seo.description)
          document
            .querySelector('meta[property="og:description"]')
            .setAttribute('content', this.product.seo.description)
        }
      }
    },

    // page function
    async getPage(slug) {
      const response = await fetch(`/api/pages/${slug}`, {
        credentials: 'include',
        method: 'GET'
      })
      this.resp = await response.json()
      if (this.resp.success) {
        this.content = this.resp.result

        if (this.content.seo.title) {
          document.title = this.content.seo.title
          document.querySelector('meta[name="title"]').setAttribute('content', this.content.seo.title)
          document.querySelector('meta[property="og:title"]').setAttribute('content', this.content.seo.title)
        }
        if (this.content.seo.keywords) {
          document.querySelector('meta[name="keywords"]').setAttribute('content', this.content.seo.keywords)
        }
        if (this.content.seo.description) {
          document.querySelector('meta[name="description"]').setAttribute('content', this.content.seo.description)
          document
            .querySelector('meta[property="og:description"]')
            .setAttribute('content', this.content.seo.description)
        }
      }
    },

    showOverlay() {
      this.error = ""
      document.getElementById('overlay').classList.remove('hidden')
      document.getElementById('overlay').classList.add('flex')
    },

    hideOverlay() {
      this.error = ""
      document.getElementById('overlay').classList.remove('flex')
      document.getElementById('overlay').classList.add('hidden')
    },


    nextSlide(length) {
      this.currentSlide = (this.currentSlide + 1) % length
    },

    prevSlide(length) {
      this.currentSlide = (this.currentSlide + length - 1) % length
    },

    // other utils
    costFormat(cost) {
      return Number(cost) ? (Number(cost) / 100).toFixed(2) : '0.00'
    }
  }
}

const app = createApp(App)
app.mount('#app')

const { createApp, ref } = Vue

import FormButton from '/assets/js/form/button.js'

createApp({
  components: {
    FormButton
  },

  data() {
    return {
      // settings
      currency: sessionStorage.getItem('currency'),
      pages: JSON.parse(sessionStorage.getItem('pages')) || ref([]),
      socials: JSON.parse(sessionStorage.getItem('socials')) || ref([]),
      payments: ref([]),

      // cart
      cart: JSON.parse(localStorage.getItem('cart')) || ref([]),
      email: localStorage.getItem('email') || ref(''),
      provider: localStorage.getItem('provider') || ref(''),

      // products
      load: false,
      products: ref([]),

      // pages
      content: ref([])
    }
  },

  mounted() {
    if (!sessionStorage.getItem('currency')) {
      this.settings()
    }

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

    // init meta tags
    const title = sessionStorage.getItem('title') || 'litecart'
    document.title = title
    document.querySelector('meta[name="title"]').setAttribute('content', title)
    document.querySelector('meta[property="og:title"]').setAttribute('content', title)

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
        sessionStorage.setItem('currency', resp.result.main.currency)
        sessionStorage.setItem('title', resp.result.main.site_name)
        sessionStorage.setItem('pages', JSON.stringify(resp.result.pages))
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
    },

    showPayments() {
      if (this.payments['stripe'] && !this.payments['spectrocoin']) {
        localStorage.setItem('provider', 'stripe')
        return false
      }
      if (!this.payments['stripe'] && this.payments['spectrocoin']) {
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

    // other utils
    costFormat(cost) {
      return Number(cost) ? (Number(cost) / 100).toFixed(2) : '0.00'
    }
  }
}).mount('#app')

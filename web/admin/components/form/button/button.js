import {importTemplate} from '/assets/js/utils.js';

export default {
  data() { },
  
  setup(props) { },

  props: {
    color: {
      type: String,
      required: true
    },
    name: {
      type: String,
      default: 'Name'
    },
    ico: String,
  },

  template: await importTemplate('./components/form/button/button.html')
}

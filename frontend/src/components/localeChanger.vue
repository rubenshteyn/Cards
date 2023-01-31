<template>
  <div class="localeChanger d-flex align-items-center header-lang">
    <button class="btn btn-dark btn-sm mr-2" v-for="language in languages" :key="language" :value="language" @click="changeLocale(language)" >{{ language }}</button>
  </div>
</template>

<script>
export default {
  name: 'localeChanger',
  data () {
    return {
      languages: ['en', 'ru'],
      key: ""
    }
  },
  methods: {
    changeLocale (language) {
      this.$cookies.set('locale', language)
      this.$i18n.locale = language
    }
  }
}
</script>
<style scoped>
  .localeChanger {
    text-align: center;
  }
  button {
    padding: 5px 10px !important;
  }
</style>
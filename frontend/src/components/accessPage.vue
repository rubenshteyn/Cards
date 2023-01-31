<template>
  <p v-if="error">{{ error }}</p>
</template>

<script>
import { defineComponent } from 'vue'
import { mapActions } from "vuex";

export default defineComponent({
  name: 'accessPage',
  data() {
    return {
      error: null
    }
  },
  methods: {
    ...mapActions({ access: "auth/access" }),
  },
  created() {
    let token = this.$route.params.token
    this.access(token)
        .then(() => {
          this.$router.push("/profile");
        })
        .catch(error => {
          console.log(error)
          this.error = error
        })
  }
})
</script>
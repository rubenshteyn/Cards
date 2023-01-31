<template>
  <div class="verify-email">
    <p v-if="!success || errorMessage" class="response-error">{{ $t('verifyEmail.' + errorMessage) }}</p>
    <p v-else>{{ $t('verifyEmail.success') }}</p>
  </div>
</template>

<script>
import apiV1 from "@/axios"
export default {
  name: "verify-email",
  data() {
    return {
      errorMessage: null,
      success: false
    }
  },
  beforeMount() {
    apiV1.get("/auth/verifyEmail/" + this.$route.params.token)
        .then(response => {
          this.errorMessage = null
          this.success = response.data.success
          if (!this.success) {
            this.errorMessage = response.data.error
          }
        })
  }
};
</script>

<style scoped>
.response-error {
  color: red;
}
</style>
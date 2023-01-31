<template>
  <div class="reset-password">
    <form
        v-if="!success"
        id="reset-password"
        @submit.prevent="checkForm"
    >
      <div class="form-group">
        <label for="newPassword">{{ $t('common.password') }}</label>
        <input
            type="password"
            class="form-control form-control-lg"
            id="newPassword"
            v-model="form.newPassword"
            name="newPassword"
            :placeholder="$t('common.password') + '...'"
            required
        />
      </div>

      <div class="form-group">
        <label for="newPasswordCheck">{{ $t('signUp.retypePassword') }}</label>
        <input
            type="password"
            class="form-control form-control-lg"
            id="newPasswordCheck"
            v-model="form.newPasswordCheck"
            name="newPasswordCheck"
            :placeholder="$t('signUp.retypePassword') + '...'"
            required
        />
        <p class="hidden-error-block password-error" v-if="passwordError" >{{ $t('signUp.errors.password') }}</p>
      </div>

      <p v-if="errorMessage" class="response-error">{{ $t('resetPassword.errors.' + errorMessage) }}</p>

      <button type="submit" class="btn btn-dark btn-lg btn-block">{{ $t('forgot.resetPassword') }}</button>

    </form>
    <p v-else>{{ $t('resetPassword.success') }}</p>
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import apiV1 from "@/axios";

export default defineComponent({
  name: "reset-password",
  data() {
    return {
      errorMessage: null,
      passwordError: false,
      success: false,
      form: {
        newPassword: null,
        newPasswordCheck: null
      },
      token: null
    }
  },
  watch: {
    'form.newPassword': 'checkPasswords',
    'form.newPasswordCheck': 'checkPasswords',
  },
  methods: {
    checkPasswords: function () {
      this.passwordError = this.form.newPassword !== this.form.newPasswordCheck;
    },
    checkForm: function (e) {
      this.errorMessage = null
      this.token = this.$route.params.token
      if (!this.passwordError) {
        apiV1.post("/auth/resetPassword/" + this.token, this.form)
            .then(response => {
              this.success = response.data.success
              if (!this.success) {
                this.errorMessage = response.data.error
              }
            })
      }
      if (this.errorMessage || this.passwordError) {
        e.preventDefault();
      }
    }
  },
})
</script>

<style scoped>
.password-error, .response-error {
  color: red;
}
</style>
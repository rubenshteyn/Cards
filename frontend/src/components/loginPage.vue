<template>
<!--  <div class="vue-template loginForm d-flex align-items-center flex-column-fluid flex-lg-row">-->
    <div class="d-flex flex-center w-lg-50 pt-15 pt-lg-0 px-10">
      <!--begin::Aside-->
      <div class="d-flex flex-column align-items-center">
        <!--begin::Logo-->
        <a class="mb-5 w-350px">
          <img alt="Logo" :src="logo" style="width: 100%;">
        </a>
        <!--end::Logo-->
        <!--begin::Title-->
        <h2 class="text-white fw-normal m-0">Financial solutions for you business</h2>
        <!--end::Title-->
      </div>
      <!--begin::Aside-->
    </div>
    <div class="d-flex flex-center w-lg-50 p-10">
      <div class="card rounded-3 w-550px w-md-550px w-lg-550px w-sm-420px position-relative" style="width: 350px">
        <div class="card-body p-10 p-lg-20">
          <!--begin::Form-->
          <form class="form w-100 mb-10" novalidate="novalidate" action="#" id="login" @submit.prevent="recaptcha">
            <!--begin::Heading-->
            <div class="text-center mb-11">
              <!--begin::Title-->
              <h1 class="text-dark fw-bolder mb-3">{{ $t('signIn.main') }}</h1>
              <!--end::Title-->
            </div>
            <!--begin::Heading-->
            <!--begin::Separator-->
            <div class="separator separator-content my-14">
              <span class="w-125px text-gray-500 fw-semibold fs-7">With email</span>
            </div>
            <!--end::Separator-->
            <!--begin::Input group=-->
            <div class="row">
                <div class="fv-row fv-row mb-8 fv-plugins-icon-container">
                <!--begin::Email-->
                <input
                    type="email"
                    class="form-control form-control-lg bg-transparent"
                    id="email"
                    v-model="form.email"
                    name="email"
                    :placeholder="$t('common.emailAddress') + '...'"
                    required
                />
                <!--end::Email-->
              </div>
              <!--end::Input group=-->
              <div class="fv-row mb-3">
                <input
                    type="password"
                    class="form-control form-control-lg bg-transparent"
                    id="password"
                    v-model="form.password"
                    name="password"
                    :placeholder="$t('common.password') + '...'"
                    required
                />
              </div>
            </div>
            <!--end::Input group=-->
            <!--begin::Wrapper-->
            <div class="d-flex flex-stack flex-wrap gap-3 fs-base fw-semibold mb-8">
              <div></div>
              <router-link class="app-bar-item" href="#" v-if="!loggedIn" @click.prevent to="/forgot-password">
                {{ $t('forgot.main') }}
              </router-link>
            </div>
            <!--end::Wrapper-->

<!--            <h3 v-if="errors.responseError !== null" class="hidden-error-block response-error">-->
<!--              {{ $t('signIn.errors.' + errors.responseError) }}</h3>-->

            <!--begin::Submit button-->
            <div class="d-grid mb-10">
              <button :disabled="this.disabled" type="submit" class="btn btn-primary">
                <span v-if="this.disabled === false">{{ $t('signIn.main') }}</span>
                <span v-else class="submit-spinner"></span>
              </button>
            </div>
            <!--end::Submit button-->
            <!--begin::Sign up-->
            <div class="text-gray-500 text-center fw-semibold fs-6">{{ $t('signUp.noReadyRegistered') }}
              <router-link class="app-bar-item link-primary" href="#" v-if="!loggedIn" @click.prevent
                           to="/registration">
                {{ $t('signUp.main') }}
              </router-link>
            </div>
            <!--end::Sign up-->
          </form>
          <div class="text-gray-500 text-center fw-semibold fs-6">
            This site is protected by reCAPTCHA and the Google
            <a href="https://policies.google.com/privacy">Privacy Policy</a>
            and <a href="https://policies.google.com/terms">Terms of Service</a> apply.
          </div>
        </div>
        <div class="position-absolute" style="text-align: center; top: 10px; left: 10px">
          <localChanger/>
        </div>
        <!--end::Form-->
      </div>
    </div>
<!--  </div>-->
</template>


<script>
import {defineComponent} from 'vue'
import {mapActions, mapGetters} from "vuex";
import logo from '@/assets/media/logo.png'
import localChanger from "@/components/localeChanger";

export default defineComponent({
  name: 'loginPage',
  data() {
    return {
      logo,
      url: process.env.VUE_APP_URL,
      apiUrl: process.env.VUE_APP_API_URL,
      valid: true,
      disabled: false,
      sitekey: process.env.VUE_APP_RECAPTCHA_SITEKEY,
      form: {
        email: null,
        password: null,
        token: null,
        recaptchaToken: null
      },
      errors: {
        responseError: null
      },
      response: null
    }
  },
  created() {
    if (this.loggedIn) {
      this.$router.push("/cards")
    }
  },
  methods: {
    ...mapActions({
      login: "auth/login"
    }),
    async recaptcha() {
      await this.$recaptchaLoaded()
      const token = await this.$recaptcha('login')
      this.form.recaptchaToken = token
      if(token) {
        this.loginAction()
      }
    },
    async loginAction() {
      this.disabled = true
      this.errors.responseError = null
      this.response = null
      console.log(this.form)
      this.login(this.form)
          .then((response) => {
            console.log(response)
            this.response = response
            this.disabled = false
            this.$router.push("/cards");
          })
          .catch(error => {
            console.log(error)
            let errorText = this.$t("signIn.errors." + error)
            this.$swal({
              title: errorText,
              type: "alert message",
              icon: "error",
              confirmButtonColor: "#f1416c",
            })
            this.disabled = false
            console.log(error)
            this.errors.responseError = error
          })
    },
  },
  computed: {
    ...mapGetters({
      loggedIn: "auth/isLoggedIn",
      user: "auth/currentUser"
    })
  },
  components: {
    localChanger
  }
})
</script>

<style>
</style>
<style scoped>
@keyframes spinner-border {
  100% {
    transform: rotate(360deg);
  }
}
.submit-spinner {
  display: inline-block;
  width: 1rem;
  height: 1rem;
  vertical-align: -0.125em;
  border: 0.2em solid currentColor;
  border-right-color: transparent;
  border-radius: 50%;
  -webkit-animation: .75s linear infinite spinner-border;
  animation: .75s linear infinite spinner-border;
}

.submit-spinner_hide {
  display: none;
}
.response-error {
  color: red;
}

@media (max-width: 768px) {
  .w-sm-420px {
    width: 420px !important;
  }
}
</style>
<template>
<!--  <div class="vue-template loginForm d-flex align-items-center flex-column-fluid flex-lg-row">-->
    <div class="d-flex flex-center w-lg-50 pt-15 pt-lg-0 px-10">
      <!--begin::Aside-->
      <div class="d-flex flex-column align-items-center">
        <!--begin::Logo-->
        <a class="mb-5 w-350px">
          <img alt="Logo" src="../assets/media/logo.png" style="width: 100%;">
        </a>
        <!--end::Logo-->
        <!--begin::Title-->
        <h2 class="text-white fw-normal m-0">Financial solutions for you business</h2>
        <!--end::Title-->
      </div>
      <!--begin::Aside-->
    </div>
    <div class="d-flex flex-center w-lg-50 p-10">
          <div class="card rounded-3 w-550px w-md-550px w-lg-550px w-sm-420px">
            <div class="card-body p-10 p-lg-20">
              <!--begin::Form-->
              <form class="form w-100 mb-10" novalidate="novalidate" action="#" id="login" @submit.prevent="recaptcha">
                <!--begin::Heading-->
                <div class="text-center mb-11">
                  <!--begin::Title-->
                  <h1 class="text-dark fw-bolder mb-3">Sign In</h1>
                  <!--end::Title-->
                </div>
                <!--begin::Heading-->

                <!--begin::Separator-->
                <div class="separator separator-content my-14">
                  <span class="w-125px text-gray-500 fw-semibold fs-7">With email</span>
                </div>
                <!--end::Separator-->
                <!--begin::Input group=-->
                <div class="fv-row mb-8">
                  <!--begin::Email-->
                  <input type="email" placeholder="Email" name="email" autocomplete="off" class="form-control form-control-lg bg-transparent"
                         v-model="form.email" required />
                  <!--end::Email-->
                </div>
                <!--end::Input group=-->
                <div class="fv-row mb-3">
                  <!--begin::Password-->
                  <input type="password" placeholder="Password" name="password" autocomplete="off"
                         class="form-control form-control-lg bg-transparent" v-model="form.password" required />
                  <!--end::Password-->
                </div>
                <!--end::Input group=-->
                <!--begin::Wrapper-->
                <div class="d-flex flex-stack flex-wrap gap-3 fs-base fw-semibold mb-8">
                  <div></div>
                  <!--begin::Link-->
                  <a :href="this.url + '/forgot-password'"
                     class="link-primary">Forgot Password ?</a>
                  <!--end::Link-->
                </div>
                <!--end::Wrapper-->
                <!--begin::Submit button-->
                <div class="d-grid mb-10">
                  <button :disabled="this.disabled" type="submit" class="btn btn-primary">
                    <span v-if="this.disabled === false">Sign In</span>
                    <span v-else class="submit-spinner"></span>
                  </button>
                </div>
                <!--end::Submit button-->
                <!--begin::Sign up-->
                <div class="text-gray-500 text-center fw-semibold fs-6">Not a Member yet?
                  <a :href="this.url + 'registration'" class="link-primary">Sign up</a>
                </div>
                <!--end::Sign up-->
              </form>
              <div class="text-gray-500 text-center fw-semibold fs-6">
                This site is protected by reCAPTCHA and the Google
                <a href="https://policies.google.com/privacy">Privacy Policy</a>
                and <a href="https://policies.google.com/terms">Terms of Service</a> apply.
              </div>
              <!--end::Form-->
            </div>
          </div>
        </div>
<!--  </div>-->
</template>


<script>
  import { defineComponent } from 'vue'
  import { mapActions, mapGetters} from "vuex";
  export default defineComponent({
    name: 'loginPage',
    data() {
      return {
        url: process.env.VUE_APP_URL,
        valid: true,
        errorTitle: null,
        disabled: false,
        form: {
          email: null,
          password: null,
          recaptchaToken: null
        },
        errors: {
          responseError: null
        }
      }
    },
    created() {
      if (this.loggedIn) {
        this.$router.push("/")
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
        this.errors.responseError = null
        this.disabled = true
        this.login(this.form)
          .then(() => {
            this.disabled = false
            this.$router.push("/");
          })
          .catch(error => {
            console.log(error)
            this.errors.responseError = error
            if(this.errors.responseError.toString() === "Error: badRequest") this.errorTitle = "Something went wrong. Try again later."
            if(this.errors.responseError.toString() === "Error: userNotFound") this.errorTitle = "User with this email address is not found."
            if(this.errors.responseError.toString() === "Error: registrationNotCompleted") this.errorTitle = "This user is not completely registered. Check your email for any further instructions."
            if(this.errors.responseError.toString() === "Error: incorrectPassword") this.errorTitle = "Incorrect password."
            if(this.errors.responseError.toString() === "Error: Captcha is invalid.") this.errorTitle = "Captcha is invalid."
            console.log(this.errorTitle)
            this.$swal({
              title: this.errorTitle,
              type: "alert message",
              icon: "error",
              confirmButtonColor: "#f1416c",
            })
            this.disabled = false
          })
      }
    },
    computed: {
      ...mapGetters({
        loggedIn: "auth/isLoggedIn",
        user: "auth/currentUser"
      })
    }
  })
</script>

<style scoped>
@media (max-width: 768px) {
  .w-sm-420px {
    width: 420px !important;
  }
}
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
</style>
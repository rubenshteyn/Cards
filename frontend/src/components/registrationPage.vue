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
    <div class="card rounded-3 w-550px w-md-550px w-lg-550px w-sm-420px position-relative">
      <div class="card-body p-10 p-lg-20">
        <form class="form w-100 mb-10" v-if="!signedUp" id="registration" @submit.prevent="recaptcha">
          <!--begin::Heading-->
          <div class="text-center mb-11">
            <!--begin::Title-->
            <h1 class="text-dark fw-bolder mb-3">{{ $t('signUp.main') }}</h1>
            <!--end::Title-->
          </div>
          <!--begin::Heading-->

          <div class="form-group fv-row mb-8">
            <input type="text" class="form-control form-control-lg bg-transparent" id="name" v-model="form.name"
                   name="name" :placeholder="$t('common.fullName') + '...'" required/>
          </div>

          <div class="form-group fv-row mb-8">
            <input type="email" class="form-control form-control-lg bg-transparent" id="email" v-model="form.email"
                   name="email" :placeholder="$t('common.emailAddress') + '...'" required/>
          </div>

          <div class="form-group fv-row mb-8">
            <input type="password" class="form-control form-control-lg bg-transparent" id="password"
                   v-model="form.password" name="password" :placeholder="$t('common.password') + '...'" required/>
          </div>

          <div class="form-group fv-row mb-8">
            <input type="password" class="form-control form-control-lg bg-transparent" id="passwordCheck"
                   v-model="form.passwordCheck" name="passwordCheck" :placeholder="$t('signUp.retypePassword') + '...'"
                   required/>
            <p class="hidden-error-block password-error" v-if="errors.password" style="margin-top: 1rem; padding: 0;">
              {{ $t('signUp.errors.password') }}</p>
          </div>

          <div class="form-group fv-row mb-8 reg">
            <v-select class="bg-transparent " label="name" :options="countries" v-model="country"
                      :placeholder="$t('signUp.country') + '...'" required>
            </v-select>
          </div>

          <div class="form-group fv-row mb-8 reg">
            <v-select class="bg-transparent" label="name" :options='messengers' v-model="messenger"
                      :placeholder="$t('signUp.messenger') + '...'" required>
            </v-select>
            <br/>
            <input type="text" class="form-control form-control-lg bg-transparent" id="messenger"
                   v-model="messengerNickname" name="messenger" :placeholder="$t('signUp.messengerNickname') + '...'"
                   required/>
          </div>

          <!--begin::Submit button-->
          <div class="d-grid mb-10">
            <button :disabled="this.disabled" type="submit" class="btn btn-primary">
              <span v-if="this.disabled === false">{{ $t('signUp.main') }}</span>
              <span v-else class="submit-spinner"></span>
            </button>
          </div>

          <!--end::Submit button-->
          <!--begin::Sign up-->
          <div class="forgot-password text-gray-500 text-center fw-semibold fs-6">
            {{ $t('signUp.alreadyRegistered') }}
            <router-link class="link-primary fw-semibold" to="/login">{{ $t('signIn.main') }}</router-link>
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
    </div>
  </div>
  <!--  </div>-->
</template>

<script>
import {defineComponent} from 'vue'
import {mapActions, mapGetters} from "vuex";
import countries from '../countries.json'
import 'vue-select/dist/vue-select.css';
import localChanger from "@/components/localeChanger";
import logo from "@/assets/media/logo.png";

export default defineComponent({
  name: 'registrationPage',
  data() {
    return {
      logo,
      valid: true,
      disabled: false,
      response: null,
      errors: {
        password: false,
        responseError: null
      },
      form: {
        name: null,
        email: null,
        password: null,
        passwordCheck: null,
        country: null,
        messenger: null,
        recaptchaToken: null
      },
      signedUp: false,
      countries: countries,
      country: null,
      messengers: [{
        name: "Telegram",
        code: "TG"
      },
        {
          name: "Skype",
          code: "SK"
        }
      ],
      messenger: null,
      messengerNickname: null
    }
  },
  watch: {
    'form.password': 'checkPasswords',
    'form.passwordCheck': 'checkPasswords',
  },
  created() {
    if (this.loggedIn) {
      this.$router.push("/profile")
    }
  },
  computed: {
    ...mapGetters({
      loggedIn: "auth/isLoggedIn",
      user: "auth/currentUser"
    })
  },
  methods: {
    ...mapActions({
      signup: "auth/signup"
    }),
    checkPasswords: function () {
      this.errors.password = this.form.password !== this.form.passwordCheck;
    },
    async recaptcha() {
      await this.$recaptchaLoaded()
      const token = await this.$recaptcha('login')
      this.form.recaptchaToken = token
      if(token) {
        this.checkForm()
      }
    },
    checkForm: function (e) {
      this.errors.responseError = null
      this.response = null
      this.disabled = true
      if (!this.country.code || !this.messenger.code) {
        this.errors.responseError = "formNotValid"
      }
      if (!this.errors.password && !this.errors.responseError) {
        this.form.country = this.country.code
        this.form.messenger = this.messenger.code + ' ' + this.messengerNickname
        this.signup(this.form)
            .then(() => {
              this.signedUp = true
              this.$swal({
                title: this.$t("signUp.successfulSignUp"),
                type: "alert message",
                icon: "info",
              })
              this.$router.push("/login");
            })
            .catch(error => {
              console.log(error)
              let errorText = this.$t("signUp.errors." + error)
              this.$swal({
                title: errorText,
                type: "alert message",
                icon: "error",
                confirmButtonColor: "#f1416c",
              })
              this.errors.responseError = error
            });
      }
      if (this.errors.password || this.errors.responseError) {
        e.preventDefault();
      }
    },
  },
  components: {
    localChanger
  }
})
</script>

<style>
.reg .vs--searchable .vs__dropdown-toggle {
  border: 1px solid var(--kt-input-border-color) !important;
  min-height: calc(1.5em + 1.65rem + 2px);
  padding: 0.825rem 1.5rem;
  font-size: 1.15rem;
  border-radius: 0.625rem;
  color: var(--kt-input-color);
  background-color: var(--kt-input-bg) !important;
  border: 1px solid var(--kt-input-border-color);
}

.vs--open .vs__dropdown-menu {
  border-color: var(--kt-input-solid-bg-focus) !important;
}

.vs__dropdown-toggle {

}

.vs__selected, .vs__actions, .vs__search, .vs__search:focus, .vs__selected-options, .vs__actions {
  border: unset !important;
  padding: 0 !important;
  margin: 0 !important;
}

.vs__search {
  color: var(--kt-text-gray-500) !important;
}

.vs__search, .vs__search:focus, .vs__selected {
  font-size: 1.15rem !important;
  font-weight: 500;
}

.vs__selected {
  color: var(--kt-input-solid-color) !important;
}


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

.password-error,
.response-error {
  color: red;
}

@media (max-width: 768px) {
  .w-sm-420px {
    width: 420px !important;
  }
}
</style>
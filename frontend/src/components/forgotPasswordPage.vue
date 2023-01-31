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
      <!--begin::Card-->
      <div class="card rounded-3 w-md-550px">
        <!--begin::Card body-->
        <div class="card-body p-10 p-lg-20">
          <!--begin::Form-->
          <form id="forgot-password" class="form w-100 fv-plugins-bootstrap5 fv-plugins-framework">
            <!--begin::Heading-->
            <div class="text-center mb-10">
              <!--begin::Title-->
              <h1 class="text-dark fw-bolder mb-3">{{ $t('forgot.main') }}</h1>
              <!--end::Title-->
              <!--begin::Link-->
              <div class="text-gray-500 fw-semibold fs-6">{{ $t('resetPassword.enterEmail') }}</div>
              <!--end::Link-->
            </div>
            <!--begin::Heading-->
            <!--begin::Input group=-->
            <div class="fv-row mb-8 fv-plugins-icon-container">
              <!--begin::Email-->
              <input type="email" v-model="form.email" :placeholder="$t('common.emailAddress')" name="email" autocomplete="off" class="form-control bg-transparent form-control form-control-lg">
              <!--end::Email-->
              <div class="fv-plugins-message-container invalid-feedback"></div></div>
            <!--begin::Actions-->
            <div class="d-flex flex-wrap justify-content-center pb-lg-0">
              <button @click="checkForm()" type="button" class="btn btn-primary me-4">
                <!--begin::Indicator label-->
                <span class="indicator-label">{{ $t('common.submit') }}</span>
                <!--end::Indicator label-->
                <!--begin::Indicator progress-->
                <span class="indicator-progress">Please wait...
										<span class="spinner-border spinner-border-sm align-middle ms-2"></span></span>
                <!--end::Indicator progress-->
              </button>
              <router-link to="/login" class="btn btn-light">{{ $t('common.cancel') }}</router-link>
            </div>
            <!--end::Actions-->
            <div></div>
          </form>
          <!--end::Form-->
        </div>
        <!--end::Card body-->
      </div>
      <!--end::Card-->
    </div>
</template>

<script>
  import {defineComponent} from 'vue'
  import apiV1 from "@/axios";
  import logo from "@/assets/media/logo.png";

  export default defineComponent({
    name: 'forgotPasswordPage',
    data() {
      return {
        logo,
        send: false,
        errorMessage: null,
        errorTitle: null,
        form: {
          email: null
        }
      }
    },
    methods: {
      checkForm: function (e) {
        this.errorMessage = null
        apiV1.post("/auth/forgotPassword", this.form)
          .then(response => {
            this.send = response.data.success
            if (!this.send) {
              this.errorMessage = response.data.error
              if(this.errorMessage.toString() === "Error: badRequest") this.errorTitle = "Something went wrong. Try again later."
              if(this.errorMessage.toString() === "Error: userNotFound") this.errorTitle = "User with this email address is not found."
              if(this.errorMessage.toString() === "Error: registrationNotCompleted") this.errorTitle = "This user is not completely registered. Check your email for any further instructions."
              if(this.errorMessage.toString() === "Error: incorrectPassword") this.errorTitle = "Incorrect password."
              this.$swal({
                title: this.errorTitle,
                type: "alert message",
                icon: "error",
                confirmButtonColor: "#f1416c",
              })
            }
            this.$swal({
              title: this.$t('forgot.success') ,
              confirmButtonColor: "#50cd89",
              type: "alert message",
              icon: "success",
            })
          })
        if (this.errorMessage) {
          e.preventDefault();
        }
      }
    }
  })
</script>

<style scoped>
  .response-error {
    color: red;
  }
</style>
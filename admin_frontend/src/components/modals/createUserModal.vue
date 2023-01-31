<template>
  <!-- из шаблона -->
  <div v-if="show">
    <div :class="{show: this.isShow}" class="modal fade d-block" id="kt_modal_view_users" tabindex="-1"
         aria-hidden="true" >
      <div class="modal-dialog modal-dialog-centered mw-550px" :style="{transform: this.transform}">
        <!--begin::Modal content-->
        <div class="modal-content">
          <!--begin::Modal header-->
          <div class="modal-header">
            <div class="text-center">
              <!--begin::Title-->
              <h1 class="text-dark fw-bolder">User creation</h1>
              <!--end::Title-->
            </div>
            <div @click="closeModal" class="btn btn-sm btn-icon btn-active-color-primary"
                 data-bs-dismiss="modal">
              <!--begin::Svg Icon | path: icons/duotune/arrows/arr061.svg-->
              <span class="svg-icon svg-icon-1">
                            <svg width="24" height="24" viewBox="0 0 24 24" fill="none"
                                 xmlns="http://www.w3.org/2000/svg">
                                <rect opacity="0.5" x="6" y="17.3137" width="16" height="2" rx="1"
                                      transform="rotate(-45 6 17.3137)" fill="currentColor"/>
                                <rect x="7.41422" y="6" width="16" height="2" rx="1" transform="rotate(45 7.41422 6)"
                                      fill="currentColor"/>
                            </svg>
                        </span>
              <!--end::Svg Icon-->
            </div>
          </div>

          <div class="modal-body scroll-y">
            <div>
              <!--begin::List-->
              <div class="modal-height scroll-y me-n7 pe-7">
                <div class="card-body p-0">
                  <div class="card rounded-3">
                    <div class="card-body">
                      <form class="form w-100"
                            id="addUserRegistration">
                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="name">Username</label>
                          <input type="text" class="form-control form-control-lg form-control-solid"
                                 id="name" v-model="form.name" name="name" placeholder="Enter username"
                                 required/>
                        </div>

                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="email">User email</label>
                          <input type="email" class="form-control form-control-lg form-control-solid"
                                 id="email" v-model="form.email" name="email"
                                 placeholder="Enter e-mail" required/>
                        </div>

                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="status">User status</label>
                          <v-select class="bg-transparent " label="name" :options="statuses"
                                    placeholder="Select user status" v-model="status" required>
                          </v-select>
                        </div>

                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="role">User role</label>
                          <v-select class="bg-transparent " label="name" :options="roles"
                                    placeholder="Select user role" v-model="role" required>
                          </v-select>
                        </div>

                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="password">User password</label>
                          <input type="password"
                                 class="form-control form-control-lg form-control-solid" id="password"
                                 v-model="form.password" name="password" placeholder="Enter password"/>
                        </div>

                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2">Repeat user password</label>
                          <input type="password"
                                 class="form-control form-control-lg form-control-solid"
                                 id="passwordCheck" v-model="form.passwordCheck" name="passwordCheck"
                                 placeholder="Enter user password" required/>
                          <p class="hidden-error-block password-error" v-if="errors.password"
                             style="color: red; text-align: center; margin-top: 1rem">
                            The entered passwords do not match.</p>
                        </div>

                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="Country">User country</label>
                          <v-select class="bg-transparent " label="name" :options="countries"
                                    v-model="country" placeholder="Enter user country"
                                    required>
                          </v-select>
                        </div>

                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="Messenger">User messenger</label>
                          <v-select class="bg-transparent" label="name" :options='messengers'
                                    placeholder="Select user messenger" v-model="messenger" required>
                          </v-select>
                          <br/>
                          <input type="text" class="form-control form-control-lg form-control-solid"
                                 id="messenger" v-model="messengerNickname" name="messenger"
                                 placeholder="Enter nickname" required/>
                        </div>
                        <div class="form-group text-center fv-row fs-6 fw-semibold" style="margin-bottom: 0">
                          <div v-if="isDisabled()" style="color:red;">Not all fields are filled</div>
                        </div>
                      </form>
                    </div>
                  </div>
                </div>
              </div>
              <!--end::List-->
            </div>
            <!--end::Users-->
          </div>
          <!--end::Modal body-->
          <div class="modal-footer flex-center">
            <button @click="closeModal" class="btn btn-light me-3">Discard</button>
            <button :disabled="isDisabled()" @click="checkForm()" type="submit" id="kt_modal_new_target_submit"
                    class="btn btn-primary">
              <span class="indicator-label">Create</span>
              <span class="indicator-progress">Please wait...
									<span class="spinner-border spinner-border-sm align-middle ms-2"></span></span>
            </button>
          </div>
        </div>
        <!--end::Modal content-->
      </div>
    </div>
  </div>
</template>

<script>
import {defineComponent} from 'vue'
import {mapActions, mapGetters} from "vuex";
import countries from '../../countries.json'
import 'vue-select/dist/vue-select.css';

const initialState = () => {
  return {
    show: false,
    isShow: false,
    transform: "translate(0, -50px)",
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
      status: null,
      role: null
    },
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
    statuses: [{
      name: "Active",
      code: 2
    },
      {
        name: "Moderation",
        code: 1
      },
      {
        name: "Unverified email",
        code: 0
      },
      {
        name: "Blocked",
        code: 3
      },
    ],
    status: null,
    roles: [
      {
        name: "Administrator",
        code: 2
      },
      {
        name: "User",
        code: 1
      },
      {
        name: "Manager",
        code: 3
      },
      {
        name: "Invited",
        code: 0
      }
    ],
    role: null,
    messengerNickname: null
  }
}

export default defineComponent({
  name: "createUserModal",
  data() {
    return initialState();
  },
  methods: {
    isDisabled() {
      return !(this.messenger && this.status && this.role
          && this.country && this.form.name && this.country
          && this.form.email && this.form.password && this.form.passwordCheck
          && this.form.password && this.messengerNickname);
    },
    openModal() {
      this.show = true
      setTimeout(()=> {
        this.isShow = this.show
        this.transform = "none"
      },200)
    },
    closeModal: function () {
      this.isShow = false
      this.transform = "translate(0, -50px)"
      setTimeout(()=> {
        this.show = false
      },200)
    },
    ...mapActions({
      addUser: "admin/addUser"
    }),
    checkPasswords: function () {
      this.errors.password = this.form.password !== this.form.passwordCheck;
    },
    checkForm: function (e) {
      this.errors.responseError = null
      if (!this.country.code || !this.messenger.code || !this.role.code || !this.status.code) {
        this.errors.responseError = "formNotValid"
      }
      if (!this.errors.password && !this.errors.responseError) {
        this.form.country = this.country.code
        this.form.messenger = this.messenger.code + ' ' + this.messengerNickname
        this.form.status = this.status.code
        this.form.role = this.role.code
        this.addUser(this.form)
            .then(() => {
              Object.assign(this.$data, initialState());
              this.closeModal()
              this.$swal({
                title: "You created a user",
                confirmButtonColor: "#50cd89",
                type: "alert message",
                icon: "success",
              })
            })
            .catch(error => {
              console.log(error)
              this.errors.responseError = error
              this.$swal({
                title: this.errors.responseError,
                type: "alert message",
                confirmButtonColor: "#f1416c",
                icon: "error",
              })
            });
      }
      if (this.errors.password || this.errors.responseError) {
        e.preventDefault();
      }
    },
  },
  watch: {
    'form.password': 'checkPasswords',
    'form.passwordCheck': 'checkPasswords',
  },
  computed: {
    ...mapGetters({
      loggedIn: "auth/isLoggedIn",
      user: "auth/currentUser"
    })
  }
})
</script>

<style>
.table.gs-0 th:first-child,
.table.gs-0 td:first-child {
  padding-left: 1rem !important;
}

.fade {
  transition: opacity 1s linear;
}

.modal {
  background: rgba(0, 0, 0, .5);
}
</style>
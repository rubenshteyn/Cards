<template>
  <!-- из шаблона -->
  <div v-if="show">
    <div :class="{show: this.isShow}" class="modal fade d-block" id="kt_modal_view_users" tabindex="-1"
         aria-hidden="true">
      <div class="modal-dialog modal-dialog-centered mw-550px" :style="{transform: this.transform}">
        <!--begin::Modal content-->
        <div class="modal-content">
          <!--begin::Modal header-->
          <div class="modal-header">
            <h1 v-if="styles.data === 'block'" class="text-dark fw-bolder">Change user data</h1>
            <h1 v-else class="text-dark fw-bolder">Change personal prices</h1>
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
                    <div class="card-body" style="padding: 2rem 1rem;">
                      <div class="edit-card-actions form-group fv-row mb-8 d-flex">
                        <div class="edit-card-select-action d-flex align-items-center fs-6 fw-semibold"
                             style="white-space: nowrap">
                          Select an action:
                        </div>
                        <div class="d-flex align-items-center ml-5 edit-card-group-actions">
                          <div class="d-flex align-items-center mr-5">
                            <input @click="tabClicked('data')" class="form-check-input" type="radio" name="exampleRadios"
                                   id="exampleRadios1" value="option1" style="width: 1em; height: 1em;" checked>
                            <label class="form-check-label edit-card-label" for="exampleRadios1">
                              Personal data
                            </label>
                          </div>
                          <div class="d-flex align-items-center mr-5">
                            <input @click="tabClicked('prices')" class="form-check-input" type="radio" name="exampleRadios"
                                   id="exampleRadios2" value="option2" style="width: 1em; height: 1em;">
                            <label class="form-check-label" for="exampleRadios2">
                              Personal Prices
                            </label>
                          </div>
                        </div>
                      </div>
                      <div :style="{ display: styles.prices }" class="form w-100">
                        <div class="d-flex">
                          <div class="form-group col-md-6 fv-row flex-column d-flex" style="background-color: var(--kt-info-light); margin-right: 5px;">
                            <span class="fs-3 fw-semibold pb-5 mb-5 mt-5" for="name" style="color: var(--kt-info);border-bottom: 1px solid var(--kt-modal-header-border-color);">Universal</span>
                            <div class="d-flex flex-column mb-8">
                              <label class="required fs-6 fw-semibold mb-2" for="name">Card price</label>
                            <div class="position-relative">
                              <span class="fs-6 fw-semibold commission">$</span>
                              <input step="any" type="number" class="form-control form-control-lg form-control-solid"
                                     id="name" v-model="formPrices.universalCost" name="name" placeholder="Enter commission"
                                     required/>
                            </div>
                            </div>
                            <div class="form-group fv-row mb-8">
                              <label class="required fs-6 fw-semibold mb-2" for="name">Monthly maintenance</label>
                              <div class="position-relative monthlyMaintenance">
                                <span class="fs-6 fw-semibold commission">$</span>
                                <input step="any" type="number" class="form-control form-control-lg form-control-solid"
                                       id="name" v-model="formPrices.universalMonthlyMaintenance" name="name" placeholder="Enter sum"
                                       required/>
                              </div>
                            </div>
                            <div class="form-group fv-row">
                              <label class="required fs-6 fw-semibold mb-2" for="name">Replenishment percentage</label>
                              <div class="position-relative popUpCommission">
                                <span class="fs-6 fw-semibold commission">%</span>
                                <input step="any" type="number" class="form-control form-control-lg form-control-solid"
                                       id="name" v-model="formPrices.universalCommission" name="name" placeholder="Enter percent"
                                       required/>
                              </div>
                            </div>
                          </div>
                          <div class="form-group col-md-6 fv-row flex-column d-flex" style="background-color: var(--kt-danger-light); margin-left: 5px;">
                            <span class="fs-3 fw-semibold mt-5 mb-5 pb-5" for="name" style="color: var(--kt-danger);border-bottom: 1px solid var(--kt-modal-header-border-color);">Advertising</span>
                            <div class="d-flex flex-column mb-8">
                              <label class="required fs-6 fw-semibold mb-2" for="name">Card price</label>
                              <div class="position-relative">
                                <span class="fs-6 fw-semibold commission">$</span>
                                <input step="any" type="number" class="form-control form-control-lg form-control-solid"
                                       id="name" v-model="formPrices.advertisingCost" name="name" placeholder="Enter commission"
                                       required/>
                              </div>
                            </div>
                            <div class="form-group fv-row mb-8">
                              <label class="required fs-6 fw-semibold mb-2" for="name">Monthly maintenance</label>
                              <div class="position-relative monthlyMaintenance">
                                <span class="fs-6 fw-semibold commission">$</span>
                                <input step="any" type="number" class="form-control form-control-lg form-control-solid"
                                       id="name" v-model="formPrices.advertisingMonthlyMaintenance" name="name" placeholder="Enter sum"
                                       required/>
                              </div>
                            </div>
                            <div class="form-group fv-row">
                              <label class="required fs-6 fw-semibold mb-2" for="name">Replenishment percentage</label>
                              <div class="position-relative popUpCommission">
                                <span class="fs-6 fw-semibold commission">%</span>
                                <input step="any" type="number" class="form-control form-control-lg form-control-solid"
                                       id="name" v-model="formPrices.advertisingCommission" name="name" placeholder="Enter percent"
                                       required/>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                      <form class="form w-100" id="registration" :style="{ display: styles.data }">

                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="name">Username</label>
                          <input type="text" class="form-control form-control-lg form-control-solid"
                                 id="name" v-model="formData.name" name="name" placeholder="Enter username"
                                 required/>
                        </div>

                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="email">User email</label>
                          <input type="email" class="form-control form-control-lg form-control-solid"
                                 id="email" v-model="formData.email" name="email"
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
                                 v-model="formData.password" name="password" placeholder="Enter password"/>
                        </div>

                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="Country">User country</label>
                          <v-select class="bg-transparent " label="name" :options="countries"
                                    v-model="country" placeholder="Enter user country"
                                    required>
                          </v-select>
                        </div>

                        <div class="form-group fv-row" style="Datamargin-bottom: 0">
                          <label class="required fs-6 fw-semibold mb-2" for="Messenger">User messenger</label>
                          <v-select class="bg-transparent" label="name" :options='messengers'
                                    placeholder="Select user messenger" v-model="messenger" required>
                          </v-select>
                          <br/>
                          <input type="text" class="form-control form-control-lg form-control-solid"
                                 id="messenger" v-model="messengerNickname" name="messenger"
                                 placeholder="Enter nickname" required/>
                        </div>
                      </form >
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
            <button @click="checkForm()" type="submit" id="kt_modal_new_target_submit"
                    class="btn btn-primary">
              <span class="indicator-label">Save changes</span>
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
import {
  defineComponent
} from 'vue'
import {
  mapActions
} from "vuex";
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
    errorTitle: null,
    activeTab: 'data',
    styles: {
      prices: "none",
      data: "block"
    },
    formData: {
      id: 0,
      name: null,
      email: null,
      password: null,
      country: null,
      messenger: null,
      status: null,
      role: null,
    },
    formPrices: {
      id: 0,
      universalCost: 15,
      advertisingCost: 10,
      universalMonthlyMaintenance: 15,
      advertisingMonthlyMaintenance: 10,
      advertisingCommission: 6,
      universalCommission: 8,
    },
    countries: countries,
    country: {
      name: null,
      code: null
    },
    messengers: [{
      name: "Telegram",
      code: "TG"
    },
      {
        name: "Skype",
        code: "SK"
      }
    ],
    messenger: {
      name: null,
      code: null
    },
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
    status: {
      name: null,
      code: null
    },
    roles: [{
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
    role: {
      name: null,
      code: null
    },
    messengerNickname: null
  }
}

export default defineComponent({
  name: "editUserModal",
  data() {
    return initialState();
  },
  methods: {
    tabClicked(element) {
      for (const style in this.styles) {
        this.styles[style] = "none"
        if (style === element) {
          this.activeTab = element
          this.styles[style] = "block"
        }
      }
    },
    openModal(user) {
      Object.assign(this.$data, initialState());
      console.log(user)
      this.show = true
      setTimeout(() => {
        this.isShow = this.show
        this.transform = "none"
      }, 200)
      this.formData.id = user.ID
      this.formPrices.id = user.ID
      this.formData.name = user.name
      this.formData.email = user.email
      this.status.name = user.status
      if (user.universalCost) {
        this.formPrices.universalCost = user.universalCost
      }
      if (user.advertisingCost) {
        this.formPrices.advertisingCost = user.advertisingCost
      }
      if (user.universalCommission) {
        this.formPrices.universalCommission = user.universalCommission
      }
      if(user.universalMonthlyMaintenance) {
        this.formPrices.universalMonthlyMaintenance = user.universalMonthlyMaintenance
      }
      if (user.advertisingCommission) {
        this.formPrices.advertisingCommission = user.advertisingCommission
      }
      if(user.advertisingMonthlyMaintenance) {
        this.formPrices.advertisingMonthlyMaintenance = user.advertisingMonthlyMaintenance
      }

      this.role.name = user.role
      let userMessengerCode = user.messenger.split(' ')[0]
      this.messenger = this.messengers.find(item => item.code === userMessengerCode);
      this.messengerNickname = user.messenger.split(' ')[1]

      let userCountryCode = user.country.split(' ')[0]
      this.country = this.countries.find(item => item.code === userCountryCode);

      let userRoleCode = user.role
      this.role = this.roles.find(item => item.code === userRoleCode);

      let userStatusCode = user.status
      this.status = this.statuses.find(item => item.code === userStatusCode);
    },
    closeModal: function () {
      this.formData.id = null
      this.formPrices.id = null
      this.formData.name = null
      this.formData.email = null
      this.status.name = null
      this.formPrices.universalCost = null
      this.formPrices.universalCommission = null
      this.formPrices.universalMonthlyMaintenance = null
      this.formPrices.advertisingCost = null
      this.formPrices.advertisingCommission = null
      this.formPrices.advertisingMonthlyMaintenance = null
      this.role.name = null
      this.messenger = null
      this.messengerNickname = null
      this.country = null
      this.role = null
      this.status = null
      this.isShow = false
      this.transform = "translate(0, -50px)"
      setTimeout(() => {
        this.show = false
      }, 200)
    },
    ...mapActions({
      updateUser: "admin/updateUser",
      updateUsersCommission: "admin/updateUsersCommission"
    }),
    checkForm: function () {
      if(this.activeTab === "data") {
        this.formData.country = this.country.code
        this.formData.messenger = this.messenger.code + ' ' + this.messengerNickname
        this.formData.status = this.status.code
        this.formData.role = this.role.code
        console.log(this.formData)
        this.updateUser(this.formData)
            .then(() => {
              Object.assign(this.$data, initialState());
              this.closeModal()
              this.$swal({
                title: "You have successfully changed user data",
                confirmButtonColor: "#50cd89",
                type: "alert message",
                icon: "success",
              })
            })
            .catch(error => {
              this.errors.responseError = error
              if (this.errors.responseError.toString() === "Error: badRequest") this.errorTitle = "Something went wrong. Try again later."
              if (this.errors.responseError.toString() === "Error: cardNotFound") this.errorTitle = "Сard not found in the system"
              if (this.errors.responseError.toString() === "Error: change is Not Approved") this.errorTitle = "Change is not approved"
              if (this.errors.responseError.toString() === "Error: userNotFound") this.errorTitle = "User not found in the system"
              if(this.errors.responseError.toString() === "Error: notEnoughBalance") this.errorTitle = "There are not enough funds on the balance to replenish"
              if(this.errors.responseError.toString() === "Error: paying not from USDT balance is disabled for now") this.errorTitle = "Paying not from USDT balance is disabled for now"
              console.log(this.errorTitle)
              this.$swal({
                title: this.errorTitle,
                type: "alert message",
                confirmButtonColor: "#f1416c",
                icon: "error",
              })
            })
      }
      if(this.activeTab === "prices") {
        console.log(this.formPrices)
        this.updateUsersCommission(this.formPrices)
            .then(() => {
              Object.assign(this.$data, initialState());
              this.closeModal()
              this.$swal({
                title: "You have successfully changed personal price",
                confirmButtonColor: "#50cd89",
                type: "alert message",
                icon: "success",
              })
            })
            .catch(error => {
              console.log(error)
              this.errors.responseError = error
              if (this.errors.responseError.toString() === "Error: badRequest") this.errorTitle = "Something went wrong. Try again later."
              if (this.errors.responseError.toString() === "Error: cardNotFound") this.errorTitle = "Сard not found in the system"
              if (this.errors.responseError.toString() === "Error: change is Not Approved") this.errorTitle = "Change is not approved"
              if (this.errors.responseError.toString() === "Error: userNotFound") this.errorTitle = "User not found in the system"
              if(this.errors.responseError.toString() === "Error: notEnoughBalance") this.errorTitle = "There are not enough funds on the balance to replenish"
              if(this.errors.responseError.toString() === "Error: paying not from USDT balance is disabled for now") this.errorTitle = "Paying not from USDT balance is disabled for now"
              console.log(this.errorTitle)
              this.$swal({
                title: this.errorTitle,
                type: "alert message",
                confirmButtonColor: "#f1416c",
                icon: "error",
              })
            });
      }
    },
  }
})
</script>

<style scoped>
.commission {
  color: var(--kt-input-solid-color);
  top: 12.8px;
  left: 2.8rem;
  position: absolute;
}

.popUpCommission .commission {
  left: 3.4rem;
}

.monthlyMaintenance .commission {
  top: 12.5px;
}
</style>

<style>
.table.gs-0 th:first-child,
.table.gs-0 td:first-child {
  padding-left: 1rem !important;
}

.modal {
  background: rgba(0, 0, 0, .5);
}
</style>
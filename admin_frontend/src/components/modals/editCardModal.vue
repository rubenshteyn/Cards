<template>
  <!-- из шаблона -->
  <div v-if="show">
    <div :class="{show: this.isShow}" class="modal fade d-block" @click.self="closeModal" id="kt_modal_view_users"
         tabindex="-1"
         aria-hidden="true">
      <div class="modal-dialog modal-dialog-centered mw-850px" :style="{transform: this.transform}">
        <!--begin::Modal content-->
        <div class="modal-content">
          <!--begin::Modal header-->
          <div class="modal-header">
            <h1 class="text-dark fw-bolder">Card change</h1>

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
          <!--begin::Modal header-->
          <!--begin::Modal body-->
          <div class="modal-body scroll-y">
            <div>
              <!--begin::List-->
              <div class="mh-375px scroll-y me-n7 pe-7">
                <div class="card-body p-0">
                  <div class="edit-card-actions form-group fv-row mb-8 d-flex">
                    <div class="edit-card-select-action d-flex align-items-center fs-6 fw-semibold"
                         style="white-space: nowrap">
                      Select an action:
                    </div>
                    <div class="d-flex align-items-center ml-5 edit-card-group-actions">
                      <div class="d-flex align-items-center mr-5">
                        <input @click="tabClicked('limit')" class="form-check-input" type="radio" name="exampleRadios"
                               id="exampleRadios1" value="option1" style="width: 1em; height: 1em;" checked>
                        <label class="form-check-label edit-card-label" for="exampleRadios1">
                          Limit
                        </label>
                      </div>
                      <div class="d-flex align-items-center mr-5">
                        <input @click="tabClicked('status')" class="form-check-input" type="radio" name="exampleRadios"
                               id="exampleRadios2" value="option2" style="width: 1em; height: 1em;">
                        <label class="form-check-label" for="exampleRadios2">
                          Status
                        </label>
                      </div>
                      <div class="d-flex align-items-center mr-5" style="white-space: nowrap">
                        <input @click="tabClicked('userData'), getAdress()" class="form-check-input" type="radio"
                               name="exampleRadios" id="exampleRadios3" value="option3"
                               style="width: 1em; height: 1em;">
                        <label class="form-check-label edit-card-label" for="exampleRadios3">
                          Personal data
                        </label>
                      </div>
                    </div>
                  </div>
                  <div class="edit-card-limit-status form-group fv-row mb-8 w-50 ml-5"
                       :style="{ display: styles.limit }">
                    <label class="edit-card-label mr-4 min-w-115px fs-6 fw-semibold mb-2" for="limit">Limit:</label>
                    <input type="text" class="form-control form-control-lg form-control-solid"
                           id="limit" v-model="formLimit.limit" name="name" placeholder="Limit"
                           required/>
                  </div>
                  <div class="edit-card-limit-status form-group fv-row mb-8 w-50 ml-5"
                       :style="{ display: styles.status }">
                    <div class="d-flex">
                      <div class="mr-5 fs-6 fw-semibold">
                        Status:
                      </div>
                      <div class="mr-5 d-flex align-items-center">
                        <input class="form-check-input" type="radio" name="exRadios" id="exampleRadios4"
                               style="width: 1em; height: 1em;" :value="false" v-model="formStatus.status"
                               :checked="this.check1">
                        <label class="form-check-label fs-6 fw-semibold" for="exampleRadios4">
                          Blocked
                        </label>
                      </div>
                      <div class="mr-5 d-flex align-items-center">
                        <input class="form-check-input" type="radio" name="exRadios" id="exampleRadios5"
                               style="width: 1em; height: 1em;" v-model="formStatus.status" :value="true"
                               :checked="this.check2">
                        <label class="form-check-label fs-6 fw-semibold" for="exampleRadios5">
                          Active
                        </label>
                      </div>
                    </div>
                  </div>
                  <div :style="{ display: styles.userData }">
                    <div class="edit-card-group form-group fv-row mb-8 d-flex align-items-center">
                      <div class="edit-card-name-block d-flex col-md-6 flex-column">
                        <label class="edit-card-label mr-4 min-w-115px fs-6 fw-semibold mb-2" for="passwordCheck"
                               style="white-space: nowrap">Username</label>
                        <input v-model="formInfo.FirstName" type="name"
                               class="edit-card-name-input form-control form-control-lg form-control-solid mr-3"
                               placeholder="Enter username"/>
                      </div>
                      <div class="edit-card-name-block d-flex col-md-6 flex-column">
                        <label class="edit-card-label mr-4 min-w-115px fs-6 fw-semibold mb-2" for="passwordCheck"
                               style="white-space: nowrap">User last name</label>
                        <input v-model="formInfo.LastName" type="lastname"
                               class="form-control form-control-lg form-control-solid"
                               placeholder="Enter user last name"/>
                      </div>
                    </div>
                    <div class="edit-card-group form-group fv-row mb-8 col-md-12 d-flex flex-column">
                      <label class="edit-card-label mr-4 min-w-115px fs-6 fw-semibold mb-2" for="passwordCheck"
                             style="white-space: nowrap">Address</label>
                      <input v-model="formInfo.Address1" type="text"
                             class="form-control form-control-lg form-control-solid" placeholder="Enter address "/>
                    </div>
                    <div class="edit-card-group form-group fv-row mb-8 col-md-12 d-flex flex-column">
                      <label class="edit-card-label mr-4 min-w-115px fs-6 fw-semibold mb-2" for="passwordCheck"
                             style="white-space: nowrap">City</label>
                      <input v-model="formInfo.City" type="text" class="form-control form-control-lg form-control-solid"
                             placeholder="Enter city"/>
                    </div>
                    <div class="edit-card-group form-group fv-row mb-8 col-md-12 d-flex flex-column">
                      <label class="edit-card-label mr-4 min-w-115px fs-6 fw-semibold mb-2" for="passwordCheck"
                             style="white-space: nowrap">State</label>
                      <input v-model="formInfo.State" type="text"
                             class="form-control form-control-lg form-control-solid" placeholder="Enter state"/>
                    </div>
                    <div class="edit-card-group form-group fv-row mb-8 col-md-12 d-flex flex-column">
                      <label class="edit-card-label mr-4 min-w-115px fs-6 fw-semibold mb-2" for="passwordCheck"
                             style="white-space: nowrap">Country</label>
                      <input v-model="formInfo.Country" type="text"
                             class="form-control form-control-lg form-control-solid" placeholder="Enter country"/>
                    </div>
                    <div class="edit-card-group form-group fv-row d-flex col-md-12 flex-column">
                      <label class="edit-card-label mr-4 min-w-115px fs-6 fw-semibold mb-2" for="passwordCheck"
                             style="white-space: nowrap">ZIP</label>
                      <input v-model="formInfo.Zip" type="number"
                             class="form-control form-control-lg form-control-solid" required/>
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
              <span class="indicator-label">Save</span>
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
import countries from '../../countries.json'
import {mapActions} from "vuex";

export default {
  name: "editCardModal",
  data() {
    return {
      show: false,
      isShow: false,
      transform: "translate(0, -50px)",
      activeTab: 'limit',
      checked: false,
      check1: false,
      check2: false,
      fullname: null,
      errors: {
        responseError: null
      },
      errorTitle: null,
      styles: {
        limit: "block",
        status: "none",
        userData: "none"
      },
      countries: countries,
      formInfo: {
        cardID: null,
        FirstName: null,
        LastName: null,
        Address1: null,
        City: null,
        State: null,
        Country: null,
        Zip: null
      },
      formLimit: {
        cardID: null,
        limit: 0,
      },
      formStatus: {
        cardID: null,
        status: null
      }
    }
  },
  methods: {
    ...mapActions({
      changeCardStatus: "admin/changeCardStatus",
      changeCardInfo: "admin/changeCardInfo",
      changeCardLimit: "admin/changeCardLimit",
    }),
    openModal(card) {
      console.log(card)
      this.formStatus.cardID = card.ID
      this.formLimit.cardID = card.ID
      this.formInfo.cardID = card.ID
      this.formLimit.limit = card.limit
      this.fullname = card.name
      this.formInfo.Address1 = card.address
      if (card.status === 1) {
        this.check1 = false
        this.check2 = true
        this.formStatus.status = true
      }
      if (card.status === 0) {
        this.check1 = true
        this.check2 = false
        this.formStatus.status = false
      }
      console.log(this.formStatus.status)
      this.formInfo.FirstName = this.splitFullName(this.fullname)[0]
      this.formInfo.LastName = this.splitFullName(this.fullname)[1]
      this.formInfo.Zip = this.getZip(this.formInfo.Address1)[0]
      this.show = true
      setTimeout(() => {
        this.isShow = this.show
        this.transform = "none"
      }, 200)
    },
    splitFullName(fullname) {
      this.formInfo.FirstName = fullname.split(" ")[0]
      this.formInfo.LastName = fullname.split(" ")[1]
      return [this.formInfo.FirstName, this.formInfo.LastName]
    },

    closeModal() {
      this.formStatus.status = null
      this.check1 = false
      this.check2 = false
      this.activeTab = 'limit'
      this.isShow = false
      this.transform = "translate(0, -50px)"
      setTimeout(() => {
        this.show = false
      }, 200)
      this.styles.limit = "block"
      this.styles.status = "none"
      this.styles.userData = "none"
    },
    getAdress() {
      let country = this.countries.find(item => item.code === this.formInfo.Address1.split(",")[3].toUpperCase().replace(/\s+/g, ''));
      this.formInfo.Country = country.name
      this.formInfo.City = this.formInfo.Address1.split(",")[1]
    },
    getZip(address) {
      this.formInfo.State = address.split(",")[2].split(" ")[1]
      this.formInfo.Zip = address.split(",")[2].split(" ")[2]
      return [this.formInfo.Zip, this.formInfo.State]
    },
    tabClicked(element) {
      for (const style in this.styles) {
        this.styles[style] = "none"
        if (style === element) {
          this.activeTab = element
          this.styles[style] = "block"
        }
      }
    },
    checkForm: function () {
      this.errors.responseError = null
      if (this.activeTab === 'userData') {
        console.log(this.formInfo)
        this.changeCardInfo(this.formInfo)
            .then((response) => {
              console.log(response)
              this.closeModal()
              this.$swal({
                title: "You have successfully changed your personal data",
                confirmButtonColor: "#50cd89",
                type: "alert message",
                icon: "success",
              })
            })
            .catch(error => {
              this.errors.responseError = error
              console.log(this.errors.responseError.toString())
              if (this.errors.responseError.toString() === "Error: badRequest") this.errorTitle = "Something went wrong. Try again later."
              if (this.errors.responseError.toString() === "Error: cardNotFound") this.errorTitle = "Сard not found in the system"
              if (this.errors.responseError.toString() === "Error: change is Not Approved") this.errorTitle = "Change is not approved"
              if (this.errors.responseError.toString() === "Error: userNotFound") this.errorTitle = "User not found in the system"
              console.log(this.errorTitle)
              this.$swal({
                title: this.errorTitle,
                type: "alert message",
                confirmButtonColor: "#f1416c",
                icon: "error",
              })
            })
      }
      if (this.activeTab === 'status') {
        console.log(this.formStatus)
        this.changeCardStatus(this.formStatus)
            .then((response) => {
              console.log(response)
              this.closeModal()
              this.$swal({
                title: "You have successfully changed the status",
                confirmButtonColor: "#50cd89",
                type: "alert message",
                icon: "success",
              })
            })
            .catch(error => {
              this.errors.responseError = error
              console.log(this.errors.responseError.toString())
              if (this.errors.responseError.toString() === "Error: badRequest") this.errorTitle = "Something went wrong. Try again later."
              if (this.errors.responseError.toString() === "Error: cardNotFound") this.errorTitle = "Сard not found in the system"
              if (this.errors.responseError.toString() === "Error: change is Not Approved") this.errorTitle = "Change is not approved"
              if (this.errors.responseError.toString() === "Error: userNotFound") this.errorTitle = "User not found in the system"
              console.log(this.errorTitle)
              this.$swal({
                title: this.errorTitle,
                type: "alert message",
                confirmButtonColor: "#f1416c",
                icon: "error",
              })
            })
      }
      if (this.activeTab === 'limit') {
        this.formLimit.limit = Number(this.formLimit.limit)
        console.log(this.formLimit)
        this.changeCardLimit(this.formLimit)
            .then((response) => {
              console.log(response)
              this.closeModal()
              this.$swal({
                title: "You have successfully changed the limit",
                confirmButtonColor: "#50cd89",
                type: "alert message",
                icon: "success",
              })
            })
            .catch(error => {
              this.errors.responseError = error
              console.log(this.errors.responseError.toString())
              if (this.errors.responseError.toString() === "Error: badRequest") this.errorTitle = "Something went wrong. Try again later."
              if (this.errors.responseError.toString() === "Error: cardNotFound") this.errorTitle = "Сard not found in the system"
              if (this.errors.responseError.toString() === "Error: change is Not Approved") this.errorTitle = "Change is not approved"
              if (this.errors.responseError.toString() === "Error: userNotFound") this.errorTitle = "User not found in the system"
              console.log(this.errorTitle)
              this.$swal({
                title: this.errorTitle,
                type: "alert message",
                confirmButtonColor: "#f1416c",
                icon: "error",
              })
            })
      }
    }
  }
}
</script>

<style scoped>
/*.form-control {*/
/*  min-width: 100%;*/
/*}*/
.min-w-115px {
  min-width: 115px;
}

.form-check-input {
  margin-top: 0.1rem !important;
}
</style>

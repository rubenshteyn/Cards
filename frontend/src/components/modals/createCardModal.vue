<template>
  <!-- из шаблона -->
  <div v-if="show">
    <div :class="{show: this.isShow}" class="modal fade d-block" id="kt_modal_view_users" tabindex="-1"
         aria-hidden="true">
      <div class="modal-dialog modal-dialog-centered mw-850px" :style="{transform: this.transform}">
        <!--begin::Modal content-->
        <div class="modal-content">
          <!--begin::Modal header-->
          <div class="modal-header">
            <div class="text-center">
              <h1 class="text-dark fw-bolder">Creating a card</h1>
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
          <!--begin::Modal header-->
          <!--begin::Modal body-->
          <div v-if="getData()" class="modal-body scroll-y px-10 px-lg-15 pb-15">
            <div>
              <!--begin::List-->
              <div class="mh-400px scroll-y me-n7 pe-7">
                <div class="card-body p-0">
                  <div v-if="this.warning"
                       class="notice d-flex bg-light-warning rounded border-warning border border-dashed mb-9 p-6">
                    <!--begin::Icon-->
                    <!--begin::Svg Icon | path: icons/duotune/general/gen044.svg-->
                    <span class="svg-icon svg-icon-2tx svg-icon-warning me-4">
										<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
											<rect opacity="0.3" x="2" y="2" width="20" height="20" rx="10" fill="currentColor"></rect>
											<rect x="11" y="14" width="7" height="2" rx="1" transform="rotate(-90 11 14)"
                            fill="currentColor"></rect>
											<rect x="11" y="17" width="2" height="2" rx="1" transform="rotate(-90 11 17)"
                            fill="currentColor"></rect>
										</svg>
									</span>
                    <!--end::Svg Icon-->
                    <!--end::Icon-->
                    <!--begin::Wrapper-->
                    <div class="d-flex flex-stack flex-grow-1">
                      <!--begin::Content-->
                      <div class="fw-semibold">
                        <h4 class="text-gray-900 fw-bold">Warning</h4>
                        <div class="fs-6 text-gray-700">The selected user does not have enough funds</div>
                      </div>
                      <!--end::Content-->
                    </div>
                    <!--end::Wrapper-->
                  </div>
                  <div class="form-group fv-row mb-8">
                    <label class="required fs-6 fw-semibold mb-2" for="status">BIN</label>
                    <v-select class="bg-transparent " label="name" :options="BINsList"
                              placeholder="Select a BIN" v-model="selectedBIN" required>
                    </v-select>
                  </div>
                  <div class="form-group fv-row mb-8">
                    <label class="required fs-6 fw-semibold mb-2" for="limit">Limit</label>
                    <input type="number" class="form-control form-control-lg form-control-solid"
                           id="limit" v-model="form.limit" name="name" placeholder="Limit"
                           required/>
                  </div>
<!--                  <div class="form-group fv-row mb-8">-->
<!--                    <label class="required fs-6 fw-semibold mb-2" for="count">Count</label>-->
<!--                    <input type="number" class="form-control form-control-lg form-control-solid"-->
<!--                           id="limit" v-model="form.count" name="name" placeholder="Count"-->
<!--                           required/>-->
<!--                  </div>-->
                  <div class="form-group fv-row mb-8">
                    <label class="fs-6 fw-semibold mb-2" for="commission">Will be credited to the balance</label>
                    <div class="form-control form-control-lg form-control-solid">{{ resultLimit() }}</div>
                  </div>

                  <div class="form-group fv-row mb-8">
                    <label class="fs-6 fw-semibold mb-2" for="note">Card note</label>
                    <input type="text" class="form-control form-control-lg form-control-solid"
                           id="note" v-model="form.note" name="name" placeholder="Enter comment"
                           required/>
                  </div>
                  <div class="form-group fv-row mb-8 fs-6 fw-semibold">
                    Creation of cards:
                    <span
                        v-if="this.cardType">{{ this.cardType.sumToCreate }}</span>
                    <span v-else>0</span>$
                    +
                    <span v-if="this.form.limit">{{ this.form.limit }}</span>
                    <span v-else>0</span>
                    =
                    <span v-if="this.cardType && this.form.limit"> {{ (this.cardType.sumToCreate) + (this.form.limit) }}</span>
                    <span v-else>0</span>$
                    <span class="d-block">
                      Total:
                       <span v-if="this.cardType && this.form.limit">
                         {{ (this.cardType.sumToCreate) + (this.form.limit) }}
                       </span>
                       <span v-else>0</span>$
                    </span>
                  </div>
                  <div class="form-group fv-row mb-8 fs-6 fw-semibold">
                    <div v-if="outputResult()">There are enough funds on the balance to issue cards</div>
                    <div v-else style="color:red;">There are not enough funds on the balance</div>
                  </div>
                </div>
              </div>
              <!--end::List-->
            </div>
            <!--end::Users-->
          </div>
          <div class="modal-footer flex-center">
            <button @click="closeModal" class="btn btn-light me-3">Discard</button>
            <button :disabled="this.isDisabled" @click="checkForm()" type="submit" id="kt_modal_new_target_submit"
                    class="btn btn-primary">
              <span class="indicator-label">Creating a card</span>
              <span class="indicator-progress">Please wait...
									<span class="spinner-border spinner-border-sm align-middle ms-2"></span></span>
            </button>
          </div>
          <!--end::Modal body-->
        </div>
        <!--end::Modal content-->
      </div>
    </div>
  </div>
</template>

<script>
import {mapActions, mapGetters} from "vuex";

const initialState = () => {
  return {
    show: false,
    isShow: false,
    transform: "translate(0, -50px)",
    warning: false,
    isDisabled: true,
    field: true,
    selectedBalance: {code: "", amount: null, name: ""},
    selectedBIN: null,
    ref: null,
    errors: {
      responseError: null
    },
    cardType: null,
    endLimit: 0,
    form: {
      cardType: null,
      BIN: null,
      note: null,
      limit: 0,
    },
  }
}
export default {
  name: "createCardModal",
  data() {
    return initialState()
  },
  watch: {
    userinfo: "getData"
  },
  methods: {
    ...mapActions({
      createTaskCardCreation: "users/createTaskCardCreation"
    }),
    outputResult() {
      if(this.cardType && this.cardType.name === "Advertising" && this.userinfo.advertisingCommission) {
        this.cardType.commission = this.userinfo.advertisingCommission
      }
      if(this.cardType && this.cardType.name === "Universal" && this.userinfo.universalCommission) {
        this.cardType.commission = this.userinfo.universalCommission
      }
      if(this.cardType && this.cardType.name === "Facebook" && this.userinfo.facebookCommission) {
        this.cardType.commission = this.userinfo.facebookCommission
      }
      if ((this.cardType && this.form.limit && this.userinfo.USDTBalance) &&
          this.userinfo.USDTBalance > (this.cardType.sumToCreate) + (this.form.limit)) {
        this.isDisabled = false
        return true
      }
      if(this.selectedBIN && this.cardType && this.form.limit && this.userinfo.USDTBalance) {
        this.field = false
      }
      this.isDisabled = true
    },
    resultLimit() {
      if (this.cardType && this.userinfo.universalCost && this.userinfo.universalCommission && this.cardType.name === "Universal") {
        this.cardType.sumToCreate = this.userinfo.universalCost
        return this.endLimit = this.form.limit - ((this.form.limit * this.userinfo.universalCommission) / 100)
      }
      if (this.cardType && this.userinfo.advertisingCost && this.userinfo.advertisingCommission && this.cardType.name === "Advertising") {
        this.cardType.sumToCreate = this.userinfo.advertisingCost
        return this.endLimit = this.form.limit - ((this.form.limit * this.userinfo.advertisingCommission) / 100)
      }
      if (this.cardType && this.userinfo.facebookCost && this.userinfo.facebookCommission && this.cardType.name === "Facebook") {
        this.cardType.sumToCreate = this.userinfo.facebookCost
        return this.endLimit = this.form.limit - ((this.form.limit * this.userinfo.facebookCommission) / 100)
      }
      if(this.cardType) {
        return this.endLimit = this.form.limit - ((this.form.limit * this.cardType.commission) / 100)
      }
      return null
    },
    getData() {
      if (this.loggedIn && !this.userInfoLoaded) {
        this.$store.dispatch("users/getUserInfoMethod");
      }
      if (this.loggedIn && !this.binsLoaded) {
        this.$store.dispatch("users/getBinsMethod");
      }
      return true
    },
    openModal(ref) {
      this.show = true
      this.ref = ref
      if(ref === "universal") {
        this.cardType = {
          name: "Universal",
          commission: this.userinfo.universalCommission,
          sumToCreate: this.userinfo.universalCost,
        }
      }
      if(ref === "advertising") {
        this.cardType = {
          name: "Advertising",
          commission: this.userinfo.advertisingCommission,
          sumToCreate: this.userinfo.advertisingCost,
        }
      }
      if(ref === "facebook") {
        this.cardType = {
          name: "Facebook",
          commission: this.userinfo.facebookCommission,
          sumToCreate: this.userinfo.facebookCost,
        }
      }
      console.log(this.cardType, this.userinfo.facebookCommission, this.ref)
      setTimeout(() => {
        this.isShow = this.show
        this.transform = "none"
      }, 200)
    },
    closeModal() {
      this.isShow = false
      this.transform = "translate(0, -50px)"
      setTimeout(() => {
        this.show = false
      }, 200)
      this.selectedBalance = {code: "", name: ""}
      this.selectedBIN = null
      this.form.limit = 0
      this.endLimit = 0
      this.cardType = null
      this.form.cardType = null
      this.form.note = null
    },
    cardCommissionOutput() {
      if (this.cardType && this.user.advertisingCommission && this.cardType.name === "Advertising") {
        return this.user.advertisingCommission
      }
      if (this.cardType && this.user.universalCommission && this.cardType.name === "Universal") {
        return this.user.universalCommission
      }
      if (this.cardType) {
        return this.cardType.commission
      }
      return null
    },
    usersBalance(number) {
      if (number === 0 || !number) {
        return '0,00'
      }

      let result = []
      let arrayBalance = ("" + number).split("").map(Number)
      for (const i of arrayBalance) {
        result.push(i)
      }

      function roundingNum(x, n) {
        return parseFloat(Number.parseFloat(x).toFixed(n || 2));
      }

      let arrRoundingNum = Array.from(roundingNum(number, 2).toString())
      let endResult = Array.from(arrRoundingNum.toString().split("."))[1]
      let firstResult = Array.from(arrRoundingNum.toString().split(".")[0])
      if (firstResult.toString().replace(/[\s.,%]/g, '').split('').length > 3) {
        arrRoundingNum.splice(0, 1)
        return result.shift() + ',' + arrRoundingNum.join("")
      }
      if (typeof endResult == 'undefined') {
        return arrRoundingNum.join("") + ".00"
      }
      if (endResult.split('').length < 2) {
        return arrRoundingNum.join("") + ".0"
      }

      return arrRoundingNum.join("")
    },
    checkForm: function () {
      this.form.cardType = this.cardType.name
      this.form.BIN = this.selectedBIN.name.toString()
      if ((this.form.limit - this.cardType.sumToCreate) > this.userinfo.USDTBalance) {
        this.warning = true
        this.isDisabled = true
        return false
      }
      console.log(this.form, this.userinfo.USDTBalance)
      this.createTaskCardCreation(this.form)
          .then(() => {
            Object.assign(this.$data, initialState());
            this.closeModal()
            this.$swal({
              title: "You have successfully created a request to create a card",
              confirmButtonColor: "#50cd89",
              type: "alert message",
              icon: "success",
            })
          })
          .catch(error => {
            console.log(error)
            this.errors.responseError = error
            let errorTitle = error
            if(this.errors.responseError.toString() === "Error: badRequest") errorTitle = "Something went wrong. Try again later."
            if(this.errors.responseError.toString() === "Error: cardNotFound") errorTitle = "Сard not found in the system"
            if(this.errors.responseError.toString() === "Error: change is Not Approved") errorTitle = "Change is not approved"
            if(this.errors.responseError.toString() === "Error: userNotFound") errorTitle = "User not found in the system"
            if(this.errors.responseError.toString() === "Error: bins error") errorTitle = "Bean is not correct"
            if(this.errors.responseError.toString() === "Error: paying not from USDT balance is disabled for now") errorTitle = "Paying not from USDT balance is disabled for now"
            if(this.errors.responseError.toString() === "Error: noAvailableAddresses") errorTitle = "These addresses are not in the system"
            if(this.errors.responseError.toString() === "Error: card is Not Approved") errorTitle = "Card is not approved"
            this.$swal({
              title: errorTitle,
              type: "alert message",
              confirmButtonColor: "#f1416c",
              icon: "error",
            })
          });
    },
  },
  computed: {
    ...mapGetters({
      userinfo: "users/getUserInfo",
      userInfoLoaded: "users/getUserInfoLoaded",
      bins: "users/getBins",
      binsLoaded: "users/getBinsLoaded",
      loggedIn: "auth/isLoggedIn"
    }),
    BINsList() {
      let BINsList = []
      if (this.bins.length > 0) {
        let providerBINs = JSON.parse(this.bins)
        console.log(providerBINs)
        providerBINs.forEach(bin => {
          if (this.ref === "universal" && !bin.isUniversal) {
            return
          }
          if (this.ref !== "universal" && bin.isUniversal) {
            return
          }
          let legitBin = {name: bin.bin, isUniversal: bin.isUniversal, active: bin.active}
          BINsList.push(legitBin)
        })
      }
      return BINsList
    },
  }
}
</script>

<style scoped>
.form-check-input {
  margin-top: -0.4rem;
}
</style>

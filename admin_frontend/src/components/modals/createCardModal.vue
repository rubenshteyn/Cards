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
          <div class="modal-body scroll-y px-10 px-lg-15 pb-15">
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
                    <label class="required fs-6 fw-semibold mb-2" for="status">Client</label>
                    <v-select class="bg-transparent" label="name" :options="usersList"
                              placeholder="Select a client" v-model="selectedUser" required>
                    </v-select>
                  </div>
                  <div class="form-group fv-row mb-8">
                    <label class="required fs-6 fw-semibold mb-2" for="status">Provider</label>
                    <v-select class="bg-transparent" label="name" :options="providersList"
                              placeholder="Select a provider" v-model="selectedProvider" required>
                    </v-select>
                  </div>
                  <div class="form-group fv-row mb-8">
                    <label class="required fs-6 fw-semibold mb-2" for="status">Type card</label>
                    <v-select class="bg-transparent " label="name" :options="cardTypes"
                              placeholder="Select a type card" v-model="cardType" required>
                    </v-select>
                  </div>

                  <div class="form-group fv-row mb-8">
                    <label class="fs-6 fw-semibold mb-2" for="commission">Commission(%)</label>
                    <div class="form-control form-control-lg form-control-solid">{{ cardCommissionOutput() }}</div>
                  </div>

                  <div class="form-group fv-row mb-8">
                    <label class="mb-2 required fs-6 fw-semibold mb-2">Payment system type</label>
                    <div class="ml-4 d-flex">
                      <div class="d-flex align-items-center mr-5">
                        <input @click="resetBIN()" class="form-check-input" type="radio" name="exRadios" id="visa"
                               :value="'VISA'"
                               style="width: 1em; height: 1em;" v-model="systemType" :checked="systemType === 'VISA'">
                        <label class="form-check-label fs-6 fw-semibold mb-2" for="visa">
                          VISA
                        </label>
                      </div>
                      <div class="mr-5 d-flex align-items-center">
                        <input @click="resetBIN()" class="form-check-input" type="radio" name="exRadios" id="mastercard"
                               :value="'MasterCard'"
                               style="width: 1em; height: 1em;" v-model="systemType"
                               :checked="systemType === 'MasterCard'">
                        <label class="form-check-label fs-6 fw-semibold mb-2" for="mastercard">
                          MasterCard
                        </label>
                      </div>
                    </div>
                  </div>

                  <div class="form-group fv-row mb-8">
                    <label class="required fs-6 fw-semibold mb-2" for="status">BIN</label>
                    <v-select class="bg-transparent " label="name" :options="BINsList"
                              placeholder="Select a BIN" v-model="selectedBIN" required>
                    </v-select>
                  </div>
                  <div class="form-group fv-row mb-8">
                    <label class="required fs-6 fw-semibold mb-2" for="balanceList">Debit the account</label>
                    <v-select class="bg-transparent " label="name" :options="balanceList"
                              placeholder="Select a account" v-model="selectedBalance" required>
                    </v-select>
                  </div>
                  <div class="form-group fv-row mb-8">
                    <label class="required fs-6 fw-semibold mb-2" for="amount">Amount cards</label>
                    <input type="number" class="form-control form-control-lg form-control-solid"
                           id="amount" v-model="form.count" name="name" placeholder="Amount cards"
                           required/>
                  </div>
                  <div class="form-group fv-row mb-8">
                    <label class="required fs-6 fw-semibold mb-2" for="limit">Limit</label>
                    <input type="number" class="form-control form-control-lg form-control-solid"
                           id="limit" v-model="form.limit" name="name" placeholder="Limit"
                           required/>
                  </div>
                  <div class="form-group fv-row mb-8">
                    <label class="fs-6 fw-semibold mb-2" for="commission">Will be credited to the balance</label>
                    <div class="form-control form-control-lg form-control-solid">{{ resultLimit() }}</div>
                  </div>

                  <!--                <div class="form-group fv-row mb-8">-->
                  <!--                  <label class="required fs-6 fw-semibold mb-2" for="expMonth">expMonth</label>-->
                  <!--                  <input type="number" class="form-control form-control-lg form-control-solid"-->
                  <!--                         id="expMonth" v-model="form.expMonth" name="name" placeholder="expMonth"-->
                  <!--                         required/>-->
                  <!--                </div>-->
                  <!--                <div class="form-group fv-row mb-8">-->
                  <!--                  <label class="required fs-6 fw-semibold mb-2" for="expYear">expYear</label>-->
                  <!--                  <input type="number" class="form-control form-control-lg form-control-solid"-->
                  <!--                         id="expYear" v-model="form.expYear" name="name" placeholder="expYear"-->
                  <!--                         required/>-->
                  <!--                </div>-->
                  <div class="form-group fv-row mb-8">
                    <label class="fs-6 fw-semibold mb-2" for="limit">Card note</label>
                    <input type="text" class="form-control form-control-lg form-control-solid"
                           id="note" v-model="form.note" name="name" placeholder="Enter comment"
                           required/>
                  </div>
                  <div class="form-group fv-row mb-8 fs-6 fw-semibold">
                    Creation of cards:
                    <span v-if="this.cardType">{{ this.cardType.sumToCreate }}</span>
                    <span v-else>0</span>$*
                    <span v-if="this.form.count">{{ this.form.count }}</span>
                    <span v-else>0</span>
                    =
                    <span
                        v-if="this.cardType && this.form.count">{{ this.cardType.sumToCreate * this.form.count }}</span>
                    <span v-else>0</span>$
                    + Replenishment of cards:
                    <span v-if="this.form.limit">{{ this.form.limit }}</span>
                    <span v-else>0</span>$*

                    <span v-if="this.form.count">{{ this.form.count }}</span>
                    <span v-else>0</span>
                    =
                    <span v-if="this.form.limit && this.form.count">{{ this.form.limit * this.form.count }}</span>
                    <span v-else>0</span>$
                    <span class="d-block">
                      Total:
                       <span v-if="this.cardType && this.form.count && this.form.limit">
                         {{ (this.cardType.sumToCreate * this.form.count) + (this.form.limit * this.form.count) }}
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
    selectedUser: {code: 0, name: "", user: {}},
    selectedBalance: {code: "", amount: null, name: ""},
    selectedProvider: {code: 0, name: "", provider: {}},
    selectedBIN: {name: "", isUniversal: null, active: null},
    errors: {
      responseError: null
    },
    cardTypes: [
      {
        name: "Universal",
        commission: 8,
        sumToCreate: 15,
        code: true
      },
      {
        name: "Advertising",
        commission: 6,
        sumToCreate: 10,
        code: false
      },
      {
        name: "Facebook",
        commission: 2.8,
        sumToCreate: 2,
        code: false
      }
    ],
    cardType: null,
    systemTypes: [
      {
        bin: "5",
        code: 5,
        name: "MasterCard"
      },
      {
        bin: "4",
        code: 4,
        name: "VISA"
      }
    ],
    systemType: null,
    endLimit: 0,
    form: {
      userID: null,
      note: null,
      cardProvider: null,
      cardType: null,
      balance: null,
      systemType: null,
      BIN: null,
      limit: null,
      count: null,
    },
  }
}
export default {
  name: "createCard",
  data() {
    return initialState()
  },
  watch: {
    "selectedBIN": "selectTypesFromBin"
  },
  methods: {
    ...mapActions({
      createCards: "admin/createCards"
    }),
    outputResult() {
      if(this.cardType && this.cardType.name === "Advertising" && this.selectedUser.user.advertisingCommission) {
        this.cardType.commission = this.selectedUser.user.advertisingCommission
      }
      if(this.cardType && this.cardType.name === "Universal" && this.selectedUser.user.universalCommission) {
        this.cardType.commission = this.selectedUser.user.universalCommission
      }
      if(this.cardType && this.cardType.name === "Facebook" && this.selectedUser.user.facebookCommission) {
        this.cardType.commission = this.selectedUser.user.facebookCommission
      }
      if ((this.cardType && this.form.count && this.form.limit && this.selectedBalance.amount) &&
          this.selectedBalance.amount > (this.cardType.sumToCreate * this.form.count) + (this.form.limit * this.form.count)) {
        this.isDisabled = false
        return true
      }
      // if(this.selectedBIN && this.selectedUser.name && this.cardType && this.form.count && this.form.limit && this.selectedBalance.amount) {
      //   console.log(this.selectedBIN)
      //   this.field = false
      // }
        this.isDisabled = true
    },
    resetBIN() {
      this.selectedBIN = null
      this.form.BIN = null
    },
    selectTypesFromBin() {
      if (this.selectedBIN && this.selectedBIN.name.toString().split("")[0] === "4") {
        this.systemType = "VISA"
      }
      if (this.selectedBIN && this.selectedBIN.name.toString().split("")[0] === "5") {
        this.systemType = "MasterCard"
      }
      if (this.selectedBIN && this.selectedBIN.isUniversal) {
        this.cardType = {
          name: "Universal",
          commission: 8,
          sumToCreate: 15,
          code: true
        }
      }
      // if (this.selectedBIN && !this.selectedBIN.isUniversal) {
      //   this.cardType = {
      //     name: "Advertising",
      //     commission: 6,
      //     sumToCreate: 10,
      //     code: false
      //   }
      // }
    },
    resultLimit() {
      if (this.cardType && this.selectedUser.user.universalCost && this.selectedUser.user.universalCommission && this.cardType.name === "Universal") {
        this.cardType.sumToCreate = this.selectedUser.user.universalCost
        return this.endLimit = this.form.limit - ((this.form.limit * this.selectedUser.user.universalCommission) / 100)
      }
      if (this.cardType && this.selectedUser.user.advertisingCost && this.selectedUser.user.advertisingCommission && this.cardType.name === "Advertising") {
        this.cardType.sumToCreate = this.selectedUser.user.advertisingCost
        return this.endLimit = this.form.limit - ((this.form.limit * this.selectedUser.user.advertisingCommission) / 100)
      }
      if(this.cardType) {
        return this.endLimit = this.form.limit - ((this.form.limit * this.cardType.commission) / 100)
      }
      return null
    },
    openModal() {
      this.show = true
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
      this.selectedUser = {code: 0, name: "", user: {}}
      this.selectedBalance = {code: "", name: ""}
      this.selectedBIN = null
      this.form.limit = 0
      this.form.count = 1
      this.endLimit = 0
      this.systemType = null
      this.cardType = null
      this.form.cardType = null
      this.form.note = null
    },
    cardCommissionOutput() {
      if (this.cardType && this.selectedUser.user.advertisingCommission && this.cardType.name === "Advertising") {
        return this.selectedUser.user.advertisingCommission
      }
      if (this.cardType && this.selectedUser.user.universalCommission && this.cardType.name === "Universal") {
        return this.selectedUser.user.universalCommission
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
      this.form.userID = this.selectedUser.user.ID
      this.form.cardProvider = this.selectedProvider.provider.name
      this.form.cardType = this.cardType.name
      this.form.balance = this.selectedBalance.code
      if (this.selectedBIN) {
        this.form.BIN = this.selectedBIN.name.toString()
      }
      this.form.systemType = this.systemType
      if ((this.limit - this.cardType.sumToCreate) > this.selectedUser.user.USDTBalance) {
        this.warning = true
        this.isDisabled = true
        return false
      }
      console.log(this.form)
      this.createCards(this.form)
          .then(() => {
            Object.assign(this.$data, initialState());
            this.closeModal()
            this.$swal({
              title: "You have successfully created a card",
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
      users: "admin/getUsers",
      suppliers: "admin/getSuppliers",
    }),

    usersList() {
      let usersList = []
      this.users.forEach(user => {
        let userForList = {
          code: user.ID,
          name: "ID:" + " " + user.ID + " - " + user.name + " - " + user.email,
          user: user
        }
        if (user.status === 2) {
          usersList.push(userForList)
        }
      })
      return usersList
    },
    BINsList() {
      let BINsList = []
      let provider = this.selectedProvider.provider
      let code = {
        "VISA": "4",
        "MasterCard": "5"
      }
      if (this.selectedProvider.name) {
        let providerBINs = JSON.parse(provider.bins)
        providerBINs.forEach(bin => {
          if (!bin.active) {
            return
          }
          if (this.systemType !== null && code[this.systemType] !== bin.bin.toString().split("")[0]) {
            return
          }
          if (this.cardType !== null && this.cardType.code !== bin.isUniversal) {
            return
          }
          let legitBin = {name: bin.bin, isUniversal: bin.isUniversal, active: bin.active}
          BINsList.push(legitBin)
        })
      }
      return BINsList
    },
    providersList() {
      let providersList = []
      if (this.suppliers.length === 0) {
        this.$store.dispatch("admin/getSuppliersMethod")
      }
      this.suppliers.forEach(provider => {
        let providerForList = {code: provider.status, name: provider.name, provider: provider}
        if (provider.ID === 1) {
          this.selectedProvider = {code: provider.status, name: provider.name, provider: provider}
        }
        providersList.push(providerForList)
      })
      return providersList
    },

    balanceList() {
      let balances = []
      let user = this.selectedUser.user
      if (user) {
        balances = [
          {"code": "USDT", "amount": user.USDTBalance, "name": "USDT - " + this.usersBalance(user.USDTBalance) + " ₮"},
          {"code": "USD", "amount": user.USDBalance, "name": "USD - " + this.usersBalance(user.USDBalance) + " $"},
          {"code": "EUR", "amount": user.EURBalance, "name": "EUR - " + this.usersBalance(user.EURBalance) + " €"},
          {"code": "BTC", "amount": user.BTCBalance, "name": "BTC - " + this.usersBalance(user.BTCBalance) + " ₿"},
        ]
      }
      return balances
    }
  }
}
</script>

<style scoped>
.form-check-input {
  margin-top: -0.4rem;
}
</style>

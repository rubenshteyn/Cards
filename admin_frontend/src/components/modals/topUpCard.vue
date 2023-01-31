<template>
  <!-- из шаблона -->
  <div v-if="show">
    <div :class="{show: this.isShow}" class="modal fade d-block" @click.self="closeModal" id="kt_modal_view_users" tabindex="-1"
         aria-hidden="true">
      <div class="modal-dialog modal-dialog-centered mw-850px" :style="{transform: this.transform}">
        <!--begin::Modal content-->
        <div class="modal-content">
          <!--begin::Modal header-->
          <div class="modal-header">
            <div class="text-center">
              <h1 class="text-dark fw-bolder">Top up card ID: {{ this.form.cardID }}</h1>
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
          <div class="modal-body scroll-y">
            <div>
              <!--begin::List-->
              <div class="mh-500px scroll-y me-n7 pe-7">
                <div class="card-body p-0">
                  <div class="form-group fv-row mb-8">
                    <label class="fs-6 fw-semibold mb-2" for="ID">Card number</label>
                    <div class="form-control form-control-lg form-control-solid">
                      {{
                        this.cardNumber.toString().substring(0, 4)
                        + " " + this.cardNumber.toString().substring(4, 8)
                        + " " + this.cardNumber.toString().substring(8, 12)
                        + " " + this.cardNumber.toString().substring(12, 16)
                      }} - {{ systemTypeOutput() }}
                    </div>
                  </div>
                  <div class="row mb-8">
                    <div class="form-group fv-row col-md-6">
                      <label class="fs-6 fw-semibold mb-2" for="ID">User</label>
                      <div class="form-control form-control-lg form-control-solid">
                        ID: {{ this.user.ID }} - {{ this.user.name }} - {{ this.user.email }}
                      </div>
                    </div>
                    <div class="form-group fv-row col-md-6">
                      <label class="fs-6 fw-semibold mb-2" for="ID">Card type</label>
                      <div v-if="!this.cardNumber.toString().includes('486555')"
                           class="form-control form-control-lg form-control-solid">
                        Advertising
                      </div>
                      <div v-else class="form-control form-control-lg form-control-solid">Universal</div>
                    </div>
                  </div>
                  <div class="form-group fv-row mb-8">
                    <label class="fs-6 fw-semibold mb-2" for="status">Select account</label>
                    <v-select class="bg-transparent " label="name" :options="balanceList"
                              placeholder="Select a account" v-model="balance" required>
                    </v-select>
                  </div>
                  <div class="form-group fv-row mb-8">
                    <label class="fs-6 fw-semibold mb-2" for="amount">Amount</label>
                    <input type="number" class="form-control form-control-lg form-control-solid"
                           id="amount" v-model="form.amount" name="name" placeholder="Amount"
                           required/>
                    <span class="text-muted fw-semibold text-muted d-block fs-7" style="padding: 0.825rem 1.5rem;">
                    {{ commissionOutput() }} will be credited to the card</span>
                  </div>
                  <div class="form-group fv-row mb-8 fs-6 fw-semibold">
                    <div v-if="outputResult()" style="color:red;">There are not enough funds on the balance</div>
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
              <span class="indicator-label">Top up</span>
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
    isDisabled: true,
    show: false,
    errorTitle: null,
    errors: {
      responseError: null
    },
    isShow: false,
    transform: "translate(0, -50px)",
    user: null,
    balances: [
      {
        name: "USDT"
      }
      ,
      {
        name: "USD"
      },
      {
        name: "BTC"
      },
      {
        name: "EUR"
      },
    ],
    balance: null,
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
    systemType: {
      name: null,
      code: null,
      bin: null
    },
    cardNumber: null,
    form: {
      cardID: null,
      balance: null,
      amount: null,
    }
  }
}

export default {
  name: "topUpCard",
  data() {
    return initialState()
  },
  methods: {
    ...mapActions({
      topUpCardBalance: "admin/topUpCardBalance"
    }),
    outputResult() {
      if (!this.form.amount || this.balance.balance < this.form.amount) {
        this.isDisabled = true
        return true
      }
      this.isDisabled = false
    },
    openModal(card) {
      this.show = true
      setTimeout(()=> {
        this.isShow = this.show
        this.transform = "none"
      },200)
      this.form.cardID = card.ID
      this.cardNumber = card.cardNumber
      let users = Object.values(this.users)
      let user = users.find(item => item.ID === card.userId)
      this.user = user
      this.balance = {"code": "USDT", "balance": this.user.USDTBalance, "name": "USDT - " + this.usersBalance(this.user.USDTBalance) + " ₮"}
    },
    closeModal() {
      this.isShow = false
      this.transform = "translate(0, -50px)"
      setTimeout(()=> {
        this.show = false
      },200)
    },
    commissionOutput() {
      if (this.cardNumber.toString().includes('486555') && this.user.universalCommission) {
        return this.result = this.form.amount - ((this.form.amount * this.user.universalCommission) / 100)
      }
      if (this.user.advertisingCommission) {
        return this.result = this.form.amount - ((this.form.amount * this.user.advertisingCommission) / 100)
      }
      if (this.cardNumber.toString().includes('486555')) {
        return this.result = this.form.amount - ((this.form.amount * 8) / 100)
      }
      return this.result = this.form.amount - ((this.form.amount * 6) / 100)
    },
    systemTypeOutput() {
      for (const systemType of this.systemTypes) {
        if (this.cardNumber.toString().split('')[0] === systemType.bin) {
          return systemType.name
        }
      }
    },
    checkForm() {
      this.form.balance = this.balance.name.split(" ")[0]
      console.log(this.form)
      this.topUpCardBalance(this.form)
          .then(() => {
            Object.assign(this.$data, initialState());
            this.closeModal()
            this.$swal({
              title: "You have successfully topped up the card",
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
    },
    usersBalance(number) {
      if (number === 0) {
        return '00,00'
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
    }
  },
  computed: {
    ...mapGetters({
      users: "admin/getUsers"
    }),
    balanceList() {
      let balances = []
      if (this.user) {
        balances = [
          {"code": "USDT", "balance": this.user.USDTBalance, "name": "USDT - " + this.usersBalance(this.user.USDTBalance) + " ₮"},
          {"code": "BTC", "balance": this.user.BTCBalance, "name": "BTC - " + this.usersBalance(this.user.BTCBalance) + " ₿"},
          {"code": "EUR", "balance": this.user.EURBalance, "name": "EUR - " + this.usersBalance(this.user.EURBalance) + " €"},
          {"code": "USD", "balance": this.user.USDBalance, "name": "USD - " + this.usersBalance(this.user.USDBalance) + " $"},
        ]
      }
      return balances
    }
  }
}
</script>

<style scoped>

</style>
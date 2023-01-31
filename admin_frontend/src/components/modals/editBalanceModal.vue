<template>
  <!-- из шаблона -->
  <div v-if="show">
    <div :class="{show: this.isShow}" class="modal fade d-block"  id="kt_modal_view_users" tabindex="-1"
         aria-hidden="true">
      <div class="modal-dialog modal-dialog-centered mw-550px" :style="{transform: this.transform}">
        <!--begin::Modal content-->
        <div class="modal-content">
          <!--begin::Modal header-->
          <div class="modal-header">
            <h1 class="text-dark fw-bolder">Top up balance</h1>
            <div @click="closeModal" class="btn btn-sm btn-icon btn-active-color-primary"
                 data-bs-dismiss="modal">
              <!--begin::Svg Icon | path: icons/duotune/arrows/arr061.svg-->
              <span class="svg-icon svg-icon-1">
                            <svg width="24" height="24" viewBox="0 0 24 24" fill="none"
                                 xmlns="http://www.w3.org/2000/svg">
                                <rect opacity="0.5" x="6" y="17.3137" width="16" height="2" rx="1"
                                      transform="rotate(-45 6 17.3137)" fill="currentColor" />
                                <rect x="7.41422" y="6" width="16" height="2" rx="1" transform="rotate(45 7.41422 6)"
                                      fill="currentColor" />
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
                      <form class="form w-100" id="registration"
                            @submit.prevent="checkForm">

                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="account">Client</label>
                          <v-select class="bg-transparent " label="name" :options="usersList"
                                    placeholder="Select client" v-model="selectedUser" required>
                          </v-select>
                        </div>

                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="balance">Account selection</label>
                          <v-select class="bg-transparent " label="name" :options="balanceList"
                                    placeholder="Select account" v-model="selectedBalance" required>
                          </v-select>
                        </div>

                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="sum">Replenishment amount</label>
                          <input type="number"
                                 class="form-control form-control-lg form-control-solid" id="sum"
                                 v-model="form.sum"
                                 required
                                 name="sum" placeholder="Enter sum"/>
                        </div>

                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="commission">Commission (%)</label>
                          <input type="number"
                                 class="form-control form-control-lg form-control-solid" id="commission"
                                 v-model="form.commission"
                                 required
                                 name="commission" :placeholder=form.commission />
                        </div>

                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="amount">Total amount</label>
                          <div class="form-control-lg form-control form-control-solid" style="width: fit-content">{{endSum()}}</div>
                        </div>

                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="transaction">Transaction</label>
                          <input type="text" class="form-control form-control-lg form-control-solid"
                                 id="transaction" v-model="form.transaction" name="transaction"
                                 placeholder="Transaction link" required />
                        </div>

                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="transaction">Payment note</label>
                          <input type="text" class="form-control form-control-lg form-control-solid"
                                 id="transaction" v-model="form.notes" name="transaction"
                                 placeholder="Enter payment note" required />
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
              <span class="indicator-label">Top up</span>
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
  mapActions, mapGetters
} from "vuex";
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
      id: 0,
      balance: null,
      amount: null,
      sum: null,
      commission: 1,
      transaction: null,
      notes: null,
    },
    selectedUser: {code: 0, name: "", user: {}},
    selectedBalance: {code: "", name: ""},
  }
}

export default defineComponent({
  name: "editBalanceModal",
  data() {
    return initialState();
  },
  methods: {
    isDisabled() {
      return !(this.selectedBalance.name && this.form.sum && this.form.transaction && this.form.notes);
    },
    endSum() {
      this.form.amount = this.form.sum - ((this.form.sum * this.form.commission) / 100)
      return this.form.amount
    },
    openModal() {
      Object.assign(this.$data, initialState());
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
    openModalUser(user, balance, nameBalance, currency) {
      Object.assign(this.$data, initialState());
      this.show = true
      setTimeout(()=> {
        this.isShow = this.show
        this.transform = "none"
      },200)
      this.form.id = user.ID
      this.selectedUser = {code: user.ID, name: "ID: " + user.ID + " - " + user.name + " — " + user.email, user: user}
      this.selectedBalance = {code: nameBalance, name: nameBalance + " - " + balance + " ₮"}
      this.selectedBalance = {code: nameBalance, balance: balance, name: nameBalance + " - " + this.usersBalance(balance) + " " + currency}
    },
    ...mapActions({
      topUpBalance: "admin/topUpBalance"
    }),
    checkForm: function (e) {
      this.form.balance = this.selectedBalance.code
      this.form.id = this.selectedUser.code
      this.topUpBalance(this.form)
          .then(() => {
            Object.assign(this.$data, initialState());
            this.closeModal()
            this.$swal({
              title: "You have replenished the balance",
              confirmButtonColor: "#50cd89",
              type: "alert message",
              icon: "success",
            })
          })
          .catch(error => {
            console.log(error)
            this.errors.responseError = error
            let errorTitle = error
            if(error === "badRequest") errorTitle = "Something went wrong. Try again later."
            if(error === "cardNotFound") errorTitle = "Сard not found in the system"
            if(error === "change is Not Approved") errorTitle = "Change is not approved"
            if(error === "userNotFound") errorTitle = "User not found in the system"
            this.$swal({
              title: errorTitle,
              type: "alert message",
              confirmButtonColor: "#f1416c",
              icon: "error",
            })
          });
      if (this.errors.responseError) {
        e.preventDefault();
      }
    },
    usersBalance(number) {
      if (number === 0) {
        return '00.00'
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
      users: "admin/getUsers",
    }),
    usersList() {
      let usersList = []
      this.users.forEach(user => {
        let userForList = {code: user.ID, name:"ID: "+ user.ID + " - " + user.name + " - " + user.email, user: user}
        if (user.status === 2) {
          usersList.push(userForList)
        }
      })
      return usersList
    },
    balanceList(){
      let balances = []
      let user = this.selectedUser.user
      if (user) {
        balances = [
          {"code": "USDT", "balance": user.USDTBalance, "name": "USDT - " + this.usersBalance(user.USDTBalance) + " ₮"},
          {"code": "BTC", "balance": user.BTCBalance, "name": "BTC - " + this.usersBalance(user.BTCBalance) + " ₿"},
          {"code": "EUR", "balance": user.EURBalance, "name": "EUR - " + this.usersBalance(user.EURBalance) + " €"},
          {"code": "USD", "balance": user.USDBalance, "name": "USD - " + this.usersBalance(user.USDBalance) + " $"},
        ]
      }
      return balances
    }
  }
})
</script>

<style>
.table.gs-0 th:first-child,
.table.gs-0 td:first-child {
  padding-left: 1rem !important;
}

.modal {
  background: rgba(0, 0, 0, .5);
}
</style>
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
            <h1 v-if="styles.creation === 'block'" class="text-dark fw-bolder">Create card</h1>
            <h1 v-else class="text-dark fw-bolder">Replenishment card</h1>
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

          <div v-if="getData()" class="modal-body scroll-y">
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
                            <input @click="tabClicked('creation')" class="form-check-input" type="radio"
                                   name="exampleRadios"
                                   id="exampleRadios1" value="option1" style="width: 1em; height: 1em;" checked>
                            <label class="form-check-label edit-card-label" for="exampleRadios1">
                              Creation card
                            </label>
                          </div>
                          <div class="d-flex align-items-center mr-5">
                            <input @click="tabClicked('replenishment')" class="form-check-input" type="radio"
                                   name="exampleRadios"
                                   id="exampleRadios2" value="option2" style="width: 1em; height: 1em;">
                            <label class="form-check-label" for="exampleRadios2">
                              Replenishment card
                            </label>
                          </div>
<!--                          <div class="d-flex align-items-center mr-5">-->
<!--                            <input @click="tabClicked('replenishment')" class="form-check-input" type="radio"-->
<!--                                   name="exampleRadios"-->
<!--                                   id="exampleRadios2" value="option2" style="width: 1em; height: 1em;">-->
<!--                            <label class="form-check-label" for="exampleRadios2">-->
<!--                              Change replenishment-->
<!--                            </label>-->
<!--                          </div>-->
<!--                          <div class="d-flex align-items-center mr-5">-->
<!--                            <input @click="tabClicked('accountReplenishment')" class="form-check-input" type="radio"-->
<!--                                   name="exampleRadios"-->
<!--                                   id="exampleRadios2" value="option2" style="width: 1em; height: 1em;">-->
<!--                            <label class="form-check-label" for="exampleRadios2">-->
<!--                              Account replenishment-->
<!--                            </label>-->
<!--                          </div>-->
                        </div>
                      </div>
                      <form class="form w-100" :style="{ display: styles.creation }">
                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="status">Provider</label>
                          <v-select class="bg-transparent " label="name" :options="providersList"
                                    placeholder="Select provider" v-model="selectedProvider" required>
                          </v-select>
                        </div>
                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="status">Type card</label>
                          <v-select class="bg-transparent " label="name" :options="cardTypes"
                                    placeholder="Select card type" v-model="cardType" required>
                          </v-select>
                        </div>
                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="status">BIN</label>
                          <v-select class="bg-transparent " label="name" :options="BINsList"
                                    placeholder="Select a BIN" v-model="selectedBIN" required>
                          </v-select>
                        </div>
                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="name">Account selection</label>
                          <v-select class="bg-transparent " label="name" :options="balanceList"
                                    placeholder="Select account" v-model="selectedBalance" required>
                          </v-select>
                        </div>

                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="name">Note</label>
                          <input type="text" class="form-control form-control-lg form-control-solid"
                                 v-model="formCreationCard.note" name="name" placeholder="Enter note"
                                 required/>
                        </div>

                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="email">Limit</label>
                          <input type="number" class="form-control form-control-lg form-control-solid"
                                 v-model="formCreationCard.limit" name="email"
                                 placeholder="Enter limit" required/>
                        </div>
                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="email">Count</label>
                          <input type="number" class="form-control form-control-lg form-control-solid"
                                 v-model="formCreationCard.count" name="email"
                                 placeholder="Enter count" required/>
                        </div>
                        <div class="form-group fv-row mb-8 fs-6 fw-semibold">
                          Creation of cards:
                          <span v-if="this.cardType">{{ this.cardType.sumToCreate }}</span>
                          <span v-else>0</span>$*
                          <span v-if="this.formCreationCard.count">{{ this.formCreationCard.count }}</span>
                          <span v-else>0</span>
                          =
                          <span
                              v-if="this.cardType && this.formCreationCard.count">{{
                              this.cardType.sumToCreate * this.formCreationCard.count
                            }}</span>
                          <span v-else>0</span>$
                          + Replenishment of cards:
                          <span v-if="this.formCreationCard.limit">{{ this.formCreationCard.limit }}</span>
                          <span v-else>0</span>$*

                          <span v-if="this.formCreationCard.count">{{ this.formCreationCard.count }}</span>
                          <span v-else>0</span>
                          =
                          <span v-if="this.formCreationCard.limit && this.formCreationCard.count">{{
                              this.formCreationCard.limit * this.formCreationCard.count
                            }}</span>
                          <span v-else>0</span>$
                          <span class="d-block">
                      Total:
                       <span v-if="this.cardType && this.formCreationCard.count && this.formCreationCard.limit">
                         {{
                           (this.cardType.sumToCreate * this.formCreationCard.count) + (this.formCreationCard.limit * this.formCreationCard.count)
                         }}
                       </span>
                       <span v-else>0</span>$
                    </span>
                        </div>
                        <div class="form-group fv-row mb-8 fs-6 fw-semibold">
                          <div v-if="outputResult()">There are enough funds on the balance to issue cards</div>
                          <div v-else style="color:red;">There are not enough funds on the balance</div>
                        </div>
                      </form>
                      <form class="form w-100" :style="{ display: styles.replenishment }">
                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="status">Client</label>
                          <v-select class="bg-transparent " label="name" :options="usersList"
                                    placeholder="Select client" v-model="selectedUser" required>
                          </v-select>
                        </div>
                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="name">Card selection</label>
                          <v-select class="bg-transparent " label="name" :options="allCardList"
                                    placeholder="Select card" v-model="selectedCard" required>
                          </v-select>
                        </div>
                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="name">Account selection</label>
                          <v-select class="bg-transparent " label="name" :options="balanceList"
                                    placeholder="Select account" v-model="selectedBalance" required>
                          </v-select>
                        </div>

                        <div class="form-group fv-row mb-8">
                          <label class="required fs-6 fw-semibold mb-2" for="email">Amount</label>
                          <input type="email" class="form-control form-control-lg form-control-solid"
                                 v-model="formReplenishmentCard.amount" name="email"
                                 placeholder="Enter amount" required/>
                        </div>
                        <div class="form-group fv-row mb-8 fs-6 fw-semibold">
                          <div v-if="outputResult()">There are enough funds on the balance to issue cards</div>
                          <div v-else style="color:red;">There are not enough funds on the balance</div>
                        </div>
                      </form>
                      <!--                      <form class="form w-100" :style="{ display: styles.change }">-->

                      <!--                        <div class="form-group fv-row mb-8">-->
                      <!--                          <label class="required fs-6 fw-semibold mb-2" for="name">Username</label>-->
                      <!--                          <input type="text" class="form-control form-control-lg form-control-solid"-->
                      <!--                                 v-model="formData.name" name="name" placeholder="Enter username"-->
                      <!--                                 required/>-->
                      <!--                        </div>-->

                      <!--                        <div class="form-group fv-row mb-8">-->
                      <!--                          <label class="required fs-6 fw-semibold mb-2" for="email">User email</label>-->
                      <!--                          <input type="email" class="form-control form-control-lg form-control-solid"-->
                      <!--                                 v-model="formData.email" name="email"-->
                      <!--                                 placeholder="Enter e-mail" required/>-->
                      <!--                        </div>-->

                      <!--                        <div class="form-group fv-row mb-8">-->
                      <!--                          <label class="required fs-6 fw-semibold mb-2" for="status">User status</label>-->
                      <!--                          <v-select class="bg-transparent " label="name" :options="statuses"-->
                      <!--                                    placeholder="Select user status" v-model="status" required>-->
                      <!--                          </v-select>-->
                      <!--                        </div>-->
                      <!--                      </form >-->
                      <!--                      <form class="form w-100" :style="{ display: styles.accountReplenishment }">-->

                      <!--                        <div class="form-group fv-row mb-8">-->
                      <!--                          <label class="required fs-6 fw-semibold mb-2" for="name">Username</label>-->
                      <!--                          <input type="text" class="form-control form-control-lg form-control-solid"-->
                      <!--                                 v-model="formData.name" name="name" placeholder="Enter username"-->
                      <!--                                 required/>-->
                      <!--                        </div>-->

                      <!--                        <div class="form-group fv-row mb-8">-->
                      <!--                          <label class="required fs-6 fw-semibold mb-2" for="email">User email</label>-->
                      <!--                          <input type="email" class="form-control form-control-lg form-control-solid"-->
                      <!--                                 v-model="formData.email" name="email"-->
                      <!--                                 placeholder="Enter e-mail" required/>-->
                      <!--                        </div>-->

                      <!--                        <div class="form-group fv-row mb-8">-->
                      <!--                          <label class="required fs-6 fw-semibold mb-2" for="status">User status</label>-->
                      <!--                          <v-select class="bg-transparent " label="name" :options="statuses"-->
                      <!--                                    placeholder="Select user status" v-model="status" required>-->
                      <!--                          </v-select>-->
                      <!--                        </div>-->
                      <!--                      </form >-->

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
            <button :disabled="this.isDisabled" @click="checkForm()" type="submit" id="kt_modal_new_target_submit"
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
import {defineComponent} from 'vue'
import {mapActions, mapGetters} from "vuex";
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
    selectedUser: null,
    selectedBalance: null,
    selectedBIN: null,
    selectedProvider: {name: null, isUniversal: null, active: null},
    selectedCard: null,
    errorTitle: null,
    cardType: null,
    systemType: null,
    isDisabled: true,
    activeTab: 'creation',
    styles: {
      // accountReplenishment: "none",
      // change: "none",
      replenishment: "none",
      creation: "block",
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
    formCreationCard: {
      taskId: 0,
      userID: 0,
      cardProvider: null,
      cardType: null,
      balance: null,
      BIN: null,
      note: null,
      limit: null,
      count: null
    },
    formReplenishmentCard: {
      cardID: 0,
      balance: 0,
      amount: 0,
    },
  }
}

export default defineComponent({
  name: "editUserModal",
  data() {
    return initialState();
  },
  watch: {
    "selectedBIN": "selectTypesFromBin",
    "selectedCard": "allCardList"
  },
  methods: {
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
    outputResult() {
      console.log(this.selectedBalance)
      if (this.cardType && this.cardType.name === "Advertising" && this.selectedUser.user.advertisingCommission) {
        this.cardType.commission = this.selectedUser.user.advertisingCommission
      }
      if (this.cardType && this.cardType.name === "Universal" && this.selectedUser.user.universalCommission) {
        this.cardType.commission = this.selectedUser.user.universalCommission
      }
      if (this.cardType && this.cardType.name === "Facebook" && this.selectedUser.user.facebookCommission) {
        this.cardType.commission = this.selectedUser.user.facebookCommission
      }
      if ((this.cardType && this.formCreationCard.count && this.formCreationCard.limit && this.selectedBalance.balance) &&
          this.selectedBalance.balance > (this.cardType.sumToCreate * this.formCreationCard.count) + (this.formCreationCard.limit * this.formCreationCard.count)) {
        this.isDisabled = false
        return true
      }
      if (this.selectedBIN && this.selectedUser.name && this.cardType && this.formCreationCard.count && this.formCreationCard.limit && this.selectedBalance.balance) {
        console.log(this.selectedBIN)
        this.field = false
      }
      this.isDisabled = true
    },
    tabClicked(element) {
      this.selectedUser = null
      this.selectedBalance = null
      this.selectedBIN = null
      this.errorTitle = null
      this.cardType = null
      this.systemType = null
      this.isDisabled = null
      this.formCreationCard.limit = null
      this.formCreationCard.note = null
      this.formCreationCard.count = null
      this.formReplenishmentCard.balance = null
      this.formReplenishmentCard.amount = null
      for (const style in this.styles) {
        this.styles[style] = "none"
        if (style === element) {
          this.activeTab = element
          this.styles[style] = "block"
        }
      }
    },
    openModal() {
      Object.assign(this.$data, initialState());
      this.show = true
      setTimeout(() => {
        this.isShow = this.show
        this.transform = "none"
      }, 200)
    },
    closeModal: function () {
      this.isShow = false
      this.transform = "translate(0, -50px)"
      this.selectedUser = null
      this.selectedBalance = null
      this.selectedBIN = null
      // this.selectedProvider = {name: null, isUniversal: null, active: null}
      this.errorTitle = null
      this.cardType = null
      this.systemType = null
      this.isDisabled = null
      this.activeTab = "creation"
      this.formCreationCard.limit = null
      this.formCreationCard.note = null
      this.formCreationCard.count = null
      setTimeout(() => {
        this.show = false
      }, 200)
    },
    ...mapActions({
      createTaskCardCreation: "users/createTaskCardCreation",
      createTaskCardReplenishment: "users/createTaskCardReplenishment"
    }),
    checkForm: function () {
      if (this.activeTab === "creation") {
        this.formCreationCard.BIN = this.selectedBIN.name.toString()
        this.formCreationCard.cardProvider = this.selectedProvider.name
        this.formCreationCard.userID = this.selectedUser.code
        this.formCreationCard.balance = this.selectedBalance.code
        this.formCreationCard.cardType = this.cardType.name
        console.log(this.formCreationCard)
        // this.createTaskCardCreation(this.formCreationCard)
        //     .then(() => {
        //       Object.assign(this.$data, initialState());
        //       this.closeModal()
        //       this.$swal({
        //         title: "You have successfully created the card",
        //         confirmButtonColor: "#50cd89",
        //         type: "alert message",
        //         icon: "success",
        //       })
        //     })
        //     .catch(error => {
        //       this.errors.responseError = error
        //       if (this.errors.responseError.toString() === "Error: badRequest") this.errorTitle = "Something went wrong. Try again later."
        //       if (this.errors.responseError.toString() === "Error: cardNotFound") this.errorTitle = "Сard not found in the system"
        //       if (this.errors.responseError.toString() === "Error: change is Not Approved") this.errorTitle = "Change is not approved"
        //       if (this.errors.responseError.toString() === "Error: userNotFound") this.errorTitle = "User not found in the system"
        //       if(this.errors.responseError.toString() === "Error: notEnoughBalance") this.errorTitle = "There are not enough funds on the balance to replenish"
        //       if(this.errors.responseError.toString() === "Error: paying not from USDT balance is disabled for now") this.errorTitle = "Paying not from USDT balance is disabled for now"
        //       console.log(this.errorTitle)
        //       this.$swal({
        //         title: this.errorTitle,
        //         type: "alert message",
        //         confirmButtonColor: "#f1416c",
        //         icon: "error",
        //       })
        //     })
      }
      if (this.activeTab === "replenishment") {
        this.formReplenishmentCard.cardID = this.selectedCard.code
        this.formReplenishmentCard.balance = this.selectedBalance.code
        console.log(this.formReplenishmentCard)
        // this.createTaskCardReplenishment(this.formReplenishmentCard)
        //     .then(() => {
        //       Object.assign(this.$data, initialState());
        //       this.closeModal()
        //       this.$swal({
        //         title: "You have successfully changed personal price",
        //         confirmButtonColor: "#50cd89",
        //         type: "alert message",
        //         icon: "success",
        //       })
        //     })
        //     .catch(error => {
        //       console.log(error)
        //       this.errors.responseError = error
        //       if (this.errors.responseError.toString() === "Error: badRequest") this.errorTitle = "Something went wrong. Try again later."
        //       if (this.errors.responseError.toString() === "Error: cardNotFound") this.errorTitle = "Сard not found in the system"
        //       if (this.errors.responseError.toString() === "Error: change is Not Approved") this.errorTitle = "Change is not approved"
        //       if (this.errors.responseError.toString() === "Error: userNotFound") this.errorTitle = "User not found in the system"
        //       if(this.errors.responseError.toString() === "Error: notEnoughBalance") this.errorTitle = "There are not enough funds on the balance to replenish"
        //       if(this.errors.responseError.toString() === "Error: paying not from USDT balance is disabled for now") this.errorTitle = "Paying not from USDT balance is disabled for now"
        //       console.log(this.errorTitle)
        //       this.$swal({
        //         title: this.errorTitle,
        //         type: "alert message",
        //         confirmButtonColor: "#f1416c",
        //         icon: "error",
        //       })
        //     });
      }
    },
    userBalance(balance) {
      if (balance === 0) {
        return '0.00'
      }

      let result = []
      let arrayBalance = ("" + balance).split("").map(Number)
      for (const i of arrayBalance) {
        result.push(i)
      }

      function roundingNum(x, n) {
        return parseFloat(Number.parseFloat(x).toFixed(n || 2));
      }

      let arrRoundingNum = Array.from(roundingNum(balance, 2).toString())
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
    getData() {
      if (this.loggedIn && this.suppliers.length === 0) {
        this.$store.dispatch("users/getSuppliersMethod")
      }
      return true
    }
  },
  computed: {
    ...mapGetters({
      cards: "users/getCards",
      cardsLoaded: "users/getCardsLoaded",
      suppliers: "users/getSuppliers",
      loggedIn: "auth/isLoggedIn",
    }),
    allCardList() {
      let cardList = []
      this.cards.forEach(card => {
        console.log(card)
        if(this.selectedUser && this.selectedUser.code === card.userId) {
          let cardForList = {code: card.ID, name: "ID: " + card.ID + " - " + card.cardNumber, card: card}
          cardList.push(cardForList)
        }
      })
      return cardList
    },
    BINsList() {
      let BINsList = []
      let code = {
        "VISA": "4",
        "MasterCard": "5"
      }
      let provider = this.selectedProvider.provider
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
          console.log(legitBin)
          BINsList.push(legitBin)
        })
      }
      return BINsList
    },
    balanceList() {
      let balances = []
      if (this.selectedUser) {
        let user = this.selectedUser.user
        balances = [
          {"code": "USDT", "balance": user.USDTBalance, "name": "USDT - " + this.userBalance(user.USDTBalance) + " ₮"},
          {"code": "BTC", "balance": user.BTCBalance, "name": "BTC - " + this.userBalance(user.BTCBalance) + " ₿"},
          {"code": "EUR", "balance": user.EURBalance, "name": "EUR - " + this.userBalance(user.EURBalance) + " €"},
          {"code": "USD", "balance": user.USDBalance, "name": "USD - " + this.userBalance(user.USDBalance) + " $"},
        ]
      }
      return balances
    },
    providersList() {
      let providersList = []
      if (this.suppliers.length === 0) {
        this.$store.dispatch("users/getSuppliersMethod")
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
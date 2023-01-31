<template>
  <!-- из шаблона -->
  <div v-if="show">
    <div :class="{show: this.isShow}" class="modal fade d-block" @click.self="closeModal" id="kt_modal_view_users"
         tabindex="-1"
         aria-hidden="true">
      <div class="modal-dialog modal-dialog-centered mw-1000px" :style="{transform: this.transform}">
        <!--begin::Modal content-->
        <div class="modal-content">
          <!--begin::Modal header-->
          <div class="modal-header">
            <h2>Card transaction log</h2>
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
          <div class="modal-buttons modal-header border-bottom-0">
            <button class="btn btn-primary">Export</button>
            <button class="btn btn-dark" @click="closeModal">Exit</button>
          </div>
          <!--begin::Modal header-->
          <!--begin::Modal body-->
          <div class="modal-body scroll">
            <div class="mb-15">
              <!--begin::List-->
              <div class="mh-375px scroll me-n7 pe-7">
                <div class="card-body p-0">
                  <!--begin::Table container-->
                  <div class="table-responsive">
                    <!--begin::Table-->
                    <table class="table table-row-bordered table-row-gray-100 align-middle gs-0 gy-3">
                      <!--begin::Table head-->
                      <thead class="rounded-left">
                      <tr class="fw-bold text-muted bg-light">
                        <th class="min-w-150px">Account</th>
                        <th @click="filterTransactionBalance('amount')" class="min-w-150px cursor-pointer text-hover-primary">Amount</th>
                        <th class="min-w-150px">Settled</th>
                        <th @click="filterTransactionDate('TransactionDate')" class="min-w-150px cursor-pointer text-hover-primary">Date</th>
                        <th class="min-w-150px">Merchant</th>
                        <th class="min-w-150px">Country</th>
                      </tr>
                      </thead>
                      <!--end::Table head-->
                      <!--begin::Table body-->
                      <tbody>
                      <tr v-for="transaction in this.filteredTransactions" :key="transaction">
                        <td class="fs-6 align-middle">
                          <div class="d-flex flex-column">
                            <span class="text-dark fw-bold mb-1 fs-6">ID: {{ transaction.cardId }}</span>
                            <span
                                class="text-muted fw-semibold text-muted d-block fs-7">{{
                                this.number.toString().substring(0, 4) + " " + this.number.toString().substring(4, 8) + " " + this.number.toString().substring(8, 12) + " " + this.number.toString().substring(12, 16)
                              }}</span>
                          </div>
                        </td>
                        <td class="align-middle">
                          <div v-if="transaction.increase" class="badge badge-light-success mb-2"
                               style="width: max-content;">+ {{ transaction.amount }}
                          </div>
                          <div v-else class="badge badge-light-danger"
                               style="width: max-content;">- {{ transaction.amount }}
                          </div>
                        </td>
                        <td class="align-middle">
                          <div class="d-flex" v-if="transaction.settled">
                  <span class="badge badge-light-success me-2">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor"
                         class="bi bi-check" viewBox="0 0 16 16">
                      <path
                          d="M10.97 4.97a.75.75 0 0 1 1.07 1.05l-3.99 4.99a.75.75 0 0 1-1.08.02L4.324 8.384a.75.75 0 1 1 1.06-1.06l2.094 2.093 3.473-4.425a.267.267 0 0 1 .02-.022z"/>
                    </svg>
                    Settled
                  </span>
                          </div>
                          <div class="d-flex align-items-center" v-else>
                  <span class="badge badge-light-warning me-2">
                      <span class="svg-icon svg-icon-warning svg-icon-3">
                        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                          <path opacity="0.3"
                                d="M20.9 12.9C20.3 12.9 19.9 12.5 19.9 11.9C19.9 11.3 20.3 10.9 20.9 10.9H21.8C21.3 6.2 17.6 2.4 12.9 2V2.9C12.9 3.5 12.5 3.9 11.9 3.9C11.3 3.9 10.9 3.5 10.9 2.9V2C6.19999 2.5 2.4 6.2 2 10.9H2.89999C3.49999 10.9 3.89999 11.3 3.89999 11.9C3.89999 12.5 3.49999 12.9 2.89999 12.9H2C2.5 17.6 6.19999 21.4 10.9 21.8V20.9C10.9 20.3 11.3 19.9 11.9 19.9C12.5 19.9 12.9 20.3 12.9 20.9V21.8C17.6 21.3 21.4 17.6 21.8 12.9H20.9Z"
                                fill="currentColor"/>
                          <path
                              d="M16.9 10.9H13.6C13.4 10.6 13.2 10.4 12.9 10.2V5.90002C12.9 5.30002 12.5 4.90002 11.9 4.90002C11.3 4.90002 10.9 5.30002 10.9 5.90002V10.2C10.6 10.4 10.4 10.6 10.2 10.9H9.89999C9.29999 10.9 8.89999 11.3 8.89999 11.9C8.89999 12.5 9.29999 12.9 9.89999 12.9H10.2C10.4 13.2 10.6 13.4 10.9 13.6V13.9C10.9 14.5 11.3 14.9 11.9 14.9C12.5 14.9 12.9 14.5 12.9 13.9V13.6C13.2 13.4 13.4 13.2 13.6 12.9H16.9C17.5 12.9 17.9 12.5 17.9 11.9C17.9 11.3 17.5 10.9 16.9 10.9Z"
                              fill="currentColor"/>
                      </svg>
                      </span>
                    <!--end::Svg Icon-->
                    Pending
                  </span>
                          </div>
                        </td>
                        <td class="align-middle">
                          <span class="text-dark fw-bold fs-6">{{ splitDate(transaction.TransactionDate) }}</span>
                        </td>
                        <td class="fs-6 align-middle">
                          <Popper class="fw-semibold fs-7" arrow hover placement="top" :content="transaction.merchant">
                            <span class="fw-bold d-block mb-1 fs-6 text-dark fw-semibold">{{
                                outputMerchant(transaction.merchant)
                              }}</span>
                          </Popper>
                        </td>
                        <td class="align-middle">
                          <span class="text-dark fw-bold mb-1 fs-6">{{ transaction.country }}</span>
                        </td>
                      </tr>
                      </tbody>
                      <!--end::Table body-->
                    </table>
                    <!--end::Table-->
                  </div>
                  <!--end::Table container-->
                </div>
              </div>
              <!--end::List-->
            </div>
            <!--end::Users-->
          </div>
          <!--end::Modal body-->
        </div>
        <!--end::Modal content-->
      </div>
    </div>

  </div>
</template>

<script>

import timezones from "@/timezones.json";
import {mapGetters} from "vuex";

export default {
  name: "transactionLogModal",
  data() {
    return {
      timezones: timezones,
      show: false,
      isShow: false,
      transform: "translate(0, -50px)",
      date: null,
      number: null,
      cardID: null,
      sum: null,
      balanceName: null,
      initial: null,
      total: null,
      link: null,
      filteredTransactions: this.transactions,
      filterItems: {
        amount: 0,
        TransactionDate: 1,
        balances: [
          {
            code: 0,
            name: "USDT",
          },
          {
            code: 0,
            name: "USD",
          },
          {
            code: 0,
            name: "BTC",
          },
          {
            code: 0,
            name: "EUR",
          }
        ],
        balance: {
          name: null,
          code: null,
        },
        providers: [
          {
            name: "AdMediaCards",
          },
          {
            name: "Grats.Cards",
          }
        ],
        provider: {
          name: null
        },
        search: '',
      },
    }
  },
  created() {
    this.filteredTransactions = this.transactions
    if (this.transactions) {
      this.filter()
    }
  },
  watch: {
    transactions: "filter"
  },
  computed: {
    ...mapGetters({
      currentUser: "auth/currentUser"
    }),
  },
  methods: {
    filterTransactionDate(date) {
      this.filterItems[date] !== 0 ? this.filterItems[date] = -this.filterItems[date] : this.filterItems[date] = 1
      this.filterItems.amount = 0
      // date === "TransactionDate" ? this.filterItems.DeletedAt = 0 : this.filterItems.TransactionDate = 0
      this.filter()
    },
    filterTransactionBalance(balance) {
      if (!balance) {
        return "-"
      }
      let stat = 0
      this.filterItems[balance] !== 0 ? stat = -this.filterItems[balance] : stat = 1
      this.filterItems.amount = 0
      this.filterItems[balance] = stat
      this.filter()
    },
    filter() {
      this.filteredTransactions = this.transactions.filter(transaction => {
        let searchString = this.filterItems.search.toLowerCase()
        if ( (transaction.transaction && transaction.notes) && (transaction.transaction.toLowerCase().indexOf(searchString) !== -1
            || transaction.notes.toLowerCase().indexOf(searchString) !== -1)) {
          return transaction
        }
      })
      if (this.filterItems.search && this.filteredTransactions.length > 1) {
        this.background = "rgb(230, 255, 241)"
      } else {
        this.background = "none"
      }
      // фильтрация по дате
      if (this.filterItems.TransactionDate !== 0) {
        this.filterItems.TransactionDate === 1 ? this.filteredTransactions.sort(function (a, b) {
          return new Date(b['TransactionDate']) - new Date(a['TransactionDate'])
        }) : this.filteredTransactions.sort(function (a, b) {
          return new Date(a['TransactionDate']) - new Date(b['TransactionDate'])
        })
      }

      // фильтрация по балансам
      if (this.filterItems.amount !== 0) {
        this.filterItems.amount === 1 ? this.filteredTransactions.sort(function (a, b) {
          return new Date(b['amount']) - new Date(a['amount'])
        }) : this.filteredTransactions.sort(function (a, b) {
          return new Date(a['amount']) - new Date(b['amount'])
        })
      }
    },
    outputMerchant(merchant) {
      if (merchant) {
        return merchant.split(" ")[0] + merchant.split(" ")[1]
      }
      return "-"
    },
    openModal(card, transaction) {
      if (transaction.length !== 0) {
        this.show = true
        this.transactions = Object.values(transaction);
        this.filteredTransactions = this.transactions
        // this.date = transaction.TransactionDate
        // this.increase = transaction.increase
        this.number = card.cardNumber
        setTimeout(() => {
          this.isShow = this.show
          this.transform = "none"
        }, 200)
      }
      if (transaction.length === 0) {
        this.show = false
        this.$swal({
          title: "No data/No transactions yet",
          type: "alert message",
          confirmButtonClass: "btn-info",
          confirmButtonText: "OK",
          icon: "info"
        })
      }
      // this.cardID = card.ID
      // this.balanceName = transaction.balance
      // this.sum = card.balance
      // this.link = transaction.transaction
      // this.initial = transaction.initial
      // this.total = transaction.total
      // this.amount = transaction.amount
    },
    closeModal: function () {
      this.isShow = false
      this.transform = "translate(0, -50px)"
      setTimeout(() => {
        this.show = false
      }, 200)
    },
    splitDate(date) {
      if (date === "0001-01-01T00:00:00Z" || !date) {
        return "-"
      }
      let userTimeZone = this.currentUser.timeZone
      let timeZone = this.timezones.find(item => item.text === userTimeZone);
      if(timeZone && timeZone.utc.includes('-')) {
        let newDate = new Date(date);
        let utc = Number(timeZone.utc.split('-')[1])
        newDate.setTime(newDate.getTime() - (utc * 60 * 60 * 1000))
        return newDate.toLocaleString().split(",")[0] + " " + newDate.toLocaleString().split(",")[1]
      }
      if(timeZone && !timeZone.utc.includes('-')) {
        let newDate = new Date(date);
        let utc = Number(timeZone.utc)
        newDate.setTime(newDate.getTime() + (utc * 60 * 60 * 1000))
        return newDate.toLocaleString().split(",")[0] + " " + newDate.toLocaleString().split(",")[1]
      }
    },
    transactionBalance(number) {
      if (number === 0) {
        return '0000.00'
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
  },
}
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
<template>
  <div v-show="isHidden" @click="isHidden = false" class="position-absolute vh-100 w-100"
       style="z-index: 99;left: 0px;top: 0px;"></div>
  <div v-show="actionsShow" @click="actionsShow = false" class="position-absolute vh-100 w-100"
       style="z-index: 99;left: 0px;top: 0px;"></div>
  <div class="d-flex flex-column flex-column-fluid container-fluid">
    <div class="toolbar mb-5 mb-lg-7" id="kt_toolbar">
      <!--begin::Page title-->
      <div class="page-title d-flex flex-column me-3">
        <!--begin::Title-->
        <h1 class="d-flex text-dark fw-bold my-1 fs-3">{{ $t('transactions.title') }}</h1>
        <!--end::Title-->
        <!--begin::Breadcrumb-->
        <ul class="breadcrumb breadcrumb-dot fw-semibold text-gray-600 fs-7 my-1">
          <!--begin::Item-->
          <li class="breadcrumb-item text-gray-600">
            <router-link to="/" class="text-gray-600 text-hover-primary">{{ $t('common.crumb1') }}</router-link>
          </li>
          <!--end::Item-->
          <!--begin::Item-->
          <li class="breadcrumb-item text-gray-500">{{ $t('transactions.title') }}</li>
          <!--end::Item-->
        </ul>
        <!--end::Breadcrumb-->
      </div>
      <!--end::Page title-->
      <!--begin::Actions-->
      <div class="actions-group-buttons d-flex align-items-center py-2 py-md-1 position-relative">
        <!--begin::Wrapper-->
        <div class="me-3">
          <!--begin::Menu-->
          <button @click='isHidden = !isHidden' class="btn btn-light fw-bold">
            <!--begin::Svg Icon | path: icons/duotune/general/gen031.svg-->
            <span class="svg-icon svg-icon-5 svg-icon-gray-500 me-1">
                            <svg width="24" height="24" viewBox="0 0 24 24" fill="none"
                                 xmlns="http://www.w3.org/2000/svg">
                                <path
                                    d="M19.0759 3H4.72777C3.95892 3 3.47768 3.83148 3.86067 4.49814L8.56967 12.6949C9.17923 13.7559 9.5 14.9582 9.5 16.1819V19.5072C9.5 20.2189 10.2223 20.7028 10.8805 20.432L13.8805 19.1977C14.2553 19.0435 14.5 18.6783 14.5 18.273V13.8372C14.5 12.8089 14.8171 11.8056 15.408 10.964L19.8943 4.57465C20.3596 3.912 19.8856 3 19.0759 3Z"
                                    fill="currentColor"/>
                            </svg>
                        </span>
            <!--end::Svg Icon-->
            <span class="user-filter-span">{{ $t('common.filter') }}</span>
          </button>
          <!--begin::Menu 1-->
          <div id="filter-modal" v-show="isHidden"
               class="menu menu-sub menu-sub-dropdown show position-fixed float-left" data-kt-menu="true"
               style="width: inherit; z-index: 999; right: 40px">
            <!--begin::Header-->
            <div class="px-7 py-5">
              <div class="fs-5 text-dark fw-bold">{{ $t('transactions.filterTransactions') }}</div>
            </div>
            <!--end::Header-->
            <!--begin::Menu separator-->
            <div class="separator border-gray-200"></div>
            <!--end::Menu separator-->
            <!--begin::Form-->
            <div class="px-7 py-5">
              <div class="mb-10">
                <div class="form-group fv-row">
                  <label class="form-label fw-semibold" for="status">{{ $t('cards.tableTitle') }}:</label>
                  <v-select class="bg-transparent" label="name" :options="allCardList"
                            placeholder="Select client" v-model="filterItems.selectedCard" required>
                  </v-select>
                </div>
              </div>

              <div class="mb-10">
                <div class="form-group fv-row">
                  <label class="form-label fw-semibold" for="role">{{ $t('common.status') }}:</label>
                  <v-select class="bg-transparent" label="name" :options="filterItems.statuses"
                            v-model="filterItems.status" required>
                  </v-select>
                </div>
              </div>

              <div class="mb-10">
                <div class="form-group fv-row">
                  <label class="form-label fw-semibold" for="search">{{ $t('common.search') }}</label>
                  <input class="form-control form-control-lg form-control-solid" label="name"
                         v-model="filterItems.search">
                </div>
              </div>


              <div class="d-flex">
                <button @click="resetFilterTransactions" type="reset"
                        class="btn btn-sm btn-light btn-active-light-primary me-2">{{ $t('common.reset') }}
                </button>
                <button @click="filter" type="submit" class="btn btn-sm btn-primary"
                        data-kt-menu-dismiss="true">{{ $t('common.search') }}
                </button>
              </div>
              <!--end::Actions-->
            </div>
            <!--end::Form-->
          </div>
          <!--end::Menu 1-->
          <!--end::Menu-->
        </div>
        <!--        <button class="btn btn-dark fw-bold" data-bs-toggle="modal"-->
        <!--                data-bs-target="#kt_modal_create_app" id="kt_toolbar_primary_button">Пока хз-->
        <!--        </button>-->
        <!--end::Wrapper-->
      </div>
      <!--end::Actions-->
    </div>
    <div class="card mb-5 mb-xl-8">
      <!--begin::Header-->
      <div class="card-header align-items-center border-0 pt-5">
        <h3 class="card-title align-items-start flex-column">
          <span class="card-label fw-bold fs-3 mb-1">{{ $t('transactions.tableTitle') }}</span>
          <span class="text-muted mt-1 fw-semibold fs-7">{{ $t('transactions.amount') }}: {{
              this.transactions.length
            }}</span>
        </h3>

        <div class="d-flex align-items-center">
          <button @click="exportExcel('xls')"
                           class="btn show-delete btn-icon btn-bg-light btn-active-color-primary btn-sm">
            <!--begin::Svg Icon | path: C:/wamp64/www/keenthemes/core/html/src/media/icons/duotune/files/fil017.svg--><span
              class="svg-icon svg-icon-muted svg-icon-2hx"><svg width="24" height="24" viewBox="0 0 24 24" fill="none"
                                                                xmlns="http://www.w3.org/2000/svg"><path opacity="0.3"
                                                                                                         d="M10 4H21C21.6 4 22 4.4 22 5V7H10V4Z"
                                                                                                         fill="currentColor"></path><path
              opacity="0.3" d="M13 14.4V9C13 8.4 12.6 8 12 8C11.4 8 11 8.4 11 9V14.4H13Z" fill="currentColor"></path><path
              d="M10.4 3.60001L12 6H21C21.6 6 22 6.4 22 7V19C22 19.6 21.6 20 21 20H3C2.4 20 2 19.6 2 19V4C2 3.4 2.4 3 3 3H9.20001C9.70001 3 10.2 3.20001 10.4 3.60001ZM13 14.4V9C13 8.4 12.6 8 12 8C11.4 8 11 8.4 11 9V14.4H8L11.3 17.7C11.7 18.1 12.3 18.1 12.7 17.7L16 14.4H13Z"
              fill="currentColor"></path></svg></span><!--end::Svg Icon--><!--            {{ $t('cards.unloading') }}-->
          </button>
        </div>
      </div>
      <!--end::Header-->
      <!--begin::Body-->
      <div class="card-body py-3">
        <!--begin::Table container-->
        <div class="table-responsive">
          <!--begin::Table-->
          <table v-if="getTransactions()" class="table table-row-bordered table-row-gray-100 align-middle gs-0 gy-3">
            <!--begin::Table head-->
            <thead class="rounded-left">
            <tr class="fw-bold text-muted bg-light">
              <th class="min-w-170px cursor-pointer text-hover-primary">
                ID
              </th>
              <th class="min-w-130px">{{ $t('common.status') }}</th>
              <th @click="filterTransactionDate('TransactionDate')"
                  class="min-w-180px cursor-pointer text-hover-primary">{{ $t('transactions.date') }}
              </th>
              <th @click="filterTransactionBalance('total')" class="min-w-180px cursor-pointer text-hover-primary">
                {{ $t('transactions.movement') }}
              </th>
              <th class="min-w-150px">{{ $t('common.link') }}</th>
            </tr>
            </thead>
            <!--end::Table head-->
            <!--begin::Table body-->
            <tbody>
            <tr v-for="transaction in filteredTransactions" :key="transaction">
              <td class="align-middle">
                <span class="text-dark fw-bold d-block fs-6">
                  {{ transaction.ID }}
                </span>
              </td>
              <td class="align-middle">
                <div class="d-flex" v-if="transaction.settled = true">
                  <span class="badge badge-light-success">
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
              <td class="text-muted fw-semibold fs-6 align-middle">
                <span v-if="transaction.TransactionDate !== '0001-01-01T00:00:00Z'" class="d-block fs-6">
                  {{ splitDate(transaction.TransactionDate) }}
                </span>
                <span v-else class="d-block fs-6">
                  {{ splitDate(transaction.CreatedAt) }}
                </span>
              </td>
              <td class="fs-6 align-middle cursor-pointer">
                <div v-if="transaction.increase" class="d-flex flex-column">
                  <span class="text-dark fw-bold d-block fs-6 mb-1">
                    {{ transaction.balance }}
                  </span>
                  <div class="badge badge-light-success"
                       style="width: max-content;">+ {{ transactionBalance(transaction.amount) }}
                  </div>
                </div>
                <div v-else class="d-flex flex-column">
                  <span class="text-dark fw-bold d-block fs-6 mb-1 d-flex align-items-center">
                    <span v-if="transaction.balance">{{ transaction.balance }}</span>
                    <span v-else>USD</span>
                    <span class="ml-2 badge badge-light-danger"
                          style="width: max-content;">- {{ transactionBalance(transaction.amount) }}</span>
                  </span>
                  <div>
                    <span class="text-dark fw-bold d-block fs-6 d-flex align-items-center">
                      CardID: {{ transaction.cardId }}
                               <span class="ml-2 badge badge-light-success"
                                     style="width: max-content;">+ {{ transactionBalance(transaction.amount) }}
                    </span>
                    </span>

                  </div>
                </div>
              </td>

              <td class="align-middle">
                <span class="fw-bold d-block mb-1 fs-6 text-muted fw-semibold">{{ transaction.transaction }}</span>
              </td>
            </tr>
            </tbody>
            <!--end::Table body-->
          </table>
          <!--end::Table-->
          <table ref="exportable_table" style="display:none;" v-if="getTransactions()" class="table table-row-bordered table-row-gray-100 align-middle gs-0 gy-3">
            <!--begin::Table head-->
            <thead class="rounded-left">
            <tr class="fw-bold text-muted bg-light">
              <th class="min-w-170px cursor-pointer text-hover-primary">
                ID
              </th>
              <th class="min-w-130px">{{ $t('common.status') }}</th>
              <th class="min-w-180px cursor-pointer text-hover-primary">
                {{ $t('transactions.date') }}
              </th>
              <th class="min-w-180px cursor-pointer text-hover-primary">
                Total
              </th>
              <th class="min-w-180px cursor-pointer text-hover-primary">
                Commission
              </th>
              <th class="min-w-180px cursor-pointer text-hover-primary">
                Country
              </th>
              <th class="min-w-180px cursor-pointer text-hover-primary">
                Initial
              </th>
              <th class="min-w-180px cursor-pointer text-hover-primary">
                Merchant
              </th>
              <th class="min-w-180px cursor-pointer text-hover-primary">
                Notes
              </th>
              <th class="min-w-150px">{{ $t('common.link') }}</th>
            </tr>
            </thead>
            <!--end::Table head-->
            <!--begin::Table body-->
            <tbody>
            <tr v-for="transaction in filteredTransactions" :key="transaction">
              <td class="align-middle">
                <span class="text-dark fw-bold d-block fs-6">
                  {{ transaction.ID }}
                </span>
              </td>
              <td class="align-middle">
                <div class="d-flex" v-if="transaction.settled = true">
                  <span class="badge badge-light-success">
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
              <td class="text-muted fw-semibold fs-6 align-middle">
                <span v-if="transaction.TransactionDate !== '0001-01-01T00:00:00Z'" class="d-block fs-6">
                  {{ splitDate(transaction.TransactionDate) }}
                </span>
                <span v-else class="d-block fs-6">
                  {{ splitDate(transaction.CreatedAt) }}
                </span>
              </td>
              <td class="fs-6 align-middle cursor-pointer">
                {{transaction.total}}
              </td>
              <td class="fs-6 align-middle cursor-pointer">
                {{transaction.commission}}
              </td>
              <td class="fs-6 align-middle cursor-pointer">
                {{transaction.country}}
              </td>
              <td class="fs-6 align-middle cursor-pointer">
                {{transaction.initial}}
              </td>
              <td class="fs-6 align-middle cursor-pointer">
                {{transaction.merchant}}
              </td>
              <td class="fs-6 align-middle cursor-pointer">
                {{transaction.notes}}
              </td>
              <td class="align-middle">
                {{ transaction.transaction }}
              </td>
            </tr>
            </tbody>
            <!--end::Table body-->
          </table>
        </div>
        <!--end::Table container-->
      </div>
      <!--begin::Body-->
    </div>
  </div>
</template>

<script>
import {mapGetters, mapActions} from 'vuex'
import timezones from "@/timezones.json";
import xlsx from "xlsx"
export default {
  name: 'tableTransactions',
  data() {
    return {
      json_fields: {
        CreatedAt: "CreatedAt",
        DeletedAt: "DeletedAt",
        TransactionDate: "TransactionDate",
        UpdatedAt: "UpdatedAt",
        amount: "amount",
        balance: "balance",
        cardId: "cardId",
        commission: "commission",
        country: "country",
        increase: "increase",
        initial: "initial",
        managerId: "managerId",
        merchant: "merchant",
        notes: "notes",
        provider: "provider",
        providerId: "providerId",
        settled: "settled",
        status: "status",
        total: "total",
        transaction: "transaction",
        userId: "userId"
      },
      json_meta: [
        [
          {
            key: "charset",
            value: "utf-8",
          },
        ],
      ],
      timezones: timezones,
      actionsShow: false,
      isHidden: false,
      filteredTransactions: this.transactions,
      filterItems: {
        selectedCard: {code: 0, name: "", card: {}},
        TransactionDate: 1,
        total: 0,
        statuses: [
          {
            name: "Rejected",
            code: false
          },
          {
            name: "Settled",
            code: true
          },
        ],
        status: {
          name: null,
          code: null
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
  methods: {
    ...mapActions({
      login: "auth/login"
    }),
    exportExcel(type, fn, dl) {
      let elt = this.$refs.exportable_table;
      let wb = xlsx.utils.table_to_book(elt, {sheet:"Sheet JS"});
      return dl ?
          xlsx.write(wb, {bookType:type, bookSST:true, type: 'base64'}) :
          xlsx.writeFile(wb, fn || (("transactions" + '.'|| 'SheetJSTableExport.') + (type || 'xls')));
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
    filterTransactionDate(date) {
      this.filterItems[date] !== 0 ? this.filterItems[date] = -this.filterItems[date] : this.filterItems[date] = 1
      this.filter()
    },
    filterTransactionBalance(balance) {
      let stat = 0
      this.filterItems[balance] !== 0 ? stat = -this.filterItems[balance] : stat = 1
      this.filterItems.total = 0
      this.filterItems[balance] = stat
      this.filter()
    },
    filter() {
      this.filteredTransactions = this.transactions.filter(transaction => {
        console.log(transaction)
        let searchString = this.filterItems.search.toLowerCase()
        if ((!this.filterItems.status.name || transaction.settled === this.filterItems.status.code)
            && (!this.filterItems.selectedCard.code || transaction.cardId === this.filterItems.selectedCard.code)
            && (transaction.ID.toString().toLowerCase().indexOf(searchString) !== -1 || transaction.userId.toString().indexOf(searchString)) !== -1) {
          return transaction
        }
      })
      // фильтрация по дате
      if (this.filterItems.TransactionDate !== 0) {
        this.filterItems.TransactionDate === 1 ? this.filteredTransactions.sort(function (a, b) {
          return new Date(b['TransactionDate']) - new Date(a['TransactionDate'])
        }) : this.filteredTransactions.sort(function (a, b) {
          return new Date(a['TransactionDate']) - new Date(b['TransactionDate'])
        })
      }

      // фильтрация по балансам
      if (this.filterItems.total !== 0) {
        this.filterItems.total === 1 ? this.filteredTransactions.sort(function (a, b) {
          return new Date(b['total']) - new Date(a['total'])
        }) : this.filteredTransactions.sort(function (a, b) {
          return new Date(a['total']) - new Date(b['total'])
        })
      }
    },
    resetFilterTransactions() {
      this.filterItems = {
        selectedCard: {code: 0, name: "", card: {}},
        TransactionDate: 1,
        total: 0,
        statuses: [
          {
            name: "Settled",
            code: true
          },
          {
            name: "Pending",
            code: false
          },
        ],
        status: {
          name: null,
          code: null
        },
        search: '',
      }
      this.filter()
    },
    getTransactions() {
      if (this.loggedIn && !this.transactionsLoaded) {
        this.$store.dispatch("users/getAllTransactionsMethod")
      }
      if (this.loggedIn && !this.cardsLoaded) {
        this.$store.dispatch("users/getCardsMethod")
      }
      this.req = false
      return true
    },
    splitDate(date) {
      if (date === "0001-01-01T00:00:00Z" || !date) {
        return "-"
      }
      let userTimeZone = this.user.timeZone
      let timeZone = this.timezones.find(item => item.text === userTimeZone);
      if (timeZone && timeZone.utc.includes('-')) {
        let newDate = new Date(date);
        let utc = Number(timeZone.utc.split('-')[1])
        newDate.setTime(newDate.getTime() - (utc * 60 * 60 * 1000))
        return newDate.toLocaleString().split(",")[0] + " " + newDate.toLocaleString().split(",")[1]
      }
      if (timeZone && !timeZone.utc.includes('-')) {
        let newDate = new Date(date);
        let utc = Number(timeZone.utc)
        newDate.setTime(newDate.getTime() + (utc * 60 * 60 * 1000))
        return newDate.toLocaleString().split(",")[0] + " " + newDate.toLocaleString().split(",")[1]
      }
    },
  },
  computed: {
    ...mapGetters({
      cards: "users/getCards",
      cardsLoaded: "users/getCardsLoaded",
      transactions: "users/getAllTransactionsMethod",
      transactionsLoaded: "users/getAllTransactionsLoaded",
      loggedIn: "auth/isLoggedIn",
      getTokenHeader: "auth/getTokenHeader",
      user: "auth/currentUser"
    }),
    allCardList() {
      let cardList = []
      this.cards.forEach(card => {
        let cardForList = {code: card.ID, name: "ID: " + card.ID + " - " + card.cardNumber, card: card}
        cardList.push(cardForList)
      })
      return cardList
    },
  },
}
</script>

<style scoped>
.table.gs-0 th:first-child,
.table.gs-0 td:first-child {
  padding-left: 1rem !important;
}

.show-delete {
  transition: opacity 0.2s ease !important;
}
</style>

<style>
.min-w-180px {
  min-width: 180px !important
}

.vs--searchable .vs__dropdown-toggle {
  background-color: var(--kt-input-solid-bg);
  border-color: var(--kt-input-solid-bg);
  color: var(--kt-input-solid-color);
  min-height: calc(1.5em + 1.65rem + 2px);
  padding: 0.825rem 1.5rem;
  font-size: 1.15rem;
  border-radius: 0.625rem;
}

.vs--open .vs__dropdown-menu {
  border-color: var(--kt-input-solid-bg-focus) !important;
}

.vs__dropdown-toggle {

}

.vs__selected, .vs__actions, .vs__search, .vs__search:focus, .vs__selected-options, .vs__actions {
  border: unset !important;
  padding: 0 !important;
  margin: 0 !important;
}

.vs__search {
  color: var(--kt-text-gray-500) !important;
}

.vs__search, .vs__search:focus, .vs__selected {
  font-size: 1.15rem !important;
  font-weight: 500;
}

.vs__selected {
  color: var(--kt-input-solid-color) !important;
}

.vs__clear {
  display: flex;
}
</style>
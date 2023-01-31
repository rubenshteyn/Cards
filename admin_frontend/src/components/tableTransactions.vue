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
        <h1 class="d-flex text-dark fw-bold my-1 fs-3">Transaction management</h1>
        <!--end::Title-->
        <!--begin::Breadcrumb-->
        <ul class="breadcrumb breadcrumb-dot fw-semibold text-gray-600 fs-7 my-1">
          <!--begin::Item-->
          <li class="breadcrumb-item text-gray-600">
            <router-link to="/" class="text-gray-600 text-hover-primary">Home</router-link>
          </li>
          <!--end::Item-->
          <!--begin::Item-->
          <li class="breadcrumb-item text-gray-500">Transaction management</li>
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
          <button @click='usersListModel' class="btn btn-light fw-bold">
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
            <span class="user-filter-span">Filter</span>
          </button>
          <!--begin::Menu 1-->
          <div id="filter-modal" v-show="isHidden"
               class="menu menu-sub menu-sub-dropdown show position-fixed float-left w-md-500px transaction-filter-modal"
               data-kt-menu="true"
               style="width: inherit; z-index: 999; right: 40px; width: 300px">
            <!--begin::Header-->
            <div class="px-7 py-5">
              <div class="fs-5 text-dark fw-bold">Transaction filter</div>
            </div>
            <!--end::Header-->
            <!--begin::Menu separator-->
            <div class="separator border-gray-200"></div>
            <!--end::Menu separator-->
            <!--begin::Form-->
            <div class="px-7 py-5">
              <div class="mb-10">
                <div class="form-group fv-row">
                  <label class="form-label fw-semibold" for="status">Client:</label>
                  <v-select class="bg-transparent" label="name" :options="usersList"
                            v-model="filterItems.selectedUser" required>
                  </v-select>
                </div>
              </div>

              <div class="mb-10">
                <div class="form-group fv-row">
                  <label class="form-label fw-semibold" for="role">Account:</label>
                  <v-select class="bg-transparent" label="name" :options="filterItems.balances"
                            v-model="filterItems.balance" required>
                  </v-select>
                </div>
              </div>

              <div class="mb-10">
                <div class="form-group fv-row">
                  <label class="form-label fw-semibold" for="role">Providers:</label>
                  <v-select class="bg-transparent" label="name" :options="filterItems.providers"
                            v-model="filterItems.provider" required>
                  </v-select>
                </div>
              </div>

              <div class="mb-10">
                <div class="form-group fv-row">
                  <label class="form-label fw-semibold" for="search">Search</label>
                  <input class="form-control form-control-lg form-control-solid" label="name"
                         v-model="filterItems.search">
                </div>
              </div>


              <div class="d-flex">
                <button @click="resetFilterTransactions" type="reset"
                        class="btn btn-sm btn-light btn-active-light-primary me-2">Reset
                </button>
                <button @click="filter" type="submit" class="btn btn-sm btn-primary"
                        data-kt-menu-dismiss="true">Apply
                </button>
              </div>
              <!--end::Actions-->
            </div>
            <!--end::Form-->
          </div>
          <!--end::Menu 1-->
          <!--end::Menu-->
        </div>
        <!--end::Wrapper-->

        <button @click='showModal' class="btn btn-dark fw-bold">
          Unloading
        </button>
      </div>
      <!--end::Actions-->
    </div>
    <div class="card mb-5 mb-xl-8">
      <!--begin::Header-->
      <div class="card-header align-items-center border-0 pt-5">
        <h3 class="card-title align-items-start flex-column">
          <span class="card-label fw-bold fs-3 mb-1">Transaction</span>
          <span class="text-muted mt-1 fw-semibold fs-7">Amount transactions: {{
              this.filteredTransactions.length
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
              <th @click="filterTransactionDate('TransactionDate')"
                  class="min-w-170px cursor-pointer text-hover-primary">
                Date of creation
              </th>
              <th class="min-w-150px">Client</th>
              <th class="min-w-150px">Provider</th>
              <th class="min-w-150px">Settled</th>
              <th class="min-w-150px">
                Account
              </th>
              <th @click="filterTransactionBalance('total')" class="min-w-150px cursor-pointer text-hover-primary">
                Amount
              </th>
              <th class="min-w-150px cursor-pointer text-hover-primary">
                Merchant
              </th>
              <th class="min-w-150px cursor-pointer text-hover-primary">
                Notes
              </th>
              <th class="min-w-150px">Link</th>
            </tr>
            </thead>
            <!--end::Table head-->
            <!--begin::Table body-->
            <tbody>
            <tr v-for="transaction in filteredTransactions" :key="transaction" :style="{background: background}">
              <td class="text-muted fw-semibold fs-6 align-middle">
                <span v-if="transaction.TransactionDate !== '0001-01-01T00:00:00Z'" class="d-block fs-6">
                  {{ splitDate(transaction.TransactionDate) }}
                </span>
                <span v-else class="d-block fs-6">
                  {{ splitDateCreate(transaction.CreatedAt) }}
                </span>
              </td>
              <td class="align-middle">
                <div v-if="getUserInfo(transaction.userId)[1]" class="d-flex flex-column">
                  <span
                      class="text-dark fw-bold mb-1 fs-6">{{
                      getUserInfo(transaction.userId)[1]
                    }}, ID: {{ getUserInfo(transaction.userId)[0] }}</span>
                  <span
                      class="text-muted fw-semibold text-muted d-block fs-7">{{
                      getUserInfo(transaction.userId)[2]
                    }}</span>
                </div>
                <div v-else class="d-flex flex-column">
                  <span class="text-dark fw-bold mb-1 fs-6">-</span>
                </div>
              </td>
              <td class="align-middle">
                <span class="text-dark fw-bold d-block fs-6">
                  {{ transaction.provider }}
                </span>
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
              <td class="fs-6 align-middle cursor-pointer">
                <span v-if="transaction.provider === 'Grats.Cards'"
                      class="text-dark d-block fw-bold fs-6"> {{ transaction.balance }}</span>
                <span v-else class="text-dark d-block fw-bold fs-6">{{ getAccount(transaction) }}</span>
              </td>
              <td class="fs-6 align-middle cursor-pointer">
                <div class="d-flex flex-column">
                  <div v-if="transaction.increase" class="badge badge-light-success"
                       style="width: max-content;">+ {{ transactionBalance(transaction.amount) }}
                  </div>
                  <div v-else class="badge badge-light-danger"
                       style="width: max-content;">- {{ transactionBalance(transaction.amount) }}
                  </div>
                </div>
              </td>

              <td class="align-middle">
                <Popper class="fw-semibold fs-7" arrow hover placement="top" :content="transaction.merchant">
                  <span class="fw-bold d-block mb-1 fs-6 text-muted fw-semibold">{{
                      outputMerchant(transaction.merchant)
                    }}</span>
                </Popper>
              </td>

              <td class="align-middle">
                <span class="fw-bold d-block mb-1 fs-6 text-muted fw-semibold">{{ transaction.notes }}</span>
              </td>

              <td class="align-middle">
                <a v-if="transaction.transaction" class="btn btn-icon btn-bg-light btn-active-color-primary btn-sm" target="_blank"
                   :href="transaction.transaction">
                  <!--begin::Svg Icon | path: C:/wamp64/www/keenthemes/core/html/src/media/icons/duotune/arrows/arr095.svg-->
                  <span class="svg-icon svg-icon-muted svg-icon-3"><svg width="24" height="24" viewBox="0 0 24 24"
                                                                        fill="none" xmlns="http://www.w3.org/2000/svg">
                                      <path opacity="0.3"
                                            d="M4.7 17.3V7.7C4.7 6.59543 5.59543 5.7 6.7 5.7H9.8C10.2694 5.7 10.65 5.31944 10.65 4.85C10.65 4.38056 10.2694 4 9.8 4H5C3.89543 4 3 4.89543 3 6V19C3 20.1046 3.89543 21 5 21H18C19.1046 21 20 20.1046 20 19V14.2C20 13.7306 19.6194 13.35 19.15 13.35C18.6806 13.35 18.3 13.7306 18.3 14.2V17.3C18.3 18.4046 17.4046 19.3 16.3 19.3H6.7C5.59543 19.3 4.7 18.4046 4.7 17.3Z"
                                            fill="currentColor"/>
                                      <rect x="21.9497" y="3.46448" width="13" height="2" rx="1"
                                            transform="rotate(135 21.9497 3.46448)" fill="currentColor"/>
                                      <path
                                          d="M19.8284 4.97161L19.8284 9.93937C19.8284 10.5252 20.3033 11 20.8891 11C21.4749 11 21.9497 10.5252 21.9497 9.93937L21.9497 3.05029C21.9497 2.498 21.502 2.05028 20.9497 2.05028L14.0607 2.05027C13.4749 2.05027 13 2.52514 13 3.11094C13 3.69673 13.4749 4.17161 14.0607 4.17161L19.0284 4.17161C19.4702 4.17161 19.8284 4.52978 19.8284 4.97161Z"
                                          fill="currentColor"/>
                                      </svg>
                                  </span>
                  <!--end::Svg Icon-->
                </a>
                <span v-else class="badge badge-light-success" style="color: rgb(84 84 84)  !important; background-color: #DBDBE0FF !important;"> n/a </span>
              </td>

              <!--              <td class="border-0 align-middle">-->
              <!--                <div class="d-flex justify-content-end">-->
              <!--                  <Popper class="fw-semibold fs-7" arrow hover placement="top" content="Click to view full details">-->
              <!--                    <button @click="showTransactionModal(transaction)"-->
              <!--                            class="btn badge badge-light-primary btn-sm px-4">-->
              <!--                      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor"-->
              <!--                           class="bi bi-info-circle" viewBox="0 0 16 16">-->
              <!--                        <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"/>-->
              <!--                        <path-->
              <!--                            d="m8.93 6.588-2.29.287-.082.38.45.083c.294.07.352.176.288.469l-.738 3.468c-.194.897.105 1.319.808 1.319.545 0 1.178-.252 1.465-.598l.088-.416c-.2.176-.492.246-.686.246-.275 0-.375-.193-.304-.533L8.93 6.588zM9 4.5a1 1 0 1 1-2 0 1 1 0 0 1 2 0z"/>-->
              <!--                      </svg>-->
              <!--                    </button>-->
              <!--                  </Popper>-->
              <!--                </div>-->
              <!--              </td>-->
            </tr>
            </tbody>
            <!--end::Table body-->
          </table>
          <table ref="exportable_table" style="display:none;" v-if="getTransactions()" class="table table-row-bordered table-row-gray-100 align-middle gs-0 gy-3">
            <!--begin::Table head-->
            <thead class="rounded-left">
            <tr class="fw-bold text-muted bg-light">
              <th class="min-w-170px cursor-pointer text-hover-primary">
                ID
              </th>
              <th class="min-w-130px">Status</th>
              <th class="min-w-180px cursor-pointer text-hover-primary">
                Date of creation
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
              <th class="min-w-150px">Link</th>
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
          <ViewTransactionModal ref="view"/>
          <!--end::Table-->
        </div>
        <!--end::Table container-->
      </div>
      <!--begin::Body-->
    </div>
  </div>
</template>

<script>
import {mapGetters, mapActions} from 'vuex'
import ViewTransactionModal from './modals/viewTransactionModal.vue';
import timezones from "@/timezones.json";
import xlsx from "xlsx"

export default {
  name: 'tableTransactions',
  data() {
    return {
      background: "none",
      timezones: timezones,
      onTheTable: false,
      check: false,
      actionsShow: false,
      finalProfit: null,
      checkedTransactions: [],
      isHidden: false,
      copyNick: false,
      req: true,
      filteredTransactions: this.transactions,
      filterItems: {
        selectedUser: {code: 0, name: "", user: {}},
        initial: 0,
        total: 0,
        commission: 0,
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
  methods: {
    ...mapActions({
      deleteUserMethod: "admin/deleteUser",
      login: "auth/login",
    }),
    exportExcel(type, fn, dl) {
      let elt = this.$refs.exportable_table;
      let wb = xlsx.utils.table_to_book(elt, {sheet:"Sheet JS"});
      return dl ?
          xlsx.write(wb, {bookType:type, bookSST:true, type: 'base64'}) :
          xlsx.writeFile(wb, fn || (("transactions" + '.'|| 'SheetJSTableExport.') + (type || 'xls')));
    },
    usersListModel() {
      let users = Object.values(this.users)
      let user = users.filter(item => item)
      this.selectedUser = {code: user.ID, name: user.name + " — " + "ID:" + user.ID, user: user}
      this.isHidden = !this.isHidden
    },
    filterTransactionDate(date) {
      this.filterItems[date] !== 0 ? this.filterItems[date] = -this.filterItems[date] : this.filterItems[date] = 1
      // date === "TransactionDate" ? this.filterItems.DeletedAt = 0 : this.filterItems.TransactionDate = 0
      this.filter()
    },
    filterTransactionBalance(balance) {
      let stat = 0
      this.filterItems[balance] !== 0 ? stat = -this.filterItems[balance] : stat = 1
      this.filterItems.initial = 0
      this.filterItems.total = 0
      this.filterItems.comission = 0
      this.filterItems[balance] = stat
      this.filter()
    },
    filter() {
      this.filteredTransactions = this.transactions.filter(transaction => {
        if (transaction.userId) {
          let searchString = this.filterItems.search.toLowerCase()
          if ((!this.filterItems.balance.name || transaction.balance === this.filterItems.balance.name)
              && (transaction.userId)
              && (!this.filterItems.selectedUser.code || transaction.userId === this.filterItems.selectedUser.code)
              && (!this.filterItems.provider.name || transaction.provider === this.filterItems.provider.name)
              && (transaction.transaction.toLowerCase().indexOf(searchString) !== -1
                  || transaction.notes.toLowerCase().indexOf(searchString) !== -1
                  || transaction.userId.toString().indexOf(searchString)) !== -1) {
            return transaction
          }
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
      if (this.filterItems.commission !== 0) {
        this.filterItems.commission === 1 ? this.filteredTransactions.sort(function (a, b) {
          return new Date(b['commission']) - new Date(a['commission'])
        }) : this.filteredTransactions.sort(function (a, b) {
          return new Date(a['commission']) - new Date(b['commission'])
        })
      }

      if (this.filterItems.initial !== 0) {
        this.filterItems.initial === 1 ? this.filteredTransactions.sort(function (a, b) {
          return new Date(b['initial']) - new Date(a['initial'])
        }) : this.filteredTransactions.sort(function (a, b) {
          return new Date(a['initial']) - new Date(b['initial'])
        })
      }

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
        selectedUser: {code: 0, name: "", user: {}},
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
          code: null,
          name: null,
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
        TransactionDate: 1,
        initial: 0,
        commission: 0,
        total: 0,
        search: '',
      },
          this.filter()
    },
    showTransactionModal(transaction) {
      this.$refs.view.openModal(transaction)
    },
    getTransactions() {
      if (this.loggedIn && this.transactions.length === 0 && this.req) {
        this.$store.dispatch("admin/getAllTransactionsMethod")
      }
      if (this.loggedIn && this.users.length === 0 && this.req) {
        this.$store.dispatch("admin/getUsersMethod")
      }
      if (this.loggedIn && !this.cardsLoaded) {
        this.$store.dispatch("admin/getCardsMethod")
      }
      this.req = false
      return true
    },
    transactionBalance(number) {
      this.transactions.forEach(transaction => {
        this.finalProfit = transaction.amount - (transaction.amount - ((transaction.amount * transaction.commission) / 100))
      })
      if (number === 0) {
        return '0000.00'
      }
      if(!number) {
        return "-"
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
    splitDateCreate(date) {
      let dateList = date.split('T')
      let afterT = dateList[1].split('.')[0]
      let returnValue = dateList[0] + " " + afterT
      return returnValue
    },
    getUserInfo(ID) {
      if (!ID) {
        return "-"
      }
      let users = Object.values(this.users);
      let user = users.find(item => item.ID === ID);
      if (user === undefined || user === "" || user === null || !ID) {
        return "-"
      }
      return [user.ID, user.name, user.email]
    },
    getAccount(transaction) {
      let cards = Object.values(this.cards);
      let card = cards.find(item => item.ID === transaction.cardId);
      if (card === undefined || card === "" || card === null || !transaction) {
        return "-"
      }
      return "*" + card.cardNumber.toString().substring(12, 16)
    },
    outputMerchant(merchant) {
      if (merchant) {
        return merchant.split(" ")[0] + merchant.split(" ")[1]
      }
      return "-"
    }
  },
  computed: {
    ...mapGetters({
      currentUser: "auth/currentUser",
      users: "admin/getUsers",
      transactions: "admin/getAllTransactions",
      cards: "admin/getCards",
      cardsLoaded: "admin/getCardsLoaded",
      loggedIn: "auth/isLoggedIn",
      getTokenHeader: "auth/getTokenHeader"
    }),
    usersList() {
      let usersList = []
      let currentUsersID = []

      this.transactions.forEach(transaction => {
        currentUsersID.push(transaction.userId)
      })
      this.users.forEach(user => {
        for (const currentUserID of [...new Set(currentUsersID)]) {
          let userForList = {code: user.ID, name: "ID: " + user.ID + " - " + user.name + " - " + user.email, user: user}
          if (user.ID === currentUserID) {
            usersList.push(userForList)
          }
        }
      })
      return usersList
    },
  },
  components: {
    ViewTransactionModal
  }
}
</script>

<style scoped>
:deep(.popper) {
  background: #eaeaea !important;
  padding: 7px 10px !important;
  border-radius: 5px !important;
  z-index: 999;
  color: var(--kt-text-dark);
  /*-webkit-box-shadow: 0px 2px 3px 1px rgba(0, 0, 0, 0.14) !important;*/
  /*-moz-box-shadow: 0px 2px 3px 1px rgba(0, 0, 0, 0.14) !important;*/
  /*box-shadow: 0px 4px 8px 4px rgba(0, 0, 0, 0.14) !important;*/
}

:deep(.popper #arrow::before) {
  background: #eaeaea !important;
}

:deep(.popper:hover),
:deep(.popper:hover > #arrow::before) {
  background: #eaeaea !important;
}

.min-w-170px {
  min-width: 170px;
}

.table.gs-0 th:first-child,
.table.gs-0 td:first-child {
  padding-left: 1rem !important;
}

.show-delete {
  transition: opacity 0.2s ease !important;
}

.table th:last-child {
  padding: 0.75rem 0.75rem !important;
}
</style>
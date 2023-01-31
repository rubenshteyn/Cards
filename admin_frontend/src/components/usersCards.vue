<template>
  <div v-show="isHidden" @click="isHidden = false" class="position-absolute vh-100 w-100"
       style="z-index: 99;left: 0px;top: 0px;"></div>
  <div v-show="actionsShow" @click="actionsShow = false" class="position-absolute vh-100 w-100"
       style="z-index: 99;left: 0px;top: 0px;"></div>
  <div v-show="this.limitID" @click="this.limitID = 0" class="position-absolute vh-100 w-100"
       style="z-index: 99;left: 0px;top: 0px;"></div>
  <div class="d-flex flex-column user-cards flex-column-fluid container-fluid">
    <div class="toolbar mb-5 mb-lg-7" id="kt_toolbar">
      <!--begin::Page title-->
      <div class="page-title d-flex flex-column me-3">
        <!--begin::Title-->
        <h1 class="d-flex text-dark fw-bold my-1 fs-3">Card management</h1>
        <!--end::Title-->
        <!--begin::Breadcrumb-->
        <ul class="breadcrumb breadcrumb-dot fw-semibold text-gray-600 fs-7 my-1">
          <!--begin::Item-->
          <li class="breadcrumb-item text-gray-600">
            <router-link to="/" class="text-gray-600 text-hover-primary">Home</router-link>
          </li>
          <!--end::Item-->
          <!--begin::Item-->
          <li class="breadcrumb-item text-gray-500">Card Management</li>
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
            Filter
          </button>
          <!--begin::Menu 1-->
          <div id="filter-modal" v-show="isHidden"
               class="menu menu-sub menu-sub-dropdown show position-fixed float-left w-md-500px" data-kt-menu="true"
               style="z-index: 999; right: 40px; width: 300px">
            <!--begin::Header-->
            <div class="px-7 py-5">
              <div class="fs-5 text-dark fw-bold">Card filter</div>
            </div>
            <!--end::Header-->
            <!--begin::Menu separator-->
            <div class="separator border-gray-200"></div>
            <!--end::Menu separator-->
            <!--begin::Form-->
            <div class="px-7 py-5 scroll-y mh-500px">

              <div class="mb-10">
                <div class="form-group fv-row">
                  <label class="form-label fw-semibold" for="status">Client:</label>
                  <v-select class="bg-transparent" label="name" :options="usersList"
                            placeholder="Select client" v-model="filterItems.selectedUser" required>
                  </v-select>
                </div>
              </div>

              <div class="mb-10">
                <div class="form-group fv-row">
                  <label class="form-label fw-semibold" for="status">Card status:</label>
                  <v-select class="bg-transparent" label="name" :options="filterItems.statuses"
                            placeholder="Select card status" v-model="filterItems.status" required>
                  </v-select>
                </div>
              </div>

              <div class="mb-10">
                <div class="form-group fv-row">
                  <label class="form-label fw-semibold" for="status">Card type:</label>
                  <v-select class="bg-transparent" label="name" :options="filterItems.types"
                            placeholder="Select card type" v-model="filterItems.type" required>
                  </v-select>
                </div>
              </div>

              <div class="mb-10">
                <div class="form-group fv-row">
                  <label class="form-label fw-semibold" for="status">Payment system:</label>
                  <v-select class="bg-transparent" label="name" :options="filterItems.systemTypes"
                            placeholder="Select payment system" v-model="filterItems.systemType" required>
                  </v-select>
                </div>
              </div>

              <div class="mb-10">
                <div class="form-group fv-row">
                  <label class="form-label fw-semibold" for="search">Card search</label>
                  <input class="form-control form-control-lg form-control-solid" label="name"
                         v-model="filterItems.search">
                </div>
              </div>

              <div class="d-flex justify-content-end">
                <button @click="resetFilterCards" type="reset"
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
        <!--begin::Button-->
        <button @click='showModal("create")' class="btn btn-dark fw-bold" data-bs-toggle="modal"
                data-bs-target="#kt_modal_create_app"
                id="kt_toolbar_primary_button">Create card
        </button>
        <!--end::Button-->
      </div>
      <!--end::Actions-->
    </div>
    <div class="card mb-5 mb-xl-8">
      <!--begin::Header-->
      <div class="card-header align-items-center card-header-card border-0 pt-5">
        <h3 class="card-title align-items-start flex-column">
          <span class="card-label fw-bold fs-3 mb-1">All cards</span>
          <span class="text-muted mt-1 fw-semibold fs-7">Amount cards: {{ this.filteredCards.length }} </span>
        </h3>
        <div class="d-flex align-items-center">
          <a download="cards.txt" id="test" :class="{disabled: showExportSelectCards()}" href="#" @click="exportTXT()"
             class="btn show-delete btn-icon btn-bg-light btn-active-color-primary btn-sm">
            <!--begin::Svg Icon | path: C:/wamp64/www/keenthemes/core/html/src/media/icons/duotune/files/fil017.svg-->
            <span class="svg-icon svg-icon-muted svg-icon-2hx"><svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path opacity="0.3" d="M10 4H21C21.6 4 22 4.4 22 5V7H10V4Z" fill="currentColor"/>
                <path opacity="0.3" d="M13 14.4V9C13 8.4 12.6 8 12 8C11.4 8 11 8.4 11 9V14.4H13Z" fill="currentColor"/>
                <path d="M10.4 3.60001L12 6H21C21.6 6 22 6.4 22 7V19C22 19.6 21.6 20 21 20H3C2.4 20 2 19.6 2 19V4C2 3.4 2.4 3 3 3H9.20001C9.70001 3 10.2 3.20001 10.4 3.60001ZM13 14.4V9C13 8.4 12.6 8 12 8C11.4 8 11 8.4 11 9V14.4H8L11.3 17.7C11.7 18.1 12.3 18.1 12.7 17.7L16 14.4H13Z" fill="currentColor"/>
                </svg>
              </span>
            <!--end::Svg Icon-->
            <!--            {{ $t('cards.unloading') }}-->
          </a>
        </div>
      </div>
      <!--end::Header-->
      <!--begin::Body-->
      <div class="card-body py-3">
        <!--begin::Table container-->
        <div class="table-responsive">
          <!--begin::Table-->
          <table v-if="getCards()" class="table table-row-bordered table-row-gray-100 align-middle gs-0 gy-3">
            <!--begin::Table head-->
            <thead class="rounded-left">
            <tr class="fw-bold text-muted bg-light">
              <th class="w-25px">
                <div class="form-check form-check-sm form-check-custom form-check-solid mr-3">
                  <input @click="selectAllCards" v-model="check" class="form-check-input" type="checkbox"/>
                </div>
              </th>
              <th class="min-w-300px">Card</th>
              <th @click="filterCardsBalance('limit')" class="min-w-150px cursor-pointer">Limit</th>
              <th @click="filterCardsBalance('spend')" class="min-w-150px cursor-pointer">Spend</th>
              <th @click="filterCardsBalance('balance')" class="min-w-150px cursor-pointer">Balance</th>
              <th class="min-w-150px">User</th>
              <th @click="filterCardDate('CreatedAt')" class="min-w-150px cursor-pointer">Date of creation</th>
              <th @click="filterCardDate('deactivateDate')" class="min-w-150px cursor-pointer">Closing date</th>
              <th @click="filterCardDate('CreatedAt')" class="min-w-150px cursor-pointer">Card Age</th>
              <th class="min-w-150px">Status</th>
              <th class="min-w-100px">Actions</th>
            </tr>
            </thead>
            <!--end::Table head-->
            <!--begin::Table body-->
            <tbody>
            <tr v-for="card in this.filteredCards" :key="card" :style="{background: background}">
              <td class="align-middle">
                <div class="form-check form-check-sm form-check-custom form-check-solid">
                  <input v-bind:value="card" v-model="checkedCards"
                         class="form-check-input widget-9-check"
                         type="checkbox"/>
                </div>
              </td>
              <td class="align-middle">
                <div class="d-flex align-items-center">
                  <div class="symbol symbol-50px me-5">
                    <span class="symbol-label bg-light">
                      <!--begin::Svg Icon | path: C:/wamp64/www/keenthemes/core/html/src/media/icons/duotune/finance/fin002.svg-->
                        <span v-if="card.cardNumber.toString().split('')[0] === '5'"
                              class="svg-icon svg-icon-dark svg-icon-3hx">
                          <svg xmlns="http://www.w3.org/2000/svg" width="2.11676in" height="1.5in"
                               viewBox="0 0 152.407 108">
                              <g>
                                <rect width="152.407" height="108" style="fill: none"/>
                                <g>
                                  <rect x="60.4117" y="25.6968" width="31.5" height="56.6064" style="fill: #ff5f00"/>
                                  <path
                                      d="M382.20839,306a35.9375,35.9375,0,0,1,13.7499-28.3032,36,36,0,1,0,0,56.6064A35.938,35.938,0,0,1,382.20839,306Z"
                                      transform="translate(-319.79649 -252)" style="fill: #eb001b"/>
                                  <path
                                      d="M454.20349,306a35.99867,35.99867,0,0,1-58.2452,28.3032,36.00518,36.00518,0,0,0,0-56.6064A35.99867,35.99867,0,0,1,454.20349,306Z"
                                      transform="translate(-319.79649 -252)" style="fill: #f79e1b"/>
                                  <path
                                      d="M450.76889,328.3077v-1.1589h.4673v-.2361h-1.1901v.2361h.4675v1.1589Zm2.3105,0v-1.3973h-.3648l-.41959.9611-.41971-.9611h-.365v1.3973h.2576v-1.054l.3935.9087h.2671l.39351-.911v1.0563Z"
                                      transform="translate(-319.79649 -252)" style="fill: #f79e1b"/>
                                </g>
                              </g>
                            </svg>
                        </span>
                      <!--end::Svg Icon-->
                      <!--begin::Svg Icon | path: C:/wamp64/www/keenthemes/core/html/src/media/icons/duotune/finance/fin002.svg-->
                        <span v-else class="svg-icon svg-icon-dark svg-icon-3hx">
                          <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"
                               height="512px" style="enable-background:new 0 0 512 512;" version="1.1"
                               viewBox="0 0 512 512" width="512px" xml:space="preserve"><g id="形状_1_3_" style="enable-background:new    ;"><g id="形状_1"><g><path d="M211.328,184.445l-23.465,144.208h37.542l23.468-144.208     H211.328z M156.276,184.445l-35.794,99.185l-4.234-21.358l0.003,0.007l-0.933-4.787c-4.332-9.336-14.365-27.08-33.31-42.223     c-5.601-4.476-11.247-8.296-16.705-11.559l32.531,124.943h39.116l59.733-144.208H156.276z M302.797,224.48     c0-16.304,36.563-14.209,52.629-5.356l5.357-30.972c0,0-16.534-6.288-33.768-6.288c-18.632,0-62.875,8.148-62.875,47.739     c0,37.26,51.928,37.723,51.928,57.285c0,19.562-46.574,16.066-61.944,3.726l-5.586,32.373c0,0,16.763,8.148,42.382,8.148     c25.616,0,64.272-13.271,64.272-49.37C355.192,244.272,302.797,240.78,302.797,224.48z M455.997,184.445h-30.185     c-13.938,0-17.332,10.747-17.332,10.747l-55.988,133.461h39.131l7.828-21.419h47.728l4.403,21.419h34.472L455.997,184.445z      M410.27,277.641l19.728-53.966l11.098,53.966H410.27z" style="fill-rule:evenodd;clip-rule:evenodd;fill:#005BAC;"/></g></g></g><g
                              id="形状_1_2_" style="enable-background:new    ;"><g id="形状_1_1_"><g><path d="M104.132,198.022c0,0-1.554-13.015-18.144-13.015H25.715     l-0.706,2.446c0,0,28.972,5.906,56.767,28.033c26.562,21.148,35.227,47.51,35.227,47.51L104.132,198.022z" style="fill-rule:evenodd;clip-rule:evenodd;fill:#F6AC1D;"/></g></g></g></svg>
                        </span>
                      <!--end::Svg Icon-->
                    </span>
                  </div>
                  <div class="d-flex justify-content-start flex-column">
                    <div class="d-flex align-items-center">
                      <span class="text-dark fw-bold mb-1 fs-6 mr-2">{{ cardNumber(card.cardNumber) }}</span>
                      <span v-if="!card.cardNumber.toString().includes('486555')"
                            class="badge badge-light-danger fw-semibold">Ads</span>
                      <span v-else class="badge badge-light-info fw-semibold">Uni</span>
                    </div>
                    <span class="text-muted fw-semibold text-muted d-block fs-7">Дата:
                                                {{ card.expirationDate }} CVV: {{ card.CVV }}</span>
                  </div>
                </div>
              </td>
              <td class="fs-6 align-middle position-relative">
                <div class="d-flex  align-items-center fs-6">
                  <span class="text-dark fw-bold d-block fs-6">${{ card.limit }}</span>
                  <button @click="this.limitID = card.ID"
                          class="btn btn-bg-light btn-color-primary btn-active-color-primary btn-sm"
                          style="padding: 3px 7px; margin-left: 5px;">+
                  </button>
                </div>
                <!--                &lt;!&ndash; поле для пополнения start&ndash;&gt;-->
                <div :id="this.limitID" v-if="showEditLimit(card)"
                     class="edit-card-limit form-group d-flex align-items-center">
                  <span class="limit-plus fs-6">+</span>
                  <input type="number" class="form-control form-control-lg mr-1"
                         id="limit" name="name" v-model="form.amount" style="z-index: 9999;"/>
                  <button @click="checkForm"
                          class="btn btn-bg-light btn-color-primary btn-active-color-primary btn-sm form-control-lg"
                          style="min-height: calc(1.5em + 1.65rem + 2px)!important;">Ок
                  </button>
                </div>
                <!-- поле для пополнения end -->
              </td>
              <td class="align-middle">
                <div class="text-dark fw-bold d-block fs-6">
                  ${{ card.spend }}
                </div>
              </td>
              <td class="align-middle">
                <div class="text-dark fw-bold d-block fs-6">${{ card.balance }}</div>
              </td>
              <td class="align-middle">
                <div class="d-flex flex-column">
                  <span class="text-dark fw-bold mb-1 fs-6">{{ getUserInfo(card.userId)[0] }}</span>
                  <span class="text-muted fw-semibold text-muted d-block fs-7">{{ getUserInfo(card.userId)[1] }}</span>
                </div>
              </td>
              <td class="align-middle">
                <div class="d-flex justify-content-start flex-column">
                  <span class="text-dark fw-bold mb-1 fs-6">
                  {{ splitDate(card.CreatedAt)[0] }}
                </span>
                  <span class="text-muted fw-semibold text-muted d-block fs-7">
                  {{ splitDate(card.CreatedAt)[1] }}
                </span>
                </div>
              </td>
              <td class="align-middle">
                <span v-if="card.status === 1" class="badge badge-light-success"
                      style="color: rgb(84 84 84)  !important; background-color: #DBDBE0FF !important;"> n/a </span>
                <div v-else class="d-flex justify-content-start flex-column">
                  <span class="text-dark fw-bold fs-6">
                  {{ splitDate(card.deactivateDate)[0] }}
                </span>
                  <span class="text-muted fw-semibold text-muted d-block fs-7">
                {{ splitDate(card.deactivateDate)[1] }}
                </span>
                </div>
              </td>
              <td class="align-middle">
                <span class="text-dark fw-bold fs-6">
                  {{ cardAge(card.CreatedAt, card.deactivateDate) }}
                </span>
              </td>
              <td class="text-muted fw-semibold text-hover-primary fs-6 align-middle">
                  <span v-if="card.status === 0" class="badge badge-light-danger">
                  Blocked
                </span>
                <span v-if="card.status === 1" class="badge badge-light-success">
                  Active
                </span>
              </td>
              <td class="border-0 align-middle">
                <div class="d-flex align-items-center">
                  <div class="me-2">
                    <Popper class="fw-semibold fs-7" arrow hover placement="top" content="Click to top up balance">
                      <button @click='showTopUpModal(card)'
                              class="btn text-muted btn-bg-light btn-active-color-primary btn-sm"
                              style="white-space: nowrap; ">
                        Top up
                      </button>
                    </Popper>
                  </div>
                  <div class="me-2">
                    <Popper class="fw-semibold fs-7" arrow hover placement="top" content="Click to view log">
                      <button @click='showLogModal(card)'
                              class="btn text-muted btn-bg-light btn-active-color-primary btn-sm"
                              style="white-space: nowrap; ">
                        Log
                      </button>
                    </Popper>
                  </div>
                  <Popper class="fw-semibold fs-7" arrow hover placement="top" content="Click to change card">
                    <button @click='showEditCardModal(card)'
                            class="btn show-delete btn-icon btn-bg-light btn-active-color-primary btn-sm">
                      <!--begin::Svg Icon | path: icons/duotune/art/art005.svg-->
                      <span class="svg-icon svg-icon-3" data-v-5939e3ea="">
                    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg"
                         data-v-5939e3ea="">
                      <path opacity="0.3"
                            d="M21.4 8.35303L19.241 10.511L13.485 4.755L15.643 2.59595C16.0248 2.21423 16.5426 1.99988 17.0825 1.99988C17.6224 1.99988 18.1402 2.21423 18.522 2.59595L21.4 5.474C21.7817 5.85581 21.9962 6.37355 21.9962 6.91345C21.9962 7.45335 21.7817 7.97122 21.4 8.35303ZM3.68699 21.932L9.88699 19.865L4.13099 14.109L2.06399 20.309C1.98815 20.5354 1.97703 20.7787 2.03189 21.0111C2.08674 21.2436 2.2054 21.4561 2.37449 21.6248C2.54359 21.7934 2.75641 21.9115 2.989 21.9658C3.22158 22.0201 3.4647 22.0084 3.69099 21.932H3.68699Z"
                            fill="currentColor" data-v-5939e3ea=""></path>
                      <path
                          d="M5.574 21.3L3.692 21.928C3.46591 22.0032 3.22334 22.0141 2.99144 21.9594C2.75954 21.9046 2.54744 21.7864 2.3789 21.6179C2.21036 21.4495 2.09202 21.2375 2.03711 21.0056C1.9822 20.7737 1.99289 20.5312 2.06799 20.3051L2.696 18.422L5.574 21.3ZM4.13499 14.105L9.891 19.861L19.245 10.507L13.489 4.75098L4.13499 14.105Z"
                          fill="currentColor" data-v-5939e3ea=""></path>
                    </svg>
                  </span>
                      <!--end::Svg Icon-->
                    </button>
                  </Popper>
                </div>
              </td>
            </tr>
            </tbody>
            <!--end::Table body-->
          </table>
          <!--end::Table-->
          <TransactionLogModal ref="transactions"/>
          <CreateCardModal ref="create"/>
          <EditCardModal ref="editcard"/>
          <TopUpCard ref="topupcard"/>
        </div>
        <!--end::Table container-->
      </div>
      <!--begin::Body-->
    </div>
  </div>
</template>

<script>
import {mapActions, mapGetters} from 'vuex'
import TransactionLogModal from '@/components/modals/transactionLogModal.vue'
import CreateCardModal from "@/components/modals/createCardModal";
import EditCardModal from "@/components/modals/editCardModal";
import TopUpCard from "@/components/modals/topUpCard";
import timezones from "@/timezones.json";

export default {
  name: "usersCards",
  data() {
    return {
      userName: null,
      userEmail: null,
      background: "none",
      timezones: timezones,
      req: true,
      check: false,
      isHidden: false,
      actionsShow: false,
      checkedCards: [],
      filteredCards: this.cards,
      limitID: 0,
      form: {
        amount: null,
        cardID: 0,
        balance: "USDT"
      },
      filterItems: {
        selectedUser: {code: 0, name: "", user: {}},
        CreatedAt: 1,
        deactivateDate: 0,
        balance: 0,
        spend: 0,
        limit: 0,
        selectedBIN: {name: "", isUniversal: null, active: null},
        types: [
          {
            name: "Universal",
            code: false
          },
          {
            name: "Advertising",
            code: true
          }
        ],
        type: {
          name: null,
          code: null
        },
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
        statuses: [
          {
            name: "Active",
            code: 1
          },
          {
            name: "Blocked",
            code: 0
          }
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
    this.filteredCards = this.cards
    if (this.cards) {
      this.filter()
    }
  },
  watch: {
    cards: "filter"
  },
  methods: {
    ...mapActions({
      topUpCardBalance: "admin/topUpCardBalance"
    }),
    exportTXT() {
      let type = 'data:application/octet-stream;base64, '
      let textArr = []
      for (const card of this.checkedCards) {
        let text = card.cardNumber + "|" + card.expirationDate + "|" + card.CVV + "|" + card.name
        textArr.push(text)
      }
      let base = btoa(textArr.toString().replace(/,/g,"\n"))
      let res = type + base
      document.getElementById('test').href = res
    },
    cardAge(created, deactivated) {
      created = created.split('T')[0].replace(/[-/]/g, '/').split('/')[1] + "/" + created.split('T')[0].replace(/[-/]/g, '/').split('/')[2] + "/" + created.split('T')[0].replace(/[-/]/g, '/').split('/')[0]
      let currentDate = new Date().toLocaleDateString().replace(/[./]/g, '/')
      currentDate = currentDate.split('/')[1] + "/" + currentDate.split('/')[0] + "/" + currentDate.split('/')[2]

      if(deactivated !== "0001-01-01T00:00:00Z") {
        currentDate = deactivated.split('T')[0].replace(/[-/]/g, '/').split('/')[1] + "/" + deactivated.split('T')[0].replace(/[-/]/g, '/').split('/')[2] + "/" + deactivated.split('T')[0].replace(/[-/]/g, '/').split('/')[0]
      }
      const date1 = new Date(created);
      const date2 = new Date(currentDate);
      // One day in milliseconds
      const oneDay = 1000 * 60 * 60 * 24;

      // Calculating the time difference between two dates
      const diffInTime = date2.getTime() - date1.getTime();

      // Calculating the no. of days between two dates
      const diffInDays = Math.round(diffInTime / oneDay);

      if(diffInDays === 1) {
        return diffInDays + " day"
      }
      return diffInDays + " days";
    },
    checkForm: function () {
      this.limitID = 0
      this.form.amount = Number(this.form.amount)
      this.topUpCardBalance(this.form)
          .then((response) => {
            console.log(response)
          })
          .catch(error => {
            console.log(error)
            this.errors.responseError = error
          });
    },
    changeCheckedCards() {
      let selectCards = []
      let IDs = []
      for (let [i, currentSelectCards] in this.checkedCards) {
        currentSelectCards = this.checkedCards[i]
        selectCards.push(currentSelectCards)
        for (const selectCard of selectCards) {
          IDs.push(selectCard)
        }
      }
    },
    filterCardDate(date) {
      this.filterItems[date] !== 0 ? this.filterItems[date] = -this.filterItems[date] : this.filterItems[date] = 1
      date === "CreatedAt" ? this.filterItems.deactivateDate = 0 : this.filterItems.CreatedAt = 0
      this.filter()
    },
    filterCardsBalance(balance) {
      let stat = 0
      this.filterItems[balance] !== 0 ? stat = -this.filterItems[balance] : stat = 1
      this.filterItems.limit = 0
      this.filterItems.spend = 0
      this.filterItems.balance = 0
      this.filterItems[balance] = stat
      this.filter()
    },
    checkingValues(cardElement) {
      if (cardElement === "" || typeof cardElement === undefined || cardElement === null) {
        return cardElement === "-"
      }
      if (cardElement === 0) {
        return cardElement === 0
      }
      return cardElement
    },
    filter() {
      this.filteredCards = this.cards.filter(card => {
        let cardInCycle = Object.values(card)
        for (const cardElement of cardInCycle) {
          this.checkingValues(cardElement)
        }
        let searchString = this.filterItems.search.toLowerCase()
        this.filterItems.selectedBIN.isUniversal = this.filterItems.type.code
        let cardBIN = Number(card.cardNumber.toString().split("")[0] + card.cardNumber.toString().split("")[1] + card.cardNumber.toString().split("")[2] + card.cardNumber.toString().split("")[3] + card.cardNumber.toString().split("")[4] + card.cardNumber.toString().split("")[5])
        let BINBoolean = this.BINsList.some(item => {
          if (item.name === cardBIN) {
            return true
          }
        })
        if ((!this.filterItems.status.name || card.status === this.filterItems.status.code)
            && card.cardNumber
            && (!this.filterItems.selectedUser.code || card.userId === this.filterItems.selectedUser.code)
            && (!this.filterItems.systemType.code || Number(card.cardNumber.toString().split("")[0]) === this.filterItems.systemType.code)
            && (card.cardNumber.toLowerCase().indexOf(searchString) !== -1) && (!BINBoolean)) {
          return card
        }
      })
      if (this.filterItems.search && this.filteredCards.length > 1) {
        this.background = "rgb(230, 255, 241)"
      } else {
        this.background = "none"
      }
      // фильтрация по дате
      if (this.filterItems.CreatedAt !== 0) {
        this.filterItems.CreatedAt === 1 ? this.filteredCards.sort(function (a, b) {
          return new Date(b['CreatedAt']) - new Date(a['CreatedAt'])
        }) : this.filteredCards.sort(function (a, b) {
          return new Date(a['CreatedAt']) - new Date(b['CreatedAt'])
        })
      }

      if (this.filterItems.deactivateDate !== 0) {
        this.filterItems.deactivateDate === 1 ? this.filteredCards.sort(function (a, b) {
          return new Date(b['deactivateDate']) - new Date(a['deactivateDate'])
        }) : this.filteredCards.sort(function (a, b) {
          return new Date(a['deactivateDate']) - new Date(b['deactivateDate'])
        })
      }

      // фильтрация по балансам
      if (this.filterItems.balance !== 0) {
        this.filterItems.balance === 1 ? this.filteredCards.sort(function (a, b) {
          return new Date(b['balance']) - new Date(a['balance'])
        }) : this.filteredCards.sort(function (a, b) {
          return new Date(a['balance']) - new Date(b['balance'])
        })
      }

      if (this.filterItems.spend !== 0) {
        this.filterItems.spend === 1 ? this.filteredCards.sort(function (a, b) {
          return new Date(b['spend']) - new Date(a['spend'])
        }) : this.filteredCards.sort(function (a, b) {
          return new Date(a['spend']) - new Date(b['spend'])
        })
      }

      if (this.filterItems.limit !== 0) {
        this.filterItems.limit === 1 ? this.filteredCards.sort(function (a, b) {
          return new Date(b['limit']) - new Date(a['limit'])
        }) : this.filteredCards.sort(function (a, b) {
          return new Date(a['limit']) - new Date(b['limit'])
        })
      }
    },
    resetFilterCards() {
      this.background = "none"
      this.filterItems = {
        cardNumberData: null,
        selectedUser: {code: 0, name: "", user: {}},
        CreatedAt: 1,
        balance: 0,
        spend: 0,
        limit: 0,
        selectedBIN: {name: "", isUniversal: null, active: null},
        types: [
          {
            name: "Universal",
            code: false
          },
          {
            name: "Advertising",
            code: true
          }
        ],
        type: {
          name: null,
          code: null
        },
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
        statuses: [
          {
            name: "Active",
            code: 1
          },
          {
            name: "Blocked",
            code: 0
          }
        ],
        status: {
          name: null,
          code: null
        },
        search: '',
      },
          this.BINsList
      this.filter()
    },
    deleteSelectCards() {
      return true
    },
    showExportSelectCards() {
      return this.checkedCards.length <= 0
    },
    selectAllCards() {
      if (this.check === true) {
        return this.checkedCards = []
      }
      this.checkedCards = this.cards
      return [this.checkedCards = this.cards, this.check === true]
    },
    showModal(ref) {
      this.$refs[ref].openModal()
    },
    showEditCardModal(card) {
      let users = Object.values(this.users);
      let user = users.find(item => item.ID === card.userId);
      this.$refs.editcard.openModal(card, user)
    },
    showLogModal(card) {
      let transactions = Object.values(this.transactions);
      let transaction = transactions.filter(item => {
        if (item.cardId === card.ID && item.provider === "AdMediaCards") {
          return item
        }
      })
      this.$refs.transactions.openModal(card, transaction)
    },
    showTopUpModal(card) {
      this.$refs.topupcard.openModal(card)
    },
    showEditLimit(card) {
      if (card.ID === this.limitID) {
        this.form.cardID = card.ID
        this.cardNumberData = card.cardNumber
        return true
      }
    },
    getCards() {
      if (this.loggedIn && this.users.length === 0 && this.req) {
        this.$store.dispatch("admin/getUsersMethod")
      }
      if (this.loggedIn && !this.cardsLoaded) {
        this.$store.dispatch("admin/getCardsMethod");
      }
      if (this.loggedIn && this.transactions.length === 0 && this.req) {
        this.$store.dispatch("admin/getAllTransactionsMethod")
      }
      if (this.loggedIn && this.suppliers.length === 0 && this.req) {
        this.$store.dispatch("admin/getSuppliersMethod")
      }
      this.req = false
      return true;
    },
    cardBalance(balance) {
      if (balance === 0) {
        return "000 000,000$";
      }
    },
    cardNumber(number) {
      if (number) {
        return number.toString().substring(0, 4) + " " + number.toString().substring(4, 8) + " " + number.toString().substring(8, 12) + " " + number.toString().substring(12, 16)
      }
      return "-"
    },
    getUserInfo(ID) {
      if (!ID) {
        return ["-", "-"]
      }
      let users = Object.values(this.users);
      let user = users.find(item => item.ID === ID);
      if (user === undefined || user === "" || user === null || !ID) {
        return ["-", "-"]
      }
      return [user.name, user.email]
    },
    splitDate(date) {
      let userTimeZone = this.currentUser.timeZone

      let timeZone = this.timezones.find(item => item.text === userTimeZone);
      if (date === "0001-01-01T00:00:00Z" || date === "" || date === null || !date || !timeZone) {
        return ["-", "-"]
      }

      if(timeZone && timeZone.utc.includes('-')) {
        let newDate = new Date(date);
        let utc = Number(timeZone.utc.split('-')[1])
        newDate.setTime(newDate.getTime() - (utc * 60 * 60 * 1000))
        return [newDate.toLocaleString().split(",")[0], newDate.toLocaleString().split(",")[1]]
      }
      if(timeZone && !timeZone.utc.includes('-')) {
        let newDate = new Date(date);
        let utc = Number(timeZone.utc)
        newDate.setTime(newDate.getTime() + (utc * 60 * 60 * 1000))
        return [newDate.toLocaleString().split(",")[0], newDate.toLocaleString().split(",")[1]]
      }
    },
  },
  computed: {
    ...mapGetters({
      cardsLoaded: "admin/getCardsLoaded",
      users: "admin/getUsers",
      cards: "admin/getCards",
      transactions: "admin/getAllTransactions",
      suppliers: "admin/getSuppliers",
      loggedIn: "auth/isLoggedIn",
      currentUser: "auth/currentUser"
    }),
    usersList() {
      let usersList = []
      this.users.forEach(user => {
        let userForList = {code: user.ID, name: "ID: " + user.ID + " - " + user.name + " - " + user.email, user: user}
        if (user.status === 2) {
          usersList.push(userForList)
        }
      })
      return usersList
    },
    BINsList() {
      let BINsList = []
      let initialBINsList = []
      if (this.loggedIn && this.suppliers.length === 0 && this.req) {
        this.$store.dispatch("admin/getSuppliersMethod")
      }
      this.suppliers.forEach(provider => {
        initialBINsList = JSON.parse(provider.bins)
      })
      initialBINsList.forEach(bin => {
        let legitBin = {name: bin.bin, isUniversal: bin.isUniversal, active: bin.active}
        if (bin.isUniversal === this.filterItems.type.code) {
          BINsList.push(legitBin)
        }
      })
      return BINsList
    },
  }
  ,
  components: {
    TopUpCard,
    EditCardModal,
    TransactionLogModal,
    CreateCardModal,
  }
}
</script>

<style scoped>
.disabled {
  pointer-events: none;
  cursor: default;
  opacity: 0.5;
}
.limit-plus {
  z-index: 99999;
  position: absolute;
  top: 28%;
  left: 0.9rem;
}

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
</style>

<style>
.show-delete {
  transition: opacity 0.2s ease !important;
}

.table.gs-0 th:first-child,
.table.gs-0 td:first-child {
  padding-left: 1rem !important;
}

.vs__selected, .vs__search {
  font-size: 1.1rem;
  font-weight: 500;
}

.edit-card-limit {
  width: 100%;
  position: absolute;
  border-radius: 0.625rem;
  top: 10px;
  left: -40px;
  background: unset;
  z-index: 99;
}
</style>
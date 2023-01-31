<template>
  <div v-show="isHidden" @click="isHidden = false" class="position-absolute vh-100 w-100"
       style="z-index: 99;left: 0px;top: 0px;"></div>
  <div v-show="showTopup" @click="showTopup = false" class="position-absolute vh-100 w-100"
       style="z-index: 99;left: 0px;top: 0px;"></div>
  <div class="d-flex flex-column flex-column-fluid container-fluid">
    <div class="toolbar mb-5 mb-lg-7" id="kt_toolbar">
      <!--begin::Page title-->
      <div class="page-title d-flex flex-column me-3">
        <!--begin::Title-->
        <h1 class="d-flex text-dark fw-bold my-1 fs-3">{{ $t('cards.title') }}</h1>
        <!--end::Title-->
        <!--begin::Breadcrumb-->
        <ul class="breadcrumb breadcrumb-dot fw-semibold text-gray-600 fs-7 my-1">
          <!--begin::Item-->
          <li class="breadcrumb-item text-gray-600">
            <router-link to="/" class="text-gray-600 text-hover-primary">{{ $t('common.crumb1') }}</router-link>
          </li>
          <!--end::Item-->
          <!--begin::Item-->
          <li class="breadcrumb-item text-gray-500">{{ $t('cards.title') }}</li>
          <!--end::Item-->
        </ul>
        <!--end::Breadcrumb-->
      </div>
      <!--end::Page title-->
      <!--begin::Actions-->
      <div class="actions-group-buttons d-flex align-items-center py-2 py-md-1">
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
            {{ $t('common.filter') }}
          </button>
          <!--begin::Menu 1-->
          <div id="filter-modal" v-show="isHidden"
               class="menu menu-sub menu-sub-dropdown show position-fixed float-left" data-kt-menu="true"
               style="width: inherit; z-index: 999; right: 40px;">
            <!--begin::Header-->
            <div class="px-7 py-5">
              <div class="fs-5 text-dark fw-bold">{{ $t('cards.filterCards') }}</div>
            </div>
            <!--end::Header-->
            <!--begin::Menu separator-->
            <div class="separator border-gray-200"></div>
            <!--end::Menu separator-->
            <!--begin::Form-->
            <div class="px-7 py-5">
              <div class="mb-10">
                <div class="form-group fv-row">
                  <label class="form-label fw-semibold" for="status">{{ $t('cards.statusCard') }}:</label>
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
                  <label class="form-label fw-semibold" for="search">{{ $t('cards.cardSearch') }}</label>
                  <input class="form-control-solid form-control form-control-lg" label="name"
                         v-model="filterItems.search">
                </div>
              </div>


              <div class="d-flex">
                <button @click="resetFilterCards" type="reset"
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
        <!--end::Wrapper-->
        <!--begin::Button-->
        <router-link to="/rates" class="btn btn-dark fw-bold">{{ $t('cards.create') }}
        </router-link>
        <!--end::Button-->
      </div>
      <!--end::Actions-->
    </div>
    <div class="card mb-5 mb-xl-8">
      <!--begin::Header-->
      <div class="card-header border-0 pt-5">
        <h3 class="card-title align-items-start flex-column">
          <span class="card-label fw-bold fs-3 mb-1">{{ $t('cards.tableTitle') }}</span>
          <span class="text-muted mt-1 fw-semibold fs-7">{{ $t('cards.amount') }}: {{ this.cards.length }}</span>
        </h3>
        <div class="d-flex align-items-center">
          <h3 v-if="!showTopup" class="card-title align-items-start flex-column">
            <span class="card-label fw-bold fs-3 mb-1">{{ $t('cards.accountBalance') }}</span>
            <span class="text-muted mt-1 fw-semibold fs-7">USDT: {{ userBalance(this.userinfo.USDTBalance) }} ₮</span>
          </h3>
          <button v-if="!showTopup" @click="showTopup = true" class="btn btn-sm btn-primary">
            {{ $t('cards.topUp') }}
          </button>
          <div v-if="showTopup" class="d-flex align-items-center min-w-150px" style="z-index: 999;">
            <input class="form-control form-control-lg form-control-solid" v-model="form.balance" type="number">
            <button @click="checkForm()" class="btn btn-sm btn-primary btn-active-light">OK</button>
          </div>
        </div>
        <div class="d-flex align-items-center">
          <a download="cards.txt" :class="{disabled: showExportSelectCards()}" id="test" href="#" @click="exportTXT()"
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
          <table ref="exportable_table" v-if="getData()" class="table table-row-bordered table-row-gray-100 align-middle gs-0 gy-3">
            <!--begin::Table head-->
            <thead class="rounded-left">
            <tr class="fw-bold text-muted bg-light">
              <th class="w-25px">
                <div class="form-check form-check-sm form-check-custom form-check-solid mr-3">
                  <input @click="selectAllCards" v-model="check" class="form-check-input" type="checkbox"/>
                </div>
              </th>
              <th class="min-w-300px">{{ $t('cards.card') }}</th>
              <th @click="filterCardsBalance('limit')" class="min-w-150px cursor-pointer">{{ $t('cards.limit') }}</th>
              <th @click="filterCardsBalance('spend')" class="min-w-150px cursor-pointer">{{ $t('cards.spend') }}</th>
              <th @click="filterCardsBalance('balance')" class="min-w-150px cursor-pointer">{{
                  $t('cards.balance')
                }}
              </th>
              <th class="min-w-150px">{{ $t('cards.release') }}</th>
              <th @click="filterCardDate('deactivateDate')" class="min-w-150px cursor-pointer">{{
                  $t('cards.closing')
                }}
              </th>
              <th class="min-w-150px">{{ $t('common.status') }}</th>
              <th class="min-w-100px">{{ $t('common.actions') }}</th>
            </tr>
            </thead>
            <!--end::Table head-->
            <!--begin::Table body-->
            <tbody>
            <tr v-for="card in this.filteredCards" :key="card" :style="{background: background}">
              <td class="align-middle">
                <div class="form-check form-check-sm form-check-custom form-check-solid">
                  <input v-bind:value="card" v-model="checkedCards" class="form-check-input widget-9-check"
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
                    <span class="text-muted fw-semibold text-muted d-block fs-7">{{ $t('cards.date') }}:
                                                {{ card.expirationDate }} CVV: {{ card.CVV }}</span>
                  </div>
                </div>
              </td>
              <td class="fs-6 align-middle">
                <div class="d-flex  align-items-center fs-6">
                  <span class="text-dark fw-bold d-block fs-6">${{ card.limit }}</span>
                </div>
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
                <div v-if="card.deactivateDate" class="d-flex justify-content-start flex-column">
                  <span class="text-dark fw-bold mb-1 fs-6">
                  {{ splitDate(card.deactivateDate)[0] }}
                </span>
                  <span class="text-muted fw-semibold text-muted d-block fs-7">
                  {{ splitDate(card.deactivateDate)[1] }}
                </span>
                </div>
                <div v-else class="d-flex justify-content-start flex-column">
                  <span class="text-dark fw-bold mb-1 fs-6">
                    -
                </span>
                </div>
              </td>
              <td class="text-muted fw-semibold text-hover-primary fs-6 align-middle">
                                    <span v-if="card.status === 0" class="badge badge-light-danger">
                                        {{ $t('cards.blocked') }}
                                    </span>

                <span v-if="card.status === 1" class="badge badge-light-success">
                                        {{ $t('cards.active') }}
                                    </span>
              </td>
              <td class="border-0 align-middle">
                <div class="d-flex justify-content-end flex-shrink-0">
                  <div class="d-flex align-items-center">
                    <!--                  @click='showModal("transactions")'-->
                    <button @click="showDetailsModal(card)"
                            class="btn text-muted btn-bg-light btn-active-color-primary btn-sm me-2"
                            style="white-space: nowrap; ">
                      {{ $t('common.details') }}
                    </button>
                  </div>
                  <div class="d-flex align-items-center">
                    <!--                  @click='showModal("transactions")'-->
                    <button @click="showTopUpModal(card)"
                            class="btn text-muted btn-bg-light btn-active-color-primary btn-sm me-2"
                            style="white-space: nowrap; ">
                      {{ $t('cards.topUp') }}
                    </button>
                  </div>
                </div>
              </td>
            </tr>
            </tbody>
            <!--end::Table body-->
          </table>
          <!--end::Table-->
          <TopUpCard ref="topupcard"/>
          <DetailsCardModal ref="details"/>
          <!--                    <TransactionLogModal ref="transactions" />-->
          <!--                    <CreateCardModal ref="create" />-->
          <!--                    <EditCardModal ref="editcard" />-->
        </div>
        <!--end::Table container-->
      </div>
      <!--begin::Body-->
    </div>
  </div>
</template>

<script>
import {mapGetters} from 'vuex'
import DetailsCardModal from '@/components/modals/detailsCardModal.vue'
import cardIcon from '@/assets/media/credit-card-2-back-fill.svg'
import timezones from '@/timezones.json'
import TopUpCard from "@/components/modals/topUpCard";

export default {
  name: "tableUsers",
  data() {
    return {
      form: {
        balance: 0
      },
      isDisabled: true,
      cardIcon,
      timezones: timezones,
      background: "none",
      check: false,
      isHidden: false,
      showTopup: false,
      checkedCards: [],
      filteredCards: this.cards,
      filterItems: {
        CreatedAt: 1,
        balance: 0,
        spend: 0,
        limit: 0,
        selectedBIN: {name: "", isUniversal: null, active: null},
        types: [
          {
            name: this.$t('cards.universal'),
            code: false
          },
          {
            name: this.$t('cards.advertising'),
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
            name: this.$t('cards.active'),
            code: 1
          },
          {
            name: this.$t('cards.blocked'),
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
  },
  watch: {
    cards: "filter"
  },
  methods: {
    alertCard() {
      this.$swal({
        title: this.$t("cards.alert"),
        type: "alert message",
        icon: "info",
      })
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
    checkForm() {
      console.log(this.form.balance)
    },
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
    cardNumber(number) {
      return number.toString().substring(0, 4) + " " + number.toString().substring(4, 8) + " " + number.toString().substring(8, 12) + " " + number.toString().substring(12, 16)
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
    filter() {
      this.filteredCards = this.cards.filter(card => {
        let searchString = this.filterItems.search.toLowerCase()
        this.filterItems.selectedBIN.isUniversal = this.filterItems.type.code
        let cardBIN = Number(card.cardNumber.toString().split("")[0] + card.cardNumber.toString().split("")[1] + card.cardNumber.toString().split("")[2] + card.cardNumber.toString().split("")[3] + card.cardNumber.toString().split("")[4] + card.cardNumber.toString().split("")[5])
        let BINBoolean = this.BINsList.some(item => {
          if (item.name === cardBIN) {
            return true
          }
        })
        if ((!this.filterItems.status.name || card.status === this.filterItems.status.code)
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

      // фильтрация по балансам
      if (this.filterItems.USDTBalance !== 0) {
        this.filterItems.USDTBalance === 1 ? this.filteredCards.sort(function (a, b) {
          return new Date(b['USDTBalance']) - new Date(a['USDTBalance'])
        }) : this.filteredCards.sort(function (a, b) {
          return new Date(a['USDTBalance']) - new Date(b['USDTBalance'])
        })
      }

      if (this.filterItems.BTCBalance !== 0) {
        this.filterItems.BTCBalance === 1 ? this.filteredCards.sort(function (a, b) {
          return new Date(b['BTCBalance']) - new Date(a['BTCBalance'])
        }) : this.filteredCards.sort(function (a, b) {
          return new Date(a['BTCBalance']) - new Date(b['BTCBalance'])
        })
      }

      if (this.filterItems.USDBalance !== 0) {
        this.filterItems.USDBalance === 1 ? this.filteredCards.sort(function (a, b) {
          return new Date(b['USDBalance']) - new Date(a['USDBalance'])
        }) : this.filteredCards.sort(function (a, b) {
          return new Date(a['USDBalance']) - new Date(b['USDBalance'])
        })
      }

      if (this.filterItems.EURBalance !== 0) {
        this.filterItems.EURBalance === 1 ? this.filteredCards.sort(function (a, b) {
          return new Date(b['EURBalance']) - new Date(a['EURBalance'])
        }) : this.filteredCards.sort(function (a, b) {
          return new Date(a['EURBalance']) - new Date(b['EURBalance'])
        })
      }
    },
    resetFilterCards() {
      this.background = "none"
      this.filterItems = {
        CreatedAt: 1,
        balance: 0,
        spend: 0,
        limit: 0,
        selectedBIN: {name: "", isUniversal: null, active: null},
        types: [
          {
            name: this.$t('cards.universal'),
            code: false
          },
          {
            name: this.$t('cards.advertising'),
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
            name: this.$t('cards.active'),
            code: 1
          },
          {
            name: this.$t('cards.blocked'),
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
      return [this.checkedCards = this.cards, this.check === true]
    },
    showModal(ref) {
      this.$refs[ref].show = true
    },
    showDetailsModal(card) {
      this.$refs.details.openModal(card)
    },
    showTopUpModal(card) {
      this.$refs.topupcard.openModal(card)
    },
    getData() {
      if (this.loggedIn && !this.userInfoLoaded) {
        this.$store.dispatch("users/getUserInfoMethod");
      }
      if (this.loggedIn && !this.cardsLoaded) {
        this.$store.dispatch("users/getCardsMethod");
      }
      if (this.loggedIn && this.suppliers.length === 0) {
        this.$store.dispatch("users/getSuppliersMethod")
      }
      return true;
    },
    cardBalance(balance) {
      if (balance === 0) {
        return "000 000,000$";
      }
    },
    splitDate(date) {
      let userTimeZone = this.user.timeZone
      let timeZone = this.timezones.find(item => item.text === userTimeZone);
      if (date === "0001-01-01T00:00:00Z" || date === "" || date === null || !date || !timeZone) {
        return ["-", "-"]
      }
      if(timeZone && timeZone.utc.includes('-')) {
        let newDate = new Date(date);
        let utc = Number(timeZone.utc.split('-')[1])
        newDate.setTime(newDate.getTime() - (utc * 60 * 60 * 1000))
        console.log(newDate)
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
  components: {
    DetailsCardModal,
    TopUpCard
  },
  computed: {
    ...mapGetters({
      cardsLoaded: "users/getCardsLoaded",
      cards: "users/getCards",
      loggedIn: "auth/isLoggedIn",
      suppliers: "users/getSuppliers",
      suppliersLoaded: "users/getSuppliersLoaded",
      user: "auth/currentUser",
      userinfo: "users/getUserInfo",
      userInfoLoaded: "users/getUserInfoLoaded",
    }),
    BINsList() {
      let BINsList = []
      let initialBINsList = []
      if (this.loggedIn && this.suppliers.length === 0) {
        this.$store.dispatch("users/getSuppliersMethod")
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
  },
}
</script>

<style>
.disabled {
  pointer-events: none;
  cursor: default;
  opacity: 0.5;
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

.table.gs-0 th:first-child,
.table.gs-0 td:first-child {
  padding-left: 1rem !important;
}

.show-delete {
  transition: opacity 0.3s ease !important;
}
</style>
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
        <h1 class="d-flex text-dark fw-bold my-1 fs-3">Logs management</h1>
        <!--end::Title-->
        <!--begin::Breadcrumb-->
        <ul class="breadcrumb breadcrumb-dot fw-semibold text-gray-600 fs-7 my-1">
          <!--begin::Item-->
          <li class="breadcrumb-item text-gray-600">
            <router-link to="/" class="text-gray-600 text-hover-primary">Home</router-link>
          </li>
          <!--end::Item-->
          <!--begin::Item-->
          <li class="breadcrumb-item text-gray-500">Logs management</li>
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
            <span class="user-filter-span">Filter</span>
          </button>
          <!--begin::Menu 1-->
          <div id="filter-modal" v-show="isHidden"
               class="menu menu-sub menu-sub-dropdown show position-fixed float-left w-md-500px transaction-filter-modal" data-kt-menu="true"
               style="width: inherit; z-index: 999; right: 50px;">
            <!--begin::Header-->
            <div class="px-7 py-5">
              <div class="fs-5 text-dark fw-bold">Logs filter</div>
            </div>
            <!--end::Header-->
            <!--begin::Menu separator-->
            <div class="separator border-gray-200"></div>
            <!--end::Menu separator-->
            <!--begin::Form-->
            <div class="px-7 py-5">
              <div class="mb-10">
                <div class="form-group fv-row">
                  <label class="form-label fw-semibold" for="role">Type log:</label>
                  <v-select class="bg-transparent" label="name" :options="filterItems.types"
                            v-model="filterItems.type" required>
                  </v-select>
                </div>
              </div>
              <div class="mb-10">
                <div class="form-group fv-row">
                  <label class="form-label fw-semibold" for="status">User who performed the action:</label>
                  <v-select class="bg-transparent" label="name" :options="usersList"
                            v-model="filterItems.selectedUserWho" required>
                  </v-select>
                </div>
              </div>

              <div class="mb-10">
                <div class="form-group fv-row">
                  <label class="form-label fw-semibold" for="role">User affected by the action:</label>
                  <v-select class="bg-transparent" label="name" :options="usersList"
                            v-model="filterItems.selectedUserWhom" required>
                  </v-select>
                </div>
              </div>

              <div class="mb-10">
                <div class="form-group fv-row">
                  <label class="form-label fw-semibold" for="search">Log search</label>
                  <input class="form-control-solid form-control form-control-lg" label="name"
                         v-model="filterItems.search">
                </div>
              </div>


              <div class="d-flex justify-content-end">
                <button @click="resetFilterLogs" type="reset"
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
      </div>
      <!--end::Actions-->
    </div>
    <div class="card mb-5 mb-xl-8">
      <!--begin::Header-->
      <div class="card-header align-items-center border-0 pt-5">
        <h3 class="card-title align-items-start flex-column">
          <span class="card-label fw-bold fs-3 mb-1">Logs</span>
          <span class="text-muted mt-1 fw-semibold fs-7">Amount logs: {{ this.filteredLogs.length }}</span>
        </h3>

        <!--        <button @click="actionsShow = !actionsShow" type="button"-->
        <!--                class="d-none btn open-actions-card-btn btn-sm btn-icon btn-color-primary btn-active-light-primary"-->
        <!--                data-kt-menu-trigger="click" data-kt-menu-placement="bottom-end">-->
        <!--          &lt;!&ndash;begin::Svg Icon | path: icons/duotune/general/gen024.svg&ndash;&gt;-->
        <!--          <span class="svg-icon svg-icon-2">-->
        <!--                            <svg xmlns="http://www.w3.org/2000/svg" width="24px" height="24px" viewBox="0 0 24 24">-->
        <!--                                <g stroke="none" stroke-width="1" fill="none" fill-rule="evenodd">-->
        <!--                                    <rect x="5" y="5" width="5" height="5" rx="1" fill="currentColor"/>-->
        <!--                                    <rect x="14" y="5" width="5" height="5" rx="1" fill="currentColor" opacity="0.3"/>-->
        <!--                                    <rect x="5" y="14" width="5" height="5" rx="1" fill="currentColor" opacity="0.3"/>-->
        <!--                                    <rect x="14" y="14" width="5" height="5" rx="1" fill="currentColor" opacity="0.3"/>-->
        <!--                                </g>-->
        <!--                            </svg>-->
        <!--                        </span>-->
        <!--          &lt;!&ndash;end::Svg Icon&ndash;&gt;-->
        <!--        </button>-->

        <!--        <div v-if="actionsShow" class="card-toolbar card-toolbar-card">-->
        <!--          <button @click="showBalanceModal()"-->
        <!--                  class="btn show-delete btn-icon btn-bg-light btn-active-color-primary btn-sm">-->
        <!--            &lt;!&ndash;begin::Svg Icon | path: C:/wamp64/www/keenthemes/core/html/src/media/icons/duotune/finance/fin008.svg&ndash;&gt;-->
        <!--            <span class="svg-icon svg-icon-muted svg-icon-2hx" style="margin-right: 0">-->
        <!--                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">-->
        <!--                <path opacity="0.3" d="M3.20001 5.91897L16.9 3.01895C17.4 2.91895 18 3.219 18.1 3.819L19.2 9.01895L3.20001 5.91897Z" fill="currentColor"/>-->
        <!--                <path opacity="0.3" d="M13 13.9189C13 12.2189 14.3 10.9189 16 10.9189H21C21.6 10.9189 22 11.3189 22 11.9189V15.9189C22 16.5189 21.6 16.9189 21 16.9189H16C14.3 16.9189 13 15.6189 13 13.9189ZM16 12.4189C15.2 12.4189 14.5 13.1189 14.5 13.9189C14.5 14.7189 15.2 15.4189 16 15.4189C16.8 15.4189 17.5 14.7189 17.5 13.9189C17.5 13.1189 16.8 12.4189 16 12.4189Z" fill="currentColor"/>-->
        <!--                <path d="M13 13.9189C13 12.2189 14.3 10.9189 16 10.9189H21V7.91895C21 6.81895 20.1 5.91895 19 5.91895H3C2.4 5.91895 2 6.31895 2 6.91895V20.9189C2 21.5189 2.4 21.9189 3 21.9189H19C20.1 21.9189 21 21.0189 21 19.9189V16.9189H16C14.3 16.9189 13 15.6189 13 13.9189Z" fill="currentColor"/>-->
        <!--                </svg>-->
        <!--              </span>-->
        <!--            &lt;!&ndash;end::Svg Icon&ndash;&gt;-->
        <!--          </button>-->
        <!--        </div>-->
      </div>
      <!--end::Header-->
      <!--begin::Body-->
      <div class="card-body py-3">
        <!--begin::Table container-->
        <div class="table-responsive">
          <!--begin::Table-->
          <table v-if="getLogs()" class="table table-row-bordered table-row-gray-100 align-middle gs-0 gy-3">
            <!--begin::Table head-->
            <thead class="rounded-left">
            <tr class="fw-bold text-muted bg-light">
              <th class="min-w-150px">Data created</th>
              <th class="min-w-150px">Type</th>
              <th class="min-w-170px">
                User who performed the action
              </th>
              <th class="min-w-170px">
                User affected by the action
              </th>
              <th class="min-w-100px text-end">Actions</th>
            </tr>
            </thead>
            <!--end::Table head-->
            <!--begin::Table body-->
            <tbody>
            <tr v-for="log in filteredLogs" :key="log" :style="{background: background}">
              <td class="align-middle">
                <div class="text-dark fw-bold fs-6">
                  {{ splitDate(log.CreatedAt) }}
                </div>
              </td>
              <td class="align-middle">
                <span v-if="log.type === 'CardCreation'" class="text-dark fw-bold fs-6">
                  Card creation
                </span>
                <span v-if="log.type === 'CardChange'" class="text-dark fw-bold fs-6">
                  Card change
                </span>
                <span v-if="log.type === 'CardReplenishment'" class="text-dark fw-bold fs-6">
                  Card replenishment
                </span>
                <span v-if="log.type === 'AccountReplenishment'" class="text-dark fw-bold fs-6">
                  Account replenishment
                </span>
              </td>
              <td class="align-middle">
                <div class="d-flex align-items-center">
                  <div class="symbol symbol-50px me-5">
                                            <span class="symbol-label bg-light">
                                                <img src="../assets/media/001-boy.svg" class="h-75 align-self-end"
                                                     alt=""/>
                                            </span>
                  </div>
                  <div class="d-flex justify-content-start flex-column">
                    <span class="text-dark fw-bold text-hover-primary fs-6 cursor-pointer text-hover-primary mb-1" style="white-space: nowrap">
                      {{getUserInfo(log.creatorID)[1] }}, ID: {{getUserInfo(log.creatorID)[0]}}</span>
                    <a v-bind:href="'mailto:' + getUserInfo(log.createdID)[2]"
                       class="text-muted fw-semibold text-muted d-block fs-7">{{ getUserInfo(log.creatorID)[2] }}</a>
                  </div>
                </div>
              </td>
              <td class="align-middle">
                <div class="d-flex align-items-center">
                  <div class="symbol symbol-50px me-5">
                                            <span class="symbol-label bg-light">
                                                <img src="../assets/media/001-boy.svg" class="h-75 align-self-end"
                                                     alt=""/>
                                            </span>
                  </div>
                  <div class="d-flex justify-content-start flex-column">
                    <span class="text-dark fw-bold text-hover-primary fs-6 cursor-pointer text-hover-primary mb-1" style="white-space: nowrap">
                      {{getUserInfo(log.userID)[1] }}, ID: {{getUserInfo(log.userID)[0]}}</span>
                    <a v-bind:href="'mailto:' + getUserInfo(log.userID)[2]"
                       class="text-muted fw-semibold text-muted d-block fs-7">{{ getUserInfo(log.userID)[2] }}</a>
                  </div>
                </div>
              </td>
              <td class="border-0 end-buttons align-middle">
                <div class="d-flex align-items-center justify-content-end">
                  <button v-if="log.type === 'CardCreation'" @click="showCardCreationModal(log.ID, log.type)"
                          class="btn text-muted btn-bg-light btn-active-color-primary btn-sm me-2"
                          style="white-space: nowrap; ">
                    Details
                  </button>
                  <button v-if="log.type === 'AccountReplenishment'" @click="showAccountReplenishmentModal(log.ID, log.type)"
                          class="btn text-muted btn-bg-light btn-active-color-primary btn-sm me-2"
                          style="white-space: nowrap; ">
                    Details
                  </button>
                  <button v-if="log.type === 'CardChanges'" @click="showCardChangeModal(log.ID, log.type)"
                          class="btn text-muted btn-bg-light btn-active-color-primary btn-sm me-2"
                          style="white-space: nowrap; ">
                    Details
                  </button>
                  <button v-if="log.type === 'CardReplenishment'" @click="showCardReplenishmentModal(log.ID, log.type)"
                          class="btn text-muted btn-bg-light btn-active-color-primary btn-sm me-2"
                          style="white-space: nowrap; ">
                    Details
                  </button>
                </div>
              </td>
            </tr>
            </tbody>
            <!--end::Table body-->
          </table>
          <!--end::Table-->
        </div>
        <!--end::Table container-->
        <logCardCreation ref="creation"/>
        <logAccountReplenishment ref="account"/>
        <logCardReplenishment ref="replenishment"/>
        <logCardChange ref="change"/>
      </div>
      <!--begin::Body-->
    </div>
  </div>
</template>

<script>
import {mapGetters, mapActions} from 'vuex'
import countries from '../countries.json'
import timezones from '../timezones.json'
import logCardCreation from './modals/LogCardCreation'
import logAccountReplenishment from './modals/LogAccountReplenishment'
import logCardReplenishment from './modals/LogCardReplenishment'
import logCardChange from './modals/LogCardChange'
// import apiV1 from "@/axios";

export default {
  name: 'tableLogs',
  data() {
    return {
      timeZone: {
        id: null,
        text: null,
        utc: 0
      },
      timezones: timezones,
      background: "none",
      onTheTable: false,
      check: false,
      actionsShow: false,
      userEmailWhom: null,
      userNameWhom: null,
      userIdWhom: null,
      userEmailWho: null,
      userNameWho: null,
      userIdWho: null,
      checkedLogs: [],
      isHidden: false,
      copyNick: false,
      req: true,
      filteredLogs: this.finLogs,
      filterItems: {
        selectedUserWho: {code: 0, name: "", user: {}},
        selectedUserWhom: {code: 0, name: "", user: {}},
        CreatedAt: 1,
        USDTBalance: 0,
        BTCBalance: 0,
        USDBalance: 0,
        EURBalance: 0,
        countries: countries,
        types: [
          {
            name: "Card сhanges",
            code: "CardChanges"
          },
          {
            name: "Card replenishment",
            code: "CardReplenishment"
          },
          {
            name: "Card creation",
            code: "CardCreation"
          },
          {
            name: "Account replenishment",
            code: "AccountReplenishment"
          },
        ],
        type: {
          name: null,
          code: null
        },
        search: '',
      },
    }
  },
  created() {
    this.filteredLogs = this.finLogs
    if (this.finLogs) {
      this.filter()
    }
  },
  watch: {
    finLogs: "filter"
  },
  methods: {
    ...mapActions({
      login: "auth/login",
      getAccess: "admin/getAccess"
    }),
    filterUserDate(date) {
      this.filterItems[date] !== 0 ? this.filterItems[date] = -this.filterItems[date] : this.filterItems[date] = 1
      date === "CreatedAt" ? this.filterItems.deactivateDate = 0 : this.filterItems.CreatedAt = 0
      this.filter()
    },
    filter() {
      this.filteredLogs = this.finLogs.filter(log => {
        let searchString = this.filterItems.search.toLowerCase()
        if ((!this.filterItems.type.name || log.type === this.filterItems.type.code)
            && (!this.filterItems.selectedUserWho.code || log.creatorID === this.filterItems.selectedUserWho.code)
            && (!this.filterItems.selectedUserWhom.code || log.userID === this.filterItems.selectedUserWhom.code)
            && (log.type.toLowerCase().indexOf(searchString) !== -1)) {
          return log
        }
      })
      if(this.filterItems.search && this.filteredLogs.length > 1) {
        this.background = "rgb(230, 255, 241)"
      } else {
        this.background = "none"
      }
      // фильтрация по дате
      if (this.filterItems.CreatedAt !== 0) {
        this.filterItems.CreatedAt === 1 ? this.filteredLogs.sort(function (a, b) {
          return new Date(b['CreatedAt']) - new Date(a['CreatedAt'])
        }) : this.filteredLogs.sort(function (a, b) {
          return new Date(a['CreatedAt']) - new Date(b['CreatedAt'])
        })
      }
    },
    resetFilterLogs() {
      this.filterItems = {
        selectedUserWho: {code: 0, name: "", user: {}},
        selectedUserWhom: {code: 0, name: "", user: {}},
        CreatedAt: 1,
        USDTBalance: 0,
        BTCBalance: 0,
        USDBalance: 0,
        EURBalance: 0,
        countries: countries,
        types: [
          {
            name: "Card сhanges",
            code: "CardChanges"
          },
          {
            name: "Card replenishment",
            code: "CardReplenishment"
          },
          {
            name: "Card creation",
            code: "CardCreation"
          },
          {
            name: "Account replenishment",
            code: "AccountReplenishment"
          },
        ],
        type: {
          name: null,
          code: null
        },
        search: '',
      }
      this.filter()
    },
    showModal() {
      this.$refs.modal.openModal()
    },
    showCardCreationModal(actionID, type) {
      this.$refs.creation.openModal(actionID, type)
    },
    showAccountReplenishmentModal(actionID, type) {
      this.$refs.account.openModal(actionID, type)
    },
    showCardReplenishmentModal(actionID, type) {
      this.$refs.replenishment.openModal(actionID, type)
    },
    showCardChangeModal(actionID, type) {
      this.$refs.change.openModal(actionID, type)
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
    getLogs() {
      if (this.loggedIn && this.finLogs.length === 0 && this.req) {
        this.$store.dispatch("admin/getAllFinLogsMethod")
      }
      if (this.loggedIn && this.users.length === 0 && this.req) {
        this.$store.dispatch("admin/getUsersMethod")
      }
      this.req = false
      return true
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
  },
  computed: {
    ...mapGetters({
      users: "admin/getUsers",
      finLogs: "admin/getAllFinLogs",
      currentUser: "auth/currentUser",
      loggedIn: "auth/isLoggedIn",
      getTokenHeader: "auth/getTokenHeader",
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
  },
  components: {
    logCardCreation,
    logAccountReplenishment,
    logCardReplenishment,
    logCardChange
  }
}
</script>

<style scoped>
.table.gs-0 td:first-child {
  padding-left: 1rem !important;
}

.table th:last-child {
  padding: 0.75rem 0.75rem !important;
}

.show-delete {
  transition: opacity 0.2s ease !important;
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
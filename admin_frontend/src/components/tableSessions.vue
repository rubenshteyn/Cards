<template>
  <div v-show="isHidden" @click="isHidden = false" class="position-absolute vh-100 w-100" style="z-index: 99;left: 0px;top: 0px;"></div>
  <div class="d-flex flex-column flex-column-fluid container-fluid" >
    <div class="toolbar mb-5 mb-lg-7" id="kt_toolbar">
      <!--begin::Page title-->
      <div class="page-title d-flex flex-column me-3">
        <!--begin::Title-->
        <h1 class="d-flex text-dark fw-bold my-1 fs-3">Session management</h1>
        <!--end::Title-->
        <!--begin::Breadcrumb-->
        <ul class="breadcrumb breadcrumb-dot fw-semibold text-gray-600 fs-7 my-1">
          <!--begin::Item-->
          <li class="breadcrumb-item text-gray-600">
            <router-link to="/" class="text-gray-600 text-hover-primary">Home</router-link>
          </li>
          <!--end::Item-->
          <!--begin::Item-->
          <li class="breadcrumb-item text-gray-500">Session management</li>
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
                                    fill="currentColor" />
                            </svg>
                        </span>
            <!--end::Svg Icon-->
            Filter
          </button>
          <!--begin::Menu 1-->
          <div id="filter-modal" v-show="isHidden" class="menu menu-sub menu-sub-dropdown show position-fixed float-left" data-kt-menu="true" style="width: inherit; z-index: 999">
            <!--begin::Header-->
            <div class="px-7 py-5">
              <div class="fs-5 text-dark fw-bold">Session filter</div>
            </div>
            <!--end::Header-->
            <!--begin::Menu separator-->
            <div class="separator border-gray-200"></div>
            <!--end::Menu separator-->
            <!--begin::Form-->
            <div class="px-7 py-5">
              <div class="mb-10">
                <div class="form-group fv-row">
                  <label class="form-label fw-semibold" for="status">Session status:</label>
                  <v-select class="bg-transparent" label="name" :options="filterItems.statuses"
                            placeholder="Select session status" v-model="filterItems.status" required>
                  </v-select>
                </div>
              </div>

              <div class="mb-10">
                <div class="form-group fv-row">
                  <label class="form-label fw-semibold" for="search">User search</label>
                  <input class="form-control form-control-solid form-control-lg" label="name" v-model="filterItems.search">
                </div>
              </div>


              <div class="d-flex justify-content-end">
                <button @click="resetFilterSessions" type="reset" class="btn btn-sm btn-light btn-active-light-primary me-2">Reset</button>
                <button @click="filter" type="submit" class="btn btn-sm btn-primary"
                        data-kt-menu-dismiss="true">Apply</button>
              </div>
              <!--end::Actions-->
            </div>
            <!--end::Form-->
          </div>
          <!--end::Menu 1-->
          <!--end::Menu-->
        </div>
        <!--end::Wrapper-->

        <button @click='deleteAllSessions' class="btn btn-dark fw-bold">Delete all sessions</button>
      </div>
      <!--end::Actions-->
    </div>
    <div class="card mb-5 mb-xl-8">
      <!--begin::Header-->
      <div class="card-header border-0 pt-5">
        <h3 class="card-title align-items-start flex-column">
          <span class="card-label fw-bold fs-3 mb-1">Sessions</span>
          <span class="text-muted mt-1 fw-semibold fs-7">Amount sessions: {{this.filteredSessions.length}}</span>
        </h3>
        <div class="card-toolbar">
          <!--begin::Menu-->
          <button type="button" class="btn btn-sm btn-icon btn-color-primary btn-active-light-primary"
                  data-kt-menu-trigger="click" data-kt-menu-placement="bottom-end">
            <!--begin::Svg Icon | path: icons/duotune/general/gen024.svg-->
            <span class="svg-icon svg-icon-2">
                            <svg xmlns="http://www.w3.org/2000/svg" width="24px" height="24px" viewBox="0 0 24 24">
                                <g stroke="none" stroke-width="1" fill="none" fill-rule="evenodd">
                                    <rect x="5" y="5" width="5" height="5" rx="1" fill="currentColor" />
                                    <rect x="14" y="5" width="5" height="5" rx="1" fill="currentColor" opacity="0.3" />
                                    <rect x="5" y="14" width="5" height="5" rx="1" fill="currentColor" opacity="0.3" />
                                    <rect x="14" y="14" width="5" height="5" rx="1" fill="currentColor" opacity="0.3" />
                                </g>
                            </svg>
                        </span>
            <!--end::Svg Icon-->
          </button>
          <!--begin::Menu 2-->
          <div class="menu menu-sub menu-sub-dropdown menu-column menu-rounded menu-gray-800 menu-state-bg-light-primary fw-semibold w-200px"
               data-kt-menu="true">
            <!--begin::Menu item-->
            <div class="menu-item px-3">
              <div class="menu-content fs-6 text-dark fw-bold px-3 py-4">Quick Actions</div>
            </div>
            <!--end::Menu item-->
            <!--begin::Menu separator-->
            <div class="separator mb-3 opacity-75"></div>
            <!--end::Menu separator-->
            <!--begin::Menu item-->
            <div class="menu-item px-3">
              <a href="#" class="menu-link px-3">New Ticket</a>
            </div>
            <!--end::Menu item-->
            <!--begin::Menu item-->
            <div class="menu-item px-3">
              <a href="#" class="menu-link px-3">New Customer</a>
            </div>
            <!--end::Menu item-->
            <!--begin::Menu item-->
            <div class="menu-item px-3" data-kt-menu-trigger="hover" data-kt-menu-placement="right-start">
              <!--begin::Menu item-->
              <a href="#" class="menu-link px-3">
                <span class="menu-title">New Group</span>
                <span class="menu-arrow"></span>
              </a>
              <!--end::Menu item-->
              <!--begin::Menu sub-->
              <div class="menu-sub menu-sub-dropdown w-175px py-4">
                <!--begin::Menu item-->
                <div class="menu-item px-3">
                  <a href="#" class="menu-link px-3">Admin Group</a>
                </div>
                <!--end::Menu item-->
                <!--begin::Menu item-->
                <div class="menu-item px-3">
                  <a href="#" class="menu-link px-3">Staff Group</a>
                </div>
                <!--end::Menu item-->
                <!--begin::Menu item-->
                <div class="menu-item px-3">
                  <a href="#" class="menu-link px-3">Member Group</a>
                </div>
                <!--end::Menu item-->
              </div>
              <!--end::Menu sub-->
            </div>
            <!--end::Menu item-->
            <!--begin::Menu item-->
            <div class="menu-item px-3">
              <a href="#" class="menu-link px-3">New Contact</a>
            </div>
            <!--end::Menu item-->
            <!--begin::Menu separator-->
            <div class="separator mt-3 opacity-75"></div>
            <!--end::Menu separator-->
            <!--begin::Menu item-->
            <div class="menu-item px-3">
              <div class="menu-content px-3 py-3">
                <a class="btn btn-primary btn-sm px-4" href="#">Generate Reports</a>
              </div>
            </div>
            <!--end::Menu item-->
          </div>
          <!--end::Menu 2-->
          <!--end::Menu-->
        </div>
      </div>
      <!--end::Header-->
      <!--begin::Body-->
      <div class="card-body py-3">
        <!--begin::Table container-->
        <div class="table-responsive">
          <!--begin::Table-->
          <table v-if="getSessions()" class="table table-row-bordered table-row-gray-100 align-middle gs-0 gy-3">
            <!--begin::Table head-->
            <thead class="rounded-left">
            <tr class="fw-bold text-muted text-muted bg-light">
              <th @click="filterSessionDate('CreatedAt')" class="min-w-140px cursor-pointer" style="white-space: nowrap">Connection time</th>
              <th class="min-w-150px">Location</th>
              <th class="min-w-130px">Status</th>
              <th @click="filterSessionDate('DeletedAt')" class="min-w-140px cursor-pointer" style="white-space: nowrap">Completion time</th>
              <th class="min-w-150px text-end">Actions</th>
            </tr>
            </thead>
            <!--end::Table head-->
            <!--begin::Table body-->
            <tbody>
            <tr v-for="session in filteredSessions" :key="session" :style="{ background: ifCurrentSession(session) }">
              <td class="align-middle">
                <span class="text-dark fw-bold d-block fs-6">{{ splitDate(session.CreatedAt) }}</span>
              </td>

              <td class="text-muted text-hover-primary fs-6 align-middle">
                <Popper class="fw-semibold" hover placement="top" :content="splitLocation(session.location)[1]">
                    <span class="d-block fs-6">
                      {{ splitLocation(session.location)[0] }}
                  </span>
                </Popper>
              </td>

              <td class="align-middle">
                <span v-if="session.status === 0 || session.DeletedAt" class="badge badge-light-danger">
                    Not active
                </span>

                <span v-else-if="session.status === 1" class="badge badge-light-success">
                                          Active
                                      </span>

                <span v-else-if="session.status === 2" class="badge badge-light-success">
                                          Active from admin
                                      </span>
                <span v-else-if="session.status === 3" class="badge badge-light-danger">
                                          Expired
                                      </span>
                <span v-else-if="session.status === 4" class="badge badge-light-warning">
                                          Login from admin
                                      </span>
              </td>

              <td class="text-muted fw-semibold fs-6 align-middle">
                  <span class="d-block fs-6">
                    {{ splitDate(session.DeletedAt) }}
                  </span>
              </td>

              <td class="border-0 align-middle">
                <div class="d-flex justify-content-end">
                  <button @click="deleteSession(session)" :disabled='isDisabled(session)[0]' class="btn badge badge-light-danger btn-sm px-4">
                    {{ isDisabled(session)[1] }}
                  </button>
                </div>
              </td>
            </tr>
            </tbody>
            <!--end::Table body-->
          </table>
          <EditUserModal ref="edit" />
          <!--end::Table-->
        </div>
        <!--end::Table container-->
      </div>
      <!--begin::Body-->
    </div>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'
import countries from '../countries.json'
import EditUserModal from '@/components/modals/editUserModal.vue'
import Popper from "vue3-popper";
import timezones from "@/timezones.json";

export default {
  name: "tableSessions",
  data() {
    return {
      isHidden: false,
      content: null,
      countries: countries,
      timezones: timezones,
      req: true,
      styles: {
        active: ""
      },
      filterItems: {
        CreatedAt: 1,
        DeletedAt: 0,
        statuses: [
          {
            name: "Not active",
            code: 0
          },
          {
            name: "Active",
            code: 1
          },
          {
            name: "Active from admin",
            code: 2
          },
          {
            name: "Expired",
            code: 3
          },
          {
            name: "Login from admin",
            code: 4
          },
        ],
        status: {
          name: null,
          code: null
        },
        search: '',
      },
      filteredSessions: this.sessions
    }
  },
  watch: {
    sessions: "filter"
  },
  created() {
    this.filteredSessions = this.sessions
    if (this.sessions) {
      this.filter()
    }
  },
  methods: {
    ...mapActions({
      deleteSessionMethod: "users/deleteSessionMethod",
    }),
    filterSessionDate(date) {
      this.filterItems[date] !== 0 ? this.filterItems[date] = -this.filterItems[date] : this.filterItems[date] = 1
      date === "CreatedAt" ? this.filterItems.DeletedAt = 0 : this.filterItems.CreatedAt = 0
      this.filter()
    },
    filter() {
      this.filteredSessions = this.sessions.filter(session => {
        let searchString = this.filterItems.search.toLowerCase()
        if((!this.filterItems.status.name
            || (session.status === this.filterItems.status.code && !session.DeletedAt)
            || (this.filterItems.status.code === 0 && session.DeletedAt))
            && session.location
            && (!searchString || session.location.indexOf(searchString) !== -1)) {
          return session
        }
      })
      if (this.filterItems.CreatedAt !== 0) {
        this.filterItems.CreatedAt === 1 ? this.filteredSessions.sort(function (a, b) {
          return new Date(b['CreatedAt']) - new Date(a['CreatedAt'])
        }) : this.filteredSessions.sort(function (a, b) {
          return new Date(a['CreatedAt']) - new Date(b['CreatedAt'])
        })
      }
      if (this.filterItems.DeletedAt !== 0) {
        this.filterItems.DeletedAt === 1 ? this.filteredSessions.sort(function (a, b) {
          return new Date(b['DeletedAt']) - new Date(a['DeletedAt'])
        }) : this.filteredSessions.sort(function (a, b) {
          return new Date(a['DeletedAt']) - new Date(b['DeletedAt'])
        })
      }
    },
    resetFilterSessions() {
      this.filterItems = {
        CreatedAt: 1,
        DeletedAt: 0,
        statuses: [
          {
            name: "Not active",
            code: 0
          },
          {
            name: "Active",
            code: 1
          },
          {
            name: "Active from admin",
            code: 2
          },
          {
            name: "Expired",
            code: 3
          },
          {
            name: "Login from admin",
            code: 4
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
    showModal() {
      this.$refs.modal.show = true
    },
    showEditModal(user) {
      this.$refs.edit.openModal(user)
    },
    getSessions() {
      if (this.loggedIn && !this.sessionsLoaded && this.req) {
        this.$store.dispatch("users/getSessionsMethod")
      }
      this.req = false
      return true
    },
    deleteSession(session) {
      let payload = {
        sessionID: session.ID,
        deleteAll: false
      }
      this.deleteSessionMethod(payload)
          .catch(error => {
            console.log(error)
          });
      if (session.authenticationToken === this.user.authenticationToken) {
        this.$store.dispatch("auth/logout")
      }
    },
    deleteAllSessions() {
      let sessions = Object.values(this.sessions);
      let session = sessions.find(item => item.authenticationToken === this.user.authenticationToken);
      let payload = {
        sessionID: session.ID,
        deleteAll: true
      }
      this.deleteSessionMethod(payload)
          .catch(error => {
            console.log(error)
          });
    },
    userBalance(balance) {
      if (balance === 0) {
        return '0000.00$'
      }
    },
    isDisabled(session) {
      let status = "End session"
      if(session.status === 0 || session.DeletedAt) {
        status = "Session is inactive"
        return [true, status]
      }
      return [false, status]
    },
    splitLocation(location) {
      if (!location) {
        return "-"
      }
      let locationList = location.split(' ')
      let locationBrowser = locationList.slice(1, -1)
      locationList[1] = locationBrowser.join(' ')
      return locationList
    },
    splitDate(date) {
      if (date === "0001-01-01T00:00:00Z" || !date) {
        return "-"
      }
      let userTimeZone = this.user.timeZone
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
    ifCurrentSession(session) {
      if (session.authenticationToken === this.user.authenticationToken) {
        return "#e6fff1"
      }
      return ""
    },
  },
  computed: {
    ...mapGetters({
      sessionsLoaded: "users/getSessionsLoaded",
      sessions: "users/getSessions",
      loggedIn: "auth/isLoggedIn",
      user: "auth/currentUser"
    })
  },
  components: {
    EditUserModal,
    Popper
  }
}
</script>
<style scoped>
:deep(.popper) {
  background: #A1A5B7;
  padding: 10px;
  border-radius: 5px;
  color: #fff;
  text-transform: uppercase;
}

:deep(.popper #arrow::before) {
  background: #9397a7;
}

:deep(.popper:hover),
:deep(.popper:hover > #arrow::before) {
  background: #9397a7;
}
/*.table.gs-0 td:first-child {*/
/*  padding-left: 1rem !important;*/
/*}*/
.table th:last-child{
  padding: 0.75rem 0.75rem !important;
}
</style>

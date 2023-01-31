<template>
  <div v-show="isHidden" @click="isHidden = false" class="position-absolute vh-100 w-100"
       style="z-index: 99;left: 0px;top: 0px;"></div>
  <div v-show="actionsShow" @click="actionsShow = false" class="position-absolute vh-100 w-100"
       style="z-index: 99;left: 0px;top: 0px;"></div>
  <div v-show="this.limitID" @click="this.limitID = 0" class="position-absolute vh-100 w-100"
       style="z-index: 99;left: 0px;top: 0px;"></div>
  <div class="d-flex flex-column user-tasks flex-column-fluid container-fluid">
    <div class="toolbar mb-5 mb-lg-7" id="kt_toolbar">
      <!--begin::Page title-->
      <div class="page-title d-flex flex-column me-3">
        <!--begin::Title-->
        <h1 class="d-flex text-dark fw-bold my-1 fs-3">Task list management</h1>
        <!--end::Title-->
        <!--begin::Breadcrumb-->
        <ul class="breadcrumb breadcrumb-dot fw-semibold text-gray-600 fs-7 my-1">
          <!--begin::Item-->
          <li class="breadcrumb-item text-gray-600">
            <router-link to="/" class="text-gray-600 text-hover-primary">Home</router-link>
          </li>
          <!--end::Item-->
          <!--begin::Item-->
          <li class="breadcrumb-item text-gray-500">Task list management</li>
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
              <div class="fs-5 text-dark fw-bold">task filter</div>
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
                  <label class="form-label fw-semibold" for="status">Task status:</label>
                  <v-select class="bg-transparent" label="name" :options="filterItems.statuses"
                            placeholder="Select task status" v-model="filterItems.status" required>
                  </v-select>
                </div>
              </div>

              <div class="mb-10">
                <div class="form-group fv-row">
                  <label class="form-label fw-semibold" for="status">Task type:</label>
                  <v-select class="bg-transparent" label="name" :options="filterItems.actionTypes"
                            placeholder="Select task type" v-model="filterItems.actionType" required>
                  </v-select>
                </div>
              </div>

              <div class="mb-10">
                <div class="form-group fv-row">
                  <label class="form-label fw-semibold" for="search">Task search</label>
                  <input class="form-control form-control-lg form-control-solid" label="name"
                         v-model="filterItems.search">
                </div>
              </div>

              <div class="d-flex justify-content-end">
                <button @click="resetFilterTasks" type="reset"
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
        <button @click='showModal' class="btn btn-dark fw-bold" data-bs-toggle="modal"
                data-bs-target="#kt_modal_create_app"
                id="kt_toolbar_primary_button">Create task
        </button>
        <!--end::Button-->
      </div>
      <!--end::Actions-->
    </div>
    <div class="task mb-5 mb-xl-8">
      <!--begin::Header-->
      <div class="task-header align-items-center task-header-task border-0 pt-5">
        <h3 class="task-title align-items-start flex-column">
          <span class="d-block task-label fw-bold fs-3 mb-1">All tasks</span>
          <span class="d-block text-muted mt-1 fw-semibold fs-7">Amount tasks: {{ this.filteredTasks.length }} </span>
        </h3>
      </div>
      <!--end::Header-->
      <!--begin::Body-->
      <div class="task-body py-3">
        <!--begin::Table container-->
        <div class="table-responsive">
          <!--begin::Table-->
          <table v-if="getTasks()" class="table table-row-bordered table-row-gray-100 align-middle gs-0 gy-3">
            <!--begin::Table head-->
            <thead class="rounded-left">
            <tr class="fw-bold text-muted bg-light">
              <th class="w-25px">
                <div class="form-check form-check-sm form-check-custom form-check-solid mr-3">
                  <input @click="selectAllTasks" v-model="check" class="form-check-input" type="checkbox"/>
                </div>
              </th>
              <th @click="filterTaskDate('CreatedAt')" class="min-w-150px cursor-pointer">Date of creation</th>
              <th class="min-w-150px">Client</th>
              <th class="min-w-150px">Type of action</th>
              <th @click="filterTaskBalance('balance')" class="min-w-150px cursor-pointer">Client balance before</th>
              <th @click="filterTaskBalance('balance')" class="min-w-150px cursor-pointer">Operation cost</th>
              <th class="min-w-150px">Status</th>
              <th @click="filtertaskDate('deactivateDate')" class="min-w-150px cursor-pointer">Closing date</th>
              <th class="min-w-100px">Actions</th>
            </tr>
            </thead>
            <!--end::Table head-->
            <!--begin::Table body-->
            <tbody>
            <tr v-for="task in this.filteredTasks" :key="task" :style="{background: background}">
              <td class="align-middle">
                <div class="form-check form-check-sm form-check-custom form-check-solid">
                  <input v-bind:value="task" v-model="checkedTasks"
                         class="form-check-input widget-9-check"
                         type="checkbox"/>
                </div>
              </td>
              <td class="align-middle">
                <div class="d-flex justify-content-start flex-column">
                  <span class="text-dark fw-bold mb-1 fs-6">
                  {{ splitDate(task.CreatedAt)[0] }}
                </span>
                  <span class="text-muted fw-semibold text-muted d-block fs-7">
                  {{ splitDate(task.CreatedAt)[1] }}
                </span>
                </div>
              </td>
              <td class="align-middle">
                <div class="d-flex flex-column">
                  <span class="text-dark fw-bold mb-1 fs-6">{{ getUserInfo(task.userID)[0] }}</span>
                  <span class="text-muted fw-semibold text-muted d-block fs-7">{{ getUserInfo(task.userID)[1] }}</span>
                </div>
              </td>
              <td class="fs-6 align-middle position-relative">
                <span v-if="task.type === 'CardCreation'" class="text-dark fw-bold fs-6">
                  Card creation
                </span>
                <span v-if="task.type === 'CardReplenishment'" class="text-dark fw-bold fs-6">
                  Card replenishment
                </span>
                <span v-if="task.type === 'AccountReplenishment'" class="text-dark fw-bold fs-6">
                  Account replenishment
                </span>
              </td>
              <td class="align-middle">
                <div class="text-dark fw-bold d-block fs-6">{{ userBalance(task.initialBalance) }}₮</div>
              </td>
              <td class="align-middle">
                <div class="text-dark fw-bold d-block fs-6">{{ task.price }}</div>
              </td>
              <td class="text-muted fw-semibold text-hover-primary fs-6 align-middle">
                <span v-if="task.status === 0" class="badge badge-light-warning">
                  Rejected
                </span>
                <span v-if="task.status === 1" class="badge badge-light-primary">
                  Waiting
                </span>
                <span v-if="task.status === 2" class="badge badge-light-danger">
                  Canceled
                </span>
                <span v-if="task.status === 3" class="badge badge-light-success">
                  Completed
                </span>
              </td>
              <td class="align-middle">
                <span v-if="task.status === 1" class="badge badge-light-success"
                      style="color: rgb(84 84 84)  !important; background-color: #DBDBE0FF !important;"> n/a </span>
                <div v-else class="d-flex justify-content-start flex-column">
                  <span class="text-dark fw-bold fs-6">
                  {{ splitDate(task.deactivateDate)[0] }}
                </span>
                  <span class="text-muted fw-semibold text-muted d-block fs-7">
                {{ splitDate(task.deactivateDate)[1] }}
                </span>
                </div>
              </td>
              <td class="border-0 align-middle">
                <div class="d-flex align-items-center">
                  <div v-if="task.status === 1" class="me-2">
                    <Popper class="fw-semibold fs-7" arrow hover placement="top" content="Click to complete task">
                      <button @click='completeTask(task.id)'
                              class="btn text-muted btn-bg-light btn-active-color-primary btn-sm"
                              style="white-space: nowrap; ">
                        Complete task
                      </button>
                    </Popper>
                  </div>
                  <div v-if="task.status === 1" class="me-2">
                    <Popper class="fw-semibold fs-7" arrow hover placement="top" content="Click to cansel task">
                      <button @click='cancelTask(task.id)'
                              class="btn text-muted btn-bg-light btn-active-color-primary btn-sm"
                              style="white-space: nowrap; ">
                        Cancel task
                      </button>
                    </Popper>
                  </div>
                  <div v-if="task.status === 2" class="me-2">
                    <Popper class="fw-semibold fs-7" arrow hover placement="top" content="Click to repeat task">
                      <button @click='repeatTask(task.id)'
                              class="btn text-muted btn-bg-light btn-active-color-primary btn-sm"
                              style="white-space: nowrap; ">
                        Repeat task
                      </button>
                    </Popper>
                  </div>
                </div>
              </td>
            </tr>
            </tbody>
            <!--end::Table body-->
          </table>
          <!--end::Table-->
        </div>
        <!--end::Table container-->
      </div>
      <!--begin::Body-->
    </div>
    <TaskToCreateCard ref="create" />
  </div>
</template>

<script>
import {mapActions, mapGetters} from 'vuex'
import timezones from "@/timezones.json";
import TaskToCreateCard from "@/components/modals/createTaskModal";

export default {
  name: "taskList",
  data() {
    return {
      background: "none",
      timezones: timezones,
      req: true,
      check: false,
      isHidden: false,
      actionsShow: false,
      checkedTasks: [],
      filteredTasks: this.tasks,
      limitID: 0,
      filterItems: {
        selectedUser: {code: 0, name: "", user: {}},
        CreatedAt: 1,
        deactivateDate: 0,
        balance: 0,
        actionTypes: [
          {
            name: "Card creation",
            code: "CardCreation"
          },
          {
            name: "Card replenishment",
            code: "CardReplenishment"
          },
          {
            name: "Account replenishment",
            code: "AccountReplenishment"
          }
        ],
        actionType: {
          name: null,
          code: null
        },
        statuses: [
          {
            name: "Rejected",
            code: 0
          },
          {
            name: "Waiting",
            code: 1
          },
          {
            name: "Canceled",
            code: 2
          },
          {
            name: "Сompleted",
            code: 3
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
    this.filteredTasks = this.tasks
    if (this.tasks) {
      this.filter()
    }
  },
  watch: {
    tasks: "filter"
  },
  methods: {
    ...mapActions({
      completeTask: "admin/completeTask",
      repeatTask: "admin/repeatTask"
    }),
    showModal() {
      this.$refs.create.openModal()
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
    completeTask(id) {
      this.completeTask(id)
    },
    cancelTask(id) {
      this.completeTask(id)
    },
    repeatTask(id) {
      this.repeatTask(id)
    },
    checkForm: function () {
      console.log("test")
    },
    filterTaskDate(date) {
      this.filterItems[date] !== 0 ? this.filterItems[date] = -this.filterItems[date] : this.filterItems[date] = 1
      date === "CreatedAt" ? this.filterItems.deactivateDate = 0 : this.filterItems.CreatedAt = 0
      this.filter()
    },
    filterTaskBalance(balance) {
      let stat = 0
      this.filterItems[balance] !== 0 ? stat = -this.filterItems[balance] : stat = 1
      this.filterItems.balance = 0
      this.filterItems[balance] = stat
      this.filter()
    },
    checkingValues(taskElement) {
      if (taskElement === "" || typeof taskElement === undefined || taskElement === null) {
        return taskElement === "-"
      }
      if (taskElement === 0) {
        return taskElement === 0
      }
      return taskElement
    },
    filter() {
      this.filteredTasks = this.tasks.filter(task => {
        let taskInCycle = Object.values(task)
        for (const taskElement of taskInCycle) {
          this.checkingValues(taskElement)
        }
        // let searchString = this.filterItems.search.toLowerCase()
        if ((!this.filterItems.status.name || task.status === this.filterItems.status.code)
            && (!this.filterItems.selectedUser.code || task.userId === this.filterItems.selectedUser.code)
            && (!this.filterItems.actionType.code || task.type === this.filterItems.actionType.code)) {
          return task
        }
      })
      if (this.filterItems.search && this.filteredTasks.length > 1) {
        this.background = "rgb(230, 255, 241)"
      } else {
        this.background = "none"
      }
      // фильтрация по дате
      if (this.filterItems.CreatedAt !== 0) {
        this.filterItems.CreatedAt === 1 ? this.filteredTasks.sort(function (a, b) {
          return new Date(b['CreatedAt']) - new Date(a['CreatedAt'])
        }) : this.filteredTasks.sort(function (a, b) {
          return new Date(a['CreatedAt']) - new Date(b['CreatedAt'])
        })
      }

      if (this.filterItems.deactivateDate !== 0) {
        this.filterItems.deactivateDate === 1 ? this.filteredTasks.sort(function (a, b) {
          return new Date(b['deactivateDate']) - new Date(a['deactivateDate'])
        }) : this.filteredTasks.sort(function (a, b) {
          return new Date(a['deactivateDate']) - new Date(b['deactivateDate'])
        })
      }
      // фильтрация по балансам
      if (this.filterItems.balance !== 0) {
        this.filterItems.balance === 1 ? this.filteredTasks.sort(function (a, b) {
          return new Date(b['balance']) - new Date(a['balance'])
        }) : this.filteredTasks.sort(function (a, b) {
          return new Date(a['balance']) - new Date(b['balance'])
        })
      }
    },
    resetFilterTasks() {
      this.background = "none"
      this.filterItems = {
        selectedUser: {code: 0, name: "", user: {}},
        CreatedAt: 1,
        deactivateDate: 0,
        balance: 0,
        actionTypes: [
          {
            name: "Card creation",
            code: "CardCreation"
          },
          {
            name: "Card replenishment",
            code: "CardReplenishment"
          },
          {
            name: "Account replenishment",
            code: "AccountReplenishment"
          }
        ],
        actionType: {
          name: null,
          code: null
        },
        statuses: [
          {
            name: "Active",
            code: 1
          },
          {
            name: "Сompleted",
            code: 0
          }
        ],
        status: {
          name: null,
          code: null
        },
        search: '',
      },
          this.filter()
    },
    getTasks() {
      if (this.loggedIn && this.users.length === 0 && this.req) {
        this.$store.dispatch("admin/getUsersMethod")
      }
      if (this.loggedIn && this.tasks.length === 0 && this.req) {
        this.$store.dispatch("admin/getTasksMethod");
      }
      this.req = false
      return true;
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

      if (timeZone && timeZone.utc.includes('-')) {
        let newDate = new Date(date);
        let utc = Number(timeZone.utc.split('-')[1])
        newDate.setTime(newDate.getTime() - (utc * 60 * 60 * 1000))
        return [newDate.toLocaleString().split(",")[0], newDate.toLocaleString().split(",")[1]]
      }
      if (timeZone && !timeZone.utc.includes('-')) {
        let newDate = new Date(date);
        let utc = Number(timeZone.utc)
        newDate.setTime(newDate.getTime() + (utc * 60 * 60 * 1000))
        return [newDate.toLocaleString().split(",")[0], newDate.toLocaleString().split(",")[1]]
      }
    },
  },
  computed: {
    ...mapGetters({
      users: "admin/getUsers",
      tasks: "admin/getTasks",
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
  },
  components: {
    TaskToCreateCard
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

.edit-task-limit {
  width: 100%;
  position: absolute;
  border-radius: 0.625rem;
  top: 10px;
  left: -40px;
  background: unset;
  z-index: 99;
}
</style>
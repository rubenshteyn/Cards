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
            <h2>Card creation log</h2>
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
<!--          <div class="modal-buttons modal-header border-bottom-0">-->
<!--            <button class="btn btn-primary">Export</button>-->
<!--            <button class="btn btn-dark" @click="closeModal">Exit</button>-->
<!--          </div>-->
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
                        <th class="min-w-150px">Card</th>
                        <th v-if="log.status.split('-')[0] !== log.status.split('-')[1]" class="min-w-150px">Status before</th>
                        <th v-if="log.status.split('-')[0] !== log.status.split('-')[1]" class="min-w-150px">Status after</th>
                        <th v-if="log.name.split('-')[0] !== log.name.split('-')[1]" class="min-w-150px">Card username before</th>
                        <th v-if="log.name.split('-')[0] !== log.name.split('-')[1]" class="min-w-150px">Card username after</th>
                        <th v-if="log.lastname.split('-')[0] !== log.lastname.split('-')[1]" class="min-w-150px">Card lastname before</th>
                        <th v-if="log.lastname.split('-')[0] !== log.lastname.split('-')[1]" class="min-w-150px">Card lastname after</th>
                        <th v-if="log.address.split('-')[0] !== log.address.split('-')[1]" class="min-w-150px">Card address before</th>
                        <th v-if="log.address.split('-')[0] !== log.address.split('-')[1]" class="min-w-150px">Card address after</th>
                        <th v-if="log.city.split('-')[0] !== log.city.split('-')[1]" class="min-w-150px">Card city before</th>
                        <th v-if="log.city.split('-')[0] !== log.city.split('-')[1]" class="min-w-150px">Card city after</th>
                        <th v-if="log.country.split('-')[0] !== log.country.split('-')[1]" class="min-w-150px">Card country before</th>
                        <th v-if="log.country.split('-')[0] !== log.country.split('-')[1]" class="min-w-150px">Card country after</th>
                        <th v-if="log.zip.split('-')[0] !== log.zip.split('-')[1]" class="min-w-150px">Card zip before</th>
                        <th v-if="log.zip.split('-')[0] !== log.zip.split('-')[1]" class="min-w-150px">Card zip after</th>
                        <th v-if="log.limit.split('-')[0] !== log.limit.split('-')[1]" class="min-w-150px">Card limit before</th>
                        <th v-if="log.limit.split('-')[0] !== log.limit.split('-')[1]" class="min-w-150px">Card limit after</th>
                      </tr>
                      </thead>
                      <!--end::Table head-->
                      <!--begin::Table body-->
                      <tbody v-if="this.log">
                      <tr>
                        <td class="fs-6 align-middle">
                          {{ this.log.cardID }}
                        </td>
                        <td v-if="this.log.status.split('-')[0] !== log.status.split('-')[1]" class="text-muted fw-semibold text-hover-primary fs-6 align-middle">
                            <span v-if="this.log.status.split('-')[0] === 0" class="badge badge-light-danger">
                              Blocked
                            </span>
                            <span v-if="this.log.status.split('-')[0] === 1" class="badge badge-light-success">
                              Active
                            </span>
                        </td>
                        <td v-if="this.log.status.split('-')[0] !== log.status.split('-')[1]" class="text-muted fw-semibold text-hover-primary fs-6 align-middle">
                            <span v-if="this.log.status.split('-')[1] === 0" class="badge badge-light-danger">
                              Blocked
                            </span>
                          <span v-if="this.log.status.split('-')[1] === 1" class="badge badge-light-success">
                              Active
                            </span>
                        </td>
                        <td v-if="this.log.name.split('-')[0] !== this.log.name.split('-')[1]" class="fs-6 align-middle">
                          {{ this.log.username.split('-')[0] }}
                        </td>
                        <td v-if="this.log.name.split('-')[0] !== this.log.name.split('-')[1]" class="fs-6 align-middle">
                          {{ this.log.username.split('-')[1] }}
                        </td>
                        <td v-if="this.log.lastname.split('-')[0] !== this.log.lastname.split('-')[1]" class="fs-6 align-middle">
                          {{ this.log.lastname.split('-')[0] }}
                        </td>
                        <td v-if="this.log.lastname.split('-')[0] !== this.log.lastname.split('-')[1]" class="fs-6 align-middle">
                          {{ this.log.lastname.split('-')[1] }}
                        </td>
                        <td v-if="this.log.address.split('-')[0] !== this.log.address.split('-')[1]" class="fs-6 align-middle">
                          {{ this.log.address.split('-')[0] }}
                        </td>
                        <td v-if="this.log.address.split('-')[0] !== this.log.address.split('-')[1]" class="fs-6 align-middle">
                          {{ this.log.address.split('-')[1] }}
                        </td>
                        <td v-if="this.log.city.split('-')[0] !== this.log.city.split('-')[1]" class="fs-6 align-middle">
                          {{ this.log.city.split('-')[0] }}
                        </td>
                        <td v-if="this.log.city.split('-')[0] !== this.log.city.split('-')[1]" class="fs-6 align-middle">
                          {{ this.log.city.split('-')[1] }}
                        </td>
                        <td v-if="this.log.country.split('-')[0] !== this.log.country.split('-')[1]" class="fs-6 align-middle">
                          {{ this.log.country.split('-')[0] }}
                        </td>
                        <td v-if="this.log.country.split('-')[0] !== this.log.country.split('-')[1]" class="fs-6 align-middle">
                          {{ this.log.country.split('-')[1] }}
                        </td>
                        <td v-if="this.log.zip.split('-')[0] !== this.log.zip.split('-')[1]" class="fs-6 align-middle">
                          {{ this.log.zip.split('-')[0] }}
                        </td>
                        <td v-if="this.log.zip.split('-')[0] !== this.log.zip.split('-')[1]" class="fs-6 align-middle">
                          {{ this.log.zip.split('-')[1] }}
                        </td>
                        <td v-if="this.log.limit.split('-')[0] !== this.log.limit.split('-')[1]" class="align-middle">
                           {{ this.log.limit.split('-')[0] }}
                        </td>
                        <td v-if="this.log.limit.split('-')[0] !== this.log.limit.split('-')[1]" class="align-middle">
                             {{ this.log.limit.split('-')[1] }}
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
import { mapGetters} from "vuex";
import apiV1 from "@/axios";

export default {
  name: "LogCardChange",
  data() {
    return {
      timezones: timezones,
      show: false,
      isShow: false,
      transform: "translate(0, -50px)",
      log: null,
      form: {
        actionID: null,
        type: null
      }
    }
  },
  computed: {
    ...mapGetters({
      currentUser: "auth/currentUser"
    }),
  },
  methods: {
    getFinLog(form) {
      this.$store.dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
          .then(() => {
            return apiV1.post("admin/getActionLog", form,
                {
                  headers: {
                    Authorization: this.getTokenHeader,
                    "Content-Type": "application/json"
                  }
                }
            )
                .then(response => {
                  if (!response.data.success) {
                    throw new Error(response.data.error);
                  }
                  console.log(response.data)
                  this.log = response.data.data
                  return response.data;
                })
                .catch(error => {
                  console.log(error)
                  throw error;
                });
          })
    },
    openModal(actionID, type) {
      this.form.actionID = actionID
      this.form.type = type
      this.getFinLog(this.form)
      this.show = true
      setTimeout(() => {
        this.isShow = this.show
        this.transform = "none"
      }, 200)
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
    logBalance(number) {
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
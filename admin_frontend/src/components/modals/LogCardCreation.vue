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
                    <table v-if="getData()" class="table table-row-bordered table-row-gray-100 align-middle gs-0 gy-3">
                      <!--begin::Table head-->
                      <thead class="rounded-left">
                      <tr class="fw-bold text-muted bg-light">
                        <th class="min-w-150px">Card</th>
                        <th class="min-w-150px">Balance</th>
                        <th class="min-w-150px">Account</th>
                        <th class="min-w-150px">Amount</th>
                      </tr>
                      </thead>
                      <!--end::Table head-->
                      <!--begin::Table body-->
                      <tbody v-if="this.log">
                      <tr>
                        <td class="text-dark fw-bold fs-6 align-middle">
                          {{ getCardData(this.log.cardID)[0] }}
                        </td>
                        <td class="text-dark fw-bold fs-6 align-middle">
                          {{ usersBalance(getCardData(this.log.cardID)[1]) + "$"}}
                        </td>
                        <td class="text-dark fw-bold fs-6 align-middle">
                          USDT
                        </td>
                        <td class="align-middle">
                          <div class="badge badge-light-danger"
                               style="width: max-content;">- {{ usersBalance(getCardData(this.log.cardID)[2])}}
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
import apiV1 from "@/axios";

export default {
  name: "LogCardCreation",
  data() {
    return {
      timezones: timezones,
      show: false,
      isShow: false,
      transform: "translate(0, -50px)",
      log: null,
      balanceName: null,
      form: {
        actionID: null,
        type: null,
      }
    }
  },
  computed: {
    ...mapGetters({
      currentUser: "auth/currentUser",
      cards: "admin/getCards",
      cardsLoaded: "admin/getCardsLoaded",
      getTokenHeader: "auth/getTokenHeader",
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
                  console.log(response.data.data)
                  this.log = response.data.data
                  return response.data;
                })
                .catch(error => {
                  console.log(error)
                  throw error;
                });
          })
    },
    getData() {
      if (!this.cardsLoaded && this.cards.length === 0) {
        this.$store.dispatch("admin/getCardsMethod");
      }
      return true
    },
    getCardData(ID) {
      if (!ID) {
        return "-"
      }
      let cards = Object.values(this.cards);
      let card = cards.find(item => item.ID === ID);
      console.log(card)
      if (card === undefined || card === "" || card === null || !ID) {
        return "-"
      }
      return [card.cardNumber, card.balance, card.balance]
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
    usersBalance(number) {
      if (number === 0 || !number) {
        return '0,00'
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
  }
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
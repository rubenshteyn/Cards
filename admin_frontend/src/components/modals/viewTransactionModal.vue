<template>
  <div v-if="show">
    <div :class="{show: this.isShow}" class="modal fade d-block" @click.self="closeModal" id="kt_modal_view_users" tabindex="-1"
         aria-hidden="true">
      <div class="modal-dialog modal-dialog-centered mw-850px" :style="{transform: this.transform}">
        <!--begin::Modal content-->
        <div class="modal-content">
          <!--begin::Modal header-->
          <div class="modal-header">
            <h2>Additional information</h2>
            <div @click="closeModal" class="btn btn-sm btn-icon btn-active-color-primary"
                 data-bs-dismiss="modal">
              <!--begin::Svg Icon | path: icons/duotune/arrows/arr061.svg-->
              <span class="svg-icon svg-icon-1">
                            <svg width="24" height="24" viewBox="0 0 24 24" fill="none"
                                 xmlns="http://www.w3.org/2000/svg">
                                <rect opacity="0.5" x="6" y="17.3137" width="16" height="2" rx="1"
                                      transform="rotate(-45 6 17.3137)" fill="currentColor" />
                                <rect x="7.41422" y="6" width="16" height="2" rx="1" transform="rotate(45 7.41422 6)"
                                      fill="currentColor" />
                            </svg>
                        </span>
              <!--end::Svg Icon-->
            </div>
          </div>
          <!--        <div class="modal-buttons modal-header border-bottom-0">-->
          <!--          <button class="btn btn-primary">Export</button>-->
          <!--          <button class="btn btn-dark" @click="closeModal">Exit</button>-->
          <!--        </div>-->
          <!--begin::Modal header-->
          <!--begin::Modal body-->
          <div class="modal-body scroll-y">
            <div class="mb-15">
              <!--begin::List-->
              <div class="mh-375px scroll-y me-n7 pe-7">
                <div class="card-body p-0">
                  <!--begin::Table container-->
                  <div class="table-responsive">
                    <!--begin::Table-->
                    <table class="table table-row-bordered table-row-gray-100 align-middle gs-0 gy-3">
                      <!--begin::Table head-->
                      <thead class="rounded-left">
                      <tr class="fw-bold text-muted bg-light">
                        <th class="min-w-150px">ID</th>
                        <th class="min-w-150px">Manager ID</th>
                        <th class="min-w-150px">Country</th>
                        <th class="min-w-150px">Balance up to</th>
                        <th class="min-w-150px">Merchant</th>
                      </tr>
                      </thead>
                      <!--end::Table head-->
                      <!--begin::Table body-->
                      <tbody>
                      <tr>
                        <td class="align-middle">
                        <span class="text-dark fw-bold d-block fs-6 mb-1">
                          {{ this.ID }}
                        </span>
                        </td>
                        <td class="align-middle">
                        <span class="text-dark fw-bold d-block fs-6 mb-1">
                          {{ this.managerId }}
                        </span>
                        </td>
                        <td class="align-middle">
                        <span class="text-dark fw-bold d-block fs-6 mb-1">
                          {{ this.country }}
                        </span>
                        </td>
                        <td class="align-middle">
                        <span class="text-dark fw-bold d-block fs-6 mb-1">
                          {{ transactionBalance(this.initial) + "$"}}
                        </span>
                        </td>
                        <td class="align-middle">
                        <span class="text-dark fw-bold d-block fs-6 mb-1">
                          {{ this.merchant }}
                        </span>
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
export default {
  name: "viewTransactionModal",
  data() {
    return {
      show: false,
      isShow: false,
      transform: "translate(0, -50px)",
      country: null,
      initial: null,
      merchant: null,
      ID: null,
      managerId: null
    }
  },
  methods: {
    openModal(transaction) {
      this.show = true
      setTimeout(()=> {
        this.isShow = this.show
        this.transform = "none"
      },200)
      this.ID = transaction.ID
      this.managerId = transaction.managerId
      this.country = transaction.country
      this.initial = transaction.initial
      this.merchant = transaction.merchant
    },
    closeModal: function () {
      this.isShow = false
      this.transform = "translate(0, -50px)"
      setTimeout(()=> {
        this.show = false
      },200)
    },
    transactionBalance(number) {
      if (number === 0) {
        return '0000.00'
      }

      let result = []
      let arrayBalance = (""+number).split("").map(Number)
      for (const i of arrayBalance) {
        result.push(i)
      }

      function roundingNum(x,n) {
        return parseFloat(Number.parseFloat(x).toFixed(n || 2));
      }
      let arrRoundingNum = Array.from(roundingNum(number, 2).toString())
      let endResult = Array.from(arrRoundingNum.toString().split("."))[1]
      let firstResult = Array.from(arrRoundingNum.toString().split(".")[0])
      if(firstResult.toString().replace(/[\s.,%]/g, '').split('').length > 3) {
        arrRoundingNum.splice(0, 1)
        return result.shift() + ',' + arrRoundingNum.join("")
      }
      if(typeof endResult == 'undefined') {
        return arrRoundingNum.join("") + ".00"
      }
      if(endResult.split('').length < 2) {
        return arrRoundingNum.join("") + ".0"
      }

      return arrRoundingNum.join("")
    },
  }
}
</script>

<style scoped>

</style>
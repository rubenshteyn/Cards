<template>
  <div v-if="show">
    <div :class="{show: this.isShow}" class="modal fade d-block" @click.self="closeModal" id="kt_modal_view_users" tabindex="-1"
         aria-hidden="true">
      <div class="modal-dialog modal-dialog-centered mw-850px" :style="{transform: this.transform}">
        <!--begin::Modal content-->
        <div class="modal-content">
          <!--begin::Modal header-->
          <div class="modal-header">
            <h2>Card details</h2>
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
                        <th class="min-w-150px">Full name</th>
                        <th class="min-w-150px">Address</th>
                      </tr>
                      </thead>
                      <!--end::Table head-->
                      <!--begin::Table body-->
                      <tbody>
                      <tr>
                        <td class="align-middle">
                          <span class="text-dark fw-bold fs-6">{{this.fullName}}</span>
                        </td>
                        <td class="align-middle">
                          <span class="text-dark fw-bold fs-6">{{this.address}}</span>
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
  name: "detailsCardModal",
  data() {
    return {
      show: false,
      isShow: false,
      transform: "translate(0, -50px)",
      fullName: null,
      address: null,
    }
  },
  methods: {
    openModal(card) {
      this.show = true
      setTimeout(()=> {
        this.isShow = this.show
        this.transform = "none"
      },200)
      this.fullName = card.name
      this.address = card.address
    },
    closeModal: function () {
      this.isShow = false
      this.transform = "translate(0, -50px)"
      setTimeout(()=> {
        this.show = false
      },200)
    },
  }
}
</script>

<style scoped>

</style>